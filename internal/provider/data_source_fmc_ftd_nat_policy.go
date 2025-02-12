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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FTDNATPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDNATPolicyDataSource{}
)

func NewFTDNATPolicyDataSource() datasource.DataSource {
	return &FTDNATPolicyDataSource{}
}

type FTDNATPolicyDataSource struct {
	client *fmc.Client
}

func (d *FTDNATPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_nat_policy"
}

func (d *FTDNATPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD NAT Policy.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the FTD Network Address Translation (NAT) policy.",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the object.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'FTDNatPolicy'.",
				Computed:            true,
			},
			"manual_nat_rules": schema.ListNestedAttribute{
				MarkdownDescription: "The ordered list of manual NAT rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the manual nat rule.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of Manual NAT rule.",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Indicates if the rule is enabled.",
							Computed:            true,
						},
						"section": schema.StringAttribute{
							MarkdownDescription: "Name of section to which the rule belongs.",
							Computed:            true,
						},
						"nat_type": schema.StringAttribute{
							MarkdownDescription: "Type of the rule",
							Computed:            true,
						},
						"fall_through": schema.BoolAttribute{
							MarkdownDescription: "Fallthrough to Interface PAT (Destination Interface)",
							Computed:            true,
						},
						"interface_in_original_destination": schema.BoolAttribute{
							MarkdownDescription: "Use interface address as original destination",
							Computed:            true,
						},
						"interface_in_translated_source": schema.BoolAttribute{
							MarkdownDescription: "Translate source network to destination interface address",
							Computed:            true,
						},
						"ipv6": schema.BoolAttribute{
							MarkdownDescription: "Use the IPv6 address of the destination interface for interface PAT.",
							Computed:            true,
						},
						"net_to_net": schema.BoolAttribute{
							MarkdownDescription: "Net to Net Mapping",
							Computed:            true,
						},
						"no_proxy_arp": schema.BoolAttribute{
							MarkdownDescription: "Do not proxy ARP on Destination Interface",
							Computed:            true,
						},
						"unidirectional": schema.BoolAttribute{
							MarkdownDescription: "Whether the rule is unidirectional",
							Computed:            true,
						},
						"source_interface_id": schema.StringAttribute{
							MarkdownDescription: "ID of source security zone or interface group",
							Computed:            true,
						},
						"original_source_id": schema.StringAttribute{
							MarkdownDescription: "ID of original source network object (host, network or range)",
							Computed:            true,
						},
						"original_source_port_id": schema.StringAttribute{
							MarkdownDescription: "ID of original source port object",
							Computed:            true,
						},
						"original_destination_id": schema.StringAttribute{
							MarkdownDescription: "ID of original destination network object (host, network or range)",
							Computed:            true,
						},
						"original_destination_port_id": schema.StringAttribute{
							MarkdownDescription: "ID of original destination port object",
							Computed:            true,
						},
						"route_lookup": schema.BoolAttribute{
							MarkdownDescription: "Perform Route Lookup for Destination Interface",
							Computed:            true,
						},
						"destination_interface_id": schema.StringAttribute{
							MarkdownDescription: "ID of destination security zone or interface group",
							Computed:            true,
						},
						"translated_source_id": schema.StringAttribute{
							MarkdownDescription: "ID of translated source network object (host, network or range)",
							Computed:            true,
						},
						"translated_source_port_id": schema.StringAttribute{
							MarkdownDescription: "ID of translated source port object",
							Computed:            true,
						},
						"translate_dns": schema.BoolAttribute{
							MarkdownDescription: "Translate DNS replies that match this rule",
							Computed:            true,
						},
						"translated_destination_id": schema.StringAttribute{
							MarkdownDescription: "ID of translated destination network object (host, network or range)",
							Computed:            true,
						},
						"translated_destination_port_id": schema.StringAttribute{
							MarkdownDescription: "ID of translated destination port object",
							Computed:            true,
						},
					},
				},
			},
			"auto_nat_rules": schema.ListNestedAttribute{
				MarkdownDescription: "The list of auto NAT rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Auto NAT rule.",
							Computed:            true,
						},
						"nat_type": schema.StringAttribute{
							MarkdownDescription: "Type of the rule",
							Computed:            true,
						},
						"destination_interface_id": schema.StringAttribute{
							MarkdownDescription: "ID of destination security zone or interface group",
							Computed:            true,
						},
						"fall_through": schema.BoolAttribute{
							MarkdownDescription: "Fallthrough to Interface PAT (Destination Interface)",
							Computed:            true,
						},
						"ipv6": schema.BoolAttribute{
							MarkdownDescription: "Use the IPv6 address of the destination interface for interface PAT.",
							Computed:            true,
						},
						"net_to_net": schema.BoolAttribute{
							MarkdownDescription: "Net to Net Mapping",
							Computed:            true,
						},
						"no_proxy_arp": schema.BoolAttribute{
							MarkdownDescription: "Do not proxy ARP on Destination Interface",
							Computed:            true,
						},
						"original_network_id": schema.StringAttribute{
							MarkdownDescription: "ID of original network object (host, network or range)",
							Computed:            true,
						},
						"original_port": schema.Int64Attribute{
							MarkdownDescription: "Original port number",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Service protocol",
							Computed:            true,
						},
						"perform_route_lookup": schema.BoolAttribute{
							MarkdownDescription: "Perform Route Lookup for Destination Interface",
							Computed:            true,
						},
						"source_interface_id": schema.StringAttribute{
							MarkdownDescription: "ID of source security zone or interface group",
							Computed:            true,
						},
						"translate_dns": schema.BoolAttribute{
							MarkdownDescription: "Translate DNS replies that match this rule",
							Computed:            true,
						},
						"translated_network_id": schema.StringAttribute{
							MarkdownDescription: "ID of translated network object (host, network or range)",
							Computed:            true,
						},
						"translated_network_is_destination_interface": schema.BoolAttribute{
							MarkdownDescription: "Translate source network to destination interface address",
							Computed:            true,
						},
						"translated_port": schema.Int64Attribute{
							MarkdownDescription: "Translated port number",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *FTDNATPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *FTDNATPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *FTDNATPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FTDNATPolicy

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(config.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	resManualNatRules, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString())+"/manualnatrules?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	resAutoNatRules, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString())+"/autonatrules?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	replaceManualNatRules := resManualNatRules.Get("items").String()
	if replaceManualNatRules == "" {
		replaceManualNatRules = "[]"
	}

	replaceAutoNatRules := resAutoNatRules.Get("items").String()
	if replaceAutoNatRules == "" {
		replaceAutoNatRules = "[]"
	}

	replace, _ := sjson.SetRaw(res.String(), "dummy_manual_nat_rules", replaceManualNatRules)
	replace, _ = sjson.SetRaw(replace, "dummy_auto_nat_rules", replaceAutoNatRules)
	res = gjson.Parse(replace)

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
