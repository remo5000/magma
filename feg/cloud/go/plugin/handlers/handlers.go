/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package handlers

import (
	"fmt"
	"net/http"

	"magma/feg/cloud/go/feg"
	fegModels "magma/feg/cloud/go/plugin/models"
	"magma/feg/cloud/go/services/health"
	lteModels "magma/lte/cloud/go/plugin/models"
	merrors "magma/orc8r/cloud/go/errors"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/pluginimpl/handlers"
	orc8rModels "magma/orc8r/cloud/go/pluginimpl/models"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/storage"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const (
	FederationNetworks             = "feg"
	ListFegNetworksPath            = obsidian.V1Root + FederationNetworks
	ManageFegNetworkPath           = ListFegNetworksPath + "/:network_id"
	ManageFegNetworkFederationPath = ManageFegNetworkPath + obsidian.UrlSep + "federation"
	ManageFegNetworkSubscriberPath = ManageFegNetworkPath + obsidian.UrlSep + "subscriber_config"
	ManageFegNetworkBaseNamesPath  = ManageFegNetworkSubscriberPath + obsidian.UrlSep + "base_names"
	ManageFegNetworkRuleNamesPath  = ManageFegNetworkSubscriberPath + obsidian.UrlSep + "rule_names"
	ManageNetworkClusterStatusPath = ManageFegNetworkPath + obsidian.UrlSep + "cluster_status"

	Gateways                      = "gateways"
	ListGatewaysPath              = ManageFegNetworkPath + obsidian.UrlSep + Gateways
	ManageGatewayPath             = ListGatewaysPath + obsidian.UrlSep + ":gateway_id"
	ManageGatewayStatePath        = ManageGatewayPath + obsidian.UrlSep + "status"
	ManageGatewayFederationPath   = ManageGatewayPath + obsidian.UrlSep + "federation"
	ManageGatewayHealthStatusPath = ManageGatewayPath + obsidian.UrlSep + "health_status"

	FederatedLteNetworks              = "feg_lte"
	ListFegLteNetworksPath            = obsidian.V1Root + FederatedLteNetworks
	ManageFegLteNetworkPath           = ListFegLteNetworksPath + "/:network_id"
	ManageFegLteNetworkFederationPath = ManageFegLteNetworkPath + obsidian.UrlSep + "federation"
	ManageFegLteNetworkSubscriberPath = ManageFegLteNetworkPath + obsidian.UrlSep + "subscriber_config"
	ManageFegLteNetworkBaseNamesPath  = ManageFegLteNetworkSubscriberPath + obsidian.UrlSep + "base_names"
	ManageFegLteNetworkRuleNamesPath  = ManageFegLteNetworkSubscriberPath + obsidian.UrlSep + "rule_names"
)

