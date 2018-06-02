import Foundation

class Difference {
    private var elements = [Int]()
    var maximumDifference: Int

    // Write your code here
    init(a: [Int]) {
        for e in a {
            self.elements.append(e)
        }
        self.maximumDifference = 0
    }

    public func computeDifference() {
        for a in self.elements {
            for b in self.elements {
                let z = abs(a - b)
                if self.maximumDifference < z {
                    self.maximumDifference = z
                }
            }
        }
    }
} // End of Difference class

let n = Int(readLine()!)!
let a = readLine()!.components(separatedBy: " ").map{ Int($0)! }

let d = Difference(a: a)

d.computeDifference()

print(d.maximumDifference)
