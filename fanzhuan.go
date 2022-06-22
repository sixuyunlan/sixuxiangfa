package main

type LN struct {
	Val  int
	Next *LN
}

func reverselist(head *LN) *LN {
	cur := head
	//(nil),1,2,3,4,5,(nil)
	var pre *LN = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre

}