func GetHandlers() []obsidian.Handler {
	ret := []obsidian.Handler{
		handlers.GetListGatewaysHandler(ListGatewaysPath, feg.FegGatewayType, makeFederationGateways),
		{Path: ListGatewaysPath, Methods: obsidian.POST, HandlerFunc: createGateway},
		{Path: ManageGatewayPath, Methods: obsidian.GET, HandlerFunc: getGateway},
		{Path: ManageGatewayPath, Methods: obsidian.PUT, HandlerFunc: updateGateway},
		handlers.GetDeleteGatewayHandler(ManageGatewayPath, feg.FegGatewayType),

		{Path: ManageGatewayStatePath, Methods: obsidian.GET, HandlerFunc: handlers.GetStateHandler},
		{Path: ManageNetworkClusterStatusPath, Methods: obsidian.GET, HandlerFunc: getClusterStatusHandler},
		{Path: ManageGatewayHealthStatusPath, Methods: obsidian.GET, HandlerFunc: getHealthStatusHandler},
	}

	ret = append(ret, handlers.GetTypedNetworkCRUDHandlers(ListFegNetworksPath, ManageFegNetworkPath, feg.FederationNetworkType, &fegModels.FegNetwork{})...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageFegNetworkFederationPath, &fegModels.NetworkFederationConfigs{}, "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageFegNetworkSubscriberPath, &lteModels.NetworkSubscriberConfig{}, "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageFegNetworkRuleNamesPath, new(lteModels.RuleNames), "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageFegNetworkBaseNamesPath, new(lteModels.BaseNames), "")...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayFederationPath, &fegModels.GatewayFederationConfigs{})...)

	ret = append(ret, handlers.GetTypedNetworkCRUDHandlers(ListFegLteNetworksPath, ManageFegLteNetworkPath, feg.FederatedLteNetworkType, &fegModels.FegLteNetwork{})...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageFegLteNetworkFederationPath, &fegModels.FederatedNetworkConfigs{}, "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageFegLteNetworkSubscriberPath, &lteModels.NetworkSubscriberConfig{}, "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageFegLteNetworkRuleNamesPath, new(lteModels.RuleNames), "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageFegLteNetworkBaseNamesPath, new(lteModels.BaseNames), "")...)

	return ret
}

func createGateway(c echo.Context) error {
	if nerr := handlers.CreateMagmadGatewayFromModel(c, &fegModels.MutableFederationGateway{}); nerr != nil {
		return nerr
	}
	return c.NoContent(http.StatusCreated)
}

func getGateway(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	magmadModel, nerr := handlers.LoadMagmadGatewayModel(nid, gid)
	if nerr != nil {
		return nerr
	}

	ent, err := configurator.LoadEntity(
		nid, feg.FegGatewayType, gid,
		configurator.EntityLoadCriteria{LoadConfig: true, LoadAssocsFromThis: true},
	)
	if err != nil {
		return obsidian.HttpError(errors.Wrap(err, "failed to load federation gateway"), http.StatusInternalServerError)
	}

	ret := &fegModels.FederationGateway{
		ID:          magmadModel.ID,
		Name:        magmadModel.Name,
		Description: magmadModel.Description,
		Device:      magmadModel.Device,
		Status:      magmadModel.Status,
		Tier:        magmadModel.Tier,
		Magmad:      magmadModel.Magmad,
		Federation:  ent.Config.(*fegModels.GatewayFederationConfigs),
	}
	return c.JSON(http.StatusOK, ret)
}

func updateGateway(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}
	if nerr = handlers.UpdateMagmadGatewayFromModel(c, nid, gid, &fegModels.MutableFederationGateway{}); nerr != nil {
		return nerr
	}
	return c.NoContent(http.StatusNoContent)
}

type federationAndMagmadGateway struct {
	magmadGateway, federationGateway configurator.NetworkEntity
}

func makeFederationGateways(
	entsByTK map[storage.TypeAndKey]configurator.NetworkEntity,
	devicesByID map[string]interface{},
	statusesByID map[string]*orc8rModels.GatewayStatus,
) map[string]handlers.GatewayModel {
	gatewayEntsByKey := map[string]*federationAndMagmadGateway{}
	for tk, ent := range entsByTK {
		existing, found := gatewayEntsByKey[tk.Key]
		if !found {
			existing = &federationAndMagmadGateway{}
			gatewayEntsByKey[tk.Key] = existing
		}

		switch ent.Type {
		case orc8r.MagmadGatewayType:
			existing.magmadGateway = ent
		case feg.FegGatewayType:
			existing.federationGateway = ent
		}
	}

	ret := make(map[string]handlers.GatewayModel, len(gatewayEntsByKey))
	for key, ents := range gatewayEntsByKey {
		hwID := ents.magmadGateway.PhysicalID
		var devCasted *orc8rModels.GatewayDevice
		if devicesByID[hwID] != nil {
			devCasted = devicesByID[hwID].(*orc8rModels.GatewayDevice)
		}
		ret[key] = (&fegModels.FederationGateway{}).FromBackendModels(ents.magmadGateway, ents.federationGateway, devCasted, statusesByID[hwID])
	}
	return ret
}

func getClusterStatusHandler(c echo.Context) error {
	nid, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	network, err := configurator.LoadNetwork(nid, true, true)
	if err == merrors.ErrNotFound {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	if network.Type != feg.FederationNetworkType {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("network %s is not a <%s> network", nid, feg.FederationNetworkType))
	}
	activeGw, err := health.GetActiveGateway(nid)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	ret := &fegModels.FederationNetworkClusterStatus{
		ActiveGateway: activeGw,
	}
	return c.JSON(http.StatusOK, ret)
}

func getHealthStatusHandler(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}
	pid, err := configurator.GetPhysicalIDOfEntity(nid, orc8r.MagmadGatewayType, gid)
	if err == merrors.ErrNotFound || len(pid) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	res, err := health.GetHealth(nid, gid)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	ret := &fegModels.FederationGatewayHealthStatus{
		Status:      res.GetHealth().GetHealth().String(),
		Description: res.GetHealth().GetHealthMessage(),
	}
	return c.JSON(http.StatusOK, ret)
}
