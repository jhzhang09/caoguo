// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: commcn/weixin.proto

package weixin

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// OauthReq 请求结构
type OauthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jscode string `protobuf:"bytes,1,opt,name=jscode,proto3" json:"jscode,omitempty"` // 微信端jscode
}

func (x *OauthReq) Reset() {
	*x = OauthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commcn_weixin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OauthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OauthReq) ProtoMessage() {}

func (x *OauthReq) ProtoReflect() protoreflect.Message {
	mi := &file_commcn_weixin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OauthReq.ProtoReflect.Descriptor instead.
func (*OauthReq) Descriptor() ([]byte, []int) {
	return file_commcn_weixin_proto_rawDescGZIP(), []int{0}
}

func (x *OauthReq) GetJscode() string {
	if x != nil {
		return x.Jscode
	}
	return ""
}

// OauthResp 微信授权返回信息
type OauthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID   string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`      // 用户sessionid 用于当前交互使用
	OverdueTime int64  `protobuf:"varint,2,opt,name=overdueTime,proto3" json:"overdueTime,omitempty"` // 逾期时间点(时间戳)
}

func (x *OauthResp) Reset() {
	*x = OauthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commcn_weixin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OauthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OauthResp) ProtoMessage() {}

func (x *OauthResp) ProtoReflect() protoreflect.Message {
	mi := &file_commcn_weixin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OauthResp.ProtoReflect.Descriptor instead.
func (*OauthResp) Descriptor() ([]byte, []int) {
	return file_commcn_weixin_proto_rawDescGZIP(), []int{1}
}

func (x *OauthResp) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *OauthResp) GetOverdueTime() int64 {
	if x != nil {
		return x.OverdueTime
	}
	return 0
}

// WxUserinfo 用户信息
type WxUserinfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"` // 用户sessionid 用于当前交互使用
	NickName  string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`   //用户昵称
	AvatarURL string `protobuf:"bytes,3,opt,name=avatarURL,proto3" json:"avatarURL,omitempty"` //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Gender    string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`       //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City      string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`           //用户所在城市
	Province  string `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty"`   //用户所在省份
	Country   string `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`     //用户所在国家
	Language  string `protobuf:"bytes,8,opt,name=language,proto3" json:"language,omitempty"`   //用户的语言，简体中文为zh_CN
}

func (x *WxUserinfo) Reset() {
	*x = WxUserinfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commcn_weixin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxUserinfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxUserinfo) ProtoMessage() {}

func (x *WxUserinfo) ProtoReflect() protoreflect.Message {
	mi := &file_commcn_weixin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxUserinfo.ProtoReflect.Descriptor instead.
func (*WxUserinfo) Descriptor() ([]byte, []int) {
	return file_commcn_weixin_proto_rawDescGZIP(), []int{2}
}

func (x *WxUserinfo) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *WxUserinfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *WxUserinfo) GetAvatarURL() string {
	if x != nil {
		return x.AvatarURL
	}
	return ""
}

func (x *WxUserinfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *WxUserinfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *WxUserinfo) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *WxUserinfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *WxUserinfo) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

// WxUserinfoResp 用户信息返回
type WxUserinfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B bool `protobuf:"varint,1,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *WxUserinfoResp) Reset() {
	*x = WxUserinfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commcn_weixin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxUserinfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxUserinfoResp) ProtoMessage() {}

func (x *WxUserinfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_commcn_weixin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxUserinfoResp.ProtoReflect.Descriptor instead.
func (*WxUserinfoResp) Descriptor() ([]byte, []int) {
	return file_commcn_weixin_proto_rawDescGZIP(), []int{3}
}

func (x *WxUserinfoResp) GetB() bool {
	if x != nil {
		return x.B
	}
	return false
}

// GetQrcodeResp 获取二维码
type GetQrcodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`   // 小程序页面头部
	Scene string `protobuf:"bytes,2,opt,name=scene,proto3" json:"scene,omitempty"` // 附带参数
}

func (x *GetQrcodeReq) Reset() {
	*x = GetQrcodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commcn_weixin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQrcodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQrcodeReq) ProtoMessage() {}

