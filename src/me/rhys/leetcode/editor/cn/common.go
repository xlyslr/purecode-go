package leetcode

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) GoString() string {
	var b strings.Builder
	b.WriteString("[")
	node := l
	for node != nil {
		b.WriteString(strconv.Itoa(node.Val))
		b.WriteString(",")
		node = node.Next
	}
	b.WriteString("]")
	return b.String()
}
