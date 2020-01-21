package copypasta

// https://oi-wiki.org/dp/

/*
若使用滚动数组，注意在下次复用时初始化第一排所有元素
但是实际情况是使用滚动数组仅降低了内存，执行效率与不使用时无异
*/

func dpCollections() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	generalDP := func(a []int) int {
		n := len(a)
		cost := func(l, r int) int {
			return 1
		}
		const mx = 505
		dp := [mx][mx]int{}
		vis := [mx][mx]bool{}
		var f func(l, r int) int
		f = func(l, r int) (ans int) {
			// 边界检查
			if l >= r {
				return 0
			}
			if vis[l][r] {
				return dp[l][r]
			}
			vis[l][r] = true
			defer func() { dp[l][r] = ans }()
			// 转移方程
			if a[l] == a[r] {
				return f(l+1, r-1)
			}
			f1 := f(l+1, r) + cost(l, r)
			f2 := f(l, r-1) + cost(l, r)
			return min(f1, f2)
		}
		return f(0, n-1)
	}

	generalDP2 := func(x, y int) int {
		type pair struct{ x, y int }
		dp := map[pair]int{}
		var f func(x, y int) int
		f = func(x, y int) (ans int) {
			// 边界检查
			// ...
			p := pair{x, y}
			if v, ok := dp[p]; ok {
				return v
			}
			defer func() { dp[p] = ans }()
			// 转移方程
			// ...
			return
		}
		return f(x, y)
	}

	knapsack01 := func(values, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			wi := weights[i]
			for j, dpij := range dp[i] {
				if j < wi {
					dp[i+1][j] = dpij
				} else {
					dp[i+1][j] = max(dpij, dp[i][j-wi]+vi)
				}
			}
		}
		return dp[n][maxW]
	}

	// TODO: 单调队列/单调栈优化
	// https://oi-wiki.org/dp/opt/monotonous-queue-stack/

	// TODO: 斜率优化
	// https://oi-wiki.org/dp/opt/slope/

	// TODO: 四边形不等式优化
	// https://oi-wiki.org/dp/opt/quadrangle/

	// 树上最大匹配
	// https://codeforces.com/blog/entry/2059
	// g[v] = ∑{max(f[son],g[son])}
	// f[v] = max{1+g[son]+g[v]−max(f[son],g[son])}
	maxMatchingOnTree := func(n int, g [][]int) int {
		cover, nonCover := make([]int, n), make([]int, n)
		vis := make([]bool, n)
		var f func(int)
		f = func(v int) {
			vis[v] = true
			for _, w := range g[v] {
				if !vis[w] {
					f(w)
					nonCover[v] += max(cover[w], nonCover[w])
				}
			}
			for _, w := range g[v] {
				cover[v] = max(cover[v], 1+nonCover[w]+nonCover[v]-max(cover[w], nonCover[w]))
			}
		}
		f(0)
		return max(cover[0], nonCover[0])
	}

	_ = []interface{}{
		generalDP, generalDP2, knapsack01,
		maxMatchingOnTree,
	}
}
