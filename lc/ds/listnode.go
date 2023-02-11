package ds

import (
	"strconv"
	"strings"
)

func (n *ListNode) String() string {
	builder := strings.Builder{}
	node := n

	for node != nil {
		builder.WriteString("-")
		builder.WriteString(strconv.Itoa(node.Val))
		node = node.Next
	}
	return builder.String()
}
