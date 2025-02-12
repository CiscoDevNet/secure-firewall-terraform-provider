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
	"slices"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &FilePolicyResource{}
	_ resource.ResourceWithImportState = &FilePolicyResource{}
)

func NewFilePolicyResource() resource.Resource {
	return &FilePolicyResource{}
}

type FilePolicyResource struct {
	client *fmc.Client
}

func (r *FilePolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file_policy"
}

func (r *FilePolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages File Policy with corresponding rules.").AddMinimumVersionHeaderDescription().AddMinimumVersionAnyDescription().AddMinimumVersionCreateDescription("7.4").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of file policy.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("File policy description.").String,
				Optional:            true,
			},
			"first_time_file_analysis": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Analyze first-seen files while AMP cloud disposition is pending").String,
				Optional:            true,
			},
			"custom_detection_list": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable custom detection list").String,
				Optional:            true,
			},
			"clean_list": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable clean list").String,
				Optional:            true,
			},
			"threat_score": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If AMP Cloud disposition is Unknown, override disposition based upon threat score.").AddStringEnumDescription("DISABLED", "MEDIUM", "High", "VERY_HIGH").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DISABLED", "MEDIUM", "High", "VERY_HIGH"),
				},
			},
			"inspect_archives": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Inspect Archives").String,
				Optional:            true,
			},
			"block_encrypted_archives": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Block encrypted archives").String,
				Optional:            true,
			},
			"block_uninspectable_archives": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Block uninspectable Archives").String,
				Optional:            true,
			},
			"max_archive_depth": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Max archive depth").AddIntegerRangeDescription(1, 3).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3),
				},
			},
			"file_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of file rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of File Rule").String,
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of File Rule.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"application_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Defines a protocol for file inspection.").AddStringEnumDescription("ANY", "HTTP", "SMTP", "IMAP", "POP3", "FTP", "SMB").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ANY", "HTTP", "SMTP", "IMAP", "POP3", "FTP", "SMB"),
							},
						},
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Action to be performed on a file.").AddStringEnumDescription("DETECT", "BLOCK_WITH_RESET", "DETECT_MALWARE", "BLOCK_MALWARE_WITH_RESET").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("DETECT", "BLOCK_WITH_RESET", "DETECT_MALWARE", "BLOCK_MALWARE_WITH_RESET"),
							},
						},
						"store_files": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of file dispositions that should be stored (MALWARE, CUSTOM, CLEAN, UNKNOWN).").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"direction_of_transfer": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Direction of file transfer.").AddStringEnumDescription("ANY", "UPLOAD", "DOWNLOAD").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ANY", "UPLOAD", "DOWNLOAD"),
							},
						},
						"file_categories": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Defines a list of file categories for inspection.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The id of file category.").String,
										Required:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The name of file category.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The type of file category.").AddDefaultValueDescription("FileCategory").String,
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("FileCategory"),
									},
								},
							},
						},
						"file_types": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Defines a list of file types for inspection.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The id of file type.").String,
										Required:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The name of file type.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The name of file type.").AddDefaultValueDescription("FileType").String,
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("FileType"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *FilePolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *FilePolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersionCreateFilePolicy) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support File Policy creation, minumum required version is 7.4", r.client.FMCVersion))
		return
	}
	var plan FilePolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	planBody := plan.toBody(ctx, FilePolicy{})

	// Create object
	body := planBody
	body, _ = sjson.Delete(body, "dummy_file_rules")

	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	read, err := r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, read)

	state := plan
	state.FileRules = nil

	state, diags = r.updateSubresources(ctx, req.Plan, plan, planBody, tfsdk.State{}, state)
	resp.Diagnostics.Append(diags...)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *FilePolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FilePolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	resGet, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resGet.String()))
		return
	}

	resRules, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/filerules?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resGet.String()))
		return
	}

	s := resGet.String()

	replaceRules := resRules.Get("items").String()
	if replaceRules == "" {
		replaceRules = "[]"
	}
	s, _ = sjson.SetRaw(s, "dummy_file_rules", replaceRules)

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	res := gjson.Parse(s)

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *FilePolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FilePolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	planBody := plan.toBody(ctx, state)
	body := planBody
	body, _ = sjson.Delete(body, "dummy_file_rules")

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.fromBodyUnknowns(ctx, res)

	orig := state
	state = plan
	state.FileRules = orig.FileRules

	state, diags = r.updateSubresources(ctx, req.Plan, plan, planBody, req.State, state)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *FilePolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FilePolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *FilePolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

func (r *FilePolicyResource) updateSubresources(ctx context.Context, tfsdkPlan tfsdk.Plan, plan FilePolicy, planBody string, tfsdkState tfsdk.State, state FilePolicy) (FilePolicy, diag.Diagnostics) {
	var diags diag.Diagnostics

	p := gjson.Parse(planBody)
	bodyRules := p.Get("dummy_file_rules").Array()

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	keptRules := 0

	err := r.truncateRulesAt(ctx, &state, keptRules, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	if len(plan.FileRules) == 0 {
		state.FileRules = plan.FileRules
	}

	err = r.createRulesAt(ctx, plan, bodyRules, keptRules, &state, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	return state, diags
}

func (r *FilePolicyResource) truncateRulesAt(ctx context.Context, state *FilePolicy, kept int, reqMods ...func(*fmc.Req)) error {
	// File rules do not support bulk operations, hence we delete one by one
	for i := kept; i < len(state.FileRules); i++ {
		id := state.FileRules[i].Id.ValueString()

		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/filerules/"+url.QueryEscape(id), reqMods...)
		if err != nil {
			// Remove already deleted objects from the state
			state.FileRules = slices.Delete(state.FileRules, kept, i-1)
			return fmt.Errorf("failed to delete object (DELETE), got error: %s, %s", err, res.String())
		}
	}

	// Remove all objects from the state
	if len(state.FileRules)-1 > kept {
		state.FileRules = slices.Delete(state.FileRules, kept, len(state.FileRules)-1)
	}

	return nil
}

func (r *FilePolicyResource) createRulesAt(ctx context.Context, plan FilePolicy, body []gjson.Result, startIndex int, state *FilePolicy, reqMods ...func(*fmc.Req)) error {
	// File rules do not supoprt bulk operations, hence we create one by one
	for i := startIndex; i < len(body); i++ {
		rule := body[i].String()
		res, err := r.client.Post(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+"/filerules", rule, reqMods...)
		if err != nil {
			return fmt.Errorf("failed to configure object (POST), got error: %s, %s", err, res.String())
		}
		item := plan.FileRules[i]
		item.Id = types.StringValue(res.Get("id").String())

		if len(state.FileRules) <= i {
			state.FileRules = append(state.FileRules, item)
		} else {
			state.FileRules[i] = item
		}
	}

	return nil
}
