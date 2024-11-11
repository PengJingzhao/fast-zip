package main

import (
	"io"
	"os"

	"github.com/golang/snappy"
)

func snappyCompress(srcContent []byte) []byte {
	dst := snappy.Encode(nil, srcContent)
	return dst
}

func snappyCompressFile(srcPath, dstPath string) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	data, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err)
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	writer := snappy.NewBufferedWriter(dstFile)

	writer.Write(data)
}

func snappyDeCompress(srcContent []byte) []byte {
	deCompress, err := snappy.Decode(nil, srcContent)
	if err != nil {
		panic(err)
	}
	return deCompress
}
