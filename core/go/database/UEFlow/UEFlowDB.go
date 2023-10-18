package UEFlow

import (
	"container/list"
	"github.com/Freddy/sctp_flowmap/database"
	"log"
)

func init() {
	creatUeFlowTable()
}

func InsertUeFlow(UeFlowList *list.List) {
	var sqlStr string = insertUeFlowSQL
	tx, err := database.Connect.Begin()
	checkErr(err)
	stmt, err := tx.Prepare(sqlStr)
	checkErr(err)
	for info := UeFlowList.Front(); info != nil; info = info.Next() {
		fl := info.Value.(*UeFlow)
		if _, err := stmt.Exec(
			fl.FlowId,
			fl.RanUeNgapId,
			fl.TotalNum,
			fl.BeginTime,
			fl.LatestTime,
			fl.VerificationTag,
			fl.SrcIP,
			fl.DstIP,
			fl.StatusFlow,
			fl.TaskID,
		); err != nil {
			log.Fatal(err)
		}
	}
	checkErr(tx.Commit())

}

func creatUeFlowTable() {
	var sqlStr string = createUEFlowTableSQL

	_, err := database.Connect.Exec(sqlStr)
	checkErr(err)
}

func dropUeFlowTable() {
	if _, err := database.Connect.Exec(dropUEFlowTableSQL); err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
