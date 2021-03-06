// Code generated by protoc-gen-go. DO NOT EDIT.
// source: beeswax/logs/ad_log.proto

/*
Package logs is a generated protocol buffer package.

It is generated from these files:
	beeswax/logs/ad_log.proto

It has these top-level messages:
	RequestLogMessage
	ImpressionLogMessage
	ClickLogMessage
	ActivityLogMessage
	ConversionLogMessage
	AdLogMessage
*/
package logs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bid "beeswax/bid"
import bid1 "beeswax/bid"
import base "beeswax/base"
import openrtb2 "beeswax/openrtb"
import openrtb "beeswax/openrtb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClickLogMessage_Type int32

const (
	ClickLogMessage_UNKNOWN   ClickLogMessage_Type = 0
	ClickLogMessage_REGULAR   ClickLogMessage_Type = 1
	ClickLogMessage_COMPANION ClickLogMessage_Type = 2
)

var ClickLogMessage_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "REGULAR",
	2: "COMPANION",
}
var ClickLogMessage_Type_value = map[string]int32{
	"UNKNOWN":   0,
	"REGULAR":   1,
	"COMPANION": 2,
}

func (x ClickLogMessage_Type) Enum() *ClickLogMessage_Type {
	p := new(ClickLogMessage_Type)
	*p = x
	return p
}
func (x ClickLogMessage_Type) String() string {
	return proto.EnumName(ClickLogMessage_Type_name, int32(x))
}
func (x *ClickLogMessage_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClickLogMessage_Type_value, data, "ClickLogMessage_Type")
	if err != nil {
		return err
	}
	*x = ClickLogMessage_Type(value)
	return nil
}
func (ClickLogMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type ActivityLogMessage_Type int32

const (
	ActivityLogMessage_START          ActivityLogMessage_Type = 1
	ActivityLogMessage_Q1             ActivityLogMessage_Type = 2
	ActivityLogMessage_MID            ActivityLogMessage_Type = 3
	ActivityLogMessage_Q3             ActivityLogMessage_Type = 4
	ActivityLogMessage_COMPLETE       ActivityLogMessage_Type = 5
	ActivityLogMessage_SKIP           ActivityLogMessage_Type = 6
	ActivityLogMessage_MUTE           ActivityLogMessage_Type = 7
	ActivityLogMessage_UNMUTE         ActivityLogMessage_Type = 8
	ActivityLogMessage_PAUSE          ActivityLogMessage_Type = 9
	ActivityLogMessage_RESUME         ActivityLogMessage_Type = 10
	ActivityLogMessage_FULLSCREEN     ActivityLogMessage_Type = 11
	ActivityLogMessage_CLOSE          ActivityLogMessage_Type = 12
	ActivityLogMessage_COMPANION_VIEW ActivityLogMessage_Type = 13
	// Deprecated. Companion click will be tracked as ClickEvent instead of ActivityEvent
	// This should be eventually removed by end of Q2 of 2017
	ActivityLogMessage_DEPRECATED_COMPANION_CLICK ActivityLogMessage_Type = 14
)

var ActivityLogMessage_Type_name = map[int32]string{
	1:  "START",
	2:  "Q1",
	3:  "MID",
	4:  "Q3",
	5:  "COMPLETE",
	6:  "SKIP",
	7:  "MUTE",
	8:  "UNMUTE",
	9:  "PAUSE",
	10: "RESUME",
	11: "FULLSCREEN",
	12: "CLOSE",
	13: "COMPANION_VIEW",
	14: "DEPRECATED_COMPANION_CLICK",
}
var ActivityLogMessage_Type_value = map[string]int32{
	"START":                      1,
	"Q1":                         2,
	"MID":                        3,
	"Q3":                         4,
	"COMPLETE":                   5,
	"SKIP":                       6,
	"MUTE":                       7,
	"UNMUTE":                     8,
	"PAUSE":                      9,
	"RESUME":                     10,
	"FULLSCREEN":                 11,
	"CLOSE":                      12,
	"COMPANION_VIEW":             13,
	"DEPRECATED_COMPANION_CLICK": 14,
}

func (x ActivityLogMessage_Type) Enum() *ActivityLogMessage_Type {
	p := new(ActivityLogMessage_Type)
	*p = x
	return p
}
func (x ActivityLogMessage_Type) String() string {
	return proto.EnumName(ActivityLogMessage_Type_name, int32(x))
}
func (x *ActivityLogMessage_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ActivityLogMessage_Type_value, data, "ActivityLogMessage_Type")
	if err != nil {
		return err
	}
	*x = ActivityLogMessage_Type(value)
	return nil
}
func (ActivityLogMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// Next Tag: 6
// RequestLogMessage is created for every incoming bid request from the exchange.
type RequestLogMessage struct {
	Auctionid       *base.EventId                   `protobuf:"bytes,1,opt,name=auctionid" json:"auctionid,omitempty"`
	InventorySource *openrtb.Enums_Inventory_Source `protobuf:"varint,2,opt,name=inventory_source,json=inventorySource,enum=openrtb.Enums_Inventory_Source" json:"inventory_source,omitempty"`
	Bidrequest      *openrtb2.BidRequest            `protobuf:"bytes,3,opt,name=bidrequest" json:"bidrequest,omitempty"`
	// All the adcandidates that matched targeting, creative attribute
	// and budget filtering and finally sent to the exchange.
	Adcandidates     []*bid.Adcandidate               `protobuf:"bytes,4,rep,name=adcandidates" json:"adcandidates,omitempty"`
	AgentData        *bid1.BidAgentResponse_AgentData `protobuf:"bytes,5,opt,name=agent_data,json=agentData" json:"agent_data,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *RequestLogMessage) Reset()                    { *m = RequestLogMessage{} }
func (m *RequestLogMessage) String() string            { return proto.CompactTextString(m) }
func (*RequestLogMessage) ProtoMessage()               {}
func (*RequestLogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestLogMessage) GetAuctionid() *base.EventId {
	if m != nil {
		return m.Auctionid
	}
	return nil
}

func (m *RequestLogMessage) GetInventorySource() openrtb.Enums_Inventory_Source {
	if m != nil && m.InventorySource != nil {
		return *m.InventorySource
	}
	return openrtb.Enums_Inventory_UNKNOWN_SOURCE
}

func (m *RequestLogMessage) GetBidrequest() *openrtb2.BidRequest {
	if m != nil {
		return m.Bidrequest
	}
	return nil
}

func (m *RequestLogMessage) GetAdcandidates() []*bid.Adcandidate {
	if m != nil {
		return m.Adcandidates
	}
	return nil
}

func (m *RequestLogMessage) GetAgentData() *bid1.BidAgentResponse_AgentData {
	if m != nil {
		return m.AgentData
	}
	return nil
}

// Next Tag: 13
// ImpressionLogMessage is created for every won auction.
type ImpressionLogMessage struct {
	Auctionid *base.EventId `protobuf:"bytes,1,opt,name=auctionid" json:"auctionid,omitempty"`
	// The unique identifier of the auction per buzz_key.
	// Format: <auctionid.timestamp>.<auctionid.hostid>.<auctionid.tid>.<buzz_key>
	AuctionidStr *string         `protobuf:"bytes,12,opt,name=auctionid_str,json=auctionidStr" json:"auctionid_str,omitempty"`
	Adgroupid    *base.AdGroupId `protobuf:"bytes,2,opt,name=adgroupid" json:"adgroupid,omitempty"`
	CreativeId   *uint64         `protobuf:"varint,10,opt,name=creative_id,json=creativeId" json:"creative_id,omitempty"`
	AdvertiserId *uint64         `protobuf:"varint,11,opt,name=advertiser_id,json=advertiserId" json:"advertiser_id,omitempty"`
	// Exchange bid price is the final price submitted to exchange.
	ExchangeBidPriceMicrosUsd *uint64 `protobuf:"varint,9,opt,name=exchange_bid_price_micros_usd,json=exchangeBidPriceMicrosUsd" json:"exchange_bid_price_micros_usd,omitempty"`
	// Represents the bid price obtained from the bid agent.
	BidPriceMicrosUsd *uint64                         `protobuf:"varint,3,opt,name=bid_price_micros_usd,json=bidPriceMicrosUsd" json:"bid_price_micros_usd,omitempty"`
	InventorySource   *openrtb.Enums_Inventory_Source `protobuf:"varint,4,opt,name=inventory_source,json=inventorySource,enum=openrtb.Enums_Inventory_Source" json:"inventory_source,omitempty"`
	UserId            *string                         `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Represents the winning price in the auction.
	WinPriceMicrosUsd *uint64                             `protobuf:"varint,6,opt,name=win_price_micros_usd,json=winPriceMicrosUsd" json:"win_price_micros_usd,omitempty"`
	RxTimestampUsecs  *uint64                             `protobuf:"varint,7,opt,name=rx_timestamp_usecs,json=rxTimestampUsecs" json:"rx_timestamp_usecs,omitempty"`
	DataCenter        *openrtb.Enums_Inventory_DataCenter `protobuf:"varint,8,opt,name=data_center,json=dataCenter,enum=openrtb.Enums_Inventory_DataCenter" json:"data_center,omitempty"`
	XXX_unrecognized  []byte                              `json:"-"`
}

func (m *ImpressionLogMessage) Reset()                    { *m = ImpressionLogMessage{} }
func (m *ImpressionLogMessage) String() string            { return proto.CompactTextString(m) }
func (*ImpressionLogMessage) ProtoMessage()               {}
func (*ImpressionLogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ImpressionLogMessage) GetAuctionid() *base.EventId {
	if m != nil {
		return m.Auctionid
	}
	return nil
}

func (m *ImpressionLogMessage) GetAuctionidStr() string {
	if m != nil && m.AuctionidStr != nil {
		return *m.AuctionidStr
	}
	return ""
}

func (m *ImpressionLogMessage) GetAdgroupid() *base.AdGroupId {
	if m != nil {
		return m.Adgroupid
	}
	return nil
}

func (m *ImpressionLogMessage) GetCreativeId() uint64 {
	if m != nil && m.CreativeId != nil {
		return *m.CreativeId
	}
	return 0
}

func (m *ImpressionLogMessage) GetAdvertiserId() uint64 {
	if m != nil && m.AdvertiserId != nil {
		return *m.AdvertiserId
	}
	return 0
}

func (m *ImpressionLogMessage) GetExchangeBidPriceMicrosUsd() uint64 {
	if m != nil && m.ExchangeBidPriceMicrosUsd != nil {
		return *m.ExchangeBidPriceMicrosUsd
	}
	return 0
}

func (m *ImpressionLogMessage) GetBidPriceMicrosUsd() uint64 {
	if m != nil && m.BidPriceMicrosUsd != nil {
		return *m.BidPriceMicrosUsd
	}
	return 0
}

func (m *ImpressionLogMessage) GetInventorySource() openrtb.Enums_Inventory_Source {
	if m != nil && m.InventorySource != nil {
		return *m.InventorySource
	}
	return openrtb.Enums_Inventory_UNKNOWN_SOURCE
}

func (m *ImpressionLogMessage) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *ImpressionLogMessage) GetWinPriceMicrosUsd() uint64 {
	if m != nil && m.WinPriceMicrosUsd != nil {
		return *m.WinPriceMicrosUsd
	}
	return 0
}

