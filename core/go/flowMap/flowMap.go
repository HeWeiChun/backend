package flowMap

import (
	"container/list"
	"github.com/Freddy/sctp_flowmap/database/PacketDB"
	"github.com/Freddy/sctp_flowmap/database/TimeFlow"
	"github.com/Freddy/sctp_flowmap/database/UEFlow"
	"strconv"
)

var (
	FlowTable_UE   []*Flow
	FlowTable_Time []*Flow
	flowCount      uint64 = 0
	TimeMax        int64  = 300000000
	min_intervl    int64  = 1
	TimeChip       int64  = 5000000000
)

func init() {
	FlowTable_UE = make([]*Flow, TABLE_SIZE)
	FlowTable_Time = make([]*Flow, TABLE_SIZE)
}

func Count_UE_ID(packet *Packet, taskid string) (uint64, bool) {
	RAN_UE_NGAP_ID := packet.RAN_UE_NGAP_ID
	if RAN_UE_NGAP_ID == -1 {
		return 0, false
	}
	flowID := FastTwoHash([]byte(strconv.FormatInt(RAN_UE_NGAP_ID, 10)), []byte(taskid))
	return flowID, true

}

func Count_Time_ID(packet *Packet, TimeFirst int64, taskid string) (uint64, bool) {
	RAN_UE_NGAP_ID := packet.RAN_UE_NGAP_ID
	if RAN_UE_NGAP_ID == -1 {
		return 0, false
	}
	Time := packet.ArriveTimeUs - TimeFirst
	Time = Time / TimeChip
	return FastTwoHash([]byte(strconv.FormatInt(Time, 10)), []byte(taskid)), true
}

func Put(packet *Packet, flowTable []*Flow, flowID string, taskid string) bool {
	var flowInfo *FlowInfo
	var first = false // 是否流的首包
	flowInfo, isExist := loadFlow(flowID, flowTable)
	if isExist {
		packet.TimeInterval = uint64(packet.ArriveTimeUs - flowInfo.EndTimeUs + min_intervl)
		flowInfo.EndTime = packet.ArriveTime
		flowInfo.EndTimeUs = packet.ArriveTimeUs
		flowInfo.PacketList.PushBack(packet)
		flowInfo.TotalNum = flowInfo.TotalNum + 1
		if flowInfo.SrcIP == packet.SrcIP {
			packet.DirSeq = 1
		} else {
			packet.DirSeq = -1
		}
	} else {
		// 首次接收，创建流info
		flowInfo = &FlowInfo{
			FlowID:          flowID,
			RAN_UE_NGAP_ID:  packet.RAN_UE_NGAP_ID,
			FlowType:        NGAPType,
			TotalNum:        1,
			VerificationTag: packet.VerificationTag,
			SrcIP:           packet.SrcIP,
			DstIP:           packet.DstIP,
			TimeID:          packet.TimeID,
			BeginTime:       packet.ArriveTime,
			TaskID:          taskid,
		}
		packet.DirSeq = 1
		packet.TimeInterval = uint64(min_intervl)
		flowInfo.EndTime = packet.ArriveTime
		flowInfo.EndTimeUs = packet.ArriveTimeUs
		flowInfo.PacketList = list.List{}
		flowInfo.PacketList.PushBack(packet)
		storeFlow(flowID, flowInfo, flowTable)
		flowCount++
		first = true
	}
	return first
}

func UEFlowMapToStore() *list.List {
	var rubbishList = list.New()

	for _, flow := range FlowTable_UE {
		if flow == nil {
			continue
		}
		for cur := flow; cur != nil; cur = cur.next {
			flowInfo := cur.info
			rubbishList.PushBack(flowInfo)
			deleteFlow(flowInfo.FlowID, FlowTable_UE)
			flowCount--

		}
	}
	//fmt.Println("UEFlow")
	UEflowStore(rubbishList)
	return rubbishList
}

