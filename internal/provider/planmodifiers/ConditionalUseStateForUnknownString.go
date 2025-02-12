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

package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

//
// ConditionalUseStateForUnknownString sets value to `known after apply` if dependent attribute changes
//

var _ planmodifier.String = conditionalUseStateForUnknownString{}

func ConditionalUseStateForUnknownString(dependentAttribute string) planmodifier.String {
	return conditionalUseStateForUnknownString{dependentAttribute: dependentAttribute}
}

type conditionalUseStateForUnknownString struct {
	dependentAttribute string
}

func (c conditionalUseStateForUnknownString) Description(ctx context.Context) string {
	return "Conditionally set to `known after apply` based on change of `dependentAttribute`"
}

func (c conditionalUseStateForUnknownString) MarkdownDescription(ctx context.Context) string {
	return c.Description(ctx)
}

func (c conditionalUseStateForUnknownString) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Get path to the dependent attribute
	dependentAttributePath := req.Path.ParentPath().AtName(c.dependentAttribute)

	// Get state value of dependent attribute
	var stateValue types.String
	req.State.GetAttribute(ctx, dependentAttributePath, &stateValue)

	// Get plan value of dependent attribute
	var planValue types.String
	req.Plan.GetAttribute(ctx, dependentAttributePath, &planValue)

	// If state and plan values don't match, do nothing (sets value to `known after apply`)
	if stateValue.ValueString() != planValue.ValueString() {
		return
	}

	// If state and plan values match, set plan to state value
	resp.PlanValue = req.StateValue
}
