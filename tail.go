package main

import (
	"os"
	"fmt"
	"strings"
	"time"
)

func main() {
	var sRead string
	sOpenFile := os.Args[1]

	fOpenFile, err := os.Open(sOpenFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fOpenFile.Close()

	//fOpenFile.Seek(0, os.SEEK_END)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("start")

	for ;; {
		aFile := make([]byte, 1)
		_, err = fOpenFile.Read(aFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.Index(string(aFile[0]),"\n" ) >= 0{
			fmt.Println(sRead)
		}else{
			sRead += string(aFile[0])
		}
		time.Sleep(100)
	}

}
