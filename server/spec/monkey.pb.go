// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: spec/monkey.proto

package spec

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type AlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AlbumRequest) Reset() {
	*x = AlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumRequest) ProtoMessage() {}

func (x *AlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumRequest.ProtoReflect.Descriptor instead.
func (*AlbumRequest) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{0}
}

func (x *AlbumRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type AlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Albums []*Album `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *AlbumReply) Reset() {
	*x = AlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumReply) ProtoMessage() {}

func (x *AlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumReply.ProtoReflect.Descriptor instead.
func (*AlbumReply) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{1}
}

func (x *AlbumReply) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl   string  `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	AlbumTitle string  `protobuf:"bytes,2,opt,name=album_title,json=albumTitle,proto3" json:"album_title,omitempty"`
	Songs      []*Song `protobuf:"bytes,3,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{2}
}

func (x *Album) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Album) GetAlbumTitle() string {
	if x != nil {
		return x.AlbumTitle
	}
	return ""
}

func (x *Album) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongTitle   string `protobuf:"bytes,1,opt,name=song_title,json=songTitle,proto3" json:"song_title,omitempty"`
	TrackNumber string `protobuf:"bytes,2,opt,name=track_number,json=trackNumber,proto3" json:"track_number,omitempty"`
	SongUrl     string `protobuf:"bytes,3,opt,name=song_url,json=songUrl,proto3" json:"song_url,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{3}
}

func (x *Song) GetSongTitle() string {
	if x != nil {
		return x.SongTitle
	}
	return ""
}

func (x *Song) GetTrackNumber() string {
	if x != nil {
		return x.TrackNumber
	}
	return ""
}

func (x *Song) GetSongUrl() string {
	if x != nil {
		return x.SongUrl
	}
	return ""
}

var File_spec_monkey_proto protoreflect.FileDescriptor

var file_spec_monkey_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x70, 0x65, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x0a, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x70, 0x65, 0x63,
	0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x22, 0x67,
	0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x22, 0x63, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x6e, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x32, 0x52, 0x0a, 0x06,
	0x4d, 0x6f, 0x6e, 0x6b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x73, 0x12, 0x12, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x3a, 0x01, 0x2a,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a,
	0x6f, 0x65, 0x52, 0x6f, 0x75, 0x72, 0x6b, 0x65, 0x31, 0x32, 0x33, 0x2f, 0x4d, 0x6f, 0x6e, 0x6b,
	0x65, 0x79, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_monkey_proto_rawDescOnce sync.Once
	file_spec_monkey_proto_rawDescData = file_spec_monkey_proto_rawDesc
)

func file_spec_monkey_proto_rawDescGZIP() []byte {
	file_spec_monkey_proto_rawDescOnce.Do(func() {
		file_spec_monkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_monkey_proto_rawDescData)
	})
	return file_spec_monkey_proto_rawDescData
}

var file_spec_monkey_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spec_monkey_proto_goTypes = []interface{}{
	(*AlbumRequest)(nil), // 0: spec.AlbumRequest
	(*AlbumReply)(nil),   // 1: spec.AlbumReply
	(*Album)(nil),        // 2: spec.Album
	(*Song)(nil),         // 3: spec.Song
}
var file_spec_monkey_proto_depIdxs = []int32{
	2, // 0: spec.AlbumReply.albums:type_name -> spec.Album
	3, // 1: spec.Album.songs:type_name -> spec.Song
	0, // 2: spec.Monkey.GetAlbums:input_type -> spec.AlbumRequest
	1, // 3: spec.Monkey.GetAlbums:output_type -> spec.AlbumReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spec_monkey_proto_init() }
func file_spec_monkey_proto_init() {
	if File_spec_monkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spec_monkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spec_monkey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spec_monkey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spec_monkey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Song); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spec_monkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spec_monkey_proto_goTypes,
		DependencyIndexes: file_spec_monkey_proto_depIdxs,
		MessageInfos:      file_spec_monkey_proto_msgTypes,
	}.Build()
	File_spec_monkey_proto = out.File
	file_spec_monkey_proto_rawDesc = nil
	file_spec_monkey_proto_goTypes = nil
	file_spec_monkey_proto_depIdxs = nil
}
