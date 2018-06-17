import Foundation

guard let t = Int((readLine()?.trimmingCharacters(in: .whitespacesAndNewlines))!)
    else { fatalError("Bad input") }

for _ in 1...t {
    guard let nkTemp = readLine() else { fatalError("Bad input") }
    let nk = nkTemp.split(separator: " ").map{ String($0) }

    guard let n = Int(nk[0].trimmingCharacters(in: .whitespacesAndNewlines))
        else { fatalError("Bad input") }

    guard let k = Int(nk[1].trimmingCharacters(in: .whitespacesAndNewlines))
        else { fatalError("Bad input") }

    var m = 0
    for a in 1...n-1 {
        for b in a+1...n {
            let c = a & b
            if c > m && c < k {
                m = c
            }
        }
    }
    print(m)
}
