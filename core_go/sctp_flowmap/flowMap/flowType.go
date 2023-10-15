package flowMap

import (
	"container/list"
	"strconv"
	"time"
)

type PacketType string

var (
	TABLE_SIZE        uint64 = 10240000
	PACKET_TABLE_SIZE uint   = 20000
)

const (
	NGAPType PacketType = "NGAP"
)

type Flow struct {
	info *FlowInfo
	next *Flow
}

type FlowInfo struct {
	RAN_UE_NGAP_ID int64
	FlowID         string
	FlowType       PacketType
	TotalNum       uint
	PacketList     list.List

	BeginTime       time.Time
	EndTime         time.Time
	EndTimeUs       int64
	VerificationTag uint32
	SrcIP           string
	DstIP           string
	TimeID          string
	TaskID          string
}
type Packet struct {
	FlowID              string
	NgapType            string
	NgapProcedureCode   string
	NgapRoute           string
	RAN_UE_NGAP_ID      int64
	AMF_UE_NGAP_ID      int64
	ArriveTimeUs        int64
	ArriveTime          time.Time
	PacketLen           uint32
	TimeInterval        uint64
	NgapPayloadBytes    []byte
	PayloadBytes        []byte
	VerificationTag     uint32
	SrcIP               string
	DstIP               string
	TimeID              string
	DirSeq              int8
	InitiatingMessage   uint8
	SuccessfulOutcome   uint8
	UnsuccessfulOutcome uint8
}

func loadFlow(flowId string, flowTable []*Flow) (*FlowInfo, bool) {
	numFlowId, _ := strconv.ParseUint(flowId, 10, 64)
	cur := flowTable[numFlowId%TABLE_SIZE]
	if cur == nil {
		return nil, false
	}
	for ; cur != nil; cur = cur.next {
		if cur.info.FlowID == flowId {
			if cur.info.TotalNum == PACKET_TABLE_SIZE {
				return cur.info, false
			}
			return cur.info, true
		}
	}
	return nil, false
}

func storeFlow(flowId string, flowInfo *FlowInfo, flowTable []*Flow) {
	numFlowId, _ := strconv.ParseUint(flowId, 10, 64)
	cur := flowTable[numFlowId%TABLE_SIZE]
	if cur == nil {
		flowTable[numFlowId%TABLE_SIZE] = &Flow{info: flowInfo, next: nil}
	} else {
		next := cur
		for next.next != nil {
			next = next.next
		}
		next.next = &Flow{info: flowInfo, next: nil}
	}
}

func deleteFlow(flowId string, flowTable []*Flow) bool {
	numFlowId, _ := strconv.ParseUint(flowId, 10, 64)
	pre := flowTable[numFlowId%TABLE_SIZE]
	if pre.info.FlowID == flowId {
		flowTable[numFlowId%TABLE_SIZE] = pre.next
		return true
	}
	cur := pre.next
	for cur != nil && cur.info.FlowID != flowId {
		pre = pre.next
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = cur.next
	return true
}
