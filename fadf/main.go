package main

import (
	"fmt"
	"sort"
)

func main() {
	//integer := largestGoodInteger("6777133339")
	//integer := largestGoodInteger("2300019")
	//integer := largestGoodInteger("74444")
	//integer := largestGoodInteger("232220002")
	//integer := largestGoodInteger("42352338")
	//fmt.Println(integer)

	//root := &TreeNode{
	//	Val:   1,
	//	Left:  nil,
	//	Right: nil,
	//}

	//root := &TreeNode{
	//	Val: 0,
	//	Left: &TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: nil,
	//}

	//root := &TreeNode{
	//	Val: 1,
	//Left: &TreeNode{
	//	Val:   0,
	//	Left:  nil,
	//	Right: nil,
	//},
	//	Right: &TreeNode{
	//		Val:  3,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:  1,
	//			Left: nil,
	//			Right: &TreeNode{
	//				Val:   3,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//	},
	//}

	//root := &TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 8,
	//		Left: &TreeNode{
	//			Val:   0,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   1,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:  5,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	//subtree := averageOfSubtree(root)
	//fmt.Println(subtree)

	//combination := largestCombination([]int{16, 17, 71, 62, 12, 24, 14})
	//combination := largestCombination([]int{8, 8})
	//fmt.Println(combination)

	//texts := countTexts("22233")
	//texts := countTexts("222222222222222222222222222222222222")
	texts := countTexts("444479999555588866")
	fmt.Println(texts)
}
func largestGoodInteger(num string) string {
	m := make(map[byte]int)
	for i := 0; i < len(num); i++ {
		if i == 0 {
			m[num[i]]++
		} else {
			if num[i-1] != num[i] {
				if m[num[i]] >= 3 {

				} else {
					m[num[i]] = 1
				}
			} else {
				m[num[i]]++
			}
		}
	}
	res := make([]int, 0, 0)
	for key, v := range m {
		if v >= 3 {
			res = append(res, int(key))
		}
	}
	sort.Ints(res)

	ans := ""
	if len(res) == 0 {
		return ans
	}
	for i := 0; i < 3; i++ {
		ans += string(res[len(res)-1])
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	ans := 0
	res := make(map[*TreeNode]int)
	m := make(map[int][]*TreeNode)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		m[root.Val] = append(m[root.Val], root)
		travel(root.Right)
	}

	var avg func(root *TreeNode) (sum int, num int)
	avg = func(root *TreeNode) (sum int, num int) {
		if root == nil {
			return 0, 0
		}

		leftsum, leftnum := avg(root.Left)
		rightsum, rightnum := avg(root.Right)
		totalsum := leftsum + rightsum + root.Val
		totalnum := leftnum + rightnum + 1
		if totalnum != 0 {
			average := totalsum / totalnum
			if average == root.Val {
				ans++
			}
		}
		return totalsum, totalnum
	}
	avg(root)
	return ans
	return len(res)
}

func largestCombination(candidates []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		temp := 1
		ans := 0

		for j := 0; j < len(candidates); j++ {
			if (candidates[j]>>i)&temp == 1 {
				ans++
			}

		}
		if ans > res {
			res = ans
		}
	}
	return res
}

func countTexts1(pressedKeys string) int {
	var travel func(s string) int
	travel = func(s string) int {
		if len(s) == 0 {
			return 1
		}
		a, b, c, d := 0, 0, 0, 0
		if s[0]-'a' == 7 || s[0]-'a' == 9 {
			if len(s) >= 1 {
				a = travel(s[1:])
			}
			if len(s) >= 2 && s[0] == s[1] {
				b = travel(s[2:])
			}
			if len(s) >= 3 && s[0] == s[1] && s[1] == s[2] {
				c = travel(s[3:])
			}
			if len(s) >= 4 && s[0] == s[1] && s[1] == s[2] && s[2] == s[3] {
				travel(s[4:])
			}
		} else {
			if len(s) >= 1 {
				a = travel(s[1:])
			}
			if len(s) >= 2 && s[0] == s[1] {
				b = travel(s[2:])
			}
			if len(s) >= 3 && s[0] == s[1] && s[1] == s[2] {
				c = travel(s[3:])
			}
		}
		return a + b + c + d
	}
	return travel(pressedKeys)
}

func countTexts(pressedKeys string) int {
	m := make(map[int]int)
	var travel func(i int) int
	travel = func(i int) int {
		if len(pressedKeys) == i {
			return 1
		}
		a, b, c, d := 0, 0, 0, 0
		if pressedKeys[i]-'0' == 7 || pressedKeys[i]-'0' == 9 {
			if len(pressedKeys) >= i+1 {
				if v, ok := m[i+1]; ok {
					a = v
				} else {
					a = travel(i + 1)
					m[i+1] = a
				}
			}
			if len(pressedKeys) >= i+2 && pressedKeys[i] == pressedKeys[i+1] {
				if v, ok := m[i+2]; ok {
					b = v
				} else {
					b = travel(i + 2)
					m[i+2] = b
				}
			}
			if len(pressedKeys) >= i+3 && pressedKeys[i] == pressedKeys[i+1] && pressedKeys[i+1] == pressedKeys[i+2] {
				if v, ok := m[i+3]; ok {
					c = v
				} else {
					c = travel(i + 3)
					m[i+3] = c
				}
			}
			if len(pressedKeys) >= i+4 && pressedKeys[i] == pressedKeys[i+1] && pressedKeys[i+1] == pressedKeys[i+2] && pressedKeys[i+2] == pressedKeys[i+3] {
				if v, ok := m[i+4]; ok {
					d = v
				} else {
					d = travel(i + 4)
					m[i+4] = d
				}
			}
		} else {
			if len(pressedKeys) >= i+1 {
				if v, ok := m[i+1]; ok {
					a = v
				} else {
					a = travel(i + 1)
					m[i+1] = a
				}
			}
			if len(pressedKeys) >= i+2 && pressedKeys[i] == pressedKeys[i+1] {
				if v, ok := m[i+2]; ok {
					b = v
				} else {
					b = travel(i + 2)
					m[i+2] = b
				}
			}
			if len(pressedKeys) >= i+3 && pressedKeys[i] == pressedKeys[i+1] && pressedKeys[i+1] == pressedKeys[i+2] {
				if v, ok := m[i+3]; ok {
					c = v
				} else {
					c = travel(i + 3)
					m[i+3] = c
				}
			}
		}
		return (a + b + c + d) % (1e9 + 7)
	}
	return travel(0)
}
