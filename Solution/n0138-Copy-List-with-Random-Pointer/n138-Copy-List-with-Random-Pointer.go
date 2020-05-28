package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	mp := map[*Node]*Node{}
	res := &Node{Val: head.Val}
	mp[head] = res

	for p, cp := head.Next, res; p != nil; p, cp = p.Next, cp.Next {
		cp.Next = &Node{Val: p.Val}
		mp[p] = cp.Next
	}

	for p, cp := head, res; p != nil; p, cp = p.Next, cp.Next {
		if p.Random != nil {
			cp.Random = mp[p.Random]
		}
	}
	return res
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	for p := head; p != nil; p = p.Next.Next {
		cp := &Node{Val: p.Val}
		cp.Next = p.Next
		p.Next = cp
	}

	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}

	dummy := &Node{}
	cp, p := dummy, head
	for p != nil {
		tmp := p.Next
		next := p.Next.Next
		cp.Next = tmp
		cp = tmp
		p.Next = next
		p = next
	}
	return dummy.Next
}
