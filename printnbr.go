package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int){
    if n<0{
        x := n*(-1)
       z01.PrintRune('-')
       z01.PrintRune(x)
    }else if n>0{
        fz01.PrintRune(n)
    }else{
        fz01.PrintRune('0')
    }
}