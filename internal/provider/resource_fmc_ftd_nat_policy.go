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
	"time"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
	_ resource.Resource                = &FTDNATPolicyResource{}
	_ resource.ResourceWithImportState = &FTDNATPolicyResource{}
)

func NewFTDNATPolicyResource() resource.Resource {
	return &FTDNATPolicyResource{}
}

type FTDNATPolicyResource struct {
	client *fmc.Client
}

func (r *FTDNATPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_nat_policy"
}

func (r *FTDNATPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Network Address Translation (NAT) policy with corresponding Manual and Auto NAT rules.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the FTD Network Address Translation (NAT) policy.").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'FTDNatPolicy'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"manual_nat_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of manual NAT rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the manual nat rule.").String,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of Manual NAT rule.").String,
							Optional:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates if the rule is enabled.").String,
							Optional:            true,
						},
						"section": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of section to which the rule belongs.").AddStringEnumDescription("BEFORE_AUTO", "AFTER_AUTO").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("BEFORE_AUTO", "AFTER_AUTO"),
							},
						},
						"nat_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the rule").AddStringEnumDescription("STATIC", "DYNAMIC").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("STATIC", "DYNAMIC"),
							},
						},
						"fall_through": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Fallthrough to Interface PAT (Destination Interface)").String,
							Optional:            true,
						},
						"interface_in_original_destination": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use interface address as original destination").String,
							Optional:            true,
						},
						"interface_in_translated_source": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translate source network to destination interface address").String,
							Optional:            true,
						},
						"ipv6": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use the IPv6 address of the destination interface for interface PAT.").String,
							Optional:            true,
						},
						"net_to_net": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Net to Net Mapping").String,
							Optional:            true,
						},
						"no_proxy_arp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not proxy ARP on Destination Interface").String,
							Optional:            true,
						},
						"unidirectional": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the rule is unidirectional").String,
							Optional:            true,
						},
						"source_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of source security zone or interface group").String,
							Optional:            true,
						},
						"original_source_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original source network object (host, network or range)").String,
							Optional:            true,
						},
						"original_source_port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original source port object").String,
							Optional:            true,
						},
						"original_destination_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original destination network object (host, network or range)").String,
							Optional:            true,
						},
						"original_destination_port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original destination port object").String,
							Optional:            true,
						},
						"route_lookup": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Perform Route Lookup for Destination Interface").String,
							Optional:            true,
						},
						"destination_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of destination security zone or interface group").String,
							Optional:            true,
						},
						"translated_source_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated source network object (host, network or range)").String,
							Optional:            true,
						},
						"translated_source_port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated source port object").String,
							Optional:            true,
						},
						"translate_dns": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translate DNS replies that match this rule").String,
							Optional:            true,
						},
						"translated_destination_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated destination network object (host, network or range)").String,
							Optional:            true,
						},
						"translated_destination_port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated destination port object").String,
							Optional:            true,
						},
					},
				},
			},
			"auto_nat_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The list of auto NAT rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Auto NAT rule.").String,
							Computed:            true,
						},
						"nat_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the rule").AddStringEnumDescription("STATIC", "DYNAMIC").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("STATIC", "DYNAMIC"),
							},
						},
						"destination_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of destination security zone or interface group").String,
							Optional:            true,
						},
						"fall_through": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Fallthrough to Interface PAT (Destination Interface)").String,
							Optional:            true,
						},
						"ipv6": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use the IPv6 address of the destination interface for interface PAT.").String,
							Optional:            true,
						},
						"net_to_net": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Net to Net Mapping").String,
							Optional:            true,
						},
						"no_proxy_arp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not proxy ARP on Destination Interface").String,
							Optional:            true,
						},
						"original_network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original network object (host, network or range)").String,
							Optional:            true,
						},
						"original_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Original port number").String,
							Optional:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service protocol").AddStringEnumDescription("TCP", "UDP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("TCP", "UDP"),
							},
						},
						"perform_route_lookup": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Perform Route Lookup for Destination Interface").String,
							Optional:            true,
						},
						"source_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of source security zone or interface group").String,
							Optional:            true,
						},
						"translate_dns": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translate DNS replies that match this rule").String,
							Optional:            true,
						},
						"translated_network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated network object (host, network or range)").String,
							Optional:            true,
						},
						"translated_network_is_destination_interface": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translate source network to destination interface address").String,
							Optional:            true,
						},
						"translated_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translated port number").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *FTDNATPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *FTDNATPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FTDNATPolicy

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

	planBody := plan.toBody(ctx, FTDNATPolicy{})

	// Create object
	body := planBody
	body, _ = sjson.Delete(body, "dummy_manual_nat_rules")
	body, _ = sjson.Delete(body, "dummy_auto_nat_rules")

	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))

		res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Also, cannot DELETE a hanging policy object, got error: %s, %s", err, res.String()))
		}
		return
	}

	plan.fromBodyUnknowns(ctx, res)

	state := plan
	state.AutoNatRules = nil
	state.ManualNatRules = nil

	state, diags = r.updateSubresources(ctx, req.Plan, plan, planBody, tfsdk.State{}, state)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *FTDNATPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FTDNATPolicy

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

	resManualRules, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/manualnatrules?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resManualRules.String()))
		return
	}

	resAutoRules, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/autonatrules?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resAutoRules.String()))
		return
	}

	s := resGet.String()

	replaceManualRules := resManualRules.Get("items").String()
	if replaceManualRules == "" {
		replaceManualRules = "[]"
	}
	s, _ = sjson.SetRaw(s, "dummy_manual_nat_rules", replaceManualRules)

	replaceAutoRules := resAutoRules.Get("items").String()
	if replaceAutoRules == "" {
		replaceAutoRules = "[]"
	}
	s, _ = sjson.SetRaw(s, "dummy_auto_nat_rules", replaceAutoRules)

	res := gjson.Parse(s)

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

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

