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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceHAPair struct {
	Id                            types.String `tfsdk:"id"`
	Domain                        types.String `tfsdk:"domain"`
	Name                          types.String `tfsdk:"name"`
	Type                          types.String `tfsdk:"type"`
	PrimaryDeviceId               types.String `tfsdk:"primary_device_id"`
	SecondaryDeviceId             types.String `tfsdk:"secondary_device_id"`
	HaLinkInterfaceId             types.String `tfsdk:"ha_link_interface_id"`
	HaLinkInterfaceName           types.String `tfsdk:"ha_link_interface_name"`
	HaLinkInterfaceType           types.String `tfsdk:"ha_link_interface_type"`
	HaLinkLogicalName             types.String `tfsdk:"ha_link_logical_name"`
	HaLinkUseIpv6                 types.Bool   `tfsdk:"ha_link_use_ipv6"`
	HaLinkPrimaryIp               types.String `tfsdk:"ha_link_primary_ip"`
	HaLinkSecondaryIp             types.String `tfsdk:"ha_link_secondary_ip"`
	HaLinkNetmask                 types.String `tfsdk:"ha_link_netmask"`
	StateLinkUseSameAsHa          types.Bool   `tfsdk:"state_link_use_same_as_ha"`
	StateLinkInterfaceId          types.String `tfsdk:"state_link_interface_id"`
	StateLinkInterfaceName        types.String `tfsdk:"state_link_interface_name"`
	StateLinkInterfaceType        types.String `tfsdk:"state_link_interface_type"`
	StateLinkLogicalName          types.String `tfsdk:"state_link_logical_name"`
	StateLinkUseIpv6              types.Bool   `tfsdk:"state_link_use_ipv6"`
	StateLinkPrimaryIp            types.String `tfsdk:"state_link_primary_ip"`
	StateLinkSecondaryIp          types.String `tfsdk:"state_link_secondary_ip"`
	StateLinkNetmask              types.String `tfsdk:"state_link_netmask"`
	EncryptionEnabled             types.Bool   `tfsdk:"encryption_enabled"`
	EncryptionKeyGenerationScheme types.String `tfsdk:"encryption_key_generation_scheme"`
	EncryptionKey                 types.String `tfsdk:"encryption_key"`
	FailedInterfacesPercent       types.Int64  `tfsdk:"failed_interfaces_percent"`
	FailedInterfacesLimit         types.Int64  `tfsdk:"failed_interfaces_limit"`
	PeerPollTime                  types.Int64  `tfsdk:"peer_poll_time"`
	PeerPollTimeUnit              types.String `tfsdk:"peer_poll_time_unit"`
	PeerHoldTime                  types.Int64  `tfsdk:"peer_hold_time"`
	PeerHoldTimeUnit              types.String `tfsdk:"peer_hold_time_unit"`
	InterfacePollTime             types.Int64  `tfsdk:"interface_poll_time"`
	InterfacePollTimeUnit         types.String `tfsdk:"interface_poll_time_unit"`
	InterfaceHoldTime             types.Int64  `tfsdk:"interface_hold_time"`
	Action                        types.String `tfsdk:"action"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceHAPair) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicehapairs/ftddevicehapairs"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceHAPair) toBody(ctx context.Context, state DeviceHAPair) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.PrimaryDeviceId.IsNull() {
		body, _ = sjson.Set(body, "primary.id", data.PrimaryDeviceId.ValueString())
	}
	if !data.SecondaryDeviceId.IsNull() {
		body, _ = sjson.Set(body, "secondary.id", data.SecondaryDeviceId.ValueString())
	}
	if !data.HaLinkInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.interfaceObject.id", data.HaLinkInterfaceId.ValueString())
	}
	if !data.HaLinkInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.interfaceObject.name", data.HaLinkInterfaceName.ValueString())
	}
	if !data.HaLinkInterfaceType.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.interfaceObject.type", data.HaLinkInterfaceType.ValueString())
	}
	if !data.HaLinkLogicalName.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.logicalName", data.HaLinkLogicalName.ValueString())
	}
	if !data.HaLinkUseIpv6.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.useIPv6Address", data.HaLinkUseIpv6.ValueBool())
	}
	if !data.HaLinkPrimaryIp.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.activeIP", data.HaLinkPrimaryIp.ValueString())
	}
	if !data.HaLinkSecondaryIp.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.standbyIP", data.HaLinkSecondaryIp.ValueString())
	}
	if !data.HaLinkNetmask.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.subnetMask", data.HaLinkNetmask.ValueString())
	}
	if !data.StateLinkUseSameAsHa.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.useSameLinkForFailovers", data.StateLinkUseSameAsHa.ValueBool())
	}
	if !data.StateLinkInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.interfaceObject.id", data.StateLinkInterfaceId.ValueString())
	}
	if !data.StateLinkInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.interfaceObject.name", data.StateLinkInterfaceName.ValueString())
	}
	if !data.StateLinkInterfaceType.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.interfaceObject.type", data.StateLinkInterfaceType.ValueString())
	}
	if !data.StateLinkLogicalName.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.logicalName", data.StateLinkLogicalName.ValueString())
	}
	if !data.StateLinkUseIpv6.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.useIPv6Address", data.StateLinkUseIpv6.ValueBool())
	}
	if !data.StateLinkPrimaryIp.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.activeIP", data.StateLinkPrimaryIp.ValueString())
	}
	if !data.StateLinkSecondaryIp.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.standbyIP", data.StateLinkSecondaryIp.ValueString())
	}
	if !data.StateLinkNetmask.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.subnetMask", data.StateLinkNetmask.ValueString())
	}
	if !data.EncryptionEnabled.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.isEncryptionEnabled", data.EncryptionEnabled.ValueBool())
	}
	if !data.EncryptionKeyGenerationScheme.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.encKeyGenerationScheme", data.EncryptionKeyGenerationScheme.ValueString())
	}
	if !data.EncryptionKey.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.sharedKey", data.EncryptionKey.ValueString())
	}
	if !data.FailedInterfacesPercent.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.percentFailedInterfaceExceed", data.FailedInterfacesPercent.ValueInt64())
	}
	if !data.FailedInterfacesLimit.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.noOfFailedInterfaceLimit", data.FailedInterfacesLimit.ValueInt64())
	}
	if !data.PeerPollTime.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.peerPollTime", data.PeerPollTime.ValueInt64())
	}
	if !data.PeerPollTimeUnit.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.peerPollTimeUnit", data.PeerPollTimeUnit.ValueString())
	}
	if !data.PeerHoldTime.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.peerHoldTime", data.PeerHoldTime.ValueInt64())
	}
	if !data.PeerHoldTimeUnit.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.peerHoldTimeUnit", data.PeerHoldTimeUnit.ValueString())
	}
	if !data.InterfacePollTime.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.interfacePollTime", data.InterfacePollTime.ValueInt64())
	}
	if !data.InterfacePollTimeUnit.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.interfacePollTimeUnit", data.InterfacePollTimeUnit.ValueString())
	}
	if !data.InterfaceHoldTime.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.interfaceHoldTime", data.InterfaceHoldTime.ValueInt64())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, "action", data.Action.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceHAPair) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("primary.id"); value.Exists() {
		data.PrimaryDeviceId = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceId = types.StringNull()
	}
	if value := res.Get("secondary.id"); value.Exists() {
		data.SecondaryDeviceId = types.StringValue(value.String())
	} else {
		data.SecondaryDeviceId = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.interfaceObject.name"); value.Exists() {
		data.HaLinkInterfaceName = types.StringValue(value.String())
	} else {
		data.HaLinkInterfaceName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.logicalName"); value.Exists() {
		data.HaLinkLogicalName = types.StringValue(value.String())
	} else {
		data.HaLinkLogicalName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.useIPv6Address"); value.Exists() {
		data.HaLinkUseIpv6 = types.BoolValue(value.Bool())
	} else {
		data.HaLinkUseIpv6 = types.BoolValue(false)
	}
	if value := res.Get("ftdHABootstrap.lanFailover.activeIP"); value.Exists() {
		data.HaLinkPrimaryIp = types.StringValue(value.String())
	} else {
		data.HaLinkPrimaryIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.standbyIP"); value.Exists() {
		data.HaLinkSecondaryIp = types.StringValue(value.String())
	} else {
		data.HaLinkSecondaryIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.subnetMask"); value.Exists() {
		data.HaLinkNetmask = types.StringValue(value.String())
	} else {
		data.HaLinkNetmask = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.interfaceObject.name"); value.Exists() {
		data.StateLinkInterfaceName = types.StringValue(value.String())
	} else {
		data.StateLinkInterfaceName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.logicalName"); value.Exists() {
		data.StateLinkLogicalName = types.StringValue(value.String())
	} else {
		data.StateLinkLogicalName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.useIPv6Address"); value.Exists() {
		data.StateLinkUseIpv6 = types.BoolValue(value.Bool())
	} else {
		data.StateLinkUseIpv6 = types.BoolNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.activeIP"); value.Exists() {
		data.StateLinkPrimaryIp = types.StringValue(value.String())
	} else {
		data.StateLinkPrimaryIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.standbyIP"); value.Exists() {
		data.StateLinkSecondaryIp = types.StringValue(value.String())
	} else {
		data.StateLinkSecondaryIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.subnetMask"); value.Exists() {
		data.StateLinkNetmask = types.StringValue(value.String())
	} else {
		data.StateLinkNetmask = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.isEncryptionEnabled"); value.Exists() {
		data.EncryptionEnabled = types.BoolValue(value.Bool())
	} else {
		data.EncryptionEnabled = types.BoolNull()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.percentFailedInterfaceExceed"); value.Exists() {
		data.FailedInterfacesPercent = types.Int64Value(value.Int())
	} else {
		data.FailedInterfacesPercent = types.Int64Null()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.noOfFailedInterfaceLimit"); value.Exists() {
		data.FailedInterfacesLimit = types.Int64Value(value.Int())
	} else {
		data.FailedInterfacesLimit = types.Int64Null()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.peerPollTime"); value.Exists() {
		data.PeerPollTime = types.Int64Value(value.Int())
	} else {
		data.PeerPollTime = types.Int64Value(1)
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.peerPollTimeUnit"); value.Exists() {
		data.PeerPollTimeUnit = types.StringValue(value.String())
	} else {
		data.PeerPollTimeUnit = types.StringValue("SEC")
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.peerHoldTime"); value.Exists() {
		data.PeerHoldTime = types.Int64Value(value.Int())
	} else {
		data.PeerHoldTime = types.Int64Value(15)
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.peerHoldTimeUnit"); value.Exists() {
		data.PeerHoldTimeUnit = types.StringValue(value.String())
	} else {
		data.PeerHoldTimeUnit = types.StringValue("SEC")
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.interfacePollTime"); value.Exists() {
		data.InterfacePollTime = types.Int64Value(value.Int())
	} else {
		data.InterfacePollTime = types.Int64Value(5)
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.interfacePollTimeUnit"); value.Exists() {
		data.InterfacePollTimeUnit = types.StringValue(value.String())
	} else {
		data.InterfacePollTimeUnit = types.StringValue("SEC")
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.interfaceHoldTime"); value.Exists() {
		data.InterfaceHoldTime = types.Int64Value(value.Int())
	} else {
		data.InterfaceHoldTime = types.Int64Value(25)
	}
	if value := res.Get("action"); value.Exists() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceHAPair) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("primary.id"); value.Exists() && !data.PrimaryDeviceId.IsNull() {
		data.PrimaryDeviceId = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceId = types.StringNull()
	}
	if value := res.Get("secondary.id"); value.Exists() && !data.SecondaryDeviceId.IsNull() {
		data.SecondaryDeviceId = types.StringValue(value.String())
	} else {
		data.SecondaryDeviceId = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.interfaceObject.name"); value.Exists() && !data.HaLinkInterfaceName.IsNull() {
		data.HaLinkInterfaceName = types.StringValue(value.String())
	} else {
		data.HaLinkInterfaceName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.logicalName"); value.Exists() && !data.HaLinkLogicalName.IsNull() {
		data.HaLinkLogicalName = types.StringValue(value.String())
	} else {
		data.HaLinkLogicalName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.useIPv6Address"); value.Exists() && !data.HaLinkUseIpv6.IsNull() {
		data.HaLinkUseIpv6 = types.BoolValue(value.Bool())
	} else if data.HaLinkUseIpv6.ValueBool() != false {
		data.HaLinkUseIpv6 = types.BoolNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.activeIP"); value.Exists() && !data.HaLinkPrimaryIp.IsNull() {
		data.HaLinkPrimaryIp = types.StringValue(value.String())
	} else {
		data.HaLinkPrimaryIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.standbyIP"); value.Exists() && !data.HaLinkSecondaryIp.IsNull() {
		data.HaLinkSecondaryIp = types.StringValue(value.String())
	} else {
		data.HaLinkSecondaryIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.subnetMask"); value.Exists() && !data.HaLinkNetmask.IsNull() {
		data.HaLinkNetmask = types.StringValue(value.String())
	} else {
		data.HaLinkNetmask = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.interfaceObject.name"); value.Exists() && !data.StateLinkInterfaceName.IsNull() {
		data.StateLinkInterfaceName = types.StringValue(value.String())
	} else {
		data.StateLinkInterfaceName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.logicalName"); value.Exists() && !data.StateLinkLogicalName.IsNull() {
		data.StateLinkLogicalName = types.StringValue(value.String())
	} else {
		data.StateLinkLogicalName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.useIPv6Address"); value.Exists() && !data.StateLinkUseIpv6.IsNull() {
		data.StateLinkUseIpv6 = types.BoolValue(value.Bool())
	} else {
		data.StateLinkUseIpv6 = types.BoolNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.activeIP"); value.Exists() && !data.StateLinkPrimaryIp.IsNull() {
		data.StateLinkPrimaryIp = types.StringValue(value.String())
	} else {
		data.StateLinkPrimaryIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.standbyIP"); value.Exists() && !data.StateLinkSecondaryIp.IsNull() {
		data.StateLinkSecondaryIp = types.StringValue(value.String())
	} else {
		data.StateLinkSecondaryIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.subnetMask"); value.Exists() && !data.StateLinkNetmask.IsNull() {
		data.StateLinkNetmask = types.StringValue(value.String())
	} else {
		data.StateLinkNetmask = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.isEncryptionEnabled"); value.Exists() && !data.EncryptionEnabled.IsNull() {
		data.EncryptionEnabled = types.BoolValue(value.Bool())
	} else {
		data.EncryptionEnabled = types.BoolNull()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.percentFailedInterfaceExceed"); value.Exists() && !data.FailedInterfacesPercent.IsNull() {
		data.FailedInterfacesPercent = types.Int64Value(value.Int())
	} else {
		data.FailedInterfacesPercent = types.Int64Null()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.noOfFailedInterfaceLimit"); value.Exists() && !data.FailedInterfacesLimit.IsNull() {
		data.FailedInterfacesLimit = types.Int64Value(value.Int())
	} else {
		data.FailedInterfacesLimit = types.Int64Null()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.peerPollTime"); value.Exists() && !data.PeerPollTime.IsNull() {
		data.PeerPollTime = types.Int64Value(value.Int())
	} else if data.PeerPollTime.ValueInt64() != 1 {
		data.PeerPollTime = types.Int64Null()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.peerPollTimeUnit"); value.Exists() && !data.PeerPollTimeUnit.IsNull() {
		data.PeerPollTimeUnit = types.StringValue(value.String())
	} else if data.PeerPollTimeUnit.ValueString() != "SEC" {
		data.PeerPollTimeUnit = types.StringNull()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.peerHoldTime"); value.Exists() && !data.PeerHoldTime.IsNull() {
		data.PeerHoldTime = types.Int64Value(value.Int())
	} else if data.PeerHoldTime.ValueInt64() != 15 {
		data.PeerHoldTime = types.Int64Null()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.peerHoldTimeUnit"); value.Exists() && !data.PeerHoldTimeUnit.IsNull() {
		data.PeerHoldTimeUnit = types.StringValue(value.String())
	} else if data.PeerHoldTimeUnit.ValueString() != "SEC" {
		data.PeerHoldTimeUnit = types.StringNull()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.interfacePollTime"); value.Exists() && !data.InterfacePollTime.IsNull() {
		data.InterfacePollTime = types.Int64Value(value.Int())
	} else if data.InterfacePollTime.ValueInt64() != 5 {
		data.InterfacePollTime = types.Int64Null()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.interfacePollTimeUnit"); value.Exists() && !data.InterfacePollTimeUnit.IsNull() {
		data.InterfacePollTimeUnit = types.StringValue(value.String())
	} else if data.InterfacePollTimeUnit.ValueString() != "SEC" {
		data.InterfacePollTimeUnit = types.StringNull()
	}
	if value := res.Get("ftdHAFailoverTriggerCriteria.interfaceHoldTime"); value.Exists() && !data.InterfaceHoldTime.IsNull() {
		data.InterfaceHoldTime = types.Int64Value(value.Int())
	} else if data.InterfaceHoldTime.ValueInt64() != 25 {
		data.InterfaceHoldTime = types.Int64Null()
	}
	if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceHAPair) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// toBodyPutDelete generates minimal required body to brek the HA.
func (data DeviceHAPair) toBodyPutDelete(ctx context.Context, state DeviceHAPair) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "action", "HABREAK")
	body, _ = sjson.Set(body, "forceBreak", true)
	return body
}

// toBodyUpdateTimers generates minimal required body to update timers.
// Those settings, even if set, are not updated by FMC when recieved as part of toBody function.
func (data DeviceHAPair) toBodyUpdateTimers(ctx context.Context, state DeviceHAPair) string {
	body := ""
	if !data.FailedInterfacesPercent.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.percentFailedInterfaceExceed", data.FailedInterfacesPercent.ValueInt64())
	}
	if !data.FailedInterfacesLimit.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.noOfFailedInterfaceLimit", data.FailedInterfacesLimit.ValueInt64())
	}
	if !data.PeerPollTime.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.peerPollTime", data.PeerPollTime.ValueInt64())
	}
	if !data.PeerPollTimeUnit.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.peerPollTimeUnit", data.PeerPollTimeUnit.ValueString())
	}
	if !data.PeerHoldTime.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.peerHoldTime", data.PeerHoldTime.ValueInt64())
	}
	if !data.PeerHoldTimeUnit.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.peerHoldTimeUnit", data.PeerHoldTimeUnit.ValueString())
	}
	if !data.InterfacePollTime.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.interfacePollTime", data.InterfacePollTime.ValueInt64())
	}
	if !data.InterfacePollTimeUnit.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.interfacePollTimeUnit", data.InterfacePollTimeUnit.ValueString())
	}
	if !data.InterfaceHoldTime.IsNull() {
		body, _ = sjson.Set(body, "ftdHAFailoverTriggerCriteria.interfaceHoldTime", data.InterfaceHoldTime.ValueInt64())
	}
	if body != "" && data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	return body
}
