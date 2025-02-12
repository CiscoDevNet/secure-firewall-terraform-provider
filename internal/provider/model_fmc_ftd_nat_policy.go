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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDNATPolicy struct {
	Id             types.String                 `tfsdk:"id"`
	Domain         types.String                 `tfsdk:"domain"`
	Name           types.String                 `tfsdk:"name"`
	Description    types.String                 `tfsdk:"description"`
	Type           types.String                 `tfsdk:"type"`
	ManualNatRules []FTDNATPolicyManualNatRules `tfsdk:"manual_nat_rules"`
	AutoNatRules   []FTDNATPolicyAutoNatRules   `tfsdk:"auto_nat_rules"`
}

type FTDNATPolicyManualNatRules struct {
	Id                             types.String `tfsdk:"id"`
	Description                    types.String `tfsdk:"description"`
	Enabled                        types.Bool   `tfsdk:"enabled"`
	Section                        types.String `tfsdk:"section"`
	NatType                        types.String `tfsdk:"nat_type"`
	FallThrough                    types.Bool   `tfsdk:"fall_through"`
	InterfaceInOriginalDestination types.Bool   `tfsdk:"interface_in_original_destination"`
	InterfaceInTranslatedSource    types.Bool   `tfsdk:"interface_in_translated_source"`
	Ipv6                           types.Bool   `tfsdk:"ipv6"`
	NetToNet                       types.Bool   `tfsdk:"net_to_net"`
	NoProxyArp                     types.Bool   `tfsdk:"no_proxy_arp"`
	Unidirectional                 types.Bool   `tfsdk:"unidirectional"`
	SourceInterfaceId              types.String `tfsdk:"source_interface_id"`
	OriginalSourceId               types.String `tfsdk:"original_source_id"`
	OriginalSourcePortId           types.String `tfsdk:"original_source_port_id"`
	OriginalDestinationId          types.String `tfsdk:"original_destination_id"`
	OriginalDestinationPortId      types.String `tfsdk:"original_destination_port_id"`
	RouteLookup                    types.Bool   `tfsdk:"route_lookup"`
	DestinationInterfaceId         types.String `tfsdk:"destination_interface_id"`
	TranslatedSourceId             types.String `tfsdk:"translated_source_id"`
	TranslatedSourcePortId         types.String `tfsdk:"translated_source_port_id"`
	TranslateDns                   types.Bool   `tfsdk:"translate_dns"`
	TranslatedDestinationId        types.String `tfsdk:"translated_destination_id"`
	TranslatedDestinationPortId    types.String `tfsdk:"translated_destination_port_id"`
}

