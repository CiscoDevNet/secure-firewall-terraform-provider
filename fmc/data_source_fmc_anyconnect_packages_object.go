package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcAnyconnectPackageObject() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Network Objects in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_anyconnect_package\" \"PrivateVLAN\" {\n" +
			"	name = \"anyconnect-pkg\"\n" +
			"}\n" +
			"```\n" +
			"Any one of the id, name or value can be specified. The first filter in the order of id, name and value will be used, and the rest will be ignored if multiple are specified.",
		ReadContext: dataSourceFmcAnyconnectPackageObjectRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The type of this resource",
			},
		},
	}
}

func dataSourceFmcAnyconnectPackageObjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	idInput, okId := d.GetOk("id")
	nameInput, okName := d.GetOk("name")

	var (
		item *AnyconnectObjResponse
		err  error
	)
	if okId && (okName) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "More than one filter provided",
			Detail:   "The first filter in the order of id, name and port will be used, and the rest will be ignored",
		})
	}

	switch {
	case okId:
		item, err = c.GetFmcAnyconnectPackageObject(ctx, idInput.(string))
	case okName:
		item, err = c.GetFmcAnyconnectPackageByNameOrValue(ctx, nameInput.(string))
	default:
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No id, name provided, please provide any one",
			Detail:   "Please set one of the values to filter the datasource by",
		})
		return diags
	}

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get anyconnect object",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(item.ID)

	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read anyconnect object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read anyconnect object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
