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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcFTDNATPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_nat_policy.test", "name", "my_ftd_nat_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_nat_policy.test", "description", "My nat policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ftd_nat_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_nat_policy.test", "manual_nat_rules.0.description", "My manual nat rule 1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_nat_policy.test", "manual_nat_rules.0.enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_nat_policy.test", "manual_nat_rules.0.section", "BEFORE_AUTO"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_nat_policy.test", "manual_nat_rules.0.nat_type", "STATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_nat_policy.test", "manual_nat_rules.0.fall_through", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_nat_policy.test", "auto_nat_rules.0.nat_type", "STATIC"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcFTDNATPolicyPrerequisitesConfig + testAccFmcFTDNATPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFTDNATPolicyPrerequisitesConfig + testAccFmcFTDNATPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_ftd_nat_policy.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcFTDNATPolicyPrerequisitesConfig = `
resource "fmc_hosts" "test" {
  items = {
    "nat_host1" ={
      ip = "10.0.0.1"
    },
    "nat_host2" ={
      ip = "10.0.0.2"
    },
    "nat_host3" ={
      ip = "10.0.0.3"
    },
    "nat_host4" ={
      ip = "10.0.0.4"
    }
  }
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcFTDNATPolicyConfig_minimum() string {
	config := `resource "fmc_ftd_nat_policy" "test" {` + "\n"
	config += `	name = "my_ftd_nat_policy"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFTDNATPolicyConfig_all() string {
	config := `resource "fmc_ftd_nat_policy" "test" {` + "\n"
	config += `	name = "my_ftd_nat_policy"` + "\n"
	config += `	description = "My nat policy"` + "\n"
	config += `	manual_nat_rules = [{` + "\n"
	config += `		description = "My manual nat rule 1"` + "\n"
	config += `		enabled = true` + "\n"
	config += `		section = "BEFORE_AUTO"` + "\n"
	config += `		nat_type = "STATIC"` + "\n"
	config += `		fall_through = false` + "\n"
	config += `		original_source_id = fmc_hosts.test.items.nat_host1.id` + "\n"
	config += `		translated_source_id = fmc_hosts.test.items.nat_host2.id` + "\n"
	config += `	}]` + "\n"
	config += `	auto_nat_rules = [{` + "\n"
	config += `		nat_type = "STATIC"` + "\n"
	config += `		original_network_id = fmc_hosts.test.items.nat_host3.id` + "\n"
	config += `		translated_network_id = fmc_hosts.test.items.nat_host4.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcFTDNATPolicy_Sequential(t *testing.T) {
	prereq := `resource "fmc_hosts" "test" {` + "\n"
	prereq += `  items = {` + "\n"
	prereq += `    "nat_seq_host1"  = { ip = "10.1.0.1" },` + "\n"
	prereq += `    "nat_seq_host2"  = { ip = "10.1.0.2" },` + "\n"
	prereq += `    "nat_seq_host3"  = { ip = "10.1.0.3" },` + "\n"
	prereq += `    "nat_seq_host4"  = { ip = "10.1.0.4" },` + "\n"
	prereq += `    "nat_seq_host5"  = { ip = "10.1.0.5" },` + "\n"
	prereq += `    "nat_seq_host6"  = { ip = "10.1.0.6" },` + "\n"
	prereq += `    "nat_seq_host7"  = { ip = "10.1.0.7" },` + "\n"
	prereq += `    "nat_seq_host8"  = { ip = "10.1.0.8" },` + "\n"
	prereq += `    "nat_seq_host9"  = { ip = "10.1.0.9" },` + "\n"
	prereq += `    "nat_seq_host10" = { ip = "10.1.0.10" },` + "\n"
	prereq += `    "nat_seq_host11" = { ip = "10.1.0.11" },` + "\n"
	prereq += `    "nat_seq_host12" = { ip = "10.1.0.12" }` + "\n"
	prereq += `  }` + "\n"
	prereq += `}` + "\n"
	prereq += `` + "\n"
	prereq += `resource "fmc_ports" "test" {` + "\n"
	prereq += `  items = {` + "\n"
	prereq += `    "nat_seq_port1" = { port = "80", protocol = "TCP" },` + "\n"
	prereq += `    "nat_seq_port2" = { port = "8080", protocol = "TCP" },` + "\n"
	prereq += `  }` + "\n"
	prereq += `}` + "\n\n"

	// Start policy
	step_01 := `resource "fmc_ftd_nat_policy" "test" {` + "\n"
	step_01 += `  name        = "nat_seq_ftd_nat_policy"` + "\n"
	step_01 += `  description = "nat_seq_ftd_nat_policy"` + "\n"
	step_01 += `  manual_nat_rules = [` + "\n"
	step_01 += `    {` + "\n"
	step_01 += `      description                    = "My rule 1",` + "\n"
	step_01 += `      enabled                        = false,` + "\n"
	step_01 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_01 += `      nat_type                       = "STATIC",` + "\n"
	step_01 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host1.id,` + "\n"
	step_01 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host2.id,` + "\n"
	step_01 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_01 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_01 += `    },` + "\n"
	step_01 += `    {` + "\n"
	step_01 += `      description                    = "My rule 2",` + "\n"
	step_01 += `      enabled                        = false,` + "\n"
	step_01 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_01 += `      nat_type                       = "STATIC",` + "\n"
	step_01 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host3.id,` + "\n"
	step_01 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host4.id,` + "\n"
	step_01 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_01 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_01 += `    },` + "\n"
	step_01 += `    {` + "\n"
	step_01 += `      description                    = "My rule 3",` + "\n"
	step_01 += `      enabled                        = false,` + "\n"
	step_01 += `      section                        = "AFTER_AUTO",` + "\n"
	step_01 += `      nat_type                       = "STATIC",` + "\n"
	step_01 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host5.id,` + "\n"
	step_01 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host6.id,` + "\n"
	step_01 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_01 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_01 += `    },` + "\n"
	step_01 += `    {` + "\n"
	step_01 += `      description                    = "My rule 4",` + "\n"
	step_01 += `      enabled                        = false,` + "\n"
	step_01 += `      section                        = "AFTER_AUTO",` + "\n"
	step_01 += `      nat_type                       = "STATIC",` + "\n"
	step_01 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host7.id,` + "\n"
	step_01 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host8.id,` + "\n"
	step_01 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_01 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_01 += `    }` + "\n"
	step_01 += `  ]` + "\n"
	step_01 += `  auto_nat_rules = [` + "\n"
	step_01 += `    {` + "\n"
	step_01 += `      nat_type              = "STATIC",` + "\n"
	step_01 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host9.id,` + "\n"
	step_01 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host10.id,` + "\n"
	step_01 += `      original_port         = 8022,` + "\n"
	step_01 += `      translated_port       = 22` + "\n"
	step_01 += `    },` + "\n"
	step_01 += `    {` + "\n"
	step_01 += `      nat_type              = "STATIC",` + "\n"
	step_01 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host11.id,` + "\n"
	step_01 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host12.id,` + "\n"
	step_01 += `      original_port         = 8022,` + "\n"
	step_01 += `      translated_port       = 22` + "\n"
	step_01 += `    }` + "\n"
	step_01 += `  ]` + "\n"
	step_01 += `}` + "\n"

	// Add rule in between existing rules
	step_02 := `resource "fmc_ftd_nat_policy" "test" {` + "\n"
	step_02 += `  name        = "nat_seq_ftd_nat_policy"` + "\n"
	step_02 += `  description = "nat_seq_ftd_nat_policy"` + "\n"
	step_02 += `  manual_nat_rules = [` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      description                    = "My rule 1",` + "\n"
	step_02 += `      enabled                        = false,` + "\n"
	step_02 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_02 += `      nat_type                       = "STATIC",` + "\n"
	step_02 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host1.id,` + "\n"
	step_02 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host2.id,` + "\n"
	step_02 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_02 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_02 += `    },` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      description                    = "My rule 1.5",` + "\n"
	step_02 += `      enabled                        = false,` + "\n"
	step_02 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_02 += `      nat_type                       = "STATIC",` + "\n"
	step_02 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host2.id,` + "\n"
	step_02 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host3.id,` + "\n"
	step_02 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_02 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_02 += `    },` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      description                    = "My rule 2",` + "\n"
	step_02 += `      enabled                        = true,` + "\n"
	step_02 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_02 += `      nat_type                       = "STATIC",` + "\n"
	step_02 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host3.id,` + "\n"
	step_02 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host4.id,` + "\n"
	step_02 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_02 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_02 += `    },` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      description                    = "My rule 3",` + "\n"
	step_02 += `      enabled                        = false,` + "\n"
	step_02 += `      section                        = "AFTER_AUTO",` + "\n"
	step_02 += `      nat_type                       = "STATIC",` + "\n"
	step_02 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host5.id,` + "\n"
	step_02 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host6.id,` + "\n"
	step_02 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_02 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_02 += `    },` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      description                    = "My rule 3.5",` + "\n"
	step_02 += `      enabled                        = false,` + "\n"
	step_02 += `      section                        = "AFTER_AUTO",` + "\n"
	step_02 += `      nat_type                       = "STATIC",` + "\n"
	step_02 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host6.id,` + "\n"
	step_02 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host7.id,` + "\n"
	step_02 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_02 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_02 += `    },` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      description                    = "My rule 4",` + "\n"
	step_02 += `      enabled                        = true,` + "\n"
	step_02 += `      section                        = "AFTER_AUTO",` + "\n"
	step_02 += `      nat_type                       = "STATIC",` + "\n"
	step_02 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host7.id,` + "\n"
	step_02 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host8.id,` + "\n"
	step_02 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_02 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_02 += `    }` + "\n"
	step_02 += `  ]` + "\n"
	step_02 += `  auto_nat_rules = [` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      nat_type              = "STATIC",` + "\n"
	step_02 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host9.id,` + "\n"
	step_02 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host10.id,` + "\n"
	step_02 += `      original_port         = 8022,` + "\n"
	step_02 += `      translated_port       = 22` + "\n"
	step_02 += `    },` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      nat_type              = "STATIC",` + "\n"
	step_02 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host10.id,` + "\n"
	step_02 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host11.id,` + "\n"
	step_02 += `      original_port         = 8022,` + "\n"
	step_02 += `      translated_port       = 22` + "\n"
	step_02 += `    },` + "\n"
	step_02 += `    {` + "\n"
	step_02 += `      nat_type              = "STATIC",` + "\n"
	step_02 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host11.id,` + "\n"
	step_02 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host12.id,` + "\n"
	step_02 += `      original_port         = 8022,` + "\n"
	step_02 += `      translated_port       = 22` + "\n"
	step_02 += `    }` + "\n"
	step_02 += `  ]` + "\n"
	step_02 += `}` + "\n"

	// Remove rule in between existing rules
	step_03 := `resource "fmc_ftd_nat_policy" "test" {` + "\n"
	step_03 += `  name        = "nat_seq_ftd_nat_policy"` + "\n"
	step_03 += `  description = "nat_seq_ftd_nat_policy"` + "\n"
	step_03 += `  manual_nat_rules = [` + "\n"
	step_03 += `    {` + "\n"
	step_03 += `      description                    = "My rule 1",` + "\n"
	step_03 += `      enabled                        = false,` + "\n"
	step_03 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_03 += `      nat_type                       = "STATIC",` + "\n"
	step_03 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host1.id,` + "\n"
	step_03 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host2.id,` + "\n"
	step_03 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_03 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_03 += `    },` + "\n"
	step_03 += `    {` + "\n"
	step_03 += `      description                    = "My rule 2",` + "\n"
	step_03 += `      enabled                        = true,` + "\n"
	step_03 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_03 += `      nat_type                       = "STATIC",` + "\n"
	step_03 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host3.id,` + "\n"
	step_03 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host4.id,` + "\n"
	step_03 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_03 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_03 += `    },` + "\n"
	step_03 += `    {` + "\n"
	step_03 += `      description                    = "My rule 3",` + "\n"
	step_03 += `      enabled                        = false,` + "\n"
	step_03 += `      section                        = "AFTER_AUTO",` + "\n"
	step_03 += `      nat_type                       = "STATIC",` + "\n"
	step_03 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host5.id,` + "\n"
	step_03 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host6.id,` + "\n"
	step_03 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_03 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_03 += `    },` + "\n"
	step_03 += `    {` + "\n"
	step_03 += `      description                    = "My rule 4",` + "\n"
	step_03 += `      enabled                        = true,` + "\n"
	step_03 += `      section                        = "AFTER_AUTO",` + "\n"
	step_03 += `      nat_type                       = "STATIC",` + "\n"
	step_03 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host7.id,` + "\n"
	step_03 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host8.id,` + "\n"
	step_03 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_03 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_03 += `    }` + "\n"
	step_03 += `  ]` + "\n"
	step_03 += `  auto_nat_rules = [` + "\n"
	step_03 += `    {` + "\n"
	step_03 += `      nat_type              = "STATIC",` + "\n"
	step_03 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host9.id,` + "\n"
	step_03 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host10.id,` + "\n"
	step_03 += `      original_port         = 8022,` + "\n"
	step_03 += `      translated_port       = 22` + "\n"
	step_03 += `    },` + "\n"
	step_03 += `    {` + "\n"
	step_03 += `      nat_type              = "STATIC",` + "\n"
	step_03 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host11.id,` + "\n"
	step_03 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host12.id,` + "\n"
	step_03 += `      original_port         = 8022,` + "\n"
	step_03 += `      translated_port       = 22` + "\n"
	step_03 += `    }` + "\n"
	step_03 += `  ]` + "\n"
	step_03 += `}` + "\n"

	// Change existing rules
	step_04 := `resource "fmc_ftd_nat_policy" "test" {` + "\n"
	step_04 += `  name        = "nat_seq_ftd_nat_policy"` + "\n"
	step_04 += `  description = "nat_seq_ftd_nat_policy"` + "\n"
	step_04 += `  manual_nat_rules = [` + "\n"
	step_04 += `    {` + "\n"
	step_04 += `      description                    = "My rule 1",` + "\n"
	step_04 += `      enabled                        = false,` + "\n"
	step_04 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_04 += `      nat_type                       = "STATIC",` + "\n"
	step_04 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host2.id,` + "\n"
	step_04 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host1.id,` + "\n"
	step_04 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_04 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_04 += `    },` + "\n"
	step_04 += `    {` + "\n"
	step_04 += `      description                    = "My rule 2",` + "\n"
	step_04 += `      enabled                        = true,` + "\n"
	step_04 += `      section                        = "BEFORE_AUTO",` + "\n"
	step_04 += `      nat_type                       = "STATIC",` + "\n"
	step_04 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host4.id,` + "\n"
	step_04 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host3.id,` + "\n"
	step_04 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_04 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_04 += `    },` + "\n"
	step_04 += `    {` + "\n"
	step_04 += `      description                    = "My rule 3",` + "\n"
	step_04 += `      enabled                        = false,` + "\n"
	step_04 += `      section                        = "AFTER_AUTO",` + "\n"
	step_04 += `      nat_type                       = "STATIC",` + "\n"
	step_04 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host6.id,` + "\n"
	step_04 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host5.id,` + "\n"
	step_04 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_04 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_04 += `    },` + "\n"
	step_04 += `    {` + "\n"
	step_04 += `      description                    = "My rule 4",` + "\n"
	step_04 += `      enabled                        = true,` + "\n"
	step_04 += `      section                        = "AFTER_AUTO",` + "\n"
	step_04 += `      nat_type                       = "STATIC",` + "\n"
	step_04 += `      original_source_id             = fmc_hosts.test.items.nat_seq_host8.id,` + "\n"
	step_04 += `      translated_source_id           = fmc_hosts.test.items.nat_seq_host7.id,` + "\n"
	step_04 += `      original_destination_port_id   = fmc_ports.test.items.nat_seq_port1.id,` + "\n"
	step_04 += `      translated_destination_port_id = fmc_ports.test.items.nat_seq_port2.id,` + "\n"
	step_04 += `    }` + "\n"
	step_04 += `  ]` + "\n"
	step_04 += `  auto_nat_rules = [` + "\n"
	step_04 += `    {` + "\n"
	step_04 += `      nat_type              = "STATIC",` + "\n"
	step_04 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host10.id,` + "\n"
	step_04 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host9.id,` + "\n"
	step_04 += `      original_port         = 8022,` + "\n"
	step_04 += `      translated_port       = 22` + "\n"
	step_04 += `    },` + "\n"
	step_04 += `    {` + "\n"
	step_04 += `      nat_type              = "STATIC",` + "\n"
	step_04 += `      original_network_id   = fmc_hosts.test.items.nat_seq_host12.id,` + "\n"
	step_04 += `      translated_network_id = fmc_hosts.test.items.nat_seq_host11.id,` + "\n"
	step_04 += `      original_port         = 8022,` + "\n"
	step_04 += `      translated_port       = 22` + "\n"
	step_04 += `    }` + "\n"
	step_04 += `  ]` + "\n"
	step_04 += `}` + "\n"

	steps := []resource.TestStep{{
		Config: prereq + step_01,
	}, {
		Config: prereq + step_02,
	}, {
		Config: prereq + step_03,
	}, {
		Config: prereq + step_04,
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