type FTDNATPolicyAutoNatRules struct {
	Id                                      types.String `tfsdk:"id"`
	NatType                                 types.String `tfsdk:"nat_type"`
	DestinationInterfaceId                  types.String `tfsdk:"destination_interface_id"`
	FallThrough                             types.Bool   `tfsdk:"fall_through"`
	Ipv6                                    types.Bool   `tfsdk:"ipv6"`
	NetToNet                                types.Bool   `tfsdk:"net_to_net"`
	NoProxyArp                              types.Bool   `tfsdk:"no_proxy_arp"`
	OriginalNetworkId                       types.String `tfsdk:"original_network_id"`
	OriginalPort                            types.Int64  `tfsdk:"original_port"`
	Protocol                                types.String `tfsdk:"protocol"`
	PerformRouteLookup                      types.Bool   `tfsdk:"perform_route_lookup"`
	SourceInterfaceId                       types.String `tfsdk:"source_interface_id"`
	TranslateDns                            types.Bool   `tfsdk:"translate_dns"`
	TranslatedNetworkId                     types.String `tfsdk:"translated_network_id"`
	TranslatedNetworkIsDestinationInterface types.Bool   `tfsdk:"translated_network_is_destination_interface"`
	TranslatedPort                          types.Int64  `tfsdk:"translated_port"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDNATPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdnatpolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDNATPolicy) toBody(ctx context.Context, state FTDNATPolicy) string {
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
	if len(data.ManualNatRules) > 0 {
		body, _ = sjson.Set(body, "dummy_manual_nat_rules", []interface{}{})
		for _, item := range data.ManualNatRules {
			itemBody := ""
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if !item.Section.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "metadata.section", item.Section.ValueString())
			}
			if !item.NatType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "natType", item.NatType.ValueString())
			}
			if !item.FallThrough.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fallThrough", item.FallThrough.ValueBool())
			}
			if !item.InterfaceInOriginalDestination.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceInOriginalDestination", item.InterfaceInOriginalDestination.ValueBool())
			}
			if !item.InterfaceInTranslatedSource.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceInTranslatedSource", item.InterfaceInTranslatedSource.ValueBool())
			}
			if !item.Ipv6.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceIpv6", item.Ipv6.ValueBool())
			}
			if !item.NetToNet.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "netToNet", item.NetToNet.ValueBool())
			}
			if !item.NoProxyArp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "noProxyArp", item.NoProxyArp.ValueBool())
			}
			if !item.Unidirectional.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "unidirectional", item.Unidirectional.ValueBool())
			}
			if !item.SourceInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceInterface.id", item.SourceInterfaceId.ValueString())
			}
			if !item.OriginalSourceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "originalSource.id", item.OriginalSourceId.ValueString())
			}
			if !item.OriginalSourcePortId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "originalSourcePort.id", item.OriginalSourcePortId.ValueString())
			}
			if !item.OriginalDestinationId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "originalDestination.id", item.OriginalDestinationId.ValueString())
			}
			if !item.OriginalDestinationPortId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "originalDestinationPort.id", item.OriginalDestinationPortId.ValueString())
			}
			if !item.RouteLookup.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeLookup", item.RouteLookup.ValueBool())
			}
			if !item.DestinationInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationInterface.id", item.DestinationInterfaceId.ValueString())
			}
			if !item.TranslatedSourceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "translatedSource.id", item.TranslatedSourceId.ValueString())
			}
			if !item.TranslatedSourcePortId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "translatedSourcePort.id", item.TranslatedSourcePortId.ValueString())
			}
			if !item.TranslateDns.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dns", item.TranslateDns.ValueBool())
			}
			if !item.TranslatedDestinationId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "translatedDestination.id", item.TranslatedDestinationId.ValueString())
			}
			if !item.TranslatedDestinationPortId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "translatedDestinationPort.id", item.TranslatedDestinationPortId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dummy_manual_nat_rules.-1", itemBody)
		}
	}
	if len(data.AutoNatRules) > 0 {
		body, _ = sjson.Set(body, "dummy_auto_nat_rules", []interface{}{})
		for _, item := range data.AutoNatRules {
			itemBody := ""
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.NatType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "natType", item.NatType.ValueString())
			}
			if !item.DestinationInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationInterface.id", item.DestinationInterfaceId.ValueString())
			}
			if !item.FallThrough.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fallThrough", item.FallThrough.ValueBool())
			}
			if !item.Ipv6.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceIpv6", item.Ipv6.ValueBool())
			}
			if !item.NetToNet.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "netToNet", item.NetToNet.ValueBool())
			}
			if !item.NoProxyArp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "noProxyArp", item.NoProxyArp.ValueBool())
			}
			if !item.OriginalNetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "originalNetwork.id", item.OriginalNetworkId.ValueString())
			}
			if !item.OriginalPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "originalPort", item.OriginalPort.ValueInt64())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serviceProtocol", item.Protocol.ValueString())
			}
			if !item.PerformRouteLookup.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeLookup", item.PerformRouteLookup.ValueBool())
			}
			if !item.SourceInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceInterface.id", item.SourceInterfaceId.ValueString())
			}
			if !item.TranslateDns.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dns", item.TranslateDns.ValueBool())
			}
			if !item.TranslatedNetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "translatedNetwork.id", item.TranslatedNetworkId.ValueString())
			}
			if !item.TranslatedNetworkIsDestinationInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceInTranslatedNetwork", item.TranslatedNetworkIsDestinationInterface.ValueBool())
			}
			if !item.TranslatedPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "translatedPort", item.TranslatedPort.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "dummy_auto_nat_rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDNATPolicy) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("dummy_manual_nat_rules"); value.Exists() {
		data.ManualNatRules = make([]FTDNATPolicyManualNatRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDNATPolicyManualNatRules{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("description"); value.Exists() {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			if value := res.Get("enabled"); value.Exists() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolNull()
			}
			if value := res.Get("metadata.section"); value.Exists() {
				data.Section = types.StringValue(value.String())
			} else {
				data.Section = types.StringNull()
			}
			if value := res.Get("natType"); value.Exists() {
				data.NatType = types.StringValue(value.String())
			} else {
				data.NatType = types.StringNull()
			}
			if value := res.Get("fallThrough"); value.Exists() {
				data.FallThrough = types.BoolValue(value.Bool())
			} else {
				data.FallThrough = types.BoolNull()
			}
			if value := res.Get("interfaceInOriginalDestination"); value.Exists() {
				data.InterfaceInOriginalDestination = types.BoolValue(value.Bool())
			} else {
				data.InterfaceInOriginalDestination = types.BoolNull()
			}
			if value := res.Get("interfaceInTranslatedSource"); value.Exists() {
				data.InterfaceInTranslatedSource = types.BoolValue(value.Bool())
			} else {
				data.InterfaceInTranslatedSource = types.BoolNull()
			}
			if value := res.Get("interfaceIpv6"); value.Exists() {
				data.Ipv6 = types.BoolValue(value.Bool())
			} else {
				data.Ipv6 = types.BoolNull()
			}
			if value := res.Get("netToNet"); value.Exists() {
				data.NetToNet = types.BoolValue(value.Bool())
			} else {
				data.NetToNet = types.BoolNull()
			}
			if value := res.Get("noProxyArp"); value.Exists() {
				data.NoProxyArp = types.BoolValue(value.Bool())
			} else {
				data.NoProxyArp = types.BoolNull()
			}
			if value := res.Get("unidirectional"); value.Exists() {
				data.Unidirectional = types.BoolValue(value.Bool())
			} else {
				data.Unidirectional = types.BoolNull()
			}
			if value := res.Get("sourceInterface.id"); value.Exists() {
				data.SourceInterfaceId = types.StringValue(value.String())
			} else {
				data.SourceInterfaceId = types.StringNull()
			}
			if value := res.Get("originalSource.id"); value.Exists() {
				data.OriginalSourceId = types.StringValue(value.String())
			} else {
				data.OriginalSourceId = types.StringNull()
			}
			if value := res.Get("originalSourcePort.id"); value.Exists() {
				data.OriginalSourcePortId = types.StringValue(value.String())
			} else {
				data.OriginalSourcePortId = types.StringNull()
			}
			if value := res.Get("originalDestination.id"); value.Exists() {
				data.OriginalDestinationId = types.StringValue(value.String())
			} else {
				data.OriginalDestinationId = types.StringNull()
			}
			if value := res.Get("originalDestinationPort.id"); value.Exists() {
				data.OriginalDestinationPortId = types.StringValue(value.String())
			} else {
				data.OriginalDestinationPortId = types.StringNull()
			}
			if value := res.Get("routeLookup"); value.Exists() {
				data.RouteLookup = types.BoolValue(value.Bool())
			} else {
				data.RouteLookup = types.BoolNull()
			}
			if value := res.Get("destinationInterface.id"); value.Exists() {
				data.DestinationInterfaceId = types.StringValue(value.String())
			} else {
				data.DestinationInterfaceId = types.StringNull()
			}
			if value := res.Get("translatedSource.id"); value.Exists() {
				data.TranslatedSourceId = types.StringValue(value.String())
			} else {
				data.TranslatedSourceId = types.StringNull()
			}
			if value := res.Get("translatedSourcePort.id"); value.Exists() {
				data.TranslatedSourcePortId = types.StringValue(value.String())
			} else {
				data.TranslatedSourcePortId = types.StringNull()
			}
			if value := res.Get("dns"); value.Exists() {
				data.TranslateDns = types.BoolValue(value.Bool())
			} else {
				data.TranslateDns = types.BoolNull()
			}
			if value := res.Get("translatedDestination.id"); value.Exists() {
				data.TranslatedDestinationId = types.StringValue(value.String())
			} else {
				data.TranslatedDestinationId = types.StringNull()
			}
			if value := res.Get("translatedDestinationPort.id"); value.Exists() {
				data.TranslatedDestinationPortId = types.StringValue(value.String())
			} else {
				data.TranslatedDestinationPortId = types.StringNull()
			}
			(*parent).ManualNatRules = append((*parent).ManualNatRules, data)
			return true
		})
	}
	if value := res.Get("dummy_auto_nat_rules"); value.Exists() {
		data.AutoNatRules = make([]FTDNATPolicyAutoNatRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDNATPolicyAutoNatRules{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("natType"); value.Exists() {
				data.NatType = types.StringValue(value.String())
			} else {
				data.NatType = types.StringNull()
			}
			if value := res.Get("destinationInterface.id"); value.Exists() {
				data.DestinationInterfaceId = types.StringValue(value.String())
			} else {
				data.DestinationInterfaceId = types.StringNull()
			}
			if value := res.Get("fallThrough"); value.Exists() {
				data.FallThrough = types.BoolValue(value.Bool())
			} else {
				data.FallThrough = types.BoolNull()
			}
			if value := res.Get("interfaceIpv6"); value.Exists() {
				data.Ipv6 = types.BoolValue(value.Bool())
			} else {
				data.Ipv6 = types.BoolNull()
			}
			if value := res.Get("netToNet"); value.Exists() {
				data.NetToNet = types.BoolValue(value.Bool())
			} else {
				data.NetToNet = types.BoolNull()
			}
			if value := res.Get("noProxyArp"); value.Exists() {
				data.NoProxyArp = types.BoolValue(value.Bool())
			} else {
				data.NoProxyArp = types.BoolNull()
			}
			if value := res.Get("originalNetwork.id"); value.Exists() {
				data.OriginalNetworkId = types.StringValue(value.String())
			} else {
				data.OriginalNetworkId = types.StringNull()
			}
			if value := res.Get("originalPort"); value.Exists() {
				data.OriginalPort = types.Int64Value(value.Int())
			} else {
				data.OriginalPort = types.Int64Null()
			}
			if value := res.Get("serviceProtocol"); value.Exists() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("routeLookup"); value.Exists() {
				data.PerformRouteLookup = types.BoolValue(value.Bool())
			} else {
				data.PerformRouteLookup = types.BoolNull()
			}
			if value := res.Get("sourceInterface.id"); value.Exists() {
				data.SourceInterfaceId = types.StringValue(value.String())
			} else {
				data.SourceInterfaceId = types.StringNull()
			}
			if value := res.Get("dns"); value.Exists() {
				data.TranslateDns = types.BoolValue(value.Bool())
			} else {
				data.TranslateDns = types.BoolNull()
			}
			if value := res.Get("translatedNetwork.id"); value.Exists() {
				data.TranslatedNetworkId = types.StringValue(value.String())
			} else {
				data.TranslatedNetworkId = types.StringNull()
			}
			if value := res.Get("interfaceInTranslatedNetwork"); value.Exists() {
				data.TranslatedNetworkIsDestinationInterface = types.BoolValue(value.Bool())
			} else {
				data.TranslatedNetworkIsDestinationInterface = types.BoolNull()
			}
			if value := res.Get("translatedPort"); value.Exists() {
				data.TranslatedPort = types.Int64Value(value.Int())
			} else {
				data.TranslatedPort = types.Int64Null()
			}
			(*parent).AutoNatRules = append((*parent).AutoNatRules, data)
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
func (data *FTDNATPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	{
		l := len(res.Get("dummy_manual_nat_rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("dummy_manual_nat_rules array resizing from %d to %d", len(data.ManualNatRules), l))
		for i := len(data.ManualNatRules); i < l; i++ {
			data.ManualNatRules = append(data.ManualNatRules, FTDNATPolicyManualNatRules{})
		}
		if len(data.ManualNatRules) > l {
			data.ManualNatRules = data.ManualNatRules[:l]
		}
	}
	for i := range data.ManualNatRules {
		parent := &data
		data := (*parent).ManualNatRules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("dummy_manual_nat_rules.%d", i))
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("metadata.section"); value.Exists() && !data.Section.IsNull() {
			data.Section = types.StringValue(value.String())
		} else {
			data.Section = types.StringNull()
		}
		if value := res.Get("natType"); value.Exists() && !data.NatType.IsNull() {
			data.NatType = types.StringValue(value.String())
		} else {
			data.NatType = types.StringNull()
		}
		if value := res.Get("fallThrough"); value.Exists() && !data.FallThrough.IsNull() {
			data.FallThrough = types.BoolValue(value.Bool())
		} else {
			data.FallThrough = types.BoolNull()
		}
		if value := res.Get("interfaceInOriginalDestination"); value.Exists() && !data.InterfaceInOriginalDestination.IsNull() {
			data.InterfaceInOriginalDestination = types.BoolValue(value.Bool())
		} else {
			data.InterfaceInOriginalDestination = types.BoolNull()
		}
		if value := res.Get("interfaceInTranslatedSource"); value.Exists() && !data.InterfaceInTranslatedSource.IsNull() {
			data.InterfaceInTranslatedSource = types.BoolValue(value.Bool())
		} else {
			data.InterfaceInTranslatedSource = types.BoolNull()
		}
		if value := res.Get("interfaceIpv6"); value.Exists() && !data.Ipv6.IsNull() {
			data.Ipv6 = types.BoolValue(value.Bool())
		} else {
			data.Ipv6 = types.BoolNull()
		}
		if value := res.Get("netToNet"); value.Exists() && !data.NetToNet.IsNull() {
			data.NetToNet = types.BoolValue(value.Bool())
		} else {
			data.NetToNet = types.BoolNull()
		}
		if value := res.Get("noProxyArp"); value.Exists() && !data.NoProxyArp.IsNull() {
			data.NoProxyArp = types.BoolValue(value.Bool())
		} else {
			data.NoProxyArp = types.BoolNull()
		}
		if value := res.Get("unidirectional"); value.Exists() && !data.Unidirectional.IsNull() {
			data.Unidirectional = types.BoolValue(value.Bool())
		} else {
			data.Unidirectional = types.BoolNull()
		}
		if value := res.Get("sourceInterface.id"); value.Exists() && !data.SourceInterfaceId.IsNull() {
			data.SourceInterfaceId = types.StringValue(value.String())
		} else {
			data.SourceInterfaceId = types.StringNull()
		}
		if value := res.Get("originalSource.id"); value.Exists() && !data.OriginalSourceId.IsNull() {
			data.OriginalSourceId = types.StringValue(value.String())
		} else {
			data.OriginalSourceId = types.StringNull()
		}
		if value := res.Get("originalSourcePort.id"); value.Exists() && !data.OriginalSourcePortId.IsNull() {
			data.OriginalSourcePortId = types.StringValue(value.String())
		} else {
			data.OriginalSourcePortId = types.StringNull()
		}
		if value := res.Get("originalDestination.id"); value.Exists() && !data.OriginalDestinationId.IsNull() {
			data.OriginalDestinationId = types.StringValue(value.String())
		} else {
			data.OriginalDestinationId = types.StringNull()
		}
		if value := res.Get("originalDestinationPort.id"); value.Exists() && !data.OriginalDestinationPortId.IsNull() {
			data.OriginalDestinationPortId = types.StringValue(value.String())
		} else {
			data.OriginalDestinationPortId = types.StringNull()
		}
		if value := res.Get("routeLookup"); value.Exists() && !data.RouteLookup.IsNull() {
			data.RouteLookup = types.BoolValue(value.Bool())
		} else {
			data.RouteLookup = types.BoolNull()
		}
		if value := res.Get("destinationInterface.id"); value.Exists() && !data.DestinationInterfaceId.IsNull() {
			data.DestinationInterfaceId = types.StringValue(value.String())
		} else {
			data.DestinationInterfaceId = types.StringNull()
		}
		if value := res.Get("translatedSource.id"); value.Exists() && !data.TranslatedSourceId.IsNull() {
			data.TranslatedSourceId = types.StringValue(value.String())
		} else {
			data.TranslatedSourceId = types.StringNull()
		}
		if value := res.Get("translatedSourcePort.id"); value.Exists() && !data.TranslatedSourcePortId.IsNull() {
			data.TranslatedSourcePortId = types.StringValue(value.String())
		} else {
			data.TranslatedSourcePortId = types.StringNull()
		}
		if value := res.Get("dns"); value.Exists() && !data.TranslateDns.IsNull() {
			data.TranslateDns = types.BoolValue(value.Bool())
		} else {
			data.TranslateDns = types.BoolNull()
		}
		if value := res.Get("translatedDestination.id"); value.Exists() && !data.TranslatedDestinationId.IsNull() {
			data.TranslatedDestinationId = types.StringValue(value.String())
		} else {
			data.TranslatedDestinationId = types.StringNull()
		}
		if value := res.Get("translatedDestinationPort.id"); value.Exists() && !data.TranslatedDestinationPortId.IsNull() {
			data.TranslatedDestinationPortId = types.StringValue(value.String())
		} else {
			data.TranslatedDestinationPortId = types.StringNull()
		}
		(*parent).ManualNatRules[i] = data
	}
	for i := 0; i < len(data.AutoNatRules); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.AutoNatRules[i].Id.ValueString()}

		parent := &data
		data := (*parent).AutoNatRules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dummy_auto_nat_rules").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AutoNatRules[%d] = %+v",
				i,
				(*parent).AutoNatRules[i],
			))
			(*parent).AutoNatRules = slices.Delete((*parent).AutoNatRules, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("natType"); value.Exists() && !data.NatType.IsNull() {
			data.NatType = types.StringValue(value.String())
		} else {
			data.NatType = types.StringNull()
		}
		if value := res.Get("destinationInterface.id"); value.Exists() && !data.DestinationInterfaceId.IsNull() {
			data.DestinationInterfaceId = types.StringValue(value.String())
		} else {
			data.DestinationInterfaceId = types.StringNull()
		}
		if value := res.Get("fallThrough"); value.Exists() && !data.FallThrough.IsNull() {
			data.FallThrough = types.BoolValue(value.Bool())
		} else {
			data.FallThrough = types.BoolNull()
		}
		if value := res.Get("interfaceIpv6"); value.Exists() && !data.Ipv6.IsNull() {
			data.Ipv6 = types.BoolValue(value.Bool())
		} else {
			data.Ipv6 = types.BoolNull()
		}
		if value := res.Get("netToNet"); value.Exists() && !data.NetToNet.IsNull() {
			data.NetToNet = types.BoolValue(value.Bool())
		} else {
			data.NetToNet = types.BoolNull()
		}
		if value := res.Get("noProxyArp"); value.Exists() && !data.NoProxyArp.IsNull() {
			data.NoProxyArp = types.BoolValue(value.Bool())
		} else {
			data.NoProxyArp = types.BoolNull()
		}
		if value := res.Get("originalNetwork.id"); value.Exists() && !data.OriginalNetworkId.IsNull() {
			data.OriginalNetworkId = types.StringValue(value.String())
		} else {
			data.OriginalNetworkId = types.StringNull()
		}
		if value := res.Get("originalPort"); value.Exists() && !data.OriginalPort.IsNull() {
			data.OriginalPort = types.Int64Value(value.Int())
		} else {
			data.OriginalPort = types.Int64Null()
		}
		if value := res.Get("serviceProtocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("routeLookup"); value.Exists() && !data.PerformRouteLookup.IsNull() {
			data.PerformRouteLookup = types.BoolValue(value.Bool())
		} else {
			data.PerformRouteLookup = types.BoolNull()
		}
		if value := res.Get("sourceInterface.id"); value.Exists() && !data.SourceInterfaceId.IsNull() {
			data.SourceInterfaceId = types.StringValue(value.String())
		} else {
			data.SourceInterfaceId = types.StringNull()
		}
		if value := res.Get("dns"); value.Exists() && !data.TranslateDns.IsNull() {
			data.TranslateDns = types.BoolValue(value.Bool())
		} else {
			data.TranslateDns = types.BoolNull()
		}
		if value := res.Get("translatedNetwork.id"); value.Exists() && !data.TranslatedNetworkId.IsNull() {
			data.TranslatedNetworkId = types.StringValue(value.String())
		} else {
			data.TranslatedNetworkId = types.StringNull()
		}
		if value := res.Get("interfaceInTranslatedNetwork"); value.Exists() && !data.TranslatedNetworkIsDestinationInterface.IsNull() {
			data.TranslatedNetworkIsDestinationInterface = types.BoolValue(value.Bool())
		} else {
			data.TranslatedNetworkIsDestinationInterface = types.BoolNull()
		}
		if value := res.Get("translatedPort"); value.Exists() && !data.TranslatedPort.IsNull() {
			data.TranslatedPort = types.Int64Value(value.Int())
		} else {
			data.TranslatedPort = types.Int64Null()
		}
		(*parent).AutoNatRules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDNATPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	for i := range data.ManualNatRules {
		r := res.Get(fmt.Sprintf("dummy_manual_nat_rules.%d", i))
		if v := data.ManualNatRules[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.ManualNatRules[i] = v
		}
	}
	for i := range data.AutoNatRules {
		keys := [...]string{"id"}
		keyValues := [...]string{data.AutoNatRules[i].Id.ValueString()}

		var r gjson.Result
		res.Get("dummy_auto_nat_rules").ForEach(
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
					r = v
					return false
				}
				return true
			},
		)
		if v := data.AutoNatRules[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.AutoNatRules[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns
