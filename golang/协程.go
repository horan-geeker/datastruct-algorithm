package main

import "fmt"

func test(list []int, c chan int) {
    for _,v := range list {
        c <- v
    }
    close(c)
}



func main () {
    list := []int{1,-1,20,4,7,9,100,8,5,2}
    c := make(chan int, 2)
    go test(list, c)
    for i := range c {
        fmt.Println(i)
    }
}