package main

import (
	"fmt"
	"github.com/injoyai/goutil/oss"
	"github.com/injoyai/goutil/oss/tray"
	"github.com/injoyai/lorca"
	"github.com/injoyai/tool/forward"
)

func main() {
	port := 60073
	db := oss.UserInjoyDir("/forward/database/forward.db")
	go forward.Run(port, db)
	tray.Run(
		tray.WithIco(ico),
		tray.WithHint("端口转发"),
		func(s *tray.Tray) {
			s.AddMenu().SetName("配置").SetIcon(tray.IconSetting).OnClick(func(m *tray.Menu) {
				lorca.Run(&lorca.Config{
					Width:  800,
					Height: 600,
					Index:  fmt.Sprintf("http://localhost:%d", port),
				})
			})
		},
		tray.WithStartup(),
		tray.WithSeparator(),
		tray.WithExit(),
	)
}
