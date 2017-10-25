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
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/wallix/awless/logger"
)

type CreateQueue struct {
	_                 string `action: "create" entity: "queue" awsAPI: "sqs" awsCall: "CreateQueue" awsInput: "CreateQueueInput" awsOutput: "CreateQueueOutput"`
	logger            *logger.Logger
	api               sqsiface.SQSAPI
	Name              *string   `awsName: "QueueName" awsType: "awsstr" templateName: "name" required: ""`
	Delay             *struct{} `awsName: "Attributes[DelaySeconds]" awsType: "awsstringpointermap" templateName: "delay"`
	MaxMsgSize        *struct{} `awsName: "Attributes[MaximumMessageSize]" awsType: "awsstringpointermap" templateName: "max-msg-size"`
	RetentionPeriod   *struct{} `awsName: "Attributes[MessageRetentionPeriod]" awsType: "awsstringpointermap" templateName: "retention-period"`
	Policy            *struct{} `awsName: "Attributes[Policy]" awsType: "awsstringpointermap" templateName: "policy"`
	MsgWait           *struct{} `awsName: "Attributes[ReceiveMessageWaitTimeSeconds]" awsType: "awsstringpointermap" templateName: "msg-wait"`
	RedrivePolicy     *struct{} `awsName: "Attributes[RedrivePolicy]" awsType: "awsstringpointermap" templateName: "redrive-policy"`
	VisibilityTimeout *struct{} `awsName: "Attributes[VisibilityTimeout]" awsType: "awsstringpointermap" templateName: "visibility-timeout"`
}
type DeleteQueue struct {
	_      string `action: "delete" entity: "queue" awsAPI: "sqs" awsCall: "DeleteQueue" awsInput: "DeleteQueueInput" awsOutput: "DeleteQueueOutput"`
	logger *logger.Logger
	api    sqsiface.SQSAPI
	Url    *string `awsName: "QueueUrl" awsType: "awsstr" templateName: "url" required: ""`
}