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

// Section below is generated&owned by "gen/generator.go". //template:begin provider
import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// FmcProvider defines the provider implementation.
type FmcProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// FmcProviderModel describes the provider data model.
type FmcProviderModel struct {
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
	URL        types.String `tfsdk:"url"`
	Insecure   types.Bool   `tfsdk:"insecure"`
	ReqTimeout types.String `tfsdk:"req_timeout"`
	Retries    types.Int64  `tfsdk:"retries"`
}

// FmcProviderData describes the data maintained by the provider.
type FmcProviderData struct {
	Client *fmc.Client
}

// Define provider constants
const (
	// Define maximum elements in single bulk request for delete & create
	bulkSizeDelete int = 200
	bulkSizeCreate int = 1000
	// maximum payload size in bytes
	maxPayloadSize int = 2048000
	// maximum URL Param length. This is a rough estimate and does not account for the entire URL length.
	maxUrlParamLength int = 7000
)

// Metadata returns the provider type name.
func (p *FmcProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "fmc"
	resp.Version = p.version
}

func (p *FmcProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the FMC instance. This can also be set as the FMC_USERNAME environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the FMC instance. This can also be set as the FMC_PASSWORD environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL of the Cisco FMC instance. This can also be set as the FMC_URL environment variable.",
				Optional:            true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the FMC_INSECURE environment variable. Defaults to `true`.",
				Optional:            true,
			},
			"req_timeout": schema.StringAttribute{
				MarkdownDescription: "Timeout for a single HTTPS request made to REST API before it is retried. This can also be set as the FMC_REQTIMEOUT environment variable. A string like `\"1s\"` means one second. Defaults to `\"5s\"`.",
				Optional:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the FMC_RETRIES environment variable. Defaults to `3`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9),
				},
			},
		},
	}
}

func (p *FmcProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config FmcProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.IsNull() {
		username = os.Getenv("FMC_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.IsNull() {
		password = os.Getenv("FMC_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a URL to the provider
	var url string
	if config.URL.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.IsNull() {
		url = os.Getenv("FMC_URL")
	} else {
		url = config.URL.ValueString()
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	var insecure bool
	if config.Insecure.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.IsNull() {
		insecureStr := os.Getenv("FMC_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.ValueBool()
	}

	var reqTimeout time.Duration
	if config.ReqTimeout.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as req_timeout",
		)
		return
	}

	var reqTimeoutStr string
	if config.ReqTimeout.IsNull() {
		reqTimeoutStr = os.Getenv("FMC_REQTIMEOUT")
		if reqTimeoutStr == "" {
			reqTimeoutStr = "5s"
		}
	} else {
		reqTimeoutStr = config.ReqTimeout.ValueString()
	}
	reqTimeout, err := time.ParseDuration(reqTimeoutStr)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			fmt.Sprintf("Cannot parse the req_timeout string: %v", err),
		)
		return
	}

	var retries int64
	if config.Retries.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.IsNull() {
		retriesStr := os.Getenv("FMC_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.ValueInt64()
	}

	tflog.Debug(ctx, fmt.Sprint("Creating a new FMC client",
		"  url=", url,
		"  insecure=", insecure,
		"  req_timeout=", reqTimeout,
		"  retries=", retries,
	))

	// Create a new FMC client and set it to the provider client
	c, err := fmc.NewClient(url, username, password, fmc.Insecure(insecure), fmc.MaxRetries(int(retries)), fmc.RequestTimeout(reqTimeout))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create fmc client:\n\n"+err.Error(),
		)
		return
	}

	data := FmcProviderData{Client: &c}
	resp.DataSourceData = &data
	resp.ResourceData = &data
}

