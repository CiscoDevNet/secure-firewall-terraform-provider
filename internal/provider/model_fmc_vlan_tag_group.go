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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type VLANTagGroup struct {
	Id          types.String           `tfsdk:"id"`
	Domain      types.String           `tfsdk:"domain"`
	Name        types.String           `tfsdk:"name"`
	Description types.String           `tfsdk:"description"`
	Type        types.String           `tfsdk:"type"`
	Overridable types.Bool             `tfsdk:"overridable"`
	VlanTags    []VLANTagGroupVlanTags `tfsdk:"vlan_tags"`
	Literals    []VLANTagGroupLiterals `tfsdk:"literals"`
}

type VLANTagGroupVlanTags struct {
	Id types.String `tfsdk:"id"`
}

type VLANTagGroupLiterals struct {
	StartTag types.String `tfsdk:"start_tag"`
	EndTag   types.String `tfsdk:"end_tag"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VLANTagGroup) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/vlangrouptags"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VLANTagGroup) toBody(ctx context.Context, state VLANTagGroup) string {
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
	if !data.Overridable.IsNull() {
		body, _ = sjson.Set(body, "overridable", data.Overridable.ValueBool())
	}
	if len(data.VlanTags) > 0 {
		body, _ = sjson.Set(body, "objects", []interface{}{})
		for _, item := range data.VlanTags {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "objects.-1", itemBody)
		}
	}
	if len(data.Literals) > 0 {
		body, _ = sjson.Set(body, "literals", []interface{}{})
		for _, item := range data.Literals {
			itemBody := ""
			if !item.StartTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "startTag", item.StartTag.ValueString())
			}
			if !item.EndTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "endTag", item.EndTag.ValueString())
			}
			body, _ = sjson.SetRaw(body, "literals.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VLANTagGroup) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("overridable"); value.Exists() {
		data.Overridable = types.BoolValue(value.Bool())
	} else {
		data.Overridable = types.BoolNull()
	}
	if value := res.Get("objects"); value.Exists() {
		data.VlanTags = make([]VLANTagGroupVlanTags, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VLANTagGroupVlanTags{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).VlanTags = append((*parent).VlanTags, data)
			return true
		})
	}
	if value := res.Get("literals"); value.Exists() {
		data.Literals = make([]VLANTagGroupLiterals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VLANTagGroupLiterals{}
			if value := res.Get("startTag"); value.Exists() {
				data.StartTag = types.StringValue(value.String())
			} else {
				data.StartTag = types.StringNull()
			}
			if value := res.Get("endTag"); value.Exists() {
				data.EndTag = types.StringValue(value.String())
			} else {
				data.EndTag = types.StringNull()
			}
			(*parent).Literals = append((*parent).Literals, data)
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
func (data *VLANTagGroup) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("overridable"); value.Exists() && !data.Overridable.IsNull() {
		data.Overridable = types.BoolValue(value.Bool())
	} else {
		data.Overridable = types.BoolNull()
	}
	for i := 0; i < len(data.VlanTags); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.VlanTags[i].Id.ValueString()}

		parent := &data
		data := (*parent).VlanTags[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing VlanTags[%d] = %+v",
				i,
				(*parent).VlanTags[i],
			))
			(*parent).VlanTags = slices.Delete((*parent).VlanTags, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).VlanTags[i] = data
	}
	for i := 0; i < len(data.Literals); i++ {
		keys := [...]string{"startTag", "endTag"}
		keyValues := [...]string{data.Literals[i].StartTag.ValueString(), data.Literals[i].EndTag.ValueString()}

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
		if value := res.Get("startTag"); value.Exists() && !data.StartTag.IsNull() {
			data.StartTag = types.StringValue(value.String())
		} else {
			data.StartTag = types.StringNull()
		}
		if value := res.Get("endTag"); value.Exists() && !data.EndTag.IsNull() {
			data.EndTag = types.StringValue(value.String())
		} else {
			data.EndTag = types.StringNull()
		}
		(*parent).Literals[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VLANTagGroup) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns
