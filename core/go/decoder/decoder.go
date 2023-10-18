package decoder

import (
	"fmt"
	"github.com/Freddy/sctp_flowmap/decoder/NGAP"
	"github.com/Freddy/sctp_flowmap/flowMap"
	"github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"strconv"
)

func Decode(file string, task string) {
	//fmt.Println("===============This project is for decoding the pcap file of NGAP and getting the info of some necessary bytes.===============")
	//fmt.Println("===============Created by Galadriel on 6/26/2022===============")
	handle := new(pcap.Handle)
	handle, err := pcap.OpenOffline(file)
	TimeFirst := int64(0)
	//fmt.Println(err)
	//fmt.Println("SUC")
	//handle, _ := pcap.OpenOffline("5gc-ngap.pcapng")
	defer handle.Close()
	//fmt.Println("do")
	fmt.Println(&handle, handle)
	packetSource := gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	num := 0
	for packet := range packetSource.Packets() {
		applicationLayer := packet.ApplicationLayer()

		if applicationLayer != nil {
			payload := applicationLayer.Payload()

			Packet_UE := &flowMap.Packet{
				NgapPayloadBytes:    payload,
				ArriveTimeUs:        packet.Metadata().Timestamp.UnixNano(),
				ArriveTime:          packet.Metadata().Timestamp,
				RAN_UE_NGAP_ID:      -1,
				PayloadBytes:        packet.Data(),
				PacketLen:           uint32(len(packet.Data())),
				TimeID:              "0",
				InitiatingMessage:   0,
				SuccessfulOutcome:   0,
				UnsuccessfulOutcome: 0,
			}

			var DecResult *ngapType.NGAPPDU
			DecResult, err = ngap.Decoder(payload)
			if err == nil {
				num = num + 1

				if num == 1 {
					TimeFirst = Packet_UE.ArriveTimeUs
				}

				NGAP.RouteNGAP(DecResult, Packet_UE)
				sctpLayer := packet.Layer(layers.LayerTypeSCTP)
				if sctpLayer != nil {
					if sctp, ok := sctpLayer.(*layers.SCTP); ok {
						Packet_UE.VerificationTag = sctp.VerificationTag
					}
				}
				ip4Layer := packet.Layer(layers.LayerTypeIPv4)
				ip6Layer := packet.Layer(layers.LayerTypeIPv6)
				if ip4Layer != nil {
					if ipv4, ok := ip4Layer.(*layers.IPv4); ok {
						Packet_UE.DstIP = ipv4.DstIP.String()
						Packet_UE.SrcIP = ipv4.SrcIP.String()
					}
				}
				if ip6Layer != nil {
					if ipv6, ok := ip6Layer.(*layers.IPv6); ok {
						Packet_UE.DstIP = ipv6.DstIP.String()
						Packet_UE.SrcIP = ipv6.SrcIP.String()
					}
				}

				Packet_Time := &flowMap.Packet{
					NgapPayloadBytes:    Packet_UE.NgapPayloadBytes,
					ArriveTimeUs:        Packet_UE.ArriveTimeUs,
					ArriveTime:          Packet_UE.ArriveTime,
					RAN_UE_NGAP_ID:      Packet_UE.RAN_UE_NGAP_ID,
					PayloadBytes:        packet.Data(),
					PacketLen:           uint32(len(packet.Data())),
					VerificationTag:     Packet_UE.VerificationTag,
					DstIP:               Packet_UE.DstIP,
					SrcIP:               Packet_UE.SrcIP,
					NgapType:            Packet_UE.NgapType,
					NgapProcedureCode:   Packet_UE.NgapProcedureCode,
					InitiatingMessage:   Packet_UE.InitiatingMessage,
					SuccessfulOutcome:   Packet_UE.SuccessfulOutcome,
					UnsuccessfulOutcome: Packet_UE.UnsuccessfulOutcome,
				}

				flowUEID, UEID := flowMap.Count_UE_ID(Packet_UE, task)
				strFlowUEID := strconv.FormatUint(flowUEID, 10)
				Packet_UE.FlowID = strFlowUEID
				if UEID {
					flowMap.Put(Packet_UE, flowMap.FlowTable_UE, strFlowUEID, task)
				}

				flowTimeID, TimeID := flowMap.Count_Time_ID(Packet_Time, TimeFirst, task)
				strFlowTimeID := strconv.FormatUint(flowTimeID, 10)
				Packet_Time.FlowID = strFlowTimeID
				if TimeID {
					flowMap.Put(Packet_Time, flowMap.FlowTable_Time, strFlowTimeID, task)
				}
			}
		}
	}
}
