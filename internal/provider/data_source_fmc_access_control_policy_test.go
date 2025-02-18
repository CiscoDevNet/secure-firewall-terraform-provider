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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcAccessControlPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "name", "my_access_control_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "description", "My Access Control Policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action", "BLOCK"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_log_end", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_syslog_severity", "DEBUG"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "categories.0.name", "category_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.action", "ALLOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.name", "rule_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.source_network_literals.0.value", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.destination_network_literals.0.value", "10.2.2.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.vlan_tag_literals.0.start_tag", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.vlan_tag_literals.0.end_tag", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.source_port_literals.0.type", "PortLiteral"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.source_port_literals.0.port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.source_port_literals.0.protocol", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.destination_port_literals.0.type", "PortLiteral"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.destination_port_literals.0.port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.destination_port_literals.0.protocol", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.log_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.application_filters.0.types.0.id", "WEBAPP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.application_filters.0.risks.0.id", "VERY_LOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.application_filters.0.business_relevances.0.id", "LOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.application_filters.0.categories.0.id", "118"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.application_filters.0.tags.0.id", "24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcAccessControlPolicyPrerequisitesConfig + testAccDataSourceFmcAccessControlPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcAccessControlPolicyPrerequisitesConfig + testAccNamedDataSourceFmcAccessControlPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcAccessControlPolicyPrerequisitesConfig = `
resource "fmc_network" "test" {
  name   = "fmc_access_control_policy_network"
  prefix = "10.0.0.0/24"
}

resource "fmc_host" "test" {
  name = "fmc_access_control_policy_host"
  ip   = "10.1.1.1"
}

resource "fmc_port" "test" {
  name     = "fmc_access_control_policy_port"
  protocol = "UDP"
  port     = "53"
}

resource "fmc_vlan_tag" "test" {
  name      = "fmc_access_control_policy_vlan_tag"
  start_tag = "10"
  end_tag   = "11"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcAccessControlPolicyConfig() string {
	config := `resource "fmc_access_control_policy" "test" {` + "\n"
	config += `	name = "my_access_control_policy"` + "\n"
	config += `	description = "My Access Control Policy"` + "\n"
	config += `	default_action = "BLOCK"` + "\n"
	config += `	default_action_log_begin = true` + "\n"
	config += `	default_action_log_end = false` + "\n"
	config += `	default_action_send_events_to_fmc = true` + "\n"
	config += `	default_action_syslog_severity = "DEBUG"` + "\n"
	config += `	categories = [{` + "\n"
	config += `		name = "category_1"` + "\n"
	config += `	}]` + "\n"
	config += `	rules = [{` + "\n"
	config += `		action = "ALLOW"` + "\n"
	config += `		name = "rule_1"` + "\n"
	config += `		category_name = "category_1"` + "\n"
	config += `		enabled = true` + "\n"
	config += `		source_network_literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_literals = [{` + "\n"
	config += `			value = "10.2.2.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `		vlan_tag_literals = [{` + "\n"
	config += `			start_tag = "11"` + "\n"
	config += `			end_tag = "22"` + "\n"
	config += `		}]` + "\n"
	config += `		vlan_tag_objects = [{` + "\n"
	config += `			id = fmc_vlan_tag.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_network_objects = [{` + "\n"
	config += `			id = fmc_network.test.id` + "\n"
	config += `			type = fmc_network.test.type` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_objects = [{` + "\n"
	config += `			id = fmc_host.test.id` + "\n"
	config += `			type = fmc_host.test.type` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_literals = [{` + "\n"
	config += `			type = "PortLiteral"` + "\n"
	config += `			port = "80"` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_literals = [{` + "\n"
	config += `			type = "PortLiteral"` + "\n"
	config += `			port = "80"` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `		}]` + "\n"
	config += `		log_begin = true` + "\n"
	config += `		log_end = true` + "\n"
	config += `		log_files = false` + "\n"
	config += `		send_events_to_fmc = true` + "\n"
	config += `		description = ""` + "\n"
	config += `		application_filters = [{` + "\n"
	config += `			types = [{` + "\n"
	config += `				id = "WEBAPP"` + "\n"
	config += `			}]` + "\n"
	config += `			risks = [{` + "\n"
	config += `				id = "VERY_LOW"` + "\n"
	config += `			}]` + "\n"
	config += `			business_relevances = [{` + "\n"
	config += `				id = "LOW"` + "\n"
	config += `			}]` + "\n"
	config += `			categories = [{` + "\n"
	config += `				id = "118"` + "\n"
	config += `			}]` + "\n"
	config += `			tags = [{` + "\n"
	config += `				id = "24"` + "\n"
	config += `			}]` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_access_control_policy" "test" {
			id = fmc_access_control_policy.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcAccessControlPolicyConfig() string {
	config := `resource "fmc_access_control_policy" "test" {` + "\n"
	config += `	name = "my_access_control_policy"` + "\n"
	config += `	description = "My Access Control Policy"` + "\n"
	config += `	default_action = "BLOCK"` + "\n"
	config += `	default_action_log_begin = true` + "\n"
	config += `	default_action_log_end = false` + "\n"
	config += `	default_action_send_events_to_fmc = true` + "\n"
	config += `	default_action_syslog_severity = "DEBUG"` + "\n"
	config += `	categories = [{` + "\n"
	config += `		name = "category_1"` + "\n"
	config += `	}]` + "\n"
	config += `	rules = [{` + "\n"
	config += `		action = "ALLOW"` + "\n"
	config += `		name = "rule_1"` + "\n"
	config += `		category_name = "category_1"` + "\n"
	config += `		enabled = true` + "\n"
	config += `		source_network_literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_literals = [{` + "\n"
	config += `			value = "10.2.2.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `		vlan_tag_literals = [{` + "\n"
	config += `			start_tag = "11"` + "\n"
	config += `			end_tag = "22"` + "\n"
	config += `		}]` + "\n"
	config += `		vlan_tag_objects = [{` + "\n"
	config += `			id = fmc_vlan_tag.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_network_objects = [{` + "\n"
	config += `			id = fmc_network.test.id` + "\n"
	config += `			type = fmc_network.test.type` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_objects = [{` + "\n"
	config += `			id = fmc_host.test.id` + "\n"
	config += `			type = fmc_host.test.type` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_literals = [{` + "\n"
	config += `			type = "PortLiteral"` + "\n"
	config += `			port = "80"` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_literals = [{` + "\n"
	config += `			type = "PortLiteral"` + "\n"
	config += `			port = "80"` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `		}]` + "\n"
	config += `		log_begin = true` + "\n"
	config += `		log_end = true` + "\n"
	config += `		log_files = false` + "\n"
	config += `		send_events_to_fmc = true` + "\n"
	config += `		description = ""` + "\n"
	config += `		application_filters = [{` + "\n"
	config += `			types = [{` + "\n"
	config += `				id = "WEBAPP"` + "\n"
	config += `			}]` + "\n"
	config += `			risks = [{` + "\n"
	config += `				id = "VERY_LOW"` + "\n"
	config += `			}]` + "\n"
	config += `			business_relevances = [{` + "\n"
	config += `				id = "LOW"` + "\n"
	config += `			}]` + "\n"
	config += `			categories = [{` + "\n"
	config += `				id = "118"` + "\n"
	config += `			}]` + "\n"
	config += `			tags = [{` + "\n"
	config += `				id = "24"` + "\n"
	config += `			}]` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_access_control_policy" "test" {
			name = fmc_access_control_policy.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
