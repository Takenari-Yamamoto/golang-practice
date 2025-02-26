package main

import "testing"

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		input    *TreeNode
		expected [][]int
	}{
		// {
		// 	name:     "空のツリー",
		// 	input:    nil,
		// 	expected: nil,
		// },
		// {
		// 	name: "単一ノード",
		// 	input: &TreeNode{
		// 		Val: 1,
		// 	},
		// 	expected: [][]int{{1}},
		// },
		{
			name: "完全な二分木",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LevelOrder(tt.input)
			if !slicesEqual(got, tt.expected) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// スライスの比較用ヘルパー関数
func slicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