func (x *GetQrcodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_commcn_weixin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQrcodeReq.ProtoReflect.Descriptor instead.
func (*GetQrcodeReq) Descriptor() ([]byte, []int) {
	return file_commcn_weixin_proto_rawDescGZIP(), []int{4}
}

func (x *GetQrcodeReq) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *GetQrcodeReq) GetScene() string {
	if x != nil {
		return x.Scene
	}
	return ""
}

// GetQrcodeResp 获取二维码
type GetQrcodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"` // 二维码图片地址
}

func (x *GetQrcodeResp) Reset() {
	*x = GetQrcodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commcn_weixin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQrcodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQrcodeResp) ProtoMessage() {}

func (x *GetQrcodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_commcn_weixin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQrcodeResp.ProtoReflect.Descriptor instead.
func (*GetQrcodeResp) Descriptor() ([]byte, []int) {
	return file_commcn_weixin_proto_rawDescGZIP(), []int{5}
}

func (x *GetQrcodeResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_commcn_weixin_proto protoreflect.FileDescriptor

var file_commcn_weixin_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x63, 0x6e, 0x2f, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x22, 0x22, 0x0a,
	0x08, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x4b, 0x0a, 0x09, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xe2,
	0x01, 0x0a, 0x0a, 0x57, 0x78, 0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x0e, 0x57, 0x78, 0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x01, 0x62, 0x22, 0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x22, 0x21, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x32, 0x77, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x2e, 0x0a, 0x05, 0x4f, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x10, 0x2e, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x4f, 0x61, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x4f, 0x61,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x77, 0x65,
	0x69, 0x78, 0x69, 0x6e, 0x2e, 0x57, 0x78, 0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x1a,
	0x16, 0x2e, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2e, 0x57, 0x78, 0x55, 0x73, 0x65, 0x72, 0x69,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x63, 0x6e, 0x2f, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commcn_weixin_proto_rawDescOnce sync.Once
	file_commcn_weixin_proto_rawDescData = file_commcn_weixin_proto_rawDesc
)

func file_commcn_weixin_proto_rawDescGZIP() []byte {
	file_commcn_weixin_proto_rawDescOnce.Do(func() {
		file_commcn_weixin_proto_rawDescData = protoimpl.X.CompressGZIP(file_commcn_weixin_proto_rawDescData)
	})
	return file_commcn_weixin_proto_rawDescData
}

var file_commcn_weixin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_commcn_weixin_proto_goTypes = []interface{}{
	(*OauthReq)(nil),       // 0: weixin.OauthReq
	(*OauthResp)(nil),      // 1: weixin.OauthResp
	(*WxUserinfo)(nil),     // 2: weixin.WxUserinfo
	(*WxUserinfoResp)(nil), // 3: weixin.WxUserinfoResp
	(*GetQrcodeReq)(nil),   // 4: weixin.GetQrcodeReq
	(*GetQrcodeResp)(nil),  // 5: weixin.GetQrcodeResp
}
var file_commcn_weixin_proto_depIdxs = []int32{
	0, // 0: weixin.Hello.Oauth:input_type -> weixin.OauthReq
	2, // 1: weixin.Hello.UpdateUserInfo:input_type -> weixin.WxUserinfo
	1, // 2: weixin.Hello.Oauth:output_type -> weixin.OauthResp
	3, // 3: weixin.Hello.UpdateUserInfo:output_type -> weixin.WxUserinfoResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commcn_weixin_proto_init() }
func file_commcn_weixin_proto_init() {
	if File_commcn_weixin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commcn_weixin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OauthReq); i {
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
		file_commcn_weixin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OauthResp); i {
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
		file_commcn_weixin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxUserinfo); i {
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
		file_commcn_weixin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxUserinfoResp); i {
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
		file_commcn_weixin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQrcodeReq); i {
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
		file_commcn_weixin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQrcodeResp); i {
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
			RawDescriptor: file_commcn_weixin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commcn_weixin_proto_goTypes,
		DependencyIndexes: file_commcn_weixin_proto_depIdxs,
		MessageInfos:      file_commcn_weixin_proto_msgTypes,
	}.Build()
	File_commcn_weixin_proto = out.File
	file_commcn_weixin_proto_rawDesc = nil
	file_commcn_weixin_proto_goTypes = nil
	file_commcn_weixin_proto_depIdxs = nil
}
