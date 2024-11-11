package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func addFileToZip(zipWriter *zip.Writer, baseDir, inputPath string) {
	//获取相对目录
	relativePath, err := filepath.Rel(baseDir, inputPath)
	if err != nil {
		panic(err)
	}

	//创建输入文件
	inputFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	//创建压缩器
	zipFileWriter, err := zipWriter.Create(relativePath)
	if err != nil {
		panic(err)
	}

	// 使用缓冲区逐步读取并写入
	buf := make([]byte, 4096)
	var totalWrite int64
	info, err := inputFile.Stat()
	if err != nil {
		panic(err)
	}
	totalSize := info.Size()
	for {
		n, err := inputFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			return
		}
		if n == 0 {
			break
		}
		w, err := zipFileWriter.Write(buf[:n])
		if err != nil {
			fmt.Println("Error writing to zip entry:", err)
			return
		}
		totalWrite += int64(w)
		progressBar.SetValue(float64(totalWrite) / float64(totalSize))
	}

	//从输入流复制到输出流
	// _, err = io.Copy(zipFileWriter, inputFile)
	return
}

func zipCompressSingleFile(inputPath, outputPath string) error {
	//创建输入文件
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	//创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	//创建压缩器
	writer := zip.NewWriter(outputFile)
	defer writer.Close()

	//写入数据
	zipFileWriter, err := writer.Create(filepath.Base(inputPath))
	if err != nil {
		return err
	}
	// _, err = io.Copy(zipFileWriter, inputFile)
	// if err != nil {
	// 	return err
	// }

	// 使用缓冲区逐步读取并写入
	buf := make([]byte, 4096)
	var totalWrite int64
	info, err := inputFile.Stat()
	if err != nil {
		panic(err)
	}
	totalSize := info.Size()
	startTime := time.Now()
	for {
		n, err := inputFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		w, err := zipFileWriter.Write(buf[:n])
		if err != nil {
			return err
		}
		totalWrite += int64(w)
		progressBar.SetValue(float64(totalWrite) / float64(totalSize))
		usedTime := time.Since(startTime).Seconds()
		if usedTime > 0 {
			speed := float64(totalWrite) / usedTime
			remainingBytes := float64(totalSize - totalWrite)
			restTime := remainingBytes / speed
			restTimeFileLabel.SetText(fmt.Sprintf("剩余时间: %.2fs", restTime))
		}
	}
	return nil
}

func zipCompressDir(inputDir, outputDir string) error {
	//创建输出文件
	outputFile, err := os.Create(outputDir)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	//创建压缩器
	zipWriter := zip.NewWriter(outputFile)
	defer zipWriter.Close()

	// var wg sync.WaitGroup

	//如果是文件就直接压缩

	//遍历文件
	err = filepath.Walk(inputDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//不是文件夹就直接压缩
		if !info.IsDir() {
			// //获取相对目录
			// relativePath, err := filepath.Rel(inputDir, path)
			// if err != nil {
			// 	panic(err)
			// }
			// //创建压缩器
			// zipFileWriter, err := zipWriter.Create(relativePath)
			// if err != nil {
			// 	panic(err)
			// }
			// wg.Add(1)
			// go addFileToZip(zipFileWriter, path, &wg)
			addFileToZip(zipWriter, inputDir, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	// wg.Wait()
	return nil
}

func zipCompressPath(inputPath, outputPath string) error {
	info, err := os.Stat(inputPath)
	fmt.Println(info)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return zipCompressDir(inputPath, outputPath)
	} else {
		return zipCompressSingleFile(inputPath, outputPath)
	}
}

func extractFileFromZip(zipFile *zip.File, dstDir string) error {
	//获得绝对路径
	absPath := filepath.Join(dstDir, zipFile.Name)

	//如果是文件夹就直接创建
	if zipFile.FileInfo().IsDir() {
		return os.MkdirAll(absPath, os.ModePerm)
	}
	//先创建父目录
	err := os.MkdirAll(filepath.Dir(absPath), os.ModePerm)
	if err != nil {
		return err
	}

	//创建解压器
	zipFileReader, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer zipFileReader.Close()

	//创建输出文件
	outputFile, err := os.Create(absPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, zipFileReader)
	return err
}

func zipDecompressDir(inputDir, outputDir string) error {
	//打开压缩文件
	zipReader, err := zip.OpenReader(inputDir)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	unzipSize := 0
	totalSize := len(zipReader.File)

	//遍历压缩文件
	for _, file := range zipReader.File {
		//读取单个文件
		err := extractFileFromZip(file, outputDir)
		if err != nil {
			return err
		}
		unzipSize++
		restTimeFileLabel.SetText(fmt.Sprintf("当前进度: %d/%d", unzipSize, totalSize))
		progressBar.SetValue(float64(unzipSize) / float64(totalSize))
	}
	return nil
}
