package main

import (
    "fmt"
)

func dynamicArr(l1 int, l2 int) [][]int {
    arr := make([][]int, l1)
    for i := 0; i < l1; i++{
        arr[i] = make([]int, l2)
    }
    return arr
}

func longestCommonSequence(a []int, b []int) int {
    dp := dynamicArr(len(a) + 1, len(b) + 1)
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(b); j++ {
            if a[i] == b[j] {
                dp[i+1][j+1] = 1 + dp[i][j]
            } else {
                dp[i+1][j+1] = Max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    return dp[len(a)][len(b)]
}

func main() {
    a := []int{1, 3, 5, 7, 9}
    b := []int{1, 2, 3, 4, 6, 7, 8, 9}
    fmt.Println(longestCommonSequence(a, b))
}