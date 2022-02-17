package main

import (
	"io"
	"log"
	"os"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	log.Fatal("error read file : ", err)
	return p, err
}

func main() {
	fp, err := os.Open("ch01/empty")
	if err != nil {
		log.Fatal("error open file!", err)
		return
	}
	data, err := ReadFrom(fp, 10)
	if err != nil {
		log.Fatal("error read file : ", err)
		return
	}
	log.Println(data)
}
