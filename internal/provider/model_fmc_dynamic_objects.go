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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DynamicObjects struct {
	Id     types.String                   `tfsdk:"id"`
	Domain types.String                   `tfsdk:"domain"`
	Items  map[string]DynamicObjectsItems `tfsdk:"items"`
}

type DynamicObjectsItems struct {
	Id          types.String `tfsdk:"id"`
	Type        types.String `tfsdk:"type"`
	Description types.String `tfsdk:"description"`
	ObjectType  types.String `tfsdk:"object_type"`
	Mappings    types.Set    `tfsdk:"mappings"`
}

// End of section. //template:end types

type DynamicObjectsMappings struct {
	Items map[string]DynamicObjectsMappingsItems
}

type DynamicObjectsMappingsItems struct {
	Id     types.String
	Add    types.Set
	Remove types.Set
}

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DynamicObjects) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/dynamicobjects"
}

// End of section. //template:end getPath

func (data DynamicObjectsMappings) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/dynamicobjectmappings"
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DynamicObjects) toBody(ctx context.Context, state DynamicObjects) string {
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
			if !item.ObjectType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "objectType", item.ObjectType.ValueString())
			}
			if !item.Mappings.IsNull() {
				var values []string
				item.Mappings.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "mappings", values)
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DynamicObjects) fromBody(ctx context.Context, res gjson.Result) {
	for k := range data.Items {
		parent := &data
		data := (*parent).Items[k]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("name").String() == k {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: name=%v", k))
			delete((*parent).Items, k)
			continue
		}
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
		if value := res.Get("description"); value.Exists() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("objectType"); value.Exists() {
			data.ObjectType = types.StringValue(value.String())
		} else {
			data.ObjectType = types.StringNull()
		}
		if value := res.Get("mappings"); value.Exists() {
			data.Mappings = helpers.GetStringSet(value.Array())
		} else {
			data.Mappings = types.SetNull(types.StringType)
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
func (data *DynamicObjects) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("objectType"); value.Exists() && !data.ObjectType.IsNull() {
			data.ObjectType = types.StringValue(value.String())
		} else {
			data.ObjectType = types.StringNull()
		}
		if value := res.Get("mappings"); value.Exists() && !data.Mappings.IsNull() {
			data.Mappings = helpers.GetStringSet(value.Array())
		} else {
			data.Mappings = types.SetNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DynamicObjects) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
		if v := data.Items[i]; v.Type.IsUnknown() {
			if value := r.Get("type"); value.Exists() {
				v.Type = types.StringValue(value.String())
			} else {
				v.Type = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

func (data *DynamicObjects) Clone() DynamicObjects {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data DynamicObjects) toBodyNonBulk(ctx context.Context, state DynamicObjects) string {
	// This is one-by-one update, so only one element to update is expected
	if len(data.Items) > 1 {
		tflog.Error(ctx, "Found more than one element to chage. Only one will be changed.")
	}

	// Utilize existing toBody function
	body := data.toBody(ctx, state)

	// Get first element only
	return gjson.Get(body, "0").String()
}

// End of section. //template:end toBodyNonBulk

// Cast to DynamicObjectsMappings structure into request body
func (data DynamicObjectsMappings) toBody(ctx context.Context, state DynamicObjectsMappings) string {
	body := "{}"
	body, _ = sjson.SetRaw(body, "add", "[]")
	body, _ = sjson.SetRaw(body, "remove", "[]")
	if len(data.Items) > 0 {
		for _, item := range data.Items {
			if !item.Add.IsNull() {
				var values []string
				addBody, _ := sjson.Set("{}", "dynamicObject.id", item.Id.ValueString())
				item.Add.ElementsAs(ctx, &values, false)
				addBody, _ = sjson.Set(addBody, "mappings", values)
				body, _ = sjson.SetRaw(body, "add.-1", addBody)
			}
			if !item.Remove.IsNull() {
				var values []string
				removeBody, _ := sjson.Set("{}", "dynamicObject.id", item.Id.ValueString())
				item.Remove.ElementsAs(ctx, &values, false)
				removeBody, _ = sjson.Set(removeBody, "mappings", values)
				body, _ = sjson.SetRaw(body, "remove.-1", removeBody)
			}
		}
	}
	return body
}

// Read mappings from response body
func (data DynamicObjects) fromBodyMapping(ctx context.Context, res gjson.Result) types.Set {
	mappings := res.Get("items.#.mapping").Array()
	tflog.Debug(ctx, fmt.Sprintf("Mappings: %v", mappings))
	ret := helpers.GetStringSet(mappings)
	tflog.Debug(ctx, fmt.Sprintf("ret: %v", ret))
	return ret
}
