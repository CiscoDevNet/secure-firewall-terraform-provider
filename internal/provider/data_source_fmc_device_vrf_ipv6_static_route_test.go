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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcDeviceVRFIPv6StaticRoute(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_vrf_ipv6_static_route.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceVRFIPv6StaticRoutePrerequisitesConfig + testAccDataSourceFmcDeviceVRFIPv6StaticRouteConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceVRFIPv6StaticRoutePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

data "fmc_host" "test" {
  name = "any-ipv6"
}

resource "fmc_device_physical_interface" "test" {
  device_id    = var.device_id
  name         = "GigabitEthernet0/1"
  logical_name = "myinterface-0-1"
  mode         = "NONE"
  enabled      = true
}

resource "fmc_device_vrf" "test" {
  name      = "vrf_static_route_test"
  device_id = var.device_id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceVRFIPv6StaticRouteConfig() string {
	config := `resource "fmc_device_vrf_ipv6_static_route" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	vrf_id = fmc_device_vrf.test.id` + "\n"
	config += `	interface_logical_name = fmc_device_physical_interface.test.logical_name` + "\n"
	config += `	interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	destination_networks = [{` + "\n"
	config += `		id = data.fmc_host.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	metric_value = null` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_vrf_ipv6_static_route" "test" {
			id = fmc_device_vrf_ipv6_static_route.test.id
			device_id = fmc_device_physical_interface.test.device_id
			vrf_id = fmc_device_vrf.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
