// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/khulnasoft/devtool/content-service/pkg/storage (interfaces: PresignedAccess,DirectAccess,PresignedS3Client,S3Client)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	archive "github.com/khulnasoft/devtool/content-service/pkg/archive"
	storage "github.com/khulnasoft/devtool/content-service/pkg/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockPresignedAccess is a mock of PresignedAccess interface.
type MockPresignedAccess struct {
	ctrl     *gomock.Controller
	recorder *MockPresignedAccessMockRecorder
}

// MockPresignedAccessMockRecorder is the mock recorder for MockPresignedAccess.
type MockPresignedAccessMockRecorder struct {
	mock *MockPresignedAccess
}

// NewMockPresignedAccess creates a new mock instance.
func NewMockPresignedAccess(ctrl *gomock.Controller) *MockPresignedAccess {
	mock := &MockPresignedAccess{ctrl: ctrl}
	mock.recorder = &MockPresignedAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPresignedAccess) EXPECT() *MockPresignedAccessMockRecorder {
	return m.recorder
}

// BackupObject mocks base method.
func (m *MockPresignedAccess) BackupObject(arg0, arg1, arg2 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	return ret0
}

// BackupObject indicates an expected call of BackupObject.
func (mr *MockPresignedAccessMockRecorder) BackupObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupObject", reflect.TypeOf((*MockPresignedAccess)(nil).BackupObject), arg0, arg1, arg2)
}

// BlobObject mocks base method.
func (m *MockPresignedAccess) BlobObject(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlobObject", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlobObject indicates an expected call of BlobObject.
func (mr *MockPresignedAccessMockRecorder) BlobObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlobObject", reflect.TypeOf((*MockPresignedAccess)(nil).BlobObject), arg0, arg1)
}

// Bucket mocks base method.
func (m *MockPresignedAccess) Bucket(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bucket", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Bucket indicates an expected call of Bucket.
func (mr *MockPresignedAccessMockRecorder) Bucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bucket", reflect.TypeOf((*MockPresignedAccess)(nil).Bucket), arg0)
}

// DeleteBucket mocks base method.
func (m *MockPresignedAccess) DeleteBucket(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockPresignedAccessMockRecorder) DeleteBucket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockPresignedAccess)(nil).DeleteBucket), arg0, arg1, arg2)
}

// DeleteObject mocks base method.
func (m *MockPresignedAccess) DeleteObject(arg0 context.Context, arg1 string, arg2 *storage.DeleteObjectQuery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockPresignedAccessMockRecorder) DeleteObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockPresignedAccess)(nil).DeleteObject), arg0, arg1, arg2)
}

// DiskUsage mocks base method.
func (m *MockPresignedAccess) DiskUsage(arg0 context.Context, arg1, arg2 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiskUsage", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiskUsage indicates an expected call of DiskUsage.
func (mr *MockPresignedAccessMockRecorder) DiskUsage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiskUsage", reflect.TypeOf((*MockPresignedAccess)(nil).DiskUsage), arg0, arg1, arg2)
}

// EnsureExists mocks base method.
func (m *MockPresignedAccess) EnsureExists(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureExists indicates an expected call of EnsureExists.
func (mr *MockPresignedAccessMockRecorder) EnsureExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureExists", reflect.TypeOf((*MockPresignedAccess)(nil).EnsureExists), arg0, arg1)
}

// InstanceObject mocks base method.
func (m *MockPresignedAccess) InstanceObject(arg0, arg1, arg2, arg3 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceObject", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	return ret0
}

// InstanceObject indicates an expected call of InstanceObject.
func (mr *MockPresignedAccessMockRecorder) InstanceObject(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceObject", reflect.TypeOf((*MockPresignedAccess)(nil).InstanceObject), arg0, arg1, arg2, arg3)
}

