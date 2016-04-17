package main

import (
	"flag"
	"fmt"
	"github.com/nguyenvanduocit/gohoro"
)

func main() {
	var signName string
	flag.StringVar(&signName, "sign", "", "Sign name")
	flag.Parse()
	if signName == "" {
		var signId int = -1
		for k, v := range gohoro.SignMap {
			fmt.Println(k, "--", v)
		}
		isFirst:=true
		for signName == ""{
			if !isFirst{
				fmt.Print("This sign is not exist, please type a number from the list below :")
			}else{
				isFirst = false;
				fmt.Print("Please choose a horoscope (number):")
			}
			fmt.Scanf("%d", &signId)
			if signId==0{
				return
			}
			signName = gohoro.GetSignNameById(signId)
		}
	}
	fmt.Println("Please wait ...")
	content, err := gohoro.GetHoroscope(signName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
}
