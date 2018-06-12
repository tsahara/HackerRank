import Foundation

// Start of class Node
class Node {
    var data: Int
    var next: Node?

    init(d: Int) {
        data = d
    }
} // End of class Node

// Start of class LinkedList
class LinkedList {
    func insert(head: Node?, data: Int) -> Node? {
        if head == nil {
            return Node(d: data)
        }

        head?.next = insert(head: head?.next, data: data)

        return head
    }

    func display(head: Node?) {
        if head != nil {
            print(head!.data, terminator: " ")

            display(head: head?.next)
        }
    }

    // Start of function removeDuplicates
    func removeDuplicates(head: Node?) -> Node? {
        // Add your code here
        var node = head
        while node != nil && node!.next != nil {
            if node!.data == node!.next!.data {
                node!.next = node!.next!.next
            } else {
                node = node!.next
            }
        }
        return head
    } // End of function removeDuplicates
} // End of class LinkedList
var head: Node?
let linkedList = LinkedList()

let t = Int(readLine()!)!

for _ in 0..<t {
    head = linkedList.insert(head: head, data: Int(readLine()!)!)
}

linkedList.display(head: linkedList.removeDuplicates(head: head))
