package main

import (
	"encoding/json"
	"flag"
	"github.com/zyb-992/demo/tools/type2json/cmd"
	"log"
	"os"
)

func init() {
	path, _ := os.Getwd()
	log.Println("your current path: ", path)
}

func main() {
	if flag.Parse(); !flag.Parsed() {
		log.Println("parse flag failed")
		os.Exit(1)
	}

	if cmd.TypeNameFlag == nil {
		log.Println("the -t flag not be set")
		os.Exit(1)
	}

	parse := basic(*cmd.DirFlag, *cmd.TypeNameFlag)
	if parse == nil {
		log.Printf("can't find the type, dir:%v, type name:%v", *cmd.DirFlag, *cmd.TypeNameFlag)
	}

	var object = parse.object
	if object == nil {
		log.Println("couldn't acquire the type by your input typename")
		os.Exit(1)
	}

	// 3. get fields mapping
	mapping := iterator(fields(object))

	// 4. generate json string
	data, err := json.MarshalIndent(mapping, "", "    ")
	if err != nil {
		log.Println("exec json.Marshal failed, err:", err)
		return
	}

	log.Println("result: \n", string(data))
}
