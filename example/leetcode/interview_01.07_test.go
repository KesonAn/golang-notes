//给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
//
// 不占用额外内存空间能否做到？
//
//
//
// 示例 1:
//
//
//给定 matrix =
//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
//
//
// 示例 2:
//
//
//给定 matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
//
//
// 注意：本题与主站 48 题相同：https://leetcode-cn.com/problems/rotate-image/
// Related Topics 数组
// 👍 169 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	assert.Equal(t, [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}, matrix)

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	assert.Equal(t, [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}, matrix)
}

func rotate(matrix [][]int) {
	if len(matrix) < 2 {
		return
	}
	var (
		n    = len(matrix)
		list = make([][]int, n)
	)
	for x := 0; x < n; x++ {
		inner := make([]int, n)
		for y := n - 1; y >= 0; y-- {
			inner[n-1-y] = matrix[y][x]
		}
		list[x] = inner
	}
	copy(matrix, list)
}
