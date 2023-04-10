package main

import (
	"bufio"
	"fmt"
	"github.com/colinmarc/hdfs/v2"
	"os"
	"path/filepath"
)

func main() {
	address := os.Args[1]
	path := os.Args[2]

	client, err := hdfs.New(address)
	if err != nil {
		fmt.Println("err create new client")
		fmt.Println(err.Error())
		return
	}
	defer client.Close()

	dir, err := client.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, fileInfo := range dir {
		func() {
			p := filepath.Join(path, fileInfo.Name())
			file, _ := client.Open(p)
			fmt.Println("open file : " + p)
			scanner := bufio.NewScanner(file)
			//optionally, resize scanner's capacity for lines over 64K, see next example
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			err = scanner.Err()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			defer file.Close()
		}()
	}
}
