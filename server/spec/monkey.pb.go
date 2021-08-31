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

type ArtistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ArtistsRequest) Reset() {
	*x = ArtistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtistsRequest) ProtoMessage() {}

func (x *ArtistsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ArtistsRequest.ProtoReflect.Descriptor instead.
func (*ArtistsRequest) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{0}
}

type ArtistsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artists     []*Artist `protobuf:"bytes,1,rep,name=artists,proto3" json:"artists,omitempty"`
	ArtistCount int32     `protobuf:"varint,2,opt,name=artist_count,json=artistCount,proto3" json:"artist_count,omitempty"`
}

func (x *ArtistsReply) Reset() {
	*x = ArtistsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtistsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtistsReply) ProtoMessage() {}

func (x *ArtistsReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ArtistsReply.ProtoReflect.Descriptor instead.
func (*ArtistsReply) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{1}
}

func (x *ArtistsReply) GetArtists() []*Artist {
	if x != nil {
		return x.Artists
	}
	return nil
}

func (x *ArtistsReply) GetArtistCount() int32 {
	if x != nil {
		return x.ArtistCount
	}
	return 0
}

type Artist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Id       int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Artist) Reset() {
	*x = Artist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artist) ProtoMessage() {}

func (x *Artist) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Artist.ProtoReflect.Descriptor instead.
func (*Artist) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{2}
}

func (x *Artist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artist) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Artist) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AlbumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtistId int64 `protobuf:"varint,1,opt,name=artist_id,json=artistId,proto3" json:"artist_id,omitempty"`
}

func (x *AlbumsRequest) Reset() {
	*x = AlbumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumsRequest) ProtoMessage() {}

func (x *AlbumsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AlbumsRequest.ProtoReflect.Descriptor instead.
func (*AlbumsRequest) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{3}
}

func (x *AlbumsRequest) GetArtistId() int64 {
	if x != nil {
		return x.ArtistId
	}
	return 0
}

type AlbumsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Albums     []*Album `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
	ArtistName string   `protobuf:"bytes,2,opt,name=artist_name,json=artistName,proto3" json:"artist_name,omitempty"`
}

func (x *AlbumsReply) Reset() {
	*x = AlbumsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumsReply) ProtoMessage() {}

func (x *AlbumsReply) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumsReply.ProtoReflect.Descriptor instead.
func (*AlbumsReply) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{4}
}

func (x *AlbumsReply) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

func (x *AlbumsReply) GetArtistName() string {
	if x != nil {
		return x.ArtistName
	}
	return ""
}

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AlbumArt string `protobuf:"bytes,2,opt,name=album_art,json=albumArt,proto3" json:"album_art,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[5]
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
	return file_spec_monkey_proto_rawDescGZIP(), []int{5}
}

func (x *Album) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Album) GetAlbumArt() string {
	if x != nil {
		return x.AlbumArt
	}
	return ""
}

func (x *Album) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SongsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtistId int64 `protobuf:"varint,1,opt,name=artist_id,json=artistId,proto3" json:"artist_id,omitempty"`
	AlbumId  int64 `protobuf:"varint,2,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *SongsRequest) Reset() {
	*x = SongsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongsRequest) ProtoMessage() {}

func (x *SongsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongsRequest.ProtoReflect.Descriptor instead.
func (*SongsRequest) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{6}
}

func (x *SongsRequest) GetArtistId() int64 {
	if x != nil {
		return x.ArtistId
	}
	return 0
}

func (x *SongsRequest) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

type SongsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*Song `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *SongsReply) Reset() {
	*x = SongsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongsReply) ProtoMessage() {}

func (x *SongsReply) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongsReply.ProtoReflect.Descriptor instead.
func (*SongsReply) Descriptor() ([]byte, []int) {
	return file_spec_monkey_proto_rawDescGZIP(), []int{7}
}

func (x *SongsReply) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_monkey_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_spec_monkey_proto_msgTypes[8]
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
	return file_spec_monkey_proto_rawDescGZIP(), []int{8}
}

func (x *Song) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Song) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_spec_monkey_proto protoreflect.FileDescriptor

var file_spec_monkey_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x70, 0x65, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x0c, 0x41, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x2c, 0x0a, 0x0d, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x53, 0x0a,
	0x0b, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x06,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73,
	0x70, 0x65, 0x63, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x48, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x41, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x0c,
	0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0a, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x05, 0x73,
	0x6f, 0x6e, 0x67, 0x73, 0x22, 0x32, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x9d, 0x02, 0x0a, 0x06, 0x4d, 0x6f, 0x6e,
	0x6b, 0x65, 0x79, 0x12, 0x4c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x73, 0x12, 0x5b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12,
	0x13, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12,
	0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x68,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x2e, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x7b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x65, 0x52, 0x6f, 0x75, 0x72, 0x6b, 0x65,
	0x31, 0x32, 0x33, 0x2f, 0x4d, 0x6f, 0x6e, 0x6b, 0x65, 0x79, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_spec_monkey_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_spec_monkey_proto_goTypes = []interface{}{
	(*ArtistsRequest)(nil), // 0: spec.ArtistsRequest
	(*ArtistsReply)(nil),   // 1: spec.ArtistsReply
	(*Artist)(nil),         // 2: spec.Artist
	(*AlbumsRequest)(nil),  // 3: spec.AlbumsRequest
	(*AlbumsReply)(nil),    // 4: spec.AlbumsReply
	(*Album)(nil),          // 5: spec.Album
	(*SongsRequest)(nil),   // 6: spec.SongsRequest
	(*SongsReply)(nil),     // 7: spec.SongsReply
	(*Song)(nil),           // 8: spec.Song
}
var file_spec_monkey_proto_depIdxs = []int32{
	2, // 0: spec.ArtistsReply.artists:type_name -> spec.Artist
	5, // 1: spec.AlbumsReply.albums:type_name -> spec.Album
	8, // 2: spec.SongsReply.songs:type_name -> spec.Song
	0, // 3: spec.Monkey.ListArtists:input_type -> spec.ArtistsRequest
	3, // 4: spec.Monkey.ListAlbums:input_type -> spec.AlbumsRequest
	6, // 5: spec.Monkey.ListSongs:input_type -> spec.SongsRequest
	1, // 6: spec.Monkey.ListArtists:output_type -> spec.ArtistsReply
	4, // 7: spec.Monkey.ListAlbums:output_type -> spec.AlbumsReply
	7, // 8: spec.Monkey.ListSongs:output_type -> spec.SongsReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spec_monkey_proto_init() }
func file_spec_monkey_proto_init() {
	if File_spec_monkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spec_monkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtistsRequest); i {
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
			switch v := v.(*ArtistsReply); i {
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
			switch v := v.(*Artist); i {
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
			switch v := v.(*AlbumsRequest); i {
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
		file_spec_monkey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumsReply); i {
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
		file_spec_monkey_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_spec_monkey_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongsRequest); i {
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
		file_spec_monkey_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongsReply); i {
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
		file_spec_monkey_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   9,
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
