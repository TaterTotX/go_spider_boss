package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	//appMenu := menu.NewMenu()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "boss直聘自动打招呼机器人",
		Width:  400,
		Height: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		DisableResize: true,

		Mac: &mac.Options{
			//TitleBar:     mac.TitleBarHiddenInset(),
			//窗口标题栏的配置
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false, // 标题栏是否透明
				HideTitle:                  false, // 是否隐藏窗口的标题
				HideTitleBar:               false, // 是否隐藏整个标题栏
				FullSizeContent:            false, // 内容是否扩展到全尺寸，包括标题栏下面
				UseToolbar:                 false, // 是否使用工具栏
				HideToolbarSeparator:       true,  // 是否隐藏工具栏和内容之间的分隔线

			},

			// 应用程序的外观设置，这里设置为深色系
			Appearance: mac.NSAppearanceNameDarkAqua,

			// 网页视图是否透明
			WebviewIsTransparent: true,
			// 窗口是否半透明
			WindowIsTranslucent: true,
			// 关于窗口的信息
			About: &mac.AboutInfo{
				Title:   "bossapp",  // 应用程序的名称
				Message: "土豆© 2024", // 关于窗口中显示的信息，例如版权信息

			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
