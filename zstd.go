package main

import (
    "archive/tar"
    "fmt"
    "io"
    "os"
    "path/filepath"

    "github.com/klauspost/compress/zstd"
)


func zstdCompressFile(inputPath, outputPath string) error {
	// 打开输入文件
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputFile.Close()

	// 创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	//创建压缩器
	// 创建 zstd 编码器
	encoder, err := zstd.NewWriter(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create zstd encoder: %w", err)
	}
	defer encoder.Close()

	// 将输入文件内容写入到 zstd 编码器进行压缩
	if _, err := io.Copy(encoder, inputFile); err != nil {
		return fmt.Errorf("failed to compress file: %w", err)
	}

	return nil
}

func zstdDecompressFile(inputPath, outputPath string) error {
	// 打开压缩文件
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open compressed file: %w", err)
	}
	defer inputFile.Close()

	// 创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()
	compressor(inputFile, outputFile)

	return nil
}

// 将文件添加到 tar 中
func addFileToTar(tw *tar.Writer, filePath string, baseDir string) error {
    // 获取相对路径
    relativePath, err := filepath.Rel(baseDir, filePath)
    if err != nil {
        return err
    }

    // 获取文件信息
    fileInfo, err := os.Stat(filePath)
    if err != nil {
        return err
    }

    // 创建 tar 文件头
    header, err := tar.FileInfoHeader(fileInfo, relativePath)
    if err != nil {
        return err
    }
    header.Name = relativePath

    // 写入文件头到 tar
    if err := tw.WriteHeader(header); err != nil {
        return err
    }

    // 如果是文件，写入内容
    if !fileInfo.IsDir() {
        file, err := os.Open(filePath)
        if err != nil {
            return err
        }
        defer file.Close()

        _, err = io.Copy(tw, file)
        if err != nil {
            return err
        }
    }
    return nil
}

// 压缩目录
func compressDirectory(inputDir, outputZst string) error {
    // 创建输出文件
    outputFile, err := os.Create(outputZst)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    // 创建 zstd 压缩器
    encoder, err := zstd.NewWriter(outputFile)
    if err != nil {
        return err
    }
    defer encoder.Close()

    // 创建 tar 写入器
    tarWriter := tar.NewWriter(encoder)
    defer tarWriter.Close()

    // 遍历目录并将每个文件添加到 tar 中
    err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        return addFileToTar(tarWriter, path, inputDir)
    })
    return err
}


func compressor(inputFile io.Reader, outputFile io.Writer) error {
	// 创建 zstd 解码器
	decoder, err := zstd.NewReader(inputFile)
	if err != nil {
		return fmt.Errorf("failed to create zstd decoder: %w", err)
	}
	defer decoder.Close()

	// 将解压缩后的数据写入输出文件
	if _, err := io.Copy(outputFile, decoder); err != nil {
		return fmt.Errorf("failed to decompress file: %w", err)
	}
	return nil
}