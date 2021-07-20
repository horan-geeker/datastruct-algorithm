package main

import (
    "bufio"
    "os"
    "strings"
    "fmt"
    "strconv"
    "github.com/pkg/errors"
    "sort"
)

func Max(i int, j int) int {
    if i > j {
        return i
    }
    return j
}

func InStr(s string, b byte) bool {
    for i := 0; i < len(s); i++ {
        if s[i] == b {
            return true
        }
    }
    return false
}

// 最长不重复子串
func lengthOfLongestSubstring(s string) int {
    maxArr := []int{}
    for i := 0; i < len(s); i++ {
        temp := string(s[i])
        for j := i + 1; j < len(s); j++ {
            if InStr(temp, s[j]) {
                break
            } else {
                temp += string(s[j])
            }
        }
        maxArr = append(maxArr, len(temp))
    }
    max := 0
    for _, value := range maxArr {
        if value > max {
            max = value
        }
    }
    return max
}

func IsMirror(s string) bool {
    length := len(s)
    for i := 0; i < length; i++ {
        if s[i] != s[length-i-1] {
            return false
        }
    }
    return true
}

// 最长回文子串
func longestPalindrome(s string) string {
    maxArr := []string{}
    for i := 0; i < len(s); i++ {
        temp := string(s[i])
        for j := i + 1; j < len(s); j++ {
            if IsMirror(s[i : j+1]) {
                if len(temp) < len(s[i:j+1]) {
                    temp = s[i : j+1]
                }
            }
        }
        maxArr = append(maxArr, temp)
    }
    result := ""
    max := 0
    for _, value := range maxArr {
        if len(value) > max {
            max = len(value)
            result = value
        }
    }
    return result
}

func SubStr(strs []string, s string) bool {
    l := len(s)
    for _, v := range strs {
        if len(v) < l || v[:l] != s {
            return false
        }
    }
    return true
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
    res := ""
    if len(strs) <= 1 {
        return strs[0]
    }
    min := strs[0]
    for _, v := range strs {
        if len(v) < len(min) {
            min = v
        }
    }
    for j := 0; j < len(min); j++ {
        if SubStr(strs, min[:j+1]) {
            res = min[:j+1]
        }
    }
    return res
}

func isValid(s string) bool {
    mapstr := map[string]string{"{": "}", "[": "]", "(": ")"}
    t := ""
    for i := 0; i < len(s); i++ {
        if _, ok := mapstr[string(s[i])]; ok {
            t += string(s[i]) // push queue
        } else {
            if len(t) > 0 && mapstr[string(t[len(t)-1])] == string(s[i]) {
                t = t[:len(t)-1] // pop queue
            } else {
                return false
            }
        }
    }
    if len(t) != 0 {
        return false
    }
    return true
}

func Input() []string {
    args := make([]string, 0)
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        line := scanner.Text()
        args = strings.Fields(line)
    }
    return args
}

func InArray(arr []string, item string) bool {
    for _, v := range arr {
        if v == item {
            return true
        }
    }
    return false
}

func CountString(m string) {
    hash := make(map[string]int)
    strs := strings.Split(m, "")
    for i := len(strs) - 1; i >= 0; i-- {
        hash[strs[i]] = 1
    }
    fmt.Println(len(hash))
}

func ReserveString(m string) string {
    s := ""
    for i := len(m) - 1; i >= 0; i-- {
        s += string(m[i])
    }
    return s
}

func ReserveArray(strs []string) []string {
    s := make([]string, 0)
    for i := len(strs) - 1; i >= 0; i-- {
        s = append(s, strs[i])
    }
    return s
}

func QuickSort(strs []string) []string {
    if len(strs) < 2 {
        return strs
    }
    m := strs[0]
    left := make([]string, 0)
    right := make([]string, 0)
    for i := 1; i < len(strs); i++ {
        if strings.Compare(strs[i], m) == -1 {
            left = append(left, strs[i])
        } else {
            right = append(right, strs[i])
        }
    }
    return append(append(QuickSort(left), m), QuickSort(right)...)
}

func SortStrings() []string {
    var num int
    var s string
    arr := make([]string, 0)
    fmt.Scanf("%d", &num)
    for i := 0; i < num; i++ {
        fmt.Scanf("%s", &s)
        if s == "" {
            continue
        }
        arr = append(arr, s)
    }
    return QuickSort(arr)
}

type Product struct {
    v int
    p int
    q int
}