// ObjectExists mocks base method.
func (m *MockPresignedAccess) ObjectExists(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectExists", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectExists indicates an expected call of ObjectExists.
func (mr *MockPresignedAccessMockRecorder) ObjectExists(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectExists", reflect.TypeOf((*MockPresignedAccess)(nil).ObjectExists), arg0, arg1, arg2)
}

// ObjectHash mocks base method.
func (m *MockPresignedAccess) ObjectHash(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectHash", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectHash indicates an expected call of ObjectHash.
func (mr *MockPresignedAccessMockRecorder) ObjectHash(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectHash", reflect.TypeOf((*MockPresignedAccess)(nil).ObjectHash), arg0, arg1, arg2)
}

// SignDownload mocks base method.
func (m *MockPresignedAccess) SignDownload(arg0 context.Context, arg1, arg2 string, arg3 *storage.SignedURLOptions) (*storage.DownloadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignDownload", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*storage.DownloadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignDownload indicates an expected call of SignDownload.
func (mr *MockPresignedAccessMockRecorder) SignDownload(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignDownload", reflect.TypeOf((*MockPresignedAccess)(nil).SignDownload), arg0, arg1, arg2, arg3)
}

// SignUpload mocks base method.
func (m *MockPresignedAccess) SignUpload(arg0 context.Context, arg1, arg2 string, arg3 *storage.SignedURLOptions) (*storage.UploadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUpload", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*storage.UploadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUpload indicates an expected call of SignUpload.
func (mr *MockPresignedAccessMockRecorder) SignUpload(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUpload", reflect.TypeOf((*MockPresignedAccess)(nil).SignUpload), arg0, arg1, arg2, arg3)
}

// MockDirectAccess is a mock of DirectAccess interface.
type MockDirectAccess struct {
	ctrl     *gomock.Controller
	recorder *MockDirectAccessMockRecorder
}

// MockDirectAccessMockRecorder is the mock recorder for MockDirectAccess.
type MockDirectAccessMockRecorder struct {
	mock *MockDirectAccess
}

// NewMockDirectAccess creates a new mock instance.
func NewMockDirectAccess(ctrl *gomock.Controller) *MockDirectAccess {
	mock := &MockDirectAccess{ctrl: ctrl}
	mock.recorder = &MockDirectAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectAccess) EXPECT() *MockDirectAccessMockRecorder {
	return m.recorder
}

// BackupObject mocks base method.
func (m *MockDirectAccess) BackupObject(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupObject", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// BackupObject indicates an expected call of BackupObject.
func (mr *MockDirectAccessMockRecorder) BackupObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupObject", reflect.TypeOf((*MockDirectAccess)(nil).BackupObject), arg0)
}

// Bucket mocks base method.
func (m *MockDirectAccess) Bucket(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bucket", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Bucket indicates an expected call of Bucket.
func (mr *MockDirectAccessMockRecorder) Bucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bucket", reflect.TypeOf((*MockDirectAccess)(nil).Bucket), arg0)
}

// Download mocks base method.
func (m *MockDirectAccess) Download(arg0 context.Context, arg1, arg2 string, arg3 []archive.IDMapping) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockDirectAccessMockRecorder) Download(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockDirectAccess)(nil).Download), arg0, arg1, arg2, arg3)
}

// DownloadSnapshot mocks base method.
func (m *MockDirectAccess) DownloadSnapshot(arg0 context.Context, arg1, arg2 string, arg3 []archive.IDMapping) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadSnapshot", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadSnapshot indicates an expected call of DownloadSnapshot.
func (mr *MockDirectAccessMockRecorder) DownloadSnapshot(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadSnapshot", reflect.TypeOf((*MockDirectAccess)(nil).DownloadSnapshot), arg0, arg1, arg2, arg3)
}

// EnsureExists mocks base method.
func (m *MockDirectAccess) EnsureExists(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureExists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureExists indicates an expected call of EnsureExists.
func (mr *MockDirectAccessMockRecorder) EnsureExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureExists", reflect.TypeOf((*MockDirectAccess)(nil).EnsureExists), arg0)
}

// Init mocks base method.
func (m *MockDirectAccess) Init(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockDirectAccessMockRecorder) Init(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDirectAccess)(nil).Init), arg0, arg1, arg2, arg3)
}

// ListObjects mocks base method.
func (m *MockDirectAccess) ListObjects(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjects", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjects indicates an expected call of ListObjects.
func (mr *MockDirectAccessMockRecorder) ListObjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*MockDirectAccess)(nil).ListObjects), arg0, arg1)
}

// Qualify mocks base method.
func (m *MockDirectAccess) Qualify(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Qualify", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Qualify indicates an expected call of Qualify.
func (mr *MockDirectAccessMockRecorder) Qualify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Qualify", reflect.TypeOf((*MockDirectAccess)(nil).Qualify), arg0)
}

