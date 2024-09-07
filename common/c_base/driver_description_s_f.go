package c_base

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
	"text/tabwriter"
)

type SDescription struct {
	Brand      string        `json:"brand"`      // 品牌
	Model      string        `json:"model"`      // 型号
	Version    string        `json:"version"`    // 版本
	Create     string        `json:"create"`     // 创建时间
	BuildTime  string        `json:"buildTime"`  // 编译时间
	CommitHash string        `json:"commitHash"` // 提交哈希
	Author     string        `json:"author"`     // 作者
	Remark     string        `json:"remark"`     // 备注
	Telemetry  []*STelemetry `json:"telemetry"`  // 遥测
}

func BuildDescriptionFromYaml(ctx context.Context, yaml []byte) *SDescription {
	//g.Log().Debugf(ctx, "BuildDescriptionFromYaml: %s", string(yaml))
	description := &SDescription{}
	err := gyaml.DecodeTo(yaml, &description)
	if err != nil {
		panic(gerror.Newf("解析版本信息失败！请检查version.yaml文件!%v", err))
	}
	return description
}

func (s *SDescription) String() string {

	// 创建一个 strings.Builder 来构建表格内容
	var builder strings.Builder

	// 创建一个新的 tabwriter，写入 strings.Builder
	writer := tabwriter.NewWriter(&builder, 0, 0, 3, ' ', 0)
	//_, _ = writer.Write([]byte("| Basic Information |\n"))
	// 写入表格头

	_, _ = writer.Write([]byte(fmt.Sprintf("Brand\t:\t%s\t\n", s.Brand)))
	_, _ = writer.Write([]byte(fmt.Sprintf("Model\t:\t%s\t\n", s.Model)))
	_, _ = writer.Write([]byte(fmt.Sprintf("Version\t:\t%s\t\n", s.Version)))
	_, _ = writer.Write([]byte(fmt.Sprintf("Author\t:\t%s\t\n", s.Author)))
	_, _ = writer.Write([]byte(fmt.Sprintf("CreateTime\t:\t%s\t\n", s.Create)))
	if s.BuildTime != "" {
		_, _ = writer.Write([]byte(fmt.Sprintf("BuildTime\t:\t%s\t\n", s.BuildTime)))
	}
	if s.CommitHash != "" {
		_, _ = writer.Write([]byte(fmt.Sprintf("CommitHash\t:\t%s\t\n", s.CommitHash)))
	}
	_, _ = writer.Write([]byte(fmt.Sprintf("Remark\t:\t%s\t\n", s.Remark)))

	//_, _ = writer.Write([]byte("Version\tBrand\tModel\tRemark\t"))
	//_, _ = writer.Write([]byte(fmt.Sprintf("\n%s\t%s\t%s\t%s\t", s.Version, s.Brand, s.Model, s.Remark)))

	if len(s.Telemetry) != 0 {
		_, _ = writer.Write([]byte("\nTelemetry Information:\t\n"))
		_, _ = writer.Write([]byte("Name\tNationalization\tUnit\tRemark\t"))

		for _, telemetry := range s.Telemetry {
			_, _ = writer.Write([]byte("\n" + telemetry.String()))
		}

	}

	_ = writer.Flush()
	//telemetryStr := ""
	//for _, telemetry := range s.Telemetry {
	//	telemetryStr += telemetry.String() + "\n\t"
	//}

	return builder.String()
}

func (s *SDescription) GetDescription() *SDescription {
	return s
}
