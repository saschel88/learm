package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	pathName := "./testdata"
	//dirName:=""
	files, err := ioutil.ReadDir(pathName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	//
} // the end main