func UEflowStore(rubbishList *list.List) {
	var UeFlowList = list.New()
	for info := rubbishList.Front(); info != nil; info = info.Next() {
		flowInfo := info.Value.(*FlowInfo)
		fl := &UEFlow.UeFlow{
			FlowId:          flowInfo.FlowID,
			RanUeNgapId:     uint64(flowInfo.RAN_UE_NGAP_ID),
			TotalNum:        uint32(flowInfo.TotalNum),
			BeginTime:       flowInfo.BeginTime,
			LatestTime:      flowInfo.EndTime,
			VerificationTag: uint64(flowInfo.VerificationTag),
			SrcIP:           flowInfo.SrcIP,
			DstIP:           flowInfo.DstIP,
			StatusFlow:      0,
			TaskID:          flowInfo.TaskID,
		}
		UeFlowList.PushBack(fl)

	}
	//fmt.Println("UEList")
	UEFlow.InsertUeFlow(UeFlowList)

}

func TimeFlowMapToStore(taskid string) {
	var rubbishList = list.New()
	for _, flow := range FlowTable_Time {
		if flow == nil {
			continue
		}
		for cur := flow; cur != nil; cur = cur.next {
			flowInfo := cur.info
			rubbishList.PushBack(flowInfo)
			deleteFlow(flowInfo.FlowID, FlowTable_Time)
			flowCount--

		}
	}
	//fmt.Println("TimeFlow")
	TimeflowStore(rubbishList, taskid)
}

func TimeflowStore(rubbishList *list.List, taskid string) {
	var TimeFlowList = list.New()
	var PacketList = list.New()
	for info := rubbishList.Front(); info != nil; info = info.Next() {
		flowInfo := info.Value.(*FlowInfo)
		fl := &TimeFlow.TimeFlow{
			FlowId:          flowInfo.FlowID,
			RanUeNgapId:     uint64(flowInfo.RAN_UE_NGAP_ID),
			TotalNum:        uint32(flowInfo.TotalNum),
			BeginTime:       flowInfo.BeginTime,
			LatestTime:      flowInfo.EndTime,
			VerificationTag: uint64(flowInfo.VerificationTag),
			SrcIP:           flowInfo.SrcIP,
			DstIP:           flowInfo.DstIP,
			StatusFlow:      0,
			TaskID:          flowInfo.TaskID,
		}
		TimeFlowList.PushBack(fl)
		for cur := flowInfo.PacketList.Front(); cur != nil; cur = cur.Next() {
			parse := cur.Value.(*Packet)
			flowue, id := Count_UE_ID(parse, taskid)
			if !id {
				flowue = '0'
			}
			packet := &PacketDB.Packet{
				//PacketId: FnvHash([]byte(string(parse.ArriveTimeUs))),
				NgapType:            parse.NgapType,
				NgapProcedureCode:   parse.NgapProcedureCode,
				RanUeNgapId:         parse.RAN_UE_NGAP_ID,
				PacketLen:           parse.PacketLen,
				ArriveTimeUs:        parse.ArriveTimeUs,
				ArriveTime:          parse.ArriveTime,
				TimeInterval:        parse.TimeInterval,
				VerificationTag:     uint64(parse.VerificationTag),
				SrcIP:               parse.SrcIP,
				DstIP:               parse.DstIP,
				DirSeq:              parse.DirSeq,
				FlowUEID:            strconv.FormatUint(flowue, 10),
				FlowTimeID:          parse.FlowID,
				InitiatingMessage:   parse.InitiatingMessage,
				SuccessfulOutcome:   parse.SuccessfulOutcome,
				UnsuccessfulOutcome: parse.UnsuccessfulOutcome,
				StatusPacket:        0,
			}
			PacketList.PushBack(packet)
		}
	}
	//fmt.Println("TimeList")
	TimeFlow.InsertTimeFlow(TimeFlowList)
	PacketDB.InsertPacket(PacketList)

}
