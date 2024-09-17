package ammeter_acrel_10r_v1

import "common/c_base"

// 时间
var (
	Year   = &c_base.Meta{Name: "Year", Cn: "年", Addr: 128, ReadType: c_base.RBcd16, SystemType: c_base.SUint16, Desc: "年 BCD码表示"}
	Month  = &c_base.Meta{Name: "Month", Cn: "月", Addr: 129, ReadType: c_base.RBcd16, SystemType: c_base.SUint16, Desc: "月 1~12"}
	Day    = &c_base.Meta{Name: "Day", Cn: "日", Addr: 130, ReadType: c_base.RBcd16, SystemType: c_base.SUint16, Desc: "日 1~31"}
	Hour   = &c_base.Meta{Name: "Hour", Cn: "时", Addr: 131, ReadType: c_base.RBcd16, SystemType: c_base.SUint16, Desc: "时 0~23"}
	Minute = &c_base.Meta{Name: "Minute", Cn: "分", Addr: 132, ReadType: c_base.RBcd16, SystemType: c_base.SUint16, Desc: "分 0~59"}
	Second = &c_base.Meta{Name: "Second", Cn: "秒", Addr: 133, ReadType: c_base.RBcd16, SystemType: c_base.SUint16, Desc: "秒 0~59"}
)

