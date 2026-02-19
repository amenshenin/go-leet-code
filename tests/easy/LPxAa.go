package easy

import "fmt"

type LPxAa struct {
	IsComplete int
	Name       string
	Descr      string
}

func LPxAaInit() *LPxAa {
	return &LPxAa{
		IsComplete: 0,
		Name:       "https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/description/",
		Descr: `Даны два бинарных дерева: original и cloned, а также ссылка на узел target в оригинальном дереве.

Дерево cloned является копией оригинального дерева.

Верните ссылку на тот же узел в дереве cloned.

Обратите внимание, что вам не разрешается изменять какое-либо из двух деревьев или узел target, и ответ должен быть ссылкой на узел в дереве cloned.`,
	}
}

func (f *LPxAa) Run() error {
	example1()

	return nil
}

func example1() {
	tree := &Tree{
		Root: &Node{
			Value: 7,
			Level: 0,
		},
	}
	//[7,4,3,null,null,6,19]
	var pair1l = 4
	var pair1r = 3
	tree.Insert(&pair1l, &pair1r, 1)

	tree.Insert(nil, nil, 2)

	var pair2rl = 6
	var pair2rr = 19
	tree.Insert(&pair2rl, &pair2rr, 2)

	fmt.Printf("%+v\n", tree)
}

type Node struct {
	Value int
	Level int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(left, right *int, level int) {
	fmt.Print(level)
	node := t.SearchNotFilledYet(level)
	fmt.Printf("%+v\n", node)
	if node != nil {
		node.InsertNode(left, right)
	}
}

func (t *Tree) SearchNotFilledYet(level int) *Node {
	if t.Root.Level == level {
		return t.Root
	}
	if t.Root.Level > level {
		return nil
	}
	if t.Root.Left != nil {
		left := t.Root.Left.SearchNotFilledYetNode(level)
		if left != nil {
			return left
		}
	}
	if t.Root.Right != nil {
		right := t.Root.Right.SearchNotFilledYetNode(level)
		if right != nil {
			return right
		}
	}
	return nil
}

func (n *Node) SearchNotFilledYetNode(level int) *Node {
	if n.Level == level {
		return n
	}
	if n.Level > level {
		return nil
	}
	if n.Left != nil {
		left := n.Left.SearchNotFilledYetNode(level)
		if left != nil {
			return left
		}
	}
	if n.Right != nil {
		right := n.Right.SearchNotFilledYetNode(level)
		if right != nil {
			return right
		}
	}
	return nil
}

func (n *Node) InsertNode(left, right *int) {
	if left != nil {
		n.Left = &Node{
			Value: *left,
			Level: n.Level + 1,
		}
	}
	if right != nil {
		n.Right = &Node{
			Value: *right,
			Level: n.Level + 1,
		}
	}
}
