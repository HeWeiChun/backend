package NGAP

import (
	"github.com/Freddy/sctp_flowmap/flowMap"
	"github.com/free5gc/ngap/ngapType"
)

// AMFConfigurationUpdateAcknowledgeDecoder ProcedureCode = 0
func AMFConfigurationUpdateAcknowledgeDecoder() {
}

// HandoverCancelAcknowledgeDecoder ProcedureCode = 10
func HandoverCancelAcknowledgeDecoder(HandoverCancelAcknowledge ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := HandoverCancelAcknowledge.HandoverCancelAcknowledge
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

// HandoverCommandDecoder ProcedureCode = 12
func HandoverCommandDecoder(HandoverCommand ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := HandoverCommand.HandoverCommand
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

// HandoverRequestAcknowledgeDecoder ProcedureCode = 13
func HandoverRequestAcknowledgeDecoder(HandoverRequestAcknowledge ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := HandoverRequestAcknowledge.HandoverRequestAcknowledge
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

// InitialContextSetupResponseDecoder ProcedureCode = 15
func InitialContextSetupResponseDecoder(InitialContextSetupResponse ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := InitialContextSetupResponse.InitialContextSetupResponse
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

// NGResetAcknowledgeDecoder ProcedureCode = 20
func NGResetAcknowledgeDecoder() {

}

// NGSetupResponseDecoder ProcedureCode = 21
func NGSetupResponseDecoder() {

}

// PathSwitchRequestAcknowledgeDecoder ProcedureCode = 25
func PathSwitchRequestAcknowledgeDecoder(PathSwitchRequestAcknowledge ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := PathSwitchRequestAcknowledge.PathSwitchRequestAcknowledge
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

// PDUSessionResourceModifyResponseDecoder ProcedureCode = 26
func PDUSessionResourceModifyResponseDecoder(PDUSessionResourceModifyResponse ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceModifyResponse.PDUSessionResourceModifyResponse
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

// PDUSessionResourceModifyConfirmDecoder ProcedureCode = 27
func PDUSessionResourceModifyConfirmDecoder(PDUSessionResourceModifyConfirm ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceModifyConfirm.PDUSessionResourceModifyConfirm
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

// PDUSessionResourceReleaseResponseDecoder ProcedureCode = 28
func PDUSessionResourceReleaseResponseDecoder(PDUSessionResourceReleaseResponse ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceReleaseResponse.PDUSessionResourceReleaseResponse
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

// PDUSessionResourceSetupResponseDecoder ProcedureCode = 29
func PDUSessionResourceSetupResponseDecoder(PDUSessionResourceSetupResponse ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := PDUSessionResourceSetupResponse.PDUSessionResourceSetupResponse
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

// PWSCancelResponseDecoder ProcedureCode = 30
func PWSCancelResponseDecoder() {
}

// RANConfigurationUpdateAcknowledgeDecoder ProcedureCode = 35
func RANConfigurationUpdateAcknowledgeDecoder() {
}

// UEContextModificationResponseDecoder ProcedureCode = 40
func UEContextModificationResponseDecoder(UEContextModificationResponse ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := UEContextModificationResponse.UEContextModificationResponse
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

// UEContextReleaseCompleteDecoder ProcedureCode = 41
func UEContextReleaseCompleteDecoder(UEContextReleaseComplete ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := UEContextReleaseComplete.UEContextReleaseComplete
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

// UERadioCapabilityCheckResponseDecoder ProcedureCode = 43
func UERadioCapabilityCheckResponseDecoder(UERadioCapabilityCheckResponse ngapType.SuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := UERadioCapabilityCheckResponse.UERadioCapabilityCheckResponse
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

// WriteReplaceWarningResponseDecoder ProcedureCode = 45
func WriteReplaceWarningResponseDecoder() {
}
