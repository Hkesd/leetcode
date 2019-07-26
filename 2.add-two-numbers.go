package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 由于链表中存储的数据本身就是逆序的，故主要考察进位处理
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//rl1 := reverseList(l1)
	//rl2 := reverseList(l2)
	rl1 := l1
	rl2 := l2

	result := ListNode{0, nil}
	insert := &result

	carry := 0

	for rl1 != nil || rl2 != nil || carry != 0 {
		tmpV1 := 0
		tmpV2 := 0

		if rl1 != nil {
			tmpV1 = rl1.Val
			rl1 = rl1.Next
		}
		if rl2 != nil {
			tmpV2 = rl2.Val
			rl2 = rl2.Next
		}

		tmp := tmpV1 + tmpV2 + carry

		if tmp >= 10 {
			carry = tmp / 10
			tmp = tmp % 10
		} else {
			carry = 0
		}

		tmpListNode := ListNode{tmp, nil}
		insert.Next = &tmpListNode
		insert = &tmpListNode
	}

	return result.Next
}

func reverseList(l *ListNode) *ListNode {
	var p, q, r *ListNode
	p = l

	q = p.Next
	for q != nil {
		r = q.Next
		q.Next = p
		p = q
		q = r
	}

	l.Next = nil

	return p
}

//func main() {
//	l1 := ListNode{5, nil}
//	//l2 := ListNode{4, nil}
//	//l3 := ListNode{3, nil}
//
//	l4 := ListNode{5, nil}
//	//l5 := ListNode{6, nil}
//	//l6 := ListNode{4, nil}
//
//
//	l1.Next = nil
//	//l1.Next = &l2
//	//l2.Next = &l3
//	//l3.Next = nil
//
//	l4.Next = nil
//	//l4.Next = &l5
//	//l5.Next = nil
//	//l5.Next = &l6
//	//l6.Next = nil
//
//	re := addTwoNumbers(&l1, &l4)
//
//	for re != nil {
//		fmt.Println(re.Val)
//
//		re = re.Next
//	}
//}
