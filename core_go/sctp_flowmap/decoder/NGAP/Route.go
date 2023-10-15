package NGAP

import (
	"github.com/Freddy/sctp_flowmap/flowMap"
	"github.com/free5gc/ngap/ngapType"
)

// 对NGAP包进行初步的分类

func RouteNGAP(DecResult *ngapType.NGAPPDU, packet *flowMap.Packet) {
	//fmt.Println("===============RouteNGAP begins working.===============")
	switch DecResult.Present {
	case ngapType.NGAPPDUPresentInitiatingMessage:
		if DecResult.InitiatingMessage == nil {
			//fmt.Println("===============There is no initiatingMessage in this package.===============")
		}
		packet.NgapRoute = "InitiatingMessage"
		packet.InitiatingMessage = 1
		RouteInitiatingMessage(DecResult, packet)

	case ngapType.NGAPPDUPresentSuccessfulOutcome:
		if DecResult.SuccessfulOutcome == nil {
			//fmt.Println("===============There is no SuccessfulOutcome in this package.===============")
		}
		packet.NgapRoute = "SuccessfulOutcome"
		packet.SuccessfulOutcome = 1
		RouteSuccessfulOutcome(DecResult, packet)

	case ngapType.NGAPPDUPresentUnsuccessfulOutcome:
		if DecResult.UnsuccessfulOutcome == nil {
			//fmt.Println("===============There is no UnsuccessfulOutcome in this package.============")
		}
		packet.NgapRoute = "UnsuccessfulOutcome"
		packet.UnsuccessfulOutcome = 1
		RouteUnsuccessfulOutcome(DecResult, packet)
	}
}

