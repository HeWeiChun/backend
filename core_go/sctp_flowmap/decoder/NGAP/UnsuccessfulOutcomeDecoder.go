package NGAP

import (
	"github.com/Freddy/sctp_flowmap/flowMap"
	"github.com/free5gc/ngap/ngapType"
)

// AMFConfigurationUpdateFailureDecoder ProcedureCode = 0
func AMFConfigurationUpdateFailureDecoder() {
}

// HandoverPreparationFailureDecoder ProcedureCode = 12
func HandoverPreparationFailureDecoder(HandoverPreparationFailure ngapType.UnsuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := HandoverPreparationFailure.HandoverPreparationFailure
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

// HandoverFailureDecoder ProcedureCode = 13
func HandoverFailureDecoder(HandoverFailure ngapType.UnsuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := HandoverFailure.HandoverFailure
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

// InitialContextSetupFailureDecoder ProcedureCode = 14
func InitialContextSetupFailureDecoder(InitialContextSetupFailure ngapType.UnsuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := InitialContextSetupFailure.InitialContextSetupFailure
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

// NGSetupFailureDecoder ProcedureCode = 21
func NGSetupFailureDecoder() {
}

// PathSwitchRequestFailureDecoder ProcedureCode = 25
func PathSwitchRequestFailureDecoder(PathSwitchRequestFailure ngapType.UnsuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := PathSwitchRequestFailure.PathSwitchRequestFailure
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

// RANConfigurationUpdateFailureDecoder ProcedureCode = 35
func RANConfigurationUpdateFailureDecoder() {
}

// UEContextModificationFailureDecoder ProcedureCode = 40
func UEContextModificationFailureDecoder(UEContextModificationFailure ngapType.UnsuccessfulOutcomeValue, packet *flowMap.Packet) {
	ies := UEContextModificationFailure.UEContextModificationFailure
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
