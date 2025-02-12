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

func TestAccFmcFQDNObjects(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_fqdn_objects.test", "items.my_fqdn_objects.id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_fqdn_objects.test", "items.my_fqdn_objects.description", "My FQDN 1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_fqdn_objects.test", "items.my_fqdn_objects.overridable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_fqdn_objects.test", "items.my_fqdn_objects.fqdn", "www.example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_fqdn_objects.test", "items.my_fqdn_objects.dns_resolution", "IPV4_AND_IPV6"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_fqdn_objects.test", "items.my_fqdn_objects.type"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcFQDNObjectsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFQDNObjectsConfig_all(),
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

func testAccFmcFQDNObjectsConfig_minimum() string {
	config := `resource "fmc_fqdn_objects" "test" {` + "\n"
	config += `	items = { "my_fqdn_objects" = {` + "\n"
	config += `		fqdn = "www.example.com"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFQDNObjectsConfig_all() string {
	config := `resource "fmc_fqdn_objects" "test" {` + "\n"
	config += `	items = { "my_fqdn_objects" = {` + "\n"
	config += `		description = "My FQDN 1"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		fqdn = "www.example.com"` + "\n"
	config += `		dns_resolution = "IPV4_AND_IPV6"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
