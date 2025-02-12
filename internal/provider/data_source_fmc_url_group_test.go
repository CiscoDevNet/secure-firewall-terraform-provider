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

func TestAccDataSourceFmcURLGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url_group.test", "name", "my_url_group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url_group.test", "description", "My URL group"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_url_group.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url_group.test", "literals.0.url", "https://www.example.com/app"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcURLGroupPrerequisitesConfig + testAccDataSourceFmcURLGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcURLGroupPrerequisitesConfig + testAccNamedDataSourceFmcURLGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcURLGroupPrerequisitesConfig = `
resource "fmc_url" "test" {
  name        = "fmc_url_group_url"
  url         = "https://www.example.com/app"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcURLGroupConfig() string {
	config := `resource "fmc_url_group" "test" {` + "\n"
	config += `	name = "my_url_group"` + "\n"
	config += `	description = "My URL group"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	urls = [{` + "\n"
	config += `		id = fmc_url.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	literals = [{` + "\n"
	config += `		url = "https://www.example.com/app"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_url_group" "test" {
			id = fmc_url_group.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcURLGroupConfig() string {
	config := `resource "fmc_url_group" "test" {` + "\n"
	config += `	name = "my_url_group"` + "\n"
	config += `	description = "My URL group"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	urls = [{` + "\n"
	config += `		id = fmc_url.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	literals = [{` + "\n"
	config += `		url = "https://www.example.com/app"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_url_group" "test" {
			name = fmc_url_group.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
