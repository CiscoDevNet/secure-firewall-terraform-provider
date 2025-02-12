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

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"slices"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type PrefilterPolicy struct {
	Id                           types.String           `tfsdk:"id"`
	Domain                       types.String           `tfsdk:"domain"`
	Name                         types.String           `tfsdk:"name"`
	Description                  types.String           `tfsdk:"description"`
	Type                         types.String           `tfsdk:"type"`
	DefaultAction                types.String           `tfsdk:"default_action"`
	DefaultActionId              types.String           `tfsdk:"default_action_id"`
	DefaultActionLogBegin        types.Bool             `tfsdk:"default_action_log_begin"`
	DefaultActionLogEnd          types.Bool             `tfsdk:"default_action_log_end"`
	DefaultActionSendEventsToFmc types.Bool             `tfsdk:"default_action_send_events_to_fmc"`
	DefaultActionSyslogConfigId  types.String           `tfsdk:"default_action_syslog_config_id"`
	DefaultActionSnmpConfigId    types.String           `tfsdk:"default_action_snmp_config_id"`
	Rules                        []PrefilterPolicyRules `tfsdk:"rules"`
}

type PrefilterPolicyRules struct {
	Id                         types.String                                     `tfsdk:"id"`
	Name                       types.String                                     `tfsdk:"name"`
	RuleType                   types.String                                     `tfsdk:"rule_type"`
	Enabled                    types.Bool                                       `tfsdk:"enabled"`
	Action                     types.String                                     `tfsdk:"action"`
	Bidirectional              types.Bool                                       `tfsdk:"bidirectional"`
	TunnelZoneId               types.String                                     `tfsdk:"tunnel_zone_id"`
	TimeRangeId                types.String                                     `tfsdk:"time_range_id"`
	SourceInterfaces           []PrefilterPolicyRulesSourceInterfaces           `tfsdk:"source_interfaces"`
	DestinationInterfaces      []PrefilterPolicyRulesDestinationInterfaces      `tfsdk:"destination_interfaces"`
	SourceNetworkLiterals      []PrefilterPolicyRulesSourceNetworkLiterals      `tfsdk:"source_network_literals"`
	SourceNetworkObjects       []PrefilterPolicyRulesSourceNetworkObjects       `tfsdk:"source_network_objects"`
	DestinationNetworkLiterals []PrefilterPolicyRulesDestinationNetworkLiterals `tfsdk:"destination_network_literals"`
	DestinationNetworkObjects  []PrefilterPolicyRulesDestinationNetworkObjects  `tfsdk:"destination_network_objects"`
	VlanTagLiterals            []PrefilterPolicyRulesVlanTagLiterals            `tfsdk:"vlan_tag_literals"`
	VlanTagObjects             []PrefilterPolicyRulesVlanTagObjects             `tfsdk:"vlan_tag_objects"`
	SourcePortLiterals         []PrefilterPolicyRulesSourcePortLiterals         `tfsdk:"source_port_literals"`
	SourcePortObjects          []PrefilterPolicyRulesSourcePortObjects          `tfsdk:"source_port_objects"`
	DestinationPortLiterals    []PrefilterPolicyRulesDestinationPortLiterals    `tfsdk:"destination_port_literals"`
	DestinationPortObjects     []PrefilterPolicyRulesDestinationPortObjects     `tfsdk:"destination_port_objects"`
	EncapsulationPorts         types.Set                                        `tfsdk:"encapsulation_ports"`
	LogBegin                   types.Bool                                       `tfsdk:"log_begin"`
	LogEnd                     types.Bool                                       `tfsdk:"log_end"`
	SendEventsToFmc            types.Bool                                       `tfsdk:"send_events_to_fmc"`
	SendSyslog                 types.Bool                                       `tfsdk:"send_syslog"`
	SyslogConfigId             types.String                                     `tfsdk:"syslog_config_id"`
	SyslogSeverity             types.String                                     `tfsdk:"syslog_severity"`
	SnmpConfigId               types.String                                     `tfsdk:"snmp_config_id"`
}

