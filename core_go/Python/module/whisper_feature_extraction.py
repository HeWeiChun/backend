import json
import math
import numpy as np

# 协议类型
weight = [12.0, 24.0, 1000.0]
min_interval_time = 1e-5

NGAP_TYPE = dict({'PDUSessionResourceSetupRequest': 1, 'PDUSessionResourceSetupResponse': 100001,
                  'PDUSessionResourceReleaseCommand': 40, 'PDUSessionResourceReleaseResponse': 100040,
                  'PDUSessionResourceModifyRequest': 75, 'PDUSessionResourceModifyResponse': 100075,
                  'PDUSessionResourceNotify': 100,
                  'PDUSessionResourceModifyIndication': 301, 'PDUSessionResourceModifyConfirm': 100301,
                  'InitialContextSetupRequest': 550, 'InitialContextSetupResponse': 100550,
                  'InitialContextSetupFailure': 9900550,
                  'UEContextReleaseRequest': 799,
                  'UEContextReleaseCommand': 1000, 'UEContextReleaseComplete': 101000,
                  'UEContextModificationRequest': 2001, 'UEContextModificationResponse': 102001,
                  'UEContextModificationFailure': 9902001,
                  'RRCInactiveTransitionReport': 3050,
                  'HandoverRequired': 3099, 'HandoverCommand': 103099, 'HandoverPreparationFailure': 9903099,
                  'HandoverRequestAcknowledge': 4100,
                  'HandoverNotify': 5301,
                  'PathSwitchRequest': 6550, 'PathSwitchRequestAcknowledge': 106550,
                  'PathSwitchRequestFailure': 9906550,
                  'HandoverCancel': 7799, 'HandoverCancelAcknowledge': 107799,
                  'UplinkRANStatusTransfer': 10000,
                  'DownlinkRANStatusTransfer': 110000,
                  'InitialUEMessage': 20001,
                  'DownlinkNASTransport': 30050,
                  'UplinkNASTransport': 130050,
                  'NASNonDeliveryIndication': 40099,
                  'RerouteNASRequest': 50100,
                  'ErrorIndication': 9910000,
                  'DownlinkUEAssociatedNRPPaTransport': 50301,
                  'UplinkUEAssociatedNRPPaTransport': 150301,
                  'TraceStart': 60550,
                  'TraceFailureIndication': 60799,
                  'DeactivateTrace': 61000,
                  'CellTrafficTrace': 62001,
                  'LocationReportingControl': 63050,
                  'LocationReportingFailureIndication': 9990000,
                  'LocationReport': 74099,
                  'UETNLABindingReleaseRequest': 75301,
                  'UERadioCapabilityInfoIndication': 76550,
                  'UERadioCapabilityCheckRequest': 77799, 'UERadioCapabilityCheckResponse': 177799,
                  'SecondaryRATDataUsageReport': 80000, })


def extraction(packet):
    N = len(packet)  # 数据包数量
    feature = []
    # print(N)
    # 定义超参数
    C = 10
    Wseg = 18
    if N < Wseg:
        for i in range(N, Wseg):
            packet.append([1, 0, 1])
    N = len(packet)
    Kf = int(Wseg / 2) + 1
    w = weight
    # Packet Feature Encoding
    v = []
    for i in range(0, N):
        p0 = packet[i][0]
        p1 = float(packet[i][1]) / 1000000000 + min_interval_time
        if packet[i][2] in NGAP_TYPE:
            p2 = NGAP_TYPE[packet[i][2]]
        else:
            p2 = 11
        v.append(p2 * w[2] + p0 * w[0] + -math.log2(p1) * w[1])

    # Vector Framing
    f = []
    i = int(0)

    while i < N - Wseg:
        f.append(v[i:i + Wseg])
        i = i + Wseg
    f.append(v[N - Wseg:N])
    Nf = len(f)
    # Discrete Fourier Transformation

    for i in range(0, Nf):

        temp_ri = []
        for k in range(0, Wseg):
            temp_aik = 0
            temp_bik = 0
            for n in range(1, Wseg - 1):
                temp_aik = temp_aik + f[i][n - 1] * math.cos(2 * np.pi * (n - 1) * k / Wseg)
                temp_bik = temp_bik - f[i][n - 1] * math.sin(2 * np.pi * (n - 1) * k / Wseg)
            temp_ri.append((math.log(temp_aik * temp_aik + temp_bik * temp_bik) + 1) / C)
        temp_ri = temp_ri[0:Kf]
        feature.append(temp_ri)
    return feature


