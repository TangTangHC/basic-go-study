package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func Test_2(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
}

func Test_3(t *testing.T) {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}

func Test_4(t *testing.T) {
	//maxSlidingWindow([]int{3, 3, 3, 30, 3, 3, 3, 3}, 3)
	//search([]int{5, 7, 7, 8, 8, 10}, 8)

	s := "abcde"
	for _, v := range s {
		fmt.Println(v)
	}
	for _, v := range s[:] {
		fmt.Println(v)
	}
	for _, v := range s {
		fmt.Println(string(v))
	}
	str := "Hello, 世界！"

	for i, ch := range str {
		//fmt.Println(ch)
		fmt.Printf("%v %c \n", i, ch)
	}
	//
	//for i, v := range utf8.RuneCountInString(str) {
	//
	//}
}

func maxSlidingWindow(nums []int, k int) []int {
	windows := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(windows) > 0 && windows[len(windows)-1] < nums[i] {
			windows = windows[:len(windows)-1]
		}
		windows = append(windows, nums[i])

		if i >= k && windows[0] == nums[i-k] {
			windows = windows[1:]
		}

		if i >= k-1 {
			res = append(res, windows[0])
		}

	}
	return res
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			// 前面找
			tmp := mid
			i := 0
			for nums[tmp] == target {
				i++
				tmp--
			}
			tmp = mid + 1
			for nums[tmp] == target {
				i++
				tmp++
			}
			// 后面找
		} else if nums[mid] > target {
			high = high - 1
		} else if nums[mid] < target {
			low = low + 1
		}
	}
	return 0
}

// 1,2,3,4,5
// 0,1,2,3,4
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印
	var res [][]int
	queue := []*TreeNode{root}
	var i int
	res = append(res, []int{root.Val})
	for len(queue) > 0 {
		lastNode := queue[len(queue)-1]
		var tmp []int
		if i%2 == 0 {
			tmp = append(tmp, lastNode.Right.Val, lastNode.Left.Val)
		} else {
			tmp = append(tmp, lastNode.Left.Val, lastNode.Right.Val)
		}
		i++
		res = append(res, tmp)

	}

	return res
}

func exist(board [][]byte, word string) bool {
	var res bool
	h := len(board)
	l := len(board[0])
	track := make([][]int, h)
	for i := range track {
		track[i] = make([]int, l)
	}

	var backTrack func([][]int, int, int, int)
	backTrack = func(track [][]int, i, x, y int) {
		if res {
			return
		}
		if y < 0 || x < 0 || y == len(track) || x == len(track[0]) {
			return
		}
		if track[y][x] == 1 {
			return
		}
		if word[i] != board[y][x] {
			return
		}
		if word[i] == board[y][x] {
			track[y][x] = 1
			i++
			if i == len(word) {
				res = true
				return
			}
		}
		backTrack(track, i, x+1, y)
		backTrack(track, i, x-1, y)
		backTrack(track, i, x, y+1)
		backTrack(track, i, x, y-1)
		track[y][x] = 0
		if i > 0 {
			i--
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			backTrack(track, 0, j, i)
			if res {
				return res
			}
		}
	}
	// backTrack(track, 0, 0)

	return false
}

func Test_backTrack(t *testing.T) {
	//a := [][]byte{}
	//a = append(a, []byte{'A', 'B', 'C', 'E'})
	//a = append(a, []byte{'S', 'F', 'C', 'S'})
	//a = append(a, []byte{'A', 'D', 'E', 'E'})
	////b := exist(a, "ABCCED")
	////b := exist(a, "SEE")
	//b := exist(a, "SFD")
	//println(b)
	arr := []interface{}{-6, nil, -3, -6, 0, -6, -5, 4, nil, nil, nil, -1, 7}
	root := buildTree(arr)
	pathSum(root, -21)
}

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var backTrack func([]int, *TreeNode, int)
	backTrack = func(tmp []int, node *TreeNode, sum int) {
		if node == nil {
			return
		}
		// if sum > target {
		//     return
		// }
		if node.Left == nil && node.Right == nil {
			if sum+node.Val == target {
				tmp = append(tmp, node.Val)
				var ts []int
				copy(tmp, ts)
				res = append(res, ts)
			}
			return
		}
		tmp = append(tmp, node.Val)
		backTrack(tmp, node.Left, sum+node.Val)
		backTrack(tmp, node.Right, sum+node.Val)
		tmp = tmp[:len(tmp)-1]
	}
	backTrack([]int{}, root, 0)
	return res
}

func buildTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{Val: -6}
	queue := []*TreeNode{root}

	for i := 1; i < len(arr); i += 2 {
		node := queue[0]
		queue = queue[1:]

		if val, ok := arr[i].(int); ok {
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}

		if i+1 < len(arr) {
			if val, ok := arr[i+1].(int); ok {
				node.Right = &TreeNode{Val: val}
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

func Test_JSON(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	u := User{
		Name: "123",
		Age:  12,
	}

	m1, _ := json.Marshal(u)
	m2, _ := json.Marshal(&u)
	t.Log(m1)
	t.Log(m2)
}
