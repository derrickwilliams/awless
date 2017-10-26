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

package awsspec

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func init() {
	genTestsParams["createsecuritygroup"] = map[string]interface{}{
		"name":        "my-sg-name",
		"vpc":         "my-vpc-id",
		"description": "security group description",
	}
	genTestsExpected["createsecuritygroup"] = &ec2.CreateSecurityGroupInput{
		GroupName:   String("my-sg-name"),
		VpcId:       String("my-vpc-id"),
		Description: String("security group description"),
	}
	genTestsOutputExtractFunc["createsecuritygroup"] = func() interface{} {
		return &ec2.CreateSecurityGroupOutput{GroupId: String("id-my-secgroup")}
	}
	genTestsOutput["createsecuritygroup"] = "id-my-secgroup"

	genTestsParams["deletesecuritygroup"] = map[string]interface{}{
		"id": "my-sg-id",
	}
	genTestsExpected["deletesecuritygroup"] = &ec2.DeleteSecurityGroupInput{
		GroupId: String("my-sg-id"),
	}
	genTestsParams["updatesecuritygroup"] = map[string]interface{}{
		"id":        "my-sg-id",
		"inbound":   "authorize",
		"protocol":  "tcp",
		"cidr":      "10.0.0.0/10",
		"portrange": "10-22",
	}
	genTestsExpected["updatesecuritygroup"] = &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: String("my-sg-id"),
		IpPermissions: []*ec2.IpPermission{
			{
				IpProtocol: String("tcp"),
				IpRanges: []*ec2.IpRange{
					{CidrIp: String("10.0.0.0/10")},
				},
				FromPort: Int64(10),
				ToPort:   Int64(22),
			},
		},
	}
	genTestsParams["attachsecuritygroup"] = map[string]interface{}{
		"id":       "secgroup-3",
		"instance": "attach/detachsecgroup-instance-id",
	}
	genTestsExpected["attachsecuritygroup"] = &ec2.ModifyInstanceAttributeInput{
		InstanceId: String("attach/detachsecgroup-instance-id"),
		Groups:     []*string{String("secgroup-1"), String("secgroup-2"), String("secgroup-3")},
	}
	genTestsParams["detachsecuritygroup"] = map[string]interface{}{
		"id":       "secgroup-2",
		"instance": "attach/detachsecgroup-instance-id",
	}
	genTestsExpected["detachsecuritygroup"] = &ec2.ModifyInstanceAttributeInput{
		InstanceId: String("attach/detachsecgroup-instance-id"),
		Groups:     []*string{String("secgroup-1")},
	}
	genTestsParams["checksecuritygroup"] = map[string]interface{}{
		"id":      "my-check-secgroup",
		"state":   "unused",
		"timeout": 0,
	}
}

