package main

import (
	"bufio"
	"fmt"
	"github.com/colinmarc/hdfs/v2"
	"os"
	"path/filepath"
)

const (
	_  = 1 << (10 * iota) // ignore first value
	KB                    // 1 << 10
	MB                    // 1 << 20
	GB                    // 1 << 30
	TB                    // 1 << 40
	PB                    // 1 << 50
	EB                    // 1 << 60
)

func newClient(nameNode string) (*hdfs.Client, error) {
	client, err := hdfs.New(nameNode)
	if err != nil {
		fmt.Println("err create new client")
		fmt.Println(err.Error())
		return nil, err
	}
	return client, nil
}

func get() {
	address := os.Args[2]
	path := os.Args[3]

	client, err := newClient(address)
	if err != nil {
		return
	}
	defer client.Close()

	dir, err := client.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("MB = %d", MB)

	buf := make([]byte, 10*MB)

	for _, fileInfo := range dir {
		func() {
			p := filepath.Join(path, fileInfo.Name())
			file, _ := client.Open(p)
			fmt.Println("open file : " + p)
			scanner := bufio.NewScanner(file)
			scanner.Buffer(buf, 64*MB)
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

func put() {
	address := os.Args[2]
	src := os.Args[3]
	dst := os.Args[4]
	client, err := newClient(address)
	if err != nil {
		return
	}
	defer client.Close()
	err = client.CopyToRemote(src, dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("put [%s] to [%s]\n", src, dst)
}

func del() {
	address := os.Args[2]
	dst := os.Args[3]
	client, err := newClient(address)
	if err != nil {
		return
	}
	defer client.Close()
	err = client.Remove(dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("del [%s]\n", dst)
}

func mkdir() {
	address := os.Args[2]
	dir := os.Args[3]
	client, err := newClient(address)
	if err != nil {
		return
	}
	defer client.Close()
	err = client.Mkdir(dir, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("mkdir [%s]\n", dir)
}

func ls() {
	address := os.Args[2]
	path := os.Args[3]

	client, err := newClient(address)
	if err != nil {
		return
	}
	defer client.Close()

	dir, err := client.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, fileInfo := range dir {
		fmt.Printf("file path = :%s\n", fileInfo.Name())
	}
}

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "get":
		get()
	case "put":
		put()
	case "del":
		del()
	case "mkdir":
		mkdir()
	case "ls":
		ls()
	default:
		fmt.Printf("not supported command : [%s]\n", cmd)
	}
}
