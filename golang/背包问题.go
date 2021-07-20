package main

import "fmt"

func packageQuestion() {
    var N, m int
    var v,p,q int
    arr := make([]Product, 0)
    _, err := fmt.Scanf("%d %d", &N, &m)
    if err != nil {
        panic(err)
    }
    var result [][]int
    for i:=0; i<m; i++ {
        fmt.Scanf("%d %d %d", &v, &p, &q)
        arr = append(arr, Product{v, p, q})
    }
    for j:=1; j<len(arr); j++ {
        for k:=0; k<N; k+=10 {
            if k < N {
                result[j][N-k] = arr[j].v
            }
        }
    }
    fmt.Println(result[N][m])
}
