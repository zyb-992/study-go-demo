package main

import (
	"fmt"
	"log"
	"time"
)

const pathPrefix = "strings_"
const pathSuffix = ".log"

func main() {
	common()
}

func getDate() string {
	return time.Now().Format("2006-01-02")
}

func readAndWrite() {

	common()
	path := fmt.Sprintf("%s%s%s", pathPrefix, getDate(), pathSuffix)
	f, err := createFile(path)
	if err != nil {
		log.Fatalf("create err:%v", err)
		return
	}
	defer f.Close()

	err = write2File(f, "")
	if err != nil {
		log.Fatalf("write err: %v", err)
		return
	}

	//f, err := os.Open(path)
	//if err != nil {
	//	log.Fatalf("open file err:%v", err)
	//}

	// after write
	data, err := readFrom(f)
	if err != nil {
		log.Fatalf("read err:%v", err)
		return
	}
	fmt.Printf("[file]:[%s], [msg]:%s\n", path, data)
}
