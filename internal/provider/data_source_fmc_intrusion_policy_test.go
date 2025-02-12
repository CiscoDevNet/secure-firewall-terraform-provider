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

func TestAccDataSourceFmcIntrusionPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_intrusion_policy.test", "name", "my_intrusion_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_intrusion_policy.test", "description", "My IPS Policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_intrusion_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_intrusion_policy.test", "inspection_mode", "PREVENTION"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcIntrusionPolicyPrerequisitesConfig + testAccDataSourceFmcIntrusionPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcIntrusionPolicyPrerequisitesConfig + testAccNamedDataSourceFmcIntrusionPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcIntrusionPolicyPrerequisitesConfig = `
data "fmc_intrusion_policy" "builtin" {
  name = "Balanced Security and Connectivity"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcIntrusionPolicyConfig() string {
	config := `resource "fmc_intrusion_policy" "test" {` + "\n"
	config += `	name = "my_intrusion_policy"` + "\n"
	config += `	description = "My IPS Policy"` + "\n"
	config += `	base_policy_id = data.fmc_intrusion_policy.builtin.id` + "\n"
	config += `	inspection_mode = "PREVENTION"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_intrusion_policy" "test" {
			id = fmc_intrusion_policy.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcIntrusionPolicyConfig() string {
	config := `resource "fmc_intrusion_policy" "test" {` + "\n"
	config += `	name = "my_intrusion_policy"` + "\n"
	config += `	description = "My IPS Policy"` + "\n"
	config += `	base_policy_id = data.fmc_intrusion_policy.builtin.id` + "\n"
	config += `	inspection_mode = "PREVENTION"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_intrusion_policy" "test" {
			name = fmc_intrusion_policy.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
