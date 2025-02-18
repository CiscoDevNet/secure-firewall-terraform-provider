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

type ApplicationFilter struct {
	Id           types.String                    `tfsdk:"id"`
	Domain       types.String                    `tfsdk:"domain"`
	Name         types.String                    `tfsdk:"name"`
	Type         types.String                    `tfsdk:"type"`
	Applications []ApplicationFilterApplications `tfsdk:"applications"`
	Filters      []ApplicationFilterFilters      `tfsdk:"filters"`
}

type ApplicationFilterApplications struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ApplicationFilterFilters struct {
	Types              []ApplicationFilterFiltersTypes              `tfsdk:"types"`
	Risks              []ApplicationFilterFiltersRisks              `tfsdk:"risks"`
	BusinessRelevances []ApplicationFilterFiltersBusinessRelevances `tfsdk:"business_relevances"`
	Categories         []ApplicationFilterFiltersCategories         `tfsdk:"categories"`
	Tags               []ApplicationFilterFiltersTags               `tfsdk:"tags"`
}

type ApplicationFilterFiltersTypes struct {
	Id types.String `tfsdk:"id"`
}
type ApplicationFilterFiltersRisks struct {
	Id types.String `tfsdk:"id"`
}
type ApplicationFilterFiltersBusinessRelevances struct {
	Id types.String `tfsdk:"id"`
}
type ApplicationFilterFiltersCategories struct {
	Id types.String `tfsdk:"id"`
}
type ApplicationFilterFiltersTags struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplicationFilter) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/applicationfilters"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplicationFilter) toBody(ctx context.Context, state ApplicationFilter) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if len(data.Applications) > 0 {
		body, _ = sjson.Set(body, "applications", []interface{}{})
		for _, item := range data.Applications {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "applications.-1", itemBody)
		}
	}
	if len(data.Filters) > 0 {
		body, _ = sjson.Set(body, "appConditions", []interface{}{})
		for _, item := range data.Filters {
			itemBody := ""
			if len(item.Types) > 0 {
				itemBody, _ = sjson.Set(itemBody, "applicationTypes", []interface{}{})
				for _, childItem := range item.Types {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "applicationTypes.-1", itemChildBody)
				}
			}
			if len(item.Risks) > 0 {
				itemBody, _ = sjson.Set(itemBody, "risks", []interface{}{})
				for _, childItem := range item.Risks {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "risks.-1", itemChildBody)
				}
			}
			if len(item.BusinessRelevances) > 0 {
				itemBody, _ = sjson.Set(itemBody, "productivities", []interface{}{})
				for _, childItem := range item.BusinessRelevances {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "productivities.-1", itemChildBody)
				}
			}
			if len(item.Categories) > 0 {
				itemBody, _ = sjson.Set(itemBody, "categories", []interface{}{})
				for _, childItem := range item.Categories {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "categories.-1", itemChildBody)
				}
			}
			if len(item.Tags) > 0 {
				itemBody, _ = sjson.Set(itemBody, "tags", []interface{}{})
				for _, childItem := range item.Tags {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "tags.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "appConditions.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplicationFilter) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("applications"); value.Exists() {
		data.Applications = make([]ApplicationFilterApplications, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplicationFilterApplications{}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).Applications = append((*parent).Applications, data)
			return true
		})
	}
	if value := res.Get("appConditions"); value.Exists() {
		data.Filters = make([]ApplicationFilterFilters, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplicationFilterFilters{}
			if value := res.Get("applicationTypes"); value.Exists() {
				data.Types = make([]ApplicationFilterFiltersTypes, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplicationFilterFiltersTypes{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Types = append((*parent).Types, data)
					return true
				})
			}
			if value := res.Get("risks"); value.Exists() {
				data.Risks = make([]ApplicationFilterFiltersRisks, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplicationFilterFiltersRisks{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Risks = append((*parent).Risks, data)
					return true
				})
			}
			if value := res.Get("productivities"); value.Exists() {
				data.BusinessRelevances = make([]ApplicationFilterFiltersBusinessRelevances, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplicationFilterFiltersBusinessRelevances{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).BusinessRelevances = append((*parent).BusinessRelevances, data)
					return true
				})
			}
			if value := res.Get("categories"); value.Exists() {
				data.Categories = make([]ApplicationFilterFiltersCategories, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplicationFilterFiltersCategories{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Categories = append((*parent).Categories, data)
					return true
				})
			}
			if value := res.Get("tags"); value.Exists() {
				data.Tags = make([]ApplicationFilterFiltersTags, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplicationFilterFiltersTags{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Tags = append((*parent).Tags, data)
					return true
				})
			}
			(*parent).Filters = append((*parent).Filters, data)
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
func (data *ApplicationFilter) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	for i := 0; i < len(data.Applications); i++ {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Applications[i].Name.ValueString()}

		parent := &data
		data := (*parent).Applications[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("applications").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Applications[%d] = %+v",
				i,
				(*parent).Applications[i],
			))
			(*parent).Applications = slices.Delete((*parent).Applications, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Applications[i] = data
	}
	{
		l := len(res.Get("appConditions").Array())
		tflog.Debug(ctx, fmt.Sprintf("appConditions array resizing from %d to %d", len(data.Filters), l))
		for i := len(data.Filters); i < l; i++ {
			data.Filters = append(data.Filters, ApplicationFilterFilters{})
		}
		if len(data.Filters) > l {
			data.Filters = data.Filters[:l]
		}
	}
	for i := range data.Filters {
		parent := &data
		data := (*parent).Filters[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("appConditions.%d", i))
		for i := 0; i < len(data.Types); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Types[i].Id.ValueString()}

			parent := &data
			data := (*parent).Types[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("applicationTypes").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Types[%d] = %+v",
					i,
					(*parent).Types[i],
				))
				(*parent).Types = slices.Delete((*parent).Types, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Types[i] = data
		}
		for i := 0; i < len(data.Risks); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Risks[i].Id.ValueString()}

			parent := &data
			data := (*parent).Risks[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("risks").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Risks[%d] = %+v",
					i,
					(*parent).Risks[i],
				))
				(*parent).Risks = slices.Delete((*parent).Risks, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Risks[i] = data
		}
		for i := 0; i < len(data.BusinessRelevances); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.BusinessRelevances[i].Id.ValueString()}

			parent := &data
			data := (*parent).BusinessRelevances[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("productivities").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing BusinessRelevances[%d] = %+v",
					i,
					(*parent).BusinessRelevances[i],
				))
				(*parent).BusinessRelevances = slices.Delete((*parent).BusinessRelevances, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).BusinessRelevances[i] = data
		}
		for i := 0; i < len(data.Categories); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Categories[i].Id.ValueString()}

			parent := &data
			data := (*parent).Categories[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("categories").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Categories[%d] = %+v",
					i,
					(*parent).Categories[i],
				))
				(*parent).Categories = slices.Delete((*parent).Categories, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Categories[i] = data
		}
		for i := 0; i < len(data.Tags); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Tags[i].Id.ValueString()}

			parent := &data
			data := (*parent).Tags[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("tags").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Tags[%d] = %+v",
					i,
					(*parent).Tags[i],
				))
				(*parent).Tags = slices.Delete((*parent).Tags, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Tags[i] = data
		}
		(*parent).Filters[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplicationFilter) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
