package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcVTEPPoliciess() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for VTEP Policies in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_device_VTEP_Policies\" \"acp\" {\n" +
			"	name = \"FTD ACP\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcVTEPPoliciesRead,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of this VTEP Policies",
			},
			// "name": {
			// 	Type:        schema.TypeString,
			// 	Required:    true,
			// 	Description: "Name of the VTEP Policies",
			// },
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this VTEP Policies",
			},
		},
	}
}

func dataSourceFmcVTEPPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	resp, err := c.GetVTEPPolicies(ctx, d.Get("device_id").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get VTEP Policies",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(resp.ID)

	// if err := d.Set("name", resp.Name); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "unable to read VTEP Policies",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	if err := d.Set("type", resp.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read VTEP Policies",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
