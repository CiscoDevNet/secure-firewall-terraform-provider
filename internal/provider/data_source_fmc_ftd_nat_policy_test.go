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

func TestAccDataSourceFmcFTDNATPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_nat_policy.test", "name", "my_ftd_nat_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_nat_policy.test", "description", "My nat policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ftd_nat_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_nat_policy.test", "manual_nat_rules.0.description", "My manual nat rule 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_nat_policy.test", "manual_nat_rules.0.enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_nat_policy.test", "manual_nat_rules.0.section", "BEFORE_AUTO"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_nat_policy.test", "manual_nat_rules.0.nat_type", "STATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_nat_policy.test", "manual_nat_rules.0.fall_through", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_nat_policy.test", "auto_nat_rules.0.nat_type", "STATIC"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFTDNATPolicyPrerequisitesConfig + testAccDataSourceFmcFTDNATPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcFTDNATPolicyPrerequisitesConfig + testAccNamedDataSourceFmcFTDNATPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcFTDNATPolicyPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFTDNATPolicyConfig() string {
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

	config += `
		data "fmc_ftd_nat_policy" "test" {
			id = fmc_ftd_nat_policy.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcFTDNATPolicyConfig() string {
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

	config += `
		data "fmc_ftd_nat_policy" "test" {
			name = fmc_ftd_nat_policy.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
