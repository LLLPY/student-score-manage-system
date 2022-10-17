package main

import (
	"fmt"
	"os"
	"student-score-manage-system/model"
	"student-score-manage-system/template"
)

//初始化操作
func init() {
	score_init_err := model.Score{}.Read_to_buffer("model/score.txt")

	student_init_err := model.Student{}.Read_to_buffer("model/student.txt")
	teacher_init_err := model.Teacher{}.Read_to_buffer("model/teacher.txt")
	manager_init_err := model.Manager{}.Read_to_buffer("model/manager.txt")
	if score_init_err != nil || student_init_err != nil || teacher_init_err != nil || manager_init_err != nil {
		print("程序初始化失败...")
		os.Exit(0) //直接退出程序

	}
}

func main() {

	// number, password, err := template.Login_menu()

	// if err == nil {
	// 	fmt.Printf("number: %v\n", number)
	// 	fmt.Printf("password: %v\n", password)
	// }
	number := "2022000011"
	password := "1234"
	// var user model.User
	// var user model.User
	user_type := string(number[len(number)-1])

	switch user_type {
	case "1": //学生
		user := model.STUDENT_BUF[number]
		ok := user.Login(number, password)
		if ok {

			for {
				func_choice := template.Student_main_menu()
				switch func_choice {
				case "1":
					user.Show_info()
				case "2":
					user.Find()
				case "3":
					user.Score_Pk()
				case "4":
				case "0":
					fmt.Printf("欢迎再次使用...\n")
					os.Exit(0)
				}
			}

		}
	// case "2":
	// 	user := TEACHER_BUF[number] //教师
	// case "3":
	// 	user := MANAGER_BUF[number] //管理员
	default:
		println("用户身份未识别，程序已退出...\n")
		os.Exit(0)

	}

}
