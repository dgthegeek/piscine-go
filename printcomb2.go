package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
for a := 0; a <= 98; a++ {
	for b := 1; b <= 99; b++ {
		if a<b {
			if a!=98 && b!=99  {
				fmt.Printf("%02d %02d, ",a,b)
			
				}else if a==98 && b==99 {
				fmt.Printf("%02d %02d\n",a,b)
				
	
				}
			}
		}		
	}
}