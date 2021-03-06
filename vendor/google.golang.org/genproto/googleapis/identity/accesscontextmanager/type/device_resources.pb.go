// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/identity/accesscontextmanager/type/device_resources.proto

package _type

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The encryption state of the device.
type DeviceEncryptionStatus int32

const (
	// The encryption status of the device is not specified or not known.
	DeviceEncryptionStatus_ENCRYPTION_UNSPECIFIED DeviceEncryptionStatus = 0
	// The device does not support encryption.
	DeviceEncryptionStatus_ENCRYPTION_UNSUPPORTED DeviceEncryptionStatus = 1
	// The device supports encryption, but is currently unencrypted.
	DeviceEncryptionStatus_UNENCRYPTED DeviceEncryptionStatus = 2
	// The device is encrypted.
	DeviceEncryptionStatus_ENCRYPTED DeviceEncryptionStatus = 3
)

var DeviceEncryptionStatus_name = map[int32]string{
	0: "ENCRYPTION_UNSPECIFIED",
	1: "ENCRYPTION_UNSUPPORTED",
	2: "UNENCRYPTED",
	3: "ENCRYPTED",
}

var DeviceEncryptionStatus_value = map[string]int32{
	"ENCRYPTION_UNSPECIFIED": 0,
	"ENCRYPTION_UNSUPPORTED": 1,
	"UNENCRYPTED":            2,
	"ENCRYPTED":              3,
}

func (x DeviceEncryptionStatus) String() string {
	return proto.EnumName(DeviceEncryptionStatus_name, int32(x))
}

func (DeviceEncryptionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08acdf01f36de536, []int{0}
}

// The operating system type of the device.
// Next id: 7
type OsType int32

const (
	// The operating system of the device is not specified or not known.
	OsType_OS_UNSPECIFIED OsType = 0
	// A desktop Mac operating system.
	OsType_DESKTOP_MAC OsType = 1
	// A desktop Windows operating system.
	OsType_DESKTOP_WINDOWS OsType = 2
	// A desktop Linux operating system.
	OsType_DESKTOP_LINUX OsType = 3
	// A desktop ChromeOS operating system.
	OsType_DESKTOP_CHROME_OS OsType = 6
	// An Android operating system.
	OsType_ANDROID OsType = 4
	// An iOS operating system.
	OsType_IOS OsType = 5
)

var OsType_name = map[int32]string{
	0: "OS_UNSPECIFIED",
	1: "DESKTOP_MAC",
	2: "DESKTOP_WINDOWS",
	3: "DESKTOP_LINUX",
	6: "DESKTOP_CHROME_OS",
	4: "ANDROID",
	5: "IOS",
}

var OsType_value = map[string]int32{
	"OS_UNSPECIFIED":    0,
	"DESKTOP_MAC":       1,
	"DESKTOP_WINDOWS":   2,
	"DESKTOP_LINUX":     3,
	"DESKTOP_CHROME_OS": 6,
	"ANDROID":           4,
	"IOS":               5,
}

func (x OsType) String() string {
	return proto.EnumName(OsType_name, int32(x))
}

func (OsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08acdf01f36de536, []int{1}
}

// The degree to which the device is managed by the Cloud organization.
type DeviceManagementLevel int32

const (
	// The device's management level is not specified or not known.
	DeviceManagementLevel_MANAGEMENT_UNSPECIFIED DeviceManagementLevel = 0
	// The device is not managed.
	DeviceManagementLevel_NONE DeviceManagementLevel = 1
	// Basic management is enabled, which is generally limited to monitoring and
	// wiping the corporate account.
	DeviceManagementLevel_BASIC DeviceManagementLevel = 2
	// Complete device management. This includes more thorough monitoring and the
	// ability to directly manage the device (such as remote wiping). This can be
	// enabled through the Android Enterprise Platform.
	DeviceManagementLevel_COMPLETE DeviceManagementLevel = 3
)

var DeviceManagementLevel_name = map[int32]string{
	0: "MANAGEMENT_UNSPECIFIED",
	1: "NONE",
	2: "BASIC",
	3: "COMPLETE",
}

var DeviceManagementLevel_value = map[string]int32{
	"MANAGEMENT_UNSPECIFIED": 0,
	"NONE":                   1,
	"BASIC":                  2,
	"COMPLETE":               3,
}

func (x DeviceManagementLevel) String() string {
	return proto.EnumName(DeviceManagementLevel_name, int32(x))
}

func (DeviceManagementLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08acdf01f36de536, []int{2}
}

func init() {
	proto.RegisterEnum("google.identity.accesscontextmanager.type.DeviceEncryptionStatus", DeviceEncryptionStatus_name, DeviceEncryptionStatus_value)
	proto.RegisterEnum("google.identity.accesscontextmanager.type.OsType", OsType_name, OsType_value)
	proto.RegisterEnum("google.identity.accesscontextmanager.type.DeviceManagementLevel", DeviceManagementLevel_name, DeviceManagementLevel_value)
}

func init() {
	proto.RegisterFile("google/identity/accesscontextmanager/type/device_resources.proto", fileDescriptor_08acdf01f36de536)
}

