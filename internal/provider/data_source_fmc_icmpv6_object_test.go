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

func TestAccDataSourceFmcICMPv6Object(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_icmpv6_object.test", "icmp_type", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_icmpv6_object.test", "code", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_icmpv6_object.test", "name", "my_icmpv6_object"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_icmpv6_object.test", "description", "ICMPv6 address unreachable response, type 1, code 3"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_icmpv6_object.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcICMPv6ObjectConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcICMPv6ObjectConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcICMPv6ObjectConfig() string {
	config := `resource "fmc_icmpv6_object" "test" {` + "\n"
	config += `	icmp_type = 1` + "\n"
	config += `	code = 3` + "\n"
	config += `	name = "my_icmpv6_object"` + "\n"
	config += `	description = "ICMPv6 address unreachable response, type 1, code 3"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_icmpv6_object" "test" {
			id = fmc_icmpv6_object.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcICMPv6ObjectConfig() string {
	config := `resource "fmc_icmpv6_object" "test" {` + "\n"
	config += `	icmp_type = 1` + "\n"
	config += `	code = 3` + "\n"
	config += `	name = "my_icmpv6_object"` + "\n"
	config += `	description = "ICMPv6 address unreachable response, type 1, code 3"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_icmpv6_object" "test" {
			name = fmc_icmpv6_object.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
