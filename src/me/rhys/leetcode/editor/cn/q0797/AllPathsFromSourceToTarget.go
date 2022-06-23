//Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find
//all possible paths from node 0 to node n - 1 and return them in any order.
//
// The graph is given as follows: graph[i] is a list of all nodes you can visit
//from node i (i.e., there is a directed edge from node i to node graph[i][j]).
//
//
// Example 1:
//
//
//Input: graph = [[1,2],[3],[3],[]]
//Output: [[0,1,3],[0,2,3]]
//Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
//
//
// Example 2:
//
//
//Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
//Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
//
//
//
// Constraints:
//
//
// n == graph.length
// 2 <= n <= 15
// 0 <= graph[i][j] < n
// graph[i][j] != i (i.e., there will be no self-loops).
// All the elements of graph[i] are unique.
// The input graph is guaranteed to be a DAG.
//
// Related Topics 深度优先搜索 广度优先搜索 图 回溯 👍 298 👎 0

package q0797

//leetcode submit region begin(Prohibit modification and deletion)
var paths [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	paths = make([][]int, 0)
	var path []int
	travers(graph, 0, path)
	return paths
}

func travers(graph [][]int, node int, path []int) {
	path = append(path, node)

	outdegrees := graph[node]
	// 如果是遍历所有路径就是判断出度，题目是从0到n-1这两个节点的路径，那就直接判断终点是不是n-1就完了
	//如果出度为空，则是终点
	//if len(outdegrees) == 0 {
	//	paths = append(paths, path)
	//}

	if node == len(graph)-1 {
		paths = append(paths, path)
	}
	for _, outdegree := range outdegrees {
		newPath := make([]int, len(path))
		copy(newPath, path)
		travers(graph, outdegree, newPath)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