func SumV(pList map[int]Product) int {
    sum := 0
    for _, p := range pList {
        sum += p.v
    }
    return sum
}

func CountVP(pList map[int]Product) int {
    sum := 0
    for _, p := range pList {
        sum += p.v * p.p
    }
    return sum
}

func positionXY() {
    var s string
    var strs []string
    fmt.Scanf("%s", &s)
    strs = strings.Split(s, ";")
    x := 0
    y := 0
    for _, v := range strs {
        if len(v) < 2 {
            continue
        }
        b := string(v[0:1])
        i, err := strconv.ParseInt(v[1:], 10, 32)
        if i >= 100 || err != nil {
            continue
        }
        if b == "A" {
            x -= int(i)
        } else if b == "S" {
            y -= int(i)
        } else if b == "W" {
            y += int(i)
        } else if b == "D" {
            x += int(i)
        }
    }
    fmt.Printf("%d,%d", x, y)
}

type errorStack struct {
    filename string
    codeLine int
}

func InArrayErrorStack(es []errorStack, t errorStack) bool {
    for _, v := range es {
        if v.filename == t.filename && v.codeLine == t.codeLine {
            return true
        }
    }
    fmt.Println(es, t)
    return false
}

func LogLineNums() {
    var s string
    var n int
    var esArr = make([]errorStack, 0)
    var esSet = make(map[errorStack]int)
    for {
        _, err := fmt.Scanf("%s %d", &s, &n)
        if err != nil {
            break
        }
        t := strings.Split(s, "\\")
        s = t[len(t)-1]
        p := 0
        if len(s) > 16 {
            p = len(s) - 16
        }
        es := errorStack{s[p:], n}
        if _, ok := esSet[es]; ok {
            esSet[es]++
            continue
        }
        esSet[es]++
        esArr = append(esArr, es)
        if len(esArr) > 8 {
            esArr = esArr[1:]
        }
    }
    for _, v := range esArr {
        fmt.Println(v.filename, v.codeLine, esSet[v])
    }
}

func isDuplicateSubSeq(s string) bool {
    var set = make(map[string]int)
    for i := 0; i <= len(s)-3; i++ {
        if _, ok := set[s[i:i+3]]; ok {
            if set[s[i:i+3]]+3 <= i {
                return true
            }
        }
        set[s[i:i+3]] = i
    }
    return false
}

func passwordValid() {
    for {
        var s string
        _, err := fmt.Scanf("%s", &s)
        if err != nil {
            break
        }
        if len(s) <= 8 {
            fmt.Println("NG")
            continue
        }
        num := 0
        charcter := 0
        a := 0
        A := 0
        if isDuplicateSubSeq(s) {
            fmt.Println("NG")
            continue
        }
        for i := 0; i < len(s); i++ {
            if (s[i] >= 32 && s[i] <= 47) || (s[i] >= 58 && s[i] <= 64) || (s[i] >= 91 && s[i] <= 96) || (s[i] >= 123 && s[i] <= 126) {
                charcter = 1
            }
            if s[i] >= 48 && s[i] <= 57 {
                num = 1
            }
            if s[i] >= 65 && s[i] <= 90 {
                A = 1
            }
            if s[i] >= 97 && s[i] <= 122 {
                a = 1
            }
        }
        if num+charcter+a+A >= 3 {
            fmt.Println("OK")
        } else {
            fmt.Println("NG")
        }
    }
}

func Sum(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    return sum
}

func quickSort(list []int) []int {
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

func minSubsequence(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }
    sum := Sum(nums)
    p := sum / 2
    nums = quickSort(nums)
    res := make([]int, 0)
    for i := len(nums) - 1; i >= 0; i-- {
        if Sum(res) > p {
            return res
        }
        res = append(res, nums[i])
    }
    return res
}

func sortString(s string) string {
    if len(s) < 2 {
        return s
    }
    min := 0
    for i := 1; i < len(s); i++ {
        if s[min] > s[i] {
            min = i
        }
    }
    res := string(s[min])
    n := len(s)
    for {
        n--
        if n <= 0 {
            break
        }
        min, err := minS(s, min)
        if err != nil {
            continue
        }
        res += string(s[min])
    }
    return res
}

func minS(s string, prev int) (int, error) {
    min := 0
    for j:=0;j<len(s);j++{

    }
    err := errors.New("not found")
    for i := 1; i < len(s); i++ {
        if s[prev] < s[i] && s[min] > s[i] {
            min = i
            err = nil
        }
    }
    return min, err
}

func main() {

}
