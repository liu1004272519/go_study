package main

import "fmt"

func justify(n int) bool{
	if n <= 1 {
		return false
    }

    for i := 2; i < n;i++{
    	if n % i == 0{
    		return false
		}
	}

    return true
}

func example1(){
	for i := 2;i < 100;i++{
		if justify(i) == true{
			fmt.Printf("[%d] is prime\n",i)
		}
	}
}

func is_shuixianhua(n int) bool{
	first := n % 10
	second := (n % 10) % 10
	third := (n % 100) % 10

	sum := first*first*first + second * second * second + third*third*third

	return sum == n
}

func example2(){
	for i := 100;i < 1000;i++{
		if is_shuixianhua(i) {
			fmt.Printf("[%d] is shuixianhua \n",i)
		}
	}
}


func calc(str string) (charCount int,numCount int, spaceCount int,otherCount int){
	utfChars := []rune(str)

	for i := 0; i < len(utfChars);i++{
		if (utfChars[i] >= 'a' && utfChars[i] <= 'z' || utfChars[i] >= 'A' && utfChars[i] <= 'Z'){
			charCount ++
			continue
		}

		if utfChars[i] >= '0' && utfChars[i] <= '9'{
			numCount ++
			continue
		}

		if utfChars[i] == ' '{
			spaceCount ++
			continue
		}
	}
	return
}

func example3(){
	var str string = "dfka    我爱天安门 3833333"
	charCount,numCount,spCount,other := calc(str)
	fmt.Printf("char count:%d num count:%d sp count: %d other:%d \n",charCount,numCount,spCount,other)
}

func main(){
	example3()
}