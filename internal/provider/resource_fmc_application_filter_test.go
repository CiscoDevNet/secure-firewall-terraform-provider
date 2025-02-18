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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcApplicationFilter(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_application_filter.test", "name", "my_application_filter"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_application_filter.test", "type"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcApplicationFilterPrerequisitesConfig + testAccFmcApplicationFilterConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_application_filter.test",
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

const testAccFmcApplicationFilterPrerequisitesConfig = `
data "fmc_application" "test" {
  name = "1-800-Flowers"
}

data "fmc_application_risk" "test" {
  name = "Medium"
}

data "fmc_application_type" "test" {
  name = "Webapp"
}

data "fmc_application_business_relevance" "test" {
  name = "Medium"
}

data "fmc_application_category" "test" {
  name = "business"
}

data "fmc_application_tag" "test" {
  name = "SSL protocol"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcApplicationFilterConfig_all() string {
	config := `resource "fmc_application_filter" "test" {` + "\n"
	config += `	name = "my_application_filter"` + "\n"
	config += `	applications = [{` + "\n"
	config += `		id = data.fmc_application.test.id` + "\n"
	config += `		name = data.fmc_application.test.name` + "\n"
	config += `	}]` + "\n"
	config += `	filters = [{` + "\n"
	config += `		types = [{` + "\n"
	config += `			id = data.fmc_application_type.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		risks = [{` + "\n"
	config += `			id = data.fmc_application_risk.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		business_relevances = [{` + "\n"
	config += `			id = data.fmc_application_business_relevance.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		categories = [{` + "\n"
	config += `			id = data.fmc_application_category.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		tags = [{` + "\n"
	config += `			id = data.fmc_application_tag.test.id` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
