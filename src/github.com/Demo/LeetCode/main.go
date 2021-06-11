package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main(){
	Select(1)
}

func Select(i int) {
	result := ""
	switch i {
	case 1://数组中两元素的最大乘积
		result = strconv.Itoa(maxProduct([]int{3,4,5,2}))
		break
	case 2://二叉搜索树中的搜索
		searchBST(new(TreeNode),3)
		break
	case 3://二叉树的最大深度
		maxDepth(new(TreeNode))
	}
	fmt.Print(result)
}

//数组中两元素的最大乘积
func maxProduct(nums []int) int {
	max1, max2 := 1, 1
	for _, value := range nums {
		if value > max1 {
			max1, max2 = value, max1
		} else if value > max2 {
			max2 = value
		}
	}
	return (max1 - 1) * (max2 - 1)
}
//二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
//二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a ,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}