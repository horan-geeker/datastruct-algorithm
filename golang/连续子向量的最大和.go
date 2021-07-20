package main

import (
    "math"
    "fmt"
)

// 暴力
func normal(list []float64) {
    var max float64 = 0
    for index := range list {
        var sum float64 = 0
        for j := index; j < len(list); j++ {
            sum += list[j]
        }
        max = math.Max(sum, max)
    }
    fmt.Println(max)
}

// 分治
func fenzhi(list []float64, l int, r int) float64 {
    if l > r {
        return 0
    }
    if l == r {
        return math.Max(0, list[0])
    }
    m := (l + r) / 2
    var lmax, lsum float64 = 0, 0
    for i := m; i >= l; i-- {
        lsum += list[i]
        lmax = math.Max(lmax, lsum)
    }
    var rmax, rsum float64 = 0, 0
    for j := m + 1; j <= r; j++ {
        rsum += list[j]
        rmax = math.Max(rmax, rsum)
    }
    return math.Max(lmax + rmax, math.Max(fenzhi(list, l, m), fenzhi(list, m + 1, r)))
}

// 动态规划
func dongTaiGuiHua(list []float64) float64 {
    var max, sum float64 = 0, 0
    for _,i := range list {
        if sum < 0 {
            sum = i
        } else {
            sum += i
        }
        max = math.Max(sum, max)
    }
    return max
}

func main() {
    list := []float64{10, -20, -20, -10,30,100,-90, 1000}
    fmt.Println(dongTaiGuiHua(list))
}

