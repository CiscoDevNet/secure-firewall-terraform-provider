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
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceBGPGeneralSettings struct {
	Id                               types.String `tfsdk:"id"`
	Domain                           types.String `tfsdk:"domain"`
	DeviceId                         types.String `tfsdk:"device_id"`
	Name                             types.String `tfsdk:"name"`
	AsNumber                         types.String `tfsdk:"as_number"`
	RouterId                         types.String `tfsdk:"router_id"`
	ScanningInterval                 types.Int64  `tfsdk:"scanning_interval"`
	AsNumberInPathAttribute          types.Int64  `tfsdk:"as_number_in_path_attribute"`
	LogNeighborChanges               types.Bool   `tfsdk:"log_neighbor_changes"`
	TcpPathMtuDiscovery              types.Bool   `tfsdk:"tcp_path_mtu_discovery"`
	ResetSessionUponFailover         types.Bool   `tfsdk:"reset_session_upon_failover"`
	EnforceFirstPeerAs               types.Bool   `tfsdk:"enforce_first_peer_as"`
	UseDotNotation                   types.Bool   `tfsdk:"use_dot_notation"`
	AggregateTimer                   types.Int64  `tfsdk:"aggregate_timer"`
	DefaultLocalPreference           types.Int64  `tfsdk:"default_local_preference"`
	CompareMedFromDifferentNeighbors types.Bool   `tfsdk:"compare_med_from_different_neighbors"`
	CompareRouterIdInPath            types.Bool   `tfsdk:"compare_router_id_in_path"`
	PickBestMed                      types.Bool   `tfsdk:"pick_best_med"`
	MissingMedAsBest                 types.Bool   `tfsdk:"missing_med_as_best"`
	KeepaliveInterval                types.Int64  `tfsdk:"keepalive_interval"`
	HoldTime                         types.Int64  `tfsdk:"hold_time"`
	MinHoldTime                      types.Int64  `tfsdk:"min_hold_time"`
	NextHopAddressTracking           types.Bool   `tfsdk:"next_hop_address_tracking"`
	NextHopDelayInterval             types.Int64  `tfsdk:"next_hop_delay_interval"`
	GracefulRestart                  types.Bool   `tfsdk:"graceful_restart"`
	GracefulRestartRestartTime       types.Int64  `tfsdk:"graceful_restart_restart_time"`
	GracefulRestartStalePathTime     types.Int64  `tfsdk:"graceful_restart_stale_path_time"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceBGPGeneralSettings) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/bgpgeneralsettings", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceBGPGeneralSettings) toBody(ctx context.Context, state DeviceBGPGeneralSettings) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.AsNumber.IsNull() {
		body, _ = sjson.Set(body, "asNumber", data.AsNumber.ValueString())
	}
	if !data.RouterId.IsNull() {
		body, _ = sjson.Set(body, "routerId", data.RouterId.ValueString())
	}
	if !data.ScanningInterval.IsNull() {
		body, _ = sjson.Set(body, "scanTime", data.ScanningInterval.ValueInt64())
	}
	if !data.AsNumberInPathAttribute.IsNull() {
		body, _ = sjson.Set(body, "maxasLimit", data.AsNumberInPathAttribute.ValueInt64())
	}
	if !data.LogNeighborChanges.IsNull() {
		body, _ = sjson.Set(body, "logNeighborChanges", data.LogNeighborChanges.ValueBool())
	}
	if !data.TcpPathMtuDiscovery.IsNull() {
		body, _ = sjson.Set(body, "transportPathMtuDiscovery", data.TcpPathMtuDiscovery.ValueBool())
	}
	if !data.ResetSessionUponFailover.IsNull() {
		body, _ = sjson.Set(body, "fastExternalFallOver", data.ResetSessionUponFailover.ValueBool())
	}
	if !data.EnforceFirstPeerAs.IsNull() {
		body, _ = sjson.Set(body, "enforceFirstAs", data.EnforceFirstPeerAs.ValueBool())
	}
	if !data.UseDotNotation.IsNull() {
		body, _ = sjson.Set(body, "asnotationDot", data.UseDotNotation.ValueBool())
	}
	if !data.AggregateTimer.IsNull() {
		body, _ = sjson.Set(body, "aggregateTimer", data.AggregateTimer.ValueInt64())
	}
	if !data.DefaultLocalPreference.IsNull() {
		body, _ = sjson.Set(body, "bestPath.defaultLocalPreferenceValue", data.DefaultLocalPreference.ValueInt64())
	}
	if !data.CompareMedFromDifferentNeighbors.IsNull() {
		body, _ = sjson.Set(body, "bestPath.alwaysCompareMed", data.CompareMedFromDifferentNeighbors.ValueBool())
	}
	if !data.CompareRouterIdInPath.IsNull() {
		body, _ = sjson.Set(body, "bestPath.deterministicMed", data.CompareRouterIdInPath.ValueBool())
	}
	if !data.PickBestMed.IsNull() {
		body, _ = sjson.Set(body, "bestPath.bestPathCompareRouterId", data.PickBestMed.ValueBool())
	}
	if !data.MissingMedAsBest.IsNull() {
		body, _ = sjson.Set(body, "bestPath.bestPathMedMissingAsWorst", data.MissingMedAsBest.ValueBool())
	}
	if !data.KeepaliveInterval.IsNull() {
		body, _ = sjson.Set(body, "bgptimers.keepAlive", data.KeepaliveInterval.ValueInt64())
	}
	if !data.HoldTime.IsNull() {
		body, _ = sjson.Set(body, "bgptimers.holdTime", data.HoldTime.ValueInt64())
	}
	if !data.MinHoldTime.IsNull() {
		body, _ = sjson.Set(body, "bgptimers.minHoldTime", data.MinHoldTime.ValueInt64())
	}
	if !data.NextHopAddressTracking.IsNull() {
		body, _ = sjson.Set(body, "bgpNextHopTriggerEnable", data.NextHopAddressTracking.ValueBool())
	}
	if !data.NextHopDelayInterval.IsNull() {
		body, _ = sjson.Set(body, "bgpNextHopTriggerDelay", data.NextHopDelayInterval.ValueInt64())
	}
	if !data.GracefulRestart.IsNull() {
		body, _ = sjson.Set(body, "bgpGracefulRestart.gracefulRestart", data.GracefulRestart.ValueBool())
	}
	if !data.GracefulRestartRestartTime.IsNull() {
		body, _ = sjson.Set(body, "bgpGracefulRestart.gracefulRestartRestartTime", data.GracefulRestartRestartTime.ValueInt64())
	}
	if !data.GracefulRestartStalePathTime.IsNull() {
		body, _ = sjson.Set(body, "bgpGracefulRestart.gracefulRestartStalePathTime", data.GracefulRestartStalePathTime.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceBGPGeneralSettings) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("asNumber"); value.Exists() {
		data.AsNumber = types.StringValue(value.String())
	} else {
		data.AsNumber = types.StringNull()
	}
	if value := res.Get("routerId"); value.Exists() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	if value := res.Get("scanTime"); value.Exists() {
		data.ScanningInterval = types.Int64Value(value.Int())
	} else {
		data.ScanningInterval = types.Int64Null()
	}
	if value := res.Get("maxasLimit"); value.Exists() {
		data.AsNumberInPathAttribute = types.Int64Value(value.Int())
	} else {
		data.AsNumberInPathAttribute = types.Int64Null()
	}
	if value := res.Get("logNeighborChanges"); value.Exists() {
		data.LogNeighborChanges = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborChanges = types.BoolNull()
	}
	if value := res.Get("transportPathMtuDiscovery"); value.Exists() {
		data.TcpPathMtuDiscovery = types.BoolValue(value.Bool())
	} else {
		data.TcpPathMtuDiscovery = types.BoolNull()
	}
	if value := res.Get("fastExternalFallOver"); value.Exists() {
		data.ResetSessionUponFailover = types.BoolValue(value.Bool())
	} else {
		data.ResetSessionUponFailover = types.BoolNull()
	}
	if value := res.Get("enforceFirstAs"); value.Exists() {
		data.EnforceFirstPeerAs = types.BoolValue(value.Bool())
	} else {
		data.EnforceFirstPeerAs = types.BoolNull()
	}
	if value := res.Get("asnotationDot"); value.Exists() {
		data.UseDotNotation = types.BoolValue(value.Bool())
	} else {
		data.UseDotNotation = types.BoolNull()
	}
	if value := res.Get("aggregateTimer"); value.Exists() {
		data.AggregateTimer = types.Int64Value(value.Int())
	} else {
		data.AggregateTimer = types.Int64Null()
	}
	if value := res.Get("bestPath.defaultLocalPreferenceValue"); value.Exists() {
		data.DefaultLocalPreference = types.Int64Value(value.Int())
	} else {
		data.DefaultLocalPreference = types.Int64Null()
	}
	if value := res.Get("bestPath.alwaysCompareMed"); value.Exists() {
		data.CompareMedFromDifferentNeighbors = types.BoolValue(value.Bool())
	} else {
		data.CompareMedFromDifferentNeighbors = types.BoolNull()
	}
	if value := res.Get("bestPath.deterministicMed"); value.Exists() {
		data.CompareRouterIdInPath = types.BoolValue(value.Bool())
	} else {
		data.CompareRouterIdInPath = types.BoolNull()
	}
	if value := res.Get("bestPath.bestPathCompareRouterId"); value.Exists() {
		data.PickBestMed = types.BoolValue(value.Bool())
	} else {
		data.PickBestMed = types.BoolNull()
	}
	if value := res.Get("bestPath.bestPathMedMissingAsWorst"); value.Exists() {
		data.MissingMedAsBest = types.BoolValue(value.Bool())
	} else {
		data.MissingMedAsBest = types.BoolNull()
	}
	if value := res.Get("bgptimers.keepAlive"); value.Exists() {
		data.KeepaliveInterval = types.Int64Value(value.Int())
	} else {
		data.KeepaliveInterval = types.Int64Null()
	}
	if value := res.Get("bgptimers.holdTime"); value.Exists() {
		data.HoldTime = types.Int64Value(value.Int())
	} else {
		data.HoldTime = types.Int64Null()
	}
	if value := res.Get("bgptimers.minHoldTime"); value.Exists() {
		data.MinHoldTime = types.Int64Value(value.Int())
	} else {
		data.MinHoldTime = types.Int64Null()
	}
	if value := res.Get("bgpNextHopTriggerEnable"); value.Exists() {
		data.NextHopAddressTracking = types.BoolValue(value.Bool())
	} else {
		data.NextHopAddressTracking = types.BoolNull()
	}
	if value := res.Get("bgpNextHopTriggerDelay"); value.Exists() {
		data.NextHopDelayInterval = types.Int64Value(value.Int())
	} else {
		data.NextHopDelayInterval = types.Int64Null()
	}
	if value := res.Get("bgpGracefulRestart.gracefulRestart"); value.Exists() {
		data.GracefulRestart = types.BoolValue(value.Bool())
	} else {
		data.GracefulRestart = types.BoolNull()
	}
	if value := res.Get("bgpGracefulRestart.gracefulRestartRestartTime"); value.Exists() {
		data.GracefulRestartRestartTime = types.Int64Value(value.Int())
	} else {
		data.GracefulRestartRestartTime = types.Int64Null()
	}
	if value := res.Get("bgpGracefulRestart.gracefulRestartStalePathTime"); value.Exists() {
		data.GracefulRestartStalePathTime = types.Int64Value(value.Int())
	} else {
		data.GracefulRestartStalePathTime = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceBGPGeneralSettings) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("asNumber"); value.Exists() && !data.AsNumber.IsNull() {
		data.AsNumber = types.StringValue(value.String())
	} else {
		data.AsNumber = types.StringNull()
	}
	if value := res.Get("routerId"); value.Exists() && !data.RouterId.IsNull() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	if value := res.Get("scanTime"); value.Exists() && !data.ScanningInterval.IsNull() {
		data.ScanningInterval = types.Int64Value(value.Int())
	} else {
		data.ScanningInterval = types.Int64Null()
	}
	if value := res.Get("maxasLimit"); value.Exists() && !data.AsNumberInPathAttribute.IsNull() {
		data.AsNumberInPathAttribute = types.Int64Value(value.Int())
	} else {
		data.AsNumberInPathAttribute = types.Int64Null()
	}
	if value := res.Get("logNeighborChanges"); value.Exists() && !data.LogNeighborChanges.IsNull() {
		data.LogNeighborChanges = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborChanges = types.BoolNull()
	}
	if value := res.Get("transportPathMtuDiscovery"); value.Exists() && !data.TcpPathMtuDiscovery.IsNull() {
		data.TcpPathMtuDiscovery = types.BoolValue(value.Bool())
	} else {
		data.TcpPathMtuDiscovery = types.BoolNull()
	}
	if value := res.Get("fastExternalFallOver"); value.Exists() && !data.ResetSessionUponFailover.IsNull() {
		data.ResetSessionUponFailover = types.BoolValue(value.Bool())
	} else {
		data.ResetSessionUponFailover = types.BoolNull()
	}
	if value := res.Get("enforceFirstAs"); value.Exists() && !data.EnforceFirstPeerAs.IsNull() {
		data.EnforceFirstPeerAs = types.BoolValue(value.Bool())
	} else {
		data.EnforceFirstPeerAs = types.BoolNull()
	}
	if value := res.Get("asnotationDot"); value.Exists() && !data.UseDotNotation.IsNull() {
		data.UseDotNotation = types.BoolValue(value.Bool())
	} else {
		data.UseDotNotation = types.BoolNull()
	}
	if value := res.Get("aggregateTimer"); value.Exists() && !data.AggregateTimer.IsNull() {
		data.AggregateTimer = types.Int64Value(value.Int())
	} else {
		data.AggregateTimer = types.Int64Null()
	}
	if value := res.Get("bestPath.defaultLocalPreferenceValue"); value.Exists() && !data.DefaultLocalPreference.IsNull() {
		data.DefaultLocalPreference = types.Int64Value(value.Int())
	} else {
		data.DefaultLocalPreference = types.Int64Null()
	}
	if value := res.Get("bestPath.alwaysCompareMed"); value.Exists() && !data.CompareMedFromDifferentNeighbors.IsNull() {
		data.CompareMedFromDifferentNeighbors = types.BoolValue(value.Bool())
	} else {
		data.CompareMedFromDifferentNeighbors = types.BoolNull()
	}
	if value := res.Get("bestPath.deterministicMed"); value.Exists() && !data.CompareRouterIdInPath.IsNull() {
		data.CompareRouterIdInPath = types.BoolValue(value.Bool())
	} else {
		data.CompareRouterIdInPath = types.BoolNull()
	}
	if value := res.Get("bestPath.bestPathCompareRouterId"); value.Exists() && !data.PickBestMed.IsNull() {
		data.PickBestMed = types.BoolValue(value.Bool())
	} else {
		data.PickBestMed = types.BoolNull()
	}
	if value := res.Get("bestPath.bestPathMedMissingAsWorst"); value.Exists() && !data.MissingMedAsBest.IsNull() {
		data.MissingMedAsBest = types.BoolValue(value.Bool())
	} else {
		data.MissingMedAsBest = types.BoolNull()
	}
	if value := res.Get("bgptimers.keepAlive"); value.Exists() && !data.KeepaliveInterval.IsNull() {
		data.KeepaliveInterval = types.Int64Value(value.Int())
	} else {
		data.KeepaliveInterval = types.Int64Null()
	}
	if value := res.Get("bgptimers.holdTime"); value.Exists() && !data.HoldTime.IsNull() {
		data.HoldTime = types.Int64Value(value.Int())
	} else {
		data.HoldTime = types.Int64Null()
	}
	if value := res.Get("bgptimers.minHoldTime"); value.Exists() && !data.MinHoldTime.IsNull() {
		data.MinHoldTime = types.Int64Value(value.Int())
	} else {
		data.MinHoldTime = types.Int64Null()
	}
	if value := res.Get("bgpNextHopTriggerEnable"); value.Exists() && !data.NextHopAddressTracking.IsNull() {
		data.NextHopAddressTracking = types.BoolValue(value.Bool())
	} else {
		data.NextHopAddressTracking = types.BoolNull()
	}
	if value := res.Get("bgpNextHopTriggerDelay"); value.Exists() && !data.NextHopDelayInterval.IsNull() {
		data.NextHopDelayInterval = types.Int64Value(value.Int())
	} else {
		data.NextHopDelayInterval = types.Int64Null()
	}
	if value := res.Get("bgpGracefulRestart.gracefulRestart"); value.Exists() && !data.GracefulRestart.IsNull() {
		data.GracefulRestart = types.BoolValue(value.Bool())
	} else {
		data.GracefulRestart = types.BoolNull()
	}
	if value := res.Get("bgpGracefulRestart.gracefulRestartRestartTime"); value.Exists() && !data.GracefulRestartRestartTime.IsNull() {
		data.GracefulRestartRestartTime = types.Int64Value(value.Int())
	} else {
		data.GracefulRestartRestartTime = types.Int64Null()
	}
	if value := res.Get("bgpGracefulRestart.gracefulRestartStalePathTime"); value.Exists() && !data.GracefulRestartStalePathTime.IsNull() {
		data.GracefulRestartStalePathTime = types.Int64Value(value.Int())
	} else {
		data.GracefulRestartStalePathTime = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceBGPGeneralSettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Name.IsUnknown() {
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
