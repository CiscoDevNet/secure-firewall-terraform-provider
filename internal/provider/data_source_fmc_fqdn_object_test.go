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

func TestAccDataSourceFmcFQDNObject(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_fqdn_object.test", "name", "my_fqdn_object"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_fqdn_object.test", "fqdn", "www.example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_fqdn_object.test", "dns_resolution", "IPV4_AND_IPV6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_fqdn_object.test", "description", "My FQDN Object"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_fqdn_object.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFQDNObjectConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcFQDNObjectConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFQDNObjectConfig() string {
	config := `resource "fmc_fqdn_object" "test" {` + "\n"
	config += `	name = "my_fqdn_object"` + "\n"
	config += `	fqdn = "www.example.com"` + "\n"
	config += `	dns_resolution = "IPV4_AND_IPV6"` + "\n"
	config += `	description = "My FQDN Object"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_fqdn_object" "test" {
			id = fmc_fqdn_object.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcFQDNObjectConfig() string {
	config := `resource "fmc_fqdn_object" "test" {` + "\n"
	config += `	name = "my_fqdn_object"` + "\n"
	config += `	fqdn = "www.example.com"` + "\n"
	config += `	dns_resolution = "IPV4_AND_IPV6"` + "\n"
	config += `	description = "My FQDN Object"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_fqdn_object" "test" {
			name = fmc_fqdn_object.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
