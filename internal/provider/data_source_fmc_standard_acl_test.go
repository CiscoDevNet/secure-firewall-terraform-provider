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

func TestAccDataSourceFmcStandardACL(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_acl.test", "name", "my_standard_acl"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_acl.test", "description", "My standard ACL"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_standard_acl.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_acl.test", "entries.0.action", "DENY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_acl.test", "entries.0.literals.0.value", "10.1.1.0/24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcStandardACLConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcStandardACLConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcStandardACLConfig() string {
	config := `resource "fmc_standard_acl" "test" {` + "\n"
	config += `	name = "my_standard_acl"` + "\n"
	config += `	description = "My standard ACL"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "DENY"` + "\n"
	config += `		literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_standard_acl" "test" {
			id = fmc_standard_acl.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcStandardACLConfig() string {
	config := `resource "fmc_standard_acl" "test" {` + "\n"
	config += `	name = "my_standard_acl"` + "\n"
	config += `	description = "My standard ACL"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "DENY"` + "\n"
	config += `		literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_standard_acl" "test" {
			name = fmc_standard_acl.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