func (p *FmcProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAccessControlPolicyResource,
		NewApplicationFilterResource,
		NewBFDTemplateResource,
		NewDeviceResource,
		NewDeviceBFDResource,
		NewDeviceBGPResource,
		NewDeviceBGPGeneralSettingsResource,
		NewDeviceClusterResource,
		NewDeviceDeployResource,
		NewDeviceEtherChannelInterfaceResource,
		NewDeviceHAPairResource,
		NewDeviceHAPairMonitoringResource,
		NewDeviceIPv4StaticRouteResource,
		NewDeviceIPv6StaticRouteResource,
		NewDevicePhysicalInterfaceResource,
		NewDeviceSubinterfaceResource,
		NewDeviceVNIInterfaceResource,
		NewDeviceVRFResource,
		NewDeviceVRFIPv4StaticRouteResource,
		NewDeviceVRFIPv6StaticRouteResource,
		NewDeviceVTEPPolicyResource,
		NewDynamicObjectsResource,
		NewExtendedACLResource,
		NewFilePolicyResource,
		NewFQDNObjectResource,
		NewFQDNObjectsResource,
		NewFTDNATPolicyResource,
		NewHostResource,
		NewHostsResource,
		NewICMPv4ObjectResource,
		NewICMPv4ObjectsResource,
		NewICMPv6ObjectResource,
		NewInterfaceGroupResource,
		NewIntrusionPolicyResource,
		NewNetworkResource,
		NewNetworkAnalysisPolicyResource,
		NewNetworkGroupResource,
		NewNetworkGroupsResource,
		NewNetworksResource,
		NewPolicyAssignmentResource,
		NewPortResource,
		NewPortGroupResource,
		NewPortGroupsResource,
		NewPortsResource,
		NewPrefilterPolicyResource,
		NewRangeResource,
		NewRangesResource,
		NewSecurityZoneResource,
		NewSecurityZonesResource,
		NewSGTResource,
		NewSGTsResource,
		NewSmartLicenseResource,
		NewStandardACLResource,
		NewTimeRangeResource,
		NewTimeRangesResource,
		NewTunnelZoneResource,
		NewTunnelZonesResource,
		NewURLResource,
		NewURLGroupResource,
		NewURLGroupsResource,
		NewURLsResource,
		NewVLANTagResource,
		NewVLANTagGroupResource,
		NewVLANTagGroupsResource,
		NewVLANTagsResource,
	}
}

func (p *FmcProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAccessControlPolicyDataSource,
		NewApplicationDataSource,
		NewApplicationBusinessRelevanceDataSource,
		NewApplicationBusinessRelevancesDataSource,
		NewApplicationCategoriesDataSource,
		NewApplicationCategoryDataSource,
		NewApplicationFilterDataSource,
		NewApplicationRiskDataSource,
		NewApplicationRisksDataSource,
		NewApplicationTagDataSource,
		NewApplicationTagsDataSource,
		NewApplicationTypeDataSource,
		NewApplicationTypesDataSource,
		NewApplicationsDataSource,
		NewBFDTemplateDataSource,
		NewDeviceDataSource,
		NewDeviceBFDDataSource,
		NewDeviceBGPDataSource,
		NewDeviceBGPGeneralSettingsDataSource,
		NewDeviceClusterDataSource,
		NewDeviceEtherChannelInterfaceDataSource,
		NewDeviceHAPairDataSource,
		NewDeviceHAPairMonitoringDataSource,
		NewDeviceIPv4StaticRouteDataSource,
		NewDeviceIPv6StaticRouteDataSource,
		NewDevicePhysicalInterfaceDataSource,
		NewDeviceSubinterfaceDataSource,
		NewDeviceVNIInterfaceDataSource,
		NewDeviceVRFDataSource,
		NewDeviceVRFIPv4StaticRouteDataSource,
		NewDeviceVRFIPv6StaticRouteDataSource,
		NewDeviceVTEPPolicyDataSource,
		NewDynamicObjectsDataSource,
		NewExtendedACLDataSource,
		NewFileCategoriesDataSource,
		NewFileCategoryDataSource,
		NewFilePolicyDataSource,
		NewFileTypeDataSource,
		NewFileTypesDataSource,
		NewFQDNObjectDataSource,
		NewFQDNObjectsDataSource,
		NewFTDNATPolicyDataSource,
		NewHostDataSource,
		NewHostsDataSource,
		NewICMPv4ObjectDataSource,
		NewICMPv4ObjectsDataSource,
		NewICMPv6ObjectDataSource,
		NewInterfaceGroupDataSource,
		NewIntrusionPolicyDataSource,
		NewNetworkDataSource,
		NewNetworkAnalysisPolicyDataSource,
		NewNetworkGroupDataSource,
		NewNetworksDataSource,
		NewPolicyAssignmentDataSource,
		NewPortDataSource,
		NewPortGroupDataSource,
		NewPortGroupsDataSource,
		NewPortsDataSource,
		NewPrefilterPolicyDataSource,
		NewRangeDataSource,
		NewRangesDataSource,
		NewSecurityZoneDataSource,
		NewSecurityZonesDataSource,
		NewSGTDataSource,
		NewSGTsDataSource,
		NewSNMPAlertDataSource,
		NewSNMPAlertsDataSource,
		NewStandardACLDataSource,
		NewSyslogAlertDataSource,
		NewSyslogAlertsDataSource,
		NewTimeRangeDataSource,
		NewTimeRangesDataSource,
		NewTunnelZoneDataSource,
		NewTunnelZonesDataSource,
		NewURLDataSource,
		NewURLGroupDataSource,
		NewURLGroupsDataSource,
		NewURLsDataSource,
		NewVariableSetDataSource,
		NewVLANTagDataSource,
		NewVLANTagGroupDataSource,
		NewVLANTagGroupsDataSource,
		NewVLANTagsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &FmcProvider{
			version: version,
		}
	}
}

// End of section. //template:end provider