// Upload mocks base method.
func (m *MockDirectAccess) Upload(arg0 context.Context, arg1, arg2 string, arg3 ...storage.UploadOption) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Upload indicates an expected call of Upload.
func (mr *MockDirectAccessMockRecorder) Upload(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockDirectAccess)(nil).Upload), varargs...)
}

// UploadInstance mocks base method.
func (m *MockDirectAccess) UploadInstance(arg0 context.Context, arg1, arg2 string, arg3 ...storage.UploadOption) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadInstance", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UploadInstance indicates an expected call of UploadInstance.
func (mr *MockDirectAccessMockRecorder) UploadInstance(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadInstance", reflect.TypeOf((*MockDirectAccess)(nil).UploadInstance), varargs...)
}

// MockPresignedS3Client is a mock of PresignedS3Client interface.
type MockPresignedS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockPresignedS3ClientMockRecorder
}

// MockPresignedS3ClientMockRecorder is the mock recorder for MockPresignedS3Client.
type MockPresignedS3ClientMockRecorder struct {
	mock *MockPresignedS3Client
}

// NewMockPresignedS3Client creates a new mock instance.
func NewMockPresignedS3Client(ctrl *gomock.Controller) *MockPresignedS3Client {
	mock := &MockPresignedS3Client{ctrl: ctrl}
	mock.recorder = &MockPresignedS3ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPresignedS3Client) EXPECT() *MockPresignedS3ClientMockRecorder {
	return m.recorder
}

// PresignGetObject mocks base method.
func (m *MockPresignedS3Client) PresignGetObject(arg0 context.Context, arg1 *s3.GetObjectInput, arg2 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PresignGetObject", varargs...)
	ret0, _ := ret[0].(*v4.PresignedHTTPRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresignGetObject indicates an expected call of PresignGetObject.
func (mr *MockPresignedS3ClientMockRecorder) PresignGetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresignGetObject", reflect.TypeOf((*MockPresignedS3Client)(nil).PresignGetObject), varargs...)
}

// PresignPutObject mocks base method.
func (m *MockPresignedS3Client) PresignPutObject(arg0 context.Context, arg1 *s3.PutObjectInput, arg2 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PresignPutObject", varargs...)
	ret0, _ := ret[0].(*v4.PresignedHTTPRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresignPutObject indicates an expected call of PresignPutObject.
func (mr *MockPresignedS3ClientMockRecorder) PresignPutObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresignPutObject", reflect.TypeOf((*MockPresignedS3Client)(nil).PresignPutObject), varargs...)
}

// MockS3Client is a mock of S3Client interface.
type MockS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientMockRecorder
}

// MockS3ClientMockRecorder is the mock recorder for MockS3Client.
type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

// NewMockS3Client creates a new mock instance.
func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

// DeleteObjects mocks base method.
func (m *MockS3Client) DeleteObjects(arg0 context.Context, arg1 *s3.DeleteObjectsInput, arg2 ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObjects", varargs...)
	ret0, _ := ret[0].(*s3.DeleteObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjects indicates an expected call of DeleteObjects.
func (mr *MockS3ClientMockRecorder) DeleteObjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjects", reflect.TypeOf((*MockS3Client)(nil).DeleteObjects), varargs...)
}

// GetObject mocks base method.
func (m *MockS3Client) GetObject(arg0 context.Context, arg1 *s3.GetObjectInput, arg2 ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockS3ClientMockRecorder) GetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3Client)(nil).GetObject), varargs...)
}

// GetObjectAttributes mocks base method.
func (m *MockS3Client) GetObjectAttributes(arg0 context.Context, arg1 *s3.GetObjectAttributesInput, arg2 ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectAttributes", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectAttributes indicates an expected call of GetObjectAttributes.
func (mr *MockS3ClientMockRecorder) GetObjectAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectAttributes", reflect.TypeOf((*MockS3Client)(nil).GetObjectAttributes), varargs...)
}

// ListObjectsV2 mocks base method.
func (m *MockS3Client) ListObjectsV2(arg0 context.Context, arg1 *s3.ListObjectsV2Input, arg2 ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObjectsV2", varargs...)
	ret0, _ := ret[0].(*s3.ListObjectsV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectsV2 indicates an expected call of ListObjectsV2.
func (mr *MockS3ClientMockRecorder) ListObjectsV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsV2", reflect.TypeOf((*MockS3Client)(nil).ListObjectsV2), varargs...)
}
