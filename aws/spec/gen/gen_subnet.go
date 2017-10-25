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

type CreateSubnet struct {
	_                string `action: "create" entity: "subnet" awsAPI: "ec2" awsCall: "CreateSubnet" awsInput: "CreateSubnetInput" awsOutput: "CreateSubnetOutput" awsDryRun: ""`
	logger           *logger.Logger
	api              ec2iface.EC2API
	Cidr             *string   `awsName: "CidrBlock" awsType: "awsstr" templateName: "cidr" required: ""`
	Vpc              *string   `awsName: "VpcId" awsType: "awsstr" templateName: "vpc" required: ""`
	Availabilityzone *string   `awsName: "AvailabilityZone" awsType: "awsstr" templateName: "availabilityzone"`
	Name             *struct{} `awsName: "Name" templateName: "name"`
}
type UpdateSubnet struct {
	_      string `action: "update" entity: "subnet" awsAPI: "ec2" awsCall: "ModifySubnetAttribute" awsInput: "ModifySubnetAttributeInput" awsOutput: "ModifySubnetAttributeOutput"`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "SubnetId" awsType: "awsstr" templateName: "id" required: ""`
	Public *bool   `awsName: "MapPublicIpOnLaunch" awsType: "awsboolattribute" templateName: "public"`
}
type DeleteSubnet struct {
	_      string `action: "delete" entity: "subnet" awsAPI: "ec2" awsCall: "DeleteSubnet" awsInput: "DeleteSubnetInput" awsOutput: "DeleteSubnetOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "SubnetId" awsType: "awsstr" templateName: "id" required: ""`
}