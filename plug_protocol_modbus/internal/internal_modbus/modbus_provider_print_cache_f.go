package internal_modbus

import (
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"strings"
	"text/tabwriter"
)

func (p *ModbusProvider) PrintCacheValues() {
	keys, err := p.cache.Keys(p.ctx)
	if err != nil || len(keys) == 0 {
		return
	}

	// 创建一个 strings.Builder 来构建表格内容
	var builder strings.Builder

	// 创建一个新的 tabwriter，写入 strings.Builder
	writer := tabwriter.NewWriter(&builder, 0, 0, 2, ' ', 0)
	// 写入表格头
	_, _ = writer.Write([]byte("     Addr\tLevel\tName\t     Value\tDesc\t"))
	//_, _ = writer.Write([]byte("    -\t    \t            -\t            -\t\n"))

	// 写入表格内容

	message := fmt.Sprintf("共有%d个缓存点位,当前告警等级:%s 共%d条告警", len(keys), p.GetAlarmLevel(), len(p.GetAlarmDetails()))
	array := garray.NewSortedStrArray()
	for _, k := range keys {
		if k == nil {
			continue
		}
		value, err := p.cache.Get(p.ctx, k)
		if err != nil {
			continue
		}
		meta := k.(*c_base.Meta)

		var cn string
		if meta.Cn == "" {
			cn = meta.Desc
		} else {
			cn = meta.Cn + "; " + meta.Desc
		}

		if meta.Precise == 0 {
			meta.Precise = 2
		}

		//meta.ValueToString(value)
		array.Add(fmt.Sprintf("\n%5d[0x%X]\t%s\t%s\t%10s\t%s\t", meta.Addr, meta.Addr, meta.Level.Name(), meta.Name, meta.ValueToString(value), cn))
	}
	for _, i2 := range array.Slice() {
		_, _ = writer.Write([]byte(i2))
	}

	// 刷新 writer，将缓冲区的内容写入到 strings.Builder
	_ = writer.Flush()

	p.log.Noticef(p.ctx, "%v \n%v", message, builder.String())
}
