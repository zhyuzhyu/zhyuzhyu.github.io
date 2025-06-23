package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	visited := make(map[*TreeNode]bool)
	prev := make(map[*TreeNode]*TreeNode)
	pq := []*TreeNode{root}
	visited[root] = true
	found := 0
	for len(pq) > 0 {
		cur := pq[0]
		pq = pq[1:]
		if cur.Val == p.Val || cur.Val == q.Val {
			found++
			if found == 2 {
				break
			}
		}
		if !visited[cur.Left] {
			prev[cur.Left] = cur
			pq = append(pq, cur.Left)
		}
		if !visited[cur.Right] {
			prev[cur.Right] = cur
			pq = append(pq, cur.Right)
		}
	}

	for i := prev[p]; i != nil; i = prev[i] {

	}

	return prev[q]
}
