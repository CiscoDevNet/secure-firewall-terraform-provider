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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SmartLicense struct {
	Id                 types.String `tfsdk:"id"`
	RegistrationType   types.String `tfsdk:"registration_type"`
	Token              types.String `tfsdk:"token"`
	RegistrationStatus types.String `tfsdk:"registration_status"`
	Force              types.Bool   `tfsdk:"force"`
	RetainRegistration types.Bool   `tfsdk:"retain_registration"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SmartLicense) getPath() string {
	return "/api/fmc_platform/v1/license/smartlicenses"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SmartLicense) toBody(ctx context.Context, state SmartLicense) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.RegistrationType.IsNull() {
		body, _ = sjson.Set(body, "registrationType", data.RegistrationType.ValueString())
	}
	if !data.Token.IsNull() {
		body, _ = sjson.Set(body, "token", data.Token.ValueString())
	}
	if !data.Force.IsNull() {
		body, _ = sjson.Set(body, "dummy_force", data.Force.ValueBool())
	}
	if !data.RetainRegistration.IsNull() {
		body, _ = sjson.Set(body, "dummy_retain_registration", data.RetainRegistration.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SmartLicense) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("regStatus"); value.Exists() {
		data.RegistrationStatus = types.StringValue(value.String())
	} else {
		data.RegistrationStatus = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SmartLicense) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("regStatus"); value.Exists() && !data.RegistrationStatus.IsNull() {
		data.RegistrationStatus = types.StringValue(value.String())
	} else {
		data.RegistrationStatus = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SmartLicense) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.RegistrationStatus.IsUnknown() {
		if value := res.Get("regStatus"); value.Exists() {
			data.RegistrationStatus = types.StringValue(value.String())
		} else {
			data.RegistrationStatus = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
