// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceBGP struct {
	Id                                  types.String                      `tfsdk:"id"`
	Domain                              types.String                      `tfsdk:"domain"`
	DeviceId                            types.String                      `tfsdk:"device_id"`
	Name                                types.String                      `tfsdk:"name"`
	Type                                types.String                      `tfsdk:"type"`
	AsNumber                            types.String                      `tfsdk:"as_number"`
	Ipv4AddressFamilyType               types.String                      `tfsdk:"ipv4_address_family_type"`
	Ipv4LearnedRouteMapId               types.String                      `tfsdk:"ipv4_learned_route_map_id"`
	Ipv4DefaultInformationOrginate      types.Bool                        `tfsdk:"ipv4_default_information_orginate"`
	Ipv4AutoAummary                     types.Bool                        `tfsdk:"ipv4_auto_aummary"`
	Ipv4BgpSupressInactive              types.Bool                        `tfsdk:"ipv4_bgp_supress_inactive"`
	Ipv4Synchronization                 types.Bool                        `tfsdk:"ipv4_synchronization"`
	Ipv4BgpRedistributeInternal         types.Bool                        `tfsdk:"ipv4_bgp_redistribute_internal"`
	Ipv4ExternalDistance                types.Int64                       `tfsdk:"ipv4_external_distance"`
	Ipv4InternalDistance                types.Int64                       `tfsdk:"ipv4_internal_distance"`
	Ipv4LocalDistance                   types.Int64                       `tfsdk:"ipv4_local_distance"`
	Ipv4ForwardPacketsOverMultipathIbgp types.Int64                       `tfsdk:"ipv4_forward_packets_over_multipath_ibgp"`
	Ipv4ForwardPacketsOverMultipathEbgp types.Int64                       `tfsdk:"ipv4_forward_packets_over_multipath_ebgp"`
	Ipv4Neighbors                       []DeviceBGPIpv4Neighbors          `tfsdk:"ipv4_neighbors"`
	Ipv4AggregateAddresses              []DeviceBGPIpv4AggregateAddresses `tfsdk:"ipv4_aggregate_addresses"`
	Ipv4Filterings                      []DeviceBGPIpv4Filterings         `tfsdk:"ipv4_filterings"`
	Ipv4Networks                        []DeviceBGPIpv4Networks           `tfsdk:"ipv4_networks"`
	Ipv4Redistributions                 []DeviceBGPIpv4Redistributions    `tfsdk:"ipv4_redistributions"`
	Ipv4RouteInjections                 []DeviceBGPIpv4RouteInjections    `tfsdk:"ipv4_route_injections"`
}

type DeviceBGPIpv4Neighbors struct {
	NeighborAddress                           types.String                                        `tfsdk:"neighbor_address"`
	NeighborRemoteAs                          types.String                                        `tfsdk:"neighbor_remote_as"`
	NeighborBfd                               types.String                                        `tfsdk:"neighbor_bfd"`
	UpdateSourceInterfaceId                   types.String                                        `tfsdk:"update_source_interface_id"`
	EnableAddressFamily                       types.Bool                                          `tfsdk:"enable_address_family"`
	NeighborShutdown                          types.Bool                                          `tfsdk:"neighbor_shutdown"`
	NeighborDescription                       types.String                                        `tfsdk:"neighbor_description"`
	NeighborFilterAccessLists                 []DeviceBGPIpv4NeighborsNeighborFilterAccessLists   `tfsdk:"neighbor_filter_access_lists"`
	NeighborFilterRouteMapLists               []DeviceBGPIpv4NeighborsNeighborFilterRouteMapLists `tfsdk:"neighbor_filter_route_map_lists"`
	NeighborFilterPrefixLists                 []DeviceBGPIpv4NeighborsNeighborFilterPrefixLists   `tfsdk:"neighbor_filter_prefix_lists"`
	NeighborFilterAsPathLists                 []DeviceBGPIpv4NeighborsNeighborFilterAsPathLists   `tfsdk:"neighbor_filter_as_path_lists"`
	NeighborFilterMaxPrefix                   types.Int64                                         `tfsdk:"neighbor_filter_max_prefix"`
	NeighborFilterWarningOnly                 types.Bool                                          `tfsdk:"neighbor_filter_warning_only"`
	NeighborFilterThresholdValue              types.Int64                                         `tfsdk:"neighbor_filter_threshold_value"`
	NeighborFilterRestartInterval             types.Int64                                         `tfsdk:"neighbor_filter_restart_interval"`
	NeighborRoutesAdvertisementInterval       types.Int64                                         `tfsdk:"neighbor_routes_advertisement_interval"`
	NeighborRoutesRemovePrivateAs             types.Bool                                          `tfsdk:"neighbor_routes_remove_private_as"`
	NeighborGenerateDefaultRouteMapId         types.String                                        `tfsdk:"neighbor_generate_default_route_map_id"`
	NeighborRoutesAdvertiseMapUseExist        types.Bool                                          `tfsdk:"neighbor_routes_advertise_map_use_exist"`
	NeighborRoutesAdvertiseMapId              types.String                                        `tfsdk:"neighbor_routes_advertise_map_id"`
	NeighborRoutesAdvertiseExistNonexistMapId types.String                                        `tfsdk:"neighbor_routes_advertise_exist_nonexist_map_id"`
	NeighborKeepaliveInterval                 types.Int64                                         `tfsdk:"neighbor_keepalive_interval"`
	NeighborHoldTime                          types.Int64                                         `tfsdk:"neighbor_hold_time"`
	NeighborMinHoldTime                       types.Int64                                         `tfsdk:"neighbor_min_hold_time"`
	NeighborAuthenticationPassword            types.String                                        `tfsdk:"neighbor_authentication_password"`
	NeighborSendCommunityAttribute            types.Bool                                          `tfsdk:"neighbor_send_community_attribute"`
	NeighborNexthopSelf                       types.Bool                                          `tfsdk:"neighbor_nexthop_self"`
	NeighborDisableConnectionVerification     types.Bool                                          `tfsdk:"neighbor_disable_connection_verification"`
	NeighborTcpMtuPathDiscovery               types.Bool                                          `tfsdk:"neighbor_tcp_mtu_path_discovery"`
	NeighborMaxHopCount                       types.Int64                                         `tfsdk:"neighbor_max_hop_count"`
	NeighborTcpTransportMode                  types.Bool                                          `tfsdk:"neighbor_tcp_transport_mode"`
	NeighborWeight                            types.Int64                                         `tfsdk:"neighbor_weight"`
	NeighborVersion                           types.String                                        `tfsdk:"neighbor_version"`
	NeighborCustomizedLocalAsNumber           types.String                                        `tfsdk:"neighbor_customized_local_as_number"`
	NeighborCustomizedNoPrepend               types.Bool                                          `tfsdk:"neighbor_customized_no_prepend"`
	NeighborCustomizedReplaceAs               types.Bool                                          `tfsdk:"neighbor_customized_replace_as"`
	NeighborCustomizedAcceptBothAs            types.Bool                                          `tfsdk:"neighbor_customized_accept_both_as"`
}

type DeviceBGPIpv4AggregateAddresses struct {
	GenerateAs     types.Bool   `tfsdk:"generate_as"`
	Filter         types.Bool   `tfsdk:"filter"`
	NetworkId      types.String `tfsdk:"network_id"`
	AdvertiseMapId types.String `tfsdk:"advertise_map_id"`
	AttributeMapId types.String `tfsdk:"attribute_map_id"`
	SuppressMapId  types.String `tfsdk:"suppress_map_id"`
}

type DeviceBGPIpv4Filterings struct {
	AccessListId     types.String `tfsdk:"access_list_id"`
	NetworkDirection types.String `tfsdk:"network_direction"`
	Protocol         types.String `tfsdk:"protocol"`
	ProrocolProcess  types.String `tfsdk:"prorocol_process"`
}

type DeviceBGPIpv4Networks struct {
	NetworkId  types.String `tfsdk:"network_id"`
	RouteMapId types.String `tfsdk:"route_map_id"`
}

type DeviceBGPIpv4Redistributions struct {
	SourceProtocol     types.String `tfsdk:"source_protocol"`
	RouteMapId         types.String `tfsdk:"route_map_id"`
	Metric             types.Int64  `tfsdk:"metric"`
	ProcessId          types.String `tfsdk:"process_id"`
	MatchExternal1     types.Bool   `tfsdk:"match_external1"`
	MatchExternal2     types.Bool   `tfsdk:"match_external2"`
	MatchInternal      types.Bool   `tfsdk:"match_internal"`
	MatchNssaExternal1 types.Bool   `tfsdk:"match_nssa_external1"`
	MatchNssaExternal2 types.Bool   `tfsdk:"match_nssa_external2"`
}

