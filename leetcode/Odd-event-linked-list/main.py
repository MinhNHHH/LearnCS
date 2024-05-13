class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def oddEvenList(head):
    if not head or not head.next or not head.next.next:
        return head

    odd_tail = head
    even_head = head.next
    even_tail = even_head

    while even_tail and even_tail.next:
        odd_tail.next = even_tail.next
        odd_tail = odd_tail.next
        even_tail.next = odd_tail.next
        even_tail = even_tail.next

    odd_tail.next = even_head

    return head

# Example usage:
# Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
head.next.next.next = ListNode(4)
head.next.next.next.next = ListNode(5)

# Call the function to rearrange the list
head = oddEvenList(head)

# Output the result: 1 -> 3 -> 5 -> 2 -> 4
while head:
    print(head.val, end=" -> ")
    head = head.next
