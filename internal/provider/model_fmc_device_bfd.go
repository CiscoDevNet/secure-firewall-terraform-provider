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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceBFD struct {
	Id                      types.String `tfsdk:"id"`
	Domain                  types.String `tfsdk:"domain"`
	DeviceId                types.String `tfsdk:"device_id"`
	Type                    types.String `tfsdk:"type"`
	HopType                 types.String `tfsdk:"hop_type"`
	BfdTemplateId           types.String `tfsdk:"bfd_template_id"`
	InterfaceLogicalName    types.String `tfsdk:"interface_logical_name"`
	DestinationHostObjectId types.String `tfsdk:"destination_host_object_id"`
	SourceHostObjectId      types.String `tfsdk:"source_host_object_id"`
	InterfaceId             types.String `tfsdk:"interface_id"`
	SlowTimer               types.Int64  `tfsdk:"slow_timer"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionDeviceBFD = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceBFD) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/bfdpolicies", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceBFD) toBody(ctx context.Context, state DeviceBFD) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.HopType.IsNull() {
		body, _ = sjson.Set(body, "hopType", data.HopType.ValueString())
	}
	if !data.BfdTemplateId.IsNull() {
		body, _ = sjson.Set(body, "template.id", data.BfdTemplateId.ValueString())
	}
	if !data.InterfaceLogicalName.IsNull() {
		body, _ = sjson.Set(body, "interface.ifname", data.InterfaceLogicalName.ValueString())
	}
	if !data.DestinationHostObjectId.IsNull() {
		body, _ = sjson.Set(body, "destinationAddress.id", data.DestinationHostObjectId.ValueString())
	}
	if !data.SourceHostObjectId.IsNull() {
		body, _ = sjson.Set(body, "sourceAddress.id", data.SourceHostObjectId.ValueString())
	}
	if !data.InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "interface.id", data.InterfaceId.ValueString())
	}
	if !data.SlowTimer.IsNull() {
		body, _ = sjson.Set(body, "slowTimer", data.SlowTimer.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceBFD) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("hopType"); value.Exists() {
		data.HopType = types.StringValue(value.String())
	} else {
		data.HopType = types.StringNull()
	}
	if value := res.Get("template.id"); value.Exists() {
		data.BfdTemplateId = types.StringValue(value.String())
	} else {
		data.BfdTemplateId = types.StringNull()
	}
	if value := res.Get("interface.ifname"); value.Exists() {
		data.InterfaceLogicalName = types.StringValue(value.String())
	} else {
		data.InterfaceLogicalName = types.StringNull()
	}
	if value := res.Get("destinationAddress.id"); value.Exists() {
		data.DestinationHostObjectId = types.StringValue(value.String())
	} else {
		data.DestinationHostObjectId = types.StringNull()
	}
	if value := res.Get("sourceAddress.id"); value.Exists() {
		data.SourceHostObjectId = types.StringValue(value.String())
	} else {
		data.SourceHostObjectId = types.StringNull()
	}
	if value := res.Get("interface.id"); value.Exists() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("slowTimer"); value.Exists() {
		data.SlowTimer = types.Int64Value(value.Int())
	} else {
		data.SlowTimer = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceBFD) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("hopType"); value.Exists() && !data.HopType.IsNull() {
		data.HopType = types.StringValue(value.String())
	} else {
		data.HopType = types.StringNull()
	}
	if value := res.Get("template.id"); value.Exists() && !data.BfdTemplateId.IsNull() {
		data.BfdTemplateId = types.StringValue(value.String())
	} else {
		data.BfdTemplateId = types.StringNull()
	}
	if value := res.Get("interface.ifname"); value.Exists() && !data.InterfaceLogicalName.IsNull() {
		data.InterfaceLogicalName = types.StringValue(value.String())
	} else {
		data.InterfaceLogicalName = types.StringNull()
	}
	if value := res.Get("destinationAddress.id"); value.Exists() && !data.DestinationHostObjectId.IsNull() {
		data.DestinationHostObjectId = types.StringValue(value.String())
	} else {
		data.DestinationHostObjectId = types.StringNull()
	}
	if value := res.Get("sourceAddress.id"); value.Exists() && !data.SourceHostObjectId.IsNull() {
		data.SourceHostObjectId = types.StringValue(value.String())
	} else {
		data.SourceHostObjectId = types.StringNull()
	}
	if value := res.Get("interface.id"); value.Exists() && !data.InterfaceId.IsNull() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("slowTimer"); value.Exists() && !data.SlowTimer.IsNull() {
		data.SlowTimer = types.Int64Value(value.Int())
	} else {
		data.SlowTimer = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceBFD) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
