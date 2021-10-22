//----------------------------------------------------------------
//  Copyright (c) ThoughtWorks, Inc.
//  Licensed under the Apache License, Version 2.0
//  See LICENSE in the project root for license information.
//----------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: services.proto

package gauge_messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xf3, 0x10, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0c, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x12, 0x23, 0x2e, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x65,
	0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x18, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x53, 0x75, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x29, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x17, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x53, 0x70, 0x65, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x28, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x12, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x53, 0x70, 0x65, 0x63, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x53, 0x70, 0x65, 0x63, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x1b, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x16,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6b, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x65, 0x70, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a,
	0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x12, 0x22, 0x2e,
	0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x13, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x53, 0x74, 0x65, 0x70, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x13, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x53, 0x70, 0x65, 0x63, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x67, 0x61, 0x75, 0x67,
	0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67,
	0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x50, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x50, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x35, 0x2e, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65,
	0x47, 0x6c, 0x6f, 0x62, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x65, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x65, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2e, 0x2e, 0x67, 0x61, 0x75,
	0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6d, 0x70, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x49, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x75, 0x62, 0x12, 0x2d, 0x2e, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x75,
	0x62, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x61, 0x75,
	0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x69, 0x66, 0x66, 0x12, 0x4d, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x1f, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x4b, 0x69, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4b, 0x69, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xaf, 0x07, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x17, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x28,
	0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x62, 0x0a, 0x1b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x70, 0x65, 0x63, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c,
	0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x53, 0x70, 0x65, 0x63, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x6a, 0x0a, 0x1f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x69, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x62, 0x0a, 0x1b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x74, 0x65, 0x70, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c,
	0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x65, 0x70, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x5e, 0x0a, 0x19, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x74, 0x65,
	0x70, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x2a, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x66, 0x0a, 0x1d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x69, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5e, 0x0a, 0x19, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x70, 0x65, 0x63, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x56, 0x0a, 0x15, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x15,
	0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x04, 0x4b, 0x69, 0x6c, 0x6c, 0x12, 0x22, 0x2e,
	0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4b,
	0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x93, 0x01, 0x0a, 0x0a, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x73, 0x12, 0x1b, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x04, 0x4b,
	0x69, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x5c,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2f, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x61, 0x75,
	0x67, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0xaa, 0x02, 0x0e, 0x47, 0x61,
	0x75, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_services_proto_goTypes = []interface{}{
	(*StepValidateRequest)(nil),                   // 0: gauge.messages.StepValidateRequest
	(*SuiteDataStoreInitRequest)(nil),             // 1: gauge.messages.SuiteDataStoreInitRequest
	(*ExecutionStartingRequest)(nil),              // 2: gauge.messages.ExecutionStartingRequest
	(*SpecDataStoreInitRequest)(nil),              // 3: gauge.messages.SpecDataStoreInitRequest
	(*SpecExecutionStartingRequest)(nil),          // 4: gauge.messages.SpecExecutionStartingRequest
	(*ScenarioDataStoreInitRequest)(nil),          // 5: gauge.messages.ScenarioDataStoreInitRequest
	(*ScenarioExecutionStartingRequest)(nil),      // 6: gauge.messages.ScenarioExecutionStartingRequest
	(*StepExecutionStartingRequest)(nil),          // 7: gauge.messages.StepExecutionStartingRequest
	(*ExecuteStepRequest)(nil),                    // 8: gauge.messages.ExecuteStepRequest
	(*StepExecutionEndingRequest)(nil),            // 9: gauge.messages.StepExecutionEndingRequest
	(*ScenarioExecutionEndingRequest)(nil),        // 10: gauge.messages.ScenarioExecutionEndingRequest
	(*SpecExecutionEndingRequest)(nil),            // 11: gauge.messages.SpecExecutionEndingRequest
	(*ExecutionEndingRequest)(nil),                // 12: gauge.messages.ExecutionEndingRequest
	(*CacheFileRequest)(nil),                      // 13: gauge.messages.CacheFileRequest
	(*StepNameRequest)(nil),                       // 14: gauge.messages.StepNameRequest
	(*Empty)(nil),                                 // 15: gauge.messages.Empty
	(*StepNamesRequest)(nil),                      // 16: gauge.messages.StepNamesRequest
	(*StepPositionsRequest)(nil),                  // 17: gauge.messages.StepPositionsRequest
	(*StubImplementationCodeRequest)(nil),         // 18: gauge.messages.StubImplementationCodeRequest
	(*RefactorRequest)(nil),                       // 19: gauge.messages.RefactorRequest
	(*KillProcessRequest)(nil),                    // 20: gauge.messages.KillProcessRequest
	(*SuiteExecutionResult)(nil),                  // 21: gauge.messages.SuiteExecutionResult
	(*SpecDetails)(nil),                           // 22: gauge.messages.SpecDetails
	(*StepValidateResponse)(nil),                  // 23: gauge.messages.StepValidateResponse
	(*ExecutionStatusResponse)(nil),               // 24: gauge.messages.ExecutionStatusResponse
	(*StepNameResponse)(nil),                      // 25: gauge.messages.StepNameResponse
	(*ImplementationFileGlobPatternResponse)(nil), // 26: gauge.messages.ImplementationFileGlobPatternResponse
	(*StepNamesResponse)(nil),                     // 27: gauge.messages.StepNamesResponse
	(*StepPositionsResponse)(nil),                 // 28: gauge.messages.StepPositionsResponse
	(*ImplementationFileListResponse)(nil),        // 29: gauge.messages.ImplementationFileListResponse
	(*FileDiff)(nil),                              // 30: gauge.messages.FileDiff
	(*RefactorResponse)(nil),                      // 31: gauge.messages.RefactorResponse
}
var file_services_proto_depIdxs = []int32{
	0,  // 0: gauge.messages.Runner.ValidateStep:input_type -> gauge.messages.StepValidateRequest
	1,  // 1: gauge.messages.Runner.InitializeSuiteDataStore:input_type -> gauge.messages.SuiteDataStoreInitRequest
	2,  // 2: gauge.messages.Runner.StartExecution:input_type -> gauge.messages.ExecutionStartingRequest
	3,  // 3: gauge.messages.Runner.InitializeSpecDataStore:input_type -> gauge.messages.SpecDataStoreInitRequest
	4,  // 4: gauge.messages.Runner.StartSpecExecution:input_type -> gauge.messages.SpecExecutionStartingRequest
	5,  // 5: gauge.messages.Runner.InitializeScenarioDataStore:input_type -> gauge.messages.ScenarioDataStoreInitRequest
	6,  // 6: gauge.messages.Runner.StartScenarioExecution:input_type -> gauge.messages.ScenarioExecutionStartingRequest
	7,  // 7: gauge.messages.Runner.StartStepExecution:input_type -> gauge.messages.StepExecutionStartingRequest
	8,  // 8: gauge.messages.Runner.ExecuteStep:input_type -> gauge.messages.ExecuteStepRequest
	9,  // 9: gauge.messages.Runner.FinishStepExecution:input_type -> gauge.messages.StepExecutionEndingRequest
	10, // 10: gauge.messages.Runner.FinishScenarioExecution:input_type -> gauge.messages.ScenarioExecutionEndingRequest
	11, // 11: gauge.messages.Runner.FinishSpecExecution:input_type -> gauge.messages.SpecExecutionEndingRequest
	12, // 12: gauge.messages.Runner.FinishExecution:input_type -> gauge.messages.ExecutionEndingRequest
	13, // 13: gauge.messages.Runner.CacheFile:input_type -> gauge.messages.CacheFileRequest
	14, // 14: gauge.messages.Runner.GetStepName:input_type -> gauge.messages.StepNameRequest
	15, // 15: gauge.messages.Runner.GetGlobPatterns:input_type -> gauge.messages.Empty
	16, // 16: gauge.messages.Runner.GetStepNames:input_type -> gauge.messages.StepNamesRequest
	17, // 17: gauge.messages.Runner.GetStepPositions:input_type -> gauge.messages.StepPositionsRequest
	15, // 18: gauge.messages.Runner.GetImplementationFiles:input_type -> gauge.messages.Empty
	18, // 19: gauge.messages.Runner.ImplementStub:input_type -> gauge.messages.StubImplementationCodeRequest
	19, // 20: gauge.messages.Runner.Refactor:input_type -> gauge.messages.RefactorRequest
	20, // 21: gauge.messages.Runner.Kill:input_type -> gauge.messages.KillProcessRequest
	2,  // 22: gauge.messages.Reporter.NotifyExecutionStarting:input_type -> gauge.messages.ExecutionStartingRequest
	4,  // 23: gauge.messages.Reporter.NotifySpecExecutionStarting:input_type -> gauge.messages.SpecExecutionStartingRequest
	6,  // 24: gauge.messages.Reporter.NotifyScenarioExecutionStarting:input_type -> gauge.messages.ScenarioExecutionStartingRequest
	7,  // 25: gauge.messages.Reporter.NotifyStepExecutionStarting:input_type -> gauge.messages.StepExecutionStartingRequest
	9,  // 26: gauge.messages.Reporter.NotifyStepExecutionEnding:input_type -> gauge.messages.StepExecutionEndingRequest
	10, // 27: gauge.messages.Reporter.NotifyScenarioExecutionEnding:input_type -> gauge.messages.ScenarioExecutionEndingRequest
	11, // 28: gauge.messages.Reporter.NotifySpecExecutionEnding:input_type -> gauge.messages.SpecExecutionEndingRequest
	12, // 29: gauge.messages.Reporter.NotifyExecutionEnding:input_type -> gauge.messages.ExecutionEndingRequest
	21, // 30: gauge.messages.Reporter.NotifySuiteResult:input_type -> gauge.messages.SuiteExecutionResult
	20, // 31: gauge.messages.Reporter.Kill:input_type -> gauge.messages.KillProcessRequest
	22, // 32: gauge.messages.Documenter.GenerateDocs:input_type -> gauge.messages.SpecDetails
	20, // 33: gauge.messages.Documenter.Kill:input_type -> gauge.messages.KillProcessRequest
	23, // 34: gauge.messages.Runner.ValidateStep:output_type -> gauge.messages.StepValidateResponse
	24, // 35: gauge.messages.Runner.InitializeSuiteDataStore:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 36: gauge.messages.Runner.StartExecution:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 37: gauge.messages.Runner.InitializeSpecDataStore:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 38: gauge.messages.Runner.StartSpecExecution:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 39: gauge.messages.Runner.InitializeScenarioDataStore:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 40: gauge.messages.Runner.StartScenarioExecution:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 41: gauge.messages.Runner.StartStepExecution:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 42: gauge.messages.Runner.ExecuteStep:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 43: gauge.messages.Runner.FinishStepExecution:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 44: gauge.messages.Runner.FinishScenarioExecution:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 45: gauge.messages.Runner.FinishSpecExecution:output_type -> gauge.messages.ExecutionStatusResponse
	24, // 46: gauge.messages.Runner.FinishExecution:output_type -> gauge.messages.ExecutionStatusResponse
	15, // 47: gauge.messages.Runner.CacheFile:output_type -> gauge.messages.Empty
	25, // 48: gauge.messages.Runner.GetStepName:output_type -> gauge.messages.StepNameResponse
	26, // 49: gauge.messages.Runner.GetGlobPatterns:output_type -> gauge.messages.ImplementationFileGlobPatternResponse
	27, // 50: gauge.messages.Runner.GetStepNames:output_type -> gauge.messages.StepNamesResponse
	28, // 51: gauge.messages.Runner.GetStepPositions:output_type -> gauge.messages.StepPositionsResponse
	29, // 52: gauge.messages.Runner.GetImplementationFiles:output_type -> gauge.messages.ImplementationFileListResponse
	30, // 53: gauge.messages.Runner.ImplementStub:output_type -> gauge.messages.FileDiff
	31, // 54: gauge.messages.Runner.Refactor:output_type -> gauge.messages.RefactorResponse
	15, // 55: gauge.messages.Runner.Kill:output_type -> gauge.messages.Empty
	15, // 56: gauge.messages.Reporter.NotifyExecutionStarting:output_type -> gauge.messages.Empty
	15, // 57: gauge.messages.Reporter.NotifySpecExecutionStarting:output_type -> gauge.messages.Empty
	15, // 58: gauge.messages.Reporter.NotifyScenarioExecutionStarting:output_type -> gauge.messages.Empty
	15, // 59: gauge.messages.Reporter.NotifyStepExecutionStarting:output_type -> gauge.messages.Empty
	15, // 60: gauge.messages.Reporter.NotifyStepExecutionEnding:output_type -> gauge.messages.Empty
	15, // 61: gauge.messages.Reporter.NotifyScenarioExecutionEnding:output_type -> gauge.messages.Empty
	15, // 62: gauge.messages.Reporter.NotifySpecExecutionEnding:output_type -> gauge.messages.Empty
	15, // 63: gauge.messages.Reporter.NotifyExecutionEnding:output_type -> gauge.messages.Empty
	15, // 64: gauge.messages.Reporter.NotifySuiteResult:output_type -> gauge.messages.Empty
	15, // 65: gauge.messages.Reporter.Kill:output_type -> gauge.messages.Empty
	15, // 66: gauge.messages.Documenter.GenerateDocs:output_type -> gauge.messages.Empty
	15, // 67: gauge.messages.Documenter.Kill:output_type -> gauge.messages.Empty
	34, // [34:68] is the sub-list for method output_type
	0,  // [0:34] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	file_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
