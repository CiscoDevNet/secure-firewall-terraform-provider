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

type DeviceVRF struct {
	Id          types.String          `tfsdk:"id"`
	Domain      types.String          `tfsdk:"domain"`
	DeviceId    types.String          `tfsdk:"device_id"`
	Name        types.String          `tfsdk:"name"`
	Type        types.String          `tfsdk:"type"`
	Description types.String          `tfsdk:"description"`
	Interfaces  []DeviceVRFInterfaces `tfsdk:"interfaces"`
}

type DeviceVRFInterfaces struct {
	InterfaceId          types.String `tfsdk:"interface_id"`
	InterfaceName        types.String `tfsdk:"interface_name"`
	InterfaceLogicalName types.String `tfsdk:"interface_logical_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceVRF) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/virtualrouters", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceVRF) toBody(ctx context.Context, state DeviceVRF) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, "interfaces", []interface{}{})
		for _, item := range data.Interfaces {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.InterfaceId.ValueString())
			}
			if !item.InterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.InterfaceName.ValueString())
			}
			if !item.InterfaceLogicalName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ifname", item.InterfaceLogicalName.ValueString())
			}
			body, _ = sjson.SetRaw(body, "interfaces.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceVRF) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("interfaces"); value.Exists() {
		data.Interfaces = make([]DeviceVRFInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceVRFInterfaces{}
			if value := res.Get("id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() {
				data.InterfaceName = types.StringValue(value.String())
			} else {
				data.InterfaceName = types.StringNull()
			}
			if value := res.Get("ifname"); value.Exists() {
				data.InterfaceLogicalName = types.StringValue(value.String())
			} else {
				data.InterfaceLogicalName = types.StringNull()
			}
			(*parent).Interfaces = append((*parent).Interfaces, data)
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
func (data *DeviceVRF) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := 0; i < len(data.Interfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Interfaces[i].InterfaceId.ValueString()}

		parent := &data
		data := (*parent).Interfaces[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("interfaces").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Interfaces[%d] = %+v",
				i,
				(*parent).Interfaces[i],
			))
			(*parent).Interfaces = slices.Delete((*parent).Interfaces, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.InterfaceName.IsNull() {
			data.InterfaceName = types.StringValue(value.String())
		} else {
			data.InterfaceName = types.StringNull()
		}
		if value := res.Get("ifname"); value.Exists() && !data.InterfaceLogicalName.IsNull() {
			data.InterfaceLogicalName = types.StringValue(value.String())
		} else {
			data.InterfaceLogicalName = types.StringNull()
		}
		(*parent).Interfaces[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceVRF) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
