# -*- coding: utf-8 -*-
import random

import numpy as np

ProcedureCode = ({
    "AMFConfigurationUpdate": 0,
    "AMFStatusIndication": 1,
    "CellTrafficTrace": 2,
    "DeactivateTrace": 3,
    "DownlinkNASTransport": 4,
    "DownlinkNonUEAssociatedNRPPaTransport": 5,
    "DownlinkRANConfigurationTransfer": 6,
    "DownlinkRANStatusTransfer": 7,
    "DownlinkUEAssociatedNRPPaTransport": 8,
    "ErrorIndication": 9,
    "HandoverCancellation": 10,
    "HandoverCancel": 10,
    "HandoverNotification": 11,
    "HandoverPreparation": 12,
    "HandoverPrepararion": 12,
    "HandoverResourceAllocation": 13,
    "InitialContextSetup": 14,
    "InitialUEMessage": 15,
    "LocationReportingControl": 16,
    "LocationReportingFailureIndication": 17,
    "LocationReport": 18,
    "NASNonDeliveryIndication": 19,
    "NGReset": 20,
    "NGSetup": 21,
    "OverloadStart": 22,
    "OverloadStop": 23,
    "Paging": 24,
    "PathSwitchRequest": 25,
    "PDUSessionResourceModify": 26,
    "PDUSessionResourceModifyIndication": 27,
    "PDUSessionResourceRelease": 28,
    "PDUSessionResourceSetup": 29,
    "PDUSessionResourceNotify": 30,
    "PrivateMessage": 31,
    "PWSCancel": 32,
    "PWSFailureIndication": 33,
    "PWSRestartIndication": 34,
    "RANConfigurationUpdate": 35,
    "RerouteNASRequest": 36,
    "RRCInactiveTransitionReport": 37,
    "TraceFailureIndication": 38,
    "TraceStart": 39,
    "UEContextModification": 40,
    "UEContextRelease": 41,
    "UEContextReleaseRequest": 42,
    "UERadioCapabilityCheck": 43,
    "UERadioCapabilityInfoIndication": 44,
    "UETNLABindingRelease": 45,
    "UplinkNASTransport": 46,
    "UplinkNonUEAssociatedNRPPaTransport": 47,
    "UplinkRANConfigurationTransfer": 48,
    "UplinkRANStatusTransfer": 49,
    "UplinkUEAssociatedNRPPaTransport": 50,
    "WriteReplaceWarning": 51,
    "SecondaryRATDataUsageReport": 52,
})

