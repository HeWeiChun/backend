package UEFlow

import (
	"fmt"
	"time"
)

type UeFlow struct {
	FlowId          string //流哈希id
	RanUeNgapId     uint64 //包哈希id
	TotalNum        uint32
	BeginTime       time.Time
	LatestTime      time.Time
	VerificationTag uint64
	SrcIP           string
	DstIP           string
	//TimeID          uint64
	StatusFlow uint16
	TaskID     string
}

var UeFlowTable = "sctp.ue_flow"
var insertUeFlowSQL = `
	INSERT INTO ` + UeFlowTable +
	`
	(flow_id, ran_ue_ngap_id, total_num, 
	begin_time, latest_time, verification_tag, src_ip, 
	dst_ip, status_flow, task_id) 
	values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

var queryUEFlowLogSQL = `
	SELECT flow_id, ran_ue_ngap_id, total_num, 
	begin_time, latest_time, verification_tag, src_ip, 
	dst_ip, status_flow, task_id 
	FROM ` + UeFlowTable

var createUEFlowTableSQL = `
CREATE TABLE IF NOT EXISTS ` + UeFlowTable + ` (
	flow_id String,
	ran_ue_ngap_id UInt64,
	total_num UInt32,
	start_second UInt64,
	end_second UInt64,
	begin_time DateTime64(6), 
	latest_time DateTime64(6), 
	verification_tag UInt64,
	src_ip String, 
	dst_ip String,
	status_flow UInt16, 
	task_id String,
	
	INDEX i_flow_id (flow_id) TYPE minmax GRANULARITY 4, 
	INDEX i_ran_ue_ngap_id (ran_ue_ngap_id) TYPE minmax GRANULARITY 4, 
	INDEX i_total_num (total_num) TYPE minmax GRANULARITY 4, 
	INDEX i_begin_time (begin_time) TYPE minmax GRANULARITY 4, 
	INDEX i_src_ip (src_ip) TYPE minmax GRANULARITY 4, 
	INDEX i_dst_ip (dst_ip) TYPE minmax GRANULARITY 4,
	INDEX i_task_id (task_id) TYPE minmax GRANULARITY 4,
	INDEX i_status_flow (status_flow) TYPE minmax GRANULARITY 4
	)  
	ENGINE = MergeTree() 
	PARTITION BY toYYYYMMDD(begin_time)
	ORDER BY (begin_time)
`

var dropUEFlowTableSQL = "DROP TABLE " + UeFlowTable

func (fl *UeFlow) initFlowLog() {
	//TODO init
}

func (fl UeFlow) String() string {
	return fmt.Sprintf(`
            FlowId: %s,
		    RanUeNgapId: %u,
            TotalNum: %u ,
	    	BeginTime: %s , 
		    LatestTime: %s , 
		    VerificationTag: %u ,
		    SrcIP: %s , 
		    DstIP: %s ,
            StatusFlow: %u ,
            TaskID: %s 
		`, fl.FlowId, fl.RanUeNgapId, fl.TotalNum, fl.BeginTime,
		fl.LatestTime, fl.VerificationTag, fl.SrcIP, fl.DstIP, fl.StatusFlow, fl.TaskID)
}
