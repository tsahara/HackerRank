package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    line1 := scanner.Text()
    T, _ := strconv.ParseUint(line1, 10, 32)

    for j := 0; j < int(T); j++ {
        scanner.Scan()
        line := scanner.Text()
        S := strings.TrimRight(string(line), "\r\n")

        // even
        for i := 0; i < len(S); i += 2 {
            fmt.Printf("%c", S[i])
        }
        fmt.Printf(" ")
        // odd
        for i := 1; i < len(S); i += 2 {
            fmt.Printf("%c", S[i])
        }
        fmt.Printf("\n")
    }
}
