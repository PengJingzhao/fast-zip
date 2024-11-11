package main

// import (
// 	"fmt"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	// 创建应用程序实例
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("文件选择器示例")

// 	// 标签用于显示文件路径
// 	label := widget.NewLabel("选择的文件路径将在此显示")

// 	// 创建按钮并配置点击事件
// 	button := widget.NewButton("选择文件", func() {
// 		// 打开文件选择对话框
// 		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
// 			// 错误处理
// 			if err != nil {
// 				label.SetText("错误: " + err.Error())
// 				return
// 			}
// 			if reader == nil {
// 				// 用户取消选择
// 				label.SetText("未选择文件")
// 				return
// 			}
// 			// 显示文件路径
// 			label.SetText("选择的文件路径: " + reader.URI().Path())
// 			fmt.Println("选择的文件:", reader.URI().Path())
// 		}, myWindow)
// 	})

// 	// 将按钮和标签添加到窗口中
// 	myWindow.SetContent(container.NewVBox(
// 		button,
// 		label,
// 	))

// 	// 显示窗口并运行应用
// 	myWindow.ShowAndRun()
// }