var fileDescriptor_08acdf01f36de536 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xba, 0x75, 0xeb, 0x29, 0xa3, 0x9e, 0xd1, 0x76, 0x31, 0xf1, 0x02, 0xab, 0x20,
	0xb9, 0xe0, 0x2e, 0xdc, 0x90, 0x26, 0x66, 0x44, 0x34, 0x76, 0xd4, 0xa4, 0x0c, 0x50, 0xa5, 0x2a,
	0x64, 0x56, 0x14, 0xa9, 0xb5, 0xa3, 0xc4, 0x9b, 0xe8, 0x2d, 0x8f, 0xc0, 0x63, 0xf0, 0x28, 0x3c,
	0x0a, 0x4f, 0x81, 0x9c, 0x3f, 0x42, 0xda, 0x7a, 0xd1, 0x9b, 0x48, 0x39, 0xdf, 0xef, 0x7c, 0xfe,
	0x8e, 0x8f, 0xe1, 0x7d, 0x2e, 0x65, 0xbe, 0xe1, 0x76, 0x71, 0xc7, 0x85, 0x2a, 0xd4, 0xce, 0x4e,
	0xb3, 0x8c, 0xd7, 0x75, 0x26, 0x85, 0xe2, 0x3f, 0xd4, 0x36, 0x15, 0x69, 0xce, 0x2b, 0x5b, 0xed,
	0x4a, 0x6e, 0xdf, 0xf1, 0x87, 0x22, 0xe3, 0xeb, 0x8a, 0xd7, 0xf2, 0xbe, 0xca, 0x78, 0x6d, 0x95,
	0x95, 0x54, 0x12, 0x5f, 0xb7, 0x0e, 0x56, 0xef, 0x60, 0xed, 0x73, 0xb0, 0xb4, 0xc3, 0xd5, 0xab,
	0xee, 0xb0, 0xb4, 0x2c, 0xec, 0x54, 0x08, 0xa9, 0x52, 0x55, 0x48, 0xd1, 0x19, 0x4d, 0x4b, 0xb8,
	0xf4, 0x9b, 0x23, 0x88, 0xc8, 0xaa, 0x5d, 0xa9, 0xa5, 0x58, 0xa5, 0xea, 0xbe, 0xc6, 0x57, 0x70,
	0x49, 0xa8, 0xb7, 0xf8, 0x1a, 0x25, 0x01, 0xa3, 0xeb, 0x25, 0x8d, 0x23, 0xe2, 0x05, 0x1f, 0x02,
	0xe2, 0xa3, 0x67, 0x4f, 0xb5, 0x65, 0x14, 0xb1, 0x45, 0x42, 0x7c, 0x64, 0xe0, 0x09, 0x8c, 0x97,
	0xb4, 0x53, 0x89, 0x8f, 0x4c, 0x7c, 0x06, 0xa3, 0xff, 0xbf, 0x83, 0xe9, 0x4f, 0x03, 0x86, 0xac,
	0x4e, 0x76, 0x25, 0xc7, 0x18, 0x5e, 0xb0, 0xf8, 0x91, 0xf5, 0x04, 0xc6, 0x3e, 0x89, 0x3f, 0x25,
	0x2c, 0x5a, 0x87, 0xae, 0x87, 0x0c, 0xfc, 0x12, 0x26, 0x7d, 0xe1, 0x36, 0xa0, 0x3e, 0xbb, 0x8d,
	0x91, 0x89, 0xcf, 0xe1, 0xac, 0x2f, 0xce, 0x03, 0xba, 0xfc, 0x82, 0x06, 0xf8, 0x02, 0xce, 0xfb,
	0x92, 0xf7, 0x71, 0xc1, 0x42, 0xb2, 0x66, 0x31, 0x1a, 0xe2, 0x31, 0x9c, 0xb8, 0xd4, 0x5f, 0xb0,
	0xc0, 0x47, 0x47, 0xf8, 0x04, 0x06, 0x01, 0x8b, 0xd1, 0xf1, 0xf4, 0x33, 0x5c, 0xb4, 0x63, 0x87,
	0xcd, 0x55, 0x6d, 0xb9, 0x50, 0x73, 0xfe, 0xc0, 0x37, 0x7a, 0xb2, 0xd0, 0xa5, 0xee, 0x0d, 0x09,
	0x09, 0x4d, 0x1e, 0x45, 0x3b, 0x85, 0x23, 0xca, 0x28, 0x41, 0x06, 0x1e, 0xc1, 0xf1, 0xcc, 0x8d,
	0x03, 0x0f, 0x99, 0xf8, 0x39, 0x9c, 0x7a, 0x2c, 0x8c, 0xe6, 0x24, 0x21, 0x68, 0x30, 0xfb, 0x65,
	0xc2, 0x9b, 0x4c, 0x6e, 0xad, 0x83, 0xd7, 0x33, 0x1b, 0xe9, 0x9b, 0x88, 0xf4, 0x2e, 0x22, 0xe3,
	0x5b, 0xd8, 0xf5, 0xe5, 0x72, 0x93, 0x8a, 0xdc, 0x92, 0x55, 0x6e, 0xe7, 0x5c, 0x34, 0x9b, 0xb2,
	0x5b, 0x29, 0x2d, 0x8b, 0xfa, 0x80, 0x77, 0xf3, 0x4e, 0x7f, 0x7e, 0x9b, 0xd7, 0x37, 0xad, 0x5f,
	0xd0, 0xe7, 0x70, 0x9b, 0x06, 0xaf, 0x6d, 0x08, 0xbb, 0x1c, 0x3a, 0xc1, 0x9f, 0x9e, 0x5d, 0xf5,
	0xec, 0x6a, 0x1f, 0xbb, 0xd2, 0xec, 0x5f, 0xf3, 0x75, 0xcb, 0x3a, 0x4e, 0x0f, 0x3b, 0xce, 0x3e,
	0xda, 0x71, 0x34, 0xfe, 0x7d, 0xd8, 0x0c, 0xf0, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53,
	0x00, 0xdc, 0x0a, 0xf7, 0x02, 0x00, 0x00,
}
