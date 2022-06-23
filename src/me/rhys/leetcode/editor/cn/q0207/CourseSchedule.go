//There are a total of numCourses courses you have to take, labeled from 0 to
//numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai,
// bi] indicates that you must take course bi first if you want to take course ai.
//
//
//
// For example, the pair [0, 1], indicates that to take course 0 you have to
//first take course 1.
//
//
// Return true if you can finish all courses. Otherwise, return false.
//
//
// Example 1:
//
//
//Input: numCourses = 2, prerequisites = [[1,0]]
//Output: true
//Explanation: There are a total of 2 courses to take.
//To take course 1 you should have finished course 0. So it is possible.
//
//
// Example 2:
//
//
//Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
//Output: false
//Explanation: There are a total of 2 courses to take.
//To take course 1 you should have finished course 0, and to take course 0 you
//should also have finished course 1. So it is impossible.
//
//
//
// Constraints:
//
//
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// All the pairs prerequisites[i] are unique.
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 1313 👎 0

package q0207

//leetcode submit region begin(Prohibit modification and deletion)
var res bool

// 能不能完成课程，就是判断课程有没有循环依赖，也就是检测图的环
func canFinish(numCourses int, prerequisites [][]int) bool {
	res = true
	// 构建图
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph = append(graph, make([]int, 0))
	}
	for _, prerequisite := range prerequisites {
		graph[prerequisite[0]] = append(graph[prerequisite[0]], prerequisite[1])
	}
	// 探测环
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		traverse(graph, i, visited, onPath)
	}
	return res
}

func traverse(graph [][]int, s int, visited, onPath []bool) {
	if !res {
		return
	}
	// 如果OnPath在路径上，则出现环了
	if onPath[s] {
		res = false
		return
	}

	if visited[s] {
		return
	}
	visited[s] = true

	onPath[s] = true

	for _, d := range graph[s] {
		traverse(graph, d, visited, onPath)
	}
	onPath[s] = false
}

//leetcode submit region end(Prohibit modification and deletion)
