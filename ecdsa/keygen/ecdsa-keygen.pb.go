// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.0
// source: protob/ecdsa-keygen.proto

package keygen

import (
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

//
// Represents a BROADCAST message sent during Round 1 of the ECDSA TSS keygen protocol.
type KGRound1Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid          []byte `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	ViKeygen     []byte `protobuf:"bytes,2,opt,name=ViKeygen,proto3" json:"ViKeygen,omitempty"`
	Ssid         []byte `protobuf:"bytes,3,opt,name=ssid,proto3" json:"ssid,omitempty"`
	ViKeyRefresh []byte `protobuf:"bytes,4,opt,name=ViKeyRefresh,proto3" json:"ViKeyRefresh,omitempty"`
}

func (x *KGRound1Message) Reset() {
	*x = KGRound1Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_keygen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound1Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound1Message) ProtoMessage() {}

func (x *KGRound1Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_keygen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound1Message.ProtoReflect.Descriptor instead.
func (*KGRound1Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_keygen_proto_rawDescGZIP(), []int{0}
}

func (x *KGRound1Message) GetSid() []byte {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *KGRound1Message) GetViKeygen() []byte {
	if x != nil {
		return x.ViKeygen
	}
	return nil
}

func (x *KGRound1Message) GetSsid() []byte {
	if x != nil {
		return x.Ssid
	}
	return nil
}

func (x *KGRound1Message) GetViKeyRefresh() []byte {
	if x != nil {
		return x.ViKeyRefresh
	}
	return nil
}

//
// Represents a BROADCAST message sent during Round 1 of the ECDSA TSS keygen protocol.
type KGRound2Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid       []byte   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Ridi      []byte   `protobuf:"bytes,2,opt,name=ridi,proto3" json:"ridi,omitempty"`
	Ui        []byte   `protobuf:"bytes,3,opt,name=ui,proto3" json:"ui,omitempty"`
	PaillierN []byte   `protobuf:"bytes,4,opt,name=paillier_n,json=paillierN,proto3" json:"paillier_n,omitempty"`
	XiKeygen  [][]byte `protobuf:"bytes,5,rep,name=XiKeygen,proto3" json:"XiKeygen,omitempty"`
	AiKeygen  [][]byte `protobuf:"bytes,6,rep,name=AiKeygen,proto3" json:"AiKeygen,omitempty"`
	Vs        [][]byte `protobuf:"bytes,7,rep,name=vs,proto3" json:"vs,omitempty"`
	// Refresh:
	Ssid      []byte   `protobuf:"bytes,8,opt,name=ssid,proto3" json:"ssid,omitempty"`
	XiRefresh [][]byte `protobuf:"bytes,9,rep,name=XiRefresh,proto3" json:"XiRefresh,omitempty"`
	AiRefresh [][]byte `protobuf:"bytes,10,rep,name=AiRefresh,proto3" json:"AiRefresh,omitempty"`
	Yi        [][]byte `protobuf:"bytes,11,rep,name=Yi,proto3" json:"Yi,omitempty"`
	Bi        [][]byte `protobuf:"bytes,12,rep,name=Bi,proto3" json:"Bi,omitempty"`
	Ni        []byte   `protobuf:"bytes,13,opt,name=Ni,proto3" json:"Ni,omitempty"`
	Si        []byte   `protobuf:"bytes,14,opt,name=si,proto3" json:"si,omitempty"`
	Ti        []byte   `protobuf:"bytes,15,opt,name=ti,proto3" json:"ti,omitempty"`
	PsiiProof [][]byte `protobuf:"bytes,16,rep,name=psii_proof,json=psiiProof,proto3" json:"psii_proof,omitempty"`
	Rhoi      []byte   `protobuf:"bytes,17,opt,name=rhoi,proto3" json:"rhoi,omitempty"`
}

func (x *KGRound2Message) Reset() {
	*x = KGRound2Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_keygen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound2Message) ProtoMessage() {}

func (x *KGRound2Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_keygen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound2Message.ProtoReflect.Descriptor instead.
func (*KGRound2Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_keygen_proto_rawDescGZIP(), []int{1}
}

func (x *KGRound2Message) GetSid() []byte {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *KGRound2Message) GetRidi() []byte {
	if x != nil {
		return x.Ridi
	}
	return nil
}

func (x *KGRound2Message) GetUi() []byte {
	if x != nil {
		return x.Ui
	}
	return nil
}

func (x *KGRound2Message) GetPaillierN() []byte {
	if x != nil {
		return x.PaillierN
	}
	return nil
}

func (x *KGRound2Message) GetXiKeygen() [][]byte {
	if x != nil {
		return x.XiKeygen
	}
	return nil
}

func (x *KGRound2Message) GetAiKeygen() [][]byte {
	if x != nil {
		return x.AiKeygen
	}
	return nil
}

func (x *KGRound2Message) GetVs() [][]byte {
	if x != nil {
		return x.Vs
	}
	return nil
}

func (x *KGRound2Message) GetSsid() []byte {
	if x != nil {
		return x.Ssid
	}
	return nil
}

func (x *KGRound2Message) GetXiRefresh() [][]byte {
	if x != nil {
		return x.XiRefresh
	}
	return nil
}

func (x *KGRound2Message) GetAiRefresh() [][]byte {
	if x != nil {
		return x.AiRefresh
	}
	return nil
}

func (x *KGRound2Message) GetYi() [][]byte {
	if x != nil {
		return x.Yi
	}
	return nil
}

func (x *KGRound2Message) GetBi() [][]byte {
	if x != nil {
		return x.Bi
	}
	return nil
}

func (x *KGRound2Message) GetNi() []byte {
	if x != nil {
		return x.Ni
	}
	return nil
}

func (x *KGRound2Message) GetSi() []byte {
	if x != nil {
		return x.Si
	}
	return nil
}

func (x *KGRound2Message) GetTi() []byte {
	if x != nil {
		return x.Ti
	}
	return nil
}

func (x *KGRound2Message) GetPsiiProof() [][]byte {
	if x != nil {
		return x.PsiiProof
	}
	return nil
}

func (x *KGRound2Message) GetRhoi() []byte {
	if x != nil {
		return x.Rhoi
	}
	return nil
}

//
// Represents a P2P message sent to each party during Round 3 of the ECDSA TSS keygen and key refresh protocols.
type KGRound3Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid          []byte   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	PsiiSchProof [][]byte `protobuf:"bytes,2,rep,name=psii_sch_proof,json=psiiSchProof,proto3" json:"psii_sch_proof,omitempty"`
	// refresh:
	Ssid         []byte   `protobuf:"bytes,3,opt,name=ssid,proto3" json:"ssid,omitempty"`
	PsiiModProof [][]byte `protobuf:"bytes,4,rep,name=psii_mod_proof,json=psiiModProof,proto3" json:"psii_mod_proof,omitempty"`
	PhijiProof   [][]byte `protobuf:"bytes,5,rep,name=phiji_proof,json=phijiProof,proto3" json:"phiji_proof,omitempty"`
	PiiProof     [][]byte `protobuf:"bytes,6,rep,name=pii_proof,json=piiProof,proto3" json:"pii_proof,omitempty"`
	Cvssji       []byte   `protobuf:"bytes,7,opt,name=Cvssji,proto3" json:"Cvssji,omitempty"`
	RandCvssji   []byte   `protobuf:"bytes,8,opt,name=randCvssji,proto3" json:"randCvssji,omitempty"`
	Czeroji      []byte   `protobuf:"bytes,9,opt,name=Czeroji,proto3" json:"Czeroji,omitempty"`
	RandCzeroji  []byte   `protobuf:"bytes,10,opt,name=randCzeroji,proto3" json:"randCzeroji,omitempty"`
	PsijiProof   [][]byte `protobuf:"bytes,11,rep,name=psiji_proof,json=psijiProof,proto3" json:"psiji_proof,omitempty"`
}

func (x *KGRound3Message) Reset() {
	*x = KGRound3Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_keygen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound3Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound3Message) ProtoMessage() {}

func (x *KGRound3Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_keygen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound3Message.ProtoReflect.Descriptor instead.
func (*KGRound3Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_keygen_proto_rawDescGZIP(), []int{2}
}

func (x *KGRound3Message) GetSid() []byte {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *KGRound3Message) GetPsiiSchProof() [][]byte {
	if x != nil {
		return x.PsiiSchProof
	}
	return nil
}

func (x *KGRound3Message) GetSsid() []byte {
	if x != nil {
		return x.Ssid
	}
	return nil
}

func (x *KGRound3Message) GetPsiiModProof() [][]byte {
	if x != nil {
		return x.PsiiModProof
	}
	return nil
}

func (x *KGRound3Message) GetPhijiProof() [][]byte {
	if x != nil {
		return x.PhijiProof
	}
	return nil
}

func (x *KGRound3Message) GetPiiProof() [][]byte {
	if x != nil {
		return x.PiiProof
	}
	return nil
}

func (x *KGRound3Message) GetCvssji() []byte {
	if x != nil {
		return x.Cvssji
	}
	return nil
}

func (x *KGRound3Message) GetRandCvssji() []byte {
	if x != nil {
		return x.RandCvssji
	}
	return nil
}

func (x *KGRound3Message) GetCzeroji() []byte {
	if x != nil {
		return x.Czeroji
	}
	return nil
}

func (x *KGRound3Message) GetRandCzeroji() []byte {
	if x != nil {
		return x.RandCzeroji
	}
	return nil
}

func (x *KGRound3Message) GetPsijiProof() [][]byte {
	if x != nil {
		return x.PsijiProof
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Round 4 of the ECDSA TSS keygen protocol.
type KGRound4Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid       []byte `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Abort     bool   `protobuf:"varint,2,opt,name=abort,proto3" json:"abort,omitempty"`
	Mu        []byte `protobuf:"bytes,3,opt,name=mu,proto3" json:"mu,omitempty"`
	CulpritPj int32  `protobuf:"varint,4,opt,name=culpritPj,proto3" json:"culpritPj,omitempty"`
	Cji       []byte `protobuf:"bytes,5,opt,name=Cji,proto3" json:"Cji,omitempty"`
	Xji       []byte `protobuf:"bytes,6,opt,name=xji,proto3" json:"xji,omitempty"`
}

