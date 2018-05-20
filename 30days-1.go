package main


import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    var _ = strconv.Itoa

    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)

    // Declare second integer, double, and String variables.
    var i2 uint64
    var d2 float64
    var s2 string

    // Read and save an integer, double, and String to your variables.
    scanner.Scan()
    line1 := scanner.Text()
    i2, _ = strconv.ParseUint(line1, 10, 64)

    scanner.Scan()
    line2 := scanner.Text()
    d2, _ = strconv.ParseFloat(line2, 64)

    scanner.Scan()
    s2 = scanner.Text()

    // Print the sum of both integer variables on a new line.
    fmt.Println(i + i2)

    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f\n", d + d2)

    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Println(s + s2)
}
