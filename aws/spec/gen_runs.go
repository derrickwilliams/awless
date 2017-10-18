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
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/wallix/awless/logger"
)

func NewAttachPolicy(l *logger.Logger, sess *session.Session) *AttachPolicy {
	cmd := new(AttachPolicy)
	cmd.api = iam.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *AttachPolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("attach policy: BeforeRun: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("attach policy: cannot set params on command struct: %s", err)
	}

	output, err := cmd.ManualRun(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("attach policy: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("attach policy: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
}

func (cmd *AttachPolicy) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *AttachPolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateInstance(l *logger.Logger, sess *session.Session) *CreateInstance {
	cmd := new(CreateInstance)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("create instance: BeforeRun: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("create instance: cannot set params on command struct: %s", err)
	}

	input := &ec2.RunInstancesInput{}
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("create instance: cannot inject in ec2.RunInstancesInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.RunInstances(input)
	cmd.logger.ExtraVerbosef("ec2.RunInstances call took %s", time.Since(start))
	if err != nil {
		return nil, fmt.Errorf("create instance: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("create instance: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
}

func (cmd *CreateInstance) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *CreateInstance) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: create instance: cannot set params on command struct: %s", err)
	}

	input := &ec2.RunInstancesInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("dry run: create instance: cannot inject in ec2.RunInstancesInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.RunInstances(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.RunInstances call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create instance ok")
			return fakeDryRunId("instance"), nil
		}
	}

	return nil, fmt.Errorf("dry run: create instance : %s", err)
}

func (cmd *CreateInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateSubnet(l *logger.Logger, sess *session.Session) *CreateSubnet {
	cmd := new(CreateSubnet)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateSubnet) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("create subnet: BeforeRun: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("create subnet: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateSubnetInput{}
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("create subnet: cannot inject in ec2.CreateSubnetInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateSubnet(input)
	cmd.logger.ExtraVerbosef("ec2.CreateSubnet call took %s", time.Since(start))
	if err != nil {
		return nil, fmt.Errorf("create subnet: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("create subnet: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
}

func (cmd *CreateSubnet) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *CreateSubnet) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: create subnet: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateSubnetInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("dry run: create subnet: cannot inject in ec2.CreateSubnetInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateSubnet(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateSubnet call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create subnet ok")
			return fakeDryRunId("subnet"), nil
		}
	}

	return nil, fmt.Errorf("dry run: create subnet : %s", err)
}

func (cmd *CreateSubnet) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateTag(l *logger.Logger, sess *session.Session) *CreateTag {
	cmd := new(CreateTag)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateTag) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("create tag: BeforeRun: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("create tag: cannot set params on command struct: %s", err)
	}

	output, err := cmd.ManualRun(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("create tag: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("create tag: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
}

func (cmd *CreateTag) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *CreateTag) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}
