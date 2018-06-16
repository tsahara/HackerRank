import Foundation

guard let N = Int((readLine()?.trimmingCharacters(in: .whitespacesAndNewlines))!)
    else { fatalError("Bad input") }

var emails: [String] = []
for _ in 1...N {
    guard let firstNameEmailIDTemp = readLine() else { fatalError("Bad input") }
    let firstNameEmailID = firstNameEmailIDTemp.split(separator: " ").map{ String($0) }

    let firstName = firstNameEmailID[0]

    let emailID = firstNameEmailID[1]
    let regex = try? NSRegularExpression(pattern: "@gmail.com$")
    if regex?.firstMatch(in: emailID, range: NSRange(location: 0, length: emailID.count)) != nil {
        emails.append(firstName)
    }
}
emails.sort(by: <)
for name in emails {
    print(name)
}
