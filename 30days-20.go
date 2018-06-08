package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    line1, _, _ := reader.ReadLine()
    n_str := strings.TrimRight(string(line1), "\r\n")
    n, _ := strconv.ParseUint(n_str, 10, 64)

    line2, _, _ := reader.ReadLine()
    arrStr := strings.Split(string(line2), " ")

    var arr []uint64
    for _, s := range arrStr {
        i, _ := strconv.ParseUint(s, 10, 64)
        arr = append(arr, i)
    }

    numSwap := 0
    for i := uint64(0); i < n; i++ {
        numberOfSwaps := 0
        for j := uint64(0); j < n - 1; j++ {
            if arr[j] > arr[j + 1] {
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                numberOfSwaps++
            }
        }
        numSwap += numberOfSwaps
        if numberOfSwaps == 0 {
            break
        }
    }

    fmt.Printf("Array is sorted in %d swaps.\n", numSwap)
    fmt.Printf("First Element: %d\n", arr[0])
    fmt.Printf("Last Element: %d\n", arr[len(arr)-1])
}
