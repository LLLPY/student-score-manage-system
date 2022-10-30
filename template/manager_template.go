package template

import "fmt"

//管理班级--菜单
func Manager_Main_Menu() (choice int) {
	choice = -1
	fmt.Printf("================\n")
	fmt.Printf("1.查看个人信息\n")
	fmt.Printf("2.教师信息列表\n")
	fmt.Printf("3.查询教师信息\n")
	fmt.Printf("4.新增教师信息\n")
	fmt.Printf("5.更新教师信息\n")
	fmt.Printf("6.删除教师信息\n")
	fmt.Printf("0.退出\n")
	fmt.Printf("================\n")
	fmt.Printf("请输入要使用的功能：")
	fmt.Scanln(&choice)

	return
}