func (x *KGRound4Message) Reset() {
	*x = KGRound4Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_keygen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound4Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound4Message) ProtoMessage() {}

func (x *KGRound4Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_keygen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound4Message.ProtoReflect.Descriptor instead.
func (*KGRound4Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_keygen_proto_rawDescGZIP(), []int{3}
}

func (x *KGRound4Message) GetSid() []byte {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *KGRound4Message) GetAbort() bool {
	if x != nil {
		return x.Abort
	}
	return false
}

func (x *KGRound4Message) GetMu() []byte {
	if x != nil {
		return x.Mu
	}
	return nil
}

func (x *KGRound4Message) GetCulpritPj() int32 {
	if x != nil {
		return x.CulpritPj
	}
	return 0
}

func (x *KGRound4Message) GetCji() []byte {
	if x != nil {
		return x.Cji
	}
	return nil
}

func (x *KGRound4Message) GetXji() []byte {
	if x != nil {
		return x.Xji
	}
	return nil
}

var File_protob_ecdsa_keygen_proto protoreflect.FileDescriptor

var file_protob_ecdsa_keygen_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2f, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2d, 0x6b,
	0x65, 0x79, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x73, 0x73, 0x6c, 0x69, 0x62, 0x2e, 0x65, 0x63, 0x64, 0x73,
	0x61, 0x2e, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x22, 0x77, 0x0a, 0x0f, 0x4b, 0x47, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x31, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x56, 0x69, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x56, 0x69, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x56, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x56, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x22, 0x81, 0x03, 0x0a, 0x0f, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x69, 0x64, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x69, 0x64, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x75,
	0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x75, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x58, 0x69,
	0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x58, 0x69,
	0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x69, 0x4b, 0x65, 0x79, 0x67,
	0x65, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x41, 0x69, 0x4b, 0x65, 0x79, 0x67,
	0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x02,
	0x76, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x58, 0x69, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x58, 0x69, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x69, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x41, 0x69, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x59, 0x69, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x02,
	0x59, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x42, 0x69, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x02,
	0x42, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x4e, 0x69, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x4e, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x69, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x73, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x69, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x74, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x73, 0x69, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x73, 0x69, 0x69, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x68, 0x6f, 0x69, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x72, 0x68, 0x6f, 0x69, 0x22, 0xd6, 0x02, 0x0a, 0x0f, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70,
	0x73, 0x69, 0x69, 0x5f, 0x73, 0x63, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x73, 0x69, 0x69, 0x53, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x73, 0x69, 0x69, 0x5f, 0x6d, 0x6f,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x70,
	0x73, 0x69, 0x69, 0x4d, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x68, 0x69, 0x6a, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0a, 0x70, 0x68, 0x69, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x69, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x08, 0x70, 0x69, 0x69, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x76, 0x73,
	0x73, 0x6a, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x43, 0x76, 0x73, 0x73, 0x6a,
	0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x43, 0x76, 0x73, 0x73, 0x6a, 0x69, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x43, 0x76, 0x73, 0x73, 0x6a,
	0x69, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x7a, 0x65, 0x72, 0x6f, 0x6a, 0x69, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x43, 0x7a, 0x65, 0x72, 0x6f, 0x6a, 0x69, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x61, 0x6e, 0x64, 0x43, 0x7a, 0x65, 0x72, 0x6f, 0x6a, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x72, 0x61, 0x6e, 0x64, 0x43, 0x7a, 0x65, 0x72, 0x6f, 0x6a, 0x69, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x73, 0x69, 0x6a, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0a, 0x70, 0x73, 0x69, 0x6a, 0x69, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x8b,
	0x01, 0x0a, 0x0f, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x34, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x73, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x75,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x6d, 0x75, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x75,
	0x6c, 0x70, 0x72, 0x69, 0x74, 0x50, 0x6a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x50, 0x6a, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6a, 0x69, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x43, 0x6a, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x78, 0x6a,
	0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x78, 0x6a, 0x69, 0x42, 0x0e, 0x5a, 0x0c,
	0x65, 0x63, 0x64, 0x73, 0x61, 0x2f, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protob_ecdsa_keygen_proto_rawDescOnce sync.Once
	file_protob_ecdsa_keygen_proto_rawDescData = file_protob_ecdsa_keygen_proto_rawDesc
)

