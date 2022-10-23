package template

import (
	"fmt"
	"os"
	"student-score-manage-system/utils"
)

//登录界面
func Login_menu() (number string, password string, tmp_err error) {
	fmt.Printf("\n%-32v\n", "==============欢迎使用==============")
	fmt.Printf("%-32v\n", "#                                  #")
	fmt.Printf("%-32v\n", "#                                  #")
	fmt.Printf("%-32v\n", "#    <<学 生 成 绩 管 理 系 统>>   #")
	fmt.Printf("%-32v\n", "#   student-score-manage-system    #")
	fmt.Printf("%-32v\n", "#                                  #")
	fmt.Printf("%-32v\n", "#     🐹🐹🐹               V 1.0   #")
	fmt.Printf("%-32v\n", "#   2022-10-12     powered by go   #")
	fmt.Printf("%-32v\n", "===================================")
	fmt.Printf("%-32v\n", "    1.登录               0.退出     ")

	a := utils.Legal_input_string("\n请输入你的选择：", map[string]string{"1": "", "0": ""})

	if a == "0" {
		fmt.Printf("欢迎再次使用...\n")
		os.Exit(0)
	}

	print("请输入账号(学号/教职工号/管理员账号)：")
	_, err := fmt.Scanln(&number)
	print("请输入密码：")
	_, err2 := fmt.Scanln(&password)

	if err == nil && err2 == nil {
		return number, password, nil
	} else {
		if err != nil {
			tmp_err = err

		} else {
			tmp_err = err2

		}
	}
	return
}
