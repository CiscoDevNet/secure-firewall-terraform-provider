package fmc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcSubInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSubInterfaceRead,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the VNI",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the VNI",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the VNI",
			},
			"if_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ifname of the VNI",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "type of the VNI",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "enabled of the sub interface",
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Priority of the VNI",
			},

			"sub_interface_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "SubInterfaceId of the VNI",
			},
			"vnlanid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "vnlanid of the VNI",
			},

			"security_zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "securityZone of the VNI",
			},
		},
	}
}

func dataSourceSubInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("SUB Interface Datasource")

	c := m.(*Client)

	var diags diag.Diagnostics
	sinterface, err := c.GetFmcSubInterface(ctx, d.Get("device_id").(string), d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get VNI",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("ID=%s Name=%s", sinterface.ID, sinterface.Name)

	d.SetId(sinterface.ID)
	d.Set("name", sinterface.Name)

	if err := d.Set("name", sinterface.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read sub interface",
			Detail:   err.Error(),
		})
		return diags
	}

	subinterfacedetail, err := c.GetFmcSubInterfaceDetails(ctx, d.Get("device_id").(string), sinterface.ID)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get VNI",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("VNI Details ID=%s Name=%s Description=%s", subinterfacedetail.ID, subinterfacedetail.Name, subinterfacedetail.Description)

	d.SetId(subinterfacedetail.ID)
	d.Set("device_id", d.Get("device_id").(string))
	d.Set("name", subinterfacedetail.Name)
	d.Set("description", subinterfacedetail.Description)
	d.Set("segment_id", subinterfacedetail.VLANId)
	d.Set("vnlanid", subinterfacedetail.VLANId)
	d.Set("enabled", subinterfacedetail.Enabled)
	d.Set("priority", subinterfacedetail.Priority)
	d.Set("if_name", subinterfacedetail.Ifname)
	d.Set("security_zone_id", subinterfacedetail.SecurityZone.ID)
	d.Set("sub_interface_id", subinterfacedetail.SubInterfaceId)
	return diags
}
