package c_base

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcmd"
)

func PluginDriverCommand(getDriver func() IDriver) *gcmd.Command {
	Main := &gcmd.Command{
		Name: "main",
	}
	Version := &gcmd.Command{
		Name:  "version",
		Brief: "Show version",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			des := getDriver().GetDescription()
			if des == nil || des.Version == "" {
				fmt.Println("unknown")
			} else {
				fmt.Println(des.Version)
			}
			return err
		},
	}
	Info := &gcmd.Command{
		Name: "info",
		Arguments: []gcmd.Argument{
			{Name: "debug", Short: "d", Brief: "Show debug info"},
		},
		Brief: "Show info",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			if parser.GetOpt("debug").Interface() != nil {
				fmt.Println("debug")
			}

			des := getDriver().GetDescription()
			if des == nil {
				fmt.Println("unknown")
			} else {
				fmt.Printf("%+v\n", des)
			}
			return err
		},
	}
	_ = Main.AddCommand(Version, Info)
	return Main
}