def ngap_feature_extract(item):
    """
    item 是一个字典，包含 'Dirseq','Time','Length','Label'
    """
    # if max(list(item['Label'])) >= 1:
    #     label = 1
    # else:
    #     label = 0
    # Dirseq: 1表示 NG-RAN -> AMF, -1表示 AMF -> NG-RAN
    dirseq = list(item["DirSeq"])
    timeseq = list(item["Time"])  # 时间戳列表
    length = list(item["PacketLen"])  # 包长度列表
    total_packet = len(dirseq)  # 总数据包数

    timeseq = [item - timeseq[0] + 1e-5 for item in timeseq]

    # f1 f2 First Request Content Size and RTT.
    outgoing_index = [j for j, x in enumerate(
        dirseq) if int(x) == 1]  # 传出数据包的index
    incoming_index = [j for j, x in enumerate(
        dirseq) if int(x) == -1]  # 传入数据包的index
    # f2:第一个传出数据包和第二个传出数据包之间传入数据包的大小

    # f1:第一个传出数据包和第一个传入数据包之间的延迟

    # f3~f10 Statistics of packets size and number.
    f3 = sum([length[j]
             for j, x in enumerate(dirseq) if int(x) == 1])  # f3:传出数据包大小
    f303 = sum([length[j] for j, x in enumerate(
        dirseq) if int(x) == -1])  # f303:传入数据包大小
    f6 = len([j for j, x in enumerate(dirseq) if int(x) == -1])  # f6:传入数据包数
    f7 = len([j for j, x in enumerate(dirseq) if int(x) == 1])  # f7:传出数据包数
    f8 = total_packet  # f8:总数据包数
    f4 = f303 / sum(length)  # f4:传入数据包大小/总数据包大小
    f5 = f3 / sum(length)  # f5:传出数据包大小/总数据包大小
    f9 = f6 / f8  # f9:传入数据包数/总数据包数
    f10 = f7 / f8  # f10:传出数据包数/总数据包数

    # f11~f14 The number of incoming, outgoing packets, the fraction of the number of incoming packets, and the fraction of the number of outgoing packets in the first 20 packets of the network flows.
    # f11 = len([j for j, x in enumerate(dirseq[:20]) if int(x) == -1]) if (
    #         total_packet > 20 or total_packet == 20) else len(
    #     [j for j, x in enumerate(dirseq) if int(x) == -1])  # 前20个包的传入数据包数
    # f12 = len([j for j, x in enumerate(dirseq[:20]) if int(x) == 1]) if (
    #         total_packet > 20 or total_packet == 20) else len(
    #     [j for j, x in enumerate(dirseq) if int(x) == 1])  # 前20个包的传出数据包数
    # f13 = f11 / 20 if (total_packet > 20 or total_packet ==
    #                    20) else f11 / len(dirseq)  # 传入数据包占前20个包的比例
    # f14 = f12 / 20 if (total_packet > 20 or total_packet ==
    #                    20) else f12 / len(dirseq)  # 传出数据包占前20个包的比例

    # f15~f18 We generate two lists by recording the number of packets before every incoming and outgoing packet. Then we compute the average and standard deviation values of these two lists respectively
    l1 = [j for j, x in enumerate(dirseq) if int(x) == 1]  # 每个传出数据包前的数据包数量
    l2 = [j for j, x in enumerate(dirseq) if int(x) == -1]  # 每个传入数据包前的数据包数量
    if len(l1) == 0:
        f15 = 0
        f17 = 0
    else:
        f15 = np.mean(l1)  # 列表平均值
        f17 = np.std(l1)  # 列表标准差
    if len(l2) == 0:
        f16 = 0
        f18 = 0
    else:
        f16 = np.mean(l2)  # 列表平均值
        f18 = np.std(l2)  # 列表标准差

    # f19~f33 Statistics of packet inter-arrival time.
    lll1 = timeseq  # 总数据包时间序列
    lll1 = [data - lll1[0] for data in lll1]
    lll4 = [lll1[1] - lll1[0]]
    lll4.extend([lll1[n] - lll1[n - 1] for n in range(2, len(lll1))])
    f19 = np.max(lll4)
    f20 = np.min(lll4)
    f21 = np.mean(lll4)
    f22 = np.std(lll4)
    f23 = np.quantile(lll4, .75)

    if len([timeseq[j] for j, x in enumerate(dirseq) if int(x) == 1]) <= 1:  # 如果没有传出数据包/只有一个
        lll2 = [0]
        lll5 = [0]
        f24 = 0
        f25 = 0
        f26 = 0
        f27 = 0
        f28 = 0
    else:
        lll2 = [timeseq[j]
                for j, x in enumerate(dirseq) if int(x) == 1]  # 传出数据包时间序列
        lll2 = [data - lll1[0] for data in lll2]
        lll5 = [lll2[1] - lll2[0]]
        lll5.extend([lll2[n] - lll2[n - 1] for n in range(2, len(lll2))])
        f24 = np.max(lll5)
        f25 = np.min(lll5)
        f26 = np.mean(lll5)
        f27 = np.std(lll5)
        f28 = np.quantile(lll5, .75)

    if len([timeseq[j] for j, x in enumerate(dirseq) if int(x) == -1]) <= 1:
        lll3 = [0]
        lll6 = [0]
        f29 = 0
        f30 = 0
        f31 = 0
        f32 = 0
        f33 = 0
    else:
        lll3 = [timeseq[j]
                for j, x in enumerate(dirseq) if int(x) == -1]  # 传入数据包时间序列
        lll3 = [data - lll1[0] for data in lll3]
        lll6 = [lll3[1] - lll3[0]]
        lll6.extend([lll3[n] - lll3[n - 1] for n in range(2, len(lll3))])
        f29 = np.max(lll6)
        f30 = np.min(lll6)
        f31 = np.mean(lll6)
        f32 = np.std(lll6)
        f33 = np.quantile(lll6, .75)

    # f34~f42 Statistics of transmission time
    f34 = np.quantile(lll1, .25)
    f35 = np.quantile(lll1, .5)
    f36 = np.quantile(lll1, .75)
    f37 = np.quantile(lll2, .25)
    f38 = np.quantile(lll2, .5)
    f39 = np.quantile(lll2, .75)
    f40 = np.quantile(lll3, .25)
    f41 = np.quantile(lll3, .5)
    f42 = np.quantile(lll3, .75)
    feature = [f3, f303, f4, f5, f6, f7, f8, f9, f10, f15, f16, f17, f18, f19, f20, f21, f22,
               f23, f24, f25, f26, f27, f28, f29, f30, f31, f32, f33, f34, f35, f36, f37, f38, f39, f40, f41, f42]

    # f43~102 The quantity and the transmission size speed of incoming, outgoing and total packets sequences.（用1除的）
    lll7 = [1 / j for j in lll4 if j != 0]
    lll8 = [1 / j for j in lll5 if j != 0]
    lll9 = [1 / j for j in lll6 if j != 0]
    f43_62 = random.sample(lll7, 20) if len(lll7) > 20 or len(
        lll7) == 20 else lll7 + [0] * (20 - len(lll7))
    f63_82 = random.sample(lll8, 20) if len(lll8) > 20 or len(
        lll8) == 20 else lll8 + [0] * (20 - len(lll8))
    f83_102 = random.sample(lll9, 20) if len(lll9) > 20 or len(
        lll9) == 20 else lll9 + [0] * (20 - len(lll9))
    feature.extend(f43_62)
    feature.extend(f63_82)
    feature.extend(f83_102)

    # f103~162 The quantity and the transmission size speed of incoming, outgoing and total packets sequences.
    llll4 = length
    lll10 = [llll4[j + 1] / x for j, x in enumerate(lll4) if x != 0]
    f103_122 = random.sample(lll10, 20) if len(lll10) > 20 or len(
        lll10) == 20 else lll10 + [0] * (20 - len(lll10))

    if len([timeseq[j] for j, x in enumerate(dirseq) if int(x) == 1]) <= 1:  # 如果传出数据包数为0/1
        f123_142 = [0] * 20
    else:
        llll5 = [length[j] for j, x in enumerate(dirseq) if int(x) == 1]
        lll11 = [llll5[j + 1] / x for j, x in enumerate(lll5) if x != 0]
        f123_142 = random.sample(lll11, 20) if len(lll11) > 20 or len(
            lll11) == 20 else lll11 + [0] * (20 - len(lll11))

    if len([timeseq[j] for j, x in enumerate(dirseq) if int(x) == -1]) <= 1:
        f143_162 = [0] * 20
    else:
        llll6 = [length[j] for j, x in enumerate(dirseq) if int(x) == -1]
        lll12 = [llll6[j + 1] / x for j, x in enumerate(lll6) if x != 0]
        f143_162 = random.sample(lll12, 20) if len(lll12) > 20 or len(
            lll12) == 20 else lll12 + [0] * (20 - len(lll12))

    feature.extend(f103_122)
    feature.extend(f123_142)
    feature.extend(f143_162)

    # f163~f262 The cumulative size of packets
    s = [0]
    for i in range(total_packet):
        s.append(int(dirseq[i]) * length[i] + s[i])
    x = np.linspace(0, total_packet, total_packet + 1)
    xvals = np.linspace(0, total_packet, 50)
    yinterp = np.interp(xvals, x, np.array(s))
    feature.extend(yinterp.tolist())

    for i in range(len(feature)):
        if np.isnan(feature[i]):
            feature[i] = 0

    # NGAP特征
    f303_355 = np.zeros(53, dtype=int)  # 所有53个信令的数量
    init303_355 = np.zeros(53, dtype=int)  # 所有53个信令的init
    # procedureCode[0,10,12,13,14,20,21,25,26,27,28,29,32,35,40,41,43,51]
    success303_355 = np.zeros(53, dtype=int)
    # procedureCode[0,12,13,14,21,25,35,40]
    fail303_355 = np.zeros(53, dtype=int)

    for _, row in item.iterrows():
        procedureCode = ProcedureCode[row["ProcedureCode"]]
        initiatingMessage = int(row["InitiatingMessage"])
        successful = int(row["SuccessfulOutcome"])
        unsuccessful = int(row["UnsuccessfulOutcome"])

        f303_355[procedureCode] += 1
        if initiatingMessage == 1:
            init303_355[procedureCode] += 1
        if successful == 1:
            success303_355[procedureCode] += 1
        if unsuccessful == 1:
            fail303_355[procedureCode] += 1

    feature.extend(f303_355)  # 所有53个信令的数量

    f356_373 = np.zeros(53)  # 信令成功率
    f374_381 = np.zeros(53)  # 信令失败率
    f382_399 = np.zeros(53)  # 信令响应率

    non_zero_init_indices = np.nonzero(init303_355)
    f356_373[non_zero_init_indices] = success303_355[non_zero_init_indices] / \
        init303_355[non_zero_init_indices]
    f374_381[non_zero_init_indices] = fail303_355[non_zero_init_indices] / \
        init303_355[non_zero_init_indices]
    f382_399 = f356_373 + f374_381
    f356_373 = f356_373[[0, 10, 12, 13, 14, 20, 21,
                         25, 26, 27, 28, 29, 32, 35, 40, 41, 43, 51]]
    f374_381 = f374_381[[0, 12, 13, 14, 21, 25, 35, 40]]
    f382_399 = f382_399[[0, 10, 12, 13, 14, 20, 21,
                         25, 26, 27, 28, 29, 32, 35, 40, 41, 43, 51]]

    feature.extend(f356_373)
    feature.extend(f374_381)
    feature.extend(f382_399)
    return feature
