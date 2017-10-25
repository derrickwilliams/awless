/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// DO NOT EDIT
// This file was automatically generated with go generate
package awsspec

import (
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/wallix/awless/logger"
)

type CreateAppscalingpolicy struct {
	_                                 string `action: "create" entity: "appscalingpolicy" awsAPI: "applicationautoscaling" awsCall: "PutScalingPolicy" awsInput: "PutScalingPolicyInput" awsOutput: "PutScalingPolicyOutput"`
	logger                            *logger.Logger
	api                               applicationautoscalingiface.ApplicationAutoScalingAPI
	Name                              *string   `awsName: "PolicyName" awsType: "awsstr" templateName: "name" required: ""`
	Type                              *string   `awsName: "PolicyType" awsType: "awsstr" templateName: "type" required: ""`
	Resource                          *string   `awsName: "ResourceId" awsType: "awsstr" templateName: "resource" required: ""`
	Dimension                         *string   `awsName: "ScalableDimension" awsType: "awsstr" templateName: "dimension" required: ""`
	ServiceNamespace                  *string   `awsName: "ServiceNamespace" awsType: "awsstr" templateName: "service-namespace" required: ""`
	StepscalingAdjustmentType         *string   `awsName: "StepScalingPolicyConfiguration.AdjustmentType" awsType: "awsstr" templateName: "stepscaling-adjustment-type" required: ""`
	StepscalingAdjustments            *struct{} `awsName: "StepScalingPolicyConfiguration.StepAdjustments" awsType: "awsstepadjustments" templateName: "stepscaling-adjustments" required: ""`
	StepscalingCooldown               *int64    `awsName: "StepScalingPolicyConfiguration.Cooldown" awsType: "awsint64" templateName: "stepscaling-cooldown"`
	StepscalingAggregationType        *string   `awsName: "StepScalingPolicyConfiguration.MetricAggregationType" awsType: "awsstr" templateName: "stepscaling-aggregation-type"`
	StepscalingMinAdjustmentMagnitude *int64    `awsName: "StepScalingPolicyConfiguration.MinAdjustmentMagnitude" awsType: "awsint64" templateName: "stepscaling-min-adjustment-magnitude"`
}
type DeleteAppscalingpolicy struct {
	_                string `action: "delete" entity: "appscalingpolicy" awsAPI: "applicationautoscaling" awsCall: "DeleteScalingPolicy" awsInput: "DeleteScalingPolicyInput" awsOutput: "DeleteScalingPolicyOutput"`
	logger           *logger.Logger
	api              applicationautoscalingiface.ApplicationAutoScalingAPI
	Name             *string `awsName: "PolicyName" awsType: "awsstr" templateName: "name" required: ""`
	Resource         *string `awsName: "ResourceId" awsType: "awsstr" templateName: "resource" required: ""`
	Dimension        *string `awsName: "ScalableDimension" awsType: "awsstr" templateName: "dimension" required: ""`
	ServiceNamespace *string `awsName: "ServiceNamespace" awsType: "awsstr" templateName: "service-namespace" required: ""`
}