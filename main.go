package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func dirTree(output io.Writer, currDir string, printFiles bool) error {
	recursionPrintService("", output, currDir, printFiles)
	return nil
}

func recursionPrintService(prependingString string, output io.Writer, currDir string, printFiles bool) {
	fileObj, err := os.Open(currDir)
	defer fileObj.Close()
	if err != nil {
		log.Fatalf("Could not open %s: %s", currDir, err.Error())
	}
	fileName := fileObj.Name()
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatalf("Could not read dir names in %s: %s", currDir, err.Error())
	}
	var filesMap map[string]os.FileInfo = map[string]os.FileInfo{}
	var unSortedFilesNameArr []string = []string{}
	for _, file := range files {
		unSortedFilesNameArr = append(unSortedFilesNameArr, file.Name())
		filesMap[file.Name()] = file
	}
	sort.Strings(unSortedFilesNameArr)
	var sortedFilesArr []os.FileInfo = []os.FileInfo{}
	for _, stringName := range unSortedFilesNameArr {
		sortedFilesArr = append(sortedFilesArr, filesMap[stringName])
	}
	files = sortedFilesArr
	var newFileList []os.FileInfo = []os.FileInfo{}
	var length int
	if !printFiles {
		for _, file := range files {
			if file.IsDir() {
				newFileList = append(newFileList, file)
			}
		}
		files = newFileList
	}
	length = len(files)
	for i, file := range files {
		if file.IsDir() {
			var stringPrepender string
			if length > i+1 {
				fmt.Fprintf(output, prependingString+"├───"+"%s\n", file.Name())
				stringPrepender = prependingString + "│\t"
			} else {
				fmt.Fprintf(output, prependingString+"└───"+"%s\n", file.Name())
				stringPrepender = prependingString + "\t"
			}
			newDir := filepath.Join(currDir, file.Name())
			recursionPrintService(stringPrepender, output, newDir, printFiles)
		} else if printFiles {
			if file.Size() > 0 {
				if length > i+1 {
					fmt.Fprintf(output, prependingString+"├───%s (%vb)\n", file.Name(), file.Size())
				} else {
					fmt.Fprintf(output, prependingString+"└───%s (%vb)\n", file.Name(), file.Size())
				}
			} else {
				if length > i+1 {
					fmt.Fprintf(output, prependingString+"├───%s (empty)\n", file.Name())
				} else {
					fmt.Fprintf(output, prependingString+"└───%s (empty)\n", file.Name())
				}
			}
		}
	}
}

func main() {
	/*	pathName := "./src/testdata"
		//dirName:=""
		files, err := ioutil.ReadDir(pathName)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}
		// hello world
	*/

	/*
		arrayOfBet:=[]float32{ 0.00017143,0.00028571,0.00045714, 0.00057143}
		arraOfPrize:=[]float32{0.0003,0.0005,0.0008,0.001}
		arrayOfFee:=[]float32 {0.00000678,0.00000904}
		fmt.Println(arrayOfBet, arraOfPrize)
		fmt.Println(float32(2)*arrayOfBet[0],float32(2)*arrayOfBet[1],float32(2)*arrayOfBet[2])

		for i:=0;i<3 ;i++ {
			for j:=1;j<11;j++ {

				if (float32(j)*arraOfPrize[i]-float32(j)*arrayOfBet[i]-float32(j)*arrayOfFee[1]-arrayOfBet[3]-arrayOfFee[1]>0)&& (arraOfPrize[3]-float32(2)*arrayOfFee[1]-float32(j)*arrayOfBet[i]-float32(j)*arrayOfFee[1]-arrayOfBet[3]>0) {
					fmt.Println("Количество ставок=",j, "Ставка=", arrayOfBet[i], "delta=",float32(j)*arraOfPrize[i]-float32(j)*arrayOfBet[i]-arrayOfBet[3])
				}
			}


		}*/

} // the end main
