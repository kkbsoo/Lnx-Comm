package main

import(
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var sComm string
    var sCat string
    var sFileRead string
    iReadCount := 0

    if len(os.Args) == 3{
        sCat := string(os.Args[1])
        sFile := string(os.Args[2])

        //파일 오픈
        fOpenFile, err := os.Open(sFile)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer fOpenFile.Close()

        //파일 정보
        fFileStat, err := fOpenFile.Stat()
        if err != nil {
            fmt.Println(err)
            return
        }
        iFileSize := int(fFileStat.Size())

        aFile := make([]byte, iFileSize)
        _, err = fOpenFile.Read(aFile)
        if err != nil {
            fmt.Println(err)
            return
        }
        for iReadCount=0; iReadCount < iFileSize; iReadCount++ {
            if string(aFile[iReadCount]) == "\n" {
                if strings.Index(sFileRead, sCat) >= 0 {
                    fmt.Println(sFileRead)
                    sFileRead = ""
                } else {
                    sFileRead = ""
                }
            }else{
                sFileRead += string(aFile[iReadCount])
            }
        }
    }else{
        scanner := bufio.NewScanner(os.Stdin)

        for scanner.Scan() {
            sComm = scanner.Text()
            sCat = string(os.Args[1])

            if strings.Index(sComm,sCat) >= 0{
                fmt.Println(sComm)
            }
        }

        if err := scanner.Err(); err != nil {
            fmt.Println(os.Stderr, "error:", err)
            os.Exit(1)
        }
    }
}