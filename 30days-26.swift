import Foundation

// Enter your code here
let actual = readLine()!.components(separatedBy: " ").map{ Int($0)! }
let (aday, amonth, ayear) = (actual[0], actual[1], actual[2])

let expected = readLine()!.components(separatedBy: " ").map{ Int($0)! }
let (eday, emonth, eyear) = (expected[0], expected[1], expected[2])

var fine: Int
if ayear < eyear {
    fine = 0
} else if ayear > eyear {
    fine = 10000
} else {
    if amonth < emonth {
        fine = 0
    } else if amonth > emonth {
        fine = 500 * (amonth - emonth)
    } else {
        if aday <= eday {
            fine = 0
        } else {
            fine = 15  * (aday - eday)
        }
    }
}
print(fine)