func file_protob_ecdsa_keygen_proto_rawDescGZIP() []byte {
	file_protob_ecdsa_keygen_proto_rawDescOnce.Do(func() {
		file_protob_ecdsa_keygen_proto_rawDescData = protoimpl.X.CompressGZIP(file_protob_ecdsa_keygen_proto_rawDescData)
	})
	return file_protob_ecdsa_keygen_proto_rawDescData
}

var file_protob_ecdsa_keygen_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protob_ecdsa_keygen_proto_goTypes = []interface{}{
	(*KGRound1Message)(nil), // 0: binance.tsslib.ecdsa.keygen.KGRound1Message
	(*KGRound2Message)(nil), // 1: binance.tsslib.ecdsa.keygen.KGRound2Message
	(*KGRound3Message)(nil), // 2: binance.tsslib.ecdsa.keygen.KGRound3Message
	(*KGRound4Message)(nil), // 3: binance.tsslib.ecdsa.keygen.KGRound4Message
}
var file_protob_ecdsa_keygen_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protob_ecdsa_keygen_proto_init() }
func file_protob_ecdsa_keygen_proto_init() {
	if File_protob_ecdsa_keygen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protob_ecdsa_keygen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGRound1Message); i {
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
		file_protob_ecdsa_keygen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGRound2Message); i {
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
		file_protob_ecdsa_keygen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGRound3Message); i {
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
		file_protob_ecdsa_keygen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGRound4Message); i {
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
			RawDescriptor: file_protob_ecdsa_keygen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protob_ecdsa_keygen_proto_goTypes,
		DependencyIndexes: file_protob_ecdsa_keygen_proto_depIdxs,
		MessageInfos:      file_protob_ecdsa_keygen_proto_msgTypes,
	}.Build()
	File_protob_ecdsa_keygen_proto = out.File
	file_protob_ecdsa_keygen_proto_rawDesc = nil
	file_protob_ecdsa_keygen_proto_goTypes = nil
	file_protob_ecdsa_keygen_proto_depIdxs = nil
}
