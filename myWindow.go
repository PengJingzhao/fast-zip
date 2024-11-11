package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"path/filepath"
	"strings"
)

func CreateMainWindow() {
	myWindow := myApp.NewWindow("FastZip")
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	// 标签
	selectedFileLabel = widget.NewLabel("输入：未选择文件")
	targetFileLabel = widget.NewLabel("输出：未选择文件")
	restTimeFileLabel = widget.NewLabel("剩余时间: 0s")
	// 进度条
	progressBar = widget.NewProgressBar()

	// 压缩文件按钮
	var inputPath string
	var outputPath string
	compressDirButton := widget.NewButton("压缩文件夹", func() {
		//选择要压缩的源文件
		dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if lu == nil {
				return
			}
			inputPath = lu.Path()
			selectedFileLabel.SetText("要压缩的文件: " + inputPath)

			if inputPath != "" {
				initialURI, err := storage.ListerForURI(storage.NewFileURI(filepath.Dir(inputPath)))
				if err != nil {
					panic(err)
				}
				//选择目标位置
				saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.ShowError(err, myWindow)
						return
					}
					if writer == nil {
						return
					}

					progressBar.SetValue(0)
					go func() {
						outputPath = writer.URI().Path()
						targetFileLabel.SetText("保存位置：" + outputPath)
						err := zipCompressPath(inputPath, outputPath)

						if err != nil {
							dialog.ShowError(err, myWindow)
						} else {
							dialog.ShowInformation("提示", "文件压缩成功", myWindow)
						}
					}()
				}, myWindow)
				filename := filepath.Base(inputPath)
				saveDialog.SetLocation(initialURI)
				saveDialog.SetFileName(filename + ".zip")
				saveDialog.Show()
			}
		}, myWindow)
	})
	compressButton := widget.NewButton("压缩文件", func() {
		//选择要压缩的源文件
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				return
			}
			inputPath = reader.URI().Path()
			selectedFileLabel.SetText("要压缩的文件: " + inputPath)

			if inputPath != "" {
				initialURI, err := storage.ListerForURI(storage.NewFileURI(filepath.Dir(inputPath)))
				if err != nil {
					panic(err)
				}
				//选择目标位置
				saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.ShowError(err, myWindow)
						return
					}
					if writer == nil {
						return
					}

					progressBar.SetValue(0)
					go func() {
						outputPath = writer.URI().Path()
						targetFileLabel.SetText("保存位置：" + outputPath)
						err := zipCompressPath(inputPath, outputPath)

						if err != nil {
							dialog.ShowError(err, myWindow)
						} else {
							dialog.ShowInformation("提示", "文件压缩成功", myWindow)
						}
					}()
				}, myWindow)
				filename := filepath.Base(inputPath)
				filename = strings.TrimSuffix(filename, filepath.Ext(filename))
				saveDialog.SetLocation(initialURI)
				saveDialog.SetFileName(filename + ".zip")
				saveDialog.Show()
			}
		}, myWindow)
	})
	compressButton.Resize(fyne.NewSize(140, 140))

	// 解压文件按钮
	decompressButton := widget.NewButton("解压文件", func() {
		//选择输入文件
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				return
			}
			inputPath = reader.URI().Path()
			selectedFileLabel.SetText("要解压的文件: " + inputPath)

			if inputPath != "" {
				initialURI, err := storage.ListerForURI(storage.NewFileURI(filepath.Dir(inputPath)))
				if err != nil {
					panic(err)
				}
				//选择保存位置
				saveDialog := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
					if err != nil {
						dialog.ShowError(err, myWindow)
						return
					}
					if folder == nil {
						return
					}

					progressBar.SetValue(0)
					targetFileLabel.SetText("保存位置：" + folder.Path())
					//解压文件
					go func() {
						err := zipDecompressDir(inputPath, folder.Path())

						if err != nil {
							dialog.ShowError(err, myWindow)
						} else {
							dialog.ShowInformation("提示", "文件解压成功", myWindow)
						}
					}()
				}, myWindow)
				saveDialog.SetLocation(initialURI)
				saveDialog.Show()
			}

		}, myWindow)
	})
	decompressButton.Resize(fyne.NewSize(140, 140))

	// 居中布局
	content := container.NewVBox(
		widget.NewLabelWithStyle("FastZip", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		selectedFileLabel,
		targetFileLabel,
		container.NewGridWithColumns(5,
			compressDirButton,
			compressButton,
			decompressButton,
		),
		restTimeFileLabel,
		progressBar,
	)

	myWindow.SetContent(container.NewCenter(content))
	myWindow.ShowAndRun()
}
