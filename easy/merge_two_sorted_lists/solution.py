class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __str__(self):
        return f"{self.val} - {self.next}"

class Solution:
    def mergeTwoLists(self, list1: ListNode, list2: ListNode) -> ListNode:
            dummy = ListNode()
            tail = dummy

            while list1 != None and list2 != None:
                if list1.val < list2.val:
                    tail.next = list1
                    list1 = list1.next
                else:
                    tail.next = list2
                    list2 = list2.next
                tail = tail.next

            if list1 == None:
                tail.next = list2
            elif list2 == None:
                tail.next = list1

            return dummy.next
            




answer = Solution()

# list2
ListTwoNode3 = ListNode(val=4)
ListTwoNode2 = ListNode(val=3, next=ListTwoNode3)
ListTwoNode1 = ListNode(val=1, next=ListTwoNode2)

# list1
ListOneNode3 = ListNode(val=4)
ListOneNode2 = ListNode(val=2, next=ListOneNode3)
ListOneNode1 = ListNode(val=1, next=ListOneNode2)

answer.mergeTwoLists(ListOneNode1, ListTwoNode1)
print(ListOneNode1)