type DeviceBGPIpv4RouteInjections struct {
	InjectRouteMapId types.String `tfsdk:"inject_route_map_id"`
	ExistRouteMapId  types.String `tfsdk:"exist_route_map_id"`
}

type DeviceBGPIpv4NeighborsNeighborFilterAccessLists struct {
	AccessListId    types.String `tfsdk:"access_list_id"`
	UpdateDirection types.String `tfsdk:"update_direction"`
}
type DeviceBGPIpv4NeighborsNeighborFilterRouteMapLists struct {
	RouteMapId      types.String `tfsdk:"route_map_id"`
	UpdateDirection types.String `tfsdk:"update_direction"`
}
type DeviceBGPIpv4NeighborsNeighborFilterPrefixLists struct {
	PrefixListId    types.String `tfsdk:"prefix_list_id"`
	UpdateDirection types.String `tfsdk:"update_direction"`
}
type DeviceBGPIpv4NeighborsNeighborFilterAsPathLists struct {
	UpdateDirection types.String `tfsdk:"update_direction"`
	AsPathId        types.String `tfsdk:"as_path_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceBGP) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/bgp", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceBGP) toBody(ctx context.Context, state DeviceBGP) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Ipv4LearnedRouteMapId.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.aftableMap.id", data.Ipv4LearnedRouteMapId.ValueString())
	}
	if !data.Ipv4DefaultInformationOrginate.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.defaultInformationOrginate", data.Ipv4DefaultInformationOrginate.ValueBool())
	}
	if !data.Ipv4AutoAummary.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.autoSummary", data.Ipv4AutoAummary.ValueBool())
	}
	if !data.Ipv4BgpSupressInactive.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.bgpSupressInactive", data.Ipv4BgpSupressInactive.ValueBool())
	}
	if !data.Ipv4Synchronization.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.synchronization", data.Ipv4Synchronization.ValueBool())
	}
	if !data.Ipv4BgpRedistributeInternal.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.bgpRedistributeInternal", data.Ipv4BgpRedistributeInternal.ValueBool())
	}
	if !data.Ipv4ExternalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.externalDistance", data.Ipv4ExternalDistance.ValueInt64())
	}
	if !data.Ipv4InternalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.internalDistance", data.Ipv4InternalDistance.ValueInt64())
	}
	if !data.Ipv4LocalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.localDistance", data.Ipv4LocalDistance.ValueInt64())
	}
	if !data.Ipv4ForwardPacketsOverMultipathIbgp.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.ibgp", data.Ipv4ForwardPacketsOverMultipathIbgp.ValueInt64())
	}
	if !data.Ipv4ForwardPacketsOverMultipathEbgp.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.ebgp", data.Ipv4ForwardPacketsOverMultipathEbgp.ValueInt64())
	}
	if len(data.Ipv4Neighbors) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.neighbors", []interface{}{})
		for _, item := range data.Ipv4Neighbors {
			itemBody := ""
			if !item.NeighborAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4Address", item.NeighborAddress.ValueString())
			}
			if !item.NeighborRemoteAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "remoteAs", item.NeighborRemoteAs.ValueString())
			}
			if !item.NeighborBfd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.fallOverBFD", item.NeighborBfd.ValueString())
			}
			if !item.UpdateSourceInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.updateSource.id", item.UpdateSourceInterfaceId.ValueString())
			}
			if !item.EnableAddressFamily.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.enableAddress", item.EnableAddressFamily.ValueBool())
			}
			if !item.NeighborShutdown.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.shutdown", item.NeighborShutdown.ValueBool())
			}
			if !item.NeighborDescription.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.description", item.NeighborDescription.ValueString())
			}
			if len(item.NeighborFilterAccessLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborDistributeLists", []interface{}{})
				for _, childItem := range item.NeighborFilterAccessLists {
					itemChildBody := ""
					if !childItem.AccessListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "accessList.id", childItem.AccessListId.ValueString())
					}
					if !childItem.UpdateDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterUpdateAction", childItem.UpdateDirection.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborFiltering.neighborDistributeLists.-1", itemChildBody)
				}
			}
			if len(item.NeighborFilterRouteMapLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborRouteMap", []interface{}{})
				for _, childItem := range item.NeighborFilterRouteMapLists {
					itemChildBody := ""
					if !childItem.RouteMapId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "RouteMap.id", childItem.RouteMapId.ValueString())
					}
					if !childItem.UpdateDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterUpdateAction", childItem.UpdateDirection.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborFiltering.neighborRouteMap.-1", itemChildBody)
				}
			}
			if len(item.NeighborFilterPrefixLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.ipv4PrefixListFilter", []interface{}{})
				for _, childItem := range item.NeighborFilterPrefixLists {
					itemChildBody := ""
					if !childItem.PrefixListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ipv4PrefixList.id", childItem.PrefixListId.ValueString())
					}
					if !childItem.UpdateDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterUpdateAction", childItem.UpdateDirection.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborFiltering.ipv4PrefixListFilter.-1", itemChildBody)
				}
			}
			if len(item.NeighborFilterAsPathLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborFilterList.neighborFilterList", []interface{}{})
				for _, childItem := range item.NeighborFilterAsPathLists {
					itemChildBody := ""
					if !childItem.UpdateDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterUpdateAction", childItem.UpdateDirection.ValueString())
					}
					if !childItem.AsPathId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "asPathList.id", childItem.AsPathId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborFilterList.neighborFilterList.-1", itemChildBody)
				}
			}
			if !item.NeighborFilterMaxPrefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborMaximumPrefix.maxPrefixLimit", item.NeighborFilterMaxPrefix.ValueInt64())
			}
			if !item.NeighborFilterWarningOnly.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborMaximumPrefix.warningOnly", item.NeighborFilterWarningOnly.ValueBool())
			}
			if !item.NeighborFilterThresholdValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborMaximumPrefix.thresholdValue", item.NeighborFilterThresholdValue.ValueInt64())
			}
			if !item.NeighborFilterRestartInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborMaximumPrefix.restartInterval", item.NeighborFilterRestartInterval.ValueInt64())
			}
			if !item.NeighborRoutesAdvertisementInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborRoutes.advertisementInterval", item.NeighborRoutesAdvertisementInterval.ValueInt64())
			}
			if !item.NeighborRoutesRemovePrivateAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborRoutes.removePrivateAs", item.NeighborRoutesRemovePrivateAs.ValueBool())
			}
			if !item.NeighborGenerateDefaultRouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborDefaultOriginate.routeMap.id", item.NeighborGenerateDefaultRouteMapId.ValueString())
			}
			if !item.NeighborRoutesAdvertiseMapUseExist.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborRoutes.neighborAdvertiseMaps.existMap", item.NeighborRoutesAdvertiseMapUseExist.ValueBool())
			}
			if !item.NeighborRoutesAdvertiseMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborRoutes.neighborAdvertiseMaps.routeMap.id", item.NeighborRoutesAdvertiseMapId.ValueString())
			}
			if !item.NeighborRoutesAdvertiseExistNonexistMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborRoutes.neighborAdvertiseMaps.existRouteMap.id", item.NeighborRoutesAdvertiseExistNonexistMapId.ValueString())
			}
			if !item.NeighborKeepaliveInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborTimers.keepAliveInterval", item.NeighborKeepaliveInterval.ValueInt64())
			}
			if !item.NeighborHoldTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborTimers.holdTime", item.NeighborHoldTime.ValueInt64())
			}
			if !item.NeighborMinHoldTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborTimers.minimumHoldTime", item.NeighborMinHoldTime.ValueInt64())
			}
			if !item.NeighborAuthenticationPassword.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborSecret", item.NeighborAuthenticationPassword.ValueString())
			}
			if !item.NeighborSendCommunityAttribute.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.sendCommunity", item.NeighborSendCommunityAttribute.ValueBool())
			}
			if !item.NeighborNexthopSelf.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.nextHopSelf", item.NeighborNexthopSelf.ValueBool())
			}
			if !item.NeighborDisableConnectionVerification.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborHops.disableConnectedCheck", item.NeighborDisableConnectionVerification.ValueBool())
			}
			if !item.NeighborTcpMtuPathDiscovery.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborTransportPathMTUDiscovery.disable", item.NeighborTcpMtuPathDiscovery.ValueBool())
			}
			if !item.NeighborMaxHopCount.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborHops.maxHopCount", item.NeighborMaxHopCount.ValueInt64())
			}
			if !item.NeighborTcpTransportMode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborTransportConnectionMode.establishTCPSession", item.NeighborTcpTransportMode.ValueBool())
			}
			if !item.NeighborWeight.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborWeight", item.NeighborWeight.ValueInt64())
			}
			if !item.NeighborVersion.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborVersion", item.NeighborVersion.ValueString())
			}
			if !item.NeighborCustomizedLocalAsNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborLocalAs.asNumber", item.NeighborCustomizedLocalAsNumber.ValueString())
			}
			if !item.NeighborCustomizedNoPrepend.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborLocalAs.noPrepend", item.NeighborCustomizedNoPrepend.ValueBool())
			}
			if !item.NeighborCustomizedReplaceAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborLocalAs.replaceAs", item.NeighborCustomizedReplaceAs.ValueBool())
			}
			if !item.NeighborCustomizedAcceptBothAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborLocalAs.dualAs", item.NeighborCustomizedAcceptBothAs.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.neighbors.-1", itemBody)
		}
	}
	if len(data.Ipv4AggregateAddresses) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.aggregateAddressesIPv4s", []interface{}{})
		for _, item := range data.Ipv4AggregateAddresses {
			itemBody := ""
			if !item.GenerateAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "asSet", item.GenerateAs.ValueBool())
			}
			if !item.Filter.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "summaryOnly", item.Filter.ValueBool())
			}
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4AggregateNetwork.id", item.NetworkId.ValueString())
			}
			if !item.AdvertiseMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "advertiseMap.id", item.AdvertiseMapId.ValueString())
			}
			if !item.AttributeMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "attributeMap.id", item.AttributeMapId.ValueString())
			}
			if !item.SuppressMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "suppressMap.id", item.SuppressMapId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.aggregateAddressesIPv4s.-1", itemBody)
		}
	}
	if len(data.Ipv4Filterings) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distributeLists", []interface{}{})
		for _, item := range data.Ipv4Filterings {
			itemBody := ""
			if !item.AccessListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4AggregateNetwork.id", item.AccessListId.ValueString())
			}
			if !item.NetworkDirection.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.NetworkDirection.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol.protocol", item.Protocol.ValueString())
			}
			if !item.ProrocolProcess.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol.processId", item.ProrocolProcess.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.distributeLists.-1", itemBody)
		}
	}
	if len(data.Ipv4Networks) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.networks", []interface{}{})
		for _, item := range data.Ipv4Networks {
			itemBody := ""
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4Address.id", item.NetworkId.ValueString())
			}
			if !item.RouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMap.id", item.RouteMapId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.networks.-1", itemBody)
		}
	}
	if len(data.Ipv4Redistributions) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.redistributeProtocols", []interface{}{})
		for _, item := range data.Ipv4Redistributions {
			itemBody := ""
			if !item.SourceProtocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.SourceProtocol.ValueString())
			}
			if !item.RouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMap.id", item.RouteMapId.ValueString())
			}
			if !item.Metric.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.metricValue", item.Metric.ValueInt64())
			}
			if !item.ProcessId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "processId", item.ProcessId.ValueString())
			}
			if !item.MatchExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchExternal1", item.MatchExternal1.ValueBool())
			}
			if !item.MatchExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchExternal2", item.MatchExternal2.ValueBool())
			}
			if !item.MatchInternal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchInternal", item.MatchInternal.ValueBool())
			}
			if !item.MatchNssaExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchNssaExternal1", item.MatchNssaExternal1.ValueBool())
			}
			if !item.MatchNssaExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchNssaExternal2", item.MatchNssaExternal2.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.redistributeProtocols.-1", itemBody)
		}
	}
	if len(data.Ipv4RouteInjections) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.injectMaps", []interface{}{})
		for _, item := range data.Ipv4RouteInjections {
			itemBody := ""
			if !item.InjectRouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "injectMap.routeMap.id", item.InjectRouteMapId.ValueString())
			}
			if !item.ExistRouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "existMap.routeMap.id", item.ExistRouteMapId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.injectMaps.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceBGP) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("asNumber"); value.Exists() {
		data.AsNumber = types.StringValue(value.String())
	} else {
		data.AsNumber = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.type"); value.Exists() {
		data.Ipv4AddressFamilyType = types.StringValue(value.String())
	} else {
		data.Ipv4AddressFamilyType = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.aftableMap.id"); value.Exists() {
		data.Ipv4LearnedRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4LearnedRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.defaultInformationOrginate"); value.Exists() {
		data.Ipv4DefaultInformationOrginate = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DefaultInformationOrginate = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.autoSummary"); value.Exists() {
		data.Ipv4AutoAummary = types.BoolValue(value.Bool())
	} else {
		data.Ipv4AutoAummary = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpSupressInactive"); value.Exists() {
		data.Ipv4BgpSupressInactive = types.BoolValue(value.Bool())
	} else {
		data.Ipv4BgpSupressInactive = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.synchronization"); value.Exists() {
		data.Ipv4Synchronization = types.BoolValue(value.Bool())
	} else {
		data.Ipv4Synchronization = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpRedistributeInternal"); value.Exists() {
		data.Ipv4BgpRedistributeInternal = types.BoolValue(value.Bool())
	} else {
		data.Ipv4BgpRedistributeInternal = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.distance.externalDistance"); value.Exists() {
		data.Ipv4ExternalDistance = types.Int64Value(value.Int())
	} else {
		data.Ipv4ExternalDistance = types.Int64Value(20)
	}
	if value := res.Get("addressFamilyIPv4.distance.internalDistance"); value.Exists() {
		data.Ipv4InternalDistance = types.Int64Value(value.Int())
	} else {
		data.Ipv4InternalDistance = types.Int64Value(200)
	}
	if value := res.Get("addressFamilyIPv4.distance.localDistance"); value.Exists() {
		data.Ipv4LocalDistance = types.Int64Value(value.Int())
	} else {
		data.Ipv4LocalDistance = types.Int64Value(200)
	}
	if value := res.Get("addressFamilyIPv4.ibgp"); value.Exists() {
		data.Ipv4ForwardPacketsOverMultipathIbgp = types.Int64Value(value.Int())
	} else {
		data.Ipv4ForwardPacketsOverMultipathIbgp = types.Int64Value(1)
	}
	if value := res.Get("addressFamilyIPv4.ebgp"); value.Exists() {
		data.Ipv4ForwardPacketsOverMultipathEbgp = types.Int64Value(value.Int())
	} else {
		data.Ipv4ForwardPacketsOverMultipathEbgp = types.Int64Value(1)
	}
	if value := res.Get("addressFamilyIPv4.neighbors"); value.Exists() {
		data.Ipv4Neighbors = make([]DeviceBGPIpv4Neighbors, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4Neighbors{}
			if value := res.Get("ipv4Address"); value.Exists() {
				data.NeighborAddress = types.StringValue(value.String())
			} else {
				data.NeighborAddress = types.StringNull()
			}
			if value := res.Get("remoteAs"); value.Exists() {
				data.NeighborRemoteAs = types.StringValue(value.String())
			} else {
				data.NeighborRemoteAs = types.StringNull()
			}
			if value := res.Get("neighborGeneral.fallOverBFD"); value.Exists() {
				data.NeighborBfd = types.StringValue(value.String())
			} else {
				data.NeighborBfd = types.StringValue("NONE")
			}
			if value := res.Get("neighborGeneral.updateSource.id"); value.Exists() {
				data.UpdateSourceInterfaceId = types.StringValue(value.String())
			} else {
				data.UpdateSourceInterfaceId = types.StringNull()
			}
			if value := res.Get("neighborGeneral.enableAddress"); value.Exists() {
				data.EnableAddressFamily = types.BoolValue(value.Bool())
			} else {
				data.EnableAddressFamily = types.BoolValue(false)
			}
			if value := res.Get("neighborGeneral.shutdown"); value.Exists() {
				data.NeighborShutdown = types.BoolValue(value.Bool())
			} else {
				data.NeighborShutdown = types.BoolValue(false)
			}
			if value := res.Get("neighborGeneral.description"); value.Exists() {
				data.NeighborDescription = types.StringValue(value.String())
			} else {
				data.NeighborDescription = types.StringNull()
			}
			if value := res.Get("neighborFiltering.neighborDistributeLists"); value.Exists() {
				data.NeighborFilterAccessLists = make([]DeviceBGPIpv4NeighborsNeighborFilterAccessLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsNeighborFilterAccessLists{}
					if value := res.Get("accessList.id"); value.Exists() {
						data.AccessListId = types.StringValue(value.String())
					} else {
						data.AccessListId = types.StringNull()
					}
					if value := res.Get("filterUpdateAction"); value.Exists() {
						data.UpdateDirection = types.StringValue(value.String())
					} else {
						data.UpdateDirection = types.StringNull()
					}
					(*parent).NeighborFilterAccessLists = append((*parent).NeighborFilterAccessLists, data)
					return true
				})
			}
			if value := res.Get("neighborFiltering.neighborRouteMap"); value.Exists() {
				data.NeighborFilterRouteMapLists = make([]DeviceBGPIpv4NeighborsNeighborFilterRouteMapLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsNeighborFilterRouteMapLists{}
					if value := res.Get("RouteMap.id"); value.Exists() {
						data.RouteMapId = types.StringValue(value.String())
					} else {
						data.RouteMapId = types.StringNull()
					}
					if value := res.Get("filterUpdateAction"); value.Exists() {
						data.UpdateDirection = types.StringValue(value.String())
					} else {
						data.UpdateDirection = types.StringNull()
					}
					(*parent).NeighborFilterRouteMapLists = append((*parent).NeighborFilterRouteMapLists, data)
					return true
				})
			}
			if value := res.Get("neighborFiltering.ipv4PrefixListFilter"); value.Exists() {
				data.NeighborFilterPrefixLists = make([]DeviceBGPIpv4NeighborsNeighborFilterPrefixLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsNeighborFilterPrefixLists{}
					if value := res.Get("ipv4PrefixList.id"); value.Exists() {
						data.PrefixListId = types.StringValue(value.String())
					} else {
						data.PrefixListId = types.StringNull()
					}
					if value := res.Get("filterUpdateAction"); value.Exists() {
						data.UpdateDirection = types.StringValue(value.String())
					} else {
						data.UpdateDirection = types.StringNull()
					}
					(*parent).NeighborFilterPrefixLists = append((*parent).NeighborFilterPrefixLists, data)
					return true
				})
			}
			if value := res.Get("neighborFilterList.neighborFilterList"); value.Exists() {
				data.NeighborFilterAsPathLists = make([]DeviceBGPIpv4NeighborsNeighborFilterAsPathLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsNeighborFilterAsPathLists{}
					if value := res.Get("filterUpdateAction"); value.Exists() {
						data.UpdateDirection = types.StringValue(value.String())
					} else {
						data.UpdateDirection = types.StringNull()
					}
					if value := res.Get("asPathList.id"); value.Exists() {
						data.AsPathId = types.StringValue(value.String())
					} else {
						data.AsPathId = types.StringNull()
					}
					(*parent).NeighborFilterAsPathLists = append((*parent).NeighborFilterAsPathLists, data)
					return true
				})
			}
			if value := res.Get("neighborFiltering.neighborMaximumPrefix.maxPrefixLimit"); value.Exists() {
				data.NeighborFilterMaxPrefix = types.Int64Value(value.Int())
			} else {
				data.NeighborFilterMaxPrefix = types.Int64Null()
			}
			if value := res.Get("neighborFiltering.neighborMaximumPrefix.warningOnly"); value.Exists() {
				data.NeighborFilterWarningOnly = types.BoolValue(value.Bool())
			} else {
				data.NeighborFilterWarningOnly = types.BoolNull()
			}
			if value := res.Get("neighborFiltering.neighborMaximumPrefix.thresholdValue"); value.Exists() {
				data.NeighborFilterThresholdValue = types.Int64Value(value.Int())
			} else {
				data.NeighborFilterThresholdValue = types.Int64Null()
			}
			if value := res.Get("neighborFiltering.neighborMaximumPrefix.restartInterval"); value.Exists() {
				data.NeighborFilterRestartInterval = types.Int64Value(value.Int())
			} else {
				data.NeighborFilterRestartInterval = types.Int64Null()
			}
			if value := res.Get("neighborRoutes.advertisementInterval"); value.Exists() {
				data.NeighborRoutesAdvertisementInterval = types.Int64Value(value.Int())
			} else {
				data.NeighborRoutesAdvertisementInterval = types.Int64Value(0)
			}
			if value := res.Get("neighborRoutes.removePrivateAs"); value.Exists() {
				data.NeighborRoutesRemovePrivateAs = types.BoolValue(value.Bool())
			} else {
				data.NeighborRoutesRemovePrivateAs = types.BoolValue(false)
			}
			if value := res.Get("neighborFiltering.neighborDefaultOriginate.routeMap.id"); value.Exists() {
				data.NeighborGenerateDefaultRouteMapId = types.StringValue(value.String())
			} else {
				data.NeighborGenerateDefaultRouteMapId = types.StringNull()
			}
			if value := res.Get("neighborRoutes.neighborAdvertiseMaps.existMap"); value.Exists() {
				data.NeighborRoutesAdvertiseMapUseExist = types.BoolValue(value.Bool())
			} else {
				data.NeighborRoutesAdvertiseMapUseExist = types.BoolNull()
			}
			if value := res.Get("neighborRoutes.neighborAdvertiseMaps.routeMap.id"); value.Exists() {
				data.NeighborRoutesAdvertiseMapId = types.StringValue(value.String())
			} else {
				data.NeighborRoutesAdvertiseMapId = types.StringNull()
			}
			if value := res.Get("neighborRoutes.neighborAdvertiseMaps.existRouteMap.id"); value.Exists() {
				data.NeighborRoutesAdvertiseExistNonexistMapId = types.StringValue(value.String())
			} else {
				data.NeighborRoutesAdvertiseExistNonexistMapId = types.StringNull()
			}
			if value := res.Get("neighborTimers.keepAliveInterval"); value.Exists() {
				data.NeighborKeepaliveInterval = types.Int64Value(value.Int())
			} else {
				data.NeighborKeepaliveInterval = types.Int64Null()
			}
			if value := res.Get("neighborTimers.holdTime"); value.Exists() {
				data.NeighborHoldTime = types.Int64Value(value.Int())
			} else {
				data.NeighborHoldTime = types.Int64Null()
			}
			if value := res.Get("neighborTimers.minimumHoldTime"); value.Exists() {
				data.NeighborMinHoldTime = types.Int64Value(value.Int())
			} else {
				data.NeighborMinHoldTime = types.Int64Null()
			}
			if value := res.Get("neighborAdvanced.neighborSecret"); value.Exists() {
				data.NeighborAuthenticationPassword = types.StringValue(value.String())
			} else {
				data.NeighborAuthenticationPassword = types.StringNull()
			}
			if value := res.Get("neighborAdvanced.sendCommunity"); value.Exists() {
				data.NeighborSendCommunityAttribute = types.BoolValue(value.Bool())
			} else {
				data.NeighborSendCommunityAttribute = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.nextHopSelf"); value.Exists() {
				data.NeighborNexthopSelf = types.BoolValue(value.Bool())
			} else {
				data.NeighborNexthopSelf = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.neighborHops.disableConnectedCheck"); value.Exists() {
				data.NeighborDisableConnectionVerification = types.BoolValue(value.Bool())
			} else {
				data.NeighborDisableConnectionVerification = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.neighborTransportPathMTUDiscovery.disable"); value.Exists() {
				data.NeighborTcpMtuPathDiscovery = types.BoolValue(value.Bool())
			} else {
				data.NeighborTcpMtuPathDiscovery = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.neighborHops.maxHopCount"); value.Exists() {
				data.NeighborMaxHopCount = types.Int64Value(value.Int())
			} else {
				data.NeighborMaxHopCount = types.Int64Value(1)
			}
			if value := res.Get("neighborAdvanced.neighborTransportConnectionMode.establishTCPSession"); value.Exists() {
				data.NeighborTcpTransportMode = types.BoolValue(value.Bool())
			} else {
				data.NeighborTcpTransportMode = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.neighborWeight"); value.Exists() {
				data.NeighborWeight = types.Int64Value(value.Int())
			} else {
				data.NeighborWeight = types.Int64Value(0)
			}
			if value := res.Get("neighborAdvanced.neighborVersion"); value.Exists() {
				data.NeighborVersion = types.StringValue(value.String())
			} else {
				data.NeighborVersion = types.StringValue("0")
			}
			if value := res.Get("neighborLocalAs.asNumber"); value.Exists() {
				data.NeighborCustomizedLocalAsNumber = types.StringValue(value.String())
			} else {
				data.NeighborCustomizedLocalAsNumber = types.StringNull()
			}
			if value := res.Get("neighborLocalAs.noPrepend"); value.Exists() {
				data.NeighborCustomizedNoPrepend = types.BoolValue(value.Bool())
			} else {
				data.NeighborCustomizedNoPrepend = types.BoolNull()
			}
			if value := res.Get("neighborLocalAs.replaceAs"); value.Exists() {
				data.NeighborCustomizedReplaceAs = types.BoolValue(value.Bool())
			} else {
				data.NeighborCustomizedReplaceAs = types.BoolNull()
			}
			if value := res.Get("neighborLocalAs.dualAs"); value.Exists() {
				data.NeighborCustomizedAcceptBothAs = types.BoolValue(value.Bool())
			} else {
				data.NeighborCustomizedAcceptBothAs = types.BoolNull()
			}
			(*parent).Ipv4Neighbors = append((*parent).Ipv4Neighbors, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.aggregateAddressesIPv4s"); value.Exists() {
		data.Ipv4AggregateAddresses = make([]DeviceBGPIpv4AggregateAddresses, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4AggregateAddresses{}
			if value := res.Get("asSet"); value.Exists() {
				data.GenerateAs = types.BoolValue(value.Bool())
			} else {
				data.GenerateAs = types.BoolNull()
			}
			if value := res.Get("summaryOnly"); value.Exists() {
				data.Filter = types.BoolValue(value.Bool())
			} else {
				data.Filter = types.BoolNull()
			}
			if value := res.Get("ipv4AggregateNetwork.id"); value.Exists() {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("advertiseMap.id"); value.Exists() {
				data.AdvertiseMapId = types.StringValue(value.String())
			} else {
				data.AdvertiseMapId = types.StringNull()
			}
			if value := res.Get("attributeMap.id"); value.Exists() {
				data.AttributeMapId = types.StringValue(value.String())
			} else {
				data.AttributeMapId = types.StringNull()
			}
			if value := res.Get("suppressMap.id"); value.Exists() {
				data.SuppressMapId = types.StringValue(value.String())
			} else {
				data.SuppressMapId = types.StringNull()
			}
			(*parent).Ipv4AggregateAddresses = append((*parent).Ipv4AggregateAddresses, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.distributeLists"); value.Exists() {
		data.Ipv4Filterings = make([]DeviceBGPIpv4Filterings, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4Filterings{}
			if value := res.Get("ipv4AggregateNetwork.id"); value.Exists() {
				data.AccessListId = types.StringValue(value.String())
			} else {
				data.AccessListId = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.NetworkDirection = types.StringValue(value.String())
			} else {
				data.NetworkDirection = types.StringNull()
			}
			if value := res.Get("protocol.protocol"); value.Exists() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("protocol.processId"); value.Exists() {
				data.ProrocolProcess = types.StringValue(value.String())
			} else {
				data.ProrocolProcess = types.StringNull()
			}
			(*parent).Ipv4Filterings = append((*parent).Ipv4Filterings, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.networks"); value.Exists() {
		data.Ipv4Networks = make([]DeviceBGPIpv4Networks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4Networks{}
			if value := res.Get("ipv4Address.id"); value.Exists() {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("routeMap.id"); value.Exists() {
				data.RouteMapId = types.StringValue(value.String())
			} else {
				data.RouteMapId = types.StringNull()
			}
			(*parent).Ipv4Networks = append((*parent).Ipv4Networks, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.redistributeProtocols"); value.Exists() {
		data.Ipv4Redistributions = make([]DeviceBGPIpv4Redistributions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4Redistributions{}
			if value := res.Get("type"); value.Exists() {
				data.SourceProtocol = types.StringValue(value.String())
			} else {
				data.SourceProtocol = types.StringNull()
			}
			if value := res.Get("routeMap.id"); value.Exists() {
				data.RouteMapId = types.StringValue(value.String())
			} else {
				data.RouteMapId = types.StringNull()
			}
			if value := res.Get("routeMetric.metricValue"); value.Exists() {
				data.Metric = types.Int64Value(value.Int())
			} else {
				data.Metric = types.Int64Null()
			}
			if value := res.Get("processId"); value.Exists() {
				data.ProcessId = types.StringValue(value.String())
			} else {
				data.ProcessId = types.StringNull()
			}
			if value := res.Get("matchExternal1"); value.Exists() {
				data.MatchExternal1 = types.BoolValue(value.Bool())
			} else {
				data.MatchExternal1 = types.BoolNull()
			}
			if value := res.Get("matchExternal2"); value.Exists() {
				data.MatchExternal2 = types.BoolValue(value.Bool())
			} else {
				data.MatchExternal2 = types.BoolNull()
			}
			if value := res.Get("matchInternal"); value.Exists() {
				data.MatchInternal = types.BoolValue(value.Bool())
			} else {
				data.MatchInternal = types.BoolNull()
			}
			if value := res.Get("matchNssaExternal1"); value.Exists() {
				data.MatchNssaExternal1 = types.BoolValue(value.Bool())
			} else {
				data.MatchNssaExternal1 = types.BoolNull()
			}
			if value := res.Get("matchNssaExternal2"); value.Exists() {
				data.MatchNssaExternal2 = types.BoolValue(value.Bool())
			} else {
				data.MatchNssaExternal2 = types.BoolNull()
			}
			(*parent).Ipv4Redistributions = append((*parent).Ipv4Redistributions, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.injectMaps"); value.Exists() {
		data.Ipv4RouteInjections = make([]DeviceBGPIpv4RouteInjections, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4RouteInjections{}
			if value := res.Get("injectMap.routeMap.id"); value.Exists() {
				data.InjectRouteMapId = types.StringValue(value.String())
			} else {
				data.InjectRouteMapId = types.StringNull()
			}
			if value := res.Get("existMap.routeMap.id"); value.Exists() {
				data.ExistRouteMapId = types.StringValue(value.String())
			} else {
				data.ExistRouteMapId = types.StringNull()
			}
			(*parent).Ipv4RouteInjections = append((*parent).Ipv4RouteInjections, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceBGP) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("asNumber"); value.Exists() && !data.AsNumber.IsNull() {
		data.AsNumber = types.StringValue(value.String())
	} else {
		data.AsNumber = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.type"); value.Exists() && !data.Ipv4AddressFamilyType.IsNull() {
		data.Ipv4AddressFamilyType = types.StringValue(value.String())
	} else {
		data.Ipv4AddressFamilyType = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.aftableMap.id"); value.Exists() && !data.Ipv4LearnedRouteMapId.IsNull() {
		data.Ipv4LearnedRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4LearnedRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.defaultInformationOrginate"); value.Exists() && !data.Ipv4DefaultInformationOrginate.IsNull() {
		data.Ipv4DefaultInformationOrginate = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DefaultInformationOrginate = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.autoSummary"); value.Exists() && !data.Ipv4AutoAummary.IsNull() {
		data.Ipv4AutoAummary = types.BoolValue(value.Bool())
	} else {
		data.Ipv4AutoAummary = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpSupressInactive"); value.Exists() && !data.Ipv4BgpSupressInactive.IsNull() {
		data.Ipv4BgpSupressInactive = types.BoolValue(value.Bool())
	} else {
		data.Ipv4BgpSupressInactive = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.synchronization"); value.Exists() && !data.Ipv4Synchronization.IsNull() {
		data.Ipv4Synchronization = types.BoolValue(value.Bool())
	} else {
		data.Ipv4Synchronization = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpRedistributeInternal"); value.Exists() && !data.Ipv4BgpRedistributeInternal.IsNull() {
		data.Ipv4BgpRedistributeInternal = types.BoolValue(value.Bool())
	} else {
		data.Ipv4BgpRedistributeInternal = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.distance.externalDistance"); value.Exists() && !data.Ipv4ExternalDistance.IsNull() {
		data.Ipv4ExternalDistance = types.Int64Value(value.Int())
	} else if data.Ipv4ExternalDistance.ValueInt64() != 20 {
		data.Ipv4ExternalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.distance.internalDistance"); value.Exists() && !data.Ipv4InternalDistance.IsNull() {
		data.Ipv4InternalDistance = types.Int64Value(value.Int())
	} else if data.Ipv4InternalDistance.ValueInt64() != 200 {
		data.Ipv4InternalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.distance.localDistance"); value.Exists() && !data.Ipv4LocalDistance.IsNull() {
		data.Ipv4LocalDistance = types.Int64Value(value.Int())
	} else if data.Ipv4LocalDistance.ValueInt64() != 200 {
		data.Ipv4LocalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.ibgp"); value.Exists() && !data.Ipv4ForwardPacketsOverMultipathIbgp.IsNull() {
		data.Ipv4ForwardPacketsOverMultipathIbgp = types.Int64Value(value.Int())
	} else if data.Ipv4ForwardPacketsOverMultipathIbgp.ValueInt64() != 1 {
		data.Ipv4ForwardPacketsOverMultipathIbgp = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.ebgp"); value.Exists() && !data.Ipv4ForwardPacketsOverMultipathEbgp.IsNull() {
		data.Ipv4ForwardPacketsOverMultipathEbgp = types.Int64Value(value.Int())
	} else if data.Ipv4ForwardPacketsOverMultipathEbgp.ValueInt64() != 1 {
		data.Ipv4ForwardPacketsOverMultipathEbgp = types.Int64Null()
	}
	for i := 0; i < len(data.Ipv4Neighbors); i++ {
		keys := [...]string{"ipv4Address", "remoteAs", "neighborRoutes.neighborAdvertiseMaps.routeMap.id", "neighborRoutes.neighborAdvertiseMaps.existRouteMap.id"}
		keyValues := [...]string{data.Ipv4Neighbors[i].NeighborAddress.ValueString(), data.Ipv4Neighbors[i].NeighborRemoteAs.ValueString(), data.Ipv4Neighbors[i].NeighborRoutesAdvertiseMapId.ValueString(), data.Ipv4Neighbors[i].NeighborRoutesAdvertiseExistNonexistMapId.ValueString()}

		parent := &data
		data := (*parent).Ipv4Neighbors[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.neighbors").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4Neighbors[%d] = %+v",
				i,
				(*parent).Ipv4Neighbors[i],
			))
			(*parent).Ipv4Neighbors = slices.Delete((*parent).Ipv4Neighbors, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipv4Address"); value.Exists() && !data.NeighborAddress.IsNull() {
			data.NeighborAddress = types.StringValue(value.String())
		} else {
			data.NeighborAddress = types.StringNull()
		}
		if value := res.Get("remoteAs"); value.Exists() && !data.NeighborRemoteAs.IsNull() {
			data.NeighborRemoteAs = types.StringValue(value.String())
		} else {
			data.NeighborRemoteAs = types.StringNull()
		}
		if value := res.Get("neighborGeneral.fallOverBFD"); value.Exists() && !data.NeighborBfd.IsNull() {
			data.NeighborBfd = types.StringValue(value.String())
		} else if data.NeighborBfd.ValueString() != "NONE" {
			data.NeighborBfd = types.StringNull()
		}
		if value := res.Get("neighborGeneral.updateSource.id"); value.Exists() && !data.UpdateSourceInterfaceId.IsNull() {
			data.UpdateSourceInterfaceId = types.StringValue(value.String())
		} else {
			data.UpdateSourceInterfaceId = types.StringNull()
		}
		if value := res.Get("neighborGeneral.enableAddress"); value.Exists() && !data.EnableAddressFamily.IsNull() {
			data.EnableAddressFamily = types.BoolValue(value.Bool())
		} else if data.EnableAddressFamily.ValueBool() != false {
			data.EnableAddressFamily = types.BoolNull()
		}
		if value := res.Get("neighborGeneral.shutdown"); value.Exists() && !data.NeighborShutdown.IsNull() {
			data.NeighborShutdown = types.BoolValue(value.Bool())
		} else if data.NeighborShutdown.ValueBool() != false {
			data.NeighborShutdown = types.BoolNull()
		}
		if value := res.Get("neighborGeneral.description"); value.Exists() && !data.NeighborDescription.IsNull() {
			data.NeighborDescription = types.StringValue(value.String())
		} else {
			data.NeighborDescription = types.StringNull()
		}
		for i := 0; i < len(data.NeighborFilterAccessLists); i++ {
			keys := [...]string{"accessList.id", "filterUpdateAction"}
			keyValues := [...]string{data.NeighborFilterAccessLists[i].AccessListId.ValueString(), data.NeighborFilterAccessLists[i].UpdateDirection.ValueString()}

			parent := &data
			data := (*parent).NeighborFilterAccessLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborFiltering.neighborDistributeLists").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing NeighborFilterAccessLists[%d] = %+v",
					i,
					(*parent).NeighborFilterAccessLists[i],
				))
				(*parent).NeighborFilterAccessLists = slices.Delete((*parent).NeighborFilterAccessLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("accessList.id"); value.Exists() && !data.AccessListId.IsNull() {
				data.AccessListId = types.StringValue(value.String())
			} else {
				data.AccessListId = types.StringNull()
			}
			if value := res.Get("filterUpdateAction"); value.Exists() && !data.UpdateDirection.IsNull() {
				data.UpdateDirection = types.StringValue(value.String())
			} else {
				data.UpdateDirection = types.StringNull()
			}
			(*parent).NeighborFilterAccessLists[i] = data
		}
		for i := 0; i < len(data.NeighborFilterRouteMapLists); i++ {
			keys := [...]string{"RouteMap.id", "filterUpdateAction"}
			keyValues := [...]string{data.NeighborFilterRouteMapLists[i].RouteMapId.ValueString(), data.NeighborFilterRouteMapLists[i].UpdateDirection.ValueString()}

			parent := &data
			data := (*parent).NeighborFilterRouteMapLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborFiltering.neighborRouteMap").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing NeighborFilterRouteMapLists[%d] = %+v",
					i,
					(*parent).NeighborFilterRouteMapLists[i],
				))
				(*parent).NeighborFilterRouteMapLists = slices.Delete((*parent).NeighborFilterRouteMapLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("RouteMap.id"); value.Exists() && !data.RouteMapId.IsNull() {
				data.RouteMapId = types.StringValue(value.String())
			} else {
				data.RouteMapId = types.StringNull()
			}
			if value := res.Get("filterUpdateAction"); value.Exists() && !data.UpdateDirection.IsNull() {
				data.UpdateDirection = types.StringValue(value.String())
			} else {
				data.UpdateDirection = types.StringNull()
			}
			(*parent).NeighborFilterRouteMapLists[i] = data
		}
		for i := 0; i < len(data.NeighborFilterPrefixLists); i++ {
			keys := [...]string{"ipv4PrefixList.id", "filterUpdateAction"}
			keyValues := [...]string{data.NeighborFilterPrefixLists[i].PrefixListId.ValueString(), data.NeighborFilterPrefixLists[i].UpdateDirection.ValueString()}

			parent := &data
			data := (*parent).NeighborFilterPrefixLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborFiltering.ipv4PrefixListFilter").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing NeighborFilterPrefixLists[%d] = %+v",
					i,
					(*parent).NeighborFilterPrefixLists[i],
				))
				(*parent).NeighborFilterPrefixLists = slices.Delete((*parent).NeighborFilterPrefixLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("ipv4PrefixList.id"); value.Exists() && !data.PrefixListId.IsNull() {
				data.PrefixListId = types.StringValue(value.String())
			} else {
				data.PrefixListId = types.StringNull()
			}
			if value := res.Get("filterUpdateAction"); value.Exists() && !data.UpdateDirection.IsNull() {
				data.UpdateDirection = types.StringValue(value.String())
			} else {
				data.UpdateDirection = types.StringNull()
			}
			(*parent).NeighborFilterPrefixLists[i] = data
		}
		for i := 0; i < len(data.NeighborFilterAsPathLists); i++ {
			keys := [...]string{"filterUpdateAction", "asPathList.id"}
			keyValues := [...]string{data.NeighborFilterAsPathLists[i].UpdateDirection.ValueString(), data.NeighborFilterAsPathLists[i].AsPathId.ValueString()}

			parent := &data
			data := (*parent).NeighborFilterAsPathLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborFilterList.neighborFilterList").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing NeighborFilterAsPathLists[%d] = %+v",
					i,
					(*parent).NeighborFilterAsPathLists[i],
				))
				(*parent).NeighborFilterAsPathLists = slices.Delete((*parent).NeighborFilterAsPathLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("filterUpdateAction"); value.Exists() && !data.UpdateDirection.IsNull() {
				data.UpdateDirection = types.StringValue(value.String())
			} else {
				data.UpdateDirection = types.StringNull()
			}
			if value := res.Get("asPathList.id"); value.Exists() && !data.AsPathId.IsNull() {
				data.AsPathId = types.StringValue(value.String())
			} else {
				data.AsPathId = types.StringNull()
			}
			(*parent).NeighborFilterAsPathLists[i] = data
		}
		if value := res.Get("neighborFiltering.neighborMaximumPrefix.maxPrefixLimit"); value.Exists() && !data.NeighborFilterMaxPrefix.IsNull() {
			data.NeighborFilterMaxPrefix = types.Int64Value(value.Int())
		} else {
			data.NeighborFilterMaxPrefix = types.Int64Null()
		}
		if value := res.Get("neighborFiltering.neighborMaximumPrefix.warningOnly"); value.Exists() && !data.NeighborFilterWarningOnly.IsNull() {
			data.NeighborFilterWarningOnly = types.BoolValue(value.Bool())
		} else {
			data.NeighborFilterWarningOnly = types.BoolNull()
		}
		if value := res.Get("neighborFiltering.neighborMaximumPrefix.thresholdValue"); value.Exists() && !data.NeighborFilterThresholdValue.IsNull() {
			data.NeighborFilterThresholdValue = types.Int64Value(value.Int())
		} else {
			data.NeighborFilterThresholdValue = types.Int64Null()
		}
		if value := res.Get("neighborFiltering.neighborMaximumPrefix.restartInterval"); value.Exists() && !data.NeighborFilterRestartInterval.IsNull() {
			data.NeighborFilterRestartInterval = types.Int64Value(value.Int())
		} else {
			data.NeighborFilterRestartInterval = types.Int64Null()
		}
		if value := res.Get("neighborRoutes.advertisementInterval"); value.Exists() && !data.NeighborRoutesAdvertisementInterval.IsNull() {
			data.NeighborRoutesAdvertisementInterval = types.Int64Value(value.Int())
		} else if data.NeighborRoutesAdvertisementInterval.ValueInt64() != 0 {
			data.NeighborRoutesAdvertisementInterval = types.Int64Null()
		}
		if value := res.Get("neighborRoutes.removePrivateAs"); value.Exists() && !data.NeighborRoutesRemovePrivateAs.IsNull() {
			data.NeighborRoutesRemovePrivateAs = types.BoolValue(value.Bool())
		} else if data.NeighborRoutesRemovePrivateAs.ValueBool() != false {
			data.NeighborRoutesRemovePrivateAs = types.BoolNull()
		}
		if value := res.Get("neighborFiltering.neighborDefaultOriginate.routeMap.id"); value.Exists() && !data.NeighborGenerateDefaultRouteMapId.IsNull() {
			data.NeighborGenerateDefaultRouteMapId = types.StringValue(value.String())
		} else {
			data.NeighborGenerateDefaultRouteMapId = types.StringNull()
		}
		if value := res.Get("neighborRoutes.neighborAdvertiseMaps.existMap"); value.Exists() && !data.NeighborRoutesAdvertiseMapUseExist.IsNull() {
			data.NeighborRoutesAdvertiseMapUseExist = types.BoolValue(value.Bool())
		} else {
			data.NeighborRoutesAdvertiseMapUseExist = types.BoolNull()
		}
		if value := res.Get("neighborRoutes.neighborAdvertiseMaps.routeMap.id"); value.Exists() && !data.NeighborRoutesAdvertiseMapId.IsNull() {
			data.NeighborRoutesAdvertiseMapId = types.StringValue(value.String())
		} else {
			data.NeighborRoutesAdvertiseMapId = types.StringNull()
		}
		if value := res.Get("neighborRoutes.neighborAdvertiseMaps.existRouteMap.id"); value.Exists() && !data.NeighborRoutesAdvertiseExistNonexistMapId.IsNull() {
			data.NeighborRoutesAdvertiseExistNonexistMapId = types.StringValue(value.String())
		} else {
			data.NeighborRoutesAdvertiseExistNonexistMapId = types.StringNull()
		}
		if value := res.Get("neighborTimers.keepAliveInterval"); value.Exists() && !data.NeighborKeepaliveInterval.IsNull() {
			data.NeighborKeepaliveInterval = types.Int64Value(value.Int())
		} else {
			data.NeighborKeepaliveInterval = types.Int64Null()
		}
		if value := res.Get("neighborTimers.holdTime"); value.Exists() && !data.NeighborHoldTime.IsNull() {
			data.NeighborHoldTime = types.Int64Value(value.Int())
		} else {
			data.NeighborHoldTime = types.Int64Null()
		}
		if value := res.Get("neighborTimers.minimumHoldTime"); value.Exists() && !data.NeighborMinHoldTime.IsNull() {
			data.NeighborMinHoldTime = types.Int64Value(value.Int())
		} else {
			data.NeighborMinHoldTime = types.Int64Null()
		}
		if value := res.Get("neighborAdvanced.neighborSecret"); value.Exists() && !data.NeighborAuthenticationPassword.IsNull() {
			data.NeighborAuthenticationPassword = types.StringValue(value.String())
		} else {
			data.NeighborAuthenticationPassword = types.StringNull()
		}
		if value := res.Get("neighborAdvanced.sendCommunity"); value.Exists() && !data.NeighborSendCommunityAttribute.IsNull() {
			data.NeighborSendCommunityAttribute = types.BoolValue(value.Bool())
		} else if data.NeighborSendCommunityAttribute.ValueBool() != false {
			data.NeighborSendCommunityAttribute = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.nextHopSelf"); value.Exists() && !data.NeighborNexthopSelf.IsNull() {
			data.NeighborNexthopSelf = types.BoolValue(value.Bool())
		} else if data.NeighborNexthopSelf.ValueBool() != false {
			data.NeighborNexthopSelf = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.neighborHops.disableConnectedCheck"); value.Exists() && !data.NeighborDisableConnectionVerification.IsNull() {
			data.NeighborDisableConnectionVerification = types.BoolValue(value.Bool())
		} else if data.NeighborDisableConnectionVerification.ValueBool() != false {
			data.NeighborDisableConnectionVerification = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.neighborTransportPathMTUDiscovery.disable"); value.Exists() && !data.NeighborTcpMtuPathDiscovery.IsNull() {
			data.NeighborTcpMtuPathDiscovery = types.BoolValue(value.Bool())
		} else if data.NeighborTcpMtuPathDiscovery.ValueBool() != false {
			data.NeighborTcpMtuPathDiscovery = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.neighborHops.maxHopCount"); value.Exists() && !data.NeighborMaxHopCount.IsNull() {
			data.NeighborMaxHopCount = types.Int64Value(value.Int())
		} else if data.NeighborMaxHopCount.ValueInt64() != 1 {
			data.NeighborMaxHopCount = types.Int64Null()
		}
		if value := res.Get("neighborAdvanced.neighborTransportConnectionMode.establishTCPSession"); value.Exists() && !data.NeighborTcpTransportMode.IsNull() {
			data.NeighborTcpTransportMode = types.BoolValue(value.Bool())
		} else if data.NeighborTcpTransportMode.ValueBool() != false {
			data.NeighborTcpTransportMode = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.neighborWeight"); value.Exists() && !data.NeighborWeight.IsNull() {
			data.NeighborWeight = types.Int64Value(value.Int())
		} else if data.NeighborWeight.ValueInt64() != 0 {
			data.NeighborWeight = types.Int64Null()
		}
		if value := res.Get("neighborAdvanced.neighborVersion"); value.Exists() && !data.NeighborVersion.IsNull() {
			data.NeighborVersion = types.StringValue(value.String())
		} else if data.NeighborVersion.ValueString() != "0" {
			data.NeighborVersion = types.StringNull()
		}
		if value := res.Get("neighborLocalAs.asNumber"); value.Exists() && !data.NeighborCustomizedLocalAsNumber.IsNull() {
			data.NeighborCustomizedLocalAsNumber = types.StringValue(value.String())
		} else {
			data.NeighborCustomizedLocalAsNumber = types.StringNull()
		}
		if value := res.Get("neighborLocalAs.noPrepend"); value.Exists() && !data.NeighborCustomizedNoPrepend.IsNull() {
			data.NeighborCustomizedNoPrepend = types.BoolValue(value.Bool())
		} else {
			data.NeighborCustomizedNoPrepend = types.BoolNull()
		}
		if value := res.Get("neighborLocalAs.replaceAs"); value.Exists() && !data.NeighborCustomizedReplaceAs.IsNull() {
			data.NeighborCustomizedReplaceAs = types.BoolValue(value.Bool())
		} else {
			data.NeighborCustomizedReplaceAs = types.BoolNull()
		}
		if value := res.Get("neighborLocalAs.dualAs"); value.Exists() && !data.NeighborCustomizedAcceptBothAs.IsNull() {
			data.NeighborCustomizedAcceptBothAs = types.BoolValue(value.Bool())
		} else {
			data.NeighborCustomizedAcceptBothAs = types.BoolNull()
		}
		(*parent).Ipv4Neighbors[i] = data
	}
	for i := 0; i < len(data.Ipv4AggregateAddresses); i++ {
		keys := [...]string{"ipv4AggregateNetwork.id", "advertiseMap.id", "attributeMap.id", "suppressMap.id"}
		keyValues := [...]string{data.Ipv4AggregateAddresses[i].NetworkId.ValueString(), data.Ipv4AggregateAddresses[i].AdvertiseMapId.ValueString(), data.Ipv4AggregateAddresses[i].AttributeMapId.ValueString(), data.Ipv4AggregateAddresses[i].SuppressMapId.ValueString()}

		parent := &data
		data := (*parent).Ipv4AggregateAddresses[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.aggregateAddressesIPv4s").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4AggregateAddresses[%d] = %+v",
				i,
				(*parent).Ipv4AggregateAddresses[i],
			))
			(*parent).Ipv4AggregateAddresses = slices.Delete((*parent).Ipv4AggregateAddresses, i, i+1)
			i--

			continue
		}
		if value := res.Get("asSet"); value.Exists() && !data.GenerateAs.IsNull() {
			data.GenerateAs = types.BoolValue(value.Bool())
		} else {
			data.GenerateAs = types.BoolNull()
		}
		if value := res.Get("summaryOnly"); value.Exists() && !data.Filter.IsNull() {
			data.Filter = types.BoolValue(value.Bool())
		} else {
			data.Filter = types.BoolNull()
		}
		if value := res.Get("ipv4AggregateNetwork.id"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("advertiseMap.id"); value.Exists() && !data.AdvertiseMapId.IsNull() {
			data.AdvertiseMapId = types.StringValue(value.String())
		} else {
			data.AdvertiseMapId = types.StringNull()
		}
		if value := res.Get("attributeMap.id"); value.Exists() && !data.AttributeMapId.IsNull() {
			data.AttributeMapId = types.StringValue(value.String())
		} else {
			data.AttributeMapId = types.StringNull()
		}
		if value := res.Get("suppressMap.id"); value.Exists() && !data.SuppressMapId.IsNull() {
			data.SuppressMapId = types.StringValue(value.String())
		} else {
			data.SuppressMapId = types.StringNull()
		}
		(*parent).Ipv4AggregateAddresses[i] = data
	}
	for i := 0; i < len(data.Ipv4Filterings); i++ {
		keys := [...]string{"ipv4AggregateNetwork.id"}
		keyValues := [...]string{data.Ipv4Filterings[i].AccessListId.ValueString()}

		parent := &data
		data := (*parent).Ipv4Filterings[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.distributeLists").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4Filterings[%d] = %+v",
				i,
				(*parent).Ipv4Filterings[i],
			))
			(*parent).Ipv4Filterings = slices.Delete((*parent).Ipv4Filterings, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipv4AggregateNetwork.id"); value.Exists() && !data.AccessListId.IsNull() {
			data.AccessListId = types.StringValue(value.String())
		} else {
			data.AccessListId = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.NetworkDirection.IsNull() {
			data.NetworkDirection = types.StringValue(value.String())
		} else {
			data.NetworkDirection = types.StringNull()
		}
		if value := res.Get("protocol.protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("protocol.processId"); value.Exists() && !data.ProrocolProcess.IsNull() {
			data.ProrocolProcess = types.StringValue(value.String())
		} else {
			data.ProrocolProcess = types.StringNull()
		}
		(*parent).Ipv4Filterings[i] = data
	}
	for i := 0; i < len(data.Ipv4Networks); i++ {
		keys := [...]string{"ipv4Address.id", "routeMap.id"}
		keyValues := [...]string{data.Ipv4Networks[i].NetworkId.ValueString(), data.Ipv4Networks[i].RouteMapId.ValueString()}

		parent := &data
		data := (*parent).Ipv4Networks[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.networks").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4Networks[%d] = %+v",
				i,
				(*parent).Ipv4Networks[i],
			))
			(*parent).Ipv4Networks = slices.Delete((*parent).Ipv4Networks, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipv4Address.id"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("routeMap.id"); value.Exists() && !data.RouteMapId.IsNull() {
			data.RouteMapId = types.StringValue(value.String())
		} else {
			data.RouteMapId = types.StringNull()
		}
		(*parent).Ipv4Networks[i] = data
	}
	for i := 0; i < len(data.Ipv4Redistributions); i++ {
		keys := [...]string{"routeMap.id"}
		keyValues := [...]string{data.Ipv4Redistributions[i].RouteMapId.ValueString()}

		parent := &data
		data := (*parent).Ipv4Redistributions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.redistributeProtocols").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4Redistributions[%d] = %+v",
				i,
				(*parent).Ipv4Redistributions[i],
			))
			(*parent).Ipv4Redistributions = slices.Delete((*parent).Ipv4Redistributions, i, i+1)
			i--

			continue
		}
		if value := res.Get("type"); value.Exists() && !data.SourceProtocol.IsNull() {
			data.SourceProtocol = types.StringValue(value.String())
		} else {
			data.SourceProtocol = types.StringNull()
		}
		if value := res.Get("routeMap.id"); value.Exists() && !data.RouteMapId.IsNull() {
			data.RouteMapId = types.StringValue(value.String())
		} else {
			data.RouteMapId = types.StringNull()
		}
		if value := res.Get("routeMetric.metricValue"); value.Exists() && !data.Metric.IsNull() {
			data.Metric = types.Int64Value(value.Int())
		} else {
			data.Metric = types.Int64Null()
		}
		if value := res.Get("processId"); value.Exists() && !data.ProcessId.IsNull() {
			data.ProcessId = types.StringValue(value.String())
		} else {
			data.ProcessId = types.StringNull()
		}
		if value := res.Get("matchExternal1"); value.Exists() && !data.MatchExternal1.IsNull() {
			data.MatchExternal1 = types.BoolValue(value.Bool())
		} else {
			data.MatchExternal1 = types.BoolNull()
		}
		if value := res.Get("matchExternal2"); value.Exists() && !data.MatchExternal2.IsNull() {
			data.MatchExternal2 = types.BoolValue(value.Bool())
		} else {
			data.MatchExternal2 = types.BoolNull()
		}
		if value := res.Get("matchInternal"); value.Exists() && !data.MatchInternal.IsNull() {
			data.MatchInternal = types.BoolValue(value.Bool())
		} else {
			data.MatchInternal = types.BoolNull()
		}
		if value := res.Get("matchNssaExternal1"); value.Exists() && !data.MatchNssaExternal1.IsNull() {
			data.MatchNssaExternal1 = types.BoolValue(value.Bool())
		} else {
			data.MatchNssaExternal1 = types.BoolNull()
		}
		if value := res.Get("matchNssaExternal2"); value.Exists() && !data.MatchNssaExternal2.IsNull() {
			data.MatchNssaExternal2 = types.BoolValue(value.Bool())
		} else {
			data.MatchNssaExternal2 = types.BoolNull()
		}
		(*parent).Ipv4Redistributions[i] = data
	}
	for i := 0; i < len(data.Ipv4RouteInjections); i++ {
		keys := [...]string{"injectMap.routeMap.id", "existMap.routeMap.id"}
		keyValues := [...]string{data.Ipv4RouteInjections[i].InjectRouteMapId.ValueString(), data.Ipv4RouteInjections[i].ExistRouteMapId.ValueString()}

		parent := &data
		data := (*parent).Ipv4RouteInjections[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.injectMaps").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4RouteInjections[%d] = %+v",
				i,
				(*parent).Ipv4RouteInjections[i],
			))
			(*parent).Ipv4RouteInjections = slices.Delete((*parent).Ipv4RouteInjections, i, i+1)
			i--

			continue
		}
		if value := res.Get("injectMap.routeMap.id"); value.Exists() && !data.InjectRouteMapId.IsNull() {
			data.InjectRouteMapId = types.StringValue(value.String())
		} else {
			data.InjectRouteMapId = types.StringNull()
		}
		if value := res.Get("existMap.routeMap.id"); value.Exists() && !data.ExistRouteMapId.IsNull() {
			data.ExistRouteMapId = types.StringValue(value.String())
		} else {
			data.ExistRouteMapId = types.StringNull()
		}
		(*parent).Ipv4RouteInjections[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceBGP) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Name.IsUnknown() {
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
	}
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.AsNumber.IsUnknown() {
		if value := res.Get("asNumber"); value.Exists() {
			data.AsNumber = types.StringValue(value.String())
		} else {
			data.AsNumber = types.StringNull()
		}
	}
	if data.Ipv4AddressFamilyType.IsUnknown() {
		if value := res.Get("addressFamilyIPv4.type"); value.Exists() {
			data.Ipv4AddressFamilyType = types.StringValue(value.String())
		} else {
			data.Ipv4AddressFamilyType = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
