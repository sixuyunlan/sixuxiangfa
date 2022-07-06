package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序遍历：先访问根节点，再访问左子树，最后访问右子树
func (h *Hero) PreOrder() { //前序遍历
	if h == nil { //如果节点为空，则返回
		return
	}
	fmt.Println(h.No, h.Name) //访问根节点
	h.Left.PreOrder()         //访问左子树
	h.Right.PreOrder()        //访问右子树
}

//中序遍历：先访问左子树，再访问根节点，最后访问右子树
func (h *Hero) InOrder() { //中序遍历
	if h == nil { //如果节点为空，则返回
		return
	}
	h.Left.InOrder()          //访问左子树
	fmt.Println(h.No, h.Name) //访问根节点
	h.Right.InOrder()         //访问右子树
}

//后序遍历：先访问左子树，再访问右子树，最后访问根 节点
func (h *Hero) PostOrder() { //后序遍历
	if h == nil { //如果节点为空，则返回
		return
	}
	h.Left.PostOrder()        //访问左子树
	h.Right.PostOrder()       //访问右子树
	fmt.Println(h.No, h.Name) //访问根节点
}
func main() {
	root := &Hero{
		1,
		"宋江",
		nil,
		nil,
	}
	left1 := &Hero{
		2,
		"吴用",
		nil,
		nil,
	}
	node10 := &Hero{
		10,
		"tom",
		nil,
		nil,
	}
	node12 := &Hero{
		12,
		"jack",
		nil,
		nil,
	}
	left1.Left = node10
	left1.Right = node12

	right1 := &Hero{
		3,
		"卢俊义",
		nil,
		nil,
	}

	root.Left = left1   //设置左子树
	root.Right = right1 //设置右子树

	right2 := &Hero{
		4,
		"林冲",
		nil,
		nil,
	}
	right1.Right = right2 //设置右子树

	root.PreOrder() //前序遍历
	fmt.Println()
	root.InOrder() //中序遍历
	fmt.Println()
	root.PostOrder() //后序遍历

}
