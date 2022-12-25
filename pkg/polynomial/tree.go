package polynomial

type Node struct {
	Value *int
	Next  *Node
	Prev  *Node
}

// func (n *Node) Insert(value *int) {
// 	if value < n.Value {
// 		if n.Left == nil {
// 			n.Left = &Node{Value: value}
// 		} else {
// 			n.Left.Insert(value)
// 		}
// 	} else {
// 		if n.Right == nil {
// 			n.Right = &Node{Value: value}
// 		} else {
// 			n.Right.Insert(value)
// 		}
// 	}
// }
