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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceVRFIPv6StaticRouteDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceVRFIPv6StaticRouteDataSource{}
)

func NewDeviceVRFIPv6StaticRouteDataSource() datasource.DataSource {
	return &DeviceVRFIPv6StaticRouteDataSource{}
}

type DeviceVRFIPv6StaticRouteDataSource struct {
	client *fmc.Client
}

func (d *DeviceVRFIPv6StaticRouteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_vrf_ipv6_static_route"
}

func (d *DeviceVRFIPv6StaticRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device VRF IPv6 Static Route.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent device .",
				Required:            true,
			},
			"vrf_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent VRF.",
				Required:            true,
			},
			"interface_logical_name": schema.StringAttribute{
				MarkdownDescription: "Logical name of the parent interface. For transparent mode, any bridge group member interface. For routed mode with bridge groups, any bridge group member interface for the BVI name.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'IPv6StaticRoute'.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Id of the interface provided in `interface_logical_name`. The value is ignored, but the attribute itself is useful for ensuring that Terraform creates interface resource before the static route resource (and destroys the interface resource only after the static route has been destroyed).",
				Computed:            true,
			},
			"destination_networks": schema.SetNestedAttribute{
				MarkdownDescription: "Set of the destination networks matching this route (Host, Networks or Ranges).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the object.",
							Computed:            true,
						},
					},
				},
			},
			"metric_value": schema.Int64Attribute{
				MarkdownDescription: "The cost of the route. The metric is used to compare routes among different routing protocols. The default administrative distance for static routes is 1, giving it precedence over routes discovered by dynamic routing protocols but not directly connected routes.",
				Computed:            true,
			},
			"gateway_host_object_id": schema.StringAttribute{
				MarkdownDescription: "Id of the next hop for this route. Exactly one of `gateway_host_object_id` or `gateway_host_literal` must be present.",
				Computed:            true,
			},
			"gateway_host_literal": schema.StringAttribute{
				MarkdownDescription: "Next hop for this route as a literal IPv6 address. Exactly one of `gateway_host_object_id` or `gateway_host_literal` must be present.",
				Computed:            true,
			},
			"is_tunneled": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether this route is a separate default route for VPN traffic. Should be used for default route only (such as when the destination_networks points to a builtin host 'any-ipv6'). Useful if you want VPN traffic to use a different default route than non-VPN traffic. When a tunnel terminates on the device, all traffic from it that cannot be routed using learned or static routes is sent to this route. You can configure only one default tunneled gateway per device. ECMP for tunneled traffic is not supported. This attribute conflicts with `metric_value` attribute.",
				Computed:            true,
			},
		},
	}
}

func (d *DeviceVRFIPv6StaticRouteDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceVRFIPv6StaticRouteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceVRFIPv6StaticRoute

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
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
