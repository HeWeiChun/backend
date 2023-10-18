package NGAP

import (
	"github.com/Freddy/sctp_flowmap/flowMap"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
)

type SupportedTAI struct {
	Tai        models.Tai
	SNssaiList []models.Snssai
}

// AMFConfigurationUpdateDecoder ProcedureCode = 0
func AMFConfigurationUpdateDecoder() {
}

// AMFStatusIndicationDecoder ProcedureCode = 1
func AMFStatusIndicationDecoder() {
}

// CellTrafficTraceDecoder ProcedureCode = 2
func CellTrafficTraceDecoder(CellTrafficTrace ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := CellTrafficTrace.CellTrafficTrace
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// DeactivateTraceDecoder ProcedureCode = 3
func DeactivateTraceDecoder(DeactivateTrace ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := DeactivateTrace.DeactivateTrace
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// DownlinkNASTransportDecoder ProcedureCode = 4
func DownlinkNASTransportDecoder(DownlinkNASTransport ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := DownlinkNASTransport.DownlinkNASTransport
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// DownlinkNonUEAssociatedNRPPaTransportDecoder ProcedureCode = 5
func DownlinkNonUEAssociatedNRPPaTransportDecoder() {
}

// DownlinkRANConfigurationTransferDecoder ProcedureCode = 6
func DownlinkRANConfigurationTransferDecoder() {
}

// DownlinkRANStatusTransferDecoder ProcedureCode = 7
func DownlinkRANStatusTransferDecoder(DownlinkRANStatusTransfer ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := DownlinkRANStatusTransfer.DownlinkRANStatusTransfer
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// DownlinkUEAssociatedNRPPaTransportDecoder ProcedureCode = 8
func DownlinkUEAssociatedNRPPaTransportDecoder(DownlinkUEAssociatedNRPPaTransport ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := DownlinkUEAssociatedNRPPaTransport.DownlinkUEAssociatedNRPPaTransport
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// ErrorIndicationDecoder ProcedureCode = 9
func ErrorIndicationDecoder(ErrorIndication ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := ErrorIndication.ErrorIndication
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// HandoverCancelDecoder ProcedureCode = 10
func HandoverCancelDecoder(HandoverCancel ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := HandoverCancel.HandoverCancel
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// HandoverNotifyDecoder ProcedureCode = 11
func HandoverNotifyDecoder(HandoverNotify ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := HandoverNotify.HandoverNotify
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// HandoverRequiredDecoder ProcedureCode = 12
func HandoverRequiredDecoder(HandoverRequired ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := HandoverRequired.HandoverRequired
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// HandoverRequestDecoder ProcedureCode = 13
func HandoverRequestDecoder(HandoverRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := HandoverRequest.HandoverRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				//RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				//packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// InitialContextSetupRequestDecoder ProcedureCode = 14
func InitialContextSetupRequestDecoder(InitialContextSetupRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := InitialContextSetupRequest.InitialContextSetupRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// InitialUEMessageDecoder ProcedureCode = 15
func InitialUEMessageDecoder(InitialUEMessage ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := InitialUEMessage.InitialUEMessage
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				//AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				//packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// LocationReportingControlDecoder ProcedureCode = 16
func LocationReportingControlDecoder(LocationReportingControl ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := LocationReportingControl.LocationReportingControl
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// LocationReportingFailureIndicationDecoder ProcedureCode = 17
func LocationReportingFailureIndicationDecoder(LocationReportingFailureIndication ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := LocationReportingFailureIndication.LocationReportingFailureIndication
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// LocationReportDecoder ProcedureCode = 18
func LocationReportDecoder(LocationReport ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := LocationReport.LocationReport
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// NASNonDeliveryIndicationDecoder ProcedureCode = 19
func NASNonDeliveryIndicationDecoder(NASNonDeliveryIndication ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := NASNonDeliveryIndication.NASNonDeliveryIndication
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// NGResetDecoder ProcedureCode = 20
func NGResetDecoder() {
}

// NGSetupRequestDecoder ProcedureCode = 21
func NGSetupRequestDecoder() {
}

// OverloadStartDecoder ProcedureCode = 22
func OverloadStartDecoder() {
}

// OverloadStopDecoder ProcedureCode = 23
func OverloadStopDecoder() {
}

// PagingDecoder ProcedureCode = 24
func PagingDecoder() {
}

// PathSwitchRequestDecoder ProcedureCode = 25
func PathSwitchRequestDecoder(PathSwitchRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := PathSwitchRequest.PathSwitchRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			// TODO: SourceAMFUENGAPID
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.SourceAMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// PDUSessionResourceModifyRequestDecoder ProcedureCode = 26
func PDUSessionResourceModifyRequestDecoder(PDUSessionResourceModifyRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceModifyRequest.PDUSessionResourceModifyRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// PDUSessionResourceModifyIndicationDecoder ProcedureCode = 27
func PDUSessionResourceModifyIndicationDecoder(PDUSessionResourceModifyIndication ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceModifyIndication.PDUSessionResourceModifyIndication
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// PDUSessionResourceReleaseCommandDecoder ProcedureCode = 28
func PDUSessionResourceReleaseCommandDecoder(PDUSessionResourceReleaseCommand ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceReleaseCommand.PDUSessionResourceReleaseCommand
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// PDUSessionResourceSetupRequestDecoder ProcedureCode = 29
func PDUSessionResourceSetupRequestDecoder(PDUSessionResourceSetupRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceSetupRequest.PDUSessionResourceSetupRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// PDUSessionResourceNotifyDecoder ProcedureCode = 30
func PDUSessionResourceNotifyDecoder(PDUSessionResourceNotify ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceNotify.PDUSessionResourceNotify
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// PrivateMessageDecoder ProcedureCode = 31
func PrivateMessageDecoder() {
}

// PWSCancelRequestDecoder ProcedureCode = 32
func PWSCancelRequestDecoder() {
}

// PWSFailureIndicationDecoder ProcedureCode = 33
func PWSFailureIndicationDecoder() {
}

// PWSRestartIndicationDecoder ProcedureCode = 34
func PWSRestartIndicationDecoder() {
}

// RANConfigurationUpdateDecoder ProcedureCode = 35
func RANConfigurationUpdateDecoder() {
}

// RerouteNASRequestDecoder ProcedureCode = 36
func RerouteNASRequestDecoder(RerouteNASRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := RerouteNASRequest.RerouteNASRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// RRCInactiveTransitionReportDecoder ProcedureCode = 37
func RRCInactiveTransitionReportDecoder(RRCInactiveTransitionReport ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := RRCInactiveTransitionReport.RRCInactiveTransitionReport
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// TraceFailureIndicationDecoder ProcedureCode = 38
func TraceFailureIndicationDecoder(TraceFailureIndication ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := TraceFailureIndication.TraceFailureIndication
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// TraceStartDecoder ProcedureCode = 39
func TraceStartDecoder(TraceStart ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := TraceStart.TraceStart
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// UEContextModificationRequestDecoder ProcedureCode = 40
func UEContextModificationRequestDecoder(UEContextModificationRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UEContextModificationRequest.UEContextModificationRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// UEContextReleaseCommandDecoder ProcedureCode = 41
func UEContextReleaseCommandDecoder(UEContextReleaseCommand ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UEContextReleaseCommand.UEContextReleaseCommand
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			// TODO: UENGAPIDs
			case ngapType.ProtocolIEIDUENGAPIDs:
				if ie.Value.UENGAPIDs.UENGAPIDPair != nil {
					AMF_UE_NGAP_ID := ie.Value.UENGAPIDs.UENGAPIDPair.AMFUENGAPID.Value
					packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
					RAN_UE_NGAP_ID := ie.Value.UENGAPIDs.UENGAPIDPair.RANUENGAPID.Value
					packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
				} else {
					AMF_UE_NGAP_ID := ie.Value.UENGAPIDs.AMFUENGAPID.Value
					packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
				}
			}
		}
	}
}

// UEContextReleaseRequestDecoder ProcedureCode = 42
func UEContextReleaseRequestDecoder(UEContextReleaseRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UEContextReleaseRequest.UEContextReleaseRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// UERadioCapabilityCheckRequestDecoder ProcedureCode = 43
func UERadioCapabilityCheckRequestDecoder(UERadioCapabilityCheckRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UERadioCapabilityCheckRequest.UERadioCapabilityCheckRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// UERadioCapabilityInfoIndicationDecoder ProcedureCode = 44
func UERadioCapabilityInfoIndicationDecoder(UERadioCapabilityInfoIndication ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UERadioCapabilityInfoIndication.UERadioCapabilityInfoIndication
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// UETNLABindingReleaseRequestDecoder ProcedureCode = 45
func UETNLABindingReleaseRequestDecoder(UETNLABindingReleaseRequest ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UETNLABindingReleaseRequest.UETNLABindingReleaseRequest
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// UplinkNASTransportDecoder ProcedureCode = 46
func UplinkNASTransportDecoder(UplinkNASTransport ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UplinkNASTransport.UplinkNASTransport
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// UplinkNonUEAssociatedNRPPaTransportDecoder ProcedureCode = 47
func UplinkNonUEAssociatedNRPPaTransportDecoder() {
}

// UplinkRANConfigurationTransferDecoder ProcedureCode = 48
func UplinkRANConfigurationTransferDecoder() {
}

// UplinkRANStatusTransferDecoder ProcedureCode = 49
func UplinkRANStatusTransferDecoder(UplinkRANStatusTransfer ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UplinkRANStatusTransfer.UplinkRANStatusTransfer
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// UplinkUEAssociatedNRPPaTransportDecoder ProcedureCode = 50
func UplinkUEAssociatedNRPPaTransportDecoder(UplinkUEAssociatedNRPPaTransport ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := UplinkUEAssociatedNRPPaTransport.UplinkUEAssociatedNRPPaTransport
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}

// WriteReplaceWarningRequestDecoder ProcedureCode = 51
func WriteReplaceWarningRequestDecoder() {
}

// SecondaryRATDataUsageReportDecoder ProcedureCode = 52
func SecondaryRATDataUsageReportDecoder(SecondaryRATDataUsageReport ngapType.InitiatingMessageValue, packet *flowMap.Packet) {
	ies := SecondaryRATDataUsageReport.SecondaryRATDataUsageReport
	if ies != nil {
		for i := 0; i < len(ies.ProtocolIEs.List); i++ {
			ie := ies.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				packet.AMF_UE_NGAP_ID = AMF_UE_NGAP_ID
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				packet.RAN_UE_NGAP_ID = RAN_UE_NGAP_ID
			}
		}
	}
}
