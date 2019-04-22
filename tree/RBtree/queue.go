package RBtree


//type Queue struct {
//	First *queue
//	Root  *queue
//	Count int
//}
//type queue struct {
//	Next *queue
//	node *Node
//}
//
//func NewQueue() *Queue {
//	return &Queue{
//		nil, nil, 0}
//}
//
//func Newqueue(node *Node) *queue {
//	return &queue{node: node, Next: nil}
//}
//func (q *Queue) Push(node *Node) {
//	q.Add(node)
//}
//func (q *Queue) Add(node *Node) bool {
//	if q.Root == nil {
//		q.Root = Newqueue(node)
//		q.First = q.Root
//		q.Count++
//		return true
//	}
//	current := q.Root
//	for {
//		if current.Next != nil {
//			current = current.Next
//		} else {
//			current.Next = Newqueue(node)
//			q.Count++
//			return true
//		}
//	}
//}
//
//func (q *Queue) Pop() *Node {
//	queue := q.First
//	q.Root = queue.Next
//	q.First = queue.Next
//	q.Count--
//	return queue.node
//}
