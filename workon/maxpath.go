/*
可以引⼊的库和版本相关请参考 “环境说明”
Please refer to the "Environmental Notes" for the libraries and versions that can be introduced.
*/

// 在一个有向图中，每个节点用一个大写字母表示。现在将一条路径的值定义为该路径上最常出现的字母数。例如， 如果一条路径为“ABACA”，因为“A”是这条路径上出现最多次数的字母，所以这条路径的的值就为3 。
// 给你一个有 n 个节点和 m 条有向边组成的图，返回这个图上所有路径里，值最大的那条路径的值。如果最大值为无限大（有环），则返回 -1

// 图用字符串和邻接列表表示。第 i 个字符代表第 i 个节点的大写字母。边列表中的每个元组（i, j）表示有一条从第 i 个节点到第 j 个节点的有向边。一条边也可以从由当前顶点指向自身。
// case 1:
// input:
// ABACA
// [(0, 1),
//  (0, 2),
//  (2, 3),
//  (3, 4)]
// output:
// 3
// Explanation: 这个图最数值最大的路径为[0, 2, 3, 4]（A， A， C，A）

// case 2:
// input:
// A
// [(0, 0)]
// output:
// None

package main

import "fmt"

const (
	white = 0
	grey  = 1
	black = 2
)

// TODO: cnt?
func dfs(input string, cnt map[byte]int, color map[int]int, g map[int][]int, start int) int {
	fmt.Printf("start at %d, %c\n", start, input[start])
	if color[start] == grey {
		fmt.Printf("loop at %v\n", start)
		return -1
	}

	color[start] = grey
	cnt[input[start]]++
	defer func() {
		cnt[input[start]]--
	}()
	var res = -1
	allLoop := false
	for _, v := range g[start] {
		allLoop = true
		//if color[v] != black {
		cur := dfs(input, cnt, color, g, v)
		// start -> v will have a loop
		if cur == -1 {
			//fmt.Printf("%d -> %d, loop\n", start, v)
			continue
		}
		allLoop = false
		res = max(res, cur)
		fmt.Printf("%d->%d: res:%v\n", start, v, res)

	}

	color[start] = black
	if allLoop {
		fmt.Printf("all loop: %v\n", start)
		return -1
	}
	for _, c := range cnt {
		res = max(res, c)
	}
	fmt.Printf("to black: %v, cnt: %v, res: %v\n", start, cnt, res)

	// too complicated?
	// for nv := range g {
	// 	if color[nv] != black {
	// 		cnt1 := make(map[byte]int)
	// 		x := dfs(input, cnt1, color, g, nv)
	// 		if x > res {
	// 			res = x
	// 		}
	// 	}
	// }
	return res
}

func MaxPath(s string, edge [][]int) int {
	g := make(map[int][]int)
	for i := 0; i < len(edge); i++ {
		g[edge[i][0]] = append(g[edge[i][0]], edge[i][1])
	}
	var res int = -1
	for start, _ := range g {
		fmt.Printf("START %d\n", start)
		cnt := make(map[byte]int)
		color := make(map[int]int)
		cur := dfs(s, cnt, color, g, start)
		//fmt.Printf("start at %d, get cnt: %v, color: %v\n", start, cnt, color)
		if cur > res {
			res = cur
		}

	}
	return res
}

// {0, 1} {2, 3} {3, 4}
func main() {
	// fmt.Println("Talk is cheap. Show me the code.")
	// fmt.Println("result", MaxPath("ABACA", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}))         // 3
	fmt.Println("result", MaxPath("ABACA", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {4, 0}})) // 3
	// fmt.Println("result", MaxPath("ABACA", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {4, 1}})) // 3
	// fmt.Println("result", MaxPath("A", [][]int{{0, 0}}))                                     // -1
	// fmt.Println("result", MaxPath("AAAAA", [][]int{{0, 1}, {2, 3}, {3, 4}}))                 // 3
}
