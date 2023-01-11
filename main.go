package main

import (
    "fmt"
    "regexp"
    "os"
    "strings"
    "strconv"
)

const BYTE = 0x100

type IPv4 struct {
    OCTET_2 int
    OCTET_3 int
    OCTET_4 int
}


func sliceAtoi(sa []string) ([]int, error) {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, err := strconv.Atoi(a)
        if err != nil {
            return si, err
        }
        si = append(si, i)
    }
    return si, nil
}

func (v4 IPv4) decimal(ip string) int{
  
    REGEX := "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)[.]){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"
  
    s_code, _ := regexp.MatchString(REGEX, ip)
    if s_code == false {
        fmt.Fprintln(os.Stderr, "Invalid IP Address!")
        os.Exit(1)
    }
  
    values, _ := sliceAtoi(strings.Split(os.Args[1], "."))
    return (values[0] * v4.OCTET_4) + (values[1] * v4.OCTET_3) + (values[2] * v4.OCTET_2) + values[3]

}

func (v4 IPv4) address(dec int) string{
   
    if dec < 0 || dec > 4294967295{
        fmt.Fprintln(os.Stderr, "Invalid decimal value!")
        os.Exit(1)
    }
    return fmt.Sprintf("%d.%d.%d.%d", (dec/v4.OCTET_4) % BYTE, 
                                      (dec/v4.OCTET_3) % BYTE, 
                                      (dec/v4.OCTET_2) % BYTE, 
                                       dec % BYTE)
}

func main() {

    cache := os.Args[1]
    ipv4 := IPv4{BYTE, 65536, 16777216}
    s_code, _ := regexp.MatchString("(?:(?:[0-9]{1,3}[.]){3}[0-9]{1,3})", cache)
    if s_code == true {
        fmt.Println(ipv4.decimal(cache))
    } else {
        code, _ := regexp.MatchString("^[0-9]{1,10}$", cache)
        if code == true {
            dec, _ := strconv.Atoi(cache)
            fmt.Println(ipv4.address(dec))
        } else {
            fmt.Fprintln(os.Stderr, "Invalid input detected!")
            os.Exit(1)
        }
    }

}
           
