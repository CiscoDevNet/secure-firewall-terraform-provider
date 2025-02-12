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

type DeviceHAPairMonitoring struct {
	Id                 types.String                          `tfsdk:"id"`
	Domain             types.String                          `tfsdk:"domain"`
	DeviceId           types.String                          `tfsdk:"device_id"`
	Type               types.String                          `tfsdk:"type"`
	LogicalName        types.String                          `tfsdk:"logical_name"`
	MonitorInterface   types.Bool                            `tfsdk:"monitor_interface"`
	Ipv4ActiveAddress  types.String                          `tfsdk:"ipv4_active_address"`
	Ipv4StandbyAddress types.String                          `tfsdk:"ipv4_standby_address"`
	Ipv4Netmask        types.String                          `tfsdk:"ipv4_netmask"`
	Ipv6Addresses      []DeviceHAPairMonitoringIpv6Addresses `tfsdk:"ipv6_addresses"`
}

type DeviceHAPairMonitoringIpv6Addresses struct {
	ActiveAddress  types.String `tfsdk:"active_address"`
	StandbyAddress types.String `tfsdk:"standby_address"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceHAPairMonitoring) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicehapairs/ftddevicehapairs/%v/monitoredinterfaces", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceHAPairMonitoring) toBody(ctx context.Context, state DeviceHAPairMonitoring) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "name", data.LogicalName.ValueString())
	}
	if !data.MonitorInterface.IsNull() {
		body, _ = sjson.Set(body, "monitorForFailures", data.MonitorInterface.ValueBool())
	}
	if !data.Ipv4StandbyAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv4Configuration.standbyIPv4Address", data.Ipv4StandbyAddress.ValueString())
	}
	if len(data.Ipv6Addresses) > 0 {
		body, _ = sjson.Set(body, "ipv6Configuration.ipv6ActiveStandbyPair", []interface{}{})
		for _, item := range data.Ipv6Addresses {
			itemBody := ""
			if !item.ActiveAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "activeIPv6", item.ActiveAddress.ValueString())
			}
			if !item.StandbyAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "standbyIPv6", item.StandbyAddress.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ipv6Configuration.ipv6ActiveStandbyPair.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceHAPairMonitoring) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("monitorForFailures"); value.Exists() {
		data.MonitorInterface = types.BoolValue(value.Bool())
	} else {
		data.MonitorInterface = types.BoolNull()
	}
	if value := res.Get("ipv4Configuration.activeIPv4Address"); value.Exists() {
		data.Ipv4ActiveAddress = types.StringValue(value.String())
	} else {
		data.Ipv4ActiveAddress = types.StringNull()
	}
	if value := res.Get("ipv4Configuration.standbyIPv4Address"); value.Exists() {
		data.Ipv4StandbyAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StandbyAddress = types.StringNull()
	}
	if value := res.Get("ipv4Configuration.activeIPv4Mask"); value.Exists() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
	if value := res.Get("ipv6Configuration.ipv6ActiveStandbyPair"); value.Exists() {
		data.Ipv6Addresses = make([]DeviceHAPairMonitoringIpv6Addresses, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceHAPairMonitoringIpv6Addresses{}
			if value := res.Get("activeIPv6"); value.Exists() {
				data.ActiveAddress = types.StringValue(value.String())
			} else {
				data.ActiveAddress = types.StringNull()
			}
			if value := res.Get("standbyIPv6"); value.Exists() {
				data.StandbyAddress = types.StringValue(value.String())
			} else {
				data.StandbyAddress = types.StringNull()
			}
			(*parent).Ipv6Addresses = append((*parent).Ipv6Addresses, data)
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
func (data *DeviceHAPairMonitoring) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.LogicalName.IsNull() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("monitorForFailures"); value.Exists() && !data.MonitorInterface.IsNull() {
		data.MonitorInterface = types.BoolValue(value.Bool())
	} else {
		data.MonitorInterface = types.BoolNull()
	}
	if value := res.Get("ipv4Configuration.activeIPv4Address"); value.Exists() && !data.Ipv4ActiveAddress.IsNull() {
		data.Ipv4ActiveAddress = types.StringValue(value.String())
	} else {
		data.Ipv4ActiveAddress = types.StringNull()
	}
	if value := res.Get("ipv4Configuration.standbyIPv4Address"); value.Exists() && !data.Ipv4StandbyAddress.IsNull() {
		data.Ipv4StandbyAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StandbyAddress = types.StringNull()
	}
	if value := res.Get("ipv4Configuration.activeIPv4Mask"); value.Exists() && !data.Ipv4Netmask.IsNull() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
	for i := 0; i < len(data.Ipv6Addresses); i++ {
		keys := [...]string{"activeIPv6", "standbyIPv6"}
		keyValues := [...]string{data.Ipv6Addresses[i].ActiveAddress.ValueString(), data.Ipv6Addresses[i].StandbyAddress.ValueString()}

		parent := &data
		data := (*parent).Ipv6Addresses[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ipv6Configuration.ipv6ActiveStandbyPair").ForEach(
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
		if value := res.Get("activeIPv6"); value.Exists() && !data.ActiveAddress.IsNull() {
			data.ActiveAddress = types.StringValue(value.String())
		} else {
			data.ActiveAddress = types.StringNull()
		}
		if value := res.Get("standbyIPv6"); value.Exists() && !data.StandbyAddress.IsNull() {
			data.StandbyAddress = types.StringValue(value.String())
		} else {
			data.StandbyAddress = types.StringNull()
		}
		(*parent).Ipv6Addresses[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceHAPairMonitoring) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.Ipv4ActiveAddress.IsUnknown() {
		if value := res.Get("ipv4Configuration.activeIPv4Address"); value.Exists() {
			data.Ipv4ActiveAddress = types.StringValue(value.String())
		} else {
			data.Ipv4ActiveAddress = types.StringNull()
		}
	}
	if data.Ipv4Netmask.IsUnknown() {
		if value := res.Get("ipv4Configuration.activeIPv4Mask"); value.Exists() {
			data.Ipv4Netmask = types.StringValue(value.String())
		} else {
			data.Ipv4Netmask = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// toBodyPutDelete generates minimal required body to reset the resource to its default state.
func (data DeviceHAPairMonitoring) toBodyPutDelete(ctx context.Context, state DeviceHAPairMonitoring) string {
	body := ""
	body, _ = sjson.Set(body, "monitorForFailures", false)
	if data.Ipv4ActiveAddress.ValueString() != "" {
		body, _ = sjson.Set(body, "ipv4Configuration.activeIPv4Address", data.Ipv4ActiveAddress.ValueString())
	}
	// There is no way of removing standby IPv6 via API now
	//if len(data.Ipv6Addresses) > 0 {
	//	body, _ = sjson.Set(body, "ipv6Configuration.ipv6ActiveStandbyPair", []interface{}{})
	//	for _, item := range data.Ipv6Addresses {
	//		itemBody := ""
	//		if !item.ActiveAddress.IsNull() {
	//			itemBody, _ = sjson.Set(itemBody, "activeIPv6", item.ActiveAddress.ValueString())
	//			itemBody, _ = sjson.Set(itemBody, "standbyIPv6", "")
	//		}
	//		body, _ = sjson.SetRaw(body, "ipv6Configuration.ipv6ActiveStandbyPair.-1", itemBody)
	//	}
	//}
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "name", data.LogicalName.ValueString())
	}
	return body
}
