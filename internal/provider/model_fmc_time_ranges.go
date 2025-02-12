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
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type TimeRanges struct {
	Id     types.String               `tfsdk:"id"`
	Domain types.String               `tfsdk:"domain"`
	Items  map[string]TimeRangesItems `tfsdk:"items"`
}

type TimeRangesItems struct {
	Id             types.String                    `tfsdk:"id"`
	Type           types.String                    `tfsdk:"type"`
	Description    types.String                    `tfsdk:"description"`
	StartTime      types.String                    `tfsdk:"start_time"`
	EndTime        types.String                    `tfsdk:"end_time"`
	RecurrenceList []TimeRangesItemsRecurrenceList `tfsdk:"recurrence_list"`
}

type TimeRangesItemsRecurrenceList struct {
	RecurrenceType types.String `tfsdk:"recurrence_type"`
	RangeStartTime types.String `tfsdk:"range_start_time"`
	RangeEndTime   types.String `tfsdk:"range_end_time"`
	RangeStartDay  types.String `tfsdk:"range_start_day"`
	RangeEndDay    types.String `tfsdk:"range_end_day"`
	DailyStartTime types.String `tfsdk:"daily_start_time"`
	DailyEndTime   types.String `tfsdk:"daily_end_time"`
	DailyDays      types.Set    `tfsdk:"daily_days"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkDeleteTimeRanges = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data TimeRanges) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/timeranges"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data TimeRanges) toBody(ctx context.Context, state TimeRanges) string {
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
			if !item.StartTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "effectiveStartDateTime", item.StartTime.ValueString())
			}
			if !item.EndTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "effectiveEndDateTime", item.EndTime.ValueString())
			}
			if len(item.RecurrenceList) > 0 {
				itemBody, _ = sjson.Set(itemBody, "recurrenceList", []interface{}{})
				for _, childItem := range item.RecurrenceList {
					itemChildBody := ""
					if !childItem.RecurrenceType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "recurrenceType", childItem.RecurrenceType.ValueString())
					}
					if !childItem.RangeStartTime.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "rangeStartTime", childItem.RangeStartTime.ValueString())
					}
					if !childItem.RangeEndTime.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "rangeEndTime", childItem.RangeEndTime.ValueString())
					}
					if !childItem.RangeStartDay.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "rangeStartDay", childItem.RangeStartDay.ValueString())
					}
					if !childItem.RangeEndDay.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "rangeEndDay", childItem.RangeEndDay.ValueString())
					}
					if !childItem.DailyStartTime.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "dailyStartTime", childItem.DailyStartTime.ValueString())
					}
					if !childItem.DailyEndTime.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "dailyEndTime", childItem.DailyEndTime.ValueString())
					}
					if !childItem.DailyDays.IsNull() {
						var values []string
						childItem.DailyDays.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "days", values)
					}
					itemBody, _ = sjson.SetRaw(itemBody, "recurrenceList.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *TimeRanges) fromBody(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("effectiveStartDateTime"); value.Exists() {
			data.StartTime = types.StringValue(value.String())
		} else {
			data.StartTime = types.StringNull()
		}
		if value := res.Get("effectiveEndDateTime"); value.Exists() {
			data.EndTime = types.StringValue(value.String())
		} else {
			data.EndTime = types.StringNull()
		}
		if value := res.Get("recurrenceList"); value.Exists() {
			data.RecurrenceList = make([]TimeRangesItemsRecurrenceList, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := TimeRangesItemsRecurrenceList{}
				if value := res.Get("recurrenceType"); value.Exists() {
					data.RecurrenceType = types.StringValue(value.String())
				} else {
					data.RecurrenceType = types.StringNull()
				}
				if value := res.Get("rangeStartTime"); value.Exists() {
					data.RangeStartTime = types.StringValue(value.String())
				} else {
					data.RangeStartTime = types.StringNull()
				}
				if value := res.Get("rangeEndTime"); value.Exists() {
					data.RangeEndTime = types.StringValue(value.String())
				} else {
					data.RangeEndTime = types.StringNull()
				}
				if value := res.Get("rangeStartDay"); value.Exists() {
					data.RangeStartDay = types.StringValue(value.String())
				} else {
					data.RangeStartDay = types.StringNull()
				}
				if value := res.Get("rangeEndDay"); value.Exists() {
					data.RangeEndDay = types.StringValue(value.String())
				} else {
					data.RangeEndDay = types.StringNull()
				}
				if value := res.Get("dailyStartTime"); value.Exists() {
					data.DailyStartTime = types.StringValue(value.String())
				} else {
					data.DailyStartTime = types.StringNull()
				}
				if value := res.Get("dailyEndTime"); value.Exists() {
					data.DailyEndTime = types.StringValue(value.String())
				} else {
					data.DailyEndTime = types.StringNull()
				}
				if value := res.Get("days"); value.Exists() {
					data.DailyDays = helpers.GetStringSet(value.Array())
				} else {
					data.DailyDays = types.SetNull(types.StringType)
				}
				(*parent).RecurrenceList = append((*parent).RecurrenceList, data)
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
func (data *TimeRanges) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("effectiveStartDateTime"); value.Exists() && !data.StartTime.IsNull() {
			data.StartTime = types.StringValue(value.String())
		} else {
			data.StartTime = types.StringNull()
		}
		if value := res.Get("effectiveEndDateTime"); value.Exists() && !data.EndTime.IsNull() {
			data.EndTime = types.StringValue(value.String())
		} else {
			data.EndTime = types.StringNull()
		}
		for i := 0; i < len(data.RecurrenceList); i++ {
			keys := [...]string{"recurrenceType", "rangeStartTime", "rangeEndTime", "rangeStartDay", "rangeEndDay", "dailyStartTime", "dailyEndTime"}
			keyValues := [...]string{data.RecurrenceList[i].RecurrenceType.ValueString(), data.RecurrenceList[i].RangeStartTime.ValueString(), data.RecurrenceList[i].RangeEndTime.ValueString(), data.RecurrenceList[i].RangeStartDay.ValueString(), data.RecurrenceList[i].RangeEndDay.ValueString(), data.RecurrenceList[i].DailyStartTime.ValueString(), data.RecurrenceList[i].DailyEndTime.ValueString()}

			parent := &data
			data := (*parent).RecurrenceList[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("recurrenceList").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing RecurrenceList[%d] = %+v",
					i,
					(*parent).RecurrenceList[i],
				))
				(*parent).RecurrenceList = slices.Delete((*parent).RecurrenceList, i, i+1)
				i--

				continue
			}
			if value := res.Get("recurrenceType"); value.Exists() && !data.RecurrenceType.IsNull() {
				data.RecurrenceType = types.StringValue(value.String())
			} else {
				data.RecurrenceType = types.StringNull()
			}
			if value := res.Get("rangeStartTime"); value.Exists() && !data.RangeStartTime.IsNull() {
				data.RangeStartTime = types.StringValue(value.String())
			} else {
				data.RangeStartTime = types.StringNull()
			}
			if value := res.Get("rangeEndTime"); value.Exists() && !data.RangeEndTime.IsNull() {
				data.RangeEndTime = types.StringValue(value.String())
			} else {
				data.RangeEndTime = types.StringNull()
			}
			if value := res.Get("rangeStartDay"); value.Exists() && !data.RangeStartDay.IsNull() {
				data.RangeStartDay = types.StringValue(value.String())
			} else {
				data.RangeStartDay = types.StringNull()
			}
			if value := res.Get("rangeEndDay"); value.Exists() && !data.RangeEndDay.IsNull() {
				data.RangeEndDay = types.StringValue(value.String())
			} else {
				data.RangeEndDay = types.StringNull()
			}
			if value := res.Get("dailyStartTime"); value.Exists() && !data.DailyStartTime.IsNull() {
				data.DailyStartTime = types.StringValue(value.String())
			} else {
				data.DailyStartTime = types.StringNull()
			}
			if value := res.Get("dailyEndTime"); value.Exists() && !data.DailyEndTime.IsNull() {
				data.DailyEndTime = types.StringValue(value.String())
			} else {
				data.DailyEndTime = types.StringNull()
			}
			if value := res.Get("days"); value.Exists() && !data.DailyDays.IsNull() {
				data.DailyDays = helpers.GetStringSet(value.Array())
			} else {
				data.DailyDays = types.SetNull(types.StringType)
			}
			(*parent).RecurrenceList[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *TimeRanges) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data *TimeRanges) Clone() TimeRanges {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data TimeRanges) toBodyNonBulk(ctx context.Context, state TimeRanges) string {
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
