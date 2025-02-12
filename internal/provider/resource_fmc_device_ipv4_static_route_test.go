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

func TestAccFmcDeviceIPv4StaticRoute(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_ipv4_static_route.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ipv4_static_route.test", "gateway_host_literal", "10.0.0.1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceIPv4StaticRoutePrerequisitesConfig + testAccFmcDeviceIPv4StaticRouteConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceIPv4StaticRoutePrerequisitesConfig + testAccFmcDeviceIPv4StaticRouteConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
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

const testAccFmcDeviceIPv4StaticRoutePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

data "fmc_network" "test" {
  name = "any-ipv4"
}

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

func testAccFmcDeviceIPv4StaticRouteConfig_minimum() string {
	config := `resource "fmc_device_ipv4_static_route" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	interface_logical_name = fmc_device_physical_interface.test.logical_name` + "\n"
	config += `	interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	destination_networks = [{` + "\n"
	config += `		id = data.fmc_network.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	metric_value = 254` + "\n"
	config += `	gateway_host_literal = "10.0.0.2"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceIPv4StaticRouteConfig_all() string {
	config := `resource "fmc_device_ipv4_static_route" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	interface_logical_name = fmc_device_physical_interface.test.logical_name` + "\n"
	config += `	interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	destination_networks = [{` + "\n"
	config += `		id = data.fmc_network.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	metric_value = null` + "\n"
	config += `	gateway_host_literal = "10.0.0.1"` + "\n"
	config += `	is_tunneled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
