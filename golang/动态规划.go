package main

import (
    "fmt"
    "math"
)

func makeDynamicArr(l1 int, l2 int) [][]int {
    arr := make([][]int, l1)
    for i := 0; i < l1; i++ {
        arr[i] = make([]int, l2)
    }
    return arr
}

func makeDynamicArrThree(l1 int, l2 int, l3 int) [][][]int {
    arr := make([][][]int, l1)
    for i := 0; i < l1; i++ {
        arr[i] = make([][]int, l2)
        for j := 0; j < l2; j++ {
            arr[i][j] = make([]int, l3)
        }
    }
    return arr
}

// 假设一个二维的地图，玩家位于初始位置 0,0 ，现在走到终点 x,y 有多少种走法，其中空地是 0 ，石头是 1 不可以走
//matrix := [][]int{
//{0, 0, 0, 0, 0, 0, 0, 0},
//{0, 0, 1, 0, 0, 0, 1, 0},
//{0, 0, 0, 0, 1, 0, 0, 0},
//{1, 0, 1, 0, 0, 1, 0, 0},
//{0, 0, 1, 0, 0, 0, 0, 0},
//{0, 0, 0, 1, 1, 0, 1, 0},
//{0, 1, 0, 0, 0, 1, 0, 0},
//{0, 0, 0, 0, 0, 0, 0, 0},
//}
func findMapPaths(matrix [][]int) int {
    x := len(matrix)
    y := len(matrix[0])
    dp := makeDynamicArr(x+1, y+1)
    for i := x - 1; i >= 0; i-- {
        for j := y - 1; j >= 0; j-- {
            if i == (x-1) && j == (y-1) {
                dp[i][j] = 0
            } else if i == (x-1) && j == (y-2) {
                if matrix[i][j] == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = 0
                }
            } else if j == (y-1) && i == (x-2) {
                if matrix[i][j] == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = 0
                }
            } else if matrix[i][j] == 0 {
                dp[i][j] = dp[i][j+1] + dp[i+1][j]
            } else {
                dp[i][j] = 0
            }
        }
    }
    return dp[0][0]
}

// 爬楼梯
func climbStairs(n int, dp map[int]int) int {
    var res int
    if _, ok := dp[n]; ok {
        return dp[n]
    }
    if n == 1 {
        res = 1
    } else if n == 2 {
        res = 2
    } else {
        res = climbStairs(n-1, dp) + climbStairs(n-2, dp)
    }
    dp[n] = res
    return res
}

// 三角形最小和 leetcode 120
//triangle := [][]int{
//    {2},
//    {3,4},
//    {6,5,7},
//    {4,1,8,3},
//}
func minimumTotal(triangle [][]int) int {
    dp := makeDynamicArr(len(triangle), len(triangle))
    for i := len(triangle) - 1; i >= 0; i-- {
        for j := i; j >= 0; j-- {
            if i == (len(triangle) - 1) {
                dp[i][j] = triangle[i][j]
            } else {
                min := dp[i+1][j]
                if min > dp[i+1][j+1] {
                    min = dp[i+1][j+1]
                }
                dp[i][j] = triangle[i][j] + min
            }
        }
    }
    return dp[0][0]
}

// 乘积最大子数组 leetcode 152
func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    res, currentMax, currentMin := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        currentMax, currentMin = maxThree(currentMax*nums[i], currentMin*nums[i], nums[i]), minThree(currentMax*nums[i], currentMin*nums[i], nums[i])
        if res < currentMax {
            res = currentMax
        }
    }
    return res
}

func maxThree(i int, j int, k int) int {
    max := i
    if max < j {
        max = j
    }
    if max < k {
        max = k
    }
    return max
}

func minThree(i int, j int, k int) int {
    min := i
    if min > j {
        min = j
    }
    if min > k {
        min = k
    }
    return min
}

func max(i int, j int) int {
    if i < j {
        return j
    }
    return i
}

// 买卖股票问题 leetcode 122
func maxProfit(prices []int) int {
    k := 2 + 1
    j := 2
    // i 天 k 买卖次数 j 股票数量 0 不持股 1 持股
    dp := makeDynamicArrThree(len(prices), k, j)
    dp[0][0][0], dp[0][0][1] = 0, -prices[0]
    dp[0][1][0], dp[0][1][1], dp[0][2][0], dp[0][2][1] = -math.MaxInt32,-math.MaxInt32,-math.MaxInt32,-math.MaxInt32
    for i := 1; i < len(prices); i++ {
        dp[i][0][0] = dp[i -1][0][0]
        dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0] - prices[i])
        dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][1] + prices[i])
        dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][1][0] - prices[i])
        dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][1][1] + prices[i])
    }
    return maxThree(dp[len(prices) - 1][0][0],dp[len(prices) - 1][1][0],dp[len(prices) - 1][2][0])
}

func main() {
    example := []int{7,1,5,3,6,4}
    fmt.Println(maxProfit(example))
}
