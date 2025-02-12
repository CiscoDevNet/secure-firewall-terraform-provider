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

func TestAccDataSourceFmcExtendedACL(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "name", "my_extended_acl"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "description", "My Extended Access Control List"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_extended_acl.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.action", "DENY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.log_level", "WARNING"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.logging", "PER_ACCESS_LIST_ENTRY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.log_interval_seconds", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.source_network_literals.0.value", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.source_network_literals.0.type", "Network"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.destination_network_literals.0.value", "10.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.destination_network_literals.0.type", "Host"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.destination_port_literals.0.type", "PortLiteral"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.destination_port_literals.0.port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.destination_port_literals.0.protocol", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.source_port_literals.0.type", "PortLiteral"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.source_port_literals.0.port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_extended_acl.test", "entries.0.source_port_literals.0.protocol", "6"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcExtendedACLPrerequisitesConfig + testAccDataSourceFmcExtendedACLConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcExtendedACLPrerequisitesConfig + testAccNamedDataSourceFmcExtendedACLConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcExtendedACLPrerequisitesConfig = `
resource "fmc_network" "test" {
  name   = "fmc_extended_acl_network"
  prefix = "10.0.0.0/24"
}

resource "fmc_host" "test" {
  name = "fmc_extended_acl_host"
  ip   = "10.1.1.1"
}

resource "fmc_port" "test" {
  name = "fmc_extended_acl_port"
  protocol = "TCP"
  port = "65000"
}

resource "fmc_sgt" "test" {
  name = "fmc_extended_acl_sgt"
  tag = "11"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcExtendedACLConfig() string {
	config := `resource "fmc_extended_acl" "test" {` + "\n"
	config += `	name = "my_extended_acl"` + "\n"
	config += `	description = "My Extended Access Control List"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "DENY"` + "\n"
	config += `		log_level = "WARNING"` + "\n"
	config += `		logging = "PER_ACCESS_LIST_ENTRY"` + "\n"
	config += `		log_interval_seconds = 120` + "\n"
	config += `		source_network_literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `			type = "Network"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_literals = [{` + "\n"
	config += `			value = "10.2.2.2"` + "\n"
	config += `			type = "Host"` + "\n"
	config += `		}]` + "\n"
	config += `		source_network_objects = [{` + "\n"
	config += `			id = fmc_network.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_sgt_objects = [{` + "\n"
	config += `			id = fmc_sgt.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_objects = [{` + "\n"
	config += `			id = fmc_host.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_objects = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_objects = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_literals = [{` + "\n"
	config += `			type = "PortLiteral"` + "\n"
	config += `			port = "80"` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_literals = [{` + "\n"
	config += `			type = "PortLiteral"` + "\n"
	config += `			port = "80"` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_extended_acl" "test" {
			id = fmc_extended_acl.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcExtendedACLConfig() string {
	config := `resource "fmc_extended_acl" "test" {` + "\n"
	config += `	name = "my_extended_acl"` + "\n"
	config += `	description = "My Extended Access Control List"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "DENY"` + "\n"
	config += `		log_level = "WARNING"` + "\n"
	config += `		logging = "PER_ACCESS_LIST_ENTRY"` + "\n"
	config += `		log_interval_seconds = 120` + "\n"
	config += `		source_network_literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `			type = "Network"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_literals = [{` + "\n"
	config += `			value = "10.2.2.2"` + "\n"
	config += `			type = "Host"` + "\n"
	config += `		}]` + "\n"
	config += `		source_network_objects = [{` + "\n"
	config += `			id = fmc_network.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_sgt_objects = [{` + "\n"
	config += `			id = fmc_sgt.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_objects = [{` + "\n"
	config += `			id = fmc_host.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_objects = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_objects = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_literals = [{` + "\n"
	config += `			type = "PortLiteral"` + "\n"
	config += `			port = "80"` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_literals = [{` + "\n"
	config += `			type = "PortLiteral"` + "\n"
	config += `			port = "80"` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_extended_acl" "test" {
			name = fmc_extended_acl.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
