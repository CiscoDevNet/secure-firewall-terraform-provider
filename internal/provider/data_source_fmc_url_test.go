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

func TestAccDataSourceFmcURL(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url.test", "name", "my_url"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url.test", "url", "https://www.example.com/app"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url.test", "description", "My URL"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_url.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcURLConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcURLConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcURLConfig() string {
	config := `resource "fmc_url" "test" {` + "\n"
	config += `	name = "my_url"` + "\n"
	config += `	url = "https://www.example.com/app"` + "\n"
	config += `	description = "My URL"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_url" "test" {
			id = fmc_url.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcURLConfig() string {
	config := `resource "fmc_url" "test" {` + "\n"
	config += `	name = "my_url"` + "\n"
	config += `	url = "https://www.example.com/app"` + "\n"
	config += `	description = "My URL"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_url" "test" {
			name = fmc_url.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