func RouteInitiatingMessage(DecResult *ngapType.NGAPPDU, packet *flowMap.Packet) {
	InitiatingMessage := DecResult.InitiatingMessage
	switch InitiatingMessage.ProcedureCode.Value {
	// Dirseq: 1表示 NG-RAN -> AMF, -1表示 AMF -> NG-RAN
	// ProcedureCode = 0
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		packet.NgapProcedureCode = "AMFConfigurationUpdate"
		packet.NgapType = "AMFConfigurationUpdate"
		packet.DirSeq = -1
		AMFConfigurationUpdateDecoder()
	// ProcedureCode = 1
	case ngapType.ProcedureCodeAMFStatusIndication:
		packet.NgapProcedureCode = "AMFStatusIndication"
		packet.NgapType = "AMFStatusIndication"
		packet.DirSeq = -1
		AMFStatusIndicationDecoder()
	// ProcedureCode = 2
	case ngapType.ProcedureCodeCellTrafficTrace:
		packet.NgapProcedureCode = "CellTrafficTrace"
		packet.NgapType = "CellTrafficTrace"
		packet.DirSeq = 1
		CellTrafficTraceDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 3
	case ngapType.ProcedureCodeDeactivateTrace:
		packet.NgapProcedureCode = "DeactivateTrace"
		packet.NgapType = "DeactivateTrace"
		packet.DirSeq = -1
		DeactivateTraceDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 4
	case ngapType.ProcedureCodeDownlinkNASTransport:
		packet.NgapProcedureCode = "DownlinkNASTransport"
		packet.NgapType = "DownlinkNASTransport"
		packet.DirSeq = -1
		DownlinkNASTransportDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 5
	case ngapType.ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport:
		packet.NgapProcedureCode = "DownlinkNonUEAssociatedNRPPaTransport"
		packet.NgapType = "DownlinkNonUEAssociatedNRPPaTransport"
		packet.DirSeq = -1
		DownlinkNonUEAssociatedNRPPaTransportDecoder()
	// ProcedureCode = 6
	case ngapType.ProcedureCodeDownlinkRANConfigurationTransfer:
		packet.NgapProcedureCode = "DownlinkRANConfigurationTransfer"
		packet.NgapType = "DownlinkRANConfigurationTransfer"
		packet.DirSeq = -1
		DownlinkRANConfigurationTransferDecoder()
	// ProcedureCode = 7
	case ngapType.ProcedureCodeDownlinkRANStatusTransfer:
		packet.NgapProcedureCode = "DownlinkRANStatusTransfer"
		packet.NgapType = "DownlinkRANStatusTransfer"
		packet.DirSeq = -1
		DownlinkRANStatusTransferDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 8
	case ngapType.ProcedureCodeDownlinkUEAssociatedNRPPaTransport:
		packet.NgapProcedureCode = "DownlinkUEAssociatedNRPPaTransport"
		packet.NgapType = "DownlinkUEAssociatedNRPPaTransport"
		packet.DirSeq = -1
		DownlinkUEAssociatedNRPPaTransportDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 9
	case ngapType.ProcedureCodeErrorIndication:
		packet.NgapProcedureCode = "ErrorIndication"
		packet.NgapType = "ErrorIndication"
		packet.DirSeq = 0 // TODO: -1 or 1?
		ErrorIndicationDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 10
	case ngapType.ProcedureCodeHandoverCancel:
		packet.NgapProcedureCode = "HandoverCancellation"
		packet.NgapType = "HandoverCancel"
		packet.DirSeq = 1
		HandoverCancelDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 11
	case ngapType.ProcedureCodeHandoverNotification:
		packet.NgapProcedureCode = "HandoverNotification"
		packet.NgapType = "HandoverNotify"
		packet.DirSeq = 1
		HandoverNotifyDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 12
	case ngapType.ProcedureCodeHandoverPreparation:
		packet.NgapProcedureCode = "HandoverPreparation"
		packet.NgapType = "HandoverRequired"
		packet.DirSeq = 1
		HandoverRequiredDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 13
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		packet.NgapProcedureCode = "HandoverResourceAllocation"
		packet.NgapType = "HandoverRequest"
		packet.DirSeq = -1
		HandoverRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 14
	case ngapType.ProcedureCodeInitialContextSetup:
		packet.NgapProcedureCode = "InitialContextSetup"
		packet.NgapType = "InitialContextSetupRequest"
		packet.DirSeq = -1
		InitialContextSetupRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 15
	case ngapType.ProcedureCodeInitialUEMessage:
		packet.NgapProcedureCode = "InitialUEMessage"
		packet.NgapType = "InitialUEMessage"
		packet.DirSeq = 1
		InitialUEMessageDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 16
	case ngapType.ProcedureCodeLocationReportingControl:
		packet.NgapProcedureCode = "LocationReportingControl"
		packet.NgapType = "LocationReportingControl"
		packet.DirSeq = -1
		LocationReportingControlDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 17
	case ngapType.ProcedureCodeLocationReportingFailureIndication:
		packet.NgapProcedureCode = "LocationReportingFailureIndication"
		packet.NgapType = "LocationReportingFailureIndication"
		packet.DirSeq = 1
		LocationReportingFailureIndicationDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 18
	case ngapType.ProcedureCodeLocationReport:
		packet.NgapProcedureCode = "LocationReport"
		packet.NgapType = "LocationReport"
		packet.DirSeq = 1
		LocationReportDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 19
	case ngapType.ProcedureCodeNASNonDeliveryIndication:
		packet.NgapProcedureCode = "NASNonDeliveryIndication"
		packet.NgapType = "NASNonDeliveryIndication"
		packet.DirSeq = 1
		NASNonDeliveryIndicationDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 20
	case ngapType.ProcedureCodeNGReset:
		packet.NgapProcedureCode = "NGReset"
		packet.NgapType = "NGReset"
		packet.DirSeq = 0 //TODO: -1 or 1?
		NGResetDecoder()
	// ProcedureCode = 21
	case ngapType.ProcedureCodeNGSetup:
		packet.NgapProcedureCode = "NGSetup"
		packet.NgapType = "NGSetupRequest"
		packet.DirSeq = 1
		NGSetupRequestDecoder()
	// ProcedureCode = 22
	case ngapType.ProcedureCodeOverloadStart:
		packet.NgapProcedureCode = "OverloadStart"
		packet.NgapType = "OverloadStart"
		packet.DirSeq = -1
		OverloadStartDecoder()
	// ProcedureCode = 23
	case ngapType.ProcedureCodeOverloadStop:
		packet.NgapProcedureCode = "OverloadStop"
		packet.NgapType = "OverloadStop"
		packet.DirSeq = -1
		OverloadStopDecoder()
	// ProcedureCode = 24
	case ngapType.ProcedureCodePaging:
		packet.NgapProcedureCode = "Paging"
		packet.NgapType = "Paging"
		packet.DirSeq = -1
		PagingDecoder()
	// ProcedureCode = 25
	case ngapType.ProcedureCodePathSwitchRequest:
		packet.NgapProcedureCode = "PathSwitchRequest"
		packet.NgapType = "PathSwitchRequest"
		packet.DirSeq = 1
		PathSwitchRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 26
	case ngapType.ProcedureCodePDUSessionResourceModify:
		packet.NgapProcedureCode = "PDUSessionResourceModify"
		packet.NgapType = "PDUSessionResourceModifyRequest"
		packet.DirSeq = -1
		PDUSessionResourceModifyRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 27
	case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
		packet.NgapProcedureCode = "PDUSessionResourceModifyIndication"
		packet.NgapType = "PDUSessionResourceModifyIndication"
		packet.DirSeq = 1
		PDUSessionResourceModifyIndicationDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 28
	case ngapType.ProcedureCodePDUSessionResourceRelease:
		packet.NgapProcedureCode = "PDUSessionResourceRelease"
		packet.NgapType = "PDUSessionResourceReleaseCommand"
		packet.DirSeq = -1
		PDUSessionResourceReleaseCommandDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 29
	case ngapType.ProcedureCodePDUSessionResourceSetup:
		packet.NgapProcedureCode = "PDUSessionResourceSetup"
		packet.NgapType = "PDUSessionResourceSetupRequest"
		packet.DirSeq = -1
		PDUSessionResourceSetupRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 30
	case ngapType.ProcedureCodePDUSessionResourceNotify:
		packet.NgapProcedureCode = "PDUSessionResourceNotify"
		packet.NgapType = "PDUSessionResourceNotify"
		packet.DirSeq = 1
		PDUSessionResourceNotifyDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 31
	case ngapType.ProcedureCodePrivateMessage:
		packet.NgapProcedureCode = "PrivateMessage"
		packet.NgapType = "PrivateMessage"
		packet.DirSeq = 0 //TODO: -1 or 1?
		PrivateMessageDecoder()
	// ProcedureCode = 32
	case ngapType.ProcedureCodePWSCancel:
		packet.NgapProcedureCode = "PWSCancel"
		packet.NgapType = "PWSCancelRequest"
		packet.DirSeq = -1
		PWSCancelRequestDecoder()
	// ProcedureCode = 33
	case ngapType.ProcedureCodePWSFailureIndication:
		packet.NgapProcedureCode = "PWSFailureIndication"
		packet.NgapType = "PWSFailureIndication"
		packet.DirSeq = 1
		PWSFailureIndicationDecoder()
	// ProcedureCode = 34
	case ngapType.ProcedureCodePWSRestartIndication:
		packet.NgapProcedureCode = "PWSRestartIndication"
		packet.NgapType = "PWSRestartIndication"
		packet.DirSeq = 1
		PWSRestartIndicationDecoder()
	// ProcedureCode = 35
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		packet.NgapProcedureCode = "RANConfigurationUpdate"
		packet.NgapType = "RANConfigurationUpdate"
		packet.DirSeq = 1
		RANConfigurationUpdateDecoder()
	// ProcedureCode = 36
	case ngapType.ProcedureCodeRerouteNASRequest:
		packet.NgapProcedureCode = "RerouteNASRequest"
		packet.NgapType = "RerouteNASRequest"
		packet.DirSeq = -1
		RerouteNASRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 37
	case ngapType.ProcedureCodeRRCInactiveTransitionReport:
		packet.NgapProcedureCode = "RRCInactiveTransitionReport"
		packet.NgapType = "RRCInactiveTransitionReport"
		packet.DirSeq = 1
		RRCInactiveTransitionReportDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 38
	case ngapType.ProcedureCodeTraceFailureIndication:
		packet.NgapProcedureCode = "TraceFailureIndication"
		packet.NgapType = "TraceFailureIndication"
		packet.DirSeq = 1
		TraceFailureIndicationDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 39
	case ngapType.ProcedureCodeTraceStart:
		packet.NgapProcedureCode = "TraceStart"
		packet.NgapType = "TraceStart"
		packet.DirSeq = -1
		TraceStartDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 40
	case ngapType.ProcedureCodeUEContextModification:
		packet.NgapProcedureCode = "UEContextModification"
		packet.NgapType = "UEContextModificationRequest"
		packet.DirSeq = -1
		UEContextModificationRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 41
	case ngapType.ProcedureCodeUEContextRelease:
		packet.NgapProcedureCode = "UEContextRelease"
		packet.NgapType = "UEContextReleaseCommand"
		packet.DirSeq = -1
		UEContextReleaseCommandDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 42
	case ngapType.ProcedureCodeUEContextReleaseRequest:
		packet.NgapProcedureCode = "UEContextReleaseRequest"
		packet.NgapType = "UEContextReleaseRequest"
		packet.DirSeq = 1
		UEContextReleaseRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 43
	case ngapType.ProcedureCodeUERadioCapabilityCheck:
		packet.NgapProcedureCode = "UERadioCapabilityCheck"
		packet.NgapType = "UERadioCapabilityCheckRequest"
		packet.DirSeq = -1
		UERadioCapabilityCheckRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 44
	case ngapType.ProcedureCodeUERadioCapabilityInfoIndication:
		packet.NgapProcedureCode = "UERadioCapabilityInfoIndication"
		packet.NgapType = "UERadioCapabilityInfoIndication"
		packet.DirSeq = 1
		UERadioCapabilityInfoIndicationDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 45
	case ngapType.ProcedureCodeUETNLABindingRelease:
		packet.NgapProcedureCode = "UETNLABindingRelease"
		packet.NgapType = "UETNLABindingReleaseRequest"
		packet.DirSeq = -1
		UETNLABindingReleaseRequestDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 46
	case ngapType.ProcedureCodeUplinkNASTransport:
		packet.NgapProcedureCode = "UplinkNASTransport"
		packet.NgapType = "UplinkNASTransport"
		packet.DirSeq = 1
		UplinkNASTransportDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 47
	case ngapType.ProcedureCodeUplinkNonUEAssociatedNRPPaTransport:
		packet.NgapProcedureCode = "UplinkNonUEAssociatedNRPPaTransport"
		packet.NgapType = "UplinkNonUEAssociatedNRPPaTransport"
		packet.DirSeq = 1
		UplinkNonUEAssociatedNRPPaTransportDecoder()
	// ProcedureCode = 48
	case ngapType.ProcedureCodeUplinkRANConfigurationTransfer:
		packet.NgapProcedureCode = "UplinkRANConfigurationTransfer"
		packet.NgapType = "UplinkRANConfigurationTransfer"
		packet.DirSeq = 1
		UplinkRANConfigurationTransferDecoder()
	// ProcedureCode = 49
	case ngapType.ProcedureCodeUplinkRANStatusTransfer:
		packet.NgapProcedureCode = "UplinkRANStatusTransfer"
		packet.NgapType = "UplinkRANStatusTransfer"
		packet.DirSeq = 1
		UplinkRANStatusTransferDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 50
	case ngapType.ProcedureCodeUplinkUEAssociatedNRPPaTransport:
		packet.NgapProcedureCode = "UplinkUEAssociatedNRPPaTransport"
		packet.NgapType = "UplinkUEAssociatedNRPPaTransport"
		packet.DirSeq = 1
		UplinkUEAssociatedNRPPaTransportDecoder(InitiatingMessage.Value, packet)
	// ProcedureCode = 51
	case ngapType.ProcedureCodeWriteReplaceWarning:
		packet.NgapProcedureCode = "WriteReplaceWarning"
		packet.NgapType = "WriteReplaceWarningRequest"
		packet.DirSeq = -1
		WriteReplaceWarningRequestDecoder()
	// ProcedureCode = 52
	case ngapType.ProcedureCodeSecondaryRATDataUsageReport:
		packet.NgapProcedureCode = "SecondaryRATDataUsageReport"
		packet.NgapType = "SecondaryRATDataUsageReport"
		packet.DirSeq = 1
		SecondaryRATDataUsageReportDecoder(InitiatingMessage.Value, packet)
	}
}

