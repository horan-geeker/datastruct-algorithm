package main

import (
    "strings"
    "strconv"
    "fmt"
    "errors"
)

const IP_A_BEGIN uint32 = 0x01000000
const IP_A_END uint32 = 0x7EFFFFFF

const IP_B_BEGIN uint32 = 0x80000000
const IP_B_END uint32 = 0xBFFFFFFF

const IP_C_BEGIN uint32 = 0xC0000000
const IP_C_END uint32 = 0xDFFFFFFF

const IP_D_BEGIN uint32 = 0xE0000000
const IP_D_END uint32 = 0xEFFFFFFF

const IP_E_BEGIN uint32 = 0xF0000000
const IP_E_END uint32 = 0xFFFFFFFF

const PRIVATE_IP_10_BEGIN uint32 = 0x0A000000
const PRIVATE_IP_10_END uint32 = 0x0AFFFFFF
const PRIVATE_IP_172_BEGIN uint32 = 0xAC100000
const PRIVATE_IP_172_END uint32 = 0xAC1FFFFF
const PRIVATE_IP_192_BEGIN uint32 = 0xC0A80000
const PRIVATE_IP_192_END uint32 = 0xC0A8FFFF

func IPAddress() {
    var A,B,C,D,E,errAddress,privateAddress int
    var input, ip,mask string
    for {
        var ipSum uint32
        _, err := fmt.Scanf("%s", &input)
        if err != nil {
            break
        }
        t := strings.Split(input, "~")
        ip = t[0]
        mask = t[1]
        if !validMask(mask) {
            errAddress++
            continue
        }
        r,err := validIP(ip)
        if err != nil {
            continue
        }
        if !r {
            errAddress++
            continue
        }
        ip0, ip1, ip2, ip3, _ := IPSplit(ip)
        ipSum += uint32(ip0) << 24
        ipSum += uint32(ip1) << 16
        ipSum += uint32(ip2) << 8
        ipSum += uint32(ip3)
        if ipSum >= IP_A_BEGIN && ipSum <= IP_A_END {
            A++
        } else if ipSum >= IP_B_BEGIN && ipSum <= IP_B_END {
            B++
        } else if ipSum >= IP_C_BEGIN && ipSum <= IP_C_END {
            C++
        } else if ipSum >= IP_D_BEGIN && ipSum <= IP_D_END {
            D++
        } else if ipSum >= IP_E_BEGIN {
            E++
        }
        if (ipSum >= PRIVATE_IP_10_BEGIN && ipSum <= PRIVATE_IP_10_END) || (ipSum >= PRIVATE_IP_172_BEGIN && ipSum <= PRIVATE_IP_172_END) || (ipSum >= PRIVATE_IP_192_BEGIN && ipSum <= PRIVATE_IP_192_END) {
            privateAddress++
        }
    }
    fmt.Println(A,B,C,D,E,errAddress,privateAddress)
}

func IPSplit(ip string) (uint64,uint64,uint64,uint64,error) {
    ips := strings.Split(ip, ".")
    if len(ips) != 4 {
        return 0,0,0,0, errors.New("ip length invalid")
    }
    ip0,err := strconv.ParseUint(ips[0], 10, 8)
    if err != nil {
        return 0,0,0,0, err
    }
    ip1,err := strconv.ParseUint(ips[1], 10, 8)
    if err != nil {
        return 0,0,0,0, err
    }
    ip2,err := strconv.ParseUint(ips[2], 10, 8)
    if err != nil {
        return 0,0,0,0, err
    }
    ip3,err := strconv.ParseUint(ips[3], 10, 8)
    if err != nil {
        return 0,0,0,0, err
    }
    return ip0,ip1,ip2,ip3,nil
}

func validIP(ip string) (bool,error) {
    ip0, _, _, _, err := IPSplit(ip)
    if err != nil {
        return false, nil // 错误 ip 但是需要统计
    }
    if ip0 == 0 || ip0 == 127 {
        return false, errors.New("环回地址") // 不统计
    }
    return true, nil
}

func validMask(mask string) bool {
    p0,p1,p2,p3,err := IPSplit(mask)
    if err != nil {
        return false
    }
    // 4 个都是 0
    if p0 == 0 && p1 == 0 && p2 == 0 && p3 == 0{
        return false
    }
    // 4 个都是 255 非法
    if p0 == 255 && p1 == 255 && p2 == 255 && p3 == 255{
        return false
    }
    // 如果b取反后+1与b进行或运算，如果结果=b则说明是合法掩码，否则为非法掩码
    var b uint32 = 0
    b += uint32(p0) << 24
    b += uint32(p1) << 16
    b += uint32(p2) << 8
    b += uint32(p3)
    return (b | (^b + 1)) == b
}

func main() {
    IPAddress()
}