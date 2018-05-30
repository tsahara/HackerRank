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

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    sumMax := int32(-100)
    for y := 0; y < 4; y++ {
        for x := 0; x < 4; x++ {
            sum := arr[y][x]
            sum += arr[y][x+1]
            sum += arr[y][x+2]
            sum += arr[y+1][x+1]
            sum += arr[y+2][x]
            sum += arr[y+2][x+1]
            sum += arr[y+2][x+2]
            if sum > sumMax {
                sumMax = sum
            }
        }
    }
    fmt.Printf("%d\n", sumMax)
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
