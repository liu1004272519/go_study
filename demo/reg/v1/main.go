package main

import (
	"regexp"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func IsIP(ip string) (b bool) {
	if m,_ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",ip); !m {
		return false
	}
	return true
}

func GetContent(){
	resp,err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println("http get error")
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	re,_ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src,strings.ToLower)

	re, _ = regexp.Compile("\\>style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src,"")

	re,_ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src,"")

	re,_ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src,"\n")

	re,_ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src,"\n")

	fmt.Println(strings.TrimSpace(src))
}


func main(){
	//if len(os.Args) == 1 {
	//	fmt.Println("Usage:regexp [string]")
	//	os.Exit(1)
	//} else if m,_ := regexp.MatchString("^[0-9]+$",os.Args[1]);m{
	//	fmt.Println("数字")
	//} else {
	//	fmt.Println("不是数字")
	//}
	//
	//if b := IsIP("127.0.0.1"); !b{
	//	log.Fatal("is not ip")
	//}
	//log.Println("is ip address")
	GetContent()
}
