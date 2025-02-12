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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type BFDTemplate struct {
	Id                               types.String `tfsdk:"id"`
	Domain                           types.String `tfsdk:"domain"`
	Name                             types.String `tfsdk:"name"`
	Type                             types.String `tfsdk:"type"`
	HopType                          types.String `tfsdk:"hop_type"`
	Echo                             types.String `tfsdk:"echo"`
	IntervalTime                     types.String `tfsdk:"interval_time"`
	MinTransmit                      types.Int64  `tfsdk:"min_transmit"`
	TxRxMultiplier                   types.Int64  `tfsdk:"tx_rx_multiplier"`
	MinReceive                       types.Int64  `tfsdk:"min_receive"`
	AuthenticationPassword           types.String `tfsdk:"authentication_password"`
	AuthenticationKeyId              types.Int64  `tfsdk:"authentication_key_id"`
	AuthenticationType               types.String `tfsdk:"authentication_type"`
	AuthenticationPasswordEncryption types.String `tfsdk:"authentication_password_encryption"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBFDTemplate = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data BFDTemplate) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/bfdtemplates"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data BFDTemplate) toBody(ctx context.Context, state BFDTemplate) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "BFDTemplate")
	if !data.HopType.IsNull() {
		body, _ = sjson.Set(body, "hopType", data.HopType.ValueString())
	}
	if !data.Echo.IsNull() {
		body, _ = sjson.Set(body, "echo", data.Echo.ValueString())
	}
	if !data.IntervalTime.IsNull() {
		body, _ = sjson.Set(body, "txRxInterval", data.IntervalTime.ValueString())
	}
	if !data.MinTransmit.IsNull() {
		body, _ = sjson.Set(body, "minTransmit", data.MinTransmit.ValueInt64())
	}
	if !data.TxRxMultiplier.IsNull() {
		body, _ = sjson.Set(body, "txRxMultiplier", data.TxRxMultiplier.ValueInt64())
	}
	if !data.MinReceive.IsNull() {
		body, _ = sjson.Set(body, "minReceive", data.MinReceive.ValueInt64())
	}
	if !data.AuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "authentication.authKey", data.AuthenticationPassword.ValueString())
	}
	if !data.AuthenticationKeyId.IsNull() {
		body, _ = sjson.Set(body, "authentication.authKeyId", data.AuthenticationKeyId.ValueInt64())
	}
	if !data.AuthenticationType.IsNull() {
		body, _ = sjson.Set(body, "authentication.authType", data.AuthenticationType.ValueString())
	}
	if !data.AuthenticationPasswordEncryption.IsNull() {
		body, _ = sjson.Set(body, "authentication.pwdEncryption", data.AuthenticationPasswordEncryption.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *BFDTemplate) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("hopType"); value.Exists() {
		data.HopType = types.StringValue(value.String())
	} else {
		data.HopType = types.StringNull()
	}
	if value := res.Get("echo"); value.Exists() {
		data.Echo = types.StringValue(value.String())
	} else {
		data.Echo = types.StringNull()
	}
	if value := res.Get("txRxInterval"); value.Exists() {
		data.IntervalTime = types.StringValue(value.String())
	} else {
		data.IntervalTime = types.StringNull()
	}
	if value := res.Get("minTransmit"); value.Exists() {
		data.MinTransmit = types.Int64Value(value.Int())
	} else {
		data.MinTransmit = types.Int64Null()
	}
	if value := res.Get("txRxMultiplier"); value.Exists() {
		data.TxRxMultiplier = types.Int64Value(value.Int())
	} else {
		data.TxRxMultiplier = types.Int64Null()
	}
	if value := res.Get("minReceive"); value.Exists() {
		data.MinReceive = types.Int64Value(value.Int())
	} else {
		data.MinReceive = types.Int64Null()
	}
	if value := res.Get("authentication.authKeyId"); value.Exists() {
		data.AuthenticationKeyId = types.Int64Value(value.Int())
	} else {
		data.AuthenticationKeyId = types.Int64Null()
	}
	if value := res.Get("authentication.authType"); value.Exists() {
		data.AuthenticationType = types.StringValue(value.String())
	} else {
		data.AuthenticationType = types.StringNull()
	}
	if value := res.Get("authentication.pwdEncryption"); value.Exists() {
		data.AuthenticationPasswordEncryption = types.StringValue(value.String())
	} else {
		data.AuthenticationPasswordEncryption = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *BFDTemplate) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("hopType"); value.Exists() && !data.HopType.IsNull() {
		data.HopType = types.StringValue(value.String())
	} else {
		data.HopType = types.StringNull()
	}
	if value := res.Get("echo"); value.Exists() && !data.Echo.IsNull() {
		data.Echo = types.StringValue(value.String())
	} else {
		data.Echo = types.StringNull()
	}
	if value := res.Get("txRxInterval"); value.Exists() && !data.IntervalTime.IsNull() {
		data.IntervalTime = types.StringValue(value.String())
	} else {
		data.IntervalTime = types.StringNull()
	}
	if value := res.Get("minTransmit"); value.Exists() && !data.MinTransmit.IsNull() {
		data.MinTransmit = types.Int64Value(value.Int())
	} else {
		data.MinTransmit = types.Int64Null()
	}
	if value := res.Get("txRxMultiplier"); value.Exists() && !data.TxRxMultiplier.IsNull() {
		data.TxRxMultiplier = types.Int64Value(value.Int())
	} else {
		data.TxRxMultiplier = types.Int64Null()
	}
	if value := res.Get("minReceive"); value.Exists() && !data.MinReceive.IsNull() {
		data.MinReceive = types.Int64Value(value.Int())
	} else {
		data.MinReceive = types.Int64Null()
	}
	if value := res.Get("authentication.authKeyId"); value.Exists() && !data.AuthenticationKeyId.IsNull() {
		data.AuthenticationKeyId = types.Int64Value(value.Int())
	} else {
		data.AuthenticationKeyId = types.Int64Null()
	}
	if value := res.Get("authentication.authType"); value.Exists() && !data.AuthenticationType.IsNull() {
		data.AuthenticationType = types.StringValue(value.String())
	} else {
		data.AuthenticationType = types.StringNull()
	}
	if value := res.Get("authentication.pwdEncryption"); value.Exists() && !data.AuthenticationPasswordEncryption.IsNull() {
		data.AuthenticationPasswordEncryption = types.StringValue(value.String())
	} else {
		data.AuthenticationPasswordEncryption = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *BFDTemplate) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
