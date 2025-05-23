// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: search.proto

package resourcepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Get statistics across multiple resources
// For these queries, we do not need authorization to see the actual values
type ResourceStatsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Namespace (tenant)
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// An optional list of group/resource identifiers
	// when empty, we assume searching across everything
	// NOTE, this query may need to federate across a few storage instances
	Kinds []string `protobuf:"bytes,2,rep,name=kinds,proto3" json:"kinds,omitempty"`
	// Limit the stats within a folder (not recursive!)
	Folder        string `protobuf:"bytes,3,opt,name=folder,proto3" json:"folder,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceStatsRequest) Reset() {
	*x = ResourceStatsRequest{}
	mi := &file_search_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceStatsRequest) ProtoMessage() {}

func (x *ResourceStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceStatsRequest.ProtoReflect.Descriptor instead.
func (*ResourceStatsRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceStatsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ResourceStatsRequest) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

func (x *ResourceStatsRequest) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

type ResourceStatsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Error details
	Error *ErrorResult `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// All results exist within this key
	Stats         []*ResourceStatsResponse_Stats `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceStatsResponse) Reset() {
	*x = ResourceStatsResponse{}
	mi := &file_search_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceStatsResponse) ProtoMessage() {}

func (x *ResourceStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceStatsResponse.ProtoReflect.Descriptor instead.
func (*ResourceStatsResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceStatsResponse) GetError() *ErrorResult {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ResourceStatsResponse) GetStats() []*ResourceStatsResponse_Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

// Search within a single resource
type ResourceSearchRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The key must include namespace + group + resource
	Options *ListOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	// To search additional resource types, add additional keys to this list
	// NOTE: queries will only support federation across kinds with common fields
	Federated []*ResourceKey `protobuf:"bytes,2,rep,name=federated,proto3" json:"federated,omitempty"`
	// When a query exists, it is parsed and used to influence
	// query string for chosen implementation (currently just bleve)
	// The score is only relevant when a query exists
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// max results
	Limit int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// where to start the query (eg, From)
	Offset int64 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	// sorting
	SortBy []*ResourceSearchRequest_Sort `protobuf:"bytes,6,rep,name=sortBy,proto3" json:"sortBy,omitempty"`
	// calculate field statistics
	Facet map[string]*ResourceSearchRequest_Facet `protobuf:"bytes,7,rep,name=facet,proto3" json:"facet,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// the return fields (empty will return everything)
	Fields []string `protobuf:"bytes,8,rep,name=fields,proto3" json:"fields,omitempty"`
	// explain each result (added to the each row)
	Explain       bool  `protobuf:"varint,9,opt,name=explain,proto3" json:"explain,omitempty"`
	IsDeleted     bool  `protobuf:"varint,10,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	Page          int64 `protobuf:"varint,11,opt,name=page,proto3" json:"page,omitempty"`
	Permission    int64 `protobuf:"varint,12,opt,name=permission,proto3" json:"permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceSearchRequest) Reset() {
	*x = ResourceSearchRequest{}
	mi := &file_search_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSearchRequest) ProtoMessage() {}

func (x *ResourceSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSearchRequest.ProtoReflect.Descriptor instead.
func (*ResourceSearchRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceSearchRequest) GetOptions() *ListOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ResourceSearchRequest) GetFederated() []*ResourceKey {
	if x != nil {
		return x.Federated
	}
	return nil
}

func (x *ResourceSearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ResourceSearchRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ResourceSearchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ResourceSearchRequest) GetSortBy() []*ResourceSearchRequest_Sort {
	if x != nil {
		return x.SortBy
	}
	return nil
}

func (x *ResourceSearchRequest) GetFacet() map[string]*ResourceSearchRequest_Facet {
	if x != nil {
		return x.Facet
	}
	return nil
}

func (x *ResourceSearchRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ResourceSearchRequest) GetExplain() bool {
	if x != nil {
		return x.Explain
	}
	return false
}

func (x *ResourceSearchRequest) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *ResourceSearchRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ResourceSearchRequest) GetPermission() int64 {
	if x != nil {
		return x.Permission
	}
	return 0
}

type ResourceSearchResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Error details
	Error *ErrorResult `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// All results exist within this key
	Key *ResourceKey `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Query results
	Results *ResourceTable `protobuf:"bytes,3,opt,name=results,proto3" json:"results,omitempty"`
	// The total hit count
	TotalHits int64 `protobuf:"varint,4,opt,name=total_hits,json=totalHits,proto3" json:"total_hits,omitempty"`
	// indicates how expensive was the query with respect to bytes read
	QueryCost float64 `protobuf:"fixed64,5,opt,name=query_cost,json=queryCost,proto3" json:"query_cost,omitempty"`
	// maximum score across all fields
	MaxScore float64 `protobuf:"fixed64,6,opt,name=max_score,json=maxScore,proto3" json:"max_score,omitempty"`
	// Facet results
	Facet         map[string]*ResourceSearchResponse_Facet `protobuf:"bytes,7,rep,name=facet,proto3" json:"facet,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceSearchResponse) Reset() {
	*x = ResourceSearchResponse{}
	mi := &file_search_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSearchResponse) ProtoMessage() {}

func (x *ResourceSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSearchResponse.ProtoReflect.Descriptor instead.
func (*ResourceSearchResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceSearchResponse) GetError() *ErrorResult {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ResourceSearchResponse) GetKey() *ResourceKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ResourceSearchResponse) GetResults() *ResourceTable {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ResourceSearchResponse) GetTotalHits() int64 {
	if x != nil {
		return x.TotalHits
	}
	return 0
}

func (x *ResourceSearchResponse) GetQueryCost() float64 {
	if x != nil {
		return x.QueryCost
	}
	return 0
}

func (x *ResourceSearchResponse) GetMaxScore() float64 {
	if x != nil {
		return x.MaxScore
	}
	return 0
}

func (x *ResourceSearchResponse) GetFacet() map[string]*ResourceSearchResponse_Facet {
	if x != nil {
		return x.Facet
	}
	return nil
}

type ResourceStatsResponse_Stats struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Resource group
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// Resource name
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// Number of items
	Count         int64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceStatsResponse_Stats) Reset() {
	*x = ResourceStatsResponse_Stats{}
	mi := &file_search_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceStatsResponse_Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceStatsResponse_Stats) ProtoMessage() {}

func (x *ResourceStatsResponse_Stats) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceStatsResponse_Stats.ProtoReflect.Descriptor instead.
func (*ResourceStatsResponse_Stats) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ResourceStatsResponse_Stats) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ResourceStatsResponse_Stats) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *ResourceStatsResponse_Stats) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ResourceSearchRequest_Sort struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Desc          bool                   `protobuf:"varint,2,opt,name=desc,proto3" json:"desc,omitempty"` // defaults to ascending
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceSearchRequest_Sort) Reset() {
	*x = ResourceSearchRequest_Sort{}
	mi := &file_search_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceSearchRequest_Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSearchRequest_Sort) ProtoMessage() {}

func (x *ResourceSearchRequest_Sort) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSearchRequest_Sort.ProtoReflect.Descriptor instead.
func (*ResourceSearchRequest_Sort) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ResourceSearchRequest_Sort) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ResourceSearchRequest_Sort) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

type ResourceSearchRequest_Facet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Limit         int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceSearchRequest_Facet) Reset() {
	*x = ResourceSearchRequest_Facet{}
	mi := &file_search_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceSearchRequest_Facet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSearchRequest_Facet) ProtoMessage() {}

func (x *ResourceSearchRequest_Facet) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSearchRequest_Facet.ProtoReflect.Descriptor instead.
func (*ResourceSearchRequest_Facet) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ResourceSearchRequest_Facet) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ResourceSearchRequest_Facet) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ResourceSearchResponse_Facet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Field string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// The distinct terms
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	// The number of documents that do *not* have this field
	Missing int64 `protobuf:"varint,3,opt,name=missing,proto3" json:"missing,omitempty"`
	// Top term stats
	Terms         []*ResourceSearchResponse_TermFacet `protobuf:"bytes,4,rep,name=terms,proto3" json:"terms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceSearchResponse_Facet) Reset() {
	*x = ResourceSearchResponse_Facet{}
	mi := &file_search_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceSearchResponse_Facet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSearchResponse_Facet) ProtoMessage() {}

func (x *ResourceSearchResponse_Facet) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSearchResponse_Facet.ProtoReflect.Descriptor instead.
func (*ResourceSearchResponse_Facet) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ResourceSearchResponse_Facet) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ResourceSearchResponse_Facet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ResourceSearchResponse_Facet) GetMissing() int64 {
	if x != nil {
		return x.Missing
	}
	return 0
}

func (x *ResourceSearchResponse_Facet) GetTerms() []*ResourceSearchResponse_TermFacet {
	if x != nil {
		return x.Terms
	}
	return nil
}

