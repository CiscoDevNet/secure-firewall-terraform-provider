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

func TestAccDataSourceFmcNetworkAnalysisPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_analysis_policy.test", "name", "my_network_analysis_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_analysis_policy.test", "description", "My network analysis policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_network_analysis_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_analysis_policy.test", "inspection_mode", "PREVENTION"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcNetworkAnalysisPolicyPrerequisitesConfig + testAccDataSourceFmcNetworkAnalysisPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcNetworkAnalysisPolicyPrerequisitesConfig + testAccNamedDataSourceFmcNetworkAnalysisPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcNetworkAnalysisPolicyPrerequisitesConfig = `
data "fmc_network_analysis_policy" "builtin" {
  name = "Security Over Connectivity"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcNetworkAnalysisPolicyConfig() string {
	config := `resource "fmc_network_analysis_policy" "test" {` + "\n"
	config += `	name = "my_network_analysis_policy"` + "\n"
	config += `	description = "My network analysis policy"` + "\n"
	config += `	base_policy_id = data.fmc_network_analysis_policy.builtin.id` + "\n"
	config += `	inspection_mode = "PREVENTION"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_network_analysis_policy" "test" {
			id = fmc_network_analysis_policy.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcNetworkAnalysisPolicyConfig() string {
	config := `resource "fmc_network_analysis_policy" "test" {` + "\n"
	config += `	name = "my_network_analysis_policy"` + "\n"
	config += `	description = "My network analysis policy"` + "\n"
	config += `	base_policy_id = data.fmc_network_analysis_policy.builtin.id` + "\n"
	config += `	inspection_mode = "PREVENTION"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_network_analysis_policy" "test" {
			name = fmc_network_analysis_policy.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
