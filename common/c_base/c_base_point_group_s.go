package c_base

type SPointGroup struct {
	GroupName string `json:"groupName" dc:"组名称"`
	GroupSort int    `json:"groupSort" dc:"组排序"`
	Display   bool   `json:"display" dc:"是否显示"` // todo  这个字段暂时没用，先留着
}
