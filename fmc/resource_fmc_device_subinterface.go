package fmc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSubInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSubInterfaceCreate,
		ReadContext:   resourceSubInterfaceRead,
		UpdateContext: resourceSubInterfaceUpdate,
		DeleteContext: resourceSubInterfaceDelete,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Device Id of SUBI",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "MulticastGroupAddress of the SUBI",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SegmentId of the SUBI",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SUBI description",
			},
			"sub_interface_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The Device Id of SUBI",
			},
			"if_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Logical Name of chosen interface",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "SUBID of the SUBI",
			},
			"priority": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Priority of the SUBI",
			},
			"management_only": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "EnableProxy of the SUBI",
			},
			"security_zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "securityZone of the SUBI",
			},

			"vlan_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "SUBID of the SUBI",
			},

			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "EnableProxy of the SUBI",
			},
			"ipv4": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP of the interface",
									},
									"netmask": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Net mask of the interface",
									},
								},
							},
						},
						"dhcp": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_default_route_dhcp": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
										Description: "Dynamic IP of the interface",
									},
									"dhcp_route_metric": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
									},
								},
							},
						},
					},
				},
				Description: "IPV4 information",
			},
		},
	}
}

func resourceSubInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	log.Printf("SUBI: Fetching SUBI details")

	if d.Get("device_id") != nil && d.Get("id") != nil {

		deviceId := d.Get("device_id").(string)
		id := d.Get("id").(string)

		log.Printf("SUBI: DeviceId=%s, Id=%s", deviceId, id)

		c := m.(*Client)

		response, err := c.GetFmcSubInterfaceDetails(ctx, deviceId, id)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read SUBI",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("SUBI: SUBI Details Id=%s Name=%s IFName=%s", response.ID, response.Name, response.Ifname)

		d.SetId(response.ID)

	} else {
		log.Printf("SUBI: No Proper Data Of Device Id=%s, Id=%s", d.Get("device_id"), d.Get("id"))
	}

	return diags
}

func resourceSubInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("SUBI: Creating sub interface")

	device_id := d.Get("device_id").(string)
	subitype := d.Get("type").(string)
	name := d.Get("name").(string)
	ifname := d.Get("if_name").(string)
	description := d.Get("description").(string)
	security_zone_id := d.Get("security_zone_id").(string)
	priority := d.Get("priority").(int)
	mtu := d.Get("mtu").(int)
	vlanId := d.Get("vlan_id").(int)
	enabled := d.Get("enabled").(bool)
	managementOnly := d.Get("management_only").(bool)
	sub_interface_id := d.Get("sub_interface_id").(int)

	var ipv4 = IPv4{}

	if d.Get("ipv4") != nil {
		if ipv4Entries, ok := d.GetOk("ipv4"); ok {
			log.Printf("IPv4s data=%v", ipv4Entries)
			entries := ipv4Entries.([]interface{})[0]
			ipV4Entry := entries.(map[string]interface{})

			isStatic := false

			if ipV4Entry["static"] != nil {
				statics := ipV4Entry["static"].([]interface{})
				if len(statics) > 0 {
					static := statics[0].(map[string]interface{})
					var IPv4Static = IPv4Static{
						Address: static["address"].(string),
						Netmask: static["netmask"].(int),
					}
					ipv4.Static = &IPv4Static
					isStatic = true
				}

			}

			if ipV4Entry["dhcp"] != nil && !isStatic {

				dhcps := ipV4Entry["dhcp"].([]interface{})
				dhcp := dhcps[0].(map[string]interface{})

				var IPv4DHCP = IPv4DHCP{
					Enable:      dhcp["enable_default_route_dhcp"].(bool),
					RouteMetric: dhcp["dhcp_route_metric"].(int),
				}
				ipv4.DHCP = &IPv4DHCP
			}

		}

	}

	log.Printf("device_id=%s name=%s, subitype=%s, ifname=%v, description=%v security_zone_id=%s priority=%v mtu=%v  vlanId=%v  sub_interface_id=%v  enabled=%v  managementOnly=%v priority=%v", device_id, name, subitype, ifname, description, security_zone_id, priority, mtu, vlanId, sub_interface_id, enabled, managementOnly, priority)

	c := m.(*Client)

	var diags diag.Diagnostics

	var securityZone = SubInterfaceSecurityZone{
		ID:   security_zone_id,
		Type: "SecurityZone",
	}

	response, err := c.CreateFmcSubInterface(ctx, device_id, &SubInterfaceRequest{
		Type:           subitype,
		Name:           name,
		Description:    description,
		VLANId:         vlanId,
		SubInterfaceId: sub_interface_id,
		MTU:            mtu,
		Priority:       priority,
		ManagementOnly: managementOnly,
		Ifname:         ifname,
		Enabled:        enabled,
		SecurityZone:   securityZone,
		IPv4:           ipv4,
	})

	if err != nil {
		log.Printf("FPU: err=%s", err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to crete SUBI",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("SUB Id=%s", response.ID)

	d.SetId(response.ID)
	return diags
}

func resourceSubInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("SUBI: Updating sub interface")

	var diags diag.Diagnostics

	id := d.Id()

	log.Printf("Updating Id=%v", id)

	if len(id) > 0 {
		device_id := d.Get("device_id").(string)
		subitype := d.Get("type").(string)
		name := d.Get("name").(string)
		ifname := d.Get("if_name").(string)
		description := d.Get("description").(string)
		security_zone_id := d.Get("security_zone_id").(string)
		priority := d.Get("priority").(int)
		mtu := d.Get("mtu").(int)
		vlanId := d.Get("vlan_id").(int)
		enabled := d.Get("enabled").(bool)
		managementOnly := d.Get("management_only").(bool)
		sub_interface_id := d.Get("sub_interface_id").(int)

		var ipv4 = IPv4{}

		if d.Get("ipv4") != nil {
			if ipv4Entries, ok := d.GetOk("ipv4"); ok {
				log.Printf("IPv4s data=%v", ipv4Entries)
				entries := ipv4Entries.([]interface{})[0]
				ipV4Entry := entries.(map[string]interface{})
				isStatic := false
				if ipV4Entry["static"] != nil {
					statics := ipV4Entry["static"].([]interface{})
					if len(statics) > 0 {
						static := statics[0].(map[string]interface{})
						var IPv4Static = IPv4Static{
							Address: static["address"].(string),
							Netmask: static["netmask"].(int),
						}
						log.Printf("IPv4Static=%v", IPv4Static)
						ipv4.Static = &IPv4Static
						isStatic = true
					}

				}

				if ipV4Entry["dhcp"] != nil && !isStatic {

					dhcps := ipV4Entry["dhcp"].([]interface{})
					if len(dhcps) > 0 {
						dhcp := dhcps[0].(map[string]interface{})

						var IPv4DHCP = IPv4DHCP{
							Enable:      dhcp["enable_default_route_dhcp"].(bool),
							RouteMetric: dhcp["dhcp_route_metric"].(int),
						}
						ipv4.DHCP = &IPv4DHCP
					}

				}

			}

		}

		log.Printf("device_id=%s name=%s, subitype=%s, ifname=%v, description=%v security_zone_id=%s priority=%v mtu=%v  vlanId=%v  sub_interface_id=%v  enabled=%v  managementOnly=%v priority=%v", device_id, name, subitype, ifname, description, security_zone_id, priority, mtu, vlanId, sub_interface_id, enabled, managementOnly, priority)

		c := m.(*Client)

		oldsubi, err := c.GetFmcSubInterfaceDetails(ctx, device_id, id)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update SUBI as not able to SUBI details based on id",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("Old sub interface=%v", oldsubi)

		var securityZone = SubInterfaceSecurityZone{
			ID:   security_zone_id,
			Type: "SecurityZone",
		}

		response, err := c.UpdateFmcSubInterface(ctx, device_id, id, &SubInterfaceRequest{
			ID:             id,
			Type:           subitype,
			Name:           name,
			Description:    description,
			VLANId:         vlanId,
			SubInterfaceId: sub_interface_id,
			MTU:            mtu,
			Priority:       priority,
			ManagementOnly: managementOnly,
			Ifname:         ifname,
			Enabled:        enabled,
			SecurityZone:   securityZone,
			IPv4:           ipv4,
		})

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update sub interface",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("Updated sub interface Id=%s", response.ID)

		d.SetId(response.ID)

	} else {
		log.Printf("Update sub interface UUID is null, so record cannot be updated")

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to update vni",
			Detail:   "sub interface UUID is null, so record cannot be updated",
		})
	}

	return diags
}

func resourceSubInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("sub interface: Delete sub interface")

	var diags diag.Diagnostics

	id := d.Id()

	log.Printf("Id=%v", id)

	if len(id) > 0 && len(d.Get("device_id").(string)) > 0 {
		device_id := d.Get("device_id").(string)

		log.Printf("device_id=%v ", device_id)

		c := m.(*Client)

		response, err := c.DeleteFmcSubInterfaceDetails(ctx, device_id, id)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to delete sub interface as not able to get sub interface details based on id",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("Deleted sub interface Id=%s", response.ID)

		d.SetId(response.ID)
	} else {
		log.Printf("Unable to delete the sub interface as device_id or id is not present")
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete the sub interface as device_id or id is not present",
			Detail:   "ID is null, so record cannot be deleted",
		})
	}

	return diags
}