func (m *ImpressionLogMessage) GetRxTimestampUsecs() uint64 {
	if m != nil && m.RxTimestampUsecs != nil {
		return *m.RxTimestampUsecs
	}
	return 0
}

func (m *ImpressionLogMessage) GetDataCenter() openrtb.Enums_Inventory_DataCenter {
	if m != nil && m.DataCenter != nil {
		return *m.DataCenter
	}
	return openrtb.Enums_Inventory_UNKNOWN
}

// Next Tag: 11
// ClickLogMessage is created for clicks received.
type ClickLogMessage struct {
	Auctionid *base.EventId `protobuf:"bytes,1,opt,name=auctionid" json:"auctionid,omitempty"`
	// The unique identifier of the auction per buzz_key.
	// Format: <auctionid.timestamp>.<auctionid.hostid>.<auctionid.tid>.<buzz_key>
	AuctionidStr    *string                         `protobuf:"bytes,9,opt,name=auctionid_str,json=auctionidStr" json:"auctionid_str,omitempty"`
	Adgroupid       *base.AdGroupId                 `protobuf:"bytes,2,opt,name=adgroupid" json:"adgroupid,omitempty"`
	CreativeId      *uint64                         `protobuf:"varint,7,opt,name=creative_id,json=creativeId" json:"creative_id,omitempty"`
	AdvertiserId    *uint64                         `protobuf:"varint,8,opt,name=advertiser_id,json=advertiserId" json:"advertiser_id,omitempty"`
	InventorySource *openrtb.Enums_Inventory_Source `protobuf:"varint,3,opt,name=inventory_source,json=inventorySource,enum=openrtb.Enums_Inventory_Source" json:"inventory_source,omitempty"`
	UserId          *string                         `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Represents optional data sent by the customer.
	UserData         *string               `protobuf:"bytes,5,opt,name=user_data,json=userData" json:"user_data,omitempty"`
	RxTimestampUsecs *uint64               `protobuf:"varint,6,opt,name=rx_timestamp_usecs,json=rxTimestampUsecs" json:"rx_timestamp_usecs,omitempty"`
	ClickType        *ClickLogMessage_Type `protobuf:"varint,10,opt,name=click_type,json=clickType,enum=logs.ClickLogMessage_Type,def=1" json:"click_type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *ClickLogMessage) Reset()                    { *m = ClickLogMessage{} }
func (m *ClickLogMessage) String() string            { return proto.CompactTextString(m) }
func (*ClickLogMessage) ProtoMessage()               {}
func (*ClickLogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_ClickLogMessage_ClickType ClickLogMessage_Type = ClickLogMessage_REGULAR

func (m *ClickLogMessage) GetAuctionid() *base.EventId {
	if m != nil {
		return m.Auctionid
	}
	return nil
}

func (m *ClickLogMessage) GetAuctionidStr() string {
	if m != nil && m.AuctionidStr != nil {
		return *m.AuctionidStr
	}
	return ""
}

func (m *ClickLogMessage) GetAdgroupid() *base.AdGroupId {
	if m != nil {
		return m.Adgroupid
	}
	return nil
}

func (m *ClickLogMessage) GetCreativeId() uint64 {
	if m != nil && m.CreativeId != nil {
		return *m.CreativeId
	}
	return 0
}

func (m *ClickLogMessage) GetAdvertiserId() uint64 {
	if m != nil && m.AdvertiserId != nil {
		return *m.AdvertiserId
	}
	return 0
}

func (m *ClickLogMessage) GetInventorySource() openrtb.Enums_Inventory_Source {
	if m != nil && m.InventorySource != nil {
		return *m.InventorySource
	}
	return openrtb.Enums_Inventory_UNKNOWN_SOURCE
}

func (m *ClickLogMessage) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *ClickLogMessage) GetUserData() string {
	if m != nil && m.UserData != nil {
		return *m.UserData
	}
	return ""
}

func (m *ClickLogMessage) GetRxTimestampUsecs() uint64 {
	if m != nil && m.RxTimestampUsecs != nil {
		return *m.RxTimestampUsecs
	}
	return 0
}

func (m *ClickLogMessage) GetClickType() ClickLogMessage_Type {
	if m != nil && m.ClickType != nil {
		return *m.ClickType
	}
	return Default_ClickLogMessage_ClickType
}

// Next Tag: 11
// ActivityLogMessage is created for video and activity events
// like video "START", "SKIP", etc.
type ActivityLogMessage struct {
	Auctionid *base.EventId `protobuf:"bytes,1,opt,name=auctionid" json:"auctionid,omitempty"`
	// The unique identifier of the auction per buzz_key.
	// Format: <auctionid.timestamp>.<auctionid.hostid>.<auctionid.tid>.<buzz_key>
	AuctionidStr    *string                         `protobuf:"bytes,10,opt,name=auctionid_str,json=auctionidStr" json:"auctionid_str,omitempty"`
	Adgroupid       *base.AdGroupId                 `protobuf:"bytes,2,opt,name=adgroupid" json:"adgroupid,omitempty"`
	CreativeId      *uint64                         `protobuf:"varint,8,opt,name=creative_id,json=creativeId" json:"creative_id,omitempty"`
	AdvertiserId    *uint64                         `protobuf:"varint,9,opt,name=advertiser_id,json=advertiserId" json:"advertiser_id,omitempty"`
	InventorySource *openrtb.Enums_Inventory_Source `protobuf:"varint,3,opt,name=inventory_source,json=inventorySource,enum=openrtb.Enums_Inventory_Source" json:"inventory_source,omitempty"`
	Type            *ActivityLogMessage_Type        `protobuf:"varint,4,opt,name=type,enum=logs.ActivityLogMessage_Type" json:"type,omitempty"`
	UserId          *string                         `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Represents optional data sent by the customer.
	UserData         *string `protobuf:"bytes,6,opt,name=user_data,json=userData" json:"user_data,omitempty"`
	RxTimestampUsecs *uint64 `protobuf:"varint,7,opt,name=rx_timestamp_usecs,json=rxTimestampUsecs" json:"rx_timestamp_usecs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ActivityLogMessage) Reset()                    { *m = ActivityLogMessage{} }
func (m *ActivityLogMessage) String() string            { return proto.CompactTextString(m) }
func (*ActivityLogMessage) ProtoMessage()               {}
func (*ActivityLogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ActivityLogMessage) GetAuctionid() *base.EventId {
	if m != nil {
		return m.Auctionid
	}
	return nil
}

func (m *ActivityLogMessage) GetAuctionidStr() string {
	if m != nil && m.AuctionidStr != nil {
		return *m.AuctionidStr
	}
	return ""
}

func (m *ActivityLogMessage) GetAdgroupid() *base.AdGroupId {
	if m != nil {
		return m.Adgroupid
	}
	return nil
}

func (m *ActivityLogMessage) GetCreativeId() uint64 {
	if m != nil && m.CreativeId != nil {
		return *m.CreativeId
	}
	return 0
}

func (m *ActivityLogMessage) GetAdvertiserId() uint64 {
	if m != nil && m.AdvertiserId != nil {
		return *m.AdvertiserId
	}
	return 0
}

func (m *ActivityLogMessage) GetInventorySource() openrtb.Enums_Inventory_Source {
	if m != nil && m.InventorySource != nil {
		return *m.InventorySource
	}
	return openrtb.Enums_Inventory_UNKNOWN_SOURCE
}

func (m *ActivityLogMessage) GetType() ActivityLogMessage_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ActivityLogMessage_START
}

func (m *ActivityLogMessage) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *ActivityLogMessage) GetUserData() string {
	if m != nil && m.UserData != nil {
		return *m.UserData
	}
	return ""
}

