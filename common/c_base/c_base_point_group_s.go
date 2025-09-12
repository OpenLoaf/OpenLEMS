package c_base

type SPointGroup struct {
	GroupKey  string `json:"groupKey" v:"required" dc:"组key"` // key是可以重复
	GroupName string `json:"groupName" v:"required" dc:"组名称"` // 名字可以重复
	GroupSort int    `json:"groupSort" dc:"组排序"`
	Disable   bool   `json:"disable" dc:"是否显示"` // todo  这个字段暂时没用，先留着
}
