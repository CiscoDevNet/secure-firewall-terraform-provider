package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var vtep_policies_type string = "VTEPPolicy"

func resourceFmcDeviceVTEPPolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Access Rules in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_VTEP_policy\" \"fmc_VTEP_policy_1\" {\n" +
			"    id = vtepPolicyUUID\n" +
			"    name = \"vtepPolic1\"\n" +
			"    type = \"VTEPPolicy\"\n" +
			"    nve_enable = false\n" +
			"    vtepEntries {\n" +
			"        sourceInterface {\n" +
			"            id = data.fmc_security_zones.inside.id\n" +
			"            type =  data.fmc_security_zones.inside.type\n" +
			"        }\n" +
			"        sourceInterface {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"\n" +
			"resource \"fmc_VTEP_policy\" \"fmc_VTEP_policy_2\" {\n" +
			"    id = vtepPolicyUUID1\n" +
			"    name = \"vtepPolicy2\"\n" +
			"    type = \"VTEPPolicy\"\n" +
			"    nveEnable = false\n" +
			"    vtepEntries {\n" +
			"        sourceInterface {\n" +
			"            id = data.fmc_security_zones.inside.id\n" +
			"            type =  data.fmc_security_zones.inside.type\n" +
			"        }\n" +
			"        sourceInterface {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"\n" +
			"}\n" +
			"```\n",
		CreateContext: resourceFmcDeviceVTEPPoliciesCreate,
		ReadContext:   resourceFmcDeviceVTEPPoliciessRead,
		UpdateContext: resourceFmcDeviceVTEPPoliciessUpdate,
		DeleteContext: resourceFmcDeviceVTEPPoliciesDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this policy",
			},
			// "name": {
			// 	Type:        schema.TypeString,
			// 	Required:    true,
			// 	Description: "The name of the policy",
			// },
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this physical interfaces",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this policy",
			},
			"nve_enable": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "The nveEnable the policy",
			},
			"vtep_entries": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_interface": {
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
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of this physical interfaces",
									},
								},
							},
						},
					},
				},
				Description: "Destination zones for this resource",
			},
		},
	}
}

func resourceFmcDeviceVTEPPoliciesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	var vtepEntry []VTEPEntry

	var sourceInterface SourceInterface

	if inputObjs, ok := d.GetOk("source_interface"); ok {
		obj := inputObjs.([]interface{})[0].(map[string]interface{})
		sourceInterface = SourceInterface{
			ID:   obj["id"].(string),
			Name: obj["name"].(string),
			Type: obj["type"].(string),
		}
	}

	if inputObjs, ok := d.GetOk("vtep_entries"); ok {
		obj := inputObjs.([]interface{})[0].(map[string]interface{})
		vtepEntry = append(vtepEntry, VTEPEntry{
			SourceInterface:          sourceInterface,
			NveVtepId:                obj["nveVtepId"].(int),
			NveDestinationPort:       obj["nveDestinationPort"].(int),
			NveEncapsulationType:     obj["NveEncapsulationType"].(string),
			// NveNeighborDiscoveryType: obj["NveNeighborDiscoveryType"].(string),
		})
	}

	res, err := c.CreateVTEPPolicies(ctx, d.Get("acp").(string), &VTEPPolicy{
		Type:        vtep_policies_type,
		NveEnable:   d.Get("nve_enable").(bool),
		VTEPEntries: vtepEntry,
	})
	if err != nil {
		return returnWithDiag(diags, err)
	}
	d.SetId(res.ID)
	return resourceFmcDeviceVTEPPoliciessRead(ctx, d, m)
}

func resourceFmcDeviceVTEPPoliciessRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//item, err := c.GetVTEPPoliciesByName(ctx, d.Get("acp").(string), d.Id())
	item, err := c.GetVTEPPolicies(ctx, d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("type", item.Type); err != nil {
		return returnWithDiag(diags, err)
	}

	return diags
}

func resourceFmcDeviceVTEPPoliciessUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("vtep_entries") {
		var vtepEntry []VTEPEntry

		var sourceInterface SourceInterface

		if inputObjs, ok := d.GetOk("source_interface"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			sourceInterface = SourceInterface{
				ID:   obj["id"].(string),
				Name: obj["name"].(string),
				Type: obj["type"].(string),
			}
		}

		if inputObjs, ok := d.GetOk("vtep_entries"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			vtepEntry = append(vtepEntry, VTEPEntry{
				SourceInterface:          sourceInterface,
				NveVtepId:                obj["nveVtepId"].(int),
				NveDestinationPort:       obj["nveDestinationPort"].(int),
				NveEncapsulationType:     obj["NveEncapsulationType"].(string),
				// NveNeighborDiscoveryType: obj["NveNeighborDiscoveryType"].(string),
			})
		}

		res, err := c.UpdateVTEPPolicies(ctx, d.Id(), &VTEPPolicy{
			Type:        vtep_policies_type,
			NveEnable:   d.Get("nve_enable").(bool),
			VTEPEntries: vtepEntry,
		})
		if err != nil {
			return returnWithDiag(diags, err)
		}
		d.SetId(res.ID)
	}
	return resourceFmcDeviceVTEPPoliciessRead(ctx, d, m)
}

func resourceFmcDeviceVTEPPoliciesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteVTEPPolicies(ctx, d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
