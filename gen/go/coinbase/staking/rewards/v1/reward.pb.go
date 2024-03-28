// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: coinbase/staking/rewards/v1/reward.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The unit of time that the reward events were aggregated by.
type AggregationUnit int32

const (
	// Aggregation unit is unknown or unspecified.
	AggregationUnit_AGGREGATION_UNIT_UNSPECIFIED AggregationUnit = 0
	// Indicates the rewards are aggregated by epoch. This means there will be a 'epoch' field displaying the epoch on this resource.
	AggregationUnit_EPOCH AggregationUnit = 1
	// Indicates the rewards are aggregated by day. This means there will be a 'date' field displaying the date on this resource.
	AggregationUnit_DAY AggregationUnit = 2
)

// Enum value maps for AggregationUnit.
var (
	AggregationUnit_name = map[int32]string{
		0: "AGGREGATION_UNIT_UNSPECIFIED",
		1: "EPOCH",
		2: "DAY",
	}
	AggregationUnit_value = map[string]int32{
		"AGGREGATION_UNIT_UNSPECIFIED": 0,
		"EPOCH":                        1,
		"DAY":                          2,
	}
)

func (x AggregationUnit) Enum() *AggregationUnit {
	p := new(AggregationUnit)
	*p = x
	return p
}

func (x AggregationUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AggregationUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_coinbase_staking_rewards_v1_reward_proto_enumTypes[0].Descriptor()
}

func (AggregationUnit) Type() protoreflect.EnumType {
	return &file_coinbase_staking_rewards_v1_reward_proto_enumTypes[0]
}

