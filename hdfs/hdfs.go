package main

import (
	"bufio"
	"fmt"
	"github.com/colinmarc/hdfs/v2"
	"github.com/colinmarc/hdfs/v2/hadoopconf"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	address := os.Args[1]
	path := os.Args[2]
	conf := hadoopconf.HadoopConf{}
	conf["fs.defaultFS"] = "hdfs://qy-ieg-4-v2"
	conf["dfs.nameservices"] = "qy-ieg-4-v2"
	conf["dfs.ha.namenodes.qy-ieg-4-v2"] = "nn1,nn2,nn3"
	conf["dfs.namenode.rpc-address.qy-ieg-4-v2.nn1"] = "qy-ieg-4-v2-nn-1.tencent-distribute.com:9000"
	conf["dfs.namenode.http-address.qy-ieg-4-v2.nn1"] = "qy-ieg-4-v2-nn-1.tencent-distribute.com:8080"
	conf["dfs.namenode.rpc-address.qy-ieg-4-v2.nn2"] = "qy-ieg-4-v2-nn-2.tencent-distribute.com:9000"
	conf["dfs.namenode.http-address.qy-ieg-4-v2.nn2"] = "qy-ieg-4-v2-nn-2.tencent-distribute.com:8080"
	conf["dfs.namenode.rpc-address.qy-ieg-4-v2.nn3"] = "qy-ieg-4-v2-nn-3.tencent-distribute.com:9000"
	conf["dfs.namenode.http-address.qy-ieg-4-v2.nn3"] = "qy-ieg-4-v2-nn-3.tencent-distribute.com:8080"
	option := hdfs.ClientOptionsFromConf(conf)
	option.Addresses = []string{address}
	u, _ := user.Current()
	option.User = u.Username
	//client, err := hdfs.NewClient(option)
	//if err != nil {
	//	fmt.Println("err create new client")
	//	fmt.Println(err.Error())
	//	return
	//}
	//defer client.Close()

	client, _ := hdfs.New("qy-ieg-4-v2-nn-1.tencent-distribute.com:9000")

	dir, err := client.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, fileInfo := range dir {
		func() {
			p := filepath.Join(path, fileInfo.Name())
			file, err := client.Open(p)
			scanner := bufio.NewScanner(file)
			// optionally, resize scanner's capacity for lines over 64K, see next example
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			defer file.Close()
		}()
	}
}
