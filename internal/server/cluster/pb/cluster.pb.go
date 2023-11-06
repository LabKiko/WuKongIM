// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: internal/server/cluster/pb/cluster.proto

package pb

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

type Role int32

const (
	Role_Unknown      Role = 0
	Role_Follower     Role = 1
	Role_Candidate    Role = 2
	Role_PreCandidate Role = 3
	Role_Leader       Role = 4
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "Unknown",
		1: "Follower",
		2: "Candidate",
		3: "PreCandidate",
		4: "Leader",
	}
	Role_value = map[string]int32{
		"Unknown":      0,
		"Follower":     1,
		"Candidate":    2,
		"PreCandidate": 3,
		"Leader":       4,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_server_cluster_pb_cluster_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_internal_server_cluster_pb_cluster_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{0}
}

type PeerState int32

const (
	PeerState_PeerStateInitial  PeerState = 0 // 初始化
	PeerState_PeerStateNormal   PeerState = 1 // 正常
	PeerState_PeerStateElecting PeerState = 2 // 选举中
	PeerState_PeerStateDown     PeerState = 3 // 下线
)

// Enum value maps for PeerState.
var (
	PeerState_name = map[int32]string{
		0: "PeerStateInitial",
		1: "PeerStateNormal",
		2: "PeerStateElecting",
		3: "PeerStateDown",
	}
	PeerState_value = map[string]int32{
		"PeerStateInitial":  0,
		"PeerStateNormal":   1,
		"PeerStateElecting": 2,
		"PeerStateDown":     3,
	}
)

func (x PeerState) Enum() *PeerState {
	p := new(PeerState)
	*p = x
	return p
}

func (x PeerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerState) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_server_cluster_pb_cluster_proto_enumTypes[1].Descriptor()
}

func (PeerState) Type() protoreflect.EnumType {
	return &file_internal_server_cluster_pb_cluster_proto_enumTypes[1]
}

func (x PeerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerState.Descriptor instead.
func (PeerState) EnumDescriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{1}
}

type SlotState int32

const (
	SlotState_SlotStateInitial  SlotState = 0 // 初始化
	SlotState_SlotStateNormal   SlotState = 1 // 正常
	SlotState_SlotStateElecting SlotState = 2 // 选举中
)

// Enum value maps for SlotState.
var (
	SlotState_name = map[int32]string{
		0: "SlotStateInitial",
		1: "SlotStateNormal",
		2: "SlotStateElecting",
	}
	SlotState_value = map[string]int32{
		"SlotStateInitial":  0,
		"SlotStateNormal":   1,
		"SlotStateElecting": 2,
	}
)

func (x SlotState) Enum() *SlotState {
	p := new(SlotState)
	*p = x
	return p
}

func (x SlotState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SlotState) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_server_cluster_pb_cluster_proto_enumTypes[2].Descriptor()
}

func (SlotState) Type() protoreflect.EnumType {
	return &file_internal_server_cluster_pb_cluster_proto_enumTypes[2]
}

func (x SlotState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SlotState.Descriptor instead.
func (SlotState) EnumDescriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{2}
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotCount    uint32  `protobuf:"varint,1,opt,name=slotCount,proto3" json:"slotCount,omitempty"`       // 槽位数量
	Leader       uint64  `protobuf:"varint,2,opt,name=leader,proto3" json:"leader,omitempty"`             // 领导者节点ID
	ReplicaCount uint32  `protobuf:"varint,3,opt,name=replicaCount,proto3" json:"replicaCount,omitempty"` // 副本数量
	Peers        []*Peer `protobuf:"bytes,4,rep,name=peers,proto3" json:"peers,omitempty"`                // 节点
	Slots        []*Slot `protobuf:"bytes,5,rep,name=slots,proto3" json:"slots,omitempty"`                // 节点槽位
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetSlotCount() uint32 {
	if x != nil {
		return x.SlotCount
	}
	return 0
}

func (x *Cluster) GetLeader() uint64 {
	if x != nil {
		return x.Leader
	}
	return 0
}

func (x *Cluster) GetReplicaCount() uint32 {
	if x != nil {
		return x.ReplicaCount
	}
	return 0
}

func (x *Cluster) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *Cluster) GetSlots() []*Slot {
	if x != nil {
		return x.Slots
	}
	return nil
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerID         uint64    `protobuf:"varint,1,opt,name=peerID,proto3" json:"peerID,omitempty"`                // 节点ID
	ServerAddr     string    `protobuf:"bytes,2,opt,name=serverAddr,proto3" json:"serverAddr,omitempty"`         // 节点服务地址
	Role           Role      `protobuf:"varint,3,opt,name=role,proto3,enum=Role" json:"role,omitempty"`          // 节点角色
	State          PeerState `protobuf:"varint,4,opt,name=state,proto3,enum=PeerState" json:"state,omitempty"`   // 节点状态
	GrpcServerAddr string    `protobuf:"bytes,5,opt,name=grpcServerAddr,proto3" json:"grpcServerAddr,omitempty"` // 节点grpc服务地址
	ApiServerAddr  string    `protobuf:"bytes,6,opt,name=apiServerAddr,proto3" json:"apiServerAddr,omitempty"`   // 节点api服务地址
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *Peer) GetPeerID() uint64 {
	if x != nil {
		return x.PeerID
	}
	return 0
}

