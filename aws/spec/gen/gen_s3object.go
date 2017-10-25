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
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/wallix/awless/logger"
)

type CreateS3object struct {
	_      string `action: "create" entity: "s3object" awsAPI: "s3"`
	logger *logger.Logger
	api    s3iface.S3API
	Bucket *string `awsName: "Bucket" awsType: "awsstr" templateName: "bucket" required: ""`
	File   *string `awsName: "Body" awsType: "awsstr" templateName: "file" required: ""`
	Name   *string `awsName: "Key" awsType: "awsstr" templateName: "name"`
	Acl    *string `awsName: "ACL" awsType: "awsstr" templateName: "acl"`
}
type UpdateS3object struct {
	_       string `action: "update" entity: "s3object" awsAPI: "s3" awsCall: "PutObjectAcl" awsInput: "PutObjectAclInput" awsOutput: "PutObjectAclOutput"`
	logger  *logger.Logger
	api     s3iface.S3API
	Bucket  *string `awsName: "Bucket" awsType: "awsstr" templateName: "bucket" required: ""`
	Name    *string `awsName: "Key" awsType: "awsstr" templateName: "name" required: ""`
	Acl     *string `awsName: "ACL" awsType: "awsstr" templateName: "acl" required: ""`
	Version *string `awsName: "VersionId" awsType: "awsstr" templateName: "version"`
}
type DeleteS3object struct {
	_      string `action: "delete" entity: "s3object" awsAPI: "s3" awsCall: "DeleteObject" awsInput: "DeleteObjectInput" awsOutput: "DeleteObjectOutput"`
	logger *logger.Logger
	api    s3iface.S3API
	Bucket *string `awsName: "Bucket" awsType: "awsstr" templateName: "bucket" required: ""`
	Name   *string `awsName: "Key" awsType: "awsstr" templateName: "name" required: ""`
}