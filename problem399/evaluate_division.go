package problem399

// 参看
// https://blog.csdn.net/fuxuemingzhu/article/details/82591165
// https://zxi.mytechroad.com/blog/graph/leetcode-399-evaluate-division/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	paths := make(map[string]map[string]float64)

	for k, v := range equations {
		if _, ok := paths[v[0]]; !ok{
			paths[v[0]] = make(map[string]float64)
		}
		paths[v[0]][v[1]] = values[k]

		if _, ok := paths[v[1]]; !ok {
			paths[v[1]] = make(map[string]float64)
		}
		paths[v[1]][v[0]] = 1/values[k]

		paths[v[0]][v[0]] = 1.0
		paths[v[1]][v[1]] = 1.0
	}

	visited := make(map[string]bool)
	var dfs func(node, end string) float64
	dfs = func(node, end string) float64{
		if v,ok := paths[node][end]; ok {
			return v
		}

		visited[node] = true
		for path := range paths[node]{
			if visited[path] {
				continue
			}

			res := dfs(path, end)
			if res != -1 {
				return paths[node][path] * res
			}
		}

		return -1
	}

	res := make([]float64, len(queries))
	for k, v := range queries{
		visited = map[string]bool{}
		res[k] = dfs(v[0], v[1])
	}

	return res
}