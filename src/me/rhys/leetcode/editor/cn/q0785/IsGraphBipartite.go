//There is an undirected graph with n nodes, where each node is numbered
//between 0 and n - 1. You are given a 2D array graph, where graph[u] is an array of
//nodes that node u is adjacent to. More formally, for each v in graph[u], there is
//an undirected edge between node u and node v. The graph has the following
//properties:
//
//
// There are no self-edges (graph[u] does not contain u).
// There are no parallel edges (graph[u] does not contain duplicate values).
// If v is in graph[u], then u is in graph[v] (the graph is undirected).
// The graph may not be connected, meaning there may be two nodes u and v such
//that there is no path between them.
//
//
// A graph is bipartite if the nodes can be partitioned into two independent
//sets A and B such that every edge in the graph connects a node in set A and a node
//in set B.
//
// Return true if and only if it is bipartite.
//
//
// Example 1:
//
//
//Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
//Output: false
//Explanation: There is no way to partition the nodes into two independent sets
//such that every edge connects a node in one and a node in the other.
//
// Example 2:
//
//
//Input: graph = [[1,3],[0,2],[1,3],[0,2]]
//Output: true
//Explanation: We can partition the nodes into two sets: {0, 2} and {1, 3}.
//
//
// Constraints:
//
//
// graph.length == n
// 1 <= n <= 100
// 0 <= graph[u].length < n
// 0 <= graph[u][i] <= n - 1
// graph[u] does not contain u.
// All the values of graph[u] are unique.
// If graph[u] contains v, then graph[v] contains u.
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 378 👎 0

package q0785

//leetcode submit region begin(Prohibit modification and deletion)
var res bool
var colored []int

func isBipartite(graph [][]int) bool {
	n := len(graph)
	res = true
	colored = make([]int, n)
	for i := 0; i < n; i++ {
		if colored[i] == 0 {
			traverse(graph, i, 1)
		}
	}

	return res
}

func traverse(graph [][]int, s int, color int) {
	if !res {
		return
	}
	colored[s] = color
	nextColor := 1
	if color == 1 {
		nextColor = 2
	}
	for _, ss := range graph[s] {
		if colored[ss] == 0 {
			traverse(graph, ss, nextColor)
		} else if colored[ss] != nextColor {
			res = false
			return
		}

	}
}

//leetcode submit region end(Prohibit modification and deletion)
