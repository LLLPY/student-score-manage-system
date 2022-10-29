package main

import (
	"fmt"
	"os"
	"student-score-manage-system/model"
	"student-score-manage-system/template"
)

// 初始化操作
func init() {
	score_init_err := model.Score{}.Read_to_buffer("model/data/score.txt")

	student_init_err := model.Student{}.Read_to_buffer("model/data/student.txt")
	teacher_init_err := model.Teacher{}.Read_to_buffer("model/data/teacher.txt")
	manager_init_err := model.Manager{}.Read_to_buffer("model/data/manager.txt")
	if score_init_err != nil || student_init_err != nil || teacher_init_err != nil || manager_init_err != nil {
		print("程序初始化失败...")
		os.Exit(0) //直接退出程序

	}
}

func main() {
label:
	fmt.Print("11111\n")
	// number, password, err := template.Login_menu()

	// if err == nil {
	// 	fmt.Printf("number: %v\n", number)
	// 	fmt.Printf("password: %v\n", password)
	// }

	number := "2022000012"
	password := "1234"

	if len(number) != 10 {
		fmt.Print("账号不正确...\n")
		goto label
	}
	user_type := string(number[len(number)-1])

	switch user_type {
	case "1": //学生
		user := model.STUDENT_BUF[number]
		ok := user.Login(number, password)
		if ok {

			for {
				func_choice := template.Student_main_menu()
				switch func_choice {
				case 1:
					user.Show_info()
				case 2:
					user.Find()
				case 3:
					user.Score_Pk()
				case 4:
					user.Score_Analyse()
				case 0:
					goto label

				}
			}

		} else {
			fmt.Print("密码不正确...\n")
			goto label
		}
	case "2":
		user := model.TEACHER_BUF[number] //教师
		ok := user.Login(number, password)
		if ok {
		teacher_loop:
			for {
				func_choice := template.Teacher_main_menu()
				switch func_choice {
				case 1:
					user.Show_Persional_info()
				case 2:
					user.Student_Score_List()
				case 3:
					user.Search_Student_Score()
				case 4:
					user.Analyse_Class_Score()
				case 5:
					for {
						func_choice := template.Teacher_Manage_Class()
						switch func_choice {
						case 1:
							user.Student_Info_List()
						case 2:
							user.Find_Student_Info()
						case 3:
							user.Add_Student_Info()
						case 4:
							user.Update_Student_Info()
						case 5:
							ok := user.Delete_Student_Info()
							if ok {
								fmt.Printf("删除成功!\n")
							} else {
								fmt.Printf("删除失败!\n")

							}
						case 0:
							goto teacher_loop
						}
					}

				case 0:
					goto label

				}

			}

		} else {
			fmt.Print("密码不正确...\n")
			goto label

		}

	// case "3":
	// 	user := MANAGER_BUF[number] //管理员
	default:
		println("用户身份未识别，程序已退出...\n")
		os.Exit(0)

	}

}
