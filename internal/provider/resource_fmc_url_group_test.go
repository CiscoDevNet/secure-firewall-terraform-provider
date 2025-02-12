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

func TestAccFmcURLGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_url_group.test", "name", "my_url_group"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_url_group.test", "description", "My URL group"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_url_group.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_url_group.test", "literals.0.url", "https://www.example.com/app"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcURLGroupPrerequisitesConfig + testAccFmcURLGroupConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcURLGroupPrerequisitesConfig + testAccFmcURLGroupConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_url_group.test",
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

const testAccFmcURLGroupPrerequisitesConfig = `
resource "fmc_url" "test" {
  name        = "fmc_url_group_url"
  url         = "https://www.example.com/app"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcURLGroupConfig_minimum() string {
	config := `resource "fmc_url_group" "test" {` + "\n"
	config += `	name = "my_url_group"` + "\n"
	config += `	urls = [{` + "\n"
	config += `		id = fmc_url.test.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcURLGroupConfig_all() string {
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
	return config
}

// End of section. //template:end testAccConfigAll
