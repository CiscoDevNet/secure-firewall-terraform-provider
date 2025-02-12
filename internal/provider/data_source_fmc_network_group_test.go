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

func TestAccDataSourceFmcNetworkGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_group.test", "name", "my_network_group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_group.test", "description", "My Network Group 1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_network_group.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_group.test", "literals.0.value", "10.1.1.0/24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcNetworkGroupPrerequisitesConfig + testAccDataSourceFmcNetworkGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcNetworkGroupPrerequisitesConfig + testAccNamedDataSourceFmcNetworkGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcNetworkGroupPrerequisitesConfig = `
resource "fmc_range" "test" {
  name      = "fmc_network_group_range"
  ip_range  = "2005::10-2005::12"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcNetworkGroupConfig() string {
	config := `resource "fmc_network_group" "test" {` + "\n"
	config += `	name = "my_network_group"` + "\n"
	config += `	description = "My Network Group 1"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	objects = [{` + "\n"
	config += `		id = fmc_range.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	literals = [{` + "\n"
	config += `		value = "10.1.1.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_network_group" "test" {
			id = fmc_network_group.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcNetworkGroupConfig() string {
	config := `resource "fmc_network_group" "test" {` + "\n"
	config += `	name = "my_network_group"` + "\n"
	config += `	description = "My Network Group 1"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	objects = [{` + "\n"
	config += `		id = fmc_range.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	literals = [{` + "\n"
	config += `		value = "10.1.1.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_network_group" "test" {
			name = fmc_network_group.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
