package main

import "fmt"


func QuickSort(list []int) []int {
    length := len(list)
    if length <= 1 {
        return list
    }
    m := list[0]
    left, right := make([]int, 0), make([]int, 0)
    for i := 1; i < length; i++ {
        if list[i] < m {
            left = append(left, list[i])
        } else {
            right = append(right, list[i])
        }
    }
    return append(append(quickSort(left), m), quickSort(right)...)
}

func main() {
    list := []int{1,-1,20,4,7,9,100,8,5,2}
    fmt.Println(quickSort(list))
}