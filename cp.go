package main

import (
	"fmt"
	"os"
)

func fnErr(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var sFileRead string
    iFileSize := 0
    iBufSize := 10240
    iBufMaxSize := 10240

    sOpenFile := os.Args[1]
    sCloseFile := os.Args[2]

    //파일 오픈
    fOpenFile, err := os.Open(sOpenFile)
    fnErr(err)
    defer fOpenFile.Close()

    fCloseFile, err := os.Create(sCloseFile)
    fnErr(err)
    defer fCloseFile.Close()

    //파일 정보
    fFileStat, err := fOpenFile.Stat()
    fnErr(err)
    iFileSize = int(fFileStat.Size())

    aFile := make([]byte, iBufMaxSize)

    //파일 정보 가져오기
    _, err = fOpenFile.Read(aFile)
    fnErr(err)
    _, err = fOpenFile.Seek(0, 0)   //중요 Seek 안 해주면 두번의 슬라이스 사용할수없음
    fnErr(err)

    i := 1
    for {
        if iFileSize < iBufMaxSize {
            aFile := make([]byte, iFileSize + 1)
            _, err = fOpenFile.Read(aFile)
            fnErr(err)
            sFileRead += string(aFile)
            fCloseFile.WriteString(sFileRead)
            //fmt.Println("##########")
            break;
        }

        sFileRead += string(aFile)
        fCloseFile.WriteString(sFileRead)

        _, err = fOpenFile.Read(aFile)
        fnErr(err)
        _, err = fOpenFile.Seek(int64(iBufSize), 0)
        fnErr(err)
        _, err = fOpenFile.Read(aFile)
        fnErr(err)

        iBufSize += iBufMaxSize

        if iBufSize > iFileSize {
            sFileRead = ""
            iBufSize = iFileSize - (iFileSize - iBufMaxSize * i )

            aFile = make([]byte, (iFileSize - iBufSize))
            _, err = fOpenFile.Seek(int64(iBufSize), 0)
            fnErr(err)
            _, err = fOpenFile.Read(aFile)
            fnErr(err)
            sFileRead += string(aFile)
            fCloseFile.WriteString(sFileRead)
            break
        }
        sFileRead = ""
        i++
        /*
        if (iFileSize / iBufSize) * 100  % 10 == 0 {
            fmt.Println("##")
        }
        */
    }
    fmt.Println(" Copy Complete")
}