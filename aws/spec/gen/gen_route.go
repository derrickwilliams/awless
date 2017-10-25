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
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateRoute struct {
	_       string `action: "create" entity: "route" awsAPI: "ec2" awsCall: "CreateRoute" awsInput: "CreateRouteInput" awsOutput: "CreateRouteOutput" awsDryRun: ""`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Table   *string `awsName: "RouteTableId" awsType: "awsstr" templateName: "table" required: ""`
	Cidr    *string `awsName: "DestinationCidrBlock" awsType: "awsstr" templateName: "cidr" required: ""`
	Gateway *string `awsName: "GatewayId" awsType: "awsstr" templateName: "gateway" required: ""`
}
type DeleteRoute struct {
	_      string `action: "delete" entity: "route" awsAPI: "ec2" awsCall: "DeleteRoute" awsInput: "DeleteRouteInput" awsOutput: "DeleteRouteOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Table  *string `awsName: "RouteTableId" awsType: "awsstr" templateName: "table" required: ""`
	Cidr   *string `awsName: "DestinationCidrBlock" awsType: "awsstr" templateName: "cidr" required: ""`
}