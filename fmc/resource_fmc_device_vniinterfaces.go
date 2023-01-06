package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var VNI_interface_type string = "VNIInterface"

func resourceFmcDeviceVNIInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for VNI Interfaces in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_VNI_interfaces\" \"VNI_interfaces_1\" {\n" +
			"    id = fpVNIinterfaceUUID1\n" +
			"    name = \"s1p1\"\n" +
			"    type = \"FPVNIInterface\"\n" +
			"    enabled = false\n" +
			"    MTU = 1500\n" +
			"    mode = true\n" +
			"    ifname = false\n" +
			"    nveOnly = true\n" +
			"    managementOnly = true\n" +
			"    source_zones {\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.inside.id\n" +
			"            type =  data.fmc_security_zones.inside.type\n" +
			"        }\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"\n" +
			"resource \"fmc_VNI_interfaces\" \"VNI_interfaces_2\" {\n" +
			"    id = fpVNIinterfaceUUID2\n" +
			"    name = \"s1p2\"\n" +
			"    type = \"FPVNIInterface\"\n" +
			"    enabled = false\n" +
			"    MTU = 1500\n" +
			"    mode = true\n" +
			"    ifname = false\n" +
			"    nveOnly = true\n" +
			"    managementOnly = true\n" +
			"    source_zones {\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.inside.id\n" +
			"            type =  data.fmc_security_zones.inside.type\n" +
			"        }\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"\n" +
			"}\n" +
			"```\n",
		UpdateContext: resourceFmcDeviceVNIInterfacesUpdate,
		ReadContext:   resourceFmcDeviceVNIInterfacesRead,
		DeleteContext: resourceFmcDeviceVNIInterfacesDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this VNI interfaces",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of VNI interfaces",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of this VNI interfaces",
			},
			"vniId": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The description of this VNI interfaces",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this VNI interfaces",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enabled of this VNI interfaces",
			},
			"MTU": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "MTU of this VNI interfaces",
			},
			"mode": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "mode of this VNI interfaces",
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ifname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ifname of this VNI interfaces",
			},
			"enableSGTPropagate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enableSGTPropagate of this VNI interfaces",
			},
			"managementOnly": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "managementOnly of this VNI interfaces",
			},
			"enableAntiSpoofing": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enableAntiSpoofing of this VNI interfaces",
			},
			"activeMACAddress": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "activeMACAddress of this VNI interfaces",
			},
			"fragmentReassembly": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "activeMACAddress of this VNI interfaces",
			},
			"standbyMACAddress": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "standbyMACAddress of this VNI interfaces",
			},
			"enableDNSLookup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enableDNSLookup of this VNI interfaces",
			},
			"enableProxy": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "enableProxy of this VNI interfaces",
			},
			"vtepID": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "vtepID of this VNI interfaces",
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "vtepID of this VNI interfaces",
			},
			"fmc_access_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_networks": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The ID of this resource",
									},
									"type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The type of this resource",
									},
									"description": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The description of this resource",
									},
								},
							},
						},
						"enableAccess": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "enableAccess of this resource",
						},
					},
				},
				Description: "fmcAccessConfig for this resource",
			},
			"security_zone": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of this resource",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of this resource",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of this resource",
						},
					},
				},
				Description: "Destination zones for this resource",
			},
		},
	}
}

func resourceFmcDeviceVNIInterfacesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	var oerrideDefaultFragmentSetting OverrideDefaultFragmentSetting
	var arpConfig []ArpConfig

	if inputObjs, ok := d.GetOk("oerrideDefaultFragmentSetting"); ok {
		obj := inputObjs.([]interface{})[0].(map[string]interface{})
		oerrideDefaultFragmentSetting = OverrideDefaultFragmentSetting{
			Size:    obj["size"].(int),
			Timeout: obj["timeout"].(int),
			Chain:   obj["chain"].(int),
		}
	}

	if inputObjs, ok := d.GetOk("arpConfig"); ok {
		obj := inputObjs.([]interface{})[0].(map[string]interface{})
		arpConfig = append(arpConfig, ArpConfig{
			IpAddress:   obj["ipAddress"].(string),
			MacAddress:  obj["macAddress"].(string),
			EnableAlias: obj["enableAlias"].(bool),
		})
	}
	res, err := c.CreateVNIInterfaces(ctx, d.Get("id").(string), &FTDVNIInterfaces{
		Type:                           VNI_interface_type,
		ArpConfig:                      arpConfig,
		OverrideDefaultFragmentSetting: oerrideDefaultFragmentSetting,
		Name:                           d.Get("name").(string),
		VNIID:                          d.Get("vniId").(int),
		MulticastGroupAddress:          d.Get("multicastGroupAddress").(string),
		VTEPID:                         d.Get("vtepID").(int),
		Enabled:                        d.Get("enables").(bool),
	})
	if err != nil {
		return returnWithDiag(diags, err)
	}
	d.SetId(res.ID)
	return resourceFmcDeviceVNIInterfacesRead(ctx, d, m)
}

func resourceFmcDeviceVNIInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetVNIInterfaces(ctx, d.Get("name").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("type", item.Type); err != nil {
		return returnWithDiag(diags, err)
	}

	return diags
}

func resourceFmcDeviceVNIInterfacesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("oerrideDefaultFragmentSetting", "arpConfig", "name", "vniId", "multicastGroupAddress", "vtepID", "enabled") {
		var oerrideDefaultFragmentSetting OverrideDefaultFragmentSetting
		var arpConfig []ArpConfig

		if inputObjs, ok := d.GetOk("oerrideDefaultFragmentSetting"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			oerrideDefaultFragmentSetting = OverrideDefaultFragmentSetting{
				Size:    obj["size"].(int),
				Timeout: obj["timeout"].(int),
				Chain:   obj["chain"].(int),
			}
		}

		if inputObjs, ok := d.GetOk("arpConfig"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			arpConfig = append(arpConfig, ArpConfig{
				IpAddress:   obj["ipAddress"].(string),
				MacAddress:  obj["macAddress"].(string),
				EnableAlias: obj["enableAlias"].(bool),
			})
		}
		res, err := c.UpdateVNIInterfaces(ctx, d.Get("id").(string), &FTDVNIInterfaces{
			Type:                           VNI_interface_type,
			ArpConfig:                      arpConfig,
			OverrideDefaultFragmentSetting: oerrideDefaultFragmentSetting,
			Name:                           d.Get("name").(string),
			VNIID:                          d.Get("vniId").(int),
			MulticastGroupAddress:          d.Get("multicastGroupAddress").(string),
			VTEPID:                         d.Get("vtepID").(int),
			Enabled:                        d.Get("enabled").(bool),
		})
		if err != nil {
			return returnWithDiag(diags, err)
		}
		d.SetId(res.ID)

	}
	return resourceFmcDeviceVNIInterfacesRead(ctx, d, m)
}

func resourceFmcDeviceVNIInterfacesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteVNIInterfaces(ctx, d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}