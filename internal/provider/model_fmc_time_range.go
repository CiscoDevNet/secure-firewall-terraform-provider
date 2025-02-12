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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type TimeRange struct {
	Id             types.String              `tfsdk:"id"`
	Domain         types.String              `tfsdk:"domain"`
	Name           types.String              `tfsdk:"name"`
	Type           types.String              `tfsdk:"type"`
	Description    types.String              `tfsdk:"description"`
	StartTime      types.String              `tfsdk:"start_time"`
	EndTime        types.String              `tfsdk:"end_time"`
	RecurrenceList []TimeRangeRecurrenceList `tfsdk:"recurrence_list"`
}

type TimeRangeRecurrenceList struct {
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

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data TimeRange) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/timeranges"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data TimeRange) toBody(ctx context.Context, state TimeRange) string {
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
	if !data.StartTime.IsNull() {
		body, _ = sjson.Set(body, "effectiveStartDateTime", data.StartTime.ValueString())
	}
	if !data.EndTime.IsNull() {
		body, _ = sjson.Set(body, "effectiveEndDateTime", data.EndTime.ValueString())
	}
	if len(data.RecurrenceList) > 0 {
		body, _ = sjson.Set(body, "recurrenceList", []interface{}{})
		for _, item := range data.RecurrenceList {
			itemBody := ""
			if !item.RecurrenceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "recurrenceType", item.RecurrenceType.ValueString())
			}
			if !item.RangeStartTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rangeStartTime", item.RangeStartTime.ValueString())
			}
			if !item.RangeEndTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rangeEndTime", item.RangeEndTime.ValueString())
			}
			if !item.RangeStartDay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rangeStartDay", item.RangeStartDay.ValueString())
			}
			if !item.RangeEndDay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rangeEndDay", item.RangeEndDay.ValueString())
			}
			if !item.DailyStartTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dailyStartTime", item.DailyStartTime.ValueString())
			}
			if !item.DailyEndTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dailyEndTime", item.DailyEndTime.ValueString())
			}
			if !item.DailyDays.IsNull() {
				var values []string
				item.DailyDays.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "days", values)
			}
			body, _ = sjson.SetRaw(body, "recurrenceList.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *TimeRange) fromBody(ctx context.Context, res gjson.Result) {
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
		data.RecurrenceList = make([]TimeRangeRecurrenceList, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := TimeRangeRecurrenceList{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *TimeRange) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	{
		l := len(res.Get("recurrenceList").Array())
		tflog.Debug(ctx, fmt.Sprintf("recurrenceList array resizing from %d to %d", len(data.RecurrenceList), l))
		for i := len(data.RecurrenceList); i < l; i++ {
			data.RecurrenceList = append(data.RecurrenceList, TimeRangeRecurrenceList{})
		}
		if len(data.RecurrenceList) > l {
			data.RecurrenceList = data.RecurrenceList[:l]
		}
	}
	for i := range data.RecurrenceList {
		parent := &data
		data := (*parent).RecurrenceList[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("recurrenceList.%d", i))
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
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *TimeRange) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
