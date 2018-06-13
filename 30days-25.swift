import Foundation

// Enter your code here

func is_prime(_ num: Int) -> Bool {
    let sqrt = Int(Double(num).squareRoot())
    if num == 1 {
        return false
    } else if num < 4 {
        return true
    }
    for i in 2...sqrt {
        if num % i == 0 {
            return false
        }
    }
    return true
}
let T = Int(readLine()!)!
for _ in 0..<T {
    let n = Int(readLine()!)!
    if is_prime(n) {
        print("Prime")
    } else {
        print("Not prime")
    }
}