type ResourceSearchResponse_TermFacet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Term          string                 `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	Count         int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceSearchResponse_TermFacet) Reset() {
	*x = ResourceSearchResponse_TermFacet{}
	mi := &file_search_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceSearchResponse_TermFacet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSearchResponse_TermFacet) ProtoMessage() {}

func (x *ResourceSearchResponse_TermFacet) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSearchResponse_TermFacet.ProtoReflect.Descriptor instead.
func (*ResourceSearchResponse_TermFacet) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3, 1}
}

func (x *ResourceSearchResponse_TermFacet) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *ResourceSearchResponse_TermFacet) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6b,
	0x69, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0xd2, 0x01, 0x0a,
	0x15, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x1a, 0x4f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x8e, 0x05, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x09,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x73, 0x6f, 0x72,
	0x74, 0x42, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x66, 0x61, 0x63, 0x65, 0x74, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x66, 0x61, 0x63, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x30, 0x0a, 0x04, 0x53, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x1a, 0x33, 0x0a, 0x05,
	0x46, 0x61, 0x63, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x1a, 0x5f, 0x0a, 0x0a, 0x46, 0x61, 0x63, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xea, 0x04, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x68, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x48, 0x69, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x41, 0x0a, 0x05, 0x66, 0x61, 0x63, 0x65, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66,
	0x61, 0x63, 0x65, 0x74, 0x1a, 0x8f, 0x01, 0x0a, 0x05, 0x46, 0x61, 0x63, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x05, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x46, 0x61, 0x63, 0x65, 0x74, 0x52,
	0x05, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x1a, 0x35, 0x0a, 0x09, 0x54, 0x65, 0x72, 0x6d, 0x46, 0x61,
	0x63, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x60, 0x0a,
	0x0a, 0x46, 0x61, 0x63, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46,
	0x61, 0x63, 0x65, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0xa9, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e,
	0x61, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData []byte
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_search_proto_rawDesc), len(file_search_proto_rawDesc)))
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_search_proto_goTypes = []any{
	(*ResourceStatsRequest)(nil),             // 0: resource.ResourceStatsRequest
	(*ResourceStatsResponse)(nil),            // 1: resource.ResourceStatsResponse
	(*ResourceSearchRequest)(nil),            // 2: resource.ResourceSearchRequest
	(*ResourceSearchResponse)(nil),           // 3: resource.ResourceSearchResponse
	(*ResourceStatsResponse_Stats)(nil),      // 4: resource.ResourceStatsResponse.Stats
	(*ResourceSearchRequest_Sort)(nil),       // 5: resource.ResourceSearchRequest.Sort
	(*ResourceSearchRequest_Facet)(nil),      // 6: resource.ResourceSearchRequest.Facet
	nil,                                      // 7: resource.ResourceSearchRequest.FacetEntry
	(*ResourceSearchResponse_Facet)(nil),     // 8: resource.ResourceSearchResponse.Facet
	(*ResourceSearchResponse_TermFacet)(nil), // 9: resource.ResourceSearchResponse.TermFacet
	nil,                                      // 10: resource.ResourceSearchResponse.FacetEntry
	(*ErrorResult)(nil),                      // 11: resource.ErrorResult
	(*ListOptions)(nil),                      // 12: resource.ListOptions
	(*ResourceKey)(nil),                      // 13: resource.ResourceKey
	(*ResourceTable)(nil),                    // 14: resource.ResourceTable
}
var file_search_proto_depIdxs = []int32{
	11, // 0: resource.ResourceStatsResponse.error:type_name -> resource.ErrorResult
	4,  // 1: resource.ResourceStatsResponse.stats:type_name -> resource.ResourceStatsResponse.Stats
	12, // 2: resource.ResourceSearchRequest.options:type_name -> resource.ListOptions
	13, // 3: resource.ResourceSearchRequest.federated:type_name -> resource.ResourceKey
	5,  // 4: resource.ResourceSearchRequest.sortBy:type_name -> resource.ResourceSearchRequest.Sort
	7,  // 5: resource.ResourceSearchRequest.facet:type_name -> resource.ResourceSearchRequest.FacetEntry
	11, // 6: resource.ResourceSearchResponse.error:type_name -> resource.ErrorResult
	13, // 7: resource.ResourceSearchResponse.key:type_name -> resource.ResourceKey
	14, // 8: resource.ResourceSearchResponse.results:type_name -> resource.ResourceTable
	10, // 9: resource.ResourceSearchResponse.facet:type_name -> resource.ResourceSearchResponse.FacetEntry
	6,  // 10: resource.ResourceSearchRequest.FacetEntry.value:type_name -> resource.ResourceSearchRequest.Facet
	9,  // 11: resource.ResourceSearchResponse.Facet.terms:type_name -> resource.ResourceSearchResponse.TermFacet
	8,  // 12: resource.ResourceSearchResponse.FacetEntry.value:type_name -> resource.ResourceSearchResponse.Facet
	2,  // 13: resource.ResourceIndex.Search:input_type -> resource.ResourceSearchRequest
	0,  // 14: resource.ResourceIndex.GetStats:input_type -> resource.ResourceStatsRequest
	3,  // 15: resource.ResourceIndex.Search:output_type -> resource.ResourceSearchResponse
	1,  // 16: resource.ResourceIndex.GetStats:output_type -> resource.ResourceStatsResponse
	15, // [15:17] is the sub-list for method output_type
	13, // [13:15] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	file_resource_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_search_proto_rawDesc), len(file_search_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}
