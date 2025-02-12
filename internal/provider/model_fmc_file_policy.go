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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FilePolicy struct {
	Id                         types.String          `tfsdk:"id"`
	Domain                     types.String          `tfsdk:"domain"`
	Name                       types.String          `tfsdk:"name"`
	Type                       types.String          `tfsdk:"type"`
	Description                types.String          `tfsdk:"description"`
	FirstTimeFileAnalysis      types.Bool            `tfsdk:"first_time_file_analysis"`
	CustomDetectionList        types.Bool            `tfsdk:"custom_detection_list"`
	CleanList                  types.Bool            `tfsdk:"clean_list"`
	ThreatScore                types.String          `tfsdk:"threat_score"`
	InspectArchives            types.Bool            `tfsdk:"inspect_archives"`
	BlockEncryptedArchives     types.Bool            `tfsdk:"block_encrypted_archives"`
	BlockUninspectableArchives types.Bool            `tfsdk:"block_uninspectable_archives"`
	MaxArchiveDepth            types.Int64           `tfsdk:"max_archive_depth"`
	FileRules                  []FilePolicyFileRules `tfsdk:"file_rules"`
}

type FilePolicyFileRules struct {
	Id                  types.String                        `tfsdk:"id"`
	Type                types.String                        `tfsdk:"type"`
	ApplicationProtocol types.String                        `tfsdk:"application_protocol"`
	Action              types.String                        `tfsdk:"action"`
	StoreFiles          types.Set                           `tfsdk:"store_files"`
	DirectionOfTransfer types.String                        `tfsdk:"direction_of_transfer"`
	FileCategories      []FilePolicyFileRulesFileCategories `tfsdk:"file_categories"`
	FileTypes           []FilePolicyFileRulesFileTypes      `tfsdk:"file_types"`
}

