package fmc

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("fmc_username").(string)
	password := d.Get("fmc_password").(string)
	iscdfmc := d.Get("is_cdfmc").(bool)
	cdotoken := d.Get("cdo_token").(string)
	cdfmcdomainuuid := d.Get("cdfmc_domain_uuid").(string)
	host := d.Get("fmc_host").(string)
	insecureSkipVerify := d.Get("fmc_insecure_skip_verify").(bool)
	var diags diag.Diagnostics

	if !iscdfmc && username != "" && password != "" && host != "" && cdotoken == "" && cdfmcdomainuuid == "" {
		client := NewClient(username, password, host, insecureSkipVerify)
		err := client.Login()
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return client, diags
	}
	if iscdfmc && cdotoken != "" && cdfmcdomainuuid != "" && host != "" && username == "" && password == "" {
		client := CDFMC_NewClient(cdotoken, cdfmcdomainuuid, host, insecureSkipVerify)
		err := client.Login()
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return client, diags
	}
	return nil, diag.FromErr(errors.New("missing fmc username, password, cdo token , cdfmc domain_uuid or base url"))
}

// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"fmc_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_USERNAME", ""),
				Description: "Username for the user to login to FMC",
			},
			"fmc_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_PASSWORD", ""),
				Description: "Password for the user to login to FMC",
			},
			"fmc_host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_HOST", nil),
				Description: "Hostname/IP address of the FMC",
			},
			"is_cdfmc": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("IS_CDFMC", false),
				Description: "set to true if is a cloud-delivered Firepower Management Center (cdFMC) with Cisco Defence Orchestrator(CDO)",
			},
			"cdo_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CDO_TOKEN", ""),
				Description: "The CDO Api token to manage cdFMC",
			},
			"cdfmc_domain_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CDFMC_DOMAIN_UUID", ""),
				Description: "The domain uuid to use API of cdFMC the default domain for global is e276abec-e0f2-11e3-8169-6d9ed49b625f",
			},
			"fmc_insecure_skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_INSECURE_SKIP_VERIFY", false),
				Description: "Skip certificate checks if the certificate is not public CA signed, or if using IP address",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fmc_url_objects":                resourceFmcURLObjects(),
			"fmc_url_object_group":           resourceFmcURLObjectGroup(),
			"fmc_port_objects":               resourceFmcPortObjects(),
			"fmc_network_objects":            resourceFmcNetworkObjects(),
			"fmc_host_objects":               resourceFmcHostObjects(),
			"fmc_range_objects":              resourceFmcRangeObjects(),
			"fmc_fqdn_objects":               resourceFmcFQDNObjects(),
			"fmc_icmpv4_objects":             resourceFmcICMPV4Objects(),
			"fmc_access_rules":               resourceFmcAccessRules(),
			"fmc_access_policies":            resourceFmcAccessPolicies(),
			"fmc_network_group_objects":      resourceFmcNetworkGroupObjects(),
			"fmc_port_group_objects":         resourceFmcPortGroupObjects(),
			"fmc_ftd_nat_policies":           resourceFmcNatPolicies(),
			"fmc_ftd_autonat_rules":          resourceFmcAutoNatRules(),
			"fmc_ftd_manualnat_rules":        resourceFmcManualNatRules(),
			"fmc_policy_devices_assignments": resourceFmcPolicyDevicesAssignments(),
			"fmc_ftd_deploy":                 resourceFmcFtdDeploy(),
			"fmc_dynamic_objects":            resourceFmcDynamicObjects(),
			"fmc_dynamic_object_mapping":     resourceFmcDynamicObjectMapping(),
			"fmc_security_zone":              resourceFmcSecurityZone(),
			"fmc_time_range_object":          resourceFmcTimeRangeObject(),
			"fmc_access_policies_category":   resourceFmcAccessPoliciesCategory(),
			"fmc_prefilter_policy":           resourceFmcPrefilterPolicy(),
			"fmc_ips_policies":               resourceFmcIPSPolicies(),
			"fmc_device_physical_interfaces": resourcePhyInterface(),
			"fmc_device_vni":                 resourceVNI(),
			"fmc_devices":                    resourceFmcDevices(),
			"fmc_device_vtep":                resourceVTEP(),
			"fmc_staticIPv4_route":           resourceFmcStaticIPv4Route(),
			"fmc_extended_acl":               resourceFmcExtendedAcl(),
			"fmc_sgt_objects":                resourceFmcSGTObjects(),
			"fmc_standard_acl":               resourceFmcStandardAcl(),
			"fmc_network_analysis_policy":    resourceFmcNetworkAnalysisPolicy(),
			"fmc_smart_license":              resourceFmcSmartLicense(),
			"fmc_device_subinterfaces":       resourceFmcSubInterface(),
			"fmc_network_objects_bulk":       resourceFmcNetworkObjectsBulk(),
			"fmc_devices_bulk":               resourceFmcDevicesBulk(),
			"fmc_device_cluster":             resourceFmcDeviceCluster(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"fmc_devices":                    dataSourceFmcDevices(),
			"fmc_access_policies":            dataSourceFmcAccessPolicies(),
			"fmc_ips_policies":               dataSourceFmcIPSPolicies(),
			"fmc_file_policies":              dataSourceFmcFilePolicies(),
			"fmc_syslog_alerts":              dataSourceFmcSyslogAlerts(),
			"fmc_security_zones":             dataSourceFmcSecurityZones(),
			"fmc_network_objects":            dataSourceFmcNetworkObjects(),
			"fmc_host_objects":               dataSourceFmcHostObjects(),
			"fmc_url_objects":                dataSourceFmcURLObjects(),
			"fmc_port_objects":               dataSourceFmcPortObjects(),
			"fmc_dynamic_objects":            dataSourceFmcDynamicObjects(),
			"fmc_network_group_objects":      dataSourceFmcNetworkGroupObjects(),
			"fmc_extended_acl":               dataSourceFmcExtendedAcl(),
			"fmc_device_vni":                 dataSourceFmcVNI(),
			"fmc_ftd_nat_policies":           dataSourceFmcNatPolicies(),
			"fmc_staticIPv4_route":           dataSourceFmcStaticIPv4Route(),
			"fmc_device_physical_interfaces": dataSourceFmcPhysicalInterface(),
			"fmc_standard_acl":               dataSourceFmcStandardAcl(),
			"fmc_sgt_objects":                dataFmcSGTObjects(),
			"fmc_ise_sgt_objects":            dataFmcIseSGTObjects(),
			"fmc_network_analysis_policy":    dataSourceFmcNetworkAnalysisPolicy(),
			"fmc_device_vtep_policies":       dataSourceFmcVTEPPolicies(),
			"fmc_smart_license":              dataSourceFmcSmartLicense(),
			"fmc_device_subinterfaces":       dataSourceFmcSubInterfaces(),
			"fmc_port_group_objects":         dataSourceFmcPortGroupObjects(),
			"fmc_device_cluster":             dataSourceFmcDeviceCluster(),
			"fmc_ravpn_policy":               dataSourceFmcRavpn(),
			"fmc_anyconnect_package":         dataSourceFmcAnyconnectPackageObject(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