type PrefilterPolicyRulesSourceInterfaces struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type PrefilterPolicyRulesDestinationInterfaces struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type PrefilterPolicyRulesSourceNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
}
type PrefilterPolicyRulesSourceNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type PrefilterPolicyRulesDestinationNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
}
type PrefilterPolicyRulesDestinationNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type PrefilterPolicyRulesVlanTagLiterals struct {
	StartTag types.String `tfsdk:"start_tag"`
	EndTag   types.String `tfsdk:"end_tag"`
}
type PrefilterPolicyRulesVlanTagObjects struct {
	Id types.String `tfsdk:"id"`
}
type PrefilterPolicyRulesSourcePortLiterals struct {
	Protocol types.String `tfsdk:"protocol"`
	Port     types.String `tfsdk:"port"`
}
type PrefilterPolicyRulesSourcePortObjects struct {
	Id types.String `tfsdk:"id"`
}
type PrefilterPolicyRulesDestinationPortLiterals struct {
	Protocol types.String `tfsdk:"protocol"`
	Port     types.String `tfsdk:"port"`
}
type PrefilterPolicyRulesDestinationPortObjects struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data PrefilterPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/prefilterpolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data PrefilterPolicy) toBody(ctx context.Context, state PrefilterPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.action", data.DefaultAction.ValueString())
	}
	if !data.DefaultActionLogBegin.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.logBegin", data.DefaultActionLogBegin.ValueBool())
	}
	if !data.DefaultActionLogEnd.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.logEnd", data.DefaultActionLogEnd.ValueBool())
	}
	if !data.DefaultActionSendEventsToFmc.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.sendEventsToFMC", data.DefaultActionSendEventsToFmc.ValueBool())
	}
	if !data.DefaultActionSyslogConfigId.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.syslogConfig.id", data.DefaultActionSyslogConfigId.ValueString())
	}
	if !data.DefaultActionSnmpConfigId.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.snmpConfig.id", data.DefaultActionSnmpConfigId.ValueString())
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "dummy_rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "PrefilterRule")
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.RuleType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ruleType", item.RuleType.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.Bidirectional.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bidirectional", item.Bidirectional.ValueBool())
			}
			if !item.TunnelZoneId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tunnelZone.objects.0.id", item.TunnelZoneId.ValueString())
			}
			if !item.TimeRangeId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "timeRangeObjects.0.id", item.TimeRangeId.ValueString())
			}
			if len(item.SourceInterfaces) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceInterfaces.objects", []interface{}{})
				for _, childItem := range item.SourceInterfaces {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceInterfaces.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationInterfaces) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationInterfaces.objects", []interface{}{})
				for _, childItem := range item.DestinationInterfaces {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationInterfaces.objects.-1", itemChildBody)
				}
			}
			if len(item.SourceNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.literals", []interface{}{})
				for _, childItem := range item.SourceNetworkLiterals {
					itemChildBody := ""
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.literals.-1", itemChildBody)
				}
			}
			if len(item.SourceNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.objects", []interface{}{})
				for _, childItem := range item.SourceNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.literals", []interface{}{})
				for _, childItem := range item.DestinationNetworkLiterals {
					itemChildBody := ""
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.literals.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.objects", []interface{}{})
				for _, childItem := range item.DestinationNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.VlanTagLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "vlanTags.literals", []interface{}{})
				for _, childItem := range item.VlanTagLiterals {
					itemChildBody := ""
					if !childItem.StartTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "startTag", childItem.StartTag.ValueString())
					}
					if !childItem.EndTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "endTag", childItem.EndTag.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "vlanTags.literals.-1", itemChildBody)
				}
			}
			if len(item.VlanTagObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "vlanTags.objects", []interface{}{})
				for _, childItem := range item.VlanTagObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "vlanTags.objects.-1", itemChildBody)
				}
			}
			if len(item.SourcePortLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourcePorts.literals", []interface{}{})
				for _, childItem := range item.SourcePortLiterals {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "PortLiteral")
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "protocol", childItem.Protocol.ValueString())
					}
					if !childItem.Port.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "port", childItem.Port.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.literals.-1", itemChildBody)
				}
			}
			if len(item.SourcePortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourcePorts.objects", []interface{}{})
				for _, childItem := range item.SourcePortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationPortLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationPorts.literals", []interface{}{})
				for _, childItem := range item.DestinationPortLiterals {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "PortLiteral")
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "protocol", childItem.Protocol.ValueString())
					}
					if !childItem.Port.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "port", childItem.Port.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.literals.-1", itemChildBody)
				}
			}
			if len(item.DestinationPortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationPorts.objects", []interface{}{})
				for _, childItem := range item.DestinationPortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.objects.-1", itemChildBody)
				}
			}
			if !item.EncapsulationPorts.IsNull() {
				var values []string
				item.EncapsulationPorts.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "encapsulationPorts", values)
			}
			if !item.LogBegin.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logBegin", item.LogBegin.ValueBool())
			}
			if !item.LogEnd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logEnd", item.LogEnd.ValueBool())
			}
			if !item.SendEventsToFmc.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sendEventsToFMC", item.SendEventsToFmc.ValueBool())
			}
			if !item.SendSyslog.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableSyslog", item.SendSyslog.ValueBool())
			}
			if !item.SyslogConfigId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "syslogConfig.id", item.SyslogConfigId.ValueString())
			}
			if !item.SyslogSeverity.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "syslogSeverity", item.SyslogSeverity.ValueString())
			}
			if !item.SnmpConfigId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "snmpConfig.id", item.SnmpConfigId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dummy_rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *PrefilterPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("defaultAction.action"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("defaultAction.id"); value.Exists() {
		data.DefaultActionId = types.StringValue(value.String())
	} else {
		data.DefaultActionId = types.StringNull()
	}
	if value := res.Get("defaultAction.logBegin"); value.Exists() {
		data.DefaultActionLogBegin = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionLogBegin = types.BoolNull()
	}
	if value := res.Get("defaultAction.logEnd"); value.Exists() {
		data.DefaultActionLogEnd = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionLogEnd = types.BoolNull()
	}
	if value := res.Get("defaultAction.sendEventsToFMC"); value.Exists() {
		data.DefaultActionSendEventsToFmc = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionSendEventsToFmc = types.BoolNull()
	}
	if value := res.Get("defaultAction.syslogConfig.id"); value.Exists() {
		data.DefaultActionSyslogConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSyslogConfigId = types.StringNull()
	}
	if value := res.Get("defaultAction.snmpConfig.id"); value.Exists() {
		data.DefaultActionSnmpConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSnmpConfigId = types.StringNull()
	}
	if value := res.Get("dummy_rules"); value.Exists() {
		data.Rules = make([]PrefilterPolicyRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PrefilterPolicyRules{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("ruleType"); value.Exists() {
				data.RuleType = types.StringValue(value.String())
			} else {
				data.RuleType = types.StringNull()
			}
			if value := res.Get("enabled"); value.Exists() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolValue(true)
			}
			if value := res.Get("action"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("bidirectional"); value.Exists() {
				data.Bidirectional = types.BoolValue(value.Bool())
			} else {
				data.Bidirectional = types.BoolNull()
			}
			if value := res.Get("tunnelZone.objects.0.id"); value.Exists() {
				data.TunnelZoneId = types.StringValue(value.String())
			} else {
				data.TunnelZoneId = types.StringNull()
			}
			if value := res.Get("timeRangeObjects.0.id"); value.Exists() {
				data.TimeRangeId = types.StringValue(value.String())
			} else {
				data.TimeRangeId = types.StringNull()
			}
			if value := res.Get("sourceInterfaces.objects"); value.Exists() {
				data.SourceInterfaces = make([]PrefilterPolicyRulesSourceInterfaces, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesSourceInterfaces{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).SourceInterfaces = append((*parent).SourceInterfaces, data)
					return true
				})
			}
			if value := res.Get("destinationInterfaces.objects"); value.Exists() {
				data.DestinationInterfaces = make([]PrefilterPolicyRulesDestinationInterfaces, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesDestinationInterfaces{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).DestinationInterfaces = append((*parent).DestinationInterfaces, data)
					return true
				})
			}
			if value := res.Get("sourceNetworks.literals"); value.Exists() {
				data.SourceNetworkLiterals = make([]PrefilterPolicyRulesSourceNetworkLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesSourceNetworkLiterals{}
					if value := res.Get("value"); value.Exists() {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					(*parent).SourceNetworkLiterals = append((*parent).SourceNetworkLiterals, data)
					return true
				})
			}
			if value := res.Get("sourceNetworks.objects"); value.Exists() {
				data.SourceNetworkObjects = make([]PrefilterPolicyRulesSourceNetworkObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesSourceNetworkObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).SourceNetworkObjects = append((*parent).SourceNetworkObjects, data)
					return true
				})
			}
			if value := res.Get("destinationNetworks.literals"); value.Exists() {
				data.DestinationNetworkLiterals = make([]PrefilterPolicyRulesDestinationNetworkLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesDestinationNetworkLiterals{}
					if value := res.Get("value"); value.Exists() {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					(*parent).DestinationNetworkLiterals = append((*parent).DestinationNetworkLiterals, data)
					return true
				})
			}
			if value := res.Get("destinationNetworks.objects"); value.Exists() {
				data.DestinationNetworkObjects = make([]PrefilterPolicyRulesDestinationNetworkObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesDestinationNetworkObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).DestinationNetworkObjects = append((*parent).DestinationNetworkObjects, data)
					return true
				})
			}
			if value := res.Get("vlanTags.literals"); value.Exists() {
				data.VlanTagLiterals = make([]PrefilterPolicyRulesVlanTagLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesVlanTagLiterals{}
					if value := res.Get("startTag"); value.Exists() {
						data.StartTag = types.StringValue(value.String())
					} else {
						data.StartTag = types.StringNull()
					}
					if value := res.Get("endTag"); value.Exists() {
						data.EndTag = types.StringValue(value.String())
					} else {
						data.EndTag = types.StringNull()
					}
					(*parent).VlanTagLiterals = append((*parent).VlanTagLiterals, data)
					return true
				})
			}
			if value := res.Get("vlanTags.objects"); value.Exists() {
				data.VlanTagObjects = make([]PrefilterPolicyRulesVlanTagObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesVlanTagObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).VlanTagObjects = append((*parent).VlanTagObjects, data)
					return true
				})
			}
			if value := res.Get("sourcePorts.literals"); value.Exists() {
				data.SourcePortLiterals = make([]PrefilterPolicyRulesSourcePortLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesSourcePortLiterals{}
					if value := res.Get("protocol"); value.Exists() {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("port"); value.Exists() {
						data.Port = types.StringValue(value.String())
					} else {
						data.Port = types.StringNull()
					}
					(*parent).SourcePortLiterals = append((*parent).SourcePortLiterals, data)
					return true
				})
			}
			if value := res.Get("sourcePorts.objects"); value.Exists() {
				data.SourcePortObjects = make([]PrefilterPolicyRulesSourcePortObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesSourcePortObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).SourcePortObjects = append((*parent).SourcePortObjects, data)
					return true
				})
			}
			if value := res.Get("destinationPorts.literals"); value.Exists() {
				data.DestinationPortLiterals = make([]PrefilterPolicyRulesDestinationPortLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesDestinationPortLiterals{}
					if value := res.Get("protocol"); value.Exists() {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("port"); value.Exists() {
						data.Port = types.StringValue(value.String())
					} else {
						data.Port = types.StringNull()
					}
					(*parent).DestinationPortLiterals = append((*parent).DestinationPortLiterals, data)
					return true
				})
			}
			if value := res.Get("destinationPorts.objects"); value.Exists() {
				data.DestinationPortObjects = make([]PrefilterPolicyRulesDestinationPortObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := PrefilterPolicyRulesDestinationPortObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).DestinationPortObjects = append((*parent).DestinationPortObjects, data)
					return true
				})
			}
			if value := res.Get("encapsulationPorts"); value.Exists() {
				data.EncapsulationPorts = helpers.GetStringSet(value.Array())
			} else {
				data.EncapsulationPorts = types.SetNull(types.StringType)
			}
			if value := res.Get("logBegin"); value.Exists() {
				data.LogBegin = types.BoolValue(value.Bool())
			} else {
				data.LogBegin = types.BoolNull()
			}
			if value := res.Get("logEnd"); value.Exists() {
				data.LogEnd = types.BoolValue(value.Bool())
			} else {
				data.LogEnd = types.BoolNull()
			}
			if value := res.Get("sendEventsToFMC"); value.Exists() {
				data.SendEventsToFmc = types.BoolValue(value.Bool())
			} else {
				data.SendEventsToFmc = types.BoolNull()
			}
			if value := res.Get("enableSyslog"); value.Exists() {
				data.SendSyslog = types.BoolValue(value.Bool())
			} else {
				data.SendSyslog = types.BoolNull()
			}
			if value := res.Get("syslogConfig.id"); value.Exists() {
				data.SyslogConfigId = types.StringValue(value.String())
			} else {
				data.SyslogConfigId = types.StringNull()
			}
			if value := res.Get("syslogSeverity"); value.Exists() {
				data.SyslogSeverity = types.StringValue(value.String())
			} else {
				data.SyslogSeverity = types.StringNull()
			}
			if value := res.Get("snmpConfig.id"); value.Exists() {
				data.SnmpConfigId = types.StringValue(value.String())
			} else {
				data.SnmpConfigId = types.StringNull()
			}
			(*parent).Rules = append((*parent).Rules, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *PrefilterPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("defaultAction.action"); value.Exists() && !data.DefaultAction.IsNull() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("defaultAction.id"); value.Exists() && !data.DefaultActionId.IsNull() {
		data.DefaultActionId = types.StringValue(value.String())
	} else {
		data.DefaultActionId = types.StringNull()
	}
	if value := res.Get("defaultAction.logBegin"); value.Exists() && !data.DefaultActionLogBegin.IsNull() {
		data.DefaultActionLogBegin = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionLogBegin = types.BoolNull()
	}
	if value := res.Get("defaultAction.logEnd"); value.Exists() && !data.DefaultActionLogEnd.IsNull() {
		data.DefaultActionLogEnd = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionLogEnd = types.BoolNull()
	}
	if value := res.Get("defaultAction.sendEventsToFMC"); value.Exists() && !data.DefaultActionSendEventsToFmc.IsNull() {
		data.DefaultActionSendEventsToFmc = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionSendEventsToFmc = types.BoolNull()
	}
	if value := res.Get("defaultAction.syslogConfig.id"); value.Exists() && !data.DefaultActionSyslogConfigId.IsNull() {
		data.DefaultActionSyslogConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSyslogConfigId = types.StringNull()
	}
	if value := res.Get("defaultAction.snmpConfig.id"); value.Exists() && !data.DefaultActionSnmpConfigId.IsNull() {
		data.DefaultActionSnmpConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSnmpConfigId = types.StringNull()
	}
	{
		l := len(res.Get("dummy_rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("dummy_rules array resizing from %d to %d", len(data.Rules), l))
		for i := len(data.Rules); i < l; i++ {
			data.Rules = append(data.Rules, PrefilterPolicyRules{})
		}
		if len(data.Rules) > l {
			data.Rules = data.Rules[:l]
		}
	}
	for i := range data.Rules {
		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("dummy_rules.%d", i))
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("ruleType"); value.Exists() && !data.RuleType.IsNull() {
			data.RuleType = types.StringValue(value.String())
		} else {
			data.RuleType = types.StringNull()
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else if data.Enabled.ValueBool() != true {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
			data.Action = types.StringValue(value.String())
		} else {
			data.Action = types.StringNull()
		}
		if value := res.Get("bidirectional"); value.Exists() && !data.Bidirectional.IsNull() {
			data.Bidirectional = types.BoolValue(value.Bool())
		} else {
			data.Bidirectional = types.BoolNull()
		}
		if value := res.Get("tunnelZone.objects.0.id"); value.Exists() && !data.TunnelZoneId.IsNull() {
			data.TunnelZoneId = types.StringValue(value.String())
		} else {
			data.TunnelZoneId = types.StringNull()
		}
		if value := res.Get("timeRangeObjects.0.id"); value.Exists() && !data.TimeRangeId.IsNull() {
			data.TimeRangeId = types.StringValue(value.String())
		} else {
			data.TimeRangeId = types.StringNull()
		}
		for i := 0; i < len(data.SourceInterfaces); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceInterfaces[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceInterfaces[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourceInterfaces.objects").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing SourceInterfaces[%d] = %+v",
					i,
					(*parent).SourceInterfaces[i],
				))
				(*parent).SourceInterfaces = slices.Delete((*parent).SourceInterfaces, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).SourceInterfaces[i] = data
		}
		for i := 0; i < len(data.DestinationInterfaces); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationInterfaces[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationInterfaces[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationInterfaces.objects").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationInterfaces[%d] = %+v",
					i,
					(*parent).DestinationInterfaces[i],
				))
				(*parent).DestinationInterfaces = slices.Delete((*parent).DestinationInterfaces, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).DestinationInterfaces[i] = data
		}
		for i := 0; i < len(data.SourceNetworkLiterals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.SourceNetworkLiterals[i].Value.ValueString()}

			parent := &data
			data := (*parent).SourceNetworkLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourceNetworks.literals").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing SourceNetworkLiterals[%d] = %+v",
					i,
					(*parent).SourceNetworkLiterals[i],
				))
				(*parent).SourceNetworkLiterals = slices.Delete((*parent).SourceNetworkLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).SourceNetworkLiterals[i] = data
		}
		for i := 0; i < len(data.SourceNetworkObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceNetworkObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceNetworkObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourceNetworks.objects").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing SourceNetworkObjects[%d] = %+v",
					i,
					(*parent).SourceNetworkObjects[i],
				))
				(*parent).SourceNetworkObjects = slices.Delete((*parent).SourceNetworkObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).SourceNetworkObjects[i] = data
		}
		for i := 0; i < len(data.DestinationNetworkLiterals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.DestinationNetworkLiterals[i].Value.ValueString()}

			parent := &data
			data := (*parent).DestinationNetworkLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationNetworks.literals").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworkLiterals[%d] = %+v",
					i,
					(*parent).DestinationNetworkLiterals[i],
				))
				(*parent).DestinationNetworkLiterals = slices.Delete((*parent).DestinationNetworkLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).DestinationNetworkLiterals[i] = data
		}
		for i := 0; i < len(data.DestinationNetworkObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationNetworkObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationNetworkObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationNetworks.objects").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworkObjects[%d] = %+v",
					i,
					(*parent).DestinationNetworkObjects[i],
				))
				(*parent).DestinationNetworkObjects = slices.Delete((*parent).DestinationNetworkObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).DestinationNetworkObjects[i] = data
		}
		for i := 0; i < len(data.VlanTagLiterals); i++ {
			keys := [...]string{"startTag", "endTag"}
			keyValues := [...]string{data.VlanTagLiterals[i].StartTag.ValueString(), data.VlanTagLiterals[i].EndTag.ValueString()}

			parent := &data
			data := (*parent).VlanTagLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("vlanTags.literals").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing VlanTagLiterals[%d] = %+v",
					i,
					(*parent).VlanTagLiterals[i],
				))
				(*parent).VlanTagLiterals = slices.Delete((*parent).VlanTagLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("startTag"); value.Exists() && !data.StartTag.IsNull() {
				data.StartTag = types.StringValue(value.String())
			} else {
				data.StartTag = types.StringNull()
			}
			if value := res.Get("endTag"); value.Exists() && !data.EndTag.IsNull() {
				data.EndTag = types.StringValue(value.String())
			} else {
				data.EndTag = types.StringNull()
			}
			(*parent).VlanTagLiterals[i] = data
		}
		for i := 0; i < len(data.VlanTagObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.VlanTagObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).VlanTagObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("vlanTags.objects").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing VlanTagObjects[%d] = %+v",
					i,
					(*parent).VlanTagObjects[i],
				))
				(*parent).VlanTagObjects = slices.Delete((*parent).VlanTagObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).VlanTagObjects[i] = data
		}
		for i := 0; i < len(data.SourcePortLiterals); i++ {
			keys := [...]string{"protocol", "port"}
			keyValues := [...]string{data.SourcePortLiterals[i].Protocol.ValueString(), data.SourcePortLiterals[i].Port.ValueString()}

			parent := &data
			data := (*parent).SourcePortLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourcePorts.literals").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing SourcePortLiterals[%d] = %+v",
					i,
					(*parent).SourcePortLiterals[i],
				))
				(*parent).SourcePortLiterals = slices.Delete((*parent).SourcePortLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			(*parent).SourcePortLiterals[i] = data
		}
		for i := 0; i < len(data.SourcePortObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourcePortObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourcePortObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourcePorts.objects").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing SourcePortObjects[%d] = %+v",
					i,
					(*parent).SourcePortObjects[i],
				))
				(*parent).SourcePortObjects = slices.Delete((*parent).SourcePortObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourcePortObjects[i] = data
		}
		for i := 0; i < len(data.DestinationPortLiterals); i++ {
			keys := [...]string{"protocol", "port"}
			keyValues := [...]string{data.DestinationPortLiterals[i].Protocol.ValueString(), data.DestinationPortLiterals[i].Port.ValueString()}

			parent := &data
			data := (*parent).DestinationPortLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationPorts.literals").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationPortLiterals[%d] = %+v",
					i,
					(*parent).DestinationPortLiterals[i],
				))
				(*parent).DestinationPortLiterals = slices.Delete((*parent).DestinationPortLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			(*parent).DestinationPortLiterals[i] = data
		}
		for i := 0; i < len(data.DestinationPortObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationPortObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationPortObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationPorts.objects").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationPortObjects[%d] = %+v",
					i,
					(*parent).DestinationPortObjects[i],
				))
				(*parent).DestinationPortObjects = slices.Delete((*parent).DestinationPortObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationPortObjects[i] = data
		}
		if value := res.Get("encapsulationPorts"); value.Exists() && !data.EncapsulationPorts.IsNull() {
			data.EncapsulationPorts = helpers.GetStringSet(value.Array())
		} else {
			data.EncapsulationPorts = types.SetNull(types.StringType)
		}
		if value := res.Get("logBegin"); value.Exists() && !data.LogBegin.IsNull() {
			data.LogBegin = types.BoolValue(value.Bool())
		} else {
			data.LogBegin = types.BoolNull()
		}
		if value := res.Get("logEnd"); value.Exists() && !data.LogEnd.IsNull() {
			data.LogEnd = types.BoolValue(value.Bool())
		} else {
			data.LogEnd = types.BoolNull()
		}
		if value := res.Get("sendEventsToFMC"); value.Exists() && !data.SendEventsToFmc.IsNull() {
			data.SendEventsToFmc = types.BoolValue(value.Bool())
		} else {
			data.SendEventsToFmc = types.BoolNull()
		}
		if value := res.Get("enableSyslog"); value.Exists() && !data.SendSyslog.IsNull() {
			data.SendSyslog = types.BoolValue(value.Bool())
		} else {
			data.SendSyslog = types.BoolNull()
		}
		if value := res.Get("syslogConfig.id"); value.Exists() && !data.SyslogConfigId.IsNull() {
			data.SyslogConfigId = types.StringValue(value.String())
		} else {
			data.SyslogConfigId = types.StringNull()
		}
		if value := res.Get("syslogSeverity"); value.Exists() && !data.SyslogSeverity.IsNull() {
			data.SyslogSeverity = types.StringValue(value.String())
		} else {
			data.SyslogSeverity = types.StringNull()
		}
		if value := res.Get("snmpConfig.id"); value.Exists() && !data.SnmpConfigId.IsNull() {
			data.SnmpConfigId = types.StringValue(value.String())
		} else {
			data.SnmpConfigId = types.StringNull()
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *PrefilterPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.DefaultActionId.IsUnknown() {
		if value := res.Get("defaultAction.id"); value.Exists() {
			data.DefaultActionId = types.StringValue(value.String())
		} else {
			data.DefaultActionId = types.StringNull()
		}
	}
	for i := range data.Rules {
		r := res.Get(fmt.Sprintf("dummy_rules.%d", i))
		if v := data.Rules[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Rules[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
