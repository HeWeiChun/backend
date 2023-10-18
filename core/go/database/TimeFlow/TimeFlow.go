package TimeFlow

import (
	"fmt"
	"time"
)

type TimeFlow struct {
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

var TimeFlowTable = "sctp.time_flow"

var insertTimeFlowSQL = `
	INSERT INTO ` + TimeFlowTable +
	`
	(flow_id, ran_ue_ngap_id, total_num, 
	begin_time, latest_time, verification_tag, src_ip, 
	dst_ip, status_flow, task_id) 
	values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

var queryTimeFlowLogSQL = `
	SELECT flow_id, ran_ue_ngap_id, total_num, 
	begin_time, latest_time, verification_tag, src_ip, 
	dst_ip, status_flow, task_id
	FROM ` + TimeFlowTable

var createTimeFlowTableSQL = `
CREATE TABLE IF NOT EXISTS ` + TimeFlowTable + ` (
	flow_id String,
	ran_ue_ngap_id UInt64,
	total_num UInt32,
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
	INDEX i_task_id (task_id) TYPE minmax GRANULARITY 4
	)  
	ENGINE = MergeTree() 
	PARTITION BY toYYYYMMDD(begin_time)
	ORDER BY (begin_time)
`

var dropTimeFlowTableSQL = "DROP TABLE " + TimeFlowTable

func (fl *TimeFlow) initFlowLog() {
	//TODO init
}

func (fl TimeFlow) String() string {
	return fmt.Sprintf(`
		flow_id: %s,
		ran_ue_ngap_id: %u,
		total_num: %u ,
		begin_time: %s , 
		latest_time: %s , 
		verification_tag: %u ,
		src_ip: %s , 
		dst_ip: %s , 
		status_flow: %u ,
		task_id: %s,
		`, fl.FlowId, fl.RanUeNgapId, fl.TotalNum, fl.BeginTime,
		fl.LatestTime, fl.VerificationTag, fl.SrcIP, fl.DstIP, fl.StatusFlow, fl.TaskID)
}
