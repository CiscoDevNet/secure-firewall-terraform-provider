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
	host := d.Get("fmc_host").(string)
	insecureSkipVerify := d.Get("fmc_insecure_skip_verify").(bool)
	var diags diag.Diagnostics

	if username != "" && password != "" && host != "" {
		client := NewClient(username, password, host, insecureSkipVerify)
		err := client.Login()
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return client, diags
	}
	return nil, diag.FromErr(errors.New("missing fmc username, password or base url"))
}

// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"fmc_username": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_USERNAME", nil),
				Description: "Username for the user to login to FMC",
			},
			"fmc_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_PASSWORD", nil),
				Description: "Password for the user to login to FMC",
			},
			"fmc_host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_HOST", nil),
				Description: "Hostname/IP address of the FMC",
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
			"fmc_dynamic_object":             resourceFmcDynamicObjects(),
			"fmc_dynamic_object_mapping":     resourceFmcDynamicObjectMapping(),
			"fmc_security_zone":              resourceFmcSecurityZone(),
			"fmc_time_range_object":          resourceFmcTimeRangeObject(),
			"fmc_access_policies_category":   resourceFmcAccessPoliciesCategory(),
			"fmc_prefilter_policy":           resourceFmcPrefilterPolicy(),
			// "fmc_device_vni_interface":       resourceFmcDeviceVNIInterfaces(),
			// "fmc_device_physical_interfaces": resourceFmcDevicePhysicalInterfaces(),
			"fmc_device_vtep_policies": resourceFmcDeviceVTEPPolicies(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fmc_devices":         dataSourceFmcDevices(),
			"fmc_access_policies": dataSourceFmcAccessPolicies(),
			"fmc_ips_policies":    dataSourceFmcIPSPolicies(),
			"fmc_file_policies":   dataSourceFmcFilePolicies(),
			"fmc_syslog_alerts":   dataSourceFmcSyslogAlerts(),
			"fmc_security_zones":  dataSourceFmcSecurityZones(),
			"fmc_network_objects": dataSourceFmcNetworkObjects(),
			"fmc_host_objects":    dataSourceFmcHostObjects(),
			"fmc_url_objects":     dataSourceFmcURLObjects(),
			"fmc_port_objects":    dataSourceFmcPortObjects(),
			"fmc_dynamic_objects": dataSourceFmcDynamicObjects(),
			"fmc_device_vni_interface":  dataSourceFmcVNIInterfaces(),
			"fmc_device_physical_interfaces":dataSourceFmcPhysicalInterface(),
			"fmc_device_vtep_policies":dataSourceFmcVTEPPoliciess(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
