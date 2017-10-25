package awsspec

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateTag struct {
	_        string `action:"create" entity:"tag" awsAPI:"ec2"` //  awsCall:"CreateTags" awsInput:"ec2.CreateTagsInput" awsOutput:"ec2.CreateTagsOutput"
	result   string
	logger   *logger.Logger
	api      ec2iface.EC2API
	sess     *session.Session
	Resource *string `awsName:"Resources" awsType:"awsstringslice" templateName:"resource" required:""`
	Key      *string `templateName:"key" required:""`
	Value    *string `templateName:"value" required:""`
}

func (cmd *CreateTag) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateTag) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateTagsInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateTagsInput: %s", err)
	}
	input.Tags = []*ec2.Tag{{Key: cmd.Key, Value: cmd.Value}}

	start := time.Now()
	_, err := cmd.api.CreateTags(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateTags call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create tag ok")
			return fakeDryRunId("tag"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateTag) ManualRun(ctx, params map[string]interface{}) (interface{}, error) {
	input := &ec2.CreateTagsInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateTagsInput: %s", err)
	}
	input.Tags = []*ec2.Tag{{Key: cmd.Key, Value: cmd.Value}}

	start := time.Now()
	req, _ := cmd.api.CreateTagsRequest(input)
	req.Retryer = createTagRetryer{}
	if err := req.Send(); err != nil {
		return nil, err
	}
	cmd.logger.ExtraVerbosef("ec2.CreateTags call took %s", time.Since(start))
	return cmd.result, nil
}

type createTagRetryer struct {
	client.DefaultRetryer
}

func (d createTagRetryer) MaxRetries() int { return 5 }
func (d createTagRetryer) ShouldRetry(r *request.Request) bool {
	if d.DefaultRetryer.ShouldRetry(r) || !(r.HTTPResponse.StatusCode < 300 && r.HTTPResponse.StatusCode >= 200) {
		return true
	}

	return false
}