func (x *Peer) GetServerAddr() string {
	if x != nil {
		return x.ServerAddr
	}
	return ""
}

func (x *Peer) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_Unknown
}

func (x *Peer) GetState() PeerState {
	if x != nil {
		return x.State
	}
	return PeerState_PeerStateInitial
}

func (x *Peer) GetGrpcServerAddr() string {
	if x != nil {
		return x.GrpcServerAddr
	}
	return ""
}

func (x *Peer) GetApiServerAddr() string {
	if x != nil {
		return x.ApiServerAddr
	}
	return ""
}

type PeerSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"` // 节点
}

func (x *PeerSet) Reset() {
	*x = PeerSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerSet) ProtoMessage() {}

func (x *PeerSet) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerSet.ProtoReflect.Descriptor instead.
func (*PeerSet) Descriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *PeerSet) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type Slot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot   uint32    `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`                  // 槽位
	Leader uint64    `protobuf:"varint,2,opt,name=leader,proto3" json:"leader,omitempty"`              // 领导者节点ID
	State  SlotState `protobuf:"varint,3,opt,name=state,proto3,enum=SlotState" json:"state,omitempty"` // 槽位状态
	Peers  []uint64  `protobuf:"varint,5,rep,packed,name=peers,proto3" json:"peers,omitempty"`         // 节点ID
}

func (x *Slot) Reset() {
	*x = Slot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slot) ProtoMessage() {}

func (x *Slot) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slot.ProtoReflect.Descriptor instead.
func (*Slot) Descriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *Slot) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *Slot) GetLeader() uint64 {
	if x != nil {
		return x.Leader
	}
	return 0
}

func (x *Slot) GetState() SlotState {
	if x != nil {
		return x.State
	}
	return SlotState_SlotStateInitial
}

func (x *Slot) GetPeers() []uint64 {
	if x != nil {
		return x.Peers
	}
	return nil
}

