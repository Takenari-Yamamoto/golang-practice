package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}        // 最終的な結果を格納する二次元スライス
	queue := []*TreeNode{root} // ツリーのノードを探索するためのキュー

	for len(queue) > 0 {
		levelSize := len(queue)                // 現在のレベルにあるノードの数
		currentLevel := make([]int, levelSize) // 現在のレベルのノードの値を格納するスライス

		for i := 0; i < levelSize; i++ {
			node := queue[0]  // キューの先頭のノードを取り出す
			queue = queue[1:] // 取り出したノードをキューから削除

			currentLevel[i] = node.Val // 現在のレベルのリストにノードの値を追加

			if node.Left != nil {
				queue = append(queue, node.Left) // 左の子ノードがあればキューに追加
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // 右の子ノードがあればキューに追加
			}
		}
		result = append(result, currentLevel) // 完成したレベルのリストを結果に追加
	}

	return result
}