func RouteSuccessfulOutcome(DecResult *ngapType.NGAPPDU, packet *flowMap.Packet) {
	SuccessfulOutcome := DecResult.SuccessfulOutcome
	switch SuccessfulOutcome.ProcedureCode.Value {
	// ProcedureCode = 0
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		packet.NgapProcedureCode = "AMFConfigurationUpdate"
		packet.NgapType = "AMFConfigurationUpdateAcknowledge"
		packet.DirSeq = 1
		AMFConfigurationUpdateAcknowledgeDecoder()
	// ProcedureCode = 10
	case ngapType.ProcedureCodeHandoverCancel:
		packet.NgapProcedureCode = "HandoverCancellation"
		packet.NgapType = "HandoverCancelAcknowledge"
		packet.DirSeq = -1
		HandoverCancelAcknowledgeDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 12
	case ngapType.ProcedureCodeHandoverPreparation:
		packet.NgapProcedureCode = "HandoverPreparation"
		packet.NgapType = "HandoverCommand"
		packet.DirSeq = -1
		HandoverCommandDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 13
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		packet.NgapProcedureCode = "HandoverResourceAllocation"
		packet.NgapType = "HandoverRequestAcknowledge"
		packet.DirSeq = 1
		HandoverRequestAcknowledgeDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 14
	case ngapType.ProcedureCodeInitialContextSetup:
		packet.NgapProcedureCode = "InitialContextSetup"
		packet.NgapType = "InitialContextSetupResponse"
		packet.DirSeq = 1
		InitialContextSetupResponseDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 20
	case ngapType.ProcedureCodeNGReset:
		packet.NgapProcedureCode = "NGReset"
		packet.NgapType = "NGResetAcknowledge"
		packet.DirSeq = 0 //TODO: -1 or 1?
		NGResetAcknowledgeDecoder()
	// ProcedureCode = 21
	case ngapType.ProcedureCodeNGSetup:
		packet.NgapProcedureCode = "NGSetup"
		packet.NgapType = "NGSetupResponse"
		packet.DirSeq = -1
		NGSetupResponseDecoder()
	// ProcedureCode = 25
	case ngapType.ProcedureCodePathSwitchRequest:
		packet.NgapProcedureCode = "PathSwitchRequest"
		packet.NgapType = "PathSwitchRequestAcknowledge"
		packet.DirSeq = -1
		PathSwitchRequestAcknowledgeDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 26
	case ngapType.ProcedureCodePDUSessionResourceModify:
		packet.NgapProcedureCode = "PDUSessionResourceModify"
		packet.NgapType = "PDUSessionResourceModifyResponse"
		packet.DirSeq = 1
		PDUSessionResourceModifyResponseDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 27
	case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
		packet.NgapProcedureCode = "PDUSessionResourceModifyIndication"
		packet.NgapType = "PDUSessionResourceModifyConfirm"
		packet.DirSeq = -1
		PDUSessionResourceModifyConfirmDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 28
	case ngapType.ProcedureCodePDUSessionResourceRelease:
		packet.NgapProcedureCode = "PDUSessionResourceRelease"
		packet.NgapType = "PDUSessionResourceReleaseResponse"
		packet.DirSeq = 1
		PDUSessionResourceReleaseResponseDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 29
	case ngapType.ProcedureCodePDUSessionResourceSetup:
		packet.NgapProcedureCode = "PDUSessionResourceSetup"
		packet.NgapType = "PDUSessionResourceSetupResponse"
		packet.DirSeq = 1
		PDUSessionResourceSetupResponseDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 32
	case ngapType.ProcedureCodePWSCancel:
		packet.NgapProcedureCode = "PWSCancel"
		packet.NgapType = "PWSCancelResponse"
		packet.DirSeq = 1
		PWSCancelResponseDecoder()
	// ProcedureCode = 35
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		packet.NgapProcedureCode = "RANConfiguration"
		packet.NgapType = "RANConfigurationUpdateAcknowledge"
		packet.DirSeq = -1
		RANConfigurationUpdateAcknowledgeDecoder()
	// ProcedureCode = 40
	case ngapType.ProcedureCodeUEContextModification:
		packet.NgapProcedureCode = "UEContextModification"
		packet.NgapType = "UEContextModificationResponse"
		packet.DirSeq = 1
		UEContextModificationResponseDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 41
	case ngapType.ProcedureCodeUEContextRelease:
		packet.NgapProcedureCode = "UEContextRelease"
		packet.NgapType = "UEContextReleaseComplete"
		packet.DirSeq = 1
		UEContextReleaseCompleteDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 43
	case ngapType.ProcedureCodeUERadioCapabilityCheck:
		packet.NgapProcedureCode = "UERadioCapabilityCheck"
		packet.NgapType = "UERadioCapabilityCheckResponse"
		packet.DirSeq = 1
		UERadioCapabilityCheckResponseDecoder(SuccessfulOutcome.Value, packet)
	// ProcedureCode = 51
	case ngapType.ProcedureCodeWriteReplaceWarning:
		packet.NgapProcedureCode = "WriteReplaceWarning"
		packet.NgapType = "WriteReplaceWarningResponse"
		packet.DirSeq = 1
		WriteReplaceWarningResponseDecoder()
	}
}

