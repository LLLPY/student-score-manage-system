package template

import (
	"fmt"
)

//学生登录后看到的主菜单
func Student_main_menu() (choice string) {

	fmt.Printf("================\n")
	fmt.Printf("1.查看个人信息\n")
	fmt.Printf("2.查询成绩\n")
	fmt.Printf("3.成绩PK\n")
	fmt.Printf("4.成绩分析\n")
	fmt.Printf("0.退出\n")
	fmt.Printf("================\n")
	fmt.Printf("请输入要使用的功能：")
	fmt.Scanln(&choice)
	return
}
