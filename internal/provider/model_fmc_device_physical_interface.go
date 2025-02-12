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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DevicePhysicalInterface struct {
	Id                                    types.String                                            `tfsdk:"id"`
	Domain                                types.String                                            `tfsdk:"domain"`
	DeviceId                              types.String                                            `tfsdk:"device_id"`
	Type                                  types.String                                            `tfsdk:"type"`
	LogicalName                           types.String                                            `tfsdk:"logical_name"`
	Enabled                               types.Bool                                              `tfsdk:"enabled"`
	ManagementOnly                        types.Bool                                              `tfsdk:"management_only"`
	Description                           types.String                                            `tfsdk:"description"`
	Mode                                  types.String                                            `tfsdk:"mode"`
	SecurityZoneId                        types.String                                            `tfsdk:"security_zone_id"`
	Name                                  types.String                                            `tfsdk:"name"`
	Mtu                                   types.Int64                                             `tfsdk:"mtu"`
	Priority                              types.Int64                                             `tfsdk:"priority"`
	EnableSgtPropagate                    types.Bool                                              `tfsdk:"enable_sgt_propagate"`
	NveOnly                               types.Bool                                              `tfsdk:"nve_only"`
	Ipv4StaticAddress                     types.String                                            `tfsdk:"ipv4_static_address"`
	Ipv4StaticNetmask                     types.String                                            `tfsdk:"ipv4_static_netmask"`
	Ipv4DhcpObtainRoute                   types.Bool                                              `tfsdk:"ipv4_dhcp_obtain_route"`
	Ipv4DhcpRouteMetric                   types.Int64                                             `tfsdk:"ipv4_dhcp_route_metric"`
	Ipv4PppoeVpdnGroupName                types.String                                            `tfsdk:"ipv4_pppoe_vpdn_group_name"`
	Ipv4PppoeUser                         types.String                                            `tfsdk:"ipv4_pppoe_user"`
	Ipv4PppoePassword                     types.String                                            `tfsdk:"ipv4_pppoe_password"`
	Ipv4PppoeAuthentication               types.String                                            `tfsdk:"ipv4_pppoe_authentication"`
	Ipv4PppoeRouteMetric                  types.Int64                                             `tfsdk:"ipv4_pppoe_route_metric"`
	Ipv4PppoeRouteSettings                types.Bool                                              `tfsdk:"ipv4_pppoe_route_settings"`
	Ipv4PppoeStoreCredentialsInFlash      types.Bool                                              `tfsdk:"ipv4_pppoe_store_credentials_in_flash"`
	Ipv6Enable                            types.Bool                                              `tfsdk:"ipv6_enable"`
	Ipv6EnforceEui                        types.Bool                                              `tfsdk:"ipv6_enforce_eui"`
	Ipv6LinkLocalAddress                  types.String                                            `tfsdk:"ipv6_link_local_address"`
	Ipv6EnableAutoConfig                  types.Bool                                              `tfsdk:"ipv6_enable_auto_config"`
	Ipv6Addresses                         []DevicePhysicalInterfaceIpv6Addresses                  `tfsdk:"ipv6_addresses"`
	Ipv6Prefixes                          []DevicePhysicalInterfaceIpv6Prefixes                   `tfsdk:"ipv6_prefixes"`
	Ipv6EnableDad                         types.Bool                                              `tfsdk:"ipv6_enable_dad"`
	Ipv6DadAttempts                       types.Int64                                             `tfsdk:"ipv6_dad_attempts"`
	Ipv6NsInterval                        types.Int64                                             `tfsdk:"ipv6_ns_interval"`
	Ipv6ReachableTime                     types.Int64                                             `tfsdk:"ipv6_reachable_time"`
	Ipv6EnableRa                          types.Bool                                              `tfsdk:"ipv6_enable_ra"`
	Ipv6RaLifeTime                        types.Int64                                             `tfsdk:"ipv6_ra_life_time"`
	Ipv6RaInterval                        types.Int64                                             `tfsdk:"ipv6_ra_interval"`
	Ipv6Dhcp                              types.Bool                                              `tfsdk:"ipv6_dhcp"`
	Ipv6DefaultRouteByDhcp                types.Bool                                              `tfsdk:"ipv6_default_route_by_dhcp"`
	Ipv6DhcpPoolId                        types.String                                            `tfsdk:"ipv6_dhcp_pool_id"`
	Ipv6DhcpPoolType                      types.String                                            `tfsdk:"ipv6_dhcp_pool_type"`
	Ipv6EnableDhcpAddressConfig           types.Bool                                              `tfsdk:"ipv6_enable_dhcp_address_config"`
	Ipv6EnableDhcpNonaddressConfig        types.Bool                                              `tfsdk:"ipv6_enable_dhcp_nonaddress_config"`
	Ipv6DhcpClientPdPrefixName            types.String                                            `tfsdk:"ipv6_dhcp_client_pd_prefix_name"`
	Ipv6DhcpClientPdHintPrefixes          types.String                                            `tfsdk:"ipv6_dhcp_client_pd_hint_prefixes"`
	IpBasedMonitoring                     types.Bool                                              `tfsdk:"ip_based_monitoring"`
	IpBasedMonitoringType                 types.String                                            `tfsdk:"ip_based_monitoring_type"`
	IpBasedMonitoringNextHop              types.String                                            `tfsdk:"ip_based_monitoring_next_hop"`
	AutoNegotiation                       types.Bool                                              `tfsdk:"auto_negotiation"`
	Duplex                                types.String                                            `tfsdk:"duplex"`
	Speed                                 types.String                                            `tfsdk:"speed"`
	LldpReceive                           types.Bool                                              `tfsdk:"lldp_receive"`
	LldpTransmit                          types.Bool                                              `tfsdk:"lldp_transmit"`
	FlowControlSend                       types.String                                            `tfsdk:"flow_control_send"`
	FecMode                               types.String                                            `tfsdk:"fec_mode"`
	ManagementAccess                      types.Bool                                              `tfsdk:"management_access"`
	ManagementAccessNetworkObjects        []DevicePhysicalInterfaceManagementAccessNetworkObjects `tfsdk:"management_access_network_objects"`
	ActiveMacAddress                      types.String                                            `tfsdk:"active_mac_address"`
	StandbyMacAddress                     types.String                                            `tfsdk:"standby_mac_address"`
	ArpTableEntries                       []DevicePhysicalInterfaceArpTableEntries                `tfsdk:"arp_table_entries"`
	EnableAntiSpoofing                    types.Bool                                              `tfsdk:"enable_anti_spoofing"`
	AllowFullFragmentReassembly           types.Bool                                              `tfsdk:"allow_full_fragment_reassembly"`
	OverrideDefaultFragmentSettingChain   types.Int64                                             `tfsdk:"override_default_fragment_setting_chain"`
	OverrideDefaultFragmentSettingSize    types.Int64                                             `tfsdk:"override_default_fragment_setting_size"`
	OverrideDefaultFragmentSettingTimeout types.Int64                                             `tfsdk:"override_default_fragment_setting_timeout"`
}