// 电能数据
var (
	Ua  = &c_base.Meta{Name: "Ua", Cn: "A相电压", Addr: 8192, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "V", Desc: "A 相电压 UA 一次侧，单位 V"}
	Ub  = &c_base.Meta{Name: "Ub", Cn: "B相电压", Addr: 8194, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "V", Desc: "B 相电压 UB 一次侧，单位 V"}
	Uc  = &c_base.Meta{Name: "Uc", Cn: "C相电压", Addr: 8196, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "V", Desc: "C 相电压 UC 一次侧，单位 V"}
	Uab = &c_base.Meta{Name: "Uab", Cn: "AB线电压", Addr: 8198, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "V", Desc: "AB 线电压 UAB 一次侧，单位 V"}
	Ubc = &c_base.Meta{Name: "Ubc", Cn: "BC线电压", Addr: 8200, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "V", Desc: "BC 线电压 UBC 一次侧，单位 V"}
	Uca = &c_base.Meta{Name: "Uca", Cn: "CA线电压", Addr: 8202, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "V", Desc: "CA 线电压 UCA 一次侧，单位 V"}
	Ia  = &c_base.Meta{Name: "Ia", Cn: "A相电流", Addr: 8204, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "A", Desc: "A 相电流 IA 一次侧，单位 A"}
	Ib  = &c_base.Meta{Name: "Ib", Cn: "B相电流", Addr: 8206, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "A", Desc: "B 相电流 IB 一次侧，单位 A"}
	Ic  = &c_base.Meta{Name: "Ic", Cn: "C相电流", Addr: 8208, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "A", Desc: "C 相电流 IC 一次侧，单位 A"}
	Pa  = &c_base.Meta{Name: "Pa", Cn: "A相有功功率", Addr: 8212, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kw", Desc: "A 相有功功率 PA 一次侧功率，单位Kw"}
	Pb  = &c_base.Meta{Name: "Pb", Cn: "B相有功功率", Addr: 8214, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kw", Desc: "B 相有功功率 PB 一次侧功率，单位Kw"}
	Pc  = &c_base.Meta{Name: "Pc", Cn: "C相有功功率", Addr: 8216, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kw", Desc: "C 相有功功率 PC 一次侧功率，单位Kw"}
	Pt  = &c_base.Meta{Name: "Pt", Cn: "总有功功率", Addr: 8218, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kw", Desc: "总有功功率 PT"}
	Qa  = &c_base.Meta{Name: "Qa", Cn: "A相无功功率", Addr: 8220, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kvar", Desc: "A 相无功功率 QA"}
	Qb  = &c_base.Meta{Name: "Qb", Cn: "B相无功功率", Addr: 8222, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kvar", Desc: "B 相无功功率 QB"}
	Qc  = &c_base.Meta{Name: "Qc", Cn: "C相无功功率", Addr: 8224, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kvar", Desc: "C 相无功功率 QC"}
	Qt  = &c_base.Meta{Name: "Qt", Cn: "总无功功率", Addr: 8226, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kvar", Desc: "总无功功率 QT"}
	Sa  = &c_base.Meta{Name: "Sa", Cn: "A相视在功率", Addr: 8228, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kva", Desc: "A 相视在功率 SA"}
	Sb  = &c_base.Meta{Name: "Sb", Cn: "B相视在功率", Addr: 8230, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kva", Desc: "B 相视在功率 SB"}
	Sc  = &c_base.Meta{Name: "Sc", Cn: "C相视在功率", Addr: 8232, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kva", Desc: "C 相视在功率 SC"}
	St  = &c_base.Meta{Name: "St", Cn: "总视在功率", Addr: 8234, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Kva", Desc: "总视在功率 ST"}
	Pfa = &c_base.Meta{Name: "Pfa", Cn: "A相功率因数", Addr: 8236, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Desc: "A 相功率因数 PFA"}
	Pfb = &c_base.Meta{Name: "Pfb", Cn: "B相功率因数", Addr: 8238, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Desc: "B 相功率因数 PFB"}
	Pfc = &c_base.Meta{Name: "Pfc", Cn: "C相功率因数", Addr: 8240, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Desc: "C 相功率因数 PFC"}
	Pft = &c_base.Meta{Name: "Pft", Cn: "总功率因数", Addr: 8242, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Desc: "总功率因数 PFT"}
	F   = &c_base.Meta{Name: "F", Cn: "频率", Addr: 8244, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Unit: "Hz", Desc: "频率 F"}
)

// 总电能
var (
	Epi = &c_base.Meta{Name: "Epi", Cn: "总正向有功电能", Addr: 12418, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.001, Unit: "Kwh", Desc: "正向有功电能 Epi"}
	Eqe = &c_base.Meta{Name: "Eqe", Cn: "总反向有功电能", Addr: 12420, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.001, Unit: "Kwh", Desc: "正向无功电能 Eqi"}
	Eql = &c_base.Meta{Name: "Eql", Cn: "总正向无功电能", Addr: 12424, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.001, Unit: "Kwh", Desc: "反向有功电能 Eqe"}
	Eqc = &c_base.Meta{Name: "Eqc", Cn: "总反向无功电能", Addr: 12426, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.001, Unit: "Kwh", Desc: "反向无功电能 Eql"}
)

// 开关量输入
var (
	SwitchIn1  = &c_base.Meta{Name: "SwitchIn1", Cn: "第一路开关量输入", Addr: 53, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "有开入时为 1，无开入时为 0"}
	SwitchIn2  = &c_base.Meta{Name: "SwitchIn2", Cn: "第二路开关量输入", Addr: 54, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "有开入时为 1，无开入时为 0"}
	SwitchIn3  = &c_base.Meta{Name: "SwitchIn3", Cn: "第三路开关量输入", Addr: 55, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "有开入时为 1，无开入时为 0"}
	SwitchIn4  = &c_base.Meta{Name: "SwitchIn4", Cn: "第四路开关量输入", Addr: 56, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "有开入时为 1，无开入时为 0"}
	SwitchIn5  = &c_base.Meta{Name: "SwitchIn5", Cn: "第五路开关量输入", Addr: 57, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "有开入时为 1，无开入时为 0"}
	SwitchIn6  = &c_base.Meta{Name: "SwitchIn6", Cn: "第六路开关量输入", Addr: 58, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "有开入时为 1，无开入时为 0"}
	SwitchIn7  = &c_base.Meta{Name: "SwitchIn7", Cn: "第七路开关量输入", Addr: 59, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "有开入时为 1，无开入时为 0"}
	SwitchIn8  = &c_base.Meta{Name: "SwitchIn8", Cn: "第八路开关量输入", Addr: 60, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "有开入时为 1，无开入时为 0"}
	SwitchOut1 = &c_base.Meta{Name: "SwitchOut1", Cn: "第一路开关量输出", Addr: 61, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "写 1 时输出继电器触点闭合，  写 0 时输出继电器触点分开"}
	SwitchOut2 = &c_base.Meta{Name: "SwitchOut2", Cn: "第二路开关量输出", Addr: 62, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "写 1 时输出继电器触点闭合，  写 0 时输出继电器触点分开"}
	SwitchOut3 = &c_base.Meta{Name: "SwitchOut3", Cn: "第三路开关量输出", Addr: 63, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "写 1 时输出继电器触点闭合，  写 0 时输出继电器触点分开"}
	SwitchOut4 = &c_base.Meta{Name: "SwitchOut4", Cn: "第四路开关量输出", Addr: 64, ReadType: c_base.RUint16, SystemType: c_base.SBool, Desc: "写 1 时输出继电器触点闭合，  写 0 时输出继电器触点分开"}
)

/*
1030	过频率	组合报警参数，-9999 – 9999 仅限第二路报  警为组合报警时有效,详见 6.2.6.4；例：显  示值为 66.00Kw,通讯值为 6600
1031	欠频率	组合报警参数，-9999 – 9999 仅限第二路报  警为组合报警时有效,详见 6.2.6.4；例：显  示值为 66.00Kw,通讯值为 6600
1032	过功率	组合报警参数，-9999 – 9999 仅限第二路报  警为组合报警时有效,详见 6.2.6.4；例：显  示值为 66.00Kw,通讯值为 6600
1033	欠功率	组合报警参数，-9999 – 9999 仅限第二路报  警为组合报警时有效,详见 6.2.6.4；例：显  示值为 66.00Kw,通讯值为 6600
1034	过电流	组合报警参数，-9999 – 9999 仅限第二路报  警为组合报警时有效,详见 6.2.6.4；例：显  示值为 66.00Kw,通讯值为 6600
1035	欠功率因数	组合报警参数，-9999 – 9999 仅限第二路报  警为组合报警时有效,详见 6.2.6.4；例：显  示值为 66.00Kw,通讯值为 6600
1036	过电压不平衡	详见 6.2.6.4 ， 例 ： 显示值为  55.00Kw,通讯值为 5500
1037	过电流不平衡	详见 6.2.6.4 ， 例 ： 显示值为  55.00Kw,通讯值为 5500
1038	组合报警状态	第 0 位表示过电压报警状态，第一位表示欠电  压报警状态，依次类推到第 9 位
*/