// 槽位领导者
type SlotLeaderRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot       uint32 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`             // 槽位
	Leader     uint64 `protobuf:"varint,2,opt,name=leader,proto3" json:"leader,omitempty"`         // 领导者节点ID
	NeedUpdate bool   `protobuf:"varint,3,opt,name=needUpdate,proto3" json:"needUpdate,omitempty"` // 是否需要更新
}

func (x *SlotLeaderRelation) Reset() {
	*x = SlotLeaderRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlotLeaderRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotLeaderRelation) ProtoMessage() {}

func (x *SlotLeaderRelation) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotLeaderRelation.ProtoReflect.Descriptor instead.
func (*SlotLeaderRelation) Descriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *SlotLeaderRelation) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *SlotLeaderRelation) GetLeader() uint64 {
	if x != nil {
		return x.Leader
	}
	return 0
}

func (x *SlotLeaderRelation) GetNeedUpdate() bool {
	if x != nil {
		return x.NeedUpdate
	}
	return false
}

type SlotLeaderRelationSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotLeaderRelations []*SlotLeaderRelation `protobuf:"bytes,1,rep,name=slotLeaderRelations,proto3" json:"slotLeaderRelations,omitempty"` // 槽位领导者列表
}

func (x *SlotLeaderRelationSet) Reset() {
	*x = SlotLeaderRelationSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlotLeaderRelationSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotLeaderRelationSet) ProtoMessage() {}

func (x *SlotLeaderRelationSet) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotLeaderRelationSet.ProtoReflect.Descriptor instead.
func (*SlotLeaderRelationSet) Descriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *SlotLeaderRelationSet) GetSlotLeaderRelations() []*SlotLeaderRelation {
	if x != nil {
		return x.SlotLeaderRelations
	}
	return nil
}

type AllocateSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot     uint32   `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`          // 槽位
	Peers    []uint64 `protobuf:"varint,2,rep,packed,name=peers,proto3" json:"peers,omitempty"` // slot分配的节点
	LeaderID uint64   `protobuf:"varint,3,opt,name=leaderID,proto3" json:"leaderID,omitempty"`  //  slot的leader（如果为0则自动选举，如果不为0则指定此节点为领导节点）
}

func (x *AllocateSlot) Reset() {
	*x = AllocateSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateSlot) ProtoMessage() {}

func (x *AllocateSlot) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateSlot.ProtoReflect.Descriptor instead.
func (*AllocateSlot) Descriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *AllocateSlot) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *AllocateSlot) GetPeers() []uint64 {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *AllocateSlot) GetLeaderID() uint64 {
	if x != nil {
		return x.LeaderID
	}
	return 0
}

type AllocateSlotSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllocateSlots []*AllocateSlot `protobuf:"bytes,1,rep,name=allocateSlots,proto3" json:"allocateSlots,omitempty"` // slot分配
}

func (x *AllocateSlotSet) Reset() {
	*x = AllocateSlotSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateSlotSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateSlotSet) ProtoMessage() {}

func (x *AllocateSlotSet) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_cluster_pb_cluster_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateSlotSet.ProtoReflect.Descriptor instead.
func (*AllocateSlotSet) Descriptor() ([]byte, []int) {
	return file_internal_server_cluster_pb_cluster_proto_rawDescGZIP(), []int{7}
}

func (x *AllocateSlotSet) GetAllocateSlots() []*AllocateSlot {
	if x != nil {
		return x.AllocateSlots
	}
	return nil
}

var File_internal_server_cluster_pb_cluster_proto protoreflect.FileDescriptor

var file_internal_server_cluster_pb_cluster_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x07, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a,
	0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53,
	0x6c, 0x6f, 0x74, 0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x04, 0x50,
	0x65, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x67, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x22, 0x26, 0x0a, 0x07, 0x50, 0x65, 0x65, 0x72, 0x53, 0x65,
	0x74, 0x12, 0x1b, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x6a,
	0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x60, 0x0a, 0x12, 0x53, 0x6c,
	0x6f, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x6e, 0x65, 0x65, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x6e, 0x65, 0x65, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x5e, 0x0a, 0x15,
	0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x13, 0x73, 0x6c, 0x6f, 0x74, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x73, 0x6c, 0x6f, 0x74, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x54, 0x0a, 0x0c,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x46, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x53, 0x6c,
	0x6f, 0x74, 0x53, 0x65, 0x74, 0x12, 0x33, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x0d, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x2a, 0x4e, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x72, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x10, 0x04, 0x2a, 0x60, 0x0a, 0x09, 0x50, 0x65,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x65, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x65, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x10, 0x03, 0x2a, 0x4d, 0x0a, 0x09,
	0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x6c, 0x6f,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_server_cluster_pb_cluster_proto_rawDescOnce sync.Once
	file_internal_server_cluster_pb_cluster_proto_rawDescData = file_internal_server_cluster_pb_cluster_proto_rawDesc
)

func file_internal_server_cluster_pb_cluster_proto_rawDescGZIP() []byte {
	file_internal_server_cluster_pb_cluster_proto_rawDescOnce.Do(func() {
		file_internal_server_cluster_pb_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_server_cluster_pb_cluster_proto_rawDescData)
	})
	return file_internal_server_cluster_pb_cluster_proto_rawDescData
}

var file_internal_server_cluster_pb_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_internal_server_cluster_pb_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_server_cluster_pb_cluster_proto_goTypes = []interface{}{
	(Role)(0),                     // 0: Role
	(PeerState)(0),                // 1: PeerState
	(SlotState)(0),                // 2: SlotState
	(*Cluster)(nil),               // 3: Cluster
	(*Peer)(nil),                  // 4: Peer
	(*PeerSet)(nil),               // 5: PeerSet
	(*Slot)(nil),                  // 6: Slot
	(*SlotLeaderRelation)(nil),    // 7: SlotLeaderRelation
	(*SlotLeaderRelationSet)(nil), // 8: SlotLeaderRelationSet
	(*AllocateSlot)(nil),          // 9: AllocateSlot
	(*AllocateSlotSet)(nil),       // 10: AllocateSlotSet
}
var file_internal_server_cluster_pb_cluster_proto_depIdxs = []int32{
	4, // 0: Cluster.peers:type_name -> Peer
	6, // 1: Cluster.slots:type_name -> Slot
	0, // 2: Peer.role:type_name -> Role
	1, // 3: Peer.state:type_name -> PeerState
	4, // 4: PeerSet.peers:type_name -> Peer
	2, // 5: Slot.state:type_name -> SlotState
	7, // 6: SlotLeaderRelationSet.slotLeaderRelations:type_name -> SlotLeaderRelation
	9, // 7: AllocateSlotSet.allocateSlots:type_name -> AllocateSlot
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_internal_server_cluster_pb_cluster_proto_init() }
func file_internal_server_cluster_pb_cluster_proto_init() {
	if File_internal_server_cluster_pb_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_server_cluster_pb_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_internal_server_cluster_pb_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
		file_internal_server_cluster_pb_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerSet); i {
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
		file_internal_server_cluster_pb_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slot); i {
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
		file_internal_server_cluster_pb_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlotLeaderRelation); i {
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
		file_internal_server_cluster_pb_cluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlotLeaderRelationSet); i {
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
		file_internal_server_cluster_pb_cluster_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateSlot); i {
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
		file_internal_server_cluster_pb_cluster_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateSlotSet); i {
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
			RawDescriptor: file_internal_server_cluster_pb_cluster_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_server_cluster_pb_cluster_proto_goTypes,
		DependencyIndexes: file_internal_server_cluster_pb_cluster_proto_depIdxs,
		EnumInfos:         file_internal_server_cluster_pb_cluster_proto_enumTypes,
		MessageInfos:      file_internal_server_cluster_pb_cluster_proto_msgTypes,
	}.Build()
	File_internal_server_cluster_pb_cluster_proto = out.File
	file_internal_server_cluster_pb_cluster_proto_rawDesc = nil
	file_internal_server_cluster_pb_cluster_proto_goTypes = nil
	file_internal_server_cluster_pb_cluster_proto_depIdxs = nil
}
