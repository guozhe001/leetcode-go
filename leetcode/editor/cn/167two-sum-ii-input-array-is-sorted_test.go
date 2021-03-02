package cn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	sum := twoSum([]int{1, 2, 3, 4, 5, 6}, 4)
	assert.Equal(t, 1, sum[0])
	assert.Equal(t, 3, sum[1])
}

//输入：numbers = [2,7,11,15], target = 9
//输出：[1,2]
//解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
func TestTwoSum1(t *testing.T) {
	sum := twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, 1, sum[0])
	assert.Equal(t, 2, sum[1])
}

// 示例 2：
//
//
//输入：numbers = [2,3,4], target = 6
//输出：[1,3]
//
func TestTwoSum2(t *testing.T) {
	sum := twoSum([]int{2, 3, 4}, 6)
	assert.Equal(t, 1, sum[0])
	assert.Equal(t, 3, sum[1])
}

// 示例 3：
//
//
//输入：numbers = [-1,0], target = -1
//输出：[1,2]
func TestTwoSum3(t *testing.T) {
	sum := twoSum([]int{-1, 0}, -1)
	assert.Equal(t, 1, sum[0])
	assert.Equal(t, 2, sum[1])
}

// [12,13,23,28,43,44,59,60,61,68,70,86,88,92,124,125,136,168,173,173,180,199,212,221,227,230,277,282,306,314,316,321,325,328,336,337,363,365,368,370,370,371,375,384,387,394,400,404,414,422,422,427,430,435,457,493,506,527,531,538,541,546,568,583,585,587,650,652,677,691,730,737,740,751,755,764,778,783,785,789,794,803,809,815,847,858,863,863,874,887,896,916,920,926,927,930,933,957,981,997]
//542
func TestTwoSum4(t *testing.T) {
	numbers := []int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997}
	sum := twoSum(numbers, 542)
	if len(sum) > 0 {
		fmt.Printf("result: first=%d, second=%d", numbers[sum[0]-1], numbers[sum[1]-1])
		assert.Equal(t, 1, sum[0])
		assert.Equal(t, 2, sum[1])
	}
}