type FilePolicyFileRulesFileCategories struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}
type FilePolicyFileRulesFileTypes struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateFilePolicy = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FilePolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/filepolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FilePolicy) toBody(ctx context.Context, state FilePolicy) string {
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
	if !data.FirstTimeFileAnalysis.IsNull() {
		body, _ = sjson.Set(body, "firstTimeFileAnalysis", data.FirstTimeFileAnalysis.ValueBool())
	}
	if !data.CustomDetectionList.IsNull() {
		body, _ = sjson.Set(body, "customDetectionList", data.CustomDetectionList.ValueBool())
	}
	if !data.CleanList.IsNull() {
		body, _ = sjson.Set(body, "cleanList", data.CleanList.ValueBool())
	}
	if !data.ThreatScore.IsNull() {
		body, _ = sjson.Set(body, "threatScore", data.ThreatScore.ValueString())
	}
	if !data.InspectArchives.IsNull() {
		body, _ = sjson.Set(body, "inspectArchives", data.InspectArchives.ValueBool())
	}
	if !data.BlockEncryptedArchives.IsNull() {
		body, _ = sjson.Set(body, "blockEncryptedArchives", data.BlockEncryptedArchives.ValueBool())
	}
	if !data.BlockUninspectableArchives.IsNull() {
		body, _ = sjson.Set(body, "archiveDepthAction", data.BlockUninspectableArchives.ValueBool())
	}
	if !data.MaxArchiveDepth.IsNull() {
		body, _ = sjson.Set(body, "archiveDepth", data.MaxArchiveDepth.ValueInt64())
	}
	if len(data.FileRules) > 0 {
		body, _ = sjson.Set(body, "dummy_file_rules", []interface{}{})
		for _, item := range data.FileRules {
			itemBody := ""
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.ApplicationProtocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.ApplicationProtocol.ValueString())
			}
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.StoreFiles.IsNull() {
				var values []string
				item.StoreFiles.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "storeFiles", values)
			}
			if !item.DirectionOfTransfer.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "direction", item.DirectionOfTransfer.ValueString())
			}
			if len(item.FileCategories) > 0 {
				itemBody, _ = sjson.Set(itemBody, "fileCategories", []interface{}{})
				for _, childItem := range item.FileCategories {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "fileCategories.-1", itemChildBody)
				}
			}
			if len(item.FileTypes) > 0 {
				itemBody, _ = sjson.Set(itemBody, "fileTypes", []interface{}{})
				for _, childItem := range item.FileTypes {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "fileTypes.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "dummy_file_rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FilePolicy) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("firstTimeFileAnalysis"); value.Exists() {
		data.FirstTimeFileAnalysis = types.BoolValue(value.Bool())
	} else {
		data.FirstTimeFileAnalysis = types.BoolNull()
	}
	if value := res.Get("customDetectionList"); value.Exists() {
		data.CustomDetectionList = types.BoolValue(value.Bool())
	} else {
		data.CustomDetectionList = types.BoolNull()
	}
	if value := res.Get("cleanList"); value.Exists() {
		data.CleanList = types.BoolValue(value.Bool())
	} else {
		data.CleanList = types.BoolNull()
	}
	if value := res.Get("threatScore"); value.Exists() {
		data.ThreatScore = types.StringValue(value.String())
	} else {
		data.ThreatScore = types.StringNull()
	}
	if value := res.Get("inspectArchives"); value.Exists() {
		data.InspectArchives = types.BoolValue(value.Bool())
	} else {
		data.InspectArchives = types.BoolNull()
	}
	if value := res.Get("blockEncryptedArchives"); value.Exists() {
		data.BlockEncryptedArchives = types.BoolValue(value.Bool())
	} else {
		data.BlockEncryptedArchives = types.BoolNull()
	}
	if value := res.Get("archiveDepthAction"); value.Exists() {
		data.BlockUninspectableArchives = types.BoolValue(value.Bool())
	} else {
		data.BlockUninspectableArchives = types.BoolNull()
	}
	if value := res.Get("archiveDepth"); value.Exists() {
		data.MaxArchiveDepth = types.Int64Value(value.Int())
	} else {
		data.MaxArchiveDepth = types.Int64Null()
	}
	if value := res.Get("dummy_file_rules"); value.Exists() {
		data.FileRules = make([]FilePolicyFileRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FilePolicyFileRules{}
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
			if value := res.Get("protocol"); value.Exists() {
				data.ApplicationProtocol = types.StringValue(value.String())
			} else {
				data.ApplicationProtocol = types.StringNull()
			}
			if value := res.Get("action"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("storeFiles"); value.Exists() {
				data.StoreFiles = helpers.GetStringSet(value.Array())
			} else {
				data.StoreFiles = types.SetNull(types.StringType)
			}
			if value := res.Get("direction"); value.Exists() {
				data.DirectionOfTransfer = types.StringValue(value.String())
			} else {
				data.DirectionOfTransfer = types.StringNull()
			}
			if value := res.Get("fileCategories"); value.Exists() {
				data.FileCategories = make([]FilePolicyFileRulesFileCategories, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := FilePolicyFileRulesFileCategories{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("name"); value.Exists() {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringValue("FileCategory")
					}
					(*parent).FileCategories = append((*parent).FileCategories, data)
					return true
				})
			}
			if value := res.Get("fileTypes"); value.Exists() {
				data.FileTypes = make([]FilePolicyFileRulesFileTypes, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := FilePolicyFileRulesFileTypes{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("name"); value.Exists() {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringValue("FileType")
					}
					(*parent).FileTypes = append((*parent).FileTypes, data)
					return true
				})
			}
			(*parent).FileRules = append((*parent).FileRules, data)
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
func (data *FilePolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("firstTimeFileAnalysis"); value.Exists() && !data.FirstTimeFileAnalysis.IsNull() {
		data.FirstTimeFileAnalysis = types.BoolValue(value.Bool())
	} else {
		data.FirstTimeFileAnalysis = types.BoolNull()
	}
	if value := res.Get("customDetectionList"); value.Exists() && !data.CustomDetectionList.IsNull() {
		data.CustomDetectionList = types.BoolValue(value.Bool())
	} else {
		data.CustomDetectionList = types.BoolNull()
	}
	if value := res.Get("cleanList"); value.Exists() && !data.CleanList.IsNull() {
		data.CleanList = types.BoolValue(value.Bool())
	} else {
		data.CleanList = types.BoolNull()
	}
	if value := res.Get("threatScore"); value.Exists() && !data.ThreatScore.IsNull() {
		data.ThreatScore = types.StringValue(value.String())
	} else {
		data.ThreatScore = types.StringNull()
	}
	if value := res.Get("inspectArchives"); value.Exists() && !data.InspectArchives.IsNull() {
		data.InspectArchives = types.BoolValue(value.Bool())
	} else {
		data.InspectArchives = types.BoolNull()
	}
	if value := res.Get("blockEncryptedArchives"); value.Exists() && !data.BlockEncryptedArchives.IsNull() {
		data.BlockEncryptedArchives = types.BoolValue(value.Bool())
	} else {
		data.BlockEncryptedArchives = types.BoolNull()
	}
	if value := res.Get("archiveDepthAction"); value.Exists() && !data.BlockUninspectableArchives.IsNull() {
		data.BlockUninspectableArchives = types.BoolValue(value.Bool())
	} else {
		data.BlockUninspectableArchives = types.BoolNull()
	}
	if value := res.Get("archiveDepth"); value.Exists() && !data.MaxArchiveDepth.IsNull() {
		data.MaxArchiveDepth = types.Int64Value(value.Int())
	} else {
		data.MaxArchiveDepth = types.Int64Null()
	}
	{
		l := len(res.Get("dummy_file_rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("dummy_file_rules array resizing from %d to %d", len(data.FileRules), l))
		for i := len(data.FileRules); i < l; i++ {
			data.FileRules = append(data.FileRules, FilePolicyFileRules{})
		}
		if len(data.FileRules) > l {
			data.FileRules = data.FileRules[:l]
		}
	}
	for i := range data.FileRules {
		parent := &data
		data := (*parent).FileRules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("dummy_file_rules.%d", i))
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
		if value := res.Get("protocol"); value.Exists() && !data.ApplicationProtocol.IsNull() {
			data.ApplicationProtocol = types.StringValue(value.String())
		} else {
			data.ApplicationProtocol = types.StringNull()
		}
		if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
			data.Action = types.StringValue(value.String())
		} else {
			data.Action = types.StringNull()
		}
		if value := res.Get("storeFiles"); value.Exists() && !data.StoreFiles.IsNull() {
			data.StoreFiles = helpers.GetStringSet(value.Array())
		} else {
			data.StoreFiles = types.SetNull(types.StringType)
		}
		if value := res.Get("direction"); value.Exists() && !data.DirectionOfTransfer.IsNull() {
			data.DirectionOfTransfer = types.StringValue(value.String())
		} else {
			data.DirectionOfTransfer = types.StringNull()
		}
		for i := 0; i < len(data.FileCategories); i++ {
			keys := [...]string{"id", "name", "type"}
			keyValues := [...]string{data.FileCategories[i].Id.ValueString(), data.FileCategories[i].Name.ValueString(), data.FileCategories[i].Type.ValueString()}

			parent := &data
			data := (*parent).FileCategories[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("fileCategories").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing FileCategories[%d] = %+v",
					i,
					(*parent).FileCategories[i],
				))
				(*parent).FileCategories = slices.Delete((*parent).FileCategories, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else if data.Type.ValueString() != "FileCategory" {
				data.Type = types.StringNull()
			}
			(*parent).FileCategories[i] = data
		}
		for i := 0; i < len(data.FileTypes); i++ {
			keys := [...]string{"id", "name", "type"}
			keyValues := [...]string{data.FileTypes[i].Id.ValueString(), data.FileTypes[i].Name.ValueString(), data.FileTypes[i].Type.ValueString()}

			parent := &data
			data := (*parent).FileTypes[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("fileTypes").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing FileTypes[%d] = %+v",
					i,
					(*parent).FileTypes[i],
				))
				(*parent).FileTypes = slices.Delete((*parent).FileTypes, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else if data.Type.ValueString() != "FileType" {
				data.Type = types.StringNull()
			}
			(*parent).FileTypes[i] = data
		}
		(*parent).FileRules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FilePolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	for i := range data.FileRules {
		r := res.Get(fmt.Sprintf("dummy_file_rules.%d", i))
		if v := data.FileRules[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.FileRules[i] = v
		}
		if v := data.FileRules[i]; v.Type.IsUnknown() {
			if value := r.Get("type"); value.Exists() {
				v.Type = types.StringValue(value.String())
			} else {
				v.Type = types.StringNull()
			}
			data.FileRules[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
