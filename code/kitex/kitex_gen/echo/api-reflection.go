// Code generated by thriftgo (0.3.2). DO NOT EDIT.

package echo

import (
	"reflect"

	"github.com/cloudwego/thriftgo/thrift_reflection"
)

// IDL Name: api
// IDL Path: api.thrift

var file_api_thrift_go_types = []interface{}{
	(*EchoRequest)(nil),                     // Struct 0: echo.EchoRequest
	(*EchoResponse)(nil),                    // Struct 1: echo.EchoResponse
	(*EchoMultiBoolResponse)(nil),           // Struct 2: echo.EchoMultiBoolResponse
	(*EchoMultiByteResponse)(nil),           // Struct 3: echo.EchoMultiByteResponse
	(*EchoMultiInt16Response)(nil),          // Struct 4: echo.EchoMultiInt16Response
	(*EchoMultiInt32Response)(nil),          // Struct 5: echo.EchoMultiInt32Response
	(*EchoMultiInt64Response)(nil),          // Struct 6: echo.EchoMultiInt64Response
	(*EchoMultiFloatResponse)(nil),          // Struct 7: echo.EchoMultiFloatResponse
	(*EchoMultiDoubleResponse)(nil),         // Struct 8: echo.EchoMultiDoubleResponse
	(*EchoMultiStringResponse)(nil),         // Struct 9: echo.EchoMultiStringResponse
	(*EchoOptionalStructRequest)(nil),       // Struct 10: echo.EchoOptionalStructRequest
	(*EchoOptionalStructResponse)(nil),      // Struct 11: echo.EchoOptionalStructResponse
	(*EchoOptionalMultiBoolRequest)(nil),    // Struct 12: echo.EchoOptionalMultiBoolRequest
	(*EchoOptionalMultiInt32Request)(nil),   // Struct 13: echo.EchoOptionalMultiInt32Request
	(*EchoOptionalMultiStringRequest)(nil),  // Struct 14: echo.EchoOptionalMultiStringRequest
	(*EchoOptionalMultiBoolResponse)(nil),   // Struct 15: echo.EchoOptionalMultiBoolResponse
	(*EchoOptionalMultiInt32Response)(nil),  // Struct 16: echo.EchoOptionalMultiInt32Response
	(*EchoOptionalMultiStringResponse)(nil), // Struct 17: echo.EchoOptionalMultiStringResponse
	(*TweetType)(nil),                       // Enum 0: echo.TweetType
}
var file_api_thrift *thrift_reflection.FileDescriptor
var file_idl_api_rawDesc = []byte{
	0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xec, 0x5d, 0xdb, 0x73, 0xdb, 0x4c,
	0x15, 0xdf, 0xcf, 0x97, 0xd8, 0x89, 0x93, 0xb8, 0x9, 0x6d, 0x4a, 0xb, 0xa5, 0x66, 0x78, 0xe0,
	0xc5, 0x63, 0x13, 0xc7, 0x98, 0xcb, 0x1b, 0x6d, 0xc2, 0x40, 0xa7, 0x81, 0x4e, 0x1a, 0x86, 0x61,
	0x18, 0x98, 0xd9, 0x38, 0xdb, 0x58, 0xd4, 0x91, 0x5c, 0x4b, 0xe, 0xf8, 0xbf, 0x67, 0x76, 0x75,
	0xb3, 0x6c, 0x6b, 0x6f, 0xda, 0x95, 0x64, 0x55, 0xdf, 0xc3, 0xe7, 0x4e, 0x2c, 0xed, 0x9e, 0xf3,
	0x3b, 0x3f, 0xad, 0x8e, 0xf6, 0xfc, 0x8e, 0xdc, 0x2, 0x3f, 0x0, 0x0, 0xe, 0xe0, 0xcc, 0xe8,
	0x39, 0x93, 0xb9, 0xf1, 0xc5, 0x39, 0x2, 0x95, 0x56, 0xb, 0x0, 0xf2, 0xe7, 0xda, 0x7f, 0xe0,
	0x13, 0x4, 0x0, 0xb4, 0xf0, 0x67, 0xf0, 0x7d, 0x35, 0xf8, 0xbe, 0xf2, 0x60, 0xe1, 0xa3, 0xd0,
	0x78, 0x62, 0xb5, 0x41, 0xed, 0x10, 0xff, 0x75, 0x63, 0xbc, 0x16, 0xa8, 0xe0, 0x11, 0x6e, 0x91,
	0xed, 0x7c, 0x46, 0xf3, 0x27, 0x63, 0x8c, 0xda, 0xa0, 0x8a, 0xf, 0xfd, 0x77, 0xcc, 0xa1, 0x8d,
	0xab, 0xf1, 0xc4, 0xfa, 0xb3, 0xe9, 0x1c, 0x82, 0x6a, 0xcc, 0x11, 0x55, 0xe3, 0x62, 0x0, 0x18,
	0x13, 0x56, 0xe7, 0xe8, 0x1b, 0x63, 0x84, 0x16, 0xa8, 0xe1, 0xe9, 0x2e, 0xd1, 0x17, 0xb8, 0x98,
	0x3a, 0x4d, 0x50, 0xc7, 0xc3, 0x1d, 0x81, 0x46, 0xab, 0xd, 0xb0, 0xc5, 0xa0, 0x89, 0x3f, 0xc0,
	0x11, 0xa8, 0xfb, 0x7f, 0xd8, 0xc3, 0x1f, 0x6d, 0xd0, 0xc0, 0xf3, 0x82, 0xa, 0xfe, 0x3e, 0x66,
	0xf8, 0x26, 0x76, 0xe1, 0x9d, 0x65, 0x4d, 0xe3, 0x2d, 0xa8, 0xdd, 0x59, 0xd6, 0x34, 0x99, 0x13,
	0xee, 0x10, 0x52, 0x5e, 0x1c, 0xf5, 0xfb, 0x9d, 0x3b, 0x68, 0xa3, 0x8e, 0xb3, 0x9c, 0x21, 0x9b,
	0xdf, 0xa5, 0xa5, 0x83, 0x68, 0xf6, 0x2c, 0x1d, 0x94, 0xd4, 0x25, 0x3c, 0x84, 0xbe, 0xc0, 0xec,
	0x7b, 0xdc, 0x3a, 0x1f, 0xd1, 0xb8, 0x71, 0x3e, 0x4a, 0xca, 0xae, 0xf3, 0x51, 0xa, 0x4e, 0x5c,
	0xc, 0x76, 0xf8, 0x12, 0xf1, 0x9d, 0x18, 0xd, 0x69, 0x26, 0x8c, 0x86, 0x49, 0x9d, 0x18, 0xd,
	0x75, 0x3b, 0xf1, 0xc7, 0xa9, 0x5, 0x29, 0x8b, 0xd5, 0xde, 0xbd, 0xb5, 0xb8, 0x9b, 0x26, 0xbc,
	0x2e, 0xfc, 0x41, 0xc4, 0x5c, 0xc1, 0x43, 0x3d, 0x9b, 0x20, 0xdb, 0x36, 0xa0, 0xd9, 0x83, 0xf3,
	0x7, 0xfb, 0x76, 0x39, 0x43, 0xfe, 0xda, 0x5d, 0x27, 0x76, 0xb, 0x38, 0x7b, 0x80, 0x9d, 0xbd,
	0x24, 0x76, 0xe4, 0xd2, 0x5b, 0x91, 0xc0, 0x11, 0x5f, 0x3e, 0x3b, 0x73, 0xc3, 0x7c, 0xa0, 0x98,
	0x61, 0x93, 0x3, 0x12, 0xfa, 0xe2, 0xd, 0xa2, 0xd9, 0x97, 0x77, 0x86, 0x9, 0xe7, 0x4b, 0x8a,
	0x19, 0x77, 0xe4, 0x80, 0x84, 0xbe, 0x78, 0x83, 0x48, 0xf9, 0xf2, 0xac, 0xdf, 0xef, 0xd8, 0x33,
	0x34, 0x36, 0xe0, 0x94, 0xff, 0xae, 0x53, 0xc3, 0xbe, 0xc5, 0x1b, 0x74, 0x88, 0xbf, 0xbe, 0x41,
	0xf6, 0xcc, 0x32, 0xed, 0x84, 0x9c, 0x6b, 0xb9, 0x43, 0x7d, 0x5b, 0x20, 0xdb, 0x91, 0x73, 0xf0,
	0x0, 0x3b, 0xe8, 0xcc, 0x17, 0x63, 0x87, 0xcb, 0xb5, 0x43, 0x3f, 0x47, 0xf8, 0x68, 0xd8, 0x94,
	0xe5, 0xa3, 0x36, 0x25, 0x5f, 0xd7, 0xa8, 0x39, 0x40, 0xc2, 0x9b, 0x2e, 0xd7, 0x14, 0x52, 0x98,
	0x9c, 0xf4, 0xfb, 0x9d, 0xb1, 0x65, 0x3a, 0xd0, 0x30, 0xd1, 0xbc, 0x83, 0xe7, 0xe1, 0x87, 0x66,
	0xe9, 0xa0, 0x84, 0xd0, 0xe0, 0x5c, 0x42, 0x33, 0x34, 0x64, 0xa, 0x7d, 0xd7, 0xf6, 0x51, 0x90,
	0xaf, 0x24, 0x82, 0x82, 0x24, 0x24, 0x5a, 0x91, 0x70, 0x67, 0xd0, 0xf, 0xc4, 0xc5, 0x20, 0x21,
	0x10, 0x17, 0x3, 0xdd, 0x40, 0xe0, 0x19, 0xf4, 0x3, 0x31, 0x1a, 0x26, 0x4, 0x62, 0x34, 0xd4,
	0xd, 0x4, 0x9e, 0x41, 0x33, 0x10, 0x24, 0x87, 0x49, 0x4, 0x84, 0x9f, 0x6b, 0x68, 0xc5, 0x22,
	0x98, 0x44, 0x1f, 0x1c, 0xc7, 0x61, 0x76, 0x56, 0xe2, 0xe1, 0xe3, 0xe1, 0x66, 0x78, 0xc9, 0xf0,
	0xf0, 0xf2, 0x37, 0xbd, 0x78, 0xf8, 0x93, 0x68, 0xc6, 0xc3, 0xcd, 0x12, 0x93, 0xe1, 0xe1, 0xe5,
	0x80, 0x7a, 0xf1, 0xf0, 0x27, 0xd1, 0x87, 0xc7, 0x33, 0x3f, 0xfd, 0x1a, 0xe0, 0xff, 0x5d, 0xc3,
	0x19, 0xe5, 0x21, 0xf2, 0x91, 0xf6, 0xad, 0x9b, 0x1e, 0xe9, 0xcd, 0xd0, 0x94, 0x58, 0x20, 0x9d,
	0xb5, 0x87, 0x9, 0xdc, 0x23, 0x9c, 0x9, 0x62, 0xbb, 0x74, 0x90, 0x56, 0x6c, 0x93, 0xa7, 0x78,
	0x4a, 0x2c, 0xd0, 0xc7, 0xd3, 0x93, 0x0, 0x4b, 0x92, 0x6, 0x6a, 0x4, 0x53, 0x41, 0x92, 0xa8,
	0xc2, 0x80, 0x94, 0xa0, 0xbc, 0x18, 0x68, 0x85, 0x32, 0x71, 0x9a, 0xa9, 0xc2, 0x80, 0x94, 0xa0,
	0x1c, 0xd, 0xb5, 0x42, 0x99, 0x38, 0x51, 0x55, 0x61, 0x40, 0x1a, 0x50, 0x92, 0x64, 0x56, 0x23,
	0x94, 0x6a, 0x52, 0x3b, 0x45, 0x36, 0xe8, 0x3, 0xf4, 0x34, 0x0, 0xd4, 0x4d, 0x87, 0x4b, 0x44,
	0xd5, 0x21, 0xea, 0x26, 0xd4, 0x3a, 0x11, 0x55, 0x92, 0x6e, 0x2b, 0xb2, 0x21, 0xd, 0x44, 0xdd,
	0x94, 0x5c, 0x27, 0xa2, 0x4a, 0x12, 0x76, 0x45, 0x36, 0xe8, 0x43, 0xf4, 0x79, 0x24, 0x9f, 0xc7,
	0x8f, 0x38, 0x3a, 0xf3, 0x4e, 0xae, 0x1d, 0xd3, 0x4c, 0xf3, 0x52, 0x3e, 0xb, 0x53, 0x89, 0x87,
	0xb7, 0x91, 0x9b, 0x65, 0x3c, 0x48, 0x92, 0x9e, 0xeb, 0x78, 0xb8, 0x16, 0xea, 0x8b, 0xc7, 0x8b,
	0xe8, 0x73, 0x44, 0xb6, 0x1, 0x71, 0x33, 0xfd, 0x1c, 0xc7, 0xc3, 0x33, 0x30, 0xa5, 0x70, 0xb8,
	0x9b, 0xda, 0x59, 0x86, 0x3, 0x3f, 0x2d, 0xe4, 0x3a, 0x1c, 0xc4, 0xc0, 0x94, 0xc2, 0xe1, 0x6e,
	0xad, 0x67, 0x19, 0xe, 0xfc, 0xc4, 0x91, 0xeb, 0x70, 0x10, 0x3, 0xd3, 0x8, 0x47, 0xb0, 0xc1,
	0x9f, 0x5d, 0x38, 0x82, 0x4, 0x3b, 0xc7, 0x11, 0x9, 0x6d, 0xd4, 0x17, 0x94, 0xb3, 0xb5, 0xe7,
	0xaa, 0x32, 0x2a, 0xf9, 0x8a, 0x4a, 0x58, 0xec, 0xc8, 0x30, 0x2a, 0xfe, 0xa3, 0x53, 0x9e, 0xa3,
	0x12, 0xd8, 0x98, 0x46, 0x54, 0xc2, 0x92, 0x4b, 0x86, 0x51, 0xf1, 0x1f, 0xbf, 0xf2, 0x1c, 0x95,
	0xc0, 0x46, 0xcd, 0x25, 0xe3, 0xeb, 0xc5, 0xd4, 0x31, 0xe8, 0xda, 0xdc, 0x17, 0x91, 0xe3, 0xd6,
	0xf5, 0x45, 0x71, 0x67, 0x35, 0xee, 0xa0, 0x8d, 0x6e, 0x54, 0x9, 0x76, 0xe3, 0x26, 0xc1, 0x28,
	0xd2, 0x27, 0x91, 0x56, 0xf3, 0x54, 0x78, 0x8d, 0xd8, 0x7b, 0x84, 0x33, 0xaa, 0xd, 0xda, 0xea,
	0x55, 0x55, 0xd5, 0x3c, 0xa0, 0xa, 0x9a, 0x5f, 0x44, 0x8e, 0x53, 0xc9, 0x3, 0x7e, 0x95, 0xb3,
	0x3e, 0x1e, 0xc4, 0x16, 0xae, 0xd2, 0xe3, 0x1, 0x36, 0x41, 0xca, 0x44, 0x35, 0x3c, 0x38, 0xe,
	0xe2, 0xcb, 0x90, 0x84, 0x9f, 0x45, 0xf, 0x54, 0xc7, 0x4, 0x1, 0xa5, 0xb8, 0x2e, 0x22, 0xc4,
	0x57, 0xdd, 0xd2, 0xe2, 0x1, 0xb1, 0x40, 0xc6, 0x40, 0xd, 0x2c, 0xa0, 0x69, 0xea, 0xcf, 0xa2,
	0x7, 0xaa, 0x64, 0x1, 0xb7, 0xd4, 0x5e, 0x1f, 0xb, 0xe2, 0xa, 0x86, 0xe9, 0xb1, 0xe0, 0x62,
	0x20, 0x55, 0xd1, 0xd4, 0xc0, 0x2, 0x5a, 0x53, 0xc2, 0x59, 0xf4, 0x40, 0x95, 0x2c, 0xe0, 0xee,
	0x55, 0xd0, 0xc7, 0x82, 0xb8, 0x5a, 0x67, 0x7a, 0x2c, 0x18, 0xd, 0xa5, 0x8a, 0xb1, 0xaa, 0x59,
	0xc0, 0xe8, 0xea, 0x38, 0x8b, 0x1e, 0xa8, 0x8e, 0x5, 0x62, 0xbd, 0xf, 0xba, 0x88, 0x40, 0x2d,
	0x2a, 0xa6, 0xc5, 0x5, 0xdf, 0x8, 0x49, 0x33, 0xe3, 0x19, 0xf1, 0x3, 0xad, 0x2f, 0xa6, 0x41,
	0x2, 0xda, 0xed, 0xa, 0xb0, 0xa6, 0x1d, 0x90, 0x81, 0xd5, 0x1e, 0xf3, 0x72, 0xed, 0xc8, 0x92,
	0x37, 0xbb, 0xc4, 0x1b, 0x91, 0x95, 0x24, 0xe4, 0x4, 0xab, 0xcd, 0xe8, 0xe5, 0xda, 0x91, 0xa,
	0x39, 0x21, 0xd4, 0x7b, 0xa4, 0x8d, 0x13, 0xb4, 0x72, 0x7a, 0x6a, 0x9c, 0xf0, 0x8c, 0x90, 0x34,
	0x53, 0xd, 0x27, 0xdc, 0xbe, 0x16, 0x68, 0xa3, 0x5c, 0xb6, 0x6, 0xd3, 0x57, 0x45, 0x3c, 0x20,
	0x82, 0xa6, 0xeb, 0xed, 0xab, 0x7e, 0xbf, 0xf3, 0x88, 0x9c, 0x89, 0x75, 0xdf, 0xc1, 0x7, 0x76,
	0xa0, 0x69, 0x5a, 0xe, 0x74, 0xc, 0xcb, 0x14, 0x83, 0x21, 0x8f, 0xed, 0xc4, 0x54, 0x18, 0xc8,
	0x80, 0xa2, 0x3b, 0xd, 0xd8, 0xd5, 0x3c, 0xf6, 0x1c, 0x53, 0x3d, 0xad, 0xdb, 0x13, 0x6b, 0x2e,
	0xd2, 0x1e, 0xba, 0xea, 0x6a, 0xee, 0x3a, 0x93, 0xa9, 0xae, 0x56, 0xd, 0x53, 0xd6, 0xd1, 0xdc,
	0x75, 0x2f, 0xd3, 0xd9, 0x3b, 0xb5, 0xcc, 0x7, 0x19, 0x4f, 0x77, 0xb5, 0xc5, 0xf9, 0x8b, 0x60,
	0x8b, 0xf3, 0xb1, 0xef, 0x70, 0x8e, 0xdb, 0x9c, 0xa9, 0x1e, 0x7b, 0x43, 0xa, 0xf7, 0x41, 0x78,
	0xf7, 0xa4, 0x42, 0xb6, 0xa2, 0x52, 0x1, 0xdb, 0xf7, 0xee, 0x6b, 0xff, 0xfc, 0x97, 0x14, 0x66,
	0x45, 0xed, 0x51, 0xa5, 0x93, 0xc, 0x8f, 0x28, 0x4, 0xd8, 0x49, 0xe4, 0x36, 0x58, 0xc0, 0x56,
	0x56, 0x7a, 0xee, 0x44, 0x6e, 0xa5, 0xd2, 0x80, 0x15, 0xb2, 0xe5, 0x95, 0xbe, 0x6e, 0x1b, 0xa6,
	0x3c, 0x5c, 0x85, 0x6c, 0x8c, 0xa5, 0x5f, 0x8f, 0xf8, 0xb6, 0x2e, 0x87, 0x57, 0xb1, 0xfb, 0x67,
	0xe9, 0x57, 0x25, 0x49, 0xe, 0x84, 0x60, 0x3b, 0x8d, 0xa6, 0x7, 0xdf, 0x27, 0x6e, 0x4d, 0x77,
	0x4c, 0x21, 0xe0, 0xa2, 0xe2, 0x6c, 0x8c, 0x60, 0xd9, 0x70, 0x29, 0x4, 0xfa, 0x9, 0x79, 0x1,
	0xdb, 0xc2, 0x31, 0xa6, 0xbd, 0x3f, 0x41, 0x7b, 0x72, 0xd, 0x67, 0xd2, 0x52, 0x6c, 0xed, 0xe8,
	0xef, 0x64, 0x4b, 0xa6, 0x3e, 0xf4, 0xd7, 0x84, 0xd7, 0x9a, 0xe1, 0xdf, 0xc9, 0x26, 0xce, 0x94,
	0xc0, 0xbf, 0x18, 0x68, 0x7, 0x7f, 0x7, 0xdb, 0x3e, 0x53, 0x2, 0x7f, 0x34, 0xd4, 0xe, 0xfe,
	0xe, 0x36, 0x8a, 0xa6, 0x1, 0x3e, 0xc9, 0xf1, 0x34, 0x83, 0xbf, 0xc3, 0x8d, 0x90, 0xfa, 0x42,
	0xb0, 0x2e, 0x98, 0x2e, 0x63, 0x90, 0x7e, 0xc, 0x4e, 0x42, 0x9, 0x1f, 0xb3, 0xee, 0x52, 0xca,
	0x3e, 0x77, 0x4f, 0xf6, 0x49, 0x65, 0xce, 0x2f, 0xbc, 0xcd, 0xbd, 0x6e, 0xb0, 0xc9, 0xd7, 0x55,
	0xc8, 0xa5, 0x52, 0x3a, 0x5a, 0x30, 0xe9, 0x28, 0x95, 0x4b, 0x6f, 0xf0, 0x88, 0x5d, 0x77, 0xe7,
	0x33, 0x11, 0x8b, 0x4e, 0x23, 0x2c, 0x2a, 0x85, 0xa7, 0xc5, 0x12, 0x9e, 0x52, 0x39, 0xf4, 0x96,
	0x6c, 0x4, 0x77, 0xbd, 0xed, 0x60, 0xb5, 0x2c, 0x2a, 0x85, 0xab, 0x5, 0x12, 0xae, 0x52, 0x59,
	0xf4, 0x13, 0xc3, 0x74, 0xba, 0x64, 0x87, 0x5c, 0x2d, 0x83, 0x4a, 0xd1, 0x6b, 0x81, 0x44, 0xaf,
	0xf4, 0x7b, 0xd9, 0xd4, 0x32, 0x1f, 0xba, 0x6e, 0xd5, 0x40, 0x21, 0x87, 0x4a, 0xc9, 0x6c, 0x71,
	0xa5, 0x8f, 0xf4, 0xfb, 0x1a, 0x29, 0xa5, 0x74, 0xbd, 0x82, 0x4a, 0x22, 0x46, 0xfd, 0x28, 0xc2,
	0xa8, 0x52, 0x4e, 0xfb, 0xdd, 0x72, 0xea, 0xe7, 0xee, 0x98, 0x5d, 0xbf, 0xda, 0x94, 0x88, 0x55,
	0xe4, 0xe5, 0xee, 0xd7, 0x44, 0xb0, 0xf8, 0x7, 0xb6, 0x46, 0x34, 0x13, 0xe5, 0xe5, 0xf1, 0x7,
	0xf8, 0x4, 0x5d, 0x13, 0xff, 0x2, 0x1f, 0x3, 0x18, 0xe, 0x42, 0xcb, 0x37, 0xa4, 0x97, 0x26,
	0x7c, 0x44, 0x82, 0xd2, 0xcb, 0x15, 0x20, 0xde, 0x69, 0x6, 0x42, 0x52, 0xa5, 0xc7, 0x8b, 0x83,
	0x68, 0xdc, 0xdf, 0xe7, 0xcf, 0xdd, 0xa, 0x4b, 0x94, 0xa8, 0x11, 0x8e, 0x4b, 0x6e, 0x38, 0x2a,
	0x71, 0x44, 0x9f, 0xa3, 0x6f, 0xe7, 0x5a, 0x37, 0xbd, 0xf0, 0xc, 0x2c, 0x21, 0x29, 0xdf, 0xc2,
	0xa8, 0x8d, 0x60, 0x44, 0x9, 0xf6, 0xd7, 0x19, 0xbe, 0xfc, 0xe0, 0x34, 0x97, 0xaa, 0x6e, 0x4f,
	0x9f, 0x7e, 0xd8, 0xef, 0x77, 0xf0, 0xa, 0xda, 0x31, 0x17, 0xd3, 0x29, 0xf7, 0x66, 0x97, 0xef,
	0x59, 0x1e, 0x25, 0xbd, 0xc2, 0xaf, 0xdb, 0xf3, 0x9d, 0x29, 0xc8, 0xcf, 0xbe, 0x3c, 0x5f, 0xa7,
	0x5e, 0x21, 0xc5, 0x9b, 0xc2, 0x6f, 0xd, 0x8a, 0x50, 0xb6, 0x80, 0xc2, 0x39, 0xe1, 0xb7, 0x90,
	0x44, 0x69, 0x5f, 0xbe, 0xb, 0x1f, 0x0, 0xf0, 0xe3, 0xf5, 0x4b, 0xa7, 0x7c, 0x7, 0x7c, 0x62,
	0x4c, 0x5f, 0x6d, 0x60, 0x5a, 0xbe, 0x64, 0x3b, 0x21, 0xa4, 0xaf, 0x37, 0x20, 0x2d, 0x5f, 0xbf,
	0xab, 0x27, 0x1f, 0x58, 0x8c, 0x29, 0x8b, 0xe2, 0xeb, 0xcd, 0x83, 0xd5, 0xfc, 0xe4, 0xd8, 0xab,
	0x6d, 0x3, 0x27, 0xf8, 0x1, 0x32, 0x2e, 0xd7, 0x7f, 0xba, 0x3a, 0xe9, 0x4a, 0xb1, 0x9b, 0xcc,
	0xab, 0x35, 0x7b, 0xa5, 0xce, 0xac, 0xd1, 0xe3, 0x37, 0x1b, 0xf3, 0x7a, 0x55, 0x10, 0x86, 0xcb,
	0xc9, 0xb3, 0x5a, 0xfa, 0xcc, 0x1a, 0x5d, 0xfe, 0xd9, 0xc6, 0xc4, 0x7e, 0x6b, 0x31, 0xc3, 0x67,
	0x25, 0xb9, 0x2f, 0x63, 0xf6, 0x54, 0x43, 0xbd, 0xaa, 0xe4, 0x10, 0x9, 0xd4, 0x36, 0x5, 0x48,
	0xea, 0xf, 0x6d, 0x92, 0xb1, 0x8e, 0x14, 0xf9, 0x44, 0xc2, 0xb4, 0xb5, 0x3a, 0x98, 0xcb, 0x7,
	0xba, 0xb7, 0xb1, 0x14, 0x63, 0x79, 0xcd, 0x3a, 0x33, 0x37, 0x8f, 0x7d, 0x47, 0xa0, 0x16, 0xec,
	0x59, 0x1c, 0x7d, 0x80, 0x4f, 0xf0, 0xfd, 0x14, 0xda, 0xf6, 0xea, 0x96, 0xc5, 0x2f, 0xad, 0xf9,
	0x43, 0xf, 0xce, 0xe0, 0x78, 0x82, 0x7a, 0xf7, 0x8b, 0xbb, 0x3b, 0xab, 0xe7, 0x20, 0xdb, 0xb1,
	0x7b, 0xd8, 0xa8, 0xbf, 0xd9, 0x68, 0xfe, 0x69, 0x6e, 0x3d, 0x19, 0xf7, 0x68, 0xde, 0x22, 0xd3,
	0xe3, 0xe1, 0xeb, 0x78, 0xf8, 0x53, 0xda, 0xee, 0x90, 0x77, 0x8d, 0x7a, 0x3f, 0xd1, 0x1d, 0x7,
	0x42, 0xdd, 0xe0, 0xd8, 0x10, 0x70, 0x31, 0x68, 0xe2, 0x11, 0x8d, 0x39, 0xba, 0x8f, 0x7, 0x81,
	0xe5, 0xe7, 0xd9, 0x57, 0xc3, 0x41, 0xff, 0xeb, 0xa1, 0xf1, 0xc4, 0xea, 0xad, 0xd8, 0xe8, 0xbb,
	0x15, 0x63, 0x43, 0xe4, 0xb7, 0x41, 0x73, 0xe5, 0xcf, 0xcb, 0xd, 0x7f, 0x5c, 0x23, 0x19, 0xe,
	0x6d, 0x97, 0xa7, 0x79, 0x9e, 0xc5, 0x59, 0xde, 0x74, 0x2b, 0x1d, 0x36, 0x33, 0x1f, 0xe3, 0xf4,
	0x2e, 0x6e, 0x1a, 0xb7, 0xd2, 0x41, 0x9d, 0x46, 0x60, 0x77, 0x60, 0xcd, 0xc, 0xee, 0x4a, 0x47,
	0x83, 0x54, 0x3a, 0x68, 0x56, 0x28, 0x7c, 0x8a, 0x5a, 0x33, 0x72, 0x5b, 0x9d, 0x83, 0xc5, 0x84,
	0x5f, 0x51, 0xae, 0xe0, 0xad, 0xd1, 0xe6, 0xa7, 0xc8, 0x8a, 0xea, 0x4c, 0x5, 0x45, 0x42, 0xd5,
	0x59, 0x96, 0x14, 0x59, 0x11, 0x75, 0x65, 0x47, 0x11, 0x1, 0xe5, 0x59, 0x6a, 0x14, 0x59, 0x89,
	0x36, 0x83, 0x22, 0x31, 0x8a, 0xb2, 0xe4, 0x1c, 0x59, 0x51, 0x94, 0x65, 0x47, 0x91, 0x55, 0xc9,
	0x56, 0x56, 0xc, 0x11, 0x50, 0x95, 0x29, 0xe0, 0xc7, 0x39, 0xf, 0x3f, 0x22, 0xa1, 0x16, 0x21,
	0x48, 0x98, 0xe, 0xaa, 0x20, 0x8, 0xff, 0x4d, 0x54, 0x23, 0x41, 0xc2, 0x3d, 0x9b, 0xec, 0x8,
	0xc2, 0x2d, 0x18, 0x4b, 0x91, 0x20, 0x61, 0xa8, 0x45, 0x8, 0x12, 0x6a, 0xc1, 0x54, 0x10, 0x24,
	0xd0, 0x82, 0x65, 0x49, 0x90, 0x50, 0x6c, 0x95, 0x1d, 0x41, 0xb8, 0xf5, 0x60, 0x29, 0x12, 0x24,
	0xc, 0x35, 0x37, 0x41, 0x22, 0x42, 0xaf, 0xe4, 0x4, 0x89, 0x8a, 0x72, 0xb2, 0xe3, 0xc8, 0x9a,
	0xde, 0x25, 0x2b, 0x9a, 0x8, 0xb, 0x73, 0xd2, 0x22, 0x4b, 0x24, 0xec, 0xc, 0xb2, 0xc4, 0x49,
	0xb8, 0x4a, 0xb6, 0x14, 0x80, 0x2d, 0x3, 0x1e, 0xb6, 0x44, 0xe3, 0xce, 0x4d, 0x97, 0xe8, 0xb6,
	0x8d, 0x2, 0xba, 0x44, 0xb6, 0x6e, 0x32, 0xa4, 0x4b, 0xb4, 0xfc, 0x91, 0x19, 0x5d, 0x44, 0x5f,
	0x98, 0x98, 0x16, 0x5d, 0xa2, 0x71, 0x67, 0xd0, 0x25, 0xbe, 0xc4, 0xc2, 0x20, 0xc, 0xff, 0x9e,
	0x6e, 0xd3, 0x1f, 0x5e, 0x90, 0x29, 0xfb, 0xce, 0x7f, 0x11, 0x72, 0x6e, 0x97, 0x33, 0xca, 0x66,
	0xe5, 0xfe, 0xad, 0x7f, 0xcc, 0xf6, 0xc9, 0x2a, 0x87, 0x60, 0xaf, 0x49, 0x4e, 0xdd, 0xab, 0x91,
	0x13, 0xc8, 0x7f, 0x7, 0xa0, 0xea, 0xff, 0xd3, 0x3d, 0xb, 0x54, 0x40, 0xdd, 0x33, 0xa7, 0x1d,
	0x8c, 0xd8, 0xbb, 0xfd, 0xfb, 0xd5, 0xd5, 0x2d, 0x50, 0x24, 0xc6, 0xa2, 0x20, 0x22, 0xc3, 0x84,
	0x21, 0x83, 0x9, 0x5b, 0x43, 0xca, 0xe0, 0x2, 0xa5, 0x8e, 0xe7, 0x91, 0x81, 0x22, 0x75, 0xe3,
	0xde, 0x40, 0x2b, 0x18, 0x1b, 0xd8, 0xa1, 0xfa, 0xb5, 0x60, 0xa8, 0xb8, 0xae, 0x5b, 0x6a, 0xad,
	0xd0, 0x8b, 0x56, 0xdc, 0xba, 0x84, 0xd7, 0x7a, 0x63, 0xcc, 0xdd, 0x91, 0x2b, 0x1b, 0xb1, 0xc6,
	0xc, 0x8e, 0xbf, 0xca, 0xce, 0xc2, 0xbf, 0x8c, 0x2b, 0x6d, 0xfc, 0x65, 0x5e, 0x99, 0x71, 0x4b,
	0x7d, 0x8a, 0x9d, 0xbf, 0x6b, 0x36, 0xd6, 0x24, 0x28, 0xf9, 0x1b, 0x4e, 0x4a, 0xae, 0x13, 0x8b,
	0x41, 0x4a, 0x7a, 0x39, 0x37, 0x39, 0x2b, 0x57, 0x37, 0x48, 0xb4, 0x91, 0x32, 0x7e, 0x92, 0xf4,
	0x38, 0x19, 0xd9, 0xe1, 0xc8, 0x27, 0x25, 0x29, 0x26, 0xca, 0x30, 0xf2, 0xb7, 0x22, 0x8c, 0x5c,
	0x65, 0x15, 0x83, 0x92, 0x8c, 0x52, 0x3b, 0x23, 0xc9, 0x11, 0x7d, 0x83, 0xb8, 0x34, 0x2b, 0x15,
	0xbf, 0x42, 0x5c, 0x96, 0xb8, 0x9a, 0x39, 0x43, 0xb7, 0x52, 0x26, 0xd, 0xfa, 0x9d, 0x8, 0x6d,
	0x22, 0x91, 0x17, 0x5e, 0xca, 0xb6, 0xd4, 0x14, 0xe3, 0xfc, 0xdc, 0xf7, 0x96, 0x32, 0xdd, 0x39,
	0x51, 0xd3, 0x5d, 0xcc, 0x24, 0xa7, 0xe1, 0x66, 0x85, 0xea, 0xda, 0xa5, 0xe4, 0x82, 0x96, 0x6e,
	0xed, 0x32, 0xe5, 0x25, 0x4d, 0xa0, 0x86, 0xc9, 0x90, 0xa5, 0x28, 0xe0, 0xa6, 0x82, 0xfb, 0x2c,
	0x7, 0x35, 0x15, 0xdc, 0x68, 0x15, 0x97, 0x3b, 0x72, 0x4a, 0x4c, 0xc5, 0xb7, 0x5a, 0xa1, 0x35,
	0x53, 0xa4, 0xee, 0xc1, 0x52, 0xe, 0x29, 0xdf, 0x82, 0x92, 0xe5, 0xa6, 0xf2, 0x2d, 0x28, 0xd9,
	0x3c, 0x51, 0x33, 0x71, 0xd4, 0xdf, 0x6f, 0x7f, 0x2f, 0x73, 0xbf, 0x8d, 0x92, 0xa7, 0xd, 0xf6,
	0xe, 0x41, 0xa8, 0xe5, 0x8a, 0x53, 0x1c, 0x85, 0xcf, 0xf2, 0x8c, 0x5, 0xad, 0x4e, 0x1e, 0xcc,
	0x57, 0x1e, 0xe2, 0x7d, 0x2f, 0x0, 0x83, 0xae, 0x8d, 0x9b, 0xab, 0xb5, 0x53, 0x2b, 0xbc, 0xa7,
	0x56, 0x2e, 0xaf, 0xc3, 0xb3, 0xe, 0x78, 0xcf, 0xaa, 0xdf, 0x5c, 0x7d, 0xfa, 0xf8, 0x8f, 0xf0,
	0xc4, 0xd6, 0xc6, 0x89, 0x1b, 0x7f, 0x68, 0x83, 0xa6, 0x7, 0xd6, 0xbe, 0xf7, 0x79, 0x40, 0x3e,
	0xc1, 0xff, 0x3, 0x0, 0x0, 0xff, 0xff, 0x8, 0xe6, 0xb2, 0x8a, 0x96, 0xa2, 0x0, 0x0,
}

