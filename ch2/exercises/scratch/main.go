package main
import (
    "fmt"
    "os"
    "strconv"
    "log"
)

var pc [256]byte

func init(){
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func PopCount(x uint64) int {
    return int(pc[byte(x>>(0*8))] + 
        pc[byte(x>>(1*8))] + 
        pc[byte(x>>(2*8))] + 
        pc[byte(x>>(3*8))] + 
        pc[byte(x>>(4*8))] + 
        pc[byte(x>>(5*8))] + 
        pc[byte(x>>(6*8))] + 
        pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
    count := 0
    for x > 0 {
        count += int(x & 1)
        x = x >> 1
    }
    return count
}


func main(){
    for _,i := range os.Args[1:] {
        num,err := strconv.Atoi(i);
        if err != nil {
            log.Fatalln(err)
        }
        fmt.Printf("V1 %d: %d\n",num,PopCount(uint64(num)))
        fmt.Printf("V2 %d: %d\n",num,PopCount2(uint64(num)))
    }
}