func (r *FTDNATPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FTDNATPolicy

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
	body, _ = sjson.Delete(body, "dummy_manual_nat_rules")
	body, _ = sjson.Delete(body, "dummy_auto_nat_rules")

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.fromBodyUnknowns(ctx, res)

	orig := state
	state = plan
	state.ManualNatRules, state.AutoNatRules = orig.ManualNatRules, orig.AutoNatRules

	state, diags = r.updateSubresources(ctx, req.Plan, plan, planBody, req.State, state)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *FTDNATPolicyResource) updateSubresources(ctx context.Context, tfsdkPlan tfsdk.Plan, plan FTDNATPolicy, planBody string, tfsdkState tfsdk.State, state FTDNATPolicy) (FTDNATPolicy, diag.Diagnostics) {
	var diags diag.Diagnostics

	p := gjson.Parse(planBody)
	bodyManualNatRules := p.Get("dummy_manual_nat_rules").Array()
	bodyAutoNatRules := p.Get("dummy_auto_nat_rules").Array()

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	// TODO: For now we remove all the rules
	keptManualNatRules := 0
	keptAutoNatRules := 0

	err := r.truncateManualNatRulesAt(ctx, &state, keptManualNatRules, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	err = r.truncateAutoNatRulesAt(ctx, &state, keptAutoNatRules, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	if len(plan.ManualNatRules) == 0 {
		state.ManualNatRules = plan.ManualNatRules
	}

	if len(plan.AutoNatRules) == 0 {
		state.AutoNatRules = plan.AutoNatRules
	}

	err = r.createManualNatRulesAt(ctx, plan, bodyManualNatRules, keptManualNatRules, &state, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	err = r.createAutoNatRulesAt(ctx, plan, bodyAutoNatRules, keptAutoNatRules, &state, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	return state, diags
}

func (r *FTDNATPolicyResource) truncateManualNatRulesAt(ctx context.Context, state *FTDNATPolicy, kept int, reqMods ...func(*fmc.Req)) error {
	var b strings.Builder
	var bulks []string
	var counts []int
	count := 0

	for i := kept; i < len(state.ManualNatRules); i++ {
		b.WriteString(state.ManualNatRules[i].Id.ValueString() + ",")
		count++
		if b.Len() >= maxUrlParamLength {
			bulks = append(bulks, b.String())
			counts = append(counts, count)
			b.Reset()
			count = 0
		}
	}

	if b.Len() > 0 {
		bulks = append(bulks, b.String())
		counts = append(counts, count)
	}

	defer func() {
		time.Sleep(2 * time.Second)
	}()

	for i, bulk := range bulks {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/manualnatrules?bulk=true&filter=ids:"+url.QueryEscape(bulk), reqMods...)
		if err != nil {
			return fmt.Errorf("failed to delete object (DELETE), got error: %s, %s", err, res.String())
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Truncate finished successfully", state.Id.ValueString()))

		state.ManualNatRules = slices.Delete(state.ManualNatRules, kept, kept+counts[i])
	}

	return nil
}

func (r *FTDNATPolicyResource) truncateAutoNatRulesAt(ctx context.Context, state *FTDNATPolicy, kept int, reqMods ...func(*fmc.Req)) error {
	var b strings.Builder
	var bulks []string
	var counts []int
	count := 0

	for i := kept; i < len(state.AutoNatRules); i++ {
		b.WriteString(state.AutoNatRules[i].Id.ValueString() + ",")
		count++
		if b.Len() >= maxUrlParamLength {
			bulks = append(bulks, b.String())
			counts = append(counts, count)
			b.Reset()
			count = 0
		}
	}

	if b.Len() > 0 {
		bulks = append(bulks, b.String())
		counts = append(counts, count)
	}

	defer func() {
		time.Sleep(2 * time.Second)
	}()

	for i, bulk := range bulks {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/autonatrules?bulk=true&filter=ids:"+url.QueryEscape(bulk), reqMods...)
		if err != nil {
			return fmt.Errorf("failed to delete object (DELETE), got error: %s, %s", err, res.String())
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Truncate finished successfully", state.Id.ValueString()))

		state.AutoNatRules = slices.Delete(state.AutoNatRules, kept, kept+counts[i])
	}

	return nil
}

func (r *FTDNATPolicyResource) createManualNatRulesAt(ctx context.Context, plan FTDNATPolicy, body []gjson.Result, startIndex int, state *FTDNATPolicy, reqMods ...func(*fmc.Req)) error {
	for i := startIndex; i < len(body); i++ {
		bulk := `{"dummy_manual_nat_rules":[]}`
		j := i
		bulkCount := 0
		bodyLen := 0
		head := plan.ManualNatRules[i]

		for ; i < len(body); i++ {
			if !head.Section.Equal(plan.ManualNatRules[i].Section) {
				i--
				break
			}
			rule := body[i].String()
			rule, _ = sjson.Delete(rule, "metadata")

			// Check if the body is too big for a single POST
			bodyLen += len(rule)
			if bodyLen >= maxPayloadSize {
				i--
				break
			}

			bulk, _ = sjson.SetRaw(bulk, "dummy_manual_nat_rules.-1", rule)
			bulkCount++
			if bulkCount >= bulkSizeCreate {
				break
			}
		}

		tflog.Debug(ctx, fmt.Sprintf("Bulk: %s", bulk))

		param := "?bulk=true&section=" + strings.ToLower(head.Section.ValueString())
		res, err := r.client.Post(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+"/manualnatrules"+param, gjson.Parse(bulk).Get("dummy_manual_nat_rules").String(), reqMods...)
		if err != nil {
			return fmt.Errorf("failed to configure object (POST), got error: %s, %s", err, res.String())
		}

		for _, v := range res.Get("items").Array() {
			item := plan.ManualNatRules[j]
			item.Id = types.StringValue(v.Get("id").String())

			if len(state.ManualNatRules) <= j {
				state.ManualNatRules = append(state.ManualNatRules, item)
			} else {
				state.ManualNatRules[j] = item
			}

			j++
		}
	}

	return nil
}

func (r *FTDNATPolicyResource) createAutoNatRulesAt(ctx context.Context, plan FTDNATPolicy, body []gjson.Result, startIndex int, state *FTDNATPolicy, reqMods ...func(*fmc.Req)) error {
	for i := startIndex; i < len(body); i++ {
		bulk := `{"dummy_auto_nat_rules":[]}`
		j := i
		bulkCount := 0
		bodyLen := 0

		for ; i < len(body); i++ {
			rule := body[i].String()

			// Check if the body is too big for a single POST
			bodyLen += len(rule)
			if bodyLen >= maxPayloadSize {
				i--
				break
			}

			bulk, _ = sjson.SetRaw(bulk, "dummy_auto_nat_rules.-1", rule)

			bulkCount++
			if bulkCount >= bulkSizeCreate {
				break
			}

		}

		param := "?bulk=true"
		res, err := r.client.Post(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+"/autonatrules"+param, gjson.Parse(bulk).Get("dummy_auto_nat_rules").String(), reqMods...)
		if err != nil {
			return fmt.Errorf("failed to configure object (POST), got error: %s, %s", err, res.String())
		}

		for _, v := range res.Get("items").Array() {
			item := plan.AutoNatRules[j]
			item.Id = types.StringValue(v.Get("id").String())

			if len(state.AutoNatRules) <= j {
				state.AutoNatRules = append(state.AutoNatRules, item)
			} else {
				state.AutoNatRules[j] = item
			}

			j++
		}
	}

	return nil
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *FTDNATPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FTDNATPolicy

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

func (r *FTDNATPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
