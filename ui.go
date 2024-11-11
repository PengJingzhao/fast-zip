package main

// import (
// 	"fmt"
// 	"os"
// 	"gioui.org/op"

// 	"gioui.org/layout"
// 	"gioui.org/widget"
// 	"gioui.org/widget/material"
// )

// // 定义按钮和状态
// var (
// 	selectFileBtn    widget.Clickable
// 	selectedFileInfo string
// )

// // 主事件循环
// func loop(w *app.Window) error {
// 	th := material.NewTheme()
// 	var ops op.Ops

// 	for {
// 		switch e := w.Event().(type) {
// 		case app.DestroyEvent:
// 			return e.Err
// 		case app.FrameEvent:
// 			gtx := app.NewContext(&ops, e)

// 			// 检查按钮点击事件
// 			if selectFileBtn.Clicked(gtx) {
// 				file, err := os.Open("选择文件的路径或通过对话框选择")
// 				if err == nil {
// 					fileInfo, _ := file.Stat()
// 					selectedFileInfo = fmt.Sprintf("文件名: %s\n文件路径: %s\n文件大小: %d 字节",
// 						fileInfo.Name(), file.Name(), fileInfo.Size())
// 					file.Close()
// 				} else {
// 					selectedFileInfo = fmt.Sprintf("文件打开失败: %v", err)
// 				}
// 			}

// 			// 布局 UI
// 			layout.Flex{
// 				Axis: layout.Vertical,
// 			}.Layout(gtx,
// 				layout.Rigid(material.Button(th, &selectFileBtn, "选择文件").Layout),
// 				layout.Rigid(material.Body1(th, selectedFileInfo).Layout),
// 			)

// 			e.Frame(gtx.Ops)
// 		}
// 	}
// }

// func main() {
// 	go func() {
// 		app := NewApp()
// 		if err := app.Loop(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()
// 	app.Main()
// }

// package main

// import (
// 	"image/color"

// 	"gioui.org/app"
// 	"gioui.org/io/key"
// 	"gioui.org/op"
// 	"gioui.org/text"

// 	"gioui.org/widget/material"
// )

// func run(window *app.Window) error {
// 	theme := material.NewTheme()
// 	var ops op.Ops
// 	for {
// 		switch e := window.Event().(type) {
// 		case key.Event:

// 		case app.DestroyEvent:
// 			return e.Err
// 		case app.FrameEvent:
// 			// This graphics context is used for managing the rendering state.
// 			gtx := app.NewContext(&ops, e)

// 			//标题
// 			title := material.H1(theme, "选择源文件夹")
// 			title.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
// 			title.Alignment = text.Middle
// 			title.Layout(gtx)

// 			// //图标
// 			// icon := widget.Icon(*theme.Icon.CheckBoxChecked)
// 			// iconSize := unit.Dp(48)
// 			// icon.Layout(gtx, color)
// 			// Pass the drawing operations to the GPU.
// 			e.Frame(gtx.Ops)
// 		}
// 	}
// }