func (x AggregationUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AggregationUnit.Descriptor instead.
func (AggregationUnit) EnumDescriptor() ([]byte, []int) {
	return file_coinbase_staking_rewards_v1_reward_proto_rawDescGZIP(), []int{0}
}

// The source of the USD price conversion.
type USDValue_Source int32

const (
	// The USD value source is unknown or unspecified.
	USDValue_SOURCE_UNSPECIFIED USDValue_Source = 0
	// The USD value source is the Coinbase exchange.
	USDValue_COINBASE_EXCHANGE USDValue_Source = 1
)

// Enum value maps for USDValue_Source.
var (
	USDValue_Source_name = map[int32]string{
		0: "SOURCE_UNSPECIFIED",
		1: "COINBASE_EXCHANGE",
	}
	USDValue_Source_value = map[string]int32{
		"SOURCE_UNSPECIFIED": 0,
		"COINBASE_EXCHANGE":  1,
	}
)

func (x USDValue_Source) Enum() *USDValue_Source {
	p := new(USDValue_Source)
	*p = x
	return p
}

func (x USDValue_Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (USDValue_Source) Descriptor() protoreflect.EnumDescriptor {
	return file_coinbase_staking_rewards_v1_reward_proto_enumTypes[1].Descriptor()
}

func (USDValue_Source) Type() protoreflect.EnumType {
	return &file_coinbase_staking_rewards_v1_reward_proto_enumTypes[1]
}

func (x USDValue_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use USDValue_Source.Descriptor instead.
func (USDValue_Source) EnumDescriptor() ([]byte, []int) {
	return file_coinbase_staking_rewards_v1_reward_proto_rawDescGZIP(), []int{1, 0}
}

// Rewards earned within a particular period of time.
// (-- api-linter: core::0123::resource-name-field=disabled
//
//	aip.dev/not-precedent: Not including a 'name' field till our data sources support a unique identifier --)
type Reward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address that earned this reward.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// The period identifier of this reward aggregation. ex: epoch number, date.
	//
	// Types that are assignable to PeriodIdentifier:
	//
	//	*Reward_Epoch
	//	*Reward_Date
	PeriodIdentifier isReward_PeriodIdentifier `protobuf_oneof:"period_identifier"`
	// The unit of time that the reward events were rolled up by.
	// Can be either "epoch" or "daily".
	AggregationUnit AggregationUnit `protobuf:"varint,4,opt,name=aggregation_unit,json=aggregationUnit,proto3,enum=coinbase.staking.rewards.v1.AggregationUnit" json:"aggregation_unit,omitempty"`
	// The starting time of this reward period. Returned when querying by epoch.
	// Timestamps are in UTC, conforming to the RFC-3339 spec (e.g. 2024-11-13T19:38:36Z).
	// Field currently unavailable. Coming soon.
	PeriodStartTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=period_start_time,json=periodStartTime,proto3" json:"period_start_time,omitempty"`
	// The ending time of this reward period. Returned when querying by epoch.
	// Timestamps are in UTC, conforming to the RFC-3339 spec (e.g. 2024-11-13T19:38:36Z).
	PeriodEndTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=period_end_time,json=periodEndTime,proto3" json:"period_end_time,omitempty"`
	// The amount earned in this time period in the native unit of the protocol.
	TotalEarnedNativeUnit *AssetAmount `protobuf:"bytes,7,opt,name=total_earned_native_unit,json=totalEarnedNativeUnit,proto3" json:"total_earned_native_unit,omitempty"`
	// The amount earned in this time period in USD. Calculated by getting each individual reward of this
	// time period and summing the USD value of each individual component. USD value is calculate at
	// the time each component was earned.
	TotalEarnedUsd []*USDValue `protobuf:"bytes,8,rep,name=total_earned_usd,json=totalEarnedUsd,proto3" json:"total_earned_usd,omitempty"`
	// A snapshot of the staking balance the end of this period.
	// Field currently unavailable. Coming soon.
	EndingBalance *Stake `protobuf:"bytes,9,opt,name=ending_balance,json=endingBalance,proto3" json:"ending_balance,omitempty"`
	// The protocol on which this reward was earned.
	Protocol string `protobuf:"bytes,11,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *Reward) Reset() {
	*x = Reward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_rewards_v1_reward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reward) ProtoMessage() {}

func (x *Reward) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_rewards_v1_reward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reward.ProtoReflect.Descriptor instead.
func (*Reward) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_rewards_v1_reward_proto_rawDescGZIP(), []int{0}
}

func (x *Reward) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (m *Reward) GetPeriodIdentifier() isReward_PeriodIdentifier {
	if m != nil {
		return m.PeriodIdentifier
	}
	return nil
}

func (x *Reward) GetEpoch() int64 {
	if x, ok := x.GetPeriodIdentifier().(*Reward_Epoch); ok {
		return x.Epoch
	}
	return 0
}

func (x *Reward) GetDate() string {
	if x, ok := x.GetPeriodIdentifier().(*Reward_Date); ok {
		return x.Date
	}
	return ""
}

func (x *Reward) GetAggregationUnit() AggregationUnit {
	if x != nil {
		return x.AggregationUnit
	}
	return AggregationUnit_AGGREGATION_UNIT_UNSPECIFIED
}

func (x *Reward) GetPeriodStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodStartTime
	}
	return nil
}

func (x *Reward) GetPeriodEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodEndTime
	}
	return nil
}

func (x *Reward) GetTotalEarnedNativeUnit() *AssetAmount {
	if x != nil {
		return x.TotalEarnedNativeUnit
	}
	return nil
}

func (x *Reward) GetTotalEarnedUsd() []*USDValue {
	if x != nil {
		return x.TotalEarnedUsd
	}
	return nil
}

func (x *Reward) GetEndingBalance() *Stake {
	if x != nil {
		return x.EndingBalance
	}
	return nil
}

func (x *Reward) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

type isReward_PeriodIdentifier interface {
	isReward_PeriodIdentifier()
}

type Reward_Epoch struct {
	// A unique identifier for the consensus-cycle of the blockchain.
	Epoch int64 `protobuf:"varint,13,opt,name=epoch,proto3,oneof"`
}

type Reward_Date struct {
	// The date of the reward in format 'YYYY-MM-DD' in UTC.
	// (-- api-linter: core::0142::time-field-type=disabled False positive. This isn't a timestamp, but a YYYY-MM-DD field --)
	Date string `protobuf:"bytes,14,opt,name=date,proto3,oneof"`
}

func (*Reward_Epoch) isReward_PeriodIdentifier() {}

func (*Reward_Date) isReward_PeriodIdentifier() {}

// Information regarding the USD value of a reward, with necessary context and metadata.
type USDValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source of the USD price conversion. Could be internal to Coinbase, and external source, or any other source.
	Source USDValue_Source `protobuf:"varint,1,opt,name=source,proto3,enum=coinbase.staking.rewards.v1.USDValue_Source" json:"source,omitempty"`
	// The timestamp at which the USD value was sourced to convert the value into USD.
	// This value is as close to the time the reward was earned as possible.
	// Timestamps are in UTC, conforming to the RFC-3339 spec (e.g. 2024-11-13T19:38:36Z).
	ConversionTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=conversion_time,json=conversionTime,proto3" json:"conversion_time,omitempty"`
	// The USD value of the reward at the conversion time.
	Amount *AssetAmount `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// The price of the native unit at the conversion time.
	ConversionPrice string `protobuf:"bytes,4,opt,name=conversion_price,json=conversionPrice,proto3" json:"conversion_price,omitempty"`
}

func (x *USDValue) Reset() {
	*x = USDValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_rewards_v1_reward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *USDValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*USDValue) ProtoMessage() {}

func (x *USDValue) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_rewards_v1_reward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use USDValue.ProtoReflect.Descriptor instead.
func (*USDValue) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_rewards_v1_reward_proto_rawDescGZIP(), []int{1}
}

func (x *USDValue) GetSource() USDValue_Source {
	if x != nil {
		return x.Source
	}
	return USDValue_SOURCE_UNSPECIFIED
}

func (x *USDValue) GetConversionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ConversionTime
	}
	return nil
}

func (x *USDValue) GetAmount() *AssetAmount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *USDValue) GetConversionPrice() string {
	if x != nil {
		return x.ConversionPrice
	}
	return ""
}

// The request message for ListRewards.
type ListRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The protocol that the rewards were earned on.
	// The response will only include rewards for the protocol specified here.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return. Maximum size of this value is 500.
	// If user supplies a value > 500, the API will truncate to 500.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token as part of the response of a previous call.
	// Provide this to retrieve the next page.
	//
	// When paginating, all other parameters must match the previous
	// request to list resources correctly.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// [AIP-160](https://google.aip.dev/160) format compliant filter. Supported protocols are 'ethereum', 'solana', and 'cosmos'.
	// Supplying other protocols will return an error.
	// * **Ethereum**:
	//   - Fields:
	//   - `address` - A ethereum validator public key.
	//   - `date` - A date in format 'YYYY-MM-DD'. Supports multiple comparisons (ex: '2024-01-15).
	//   - `period_end_time` - A timestamp in RFC-3339 format. Supports multiple comparisons (ex: '2024-01-01T00:00:00Z' and '2024-01-15T00:00:00Z').
	//   - Example(s):
	//   - `"address='0xac53512c39d0081ca4437c285305eb423f474e6153693c12fbba4a3df78bcaa3422b31d800c5bea71c1b017168a60474' AND date >= '2024-01-01' AND date < '2024-01-15'"`
	//   - `"address='0xac53512c39d0081ca4437c285305eb423f474e6153693c12fbba4a3df78bcaa3422b31d800c5bea71c1b017168a60474' AND period_end_time >= '2024-01-01T00:00:00Z' AND period_end_time < '2024-01-15T00:00:00Z'"`
	//   - `"address='0xac53512c39d0081ca4437c285305eb423f474e6153693c12fbba4a3df78bcaa3422b31d800c5bea71c1b017168a60474' AND date = '2024-01-01'"`
	//
	// * **Solana**:
	//   - Fields:
	//   - `address` - A solana validator or delegator address.
	//   - `epoch` - A solana epoch. Supports epoch comparisons (ex: `epoch >= 1000 AND epoch <= 2000`).
	//   - `period_end_time` - A timestamp in RFC-3339 format. Supports multiple comparisons (ex: '2024-01-01T00:00:00Z' and '2024-01-15T00:00:00Z').
	//   - Example(s):
	//   - `"address='beefKGBWeSpHzYBHZXwp5So7wdQGX6mu4ZHCsH3uTar' AND epoch >= 540 AND epoch < 550"`
	//   - `"address='beefKGBWeSpHzYBHZXwp5So7wdQGX6mu4ZHCsH3uTar' AND period_end_time >= '2024-01-01T00:00:00Z' AND period_end_time < '2024-01-15T00:00:00Z'"`
	//   - `"address='beefKGBWeSpHzYBHZXwp5So7wdQGX6mu4ZHCsH3uTar' AND epoch = 550"`
	//
	// * **Cosmos**:
	//   - Fields:
	//   - `address` - A cosmos validator or delegator address (ex: `cosmosvaloper1c4k24jzduc365kywrsvf5ujz4ya6mwympnc4en` and `cosmos1c4k24jzduc365kywrsvf5ujz4ya6mwymy8vq4q`)
	//   - `date` - A date in format 'YYYY-MM-DD'. Supports multiple comparisons (ex: '2024-01-15).
	//   - `period_end_time` - A timestamp in RFC-3339 format. Supports multiple comparisons (ex: '2024-01-01T00:00:00Z' and '2024-01-15T00:00:00Z').
	//   - Example(s):
	//   - `"address='cosmos1mfduj0qax6ut8rd6cfc4j0ds06z0mwlhrljhqh' AND date = '2024-11-16'"`
	//   - `"address='cosmos1mfduj0qax6ut8rd6cfc4j0ds06z0mwlhrljhqh' AND period_end_time >= '2024-01-01T00:00:00Z' AND period_end_time < '2024-01-15T00:00:00Z'"`
	//   - `"address='cosmos1mfduj0qax6ut8rd6cfc4j0ds06z0mwlhrljhqh' AND date = '2024-01-01'"`
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListRewardsRequest) Reset() {
	*x = ListRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_rewards_v1_reward_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRewardsRequest) ProtoMessage() {}

func (x *ListRewardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_rewards_v1_reward_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRewardsRequest.ProtoReflect.Descriptor instead.
func (*ListRewardsRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_rewards_v1_reward_proto_rawDescGZIP(), []int{2}
}

func (x *ListRewardsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListRewardsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRewardsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListRewardsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// The response message for ListRewards.
type ListRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rewards returned in this response.
	Rewards []*Reward `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
	// The page token the user must use in the next request if the next page is desired.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListRewardsResponse) Reset() {
	*x = ListRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_rewards_v1_reward_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRewardsResponse) ProtoMessage() {}

func (x *ListRewardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_rewards_v1_reward_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRewardsResponse.ProtoReflect.Descriptor instead.
func (*ListRewardsResponse) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_rewards_v1_reward_proto_rawDescGZIP(), []int{3}
}

func (x *ListRewardsResponse) GetRewards() []*Reward {
	if x != nil {
		return x.Rewards
	}
	return nil
}

func (x *ListRewardsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_coinbase_staking_rewards_v1_reward_proto protoreflect.FileDescriptor

var file_coinbase_staking_rewards_v1_reward_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x06, 0x0a, 0x06, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1b, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x19,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x5c, 0x0a, 0x10, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x4b, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x66, 0x0a,
	0x18, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x5f, 0x6e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x15,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x4e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x54, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65,
	0x61, 0x72, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x53,
	0x44, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x55, 0x73, 0x64, 0x12, 0x4e, 0x0a, 0x0e, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x5c, 0xea, 0x41,
	0x59, 0x0a, 0x1f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x7d, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x2f, 0x7b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x7d, 0x2a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x32, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4a,
	0x04, 0x08, 0x03, 0x10, 0x04, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0xcf, 0x02,
	0x0a, 0x08, 0x55, 0x53, 0x44, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x53, 0x44, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x45, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x37, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x49, 0x4e,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x01, 0x22,
	0xba, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x47, 0x0a, 0x0f, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x47, 0x47, 0x52,
	0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x50,
	0x4f, 0x43, 0x48, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x02, 0x42, 0x52,
	0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2d, 0x67, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_staking_rewards_v1_reward_proto_rawDescOnce sync.Once
	file_coinbase_staking_rewards_v1_reward_proto_rawDescData = file_coinbase_staking_rewards_v1_reward_proto_rawDesc
)

func file_coinbase_staking_rewards_v1_reward_proto_rawDescGZIP() []byte {
	file_coinbase_staking_rewards_v1_reward_proto_rawDescOnce.Do(func() {
		file_coinbase_staking_rewards_v1_reward_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_staking_rewards_v1_reward_proto_rawDescData)
	})
	return file_coinbase_staking_rewards_v1_reward_proto_rawDescData
}

var file_coinbase_staking_rewards_v1_reward_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_coinbase_staking_rewards_v1_reward_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_coinbase_staking_rewards_v1_reward_proto_goTypes = []interface{}{
	(AggregationUnit)(0),          // 0: coinbase.staking.rewards.v1.AggregationUnit
	(USDValue_Source)(0),          // 1: coinbase.staking.rewards.v1.USDValue.Source
	(*Reward)(nil),                // 2: coinbase.staking.rewards.v1.Reward
	(*USDValue)(nil),              // 3: coinbase.staking.rewards.v1.USDValue
	(*ListRewardsRequest)(nil),    // 4: coinbase.staking.rewards.v1.ListRewardsRequest
	(*ListRewardsResponse)(nil),   // 5: coinbase.staking.rewards.v1.ListRewardsResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*AssetAmount)(nil),           // 7: coinbase.staking.rewards.v1.AssetAmount
	(*Stake)(nil),                 // 8: coinbase.staking.rewards.v1.Stake
}
var file_coinbase_staking_rewards_v1_reward_proto_depIdxs = []int32{
	0,  // 0: coinbase.staking.rewards.v1.Reward.aggregation_unit:type_name -> coinbase.staking.rewards.v1.AggregationUnit
	6,  // 1: coinbase.staking.rewards.v1.Reward.period_start_time:type_name -> google.protobuf.Timestamp
	6,  // 2: coinbase.staking.rewards.v1.Reward.period_end_time:type_name -> google.protobuf.Timestamp
	7,  // 3: coinbase.staking.rewards.v1.Reward.total_earned_native_unit:type_name -> coinbase.staking.rewards.v1.AssetAmount
	3,  // 4: coinbase.staking.rewards.v1.Reward.total_earned_usd:type_name -> coinbase.staking.rewards.v1.USDValue
	8,  // 5: coinbase.staking.rewards.v1.Reward.ending_balance:type_name -> coinbase.staking.rewards.v1.Stake
	1,  // 6: coinbase.staking.rewards.v1.USDValue.source:type_name -> coinbase.staking.rewards.v1.USDValue.Source
	6,  // 7: coinbase.staking.rewards.v1.USDValue.conversion_time:type_name -> google.protobuf.Timestamp
	7,  // 8: coinbase.staking.rewards.v1.USDValue.amount:type_name -> coinbase.staking.rewards.v1.AssetAmount
	2,  // 9: coinbase.staking.rewards.v1.ListRewardsResponse.rewards:type_name -> coinbase.staking.rewards.v1.Reward
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_coinbase_staking_rewards_v1_reward_proto_init() }
func file_coinbase_staking_rewards_v1_reward_proto_init() {
	if File_coinbase_staking_rewards_v1_reward_proto != nil {
		return
	}
	file_coinbase_staking_rewards_v1_common_proto_init()
	file_coinbase_staking_rewards_v1_stake_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coinbase_staking_rewards_v1_reward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reward); i {
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
		file_coinbase_staking_rewards_v1_reward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*USDValue); i {
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
		file_coinbase_staking_rewards_v1_reward_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRewardsRequest); i {
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
		file_coinbase_staking_rewards_v1_reward_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRewardsResponse); i {
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
	file_coinbase_staking_rewards_v1_reward_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Reward_Epoch)(nil),
		(*Reward_Date)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_staking_rewards_v1_reward_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_staking_rewards_v1_reward_proto_goTypes,
		DependencyIndexes: file_coinbase_staking_rewards_v1_reward_proto_depIdxs,
		EnumInfos:         file_coinbase_staking_rewards_v1_reward_proto_enumTypes,
		MessageInfos:      file_coinbase_staking_rewards_v1_reward_proto_msgTypes,
	}.Build()
	File_coinbase_staking_rewards_v1_reward_proto = out.File
	file_coinbase_staking_rewards_v1_reward_proto_rawDesc = nil
	file_coinbase_staking_rewards_v1_reward_proto_goTypes = nil
	file_coinbase_staking_rewards_v1_reward_proto_depIdxs = nil
}
