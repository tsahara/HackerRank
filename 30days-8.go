package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    s := bufio.NewScanner(os.Stdin)

    s.Scan()
    n, _ := strconv.ParseUint(s.Text(), 10, 64)

    phoneBook := make(map[string]string)

    for i:= uint64(0); i < n; i++ {
        s.Scan()
        arr := strings.Split(s.Text(), " ")
        phoneBook[arr[0]] = arr[1]
    }

    for s.Scan() {
        name := s.Text()
        phoneNumber, found := phoneBook[name]
        if found {
            fmt.Println(name + "=" + phoneNumber)
        } else {
            fmt.Println("Not found")
        }
    }
}
