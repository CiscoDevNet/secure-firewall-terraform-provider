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
	"maps"
	"slices"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkGroups struct {
	Id     types.String                  `tfsdk:"id"`
	Domain types.String                  `tfsdk:"domain"`
	Items  map[string]NetworkGroupsItems `tfsdk:"items"`
}

type NetworkGroupsItems struct {
	Id            types.String                 `tfsdk:"id"`
	Description   types.String                 `tfsdk:"description"`
	Overridable   types.Bool                   `tfsdk:"overridable"`
	NetworkGroups types.Set                    `tfsdk:"network_groups"`
	Objects       []NetworkGroupsItemsObjects  `tfsdk:"objects"`
	Literals      []NetworkGroupsItemsLiterals `tfsdk:"literals"`
}

type NetworkGroupsItemsObjects struct {
	Id types.String `tfsdk:"id"`
}
type NetworkGroupsItemsLiterals struct {
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkGroups) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/networkgroups"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkGroups) toBody(ctx context.Context, state NetworkGroups) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for key, item := range data.Items {
			itemBody, _ := sjson.Set("{}", "name", key)
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.Overridable.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overridable", item.Overridable.ValueBool())
			}
			if !item.NetworkGroups.IsNull() {
				var values []string
				item.NetworkGroups.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "network_groups", values)
			}
			if len(item.Objects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "objects", []interface{}{})
				for _, childItem := range item.Objects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "objects.-1", itemChildBody)
				}
			}
			if len(item.Literals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "literals", []interface{}{})
				for _, childItem := range item.Literals {
					itemChildBody := ""
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "literals.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkGroups) fromBody(ctx context.Context, res gjson.Result) {
	for k := range data.Items {
		parent := &data
		data := (*parent).Items[k]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: uuid=%s, key=%v", data.Id, k))
			delete((*parent).Items, k)
			continue
		}
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("description"); value.Exists() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("overridable"); value.Exists() {
			data.Overridable = types.BoolValue(value.Bool())
		} else {
			data.Overridable = types.BoolNull()
		}
		if value := res.Get("network_groups"); value.Exists() {
			data.NetworkGroups = helpers.GetStringSet(value.Array())
		} else {
			data.NetworkGroups = types.SetNull(types.StringType)
		}
		if value := res.Get("objects"); value.Exists() {
			data.Objects = make([]NetworkGroupsItemsObjects, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkGroupsItemsObjects{}
				if value := res.Get("id"); value.Exists() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).Objects = append((*parent).Objects, data)
				return true
			})
		}
		if value := res.Get("literals"); value.Exists() {
			data.Literals = make([]NetworkGroupsItemsLiterals, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkGroupsItemsLiterals{}
				if value := res.Get("value"); value.Exists() {
					data.Value = types.StringValue(value.String())
				} else {
					data.Value = types.StringNull()
				}
				(*parent).Literals = append((*parent).Literals, data)
				return true
			})
		}
		(*parent).Items[k] = data
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkGroups) fromBodyPartial(ctx context.Context, res gjson.Result) {
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("overridable"); value.Exists() && !data.Overridable.IsNull() {
			data.Overridable = types.BoolValue(value.Bool())
		} else {
			data.Overridable = types.BoolNull()
		}
		if value := res.Get("network_groups"); value.Exists() && !data.NetworkGroups.IsNull() {
			data.NetworkGroups = helpers.GetStringSet(value.Array())
		} else {
			data.NetworkGroups = types.SetNull(types.StringType)
		}
		for i := 0; i < len(data.Objects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Objects[i].Id.ValueString()}

			parent := &data
			data := (*parent).Objects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Objects[%d] = %+v",
					i,
					(*parent).Objects[i],
				))
				(*parent).Objects = slices.Delete((*parent).Objects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Objects[i] = data
		}
		for i := 0; i < len(data.Literals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.Literals[i].Value.ValueString()}

			parent := &data
			data := (*parent).Literals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("literals").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Literals[%d] = %+v",
					i,
					(*parent).Literals[i],
				))
				(*parent).Literals = slices.Delete((*parent).Literals, i, i+1)
				i--

				continue
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).Literals[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkGroups) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	for i, val := range data.Items {
		var r gjson.Result
		res.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if val.Id.IsUnknown() {
					if v.Get("name").String() == i {
						r = v
						return false // break ForEach
					}
				} else {
					if v.Get("id").String() == val.Id.ValueString() && val.Id.ValueString() != "" {
						r = v
						return false // break ForEach
					}
				}

				return true
			},
		)
		if v := data.Items[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data *NetworkGroups) Clone() NetworkGroups {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}