func RouteUnsuccessfulOutcome(DecResult *ngapType.NGAPPDU, packet *flowMap.Packet) {
	UnsuccessfulOutcome := DecResult.UnsuccessfulOutcome
	switch UnsuccessfulOutcome.ProcedureCode.Value {
	// ProcedureCode = 0
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		packet.NgapProcedureCode = "AMFConfigurationUpdate"
		packet.NgapType = "AMFConfigurationUpdateFailure"
		packet.DirSeq = 1
		AMFConfigurationUpdateFailureDecoder()
	// ProcedureCode = 12
	case ngapType.ProcedureCodeHandoverPreparation:
		packet.NgapProcedureCode = "HandoverPreparation"
		packet.NgapType = "HandoverPreparationFailure"
		packet.DirSeq = -1
		HandoverPreparationFailureDecoder(UnsuccessfulOutcome.Value, packet)
	// ProcedureCode = 13
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		packet.NgapProcedureCode = "HandoverResourceAllocation"
		packet.NgapType = "HandoverFailure"
		packet.DirSeq = 1
		HandoverFailureDecoder(UnsuccessfulOutcome.Value, packet)
	// ProcedureCode = 14
	case ngapType.ProcedureCodeInitialContextSetup:
		packet.NgapProcedureCode = "InitialContextSetup"
		packet.NgapType = "InitialContextSetupFailure"
		packet.DirSeq = 1
		InitialContextSetupFailureDecoder(UnsuccessfulOutcome.Value, packet)
	// ProcedureCode = 21
	case ngapType.ProcedureCodeNGSetup:
		packet.NgapProcedureCode = "NGSetup"
		packet.NgapType = "NGSetupFailure"
		packet.DirSeq = -1
		NGSetupFailureDecoder()
	// ProcedureCode = 25
	case ngapType.ProcedureCodePathSwitchRequest:
		packet.NgapProcedureCode = "PathSwitchRequest"
		packet.NgapType = "PathSwitchRequestFailure"
		packet.DirSeq = -1
		PathSwitchRequestFailureDecoder(UnsuccessfulOutcome.Value, packet)
	// ProcedureCode = 35
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		//fmt.Println("==================== RANConfigurationUpdate ====================")
		packet.NgapProcedureCode = "RANConfigurationUpdate"
		packet.NgapType = "RANConfigurationUpdateFailure"
		packet.DirSeq = -1
		RANConfigurationUpdateFailureDecoder()
	// ProcedureCode = 40
	case ngapType.ProcedureCodeUEContextModification:
		//fmt.Println("==================== UEContextModification ====================")
		packet.NgapProcedureCode = "UEContextModification"
		packet.NgapType = "UEContextModificationFailure"
		packet.DirSeq = 1
		UEContextModificationFailureDecoder(UnsuccessfulOutcome.Value, packet)
	}
}
