package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    count := 0
    maxCount := 0
    for i := 0; i < 32; i++ {
        if n & 1 == 1 {
            count += 1
            if count > maxCount {
                maxCount = count
            }
        } else {
            count = 0
        }
        n >>= 1
    }
    fmt.Printf("%d\n", maxCount)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