def packetParse(filename, feature, weight, len_go):
    file = open(filename, 'r', encoding='utf-8')
    for line in file.readlines():
        packet = []
        dic = json.loads(line)
        for i in range(dic['TotalNum']):
            packet_length = int(dic['PacketLength'][i])
            packet_time = float(dic['TimeInterval'][i]) + min_interval_time
            Ntype = str(dic['NGAPType'][i])
            Ntype_new = Ntype.replace(',', '')
            if Ntype_new in NGAP_TYPE:
                packet_type = NGAP_TYPE[Ntype_new]
            else:
                packet_type = 11
            if packet_time >= 0:
                packet.append([packet_length, packet_time, packet_type])
        N = len(packet)  # 数据包数量

        # print(N)
        # 定义超参数
        C = 10
        Wseg = 18
        if N < Wseg:
            for i in range(N, Wseg):
                packet.append([1, min_interval_time, 1])
        N = len(packet)
        Kf = int(Wseg / 2) + 1
        w = weight
        # Packet Feature Encoding
        v = []
        for i in range(0, N):
            v.append(packet[i][2] * w[2] + packet[i][0] * w[0] + -math.log2(packet[i][1]) * w[1])

        # Vector Framing
        f = []
        i = int(0)

        while i < N - Wseg:
            f.append(v[i:i + Wseg])
            i = i + len_go
        f.append(v[N - Wseg:N])
        Nf = len(f)
        # Discrete Fourier Transformation

        for i in range(0, Nf):

            temp_ri = []
            for k in range(0, Wseg):
                temp_aik = 0
                temp_bik = 0
                for n in range(1, Wseg - 1):
                    temp_aik = temp_aik + f[i][n - 1] * math.cos(2 * np.pi * (n - 1) * k / Wseg)
                    temp_bik = temp_bik - f[i][n - 1] * math.sin(2 * np.pi * (n - 1) * k / Wseg)
                temp_ri.append((math.log(temp_aik * temp_aik + temp_bik * temp_bik) + 1) / C)
            temp_ri = temp_ri[0:Kf]
            feature.append(temp_ri)
    return feature


def N2_TestData(file_nom, file_ano, weight):
    # print("测试初始数据处理...")
    feature_nom = packetParse(file_nom, [], weight, 18)
    feature_ano = packetParse(file_ano, [], weight, 18)
    len_nom = len(feature_nom)
    len_ano = len(feature_ano)
    len_all = len_ano + len_nom
    tag_nom = np.ones(len_nom)
    tag_ano = np.zeros(len_ano)
    tag = []
    feature = []
    tag[0:len_nom] = tag_nom
    tag[len_nom:len_all] = tag_ano
    feature[0:len_nom] = feature_nom
    feature[len_nom:len_all] = feature_ano
    # print("数据处理成功！")
    return feature, tag


def N2_TestNom(file_ano, weight):
    feature_nom = packetParse(file_ano, weight, 6)
    len_nom = len(feature_nom)
    tag_nom = np.ones(len_nom)
    return feature_nom, tag_nom


def N2_TestAno(file_ano, weight):
    feature_ano = packetParse(file_ano, weight, 6)
    len_ano = len(feature_ano)
    tag_ano = np.zeros(len_ano)
    return feature_ano, tag_ano