func (m *ActivityLogMessage) GetRxTimestampUsecs() uint64 {
	if m != nil && m.RxTimestampUsecs != nil {
		return *m.RxTimestampUsecs
	}
	return 0
}

// Next Tag: 9
// Conversion tag urls can be created using the Buzz API.
// Whenever a conversion tag fires, Beeswax records that event into a
// ConversionLogMessage.
// Currently, there are two scenarios in which a conversion tag may fire.
// 1. Conversion Events
//   This occurs whenever the customer trafficks the conversion tag on a site/app
//   and the tag fires whenever a user visits the site/app. For example, a conversion tag
//   may be trafficked on the shopping cart page of a retail portal.
//   For conversion events, the following fields are always expected to be populated:
//   a) tag_id
//   b) user_id (if available)
//
// 2 Postback Events
//  In this scenario, the conversion attribution is determined by a third party and the
//  conversion tag is used by the third party to credit a particular impression/click for
//  the conversion. For postback events, the following fields are always expected to be
//  populated:
//   a) tag_id
//   b) auctionid
//
type ConversionLogMessage struct {
	UserId *string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Customer defined transaction value e.g. Price in shopping cart.
	Value *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	// Usually, the number of items purchased
	Order *string `protobuf:"bytes,3,opt,name=order" json:"order,omitempty"`
	// Customer defined unique join key. Also works as cache buster.
	Ord *string `protobuf:"bytes,4,opt,name=ord" json:"ord,omitempty"`
	// Event Id corresponding to the tag created via the Buzz API
	TagId *int32 `protobuf:"varint,5,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	// Unique buzz_key of the customer.
	BuzzKey *string `protobuf:"bytes,6,opt,name=buzz_key,json=buzzKey" json:"buzz_key,omitempty"`
	// Timestamp (usecs) at which event was received by Beeswax.
	RxTimestamp *uint64 `protobuf:"varint,7,opt,name=rx_timestamp,json=rxTimestamp" json:"rx_timestamp,omitempty"`
	// auction id (Filled only for postback events)
	Auctionid        *string `protobuf:"bytes,8,opt,name=auctionid" json:"auctionid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ConversionLogMessage) Reset()                    { *m = ConversionLogMessage{} }
func (m *ConversionLogMessage) String() string            { return proto.CompactTextString(m) }
func (*ConversionLogMessage) ProtoMessage()               {}
func (*ConversionLogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConversionLogMessage) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *ConversionLogMessage) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *ConversionLogMessage) GetOrder() string {
	if m != nil && m.Order != nil {
		return *m.Order
	}
	return ""
}

