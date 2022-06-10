
function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

var mergeTwoLists = function(list1, list2) {
    let dummy = new ListNode()
    let tail = dummy

    while (list1 != null && list2 != null){
        if (list1.val < list2.val){
            tail.next = list1
            list1 = list1.next
        } else {
            tail.next = list2
            list2 = list2.next
        }
        tail = tail.next
    }
    if (list1 == null){
        tail.next = list2
    } else if(list2 == null){
        tail.next = list1
    }

    return dummy.next
};

// list2
const ListTwoNode3 = ListNode(val=4)
const ListTwoNode2 = ListNode(val=3, next=ListTwoNode3)
const ListTwoNode1 = ListNode(val=1, next=ListTwoNode2)

// list1
const ListOneNode3 = ListNode(val=4)
const ListOneNode2 = ListNode(val=2, next=ListOneNode3)
const ListOneNode1 = ListNode(val=1, next=ListOneNode2)

mergeTwoLists(ListOneNode1, ListTwoNode1)