package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

func addFile2Zip(zipWriter *zip.Writer, fileContent, fileName string) error {
	// 创建一个内存中的文本文件
	fileWriter, err := zipWriter.Create(fileName)
	if err != nil {
		return err
	}

	// 将内存中的文件数据写入 Zip 文件
	_, err = io.Copy(fileWriter, bytes.NewReader([]byte(fileContent)))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// 创建一个内存中的 Zip 文件
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	err := addFile2Zip(zipWriter, "hello, world", "hello.txt")
	if err != nil {
		panic(err)
	}
	err = addFile2Zip(zipWriter, "hello, world", "hello2.txt")
	if err != nil {
		panic(err)
	}

	// 关闭 Zip 文件
	err = zipWriter.Close()
	if err != nil {
		panic(err)
	}

	// 将内存中的 Zip 文件写入磁盘文件
	zipContent := buf.Bytes()
	err = os.WriteFile("output.zip", zipContent, 0644)
	if err != nil {
		panic(err)
	}
}
