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

func TestAccFmcInterfaceGroup(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_interface_group.test", "name", "my_interface_group"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_interface_group.test", "interface_mode", "ROUTED"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_interface_group.test", "type"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcInterfaceGroupPrerequisitesConfig + testAccFmcInterfaceGroupConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcInterfaceGroupPrerequisitesConfig + testAccFmcInterfaceGroupConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_interface_group.test",
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

const testAccFmcInterfaceGroupPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

resource "fmc_device_physical_interface" "test" {
  device_id    = var.device_id
  name         = "GigabitEthernet0/1"
  logical_name = "myinterface-0-1"
  mode         = "NONE"
  enabled      = true
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcInterfaceGroupConfig_minimum() string {
	config := `resource "fmc_interface_group" "test" {` + "\n"
	config += `	name = "my_interface_group"` + "\n"
	config += `	interface_mode = "ROUTED"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcInterfaceGroupConfig_all() string {
	config := `resource "fmc_interface_group" "test" {` + "\n"
	config += `	name = "my_interface_group"` + "\n"
	config += `	interface_mode = "ROUTED"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		id = fmc_device_physical_interface.test.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
