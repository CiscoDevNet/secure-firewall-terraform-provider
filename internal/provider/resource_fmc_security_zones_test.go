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

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcSecurityZones(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_security_zones.test", "items.my_security_zones.id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_security_zones.test", "items.my_security_zones.interface_type", "ROUTED"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_security_zones.test", "items.my_security_zones.type"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcSecurityZonesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcSecurityZonesConfig_all(),
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcSecurityZonesConfig_minimum() string {
	config := `resource "fmc_security_zones" "test" {` + "\n"
	config += `	items = { "my_security_zones" = {` + "\n"
	config += `		interface_type = "ROUTED"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcSecurityZonesConfig_all() string {
	config := `resource "fmc_security_zones" "test" {` + "\n"
	config += `	items = { "my_security_zones" = {` + "\n"
	config += `		interface_type = "ROUTED"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcSecurityZones_Sequential(t *testing.T) {

	step_01 := `resource "fmc_security_zones" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"my_security_zones_1" = {` + "\n" +
		`			interface_type = "ROUTED",` + "\n" +
		`		},` + "\n" +
		`		"my_security_zones_2" = {` + "\n" +
		`			interface_type = "ROUTED",` + "\n" +
		`		},` + "\n" +
		`		"my_security_zones_3" = {` + "\n" +
		`			interface_type = "ROUTED",` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_security_zones.test", "items.my_security_zones_1.id"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_security_zones.test", "items.my_security_zones_1.interface_type", "ROUTED"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_security_zones.test", "items.my_security_zones_2.id"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_security_zones.test", "items.my_security_zones_2.interface_type", "ROUTED"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_security_zones.test", "items.my_security_zones_3.id"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_security_zones.test", "items.my_security_zones_3.interface_type", "ROUTED"))

	step_02 := `resource "fmc_security_zones" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"my_security_zones_1" = {` + "\n" +
		`			interface_type = "ROUTED",` + "\n" +
		`		},` + "\n" +
		`		"my_security_zones_2" = {` + "\n" +
		`			interface_type = "INLINE",` + "\n" +
		`		},` + "\n" +
		`		"my_security_zones_4" = {` + "\n" +
		`			interface_type = "ROUTED",` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step02 []resource.TestCheckFunc
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_security_zones.test", "items.my_security_zones_1.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_security_zones.test", "items.my_security_zones_1.interface_type", "ROUTED"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_security_zones.test", "items.my_security_zones_2.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_security_zones.test", "items.my_security_zones_2.interface_type", "INLINE"))
	checks_step02 = append(checks_step02, resource.TestCheckNoResourceAttr("fmc_security_zones.test", "items.my_security_zones_3.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_security_zones.test", "items.my_security_zones_4.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_security_zones.test", "items.my_security_zones_4.interface_type", "ROUTED"))

	// Zone 1 should not change
	zone1Id := statecheck.CompareValue(compare.ValuesSame())
	// Zone 2 should be recreated, due to the change in interface_type
	zone2Id := statecheck.CompareValue(compare.ValuesDiffer())

	steps := []resource.TestStep{{
		Config: step_01,
		Check:  resource.ComposeTestCheckFunc(checks_step01...),
		ConfigStateChecks: []statecheck.StateCheck{
			zone1Id.AddStateValue("fmc_security_zones.test", tfjsonpath.New("items").AtMapKey("my_security_zones_1").AtMapKey("id")),
			zone2Id.AddStateValue("fmc_security_zones.test", tfjsonpath.New("items").AtMapKey("my_security_zones_2").AtMapKey("id")),
		},
	}, {
		Config: step_02,
		Check:  resource.ComposeTestCheckFunc(checks_step02...),
		ConfigStateChecks: []statecheck.StateCheck{
			zone1Id.AddStateValue("fmc_security_zones.test", tfjsonpath.New("items").AtMapKey("my_security_zones_1").AtMapKey("id")),
			zone2Id.AddStateValue("fmc_security_zones.test", tfjsonpath.New("items").AtMapKey("my_security_zones_2").AtMapKey("id")),
		},
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
