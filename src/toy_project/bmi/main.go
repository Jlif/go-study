package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var weight float64
		var height float64
		var age int32

		fmt.Println("欢迎使用 BMI 计算器")

		fmt.Println("请输入您的体重，单位：千克")
		fmt.Scanln(&weight)

		fmt.Println("请输入您的身高，单位：米")
		fmt.Scanln(&height)

		fmt.Println("请输入您的年龄：单位：岁")
		fmt.Scanln(&age)

		bmi := weight / (height * height)
		var msg string

		if bmi <= 18.4 {
			msg = "偏瘦"
		} else if bmi <= 23.9 {
			msg = "正常"
		} else if bmi <= 27.9 {
			msg = "过重"
		} else {
			msg = "肥胖"
		}
		fmt.Printf("您的身体质量指数为 %.2f，指标为 %s\n", bmi, msg)

		var againFlag string
		fmt.Println("是否继续？y/n")
		fmt.Scanln(&againFlag)

		if strings.EqualFold(againFlag, "y") {
		} else if strings.EqualFold(againFlag, "n") {
			fmt.Println("感谢使用，再见！")
			break
		} else {
			fmt.Println("非法输入，将退出")
			break
		}
	}

}