func TestBuildIpPermissionsFromParams(t *testing.T) {
	tcases := []struct {
		params   map[string]interface{}
		expected []*ec2.IpPermission
	}{
		{
			params: map[string]interface{}{
				"protocol":  "tcp",
				"cidr":      "192.168.1.10/24",
				"portrange": 80,
			},
			expected: []*ec2.IpPermission{
				{
					IpProtocol: aws.String("tcp"),
					IpRanges:   []*ec2.IpRange{{CidrIp: aws.String("192.168.1.10/24")}},
					FromPort:   aws.Int64(int64(80)),
					ToPort:     aws.Int64(int64(80)),
				},
			},
		},
		{
			params: map[string]interface{}{
				"protocol": "any",
				"cidr":     "192.168.1.18/32",
			},
			expected: []*ec2.IpPermission{
				{
					IpProtocol: aws.String("-1"),
					IpRanges:   []*ec2.IpRange{{CidrIp: aws.String("192.168.1.18/32")}},
					FromPort:   aws.Int64(int64(-1)),
					ToPort:     aws.Int64(int64(-1)),
				},
			},
		},
		{
			params: map[string]interface{}{
				"protocol":  "udp",
				"cidr":      "0.0.0.0/0",
				"portrange": "22-23",
			},
			expected: []*ec2.IpPermission{
				{
					IpProtocol: aws.String("udp"),
					IpRanges:   []*ec2.IpRange{{CidrIp: aws.String("0.0.0.0/0")}},
					FromPort:   aws.Int64(int64(22)),
					ToPort:     aws.Int64(int64(23)),
				},
			},
		},
		{
			params: map[string]interface{}{
				"protocol":  "icmp",
				"cidr":      "10.0.0.0/16",
				"portrange": "any",
			},
			expected: []*ec2.IpPermission{
				{
					IpProtocol: aws.String("icmp"),
					IpRanges:   []*ec2.IpRange{{CidrIp: aws.String("10.0.0.0/16")}},
					FromPort:   aws.Int64(int64(-1)),
					ToPort:     aws.Int64(int64(-1)),
				},
			},
		},
		{
			params: map[string]interface{}{
				"protocol":      "icmp",
				"securitygroup": "sg-12345",
				"portrange":     "any",
			},
			expected: []*ec2.IpPermission{
				{
					IpProtocol:       aws.String("icmp"),
					UserIdGroupPairs: []*ec2.UserIdGroupPair{{GroupId: aws.String("sg-12345")}},
					FromPort:         aws.Int64(int64(-1)),
					ToPort:           aws.Int64(int64(-1)),
				},
			},
		},

		{
			params: map[string]interface{}{
				"protocol":      "tcp",
				"securitygroup": "sg-23456",
				"portrange":     80,
			},
			expected: []*ec2.IpPermission{
				{
					IpProtocol:       aws.String("tcp"),
					UserIdGroupPairs: []*ec2.UserIdGroupPair{{GroupId: aws.String("sg-23456")}},
					FromPort:         aws.Int64(int64(80)),
					ToPort:           aws.Int64(int64(80)),
				},
			},
		},
	}

	for i, tcase := range tcases {
		cmd := &UpdateSecuritygroup{}
		cmd.inject(tcase.params)
		ipPermissions, err := cmd.buildIpPermissions()
		if err != nil {
			t.Fatal(i+1, ":", err)
		}
		if got, want := ipPermissions, tcase.expected; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %+v, want %+v", i+1, got, want)
		}
	}
}

func (m *mockEc2) AuthorizeSecurityGroupIngress(input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	if got, want := input, genTestsExpected["updatesecuritygroup"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	return nil, nil
}

func (m *mockEc2) ModifyInstanceAttribute(input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
	if expectedAttach, expectedDetach := genTestsExpected["attachsecuritygroup"], genTestsExpected["detachsecuritygroup"]; !reflect.DeepEqual(input, expectedAttach) && !reflect.DeepEqual(input, expectedDetach) {
		return nil, fmt.Errorf("got %#v, want either %#v or %#v", input, expectedAttach, expectedDetach)
	}
	return nil, nil
}

func (m *mockEc2) DescribeInstanceAttribute(input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
	if StringValue(input.Attribute) == "groupSet" && StringValue(input.InstanceId) == "attach/detachsecgroup-instance-id" {
		return &ec2.DescribeInstanceAttributeOutput{Groups: []*ec2.GroupIdentifier{{GroupId: String("secgroup-1")}, {GroupId: String("secgroup-2")}}}, nil
	}
	return nil, fmt.Errorf("DescribeInstanceAttribute mock: invalid value for 'Attribute' or 'InstanceId'")
}

func (m *mockEc2) DescribeNetworkInterfaces(input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	exp := &ec2.DescribeNetworkInterfacesInput{
		Filters: []*ec2.Filter{
			{Name: String("group-id"), Values: []*string{String("my-check-secgroup")}},
		},
	}

	if reflect.DeepEqual(input, exp) {
		return &ec2.DescribeNetworkInterfacesOutput{}, nil
	}
	return nil, fmt.Errorf("DescribeNetworkInterfaces mock: invalid value for 'Filters'")
}