func (m *ConversionLogMessage) GetOrd() string {
	if m != nil && m.Ord != nil {
		return *m.Ord
	}
	return ""
}

func (m *ConversionLogMessage) GetTagId() int32 {
	if m != nil && m.TagId != nil {
		return *m.TagId
	}
	return 0
}

func (m *ConversionLogMessage) GetBuzzKey() string {
	if m != nil && m.BuzzKey != nil {
		return *m.BuzzKey
	}
	return ""
}

func (m *ConversionLogMessage) GetRxTimestamp() uint64 {
	if m != nil && m.RxTimestamp != nil {
		return *m.RxTimestamp
	}
	return 0
}

func (m *ConversionLogMessage) GetAuctionid() string {
	if m != nil && m.Auctionid != nil {
		return *m.Auctionid
	}
	return ""
}

// Note: BeeswaxIO DOES NOT aggregate the ad log objects on same "auction_id";
// All events are logged to Kinesis as and when they arrive in the system.
type AdLogMessage struct {
	Request          *RequestLogMessage    `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Impression       *ImpressionLogMessage `protobuf:"bytes,2,opt,name=impression" json:"impression,omitempty"`
	Click            *ClickLogMessage      `protobuf:"bytes,3,opt,name=click" json:"click,omitempty"`
	Activity         []*ActivityLogMessage `protobuf:"bytes,4,rep,name=activity" json:"activity,omitempty"`
	Conversion       *ConversionLogMessage `protobuf:"bytes,5,opt,name=conversion" json:"conversion,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *AdLogMessage) Reset()                    { *m = AdLogMessage{} }
func (m *AdLogMessage) String() string            { return proto.CompactTextString(m) }
func (*AdLogMessage) ProtoMessage()               {}
func (*AdLogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AdLogMessage) GetRequest() *RequestLogMessage {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AdLogMessage) GetImpression() *ImpressionLogMessage {
	if m != nil {
		return m.Impression
	}
	return nil
}

func (m *AdLogMessage) GetClick() *ClickLogMessage {
	if m != nil {
		return m.Click
	}
	return nil
}

func (m *AdLogMessage) GetActivity() []*ActivityLogMessage {
	if m != nil {
		return m.Activity
	}
	return nil
}

func (m *AdLogMessage) GetConversion() *ConversionLogMessage {
	if m != nil {
		return m.Conversion
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestLogMessage)(nil), "logs.RequestLogMessage")
	proto.RegisterType((*ImpressionLogMessage)(nil), "logs.ImpressionLogMessage")
	proto.RegisterType((*ClickLogMessage)(nil), "logs.ClickLogMessage")
	proto.RegisterType((*ActivityLogMessage)(nil), "logs.ActivityLogMessage")
	proto.RegisterType((*ConversionLogMessage)(nil), "logs.ConversionLogMessage")
	proto.RegisterType((*AdLogMessage)(nil), "logs.AdLogMessage")
	proto.RegisterEnum("logs.ClickLogMessage_Type", ClickLogMessage_Type_name, ClickLogMessage_Type_value)
	proto.RegisterEnum("logs.ActivityLogMessage_Type", ActivityLogMessage_Type_name, ActivityLogMessage_Type_value)
}

