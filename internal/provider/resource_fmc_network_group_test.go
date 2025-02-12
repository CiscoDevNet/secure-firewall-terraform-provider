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

func TestAccFmcNetworkGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_network_group.test", "name", "my_network_group"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_network_group.test", "description", "My Network Group 1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_network_group.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_network_group.test", "literals.0.value", "10.1.1.0/24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcNetworkGroupPrerequisitesConfig + testAccFmcNetworkGroupConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcNetworkGroupPrerequisitesConfig + testAccFmcNetworkGroupConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_network_group.test",
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

const testAccFmcNetworkGroupPrerequisitesConfig = `
resource "fmc_range" "test" {
  name      = "fmc_network_group_range"
  ip_range  = "2005::10-2005::12"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcNetworkGroupConfig_minimum() string {
	config := `resource "fmc_network_group" "test" {` + "\n"
	config += `	name = "my_network_group"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcNetworkGroupConfig_all() string {
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
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcNetworkGroup_Extended(t *testing.T) {
	config := `resource "fmc_host" "ng_host1" {` + "\n"
	config += `  name = "ng_host1"` + "\n"
	config += `  ip   = "1.2.3.4"` + "\n"
	config += `}` + "\n"

	config += `resource "fmc_host" "ng_host2" {` + "\n"
	config += `  name = "ng_host2"` + "\n"
	config += `  ip   = "1.2.3.40"` + "\n"
	config += `}` + "\n"

	config += `resource "fmc_network" "ng_network1" {` + "\n"
	config += `  name   = "ng_network1"` + "\n"
	config += `  prefix = "10.20.0.0/24"` + "\n"
	config += `}` + "\n"

	config += `resource "fmc_range" "ng_range1" {` + "\n"
	config += `  name     = "ng_range1"` + "\n"
	config += `  ip_range = "192.168.1.0-192.168.1.10"` + "\n"
	config += `}` + "\n"

	config += `resource "fmc_fqdn_object" "ng_fqdn1" {` + "\n"
	config += `  name  = "ng_fqdn1"` + "\n"
	config += `  fqdn = "www.example.com"` + "\n"
	config += `}` + "\n"

	config += `resource "fmc_network_group" "ng1" {` + "\n"
	config += `  name = "ng1"` + "\n"
	config += `  objects = [` + "\n"
	config += `    { id = fmc_host.ng_host2.id },` + "\n"
	config += `  ]` + "\n"
	config += `  literals = [` + "\n"
	config += `    { value = "11.10.11.12" },` + "\n"
	config += `    { value = "11.0.0.0/24" }` + "\n"
	config += `  ]` + "\n"
	config += `}` + "\n"

	config += `resource "fmc_network_group" "ng2" {` + "\n"
	config += `  name = "ng2"` + "\n"
	config += `  objects = [` + "\n"
	config += `    { id = fmc_network_group.ng1.id },` + "\n"
	config += `    { id = fmc_host.ng_host1.id },` + "\n"
	config += `    { id = fmc_network.ng_network1.id },` + "\n"
	config += `    { id = fmc_range.ng_range1.id },` + "\n"
	config += `    { id = fmc_fqdn_object.ng_fqdn1.id }` + "\n"
	config += `  ]` + "\n"
	config += `  literals = [` + "\n"
	config += `    { value = "10.10.11.12" },` + "\n"
	config += `    { value = "10.0.0.0/24" }` + "\n"
	config += `  ]` + "\n"
	config += `}` + "\n"

	steps := []resource.TestStep{{
		Config: config,
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
