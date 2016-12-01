package sliceFunc

import (
    "fmt"
)

func Slicer(i int64)  bool{
    
    if i <0 {
    	i= -1*i
    }

    if i < 9223372036854775807 {
    	a:=make([]int,10)
		s:=make([]int,0)
		for i>0{
			x:=0
			x=int(i%10)
			i=i/10
			a[x]=a[x]+1
		
			s=append(s,x)
		}
		for i := len(s)/2-1; i >= 0; i-- {
    		opp := len(s)-1-i
    		s[i], s[opp] = s[opp], s[i]
		
		}
	fmt.Println(s)
	fmt.Println(a)
	return true
	} else {
		fmt.Println("you entered an  out of range invalid no.")
		//fmt.Println("enter a valid no. less then 9223372036854775808")
    	//fmt.Println("enter the valid series of digits")
		//fmt.Scanf("%d",&i)
		//Slicer(i)
		return  false
    }

}

func main() {
    var i int64
	fmt.Println("enter the series of digits less then 9223372036854775807")
	fmt.Scanf("%d",&i)
    
    fmt.Println(i)
    Slicer(i)
}