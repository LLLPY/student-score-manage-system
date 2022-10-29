package template

import "fmt"

//老师登录后看到的主菜单
func Teacher_main_menu() (choice int) {
	choice = -1
	fmt.Printf("================\n")
	fmt.Printf("1.查看个人信息\n")
	fmt.Printf("2.学生成绩列表\n")
	fmt.Printf("3.查询学生成绩\n")
	fmt.Printf("4.成绩分析\n")
	fmt.Printf("5.管理班级\n")
	fmt.Printf("0.退出\n")
	fmt.Printf("================\n")
	fmt.Printf("请输入要使用的功能：")
	fmt.Scanln(&choice)

	return
}

//管理班级--菜单
func Teacher_Manage_Class() (choice int) {
	choice = -1
	fmt.Printf("================\n")
	fmt.Printf("1.学生信息列表\n")
	fmt.Printf("2.查询学生信息\n")
	fmt.Printf("3.新增学生信息\n")
	fmt.Printf("4.更新学生信息\n")
	fmt.Printf("5.删除学生信息\n")
	fmt.Printf("0.退出\n")
	fmt.Printf("================\n")
	fmt.Printf("请输入要使用的功能：")
	fmt.Scanln(&choice)

	return
}