func init() { proto.RegisterFile("beeswax/logs/ad_log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1062 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x72, 0xe3, 0x44,
	0x10, 0xc6, 0xb6, 0xfc, 0xa3, 0xb6, 0x93, 0x68, 0x87, 0x6c, 0xad, 0x12, 0x08, 0x09, 0x5e, 0x0e,
	0xa9, 0x5a, 0x90, 0x2b, 0xd9, 0x3d, 0xed, 0x81, 0xc2, 0x71, 0xc4, 0x96, 0x88, 0xed, 0x84, 0xb1,
	0xcd, 0x1e, 0x55, 0x63, 0xcd, 0x94, 0x99, 0xda, 0x58, 0x32, 0xd2, 0xc8, 0x89, 0xf7, 0x59, 0x78,
	0x0f, 0x6e, 0x3c, 0x08, 0x27, 0x5e, 0x84, 0x82, 0x9a, 0xd1, 0x8f, 0x9d, 0xc4, 0xae, 0xa5, 0x96,
	0x70, 0xf2, 0x4c, 0xf7, 0xd7, 0x33, 0xad, 0xfe, 0xbe, 0x69, 0x37, 0xec, 0x8d, 0x19, 0x8b, 0x6e,
	0xc8, 0x6d, 0xeb, 0x3a, 0x98, 0x44, 0x2d, 0x42, 0xdd, 0xeb, 0x60, 0x62, 0xcd, 0xc2, 0x40, 0x04,
	0x48, 0x93, 0xa6, 0xfd, 0x83, 0x0c, 0x30, 0xe6, 0xb4, 0x45, 0xa8, 0x47, 0x7c, 0xca, 0x29, 0x11,
	0x2c, 0x01, 0xed, 0xef, 0xad, 0xba, 0x43, 0xf6, 0x4b, 0xcc, 0x22, 0x91, 0xba, 0xf6, 0x73, 0x17,
	0x89, 0x58, 0x8b, 0xcd, 0x99, 0x2f, 0x38, 0x4d, 0x7d, 0xf9, 0xa9, 0xc1, 0x8c, 0xf9, 0xa1, 0x18,
	0x67, 0xbf, 0xa9, 0xfb, 0xab, 0x0d, 0x6e, 0xd7, 0x0b, 0xa6, 0xd3, 0xc0, 0x4f, 0x50, 0xcd, 0xdf,
	0x8a, 0xf0, 0x04, 0x27, 0x57, 0x76, 0x83, 0x49, 0x8f, 0x45, 0x11, 0x99, 0x30, 0xf4, 0x02, 0x74,
	0x12, 0x7b, 0x82, 0x07, 0x3e, 0xa7, 0x66, 0xe1, 0xa8, 0x70, 0x5c, 0x3f, 0xdd, 0xb2, 0x64, 0x0a,
	0x96, 0x2d, 0x53, 0x70, 0x28, 0x5e, 0xfa, 0xd1, 0x0f, 0x60, 0x70, 0x5f, 0x9a, 0x83, 0x70, 0xe1,
	0x46, 0x41, 0x1c, 0x7a, 0xcc, 0x2c, 0x1e, 0x15, 0x8e, 0xb7, 0x4f, 0x0f, 0xad, 0x2c, 0x25, 0xdb,
	0x8f, 0xa7, 0x91, 0xe5, 0x64, 0x30, 0x6b, 0xa0, 0x60, 0x78, 0x27, 0x0f, 0x4c, 0x0c, 0xe8, 0x25,
	0xc0, 0x98, 0xd3, 0xb4, 0x06, 0x66, 0x49, 0xdd, 0xfc, 0x69, 0x7e, 0xca, 0x19, 0xa7, 0x69, 0xae,
	0x78, 0x05, 0x86, 0x5e, 0x41, 0x63, 0xa5, 0xa8, 0x91, 0xa9, 0x1d, 0x95, 0x8e, 0xeb, 0xa7, 0x86,
	0x35, 0xe6, 0xd4, 0x6a, 0x2f, 0x1d, 0xf8, 0x0e, 0x0a, 0x7d, 0x0b, 0x40, 0x26, 0xcc, 0x17, 0x2e,
	0x25, 0x82, 0x98, 0x65, 0x75, 0xd5, 0xa1, 0x8a, 0x39, 0xe3, 0xb4, 0x2d, 0x3d, 0x98, 0x45, 0xb3,
	0xc0, 0x8f, 0x98, 0xa5, 0x76, 0xe7, 0x44, 0x10, 0xac, 0x93, 0x6c, 0xd9, 0xfc, 0x43, 0x83, 0x5d,
	0x67, 0x3a, 0x0b, 0x59, 0x14, 0xf1, 0xc0, 0xff, 0xd8, 0xe2, 0x3d, 0x87, 0xad, 0x7c, 0xe3, 0x46,
	0x22, 0x34, 0x1b, 0x47, 0x85, 0x63, 0x1d, 0x37, 0x72, 0xe3, 0x40, 0x84, 0xe8, 0x1b, 0xd0, 0x09,
	0x9d, 0x84, 0x41, 0x3c, 0xe3, 0x54, 0x95, 0xb6, 0x7e, 0xba, 0x93, 0x9c, 0xd8, 0xa6, 0x6f, 0xa4,
	0x59, 0x9d, 0x99, 0x21, 0xd0, 0x21, 0xd4, 0xbd, 0x90, 0x11, 0xc1, 0xe7, 0xcc, 0xe5, 0xd4, 0x84,
	0xa3, 0xc2, 0xb1, 0x86, 0x21, 0x33, 0x39, 0xc9, 0xa5, 0x74, 0xce, 0x42, 0xc1, 0x23, 0x16, 0x4a,
	0x48, 0x5d, 0x41, 0x1a, 0x4b, 0xa3, 0x43, 0xd1, 0x77, 0x70, 0xc0, 0x6e, 0xbd, 0x9f, 0x89, 0x3f,
	0x61, 0xee, 0x98, 0x53, 0x77, 0x16, 0x72, 0x8f, 0xb9, 0x53, 0xee, 0x85, 0x41, 0xe4, 0xc6, 0x11,
	0x35, 0x75, 0x15, 0xb4, 0x97, 0x81, 0xce, 0x38, 0xbd, 0x92, 0x90, 0x9e, 0x42, 0x8c, 0x22, 0x8a,
	0x5a, 0xb0, 0xbb, 0x36, 0xb0, 0xa4, 0x02, 0x9f, 0x8c, 0x1f, 0x04, 0xac, 0x53, 0x92, 0xf6, 0x91,
	0x4a, 0x7a, 0x06, 0xd5, 0x38, 0xfd, 0xba, 0xb2, 0x2a, 0x69, 0x25, 0x4e, 0xbe, 0xab, 0x05, 0xbb,
	0x37, 0xdc, 0x7f, 0x98, 0x55, 0x25, 0xc9, 0xea, 0x86, 0xfb, 0xf7, 0xb2, 0xfa, 0x1a, 0x50, 0x78,
	0xeb, 0x0a, 0x3e, 0x65, 0x91, 0x20, 0xd3, 0x99, 0x1b, 0x47, 0xcc, 0x8b, 0xcc, 0xaa, 0x82, 0x1b,
	0xe1, 0xed, 0x30, 0x73, 0x8c, 0xa4, 0x1d, 0x9d, 0x43, 0x5d, 0x0a, 0xca, 0xf5, 0x98, 0x2f, 0x58,
	0x68, 0xd6, 0x54, 0xfa, 0xcf, 0x37, 0xa6, 0x2f, 0xa5, 0xd4, 0x51, 0x50, 0x0c, 0x34, 0x5f, 0x37,
	0xff, 0x2a, 0xc1, 0x4e, 0xe7, 0x9a, 0x7b, 0xef, 0x1e, 0x4d, 0x57, 0xfa, 0xa3, 0xeb, 0xaa, 0xfa,
	0x61, 0x5d, 0xd5, 0xd6, 0xe8, 0x6a, 0x1d, 0xc9, 0xa5, 0xff, 0x4e, 0xb2, 0x76, 0x87, 0xe4, 0xcf,
	0x40, 0x57, 0x8e, 0xfc, 0x6d, 0xeb, 0xb8, 0x26, 0x0d, 0xb2, 0xdc, 0x1b, 0x08, 0xad, 0x6c, 0x20,
	0xf4, 0x0c, 0xc0, 0x93, 0x4c, 0xb8, 0x62, 0x31, 0x63, 0xea, 0x31, 0x6d, 0x9f, 0xee, 0x5b, 0xb2,
	0xaf, 0x5b, 0xf7, 0x18, 0xb2, 0x86, 0x8b, 0x19, 0x7b, 0x5d, 0xc5, 0xf6, 0x9b, 0x51, 0xb7, 0x8d,
	0xb1, 0xae, 0xc2, 0xa4, 0xad, 0xd9, 0x02, 0x4d, 0xfe, 0xa2, 0x3a, 0x54, 0x47, 0xfd, 0x8b, 0xfe,
	0xe5, 0xdb, 0xbe, 0xf1, 0x89, 0xdc, 0xa4, 0x50, 0xa3, 0x80, 0xb6, 0x40, 0xef, 0x5c, 0xf6, 0xae,
	0xda, 0x7d, 0xe7, 0xb2, 0x6f, 0x14, 0x9b, 0x7f, 0x6b, 0x80, 0xda, 0x9e, 0xe0, 0x73, 0x2e, 0x16,
	0x8f, 0x26, 0x01, 0x78, 0x74, 0x09, 0xd4, 0x3e, 0x2c, 0x01, 0xfd, 0x7f, 0x96, 0xc0, 0x09, 0x68,
	0x8a, 0x98, 0xa4, 0x4f, 0x1c, 0x24, 0xc4, 0x3c, 0x2c, 0x9d, 0xe2, 0x06, 0x2b, 0xe8, 0xe6, 0xd6,
	0x70, 0x47, 0x35, 0x95, 0x7f, 0xa5, 0x9a, 0x0d, 0x6d, 0xa0, 0xf9, 0x7b, 0x21, 0xa5, 0x5c, 0x87,
	0xf2, 0x60, 0xd8, 0xc6, 0x43, 0xa3, 0x80, 0x2a, 0x50, 0xfc, 0xf1, 0xc4, 0x28, 0xa2, 0x2a, 0x94,
	0x7a, 0xce, 0xb9, 0x51, 0x52, 0x86, 0x97, 0x86, 0x86, 0x1a, 0x50, 0x93, 0xe4, 0x77, 0xed, 0xa1,
	0x6d, 0x94, 0x51, 0x0d, 0xb4, 0xc1, 0x85, 0x73, 0x65, 0x54, 0xe4, 0xaa, 0x37, 0x1a, 0xda, 0x46,
	0x15, 0x01, 0x54, 0x46, 0x7d, 0xb5, 0xae, 0xc9, 0x13, 0xaf, 0xda, 0xa3, 0x81, 0x6d, 0xe8, 0xd2,
	0x8c, 0xed, 0xc1, 0xa8, 0x67, 0x1b, 0x80, 0xb6, 0x01, 0xbe, 0x1f, 0x75, 0xbb, 0x83, 0x0e, 0xb6,
	0xed, 0xbe, 0x51, 0x97, 0xb0, 0x4e, 0xf7, 0x72, 0x60, 0x1b, 0x0d, 0x84, 0x60, 0x3b, 0x17, 0x97,
	0xfb, 0x93, 0x63, 0xbf, 0x35, 0xb6, 0xd0, 0x17, 0xb0, 0x7f, 0x6e, 0x5f, 0x61, 0xbb, 0xd3, 0x1e,
	0xda, 0xe7, 0xee, 0xd2, 0xdd, 0xe9, 0x3a, 0x9d, 0x0b, 0x63, 0xbb, 0xf9, 0x67, 0x01, 0x76, 0x3b,
	0x81, 0x3f, 0x67, 0xe1, 0xbd, 0xbf, 0xb7, 0x95, 0xea, 0x15, 0xee, 0x54, 0x6f, 0x17, 0xca, 0x73,
	0x72, 0x1d, 0x27, 0x7f, 0xfe, 0x3a, 0x4e, 0x36, 0xd2, 0x1a, 0x84, 0x94, 0x85, 0x8a, 0x60, 0x1d,
	0x27, 0x1b, 0x64, 0x40, 0x29, 0x08, 0xb3, 0x47, 0x2b, 0x97, 0xe8, 0x29, 0x54, 0x04, 0x99, 0x64,
	0x9c, 0x94, 0x71, 0x59, 0x90, 0x89, 0x43, 0xd1, 0x1e, 0xd4, 0xc6, 0xf1, 0xfb, 0xf7, 0xee, 0x3b,
	0xb6, 0x48, 0x19, 0xa9, 0xca, 0xfd, 0x05, 0x5b, 0xa0, 0x2f, 0xa1, 0xb1, 0x4a, 0x48, 0x4a, 0x45,
	0x7d, 0x85, 0x0a, 0xf4, 0xf9, 0xea, 0x7b, 0xa9, 0xa9, 0xf0, 0xa5, 0xa1, 0xf9, 0x6b, 0x11, 0x1a,
	0x6d, 0xba, 0xf2, 0x69, 0x27, 0x50, 0xcd, 0x46, 0x8f, 0xe4, 0x71, 0x3d, 0x4b, 0xe4, 0xf4, 0x60,
	0x40, 0xc2, 0x19, 0x0e, 0xbd, 0x06, 0xe0, 0xf9, 0x10, 0x90, 0x3e, 0xa0, 0xb4, 0x3b, 0xac, 0x1b,
	0x0e, 0xf0, 0x0a, 0x1a, 0xbd, 0x80, 0xb2, 0x6a, 0x11, 0xe9, 0x9c, 0xf3, 0x74, 0x6d, 0x53, 0xc1,
	0x09, 0x06, 0xbd, 0x82, 0x1a, 0x49, 0x55, 0x9d, 0x0e, 0x38, 0xe6, 0x26, 0xad, 0xe3, 0x1c, 0x29,
	0xd3, 0xf3, 0x72, 0x12, 0xd3, 0x21, 0x27, 0x6b, 0x5e, 0x6b, 0xc8, 0xc5, 0x2b, 0xe8, 0x33, 0x04,
	0x86, 0x17, 0x4c, 0xad, 0x74, 0x8c, 0x54, 0x41, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xac, 0x3d,
	0x7c, 0x5f, 0xeb, 0x0a, 0x00, 0x00,
}