type DevicePhysicalInterfaceIpv6Addresses struct {
	Address    types.String `tfsdk:"address"`
	Prefix     types.String `tfsdk:"prefix"`
	EnforceEui types.Bool   `tfsdk:"enforce_eui"`
}

type DevicePhysicalInterfaceIpv6Prefixes struct {
	Address    types.String `tfsdk:"address"`
	Default    types.String `tfsdk:"default"`
	EnforceEui types.Bool   `tfsdk:"enforce_eui"`
}

type DevicePhysicalInterfaceManagementAccessNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type DevicePhysicalInterfaceArpTableEntries struct {
	MacAddress  types.String `tfsdk:"mac_address"`
	IpAddress   types.String `tfsdk:"ip_address"`
	EnableAlias types.Bool   `tfsdk:"enable_alias"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DevicePhysicalInterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/physicalinterfaces", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DevicePhysicalInterface) toBody(ctx context.Context, state DevicePhysicalInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "ifname", data.LogicalName.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.ManagementOnly.IsNull() {
		body, _ = sjson.Set(body, "managementOnly", data.ManagementOnly.ValueBool())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if !data.SecurityZoneId.IsNull() {
		body, _ = sjson.Set(body, "securityZone.id", data.SecurityZoneId.ValueString())
	}
	body, _ = sjson.Set(body, "securityZone.type", "SecurityZone")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Mtu.IsNull() {
		body, _ = sjson.Set(body, "MTU", data.Mtu.ValueInt64())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, "priority", data.Priority.ValueInt64())
	}
	if !data.EnableSgtPropagate.IsNull() {
		body, _ = sjson.Set(body, "enableSGTPropagate", data.EnableSgtPropagate.ValueBool())
	}
	if !data.NveOnly.IsNull() {
		body, _ = sjson.Set(body, "nveOnly", data.NveOnly.ValueBool())
	}
	if !data.Ipv4StaticAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.address", data.Ipv4StaticAddress.ValueString())
	}
	if !data.Ipv4StaticNetmask.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.netmask", data.Ipv4StaticNetmask.ValueString())
	}
	if !data.Ipv4DhcpObtainRoute.IsNull() {
		body, _ = sjson.Set(body, "ipv4.dhcp.enableDefaultRouteDHCP", data.Ipv4DhcpObtainRoute.ValueBool())
	}
	if !data.Ipv4DhcpRouteMetric.IsNull() {
		body, _ = sjson.Set(body, "ipv4.dhcp.dhcpRouteMetric", data.Ipv4DhcpRouteMetric.ValueInt64())
	}
	if !data.Ipv4PppoeVpdnGroupName.IsNull() {
		body, _ = sjson.Set(body, "ipv4.pppoe.vpdnGroupName", data.Ipv4PppoeVpdnGroupName.ValueString())
	}
	if !data.Ipv4PppoeUser.IsNull() {
		body, _ = sjson.Set(body, "ipv4.pppoe.pppoeUser", data.Ipv4PppoeUser.ValueString())
	}
	if !data.Ipv4PppoePassword.IsNull() {
		body, _ = sjson.Set(body, "ipv4.pppoe.pppoePassword", data.Ipv4PppoePassword.ValueString())
	}
	if !data.Ipv4PppoeAuthentication.IsNull() {
		body, _ = sjson.Set(body, "ipv4.pppoe.pppAuth", data.Ipv4PppoeAuthentication.ValueString())
	}
	if !data.Ipv4PppoeRouteMetric.IsNull() {
		body, _ = sjson.Set(body, "ipv4.pppoe.pppoeRouteMetric", data.Ipv4PppoeRouteMetric.ValueInt64())
	}
	if !data.Ipv4PppoeRouteSettings.IsNull() {
		body, _ = sjson.Set(body, "ipv4.pppoe.enableRouteSettings", data.Ipv4PppoeRouteSettings.ValueBool())
	}
	if !data.Ipv4PppoeStoreCredentialsInFlash.IsNull() {
		body, _ = sjson.Set(body, "ipv4.pppoe.storeCredsInFlash", data.Ipv4PppoeStoreCredentialsInFlash.ValueBool())
	}
	if !data.Ipv6Enable.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableIPV6", data.Ipv6Enable.ValueBool())
	}
	if !data.Ipv6EnforceEui.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enforceEUI64", data.Ipv6EnforceEui.ValueBool())
	}
	if !data.Ipv6LinkLocalAddress.IsNull() {
		body, _ = sjson.Set(body, "linkLocalAddress", data.Ipv6LinkLocalAddress.ValueString())
	}
	if !data.Ipv6EnableAutoConfig.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableAutoConfig", data.Ipv6EnableAutoConfig.ValueBool())
	}
	if len(data.Ipv6Addresses) > 0 {
		body, _ = sjson.Set(body, "ipv6.addresses", []interface{}{})
		for _, item := range data.Ipv6Addresses {
			itemBody := ""
			if !item.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address", item.Address.ValueString())
			}
			if !item.Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "prefix", item.Prefix.ValueString())
			}
			if !item.EnforceEui.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enforceEUI64", item.EnforceEui.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "ipv6.addresses.-1", itemBody)
		}
	}
	if len(data.Ipv6Prefixes) > 0 {
		body, _ = sjson.Set(body, "ipv6.prefixes", []interface{}{})
		for _, item := range data.Ipv6Prefixes {
			itemBody := ""
			if !item.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address", item.Address.ValueString())
			}
			if !item.Default.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "default", item.Default.ValueString())
			}
			if !item.EnforceEui.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enforceEUI64", item.EnforceEui.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "ipv6.prefixes.-1", itemBody)
		}
	}
	if !data.Ipv6EnableDad.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableDADLoopback", data.Ipv6EnableDad.ValueBool())
	}
	if !data.Ipv6DadAttempts.IsNull() {
		body, _ = sjson.Set(body, "ipv6.dadAttempts", data.Ipv6DadAttempts.ValueInt64())
	}
	if !data.Ipv6NsInterval.IsNull() {
		body, _ = sjson.Set(body, "ipv6.nsInterval", data.Ipv6NsInterval.ValueInt64())
	}
	if !data.Ipv6ReachableTime.IsNull() {
		body, _ = sjson.Set(body, "ipv6.reachableTime", data.Ipv6ReachableTime.ValueInt64())
	}
	if !data.Ipv6EnableRa.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableRA", data.Ipv6EnableRa.ValueBool())
	}
	if !data.Ipv6RaLifeTime.IsNull() {
		body, _ = sjson.Set(body, "ipv6.raLifeTime", data.Ipv6RaLifeTime.ValueInt64())
	}
	if !data.Ipv6RaInterval.IsNull() {
		body, _ = sjson.Set(body, "ipv6.raInterval", data.Ipv6RaInterval.ValueInt64())
	}
	if !data.Ipv6Dhcp.IsNull() {
		body, _ = sjson.Set(body, "ipv6.DHCP.enableDHCPClient", data.Ipv6Dhcp.ValueBool())
	}
	if !data.Ipv6DefaultRouteByDhcp.IsNull() {
		body, _ = sjson.Set(body, "ipv6.DHCP.obtainIPV6DefaultRouteDHCP", data.Ipv6DefaultRouteByDhcp.ValueBool())
	}
	if !data.Ipv6DhcpPoolId.IsNull() {
		body, _ = sjson.Set(body, "ipv6.ipv6DHCPPool.id", data.Ipv6DhcpPoolId.ValueString())
	}
	if !data.Ipv6DhcpPoolType.IsNull() {
		body, _ = sjson.Set(body, "ipv6.ipv6DHCPPool.type", data.Ipv6DhcpPoolType.ValueString())
	}
	if !data.Ipv6EnableDhcpAddressConfig.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableDHCPAddrConfig", data.Ipv6EnableDhcpAddressConfig.ValueBool())
	}
	if !data.Ipv6EnableDhcpNonaddressConfig.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableDHCPNonAddrConfig", data.Ipv6EnableDhcpNonaddressConfig.ValueBool())
	}
	if !data.Ipv6DhcpClientPdPrefixName.IsNull() {
		body, _ = sjson.Set(body, "ipv6.DHCP.clientPd.prefixName", data.Ipv6DhcpClientPdPrefixName.ValueString())
	}
	if !data.Ipv6DhcpClientPdHintPrefixes.IsNull() {
		body, _ = sjson.Set(body, "ipv6.DHCP.clientPd.hintPrefixes", data.Ipv6DhcpClientPdHintPrefixes.ValueString())
	}
	if !data.IpBasedMonitoring.IsNull() {
		body, _ = sjson.Set(body, "pathMonitoring.enable", data.IpBasedMonitoring.ValueBool())
	}
	if !data.IpBasedMonitoringType.IsNull() {
		body, _ = sjson.Set(body, "pathMonitoring.type", data.IpBasedMonitoringType.ValueString())
	}
	if !data.IpBasedMonitoringNextHop.IsNull() {
		body, _ = sjson.Set(body, "monitoredIp", data.IpBasedMonitoringNextHop.ValueString())
	}
	if !data.AutoNegotiation.IsNull() {
		body, _ = sjson.Set(body, "hardware.autoNegState", data.AutoNegotiation.ValueBool())
	}
	if !data.Duplex.IsNull() {
		body, _ = sjson.Set(body, "hardware.duplex", data.Duplex.ValueString())
	}
	if !data.Speed.IsNull() {
		body, _ = sjson.Set(body, "hardware.speed", data.Speed.ValueString())
	}
	if !data.LldpReceive.IsNull() {
		body, _ = sjson.Set(body, "LLDP.receive", data.LldpReceive.ValueBool())
	}
	if !data.LldpTransmit.IsNull() {
		body, _ = sjson.Set(body, "LLDP.transmit", data.LldpTransmit.ValueBool())
	}
	if !data.FlowControlSend.IsNull() {
		body, _ = sjson.Set(body, "hardware.flowControlSend", data.FlowControlSend.ValueString())
	}
	if !data.FecMode.IsNull() {
		body, _ = sjson.Set(body, "hardware.fecMode", data.FecMode.ValueString())
	}
	if !data.ManagementAccess.IsNull() {
		body, _ = sjson.Set(body, "fmcAccessConfig.enableAccess", data.ManagementAccess.ValueBool())
	}
	if len(data.ManagementAccessNetworkObjects) > 0 {
		body, _ = sjson.Set(body, "fmcAccessConfig.allowedNetworks", []interface{}{})
		for _, item := range data.ManagementAccessNetworkObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "fmcAccessConfig.allowedNetworks.-1", itemBody)
		}
	}
	if !data.ActiveMacAddress.IsNull() {
		body, _ = sjson.Set(body, "activeMACAddress", data.ActiveMacAddress.ValueString())
	}
	if !data.StandbyMacAddress.IsNull() {
		body, _ = sjson.Set(body, "standbyMACAddress", data.StandbyMacAddress.ValueString())
	}
	if len(data.ArpTableEntries) > 0 {
		body, _ = sjson.Set(body, "arpConfig.arpConfig", []interface{}{})
		for _, item := range data.ArpTableEntries {
			itemBody := ""
			if !item.MacAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "macAddress", item.MacAddress.ValueString())
			}
			if !item.IpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress", item.IpAddress.ValueString())
			}
			if !item.EnableAlias.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableAlias", item.EnableAlias.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "arpConfig.arpConfig.-1", itemBody)
		}
	}
	if !data.EnableAntiSpoofing.IsNull() {
		body, _ = sjson.Set(body, "enableAntiSpoofing", data.EnableAntiSpoofing.ValueBool())
	}
	if !data.AllowFullFragmentReassembly.IsNull() {
		body, _ = sjson.Set(body, "fragmentReassembly", data.AllowFullFragmentReassembly.ValueBool())
	}
	if !data.OverrideDefaultFragmentSettingChain.IsNull() {
		body, _ = sjson.Set(body, "overrideDefaultFragmentSetting.chain", data.OverrideDefaultFragmentSettingChain.ValueInt64())
	}
	if !data.OverrideDefaultFragmentSettingSize.IsNull() {
		body, _ = sjson.Set(body, "overrideDefaultFragmentSetting.size", data.OverrideDefaultFragmentSettingSize.ValueInt64())
	}
	if !data.OverrideDefaultFragmentSettingTimeout.IsNull() {
		body, _ = sjson.Set(body, "overrideDefaultFragmentSetting.timeout", data.OverrideDefaultFragmentSettingTimeout.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DevicePhysicalInterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ifname"); value.Exists() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolValue(true)
	}
	if value := res.Get("managementOnly"); value.Exists() {
		data.ManagementOnly = types.BoolValue(value.Bool())
	} else {
		data.ManagementOnly = types.BoolNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("mode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("securityZone.id"); value.Exists() {
		data.SecurityZoneId = types.StringValue(value.String())
	} else {
		data.SecurityZoneId = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("MTU"); value.Exists() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get("priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("enableSGTPropagate"); value.Exists() {
		data.EnableSgtPropagate = types.BoolValue(value.Bool())
	} else {
		data.EnableSgtPropagate = types.BoolNull()
	}
	if value := res.Get("nveOnly"); value.Exists() {
		data.NveOnly = types.BoolValue(value.Bool())
	} else {
		data.NveOnly = types.BoolNull()
	}
	if value := res.Get("ipv4.static.address"); value.Exists() {
		data.Ipv4StaticAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StaticAddress = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() {
		data.Ipv4StaticNetmask = types.StringValue(value.String())
	} else {
		data.Ipv4StaticNetmask = types.StringNull()
	}
	if value := res.Get("ipv4.dhcp.enableDefaultRouteDHCP"); value.Exists() {
		data.Ipv4DhcpObtainRoute = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DhcpObtainRoute = types.BoolNull()
	}
	if value := res.Get("ipv4.dhcp.dhcpRouteMetric"); value.Exists() {
		data.Ipv4DhcpRouteMetric = types.Int64Value(value.Int())
	} else {
		data.Ipv4DhcpRouteMetric = types.Int64Null()
	}
	if value := res.Get("ipv4.pppoe.vpdnGroupName"); value.Exists() {
		data.Ipv4PppoeVpdnGroupName = types.StringValue(value.String())
	} else {
		data.Ipv4PppoeVpdnGroupName = types.StringNull()
	}
	if value := res.Get("ipv4.pppoe.pppoeUser"); value.Exists() {
		data.Ipv4PppoeUser = types.StringValue(value.String())
	} else {
		data.Ipv4PppoeUser = types.StringNull()
	}
	if value := res.Get("ipv4.pppoe.pppoePassword"); value.Exists() {
		data.Ipv4PppoePassword = types.StringValue(value.String())
	} else {
		data.Ipv4PppoePassword = types.StringNull()
	}
	if value := res.Get("ipv4.pppoe.pppAuth"); value.Exists() {
		data.Ipv4PppoeAuthentication = types.StringValue(value.String())
	} else {
		data.Ipv4PppoeAuthentication = types.StringNull()
	}
	if value := res.Get("ipv4.pppoe.pppoeRouteMetric"); value.Exists() {
		data.Ipv4PppoeRouteMetric = types.Int64Value(value.Int())
	} else {
		data.Ipv4PppoeRouteMetric = types.Int64Null()
	}
	if value := res.Get("ipv4.pppoe.enableRouteSettings"); value.Exists() {
		data.Ipv4PppoeRouteSettings = types.BoolValue(value.Bool())
	} else {
		data.Ipv4PppoeRouteSettings = types.BoolNull()
	}
	if value := res.Get("ipv4.pppoe.storeCredsInFlash"); value.Exists() {
		data.Ipv4PppoeStoreCredentialsInFlash = types.BoolValue(value.Bool())
	} else {
		data.Ipv4PppoeStoreCredentialsInFlash = types.BoolNull()
	}
	if value := res.Get("ipv6.enableIPV6"); value.Exists() {
		data.Ipv6Enable = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enable = types.BoolNull()
	}
	if value := res.Get("ipv6.enforceEUI64"); value.Exists() {
		data.Ipv6EnforceEui = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnforceEui = types.BoolNull()
	}
	if value := res.Get("linkLocalAddress"); value.Exists() {
		data.Ipv6LinkLocalAddress = types.StringValue(value.String())
	} else {
		data.Ipv6LinkLocalAddress = types.StringNull()
	}
	if value := res.Get("ipv6.enableAutoConfig"); value.Exists() {
		data.Ipv6EnableAutoConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableAutoConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.addresses"); value.Exists() {
		data.Ipv6Addresses = make([]DevicePhysicalInterfaceIpv6Addresses, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DevicePhysicalInterfaceIpv6Addresses{}
			if value := res.Get("address"); value.Exists() {
				data.Address = types.StringValue(value.String())
			} else {
				data.Address = types.StringNull()
			}
			if value := res.Get("prefix"); value.Exists() {
				data.Prefix = types.StringValue(value.String())
			} else {
				data.Prefix = types.StringNull()
			}
			if value := res.Get("enforceEUI64"); value.Exists() {
				data.EnforceEui = types.BoolValue(value.Bool())
			} else {
				data.EnforceEui = types.BoolNull()
			}
			(*parent).Ipv6Addresses = append((*parent).Ipv6Addresses, data)
			return true
		})
	}
	if value := res.Get("ipv6.prefixes"); value.Exists() {
		data.Ipv6Prefixes = make([]DevicePhysicalInterfaceIpv6Prefixes, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DevicePhysicalInterfaceIpv6Prefixes{}
			if value := res.Get("address"); value.Exists() {
				data.Address = types.StringValue(value.String())
			} else {
				data.Address = types.StringNull()
			}
			if value := res.Get("default"); value.Exists() {
				data.Default = types.StringValue(value.String())
			} else {
				data.Default = types.StringNull()
			}
			if value := res.Get("enforceEUI64"); value.Exists() {
				data.EnforceEui = types.BoolValue(value.Bool())
			} else {
				data.EnforceEui = types.BoolNull()
			}
			(*parent).Ipv6Prefixes = append((*parent).Ipv6Prefixes, data)
			return true
		})
	}
	if value := res.Get("ipv6.enableDADLoopback"); value.Exists() {
		data.Ipv6EnableDad = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDad = types.BoolNull()
	}
	if value := res.Get("ipv6.dadAttempts"); value.Exists() {
		data.Ipv6DadAttempts = types.Int64Value(value.Int())
	} else {
		data.Ipv6DadAttempts = types.Int64Null()
	}
	if value := res.Get("ipv6.nsInterval"); value.Exists() {
		data.Ipv6NsInterval = types.Int64Value(value.Int())
	} else {
		data.Ipv6NsInterval = types.Int64Null()
	}
	if value := res.Get("ipv6.reachableTime"); value.Exists() {
		data.Ipv6ReachableTime = types.Int64Value(value.Int())
	} else {
		data.Ipv6ReachableTime = types.Int64Null()
	}
	if value := res.Get("ipv6.enableRA"); value.Exists() {
		data.Ipv6EnableRa = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableRa = types.BoolNull()
	}
	if value := res.Get("ipv6.raLifeTime"); value.Exists() {
		data.Ipv6RaLifeTime = types.Int64Value(value.Int())
	} else {
		data.Ipv6RaLifeTime = types.Int64Null()
	}
	if value := res.Get("ipv6.raInterval"); value.Exists() {
		data.Ipv6RaInterval = types.Int64Value(value.Int())
	} else {
		data.Ipv6RaInterval = types.Int64Null()
	}
	if value := res.Get("ipv6.DHCP.enableDHCPClient"); value.Exists() {
		data.Ipv6Dhcp = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Dhcp = types.BoolNull()
	}
	if value := res.Get("ipv6.DHCP.obtainIPV6DefaultRouteDHCP"); value.Exists() {
		data.Ipv6DefaultRouteByDhcp = types.BoolValue(value.Bool())
	} else {
		data.Ipv6DefaultRouteByDhcp = types.BoolNull()
	}
	if value := res.Get("ipv6.ipv6DHCPPool.id"); value.Exists() {
		data.Ipv6DhcpPoolId = types.StringValue(value.String())
	} else {
		data.Ipv6DhcpPoolId = types.StringNull()
	}
	if value := res.Get("ipv6.ipv6DHCPPool.type"); value.Exists() {
		data.Ipv6DhcpPoolType = types.StringValue(value.String())
	} else {
		data.Ipv6DhcpPoolType = types.StringNull()
	}
	if value := res.Get("ipv6.enableDHCPAddrConfig"); value.Exists() {
		data.Ipv6EnableDhcpAddressConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpAddressConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPNonAddrConfig"); value.Exists() {
		data.Ipv6EnableDhcpNonaddressConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpNonaddressConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.DHCP.clientPd.prefixName"); value.Exists() {
		data.Ipv6DhcpClientPdPrefixName = types.StringValue(value.String())
	} else {
		data.Ipv6DhcpClientPdPrefixName = types.StringNull()
	}
	if value := res.Get("ipv6.DHCP.clientPd.hintPrefixes"); value.Exists() {
		data.Ipv6DhcpClientPdHintPrefixes = types.StringValue(value.String())
	} else {
		data.Ipv6DhcpClientPdHintPrefixes = types.StringNull()
	}
	if value := res.Get("pathMonitoring.enable"); value.Exists() {
		data.IpBasedMonitoring = types.BoolValue(value.Bool())
	} else {
		data.IpBasedMonitoring = types.BoolNull()
	}
	if value := res.Get("pathMonitoring.type"); value.Exists() {
		data.IpBasedMonitoringType = types.StringValue(value.String())
	} else {
		data.IpBasedMonitoringType = types.StringNull()
	}
	if value := res.Get("monitoredIp"); value.Exists() {
		data.IpBasedMonitoringNextHop = types.StringValue(value.String())
	} else {
		data.IpBasedMonitoringNextHop = types.StringNull()
	}
	if value := res.Get("hardware.autoNegState"); value.Exists() {
		data.AutoNegotiation = types.BoolValue(value.Bool())
	} else {
		data.AutoNegotiation = types.BoolNull()
	}
	if value := res.Get("hardware.duplex"); value.Exists() {
		data.Duplex = types.StringValue(value.String())
	} else {
		data.Duplex = types.StringNull()
	}
	if value := res.Get("hardware.speed"); value.Exists() {
		data.Speed = types.StringValue(value.String())
	} else {
		data.Speed = types.StringNull()
	}
	if value := res.Get("LLDP.receive"); value.Exists() {
		data.LldpReceive = types.BoolValue(value.Bool())
	} else {
		data.LldpReceive = types.BoolNull()
	}
	if value := res.Get("LLDP.transmit"); value.Exists() {
		data.LldpTransmit = types.BoolValue(value.Bool())
	} else {
		data.LldpTransmit = types.BoolNull()
	}
	if value := res.Get("hardware.flowControlSend"); value.Exists() {
		data.FlowControlSend = types.StringValue(value.String())
	} else {
		data.FlowControlSend = types.StringNull()
	}
	if value := res.Get("hardware.fecMode"); value.Exists() {
		data.FecMode = types.StringValue(value.String())
	} else {
		data.FecMode = types.StringNull()
	}
	if value := res.Get("fmcAccessConfig.enableAccess"); value.Exists() {
		data.ManagementAccess = types.BoolValue(value.Bool())
	} else {
		data.ManagementAccess = types.BoolNull()
	}
	if value := res.Get("fmcAccessConfig.allowedNetworks"); value.Exists() {
		data.ManagementAccessNetworkObjects = make([]DevicePhysicalInterfaceManagementAccessNetworkObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DevicePhysicalInterfaceManagementAccessNetworkObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).ManagementAccessNetworkObjects = append((*parent).ManagementAccessNetworkObjects, data)
			return true
		})
	}
	if value := res.Get("activeMACAddress"); value.Exists() {
		data.ActiveMacAddress = types.StringValue(value.String())
	} else {
		data.ActiveMacAddress = types.StringNull()
	}
	if value := res.Get("standbyMACAddress"); value.Exists() {
		data.StandbyMacAddress = types.StringValue(value.String())
	} else {
		data.StandbyMacAddress = types.StringNull()
	}
	if value := res.Get("arpConfig.arpConfig"); value.Exists() {
		data.ArpTableEntries = make([]DevicePhysicalInterfaceArpTableEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DevicePhysicalInterfaceArpTableEntries{}
			if value := res.Get("macAddress"); value.Exists() {
				data.MacAddress = types.StringValue(value.String())
			} else {
				data.MacAddress = types.StringNull()
			}
			if value := res.Get("ipAddress"); value.Exists() {
				data.IpAddress = types.StringValue(value.String())
			} else {
				data.IpAddress = types.StringNull()
			}
			if value := res.Get("enableAlias"); value.Exists() {
				data.EnableAlias = types.BoolValue(value.Bool())
			} else {
				data.EnableAlias = types.BoolNull()
			}
			(*parent).ArpTableEntries = append((*parent).ArpTableEntries, data)
			return true
		})
	}
	if value := res.Get("enableAntiSpoofing"); value.Exists() {
		data.EnableAntiSpoofing = types.BoolValue(value.Bool())
	} else {
		data.EnableAntiSpoofing = types.BoolNull()
	}
	if value := res.Get("fragmentReassembly"); value.Exists() {
		data.AllowFullFragmentReassembly = types.BoolValue(value.Bool())
	} else {
		data.AllowFullFragmentReassembly = types.BoolNull()
	}
	if value := res.Get("overrideDefaultFragmentSetting.chain"); value.Exists() {
		data.OverrideDefaultFragmentSettingChain = types.Int64Value(value.Int())
	} else {
		data.OverrideDefaultFragmentSettingChain = types.Int64Null()
	}
	if value := res.Get("overrideDefaultFragmentSetting.size"); value.Exists() {
		data.OverrideDefaultFragmentSettingSize = types.Int64Value(value.Int())
	} else {
		data.OverrideDefaultFragmentSettingSize = types.Int64Null()
	}
	if value := res.Get("overrideDefaultFragmentSetting.timeout"); value.Exists() {
		data.OverrideDefaultFragmentSettingTimeout = types.Int64Value(value.Int())
	} else {
		data.OverrideDefaultFragmentSettingTimeout = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DevicePhysicalInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ifname"); value.Exists() && !data.LogicalName.IsNull() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else if data.Enabled.ValueBool() != true {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("managementOnly"); value.Exists() && !data.ManagementOnly.IsNull() {
		data.ManagementOnly = types.BoolValue(value.Bool())
	} else {
		data.ManagementOnly = types.BoolNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("securityZone.id"); value.Exists() && !data.SecurityZoneId.IsNull() {
		data.SecurityZoneId = types.StringValue(value.String())
	} else {
		data.SecurityZoneId = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("MTU"); value.Exists() && !data.Mtu.IsNull() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("enableSGTPropagate"); value.Exists() && !data.EnableSgtPropagate.IsNull() {
		data.EnableSgtPropagate = types.BoolValue(value.Bool())
	} else {
		data.EnableSgtPropagate = types.BoolNull()
	}
	if value := res.Get("nveOnly"); value.Exists() && !data.NveOnly.IsNull() {
		data.NveOnly = types.BoolValue(value.Bool())
	} else {
		data.NveOnly = types.BoolNull()
	}
	if value := res.Get("ipv4.static.address"); value.Exists() && !data.Ipv4StaticAddress.IsNull() {
		data.Ipv4StaticAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StaticAddress = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() && !data.Ipv4StaticNetmask.IsNull() {
		data.Ipv4StaticNetmask = types.StringValue(value.String())
	} else {
		data.Ipv4StaticNetmask = types.StringNull()
	}
	if value := res.Get("ipv4.dhcp.enableDefaultRouteDHCP"); value.Exists() && !data.Ipv4DhcpObtainRoute.IsNull() {
		data.Ipv4DhcpObtainRoute = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DhcpObtainRoute = types.BoolNull()
	}
	if value := res.Get("ipv4.dhcp.dhcpRouteMetric"); value.Exists() && !data.Ipv4DhcpRouteMetric.IsNull() {
		data.Ipv4DhcpRouteMetric = types.Int64Value(value.Int())
	} else {
		data.Ipv4DhcpRouteMetric = types.Int64Null()
	}
	if value := res.Get("ipv4.pppoe.vpdnGroupName"); value.Exists() && !data.Ipv4PppoeVpdnGroupName.IsNull() {
		data.Ipv4PppoeVpdnGroupName = types.StringValue(value.String())
	} else {
		data.Ipv4PppoeVpdnGroupName = types.StringNull()
	}
	if value := res.Get("ipv4.pppoe.pppoeUser"); value.Exists() && !data.Ipv4PppoeUser.IsNull() {
		data.Ipv4PppoeUser = types.StringValue(value.String())
	} else {
		data.Ipv4PppoeUser = types.StringNull()
	}
	if value := res.Get("ipv4.pppoe.pppoePassword"); value.Exists() && !data.Ipv4PppoePassword.IsNull() {
		data.Ipv4PppoePassword = types.StringValue(value.String())
	} else {
		data.Ipv4PppoePassword = types.StringNull()
	}
	if value := res.Get("ipv4.pppoe.pppAuth"); value.Exists() && !data.Ipv4PppoeAuthentication.IsNull() {
		data.Ipv4PppoeAuthentication = types.StringValue(value.String())
	} else {
		data.Ipv4PppoeAuthentication = types.StringNull()
	}
	if value := res.Get("ipv4.pppoe.pppoeRouteMetric"); value.Exists() && !data.Ipv4PppoeRouteMetric.IsNull() {
		data.Ipv4PppoeRouteMetric = types.Int64Value(value.Int())
	} else {
		data.Ipv4PppoeRouteMetric = types.Int64Null()
	}
	if value := res.Get("ipv4.pppoe.enableRouteSettings"); value.Exists() && !data.Ipv4PppoeRouteSettings.IsNull() {
		data.Ipv4PppoeRouteSettings = types.BoolValue(value.Bool())
	} else {
		data.Ipv4PppoeRouteSettings = types.BoolNull()
	}
	if value := res.Get("ipv4.pppoe.storeCredsInFlash"); value.Exists() && !data.Ipv4PppoeStoreCredentialsInFlash.IsNull() {
		data.Ipv4PppoeStoreCredentialsInFlash = types.BoolValue(value.Bool())
	} else {
		data.Ipv4PppoeStoreCredentialsInFlash = types.BoolNull()
	}
	if value := res.Get("ipv6.enableIPV6"); value.Exists() && !data.Ipv6Enable.IsNull() {
		data.Ipv6Enable = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enable = types.BoolNull()
	}
	if value := res.Get("ipv6.enforceEUI64"); value.Exists() && !data.Ipv6EnforceEui.IsNull() {
		data.Ipv6EnforceEui = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnforceEui = types.BoolNull()
	}
	if value := res.Get("linkLocalAddress"); value.Exists() && !data.Ipv6LinkLocalAddress.IsNull() {
		data.Ipv6LinkLocalAddress = types.StringValue(value.String())
	} else {
		data.Ipv6LinkLocalAddress = types.StringNull()
	}
	if value := res.Get("ipv6.enableAutoConfig"); value.Exists() && !data.Ipv6EnableAutoConfig.IsNull() {
		data.Ipv6EnableAutoConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableAutoConfig = types.BoolNull()
	}
	for i := 0; i < len(data.Ipv6Addresses); i++ {
		keys := [...]string{"address", "prefix"}
		keyValues := [...]string{data.Ipv6Addresses[i].Address.ValueString(), data.Ipv6Addresses[i].Prefix.ValueString()}

		parent := &data
		data := (*parent).Ipv6Addresses[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ipv6.addresses").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv6Addresses[%d] = %+v",
				i,
				(*parent).Ipv6Addresses[i],
			))
			(*parent).Ipv6Addresses = slices.Delete((*parent).Ipv6Addresses, i, i+1)
			i--

			continue
		}
		if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("prefix"); value.Exists() && !data.Prefix.IsNull() {
			data.Prefix = types.StringValue(value.String())
		} else {
			data.Prefix = types.StringNull()
		}
		if value := res.Get("enforceEUI64"); value.Exists() && !data.EnforceEui.IsNull() {
			data.EnforceEui = types.BoolValue(value.Bool())
		} else {
			data.EnforceEui = types.BoolNull()
		}
		(*parent).Ipv6Addresses[i] = data
	}
	for i := 0; i < len(data.Ipv6Prefixes); i++ {
		keys := [...]string{"address", "default"}
		keyValues := [...]string{data.Ipv6Prefixes[i].Address.ValueString(), data.Ipv6Prefixes[i].Default.ValueString()}

		parent := &data
		data := (*parent).Ipv6Prefixes[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ipv6.prefixes").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv6Prefixes[%d] = %+v",
				i,
				(*parent).Ipv6Prefixes[i],
			))
			(*parent).Ipv6Prefixes = slices.Delete((*parent).Ipv6Prefixes, i, i+1)
			i--

			continue
		}
		if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("default"); value.Exists() && !data.Default.IsNull() {
			data.Default = types.StringValue(value.String())
		} else {
			data.Default = types.StringNull()
		}
		if value := res.Get("enforceEUI64"); value.Exists() && !data.EnforceEui.IsNull() {
			data.EnforceEui = types.BoolValue(value.Bool())
		} else {
			data.EnforceEui = types.BoolNull()
		}
		(*parent).Ipv6Prefixes[i] = data
	}
	if value := res.Get("ipv6.enableDADLoopback"); value.Exists() && !data.Ipv6EnableDad.IsNull() {
		data.Ipv6EnableDad = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDad = types.BoolNull()
	}
	if value := res.Get("ipv6.dadAttempts"); value.Exists() && !data.Ipv6DadAttempts.IsNull() {
		data.Ipv6DadAttempts = types.Int64Value(value.Int())
	} else {
		data.Ipv6DadAttempts = types.Int64Null()
	}
	if value := res.Get("ipv6.nsInterval"); value.Exists() && !data.Ipv6NsInterval.IsNull() {
		data.Ipv6NsInterval = types.Int64Value(value.Int())
	} else {
		data.Ipv6NsInterval = types.Int64Null()
	}
	if value := res.Get("ipv6.reachableTime"); value.Exists() && !data.Ipv6ReachableTime.IsNull() {
		data.Ipv6ReachableTime = types.Int64Value(value.Int())
	} else {
		data.Ipv6ReachableTime = types.Int64Null()
	}
	if value := res.Get("ipv6.enableRA"); value.Exists() && !data.Ipv6EnableRa.IsNull() {
		data.Ipv6EnableRa = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableRa = types.BoolNull()
	}
	if value := res.Get("ipv6.raLifeTime"); value.Exists() && !data.Ipv6RaLifeTime.IsNull() {
		data.Ipv6RaLifeTime = types.Int64Value(value.Int())
	} else {
		data.Ipv6RaLifeTime = types.Int64Null()
	}
	if value := res.Get("ipv6.raInterval"); value.Exists() && !data.Ipv6RaInterval.IsNull() {
		data.Ipv6RaInterval = types.Int64Value(value.Int())
	} else {
		data.Ipv6RaInterval = types.Int64Null()
	}
	if value := res.Get("ipv6.DHCP.enableDHCPClient"); value.Exists() && !data.Ipv6Dhcp.IsNull() {
		data.Ipv6Dhcp = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Dhcp = types.BoolNull()
	}
	if value := res.Get("ipv6.DHCP.obtainIPV6DefaultRouteDHCP"); value.Exists() && !data.Ipv6DefaultRouteByDhcp.IsNull() {
		data.Ipv6DefaultRouteByDhcp = types.BoolValue(value.Bool())
	} else {
		data.Ipv6DefaultRouteByDhcp = types.BoolNull()
	}
	if value := res.Get("ipv6.ipv6DHCPPool.id"); value.Exists() && !data.Ipv6DhcpPoolId.IsNull() {
		data.Ipv6DhcpPoolId = types.StringValue(value.String())
	} else {
		data.Ipv6DhcpPoolId = types.StringNull()
	}
	if value := res.Get("ipv6.ipv6DHCPPool.type"); value.Exists() && !data.Ipv6DhcpPoolType.IsNull() {
		data.Ipv6DhcpPoolType = types.StringValue(value.String())
	} else {
		data.Ipv6DhcpPoolType = types.StringNull()
	}
	if value := res.Get("ipv6.enableDHCPAddrConfig"); value.Exists() && !data.Ipv6EnableDhcpAddressConfig.IsNull() {
		data.Ipv6EnableDhcpAddressConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpAddressConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPNonAddrConfig"); value.Exists() && !data.Ipv6EnableDhcpNonaddressConfig.IsNull() {
		data.Ipv6EnableDhcpNonaddressConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpNonaddressConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.DHCP.clientPd.prefixName"); value.Exists() && !data.Ipv6DhcpClientPdPrefixName.IsNull() {
		data.Ipv6DhcpClientPdPrefixName = types.StringValue(value.String())
	} else {
		data.Ipv6DhcpClientPdPrefixName = types.StringNull()
	}
	if value := res.Get("ipv6.DHCP.clientPd.hintPrefixes"); value.Exists() && !data.Ipv6DhcpClientPdHintPrefixes.IsNull() {
		data.Ipv6DhcpClientPdHintPrefixes = types.StringValue(value.String())
	} else {
		data.Ipv6DhcpClientPdHintPrefixes = types.StringNull()
	}
	if value := res.Get("pathMonitoring.enable"); value.Exists() && !data.IpBasedMonitoring.IsNull() {
		data.IpBasedMonitoring = types.BoolValue(value.Bool())
	} else {
		data.IpBasedMonitoring = types.BoolNull()
	}
	if value := res.Get("pathMonitoring.type"); value.Exists() && !data.IpBasedMonitoringType.IsNull() {
		data.IpBasedMonitoringType = types.StringValue(value.String())
	} else {
		data.IpBasedMonitoringType = types.StringNull()
	}
	if value := res.Get("monitoredIp"); value.Exists() && !data.IpBasedMonitoringNextHop.IsNull() {
		data.IpBasedMonitoringNextHop = types.StringValue(value.String())
	} else {
		data.IpBasedMonitoringNextHop = types.StringNull()
	}
	if value := res.Get("hardware.autoNegState"); value.Exists() && !data.AutoNegotiation.IsNull() {
		data.AutoNegotiation = types.BoolValue(value.Bool())
	} else {
		data.AutoNegotiation = types.BoolNull()
	}
	if value := res.Get("hardware.duplex"); value.Exists() && !data.Duplex.IsNull() {
		data.Duplex = types.StringValue(value.String())
	} else {
		data.Duplex = types.StringNull()
	}
	if value := res.Get("hardware.speed"); value.Exists() && !data.Speed.IsNull() {
		data.Speed = types.StringValue(value.String())
	} else {
		data.Speed = types.StringNull()
	}
	if value := res.Get("LLDP.receive"); value.Exists() && !data.LldpReceive.IsNull() {
		data.LldpReceive = types.BoolValue(value.Bool())
	} else {
		data.LldpReceive = types.BoolNull()
	}
	if value := res.Get("LLDP.transmit"); value.Exists() && !data.LldpTransmit.IsNull() {
		data.LldpTransmit = types.BoolValue(value.Bool())
	} else {
		data.LldpTransmit = types.BoolNull()
	}
	if value := res.Get("hardware.flowControlSend"); value.Exists() && !data.FlowControlSend.IsNull() {
		data.FlowControlSend = types.StringValue(value.String())
	} else {
		data.FlowControlSend = types.StringNull()
	}
	if value := res.Get("hardware.fecMode"); value.Exists() && !data.FecMode.IsNull() {
		data.FecMode = types.StringValue(value.String())
	} else {
		data.FecMode = types.StringNull()
	}
	if value := res.Get("fmcAccessConfig.enableAccess"); value.Exists() && !data.ManagementAccess.IsNull() {
		data.ManagementAccess = types.BoolValue(value.Bool())
	} else {
		data.ManagementAccess = types.BoolNull()
	}
	for i := 0; i < len(data.ManagementAccessNetworkObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.ManagementAccessNetworkObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).ManagementAccessNetworkObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("fmcAccessConfig.allowedNetworks").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ManagementAccessNetworkObjects[%d] = %+v",
				i,
				(*parent).ManagementAccessNetworkObjects[i],
			))
			(*parent).ManagementAccessNetworkObjects = slices.Delete((*parent).ManagementAccessNetworkObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).ManagementAccessNetworkObjects[i] = data
	}
	if value := res.Get("activeMACAddress"); value.Exists() && !data.ActiveMacAddress.IsNull() {
		data.ActiveMacAddress = types.StringValue(value.String())
	} else {
		data.ActiveMacAddress = types.StringNull()
	}
	if value := res.Get("standbyMACAddress"); value.Exists() && !data.StandbyMacAddress.IsNull() {
		data.StandbyMacAddress = types.StringValue(value.String())
	} else {
		data.StandbyMacAddress = types.StringNull()
	}
	for i := 0; i < len(data.ArpTableEntries); i++ {
		keys := [...]string{"macAddress", "ipAddress", "enableAlias"}
		keyValues := [...]string{data.ArpTableEntries[i].MacAddress.ValueString(), data.ArpTableEntries[i].IpAddress.ValueString(), strconv.FormatBool(data.ArpTableEntries[i].EnableAlias.ValueBool())}

		parent := &data
		data := (*parent).ArpTableEntries[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("arpConfig.arpConfig").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ArpTableEntries[%d] = %+v",
				i,
				(*parent).ArpTableEntries[i],
			))
			(*parent).ArpTableEntries = slices.Delete((*parent).ArpTableEntries, i, i+1)
			i--

			continue
		}
		if value := res.Get("macAddress"); value.Exists() && !data.MacAddress.IsNull() {
			data.MacAddress = types.StringValue(value.String())
		} else {
			data.MacAddress = types.StringNull()
		}
		if value := res.Get("ipAddress"); value.Exists() && !data.IpAddress.IsNull() {
			data.IpAddress = types.StringValue(value.String())
		} else {
			data.IpAddress = types.StringNull()
		}
		if value := res.Get("enableAlias"); value.Exists() && !data.EnableAlias.IsNull() {
			data.EnableAlias = types.BoolValue(value.Bool())
		} else {
			data.EnableAlias = types.BoolNull()
		}
		(*parent).ArpTableEntries[i] = data
	}
	if value := res.Get("enableAntiSpoofing"); value.Exists() && !data.EnableAntiSpoofing.IsNull() {
		data.EnableAntiSpoofing = types.BoolValue(value.Bool())
	} else {
		data.EnableAntiSpoofing = types.BoolNull()
	}
	if value := res.Get("fragmentReassembly"); value.Exists() && !data.AllowFullFragmentReassembly.IsNull() {
		data.AllowFullFragmentReassembly = types.BoolValue(value.Bool())
	} else {
		data.AllowFullFragmentReassembly = types.BoolNull()
	}
	if value := res.Get("overrideDefaultFragmentSetting.chain"); value.Exists() && !data.OverrideDefaultFragmentSettingChain.IsNull() {
		data.OverrideDefaultFragmentSettingChain = types.Int64Value(value.Int())
	} else {
		data.OverrideDefaultFragmentSettingChain = types.Int64Null()
	}
	if value := res.Get("overrideDefaultFragmentSetting.size"); value.Exists() && !data.OverrideDefaultFragmentSettingSize.IsNull() {
		data.OverrideDefaultFragmentSettingSize = types.Int64Value(value.Int())
	} else {
		data.OverrideDefaultFragmentSettingSize = types.Int64Null()
	}
	if value := res.Get("overrideDefaultFragmentSetting.timeout"); value.Exists() && !data.OverrideDefaultFragmentSettingTimeout.IsNull() {
		data.OverrideDefaultFragmentSettingTimeout = types.Int64Value(value.Int())
	} else {
		data.OverrideDefaultFragmentSettingTimeout = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DevicePhysicalInterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// toBodyPutDelete generates minimal required body to reset the resource to its default state.
func (data DevicePhysicalInterface) toBodyPutDelete(ctx context.Context, state DevicePhysicalInterface) string {
	body := ""
	body, _ = sjson.Set(body, "mode", "NONE")
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "ifname", data.LogicalName.ValueString())
	}

	return body
}