func init() {
	if file_api_thrift != nil {
		return
	}
	type x struct{}
	builder := &thrift_reflection.FileDescriptorBuilder{
		Bytes:         file_idl_api_rawDesc,
		GoTypes:       file_api_thrift_go_types,
		GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
	}
	file_api_thrift = thrift_reflection.BuildFileDescriptor(builder)
}

func GetFileDescriptorForApi() *thrift_reflection.FileDescriptor {
	return file_api_thrift
}
func (p *EchoRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoRequest")
}
func (p *EchoResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoResponse")
}
func (p *EchoMultiBoolResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiBoolResponse")
}
func (p *EchoMultiByteResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiByteResponse")
}
func (p *EchoMultiInt16Response) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiInt16Response")
}
func (p *EchoMultiInt32Response) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiInt32Response")
}
func (p *EchoMultiInt64Response) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiInt64Response")
}
func (p *EchoMultiFloatResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiFloatResponse")
}
func (p *EchoMultiDoubleResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiDoubleResponse")
}
func (p *EchoMultiStringResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiStringResponse")
}
func (p *EchoOptionalStructRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalStructRequest")
}
func (p *EchoOptionalStructResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalStructResponse")
}
func (p *EchoOptionalMultiBoolRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiBoolRequest")
}
func (p *EchoOptionalMultiInt32Request) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiInt32Request")
}
func (p *EchoOptionalMultiStringRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiStringRequest")
}
func (p *EchoOptionalMultiBoolResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiBoolResponse")
}
func (p *EchoOptionalMultiInt32Response) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiInt32Response")
}
func (p *EchoOptionalMultiStringResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiStringResponse")
}
func (p TweetType) GetDescriptor() *thrift_reflection.EnumDescriptor {
	return file_api_thrift.GetEnumDescriptor("TweetType")
}
