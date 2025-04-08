package main

import (
	"errors"
	_ "gf_star/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"gf_star/internal/cmd"
)

func main() {
	var err error

	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")

	// 检测数据库是否正常连接
	err = connDb()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("数据库连接失败")
	}
	return nil
}
