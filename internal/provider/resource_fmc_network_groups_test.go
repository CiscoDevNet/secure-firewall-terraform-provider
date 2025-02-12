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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcNetworkGroups(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_network_groups.test", "items.my_network_groups.id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_network_groups.test", "items.my_network_groups.description", "My Network Group 1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_network_groups.test", "items.my_network_groups.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_network_groups.test", "items.my_network_groups.literals.0.value", "10.1.1.0/24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcNetworkGroupsPrerequisitesConfig + testAccFmcNetworkGroupsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcNetworkGroupsPrerequisitesConfig + testAccFmcNetworkGroupsConfig_all(),
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

const testAccFmcNetworkGroupsPrerequisitesConfig = `
resource "fmc_range" "test" {
  name      = "test_fmc_network_groups"
  ip_range  = "2005::10-2005::12"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcNetworkGroupsConfig_minimum() string {
	config := `resource "fmc_network_groups" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcNetworkGroupsConfig_all() string {
	config := `resource "fmc_network_groups" "test" {` + "\n"
	config += `	items = { "my_network_groups" = {` + "\n"
	config += `		description = "My Network Group 1"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		objects = [{` + "\n"
	config += `			id = fmc_range.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcNetworkGroups_GroupNames(t *testing.T) {
	steps := []resource.TestStep{{
		// step 1
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		network_groups = ["g2"]` + "\n" +
			`	}` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 2
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 3
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.99.0.0/16"}]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 4
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		network_groups = ["g2","g3","g4","g5"]` + "\n" +
			`	}` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`	"g3" = {` + "\n" +
			`		network_groups = ["g2"]` + "\n" +
			`	}` + "\n" +
			`	"g4" = {` + "\n" +
			`		network_groups = ["g2"]` + "\n" +
			`	}` + "\n" +
			`	"g5" = {` + "\n" +
			`		network_groups = ["g3", "g4"]` + "\n" +
			`	}` + "\n" +
			`	"g6" = {` + "\n" +
			`		network_groups = ["g2","g3","g4","g5"]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 5
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		network_groups = ["g2","g4","g5"]` + "\n" +
			`	}` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`	"g4" = {` + "\n" +
			`		network_groups = ["g2"]` + "\n" +
			`	}` + "\n" +
			`	"g5" = {` + "\n" +
			`		network_groups = ["g4"]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 6
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 7
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g2" = {` + "\n" +
			`		network_groups = ["g2"]` + "\n" +
			`	}` + "\n" +
			`}}`,
		ExpectError: regexp.MustCompile(`Cycle in network_groups`),
	}, {
		// step 8
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g2" = {` + "\n" +
			`		network_groups = ["no_such_group"]` + "\n" +
			`	}` + "\n" +
			`}}`,
		ExpectError: regexp.MustCompile(`Failed to create`),
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
