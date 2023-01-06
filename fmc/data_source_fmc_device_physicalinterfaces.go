package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcPhysicalInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePhysicalInterfaceRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the device this resource needs",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the FTD accessPolicy",
			},
		},
	}
}

func dataSourcePhysicalInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	ifc, err := c.GetFmcPhysicalInterface(ctx, d.Get("id").(string), d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get physical interface",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(ifc.ID)

	if err := d.Set("name", ifc.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physical Interface",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
