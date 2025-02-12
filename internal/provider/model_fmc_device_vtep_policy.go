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

type DeviceVTEPPolicy struct {
	Id         types.String            `tfsdk:"id"`
	Domain     types.String            `tfsdk:"domain"`
	DeviceId   types.String            `tfsdk:"device_id"`
	NveEnabled types.Bool              `tfsdk:"nve_enabled"`
	Vteps      []DeviceVTEPPolicyVteps `tfsdk:"vteps"`
}

type DeviceVTEPPolicyVteps struct {
	SourceInterfaceId      types.String `tfsdk:"source_interface_id"`
	NveNumber              types.Int64  `tfsdk:"nve_number"`
	EncapsulationPort      types.Int64  `tfsdk:"encapsulation_port"`
	EncapsulationType      types.String `tfsdk:"encapsulation_type"`
	NeighborDiscovery      types.String `tfsdk:"neighbor_discovery"`
	NeighborAddressLiteral types.String `tfsdk:"neighbor_address_literal"`
	NeighborAddressId      types.String `tfsdk:"neighbor_address_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceVTEPPolicy) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/vteppolicies", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceVTEPPolicy) toBody(ctx context.Context, state DeviceVTEPPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.NveEnabled.IsNull() {
		body, _ = sjson.Set(body, "nveEnable", data.NveEnabled.ValueBool())
	}
	if len(data.Vteps) > 0 {
		body, _ = sjson.Set(body, "vtepEntries", []interface{}{})
		for _, item := range data.Vteps {
			itemBody := ""
			if !item.SourceInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceInterface.id", item.SourceInterfaceId.ValueString())
			}
			if !item.NveNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nveVtepId", item.NveNumber.ValueInt64())
			}
			if !item.EncapsulationPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nveDestinationPort", item.EncapsulationPort.ValueInt64())
			}
			if !item.EncapsulationType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nveEncapsulationType", item.EncapsulationType.ValueString())
			}
			if !item.NeighborDiscovery.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nveNeighborDiscoveryType", item.NeighborDiscovery.ValueString())
			}
			if !item.NeighborAddressLiteral.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nveNeighborAddress.literal.value", item.NeighborAddressLiteral.ValueString())
			}
			if !item.NeighborAddressId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nveNeighborAddress.object.id", item.NeighborAddressId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "vtepEntries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceVTEPPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("nveEnable"); value.Exists() {
		data.NveEnabled = types.BoolValue(value.Bool())
	} else {
		data.NveEnabled = types.BoolValue(true)
	}
	if value := res.Get("vtepEntries"); value.Exists() {
		data.Vteps = make([]DeviceVTEPPolicyVteps, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceVTEPPolicyVteps{}
			if value := res.Get("sourceInterface.id"); value.Exists() {
				data.SourceInterfaceId = types.StringValue(value.String())
			} else {
				data.SourceInterfaceId = types.StringNull()
			}
			if value := res.Get("nveVtepId"); value.Exists() {
				data.NveNumber = types.Int64Value(value.Int())
			} else {
				data.NveNumber = types.Int64Null()
			}
			if value := res.Get("nveDestinationPort"); value.Exists() {
				data.EncapsulationPort = types.Int64Value(value.Int())
			} else {
				data.EncapsulationPort = types.Int64Value(4789)
			}
			if value := res.Get("nveEncapsulationType"); value.Exists() {
				data.EncapsulationType = types.StringValue(value.String())
			} else {
				data.EncapsulationType = types.StringValue("VXLAN")
			}
			if value := res.Get("nveNeighborDiscoveryType"); value.Exists() {
				data.NeighborDiscovery = types.StringValue(value.String())
			} else {
				data.NeighborDiscovery = types.StringNull()
			}
			if value := res.Get("nveNeighborAddress.literal.value"); value.Exists() {
				data.NeighborAddressLiteral = types.StringValue(value.String())
			} else {
				data.NeighborAddressLiteral = types.StringNull()
			}
			if value := res.Get("nveNeighborAddress.object.id"); value.Exists() {
				data.NeighborAddressId = types.StringValue(value.String())
			} else {
				data.NeighborAddressId = types.StringNull()
			}
			(*parent).Vteps = append((*parent).Vteps, data)
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
func (data *DeviceVTEPPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("nveEnable"); value.Exists() && !data.NveEnabled.IsNull() {
		data.NveEnabled = types.BoolValue(value.Bool())
	} else if data.NveEnabled.ValueBool() != true {
		data.NveEnabled = types.BoolNull()
	}
	for i := 0; i < len(data.Vteps); i++ {
		keys := [...]string{"sourceInterface.id"}
		keyValues := [...]string{data.Vteps[i].SourceInterfaceId.ValueString()}

		parent := &data
		data := (*parent).Vteps[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("vtepEntries").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Vteps[%d] = %+v",
				i,
				(*parent).Vteps[i],
			))
			(*parent).Vteps = slices.Delete((*parent).Vteps, i, i+1)
			i--

			continue
		}
		if value := res.Get("sourceInterface.id"); value.Exists() && !data.SourceInterfaceId.IsNull() {
			data.SourceInterfaceId = types.StringValue(value.String())
		} else {
			data.SourceInterfaceId = types.StringNull()
		}
		if value := res.Get("nveVtepId"); value.Exists() && !data.NveNumber.IsNull() {
			data.NveNumber = types.Int64Value(value.Int())
		} else {
			data.NveNumber = types.Int64Null()
		}
		if value := res.Get("nveDestinationPort"); value.Exists() && !data.EncapsulationPort.IsNull() {
			data.EncapsulationPort = types.Int64Value(value.Int())
		} else if data.EncapsulationPort.ValueInt64() != 4789 {
			data.EncapsulationPort = types.Int64Null()
		}
		if value := res.Get("nveEncapsulationType"); value.Exists() && !data.EncapsulationType.IsNull() {
			data.EncapsulationType = types.StringValue(value.String())
		} else if data.EncapsulationType.ValueString() != "VXLAN" {
			data.EncapsulationType = types.StringNull()
		}
		if value := res.Get("nveNeighborDiscoveryType"); value.Exists() && !data.NeighborDiscovery.IsNull() {
			data.NeighborDiscovery = types.StringValue(value.String())
		} else {
			data.NeighborDiscovery = types.StringNull()
		}
		if value := res.Get("nveNeighborAddress.literal.value"); value.Exists() && !data.NeighborAddressLiteral.IsNull() {
			data.NeighborAddressLiteral = types.StringValue(value.String())
		} else {
			data.NeighborAddressLiteral = types.StringNull()
		}
		if value := res.Get("nveNeighborAddress.object.id"); value.Exists() && !data.NeighborAddressId.IsNull() {
			data.NeighborAddressId = types.StringValue(value.String())
		} else {
			data.NeighborAddressId = types.StringNull()
		}
		(*parent).Vteps[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceVTEPPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
