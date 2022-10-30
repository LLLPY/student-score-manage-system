package main

import (
	"fmt"
	"os"
	"student-score-manage-system/model"
	"student-score-manage-system/template"
)

const (
	SCORE_DATA_PATH   string = "model/data/score.txt"
	STUDENT_DATA_PATH string = "model/data/student.txt"
	TEACHER_DATA_PATH string = "model/data/teacher.txt"
	MANAGE_DATA_PATH  string = "model/data/manager.txt"
)

// 初始化操作
func init() {
	score_init_err := model.Score{}.Read_to_buffer(SCORE_DATA_PATH)
	student_init_err := model.Student{}.Read_to_buffer(STUDENT_DATA_PATH)
	teacher_init_err := model.Teacher{}.Read_to_buffer(TEACHER_DATA_PATH)
	manager_init_err := model.Manager{}.Read_to_buffer(MANAGE_DATA_PATH)
	if score_init_err != nil || student_init_err != nil || teacher_init_err != nil || manager_init_err != nil {
		print("程序初始化失败...")
		os.Exit(0) //直接退出程序

	}
}

func main() {
main_loop:
	number, password, _ := template.Login_menu()
	user_type := string(number[len(number)-1])
	switch user_type {
	case "1": //学生
		user := model.STUDENT_BUF[number]
		msg, ok := user.Login(number, password)
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
					goto main_loop

				}
			}

		} else {
			fmt.Printf("%v\n", msg)
			goto main_loop
		}
	case "2": //教师
		user := model.TEACHER_BUF[number]
		msg, ok := user.Login(number, password)
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
							model.Student{}.Write_To_File(STUDENT_DATA_PATH, SCORE_DATA_PATH)
						case 4:
							user.Update_Student_Info()
							model.Student{}.Write_To_File(STUDENT_DATA_PATH, SCORE_DATA_PATH)
						case 5:
							ok := user.Delete_Student_Info()
							if ok {
								fmt.Printf("删除成功!\n")
								model.Student{}.Write_To_File(STUDENT_DATA_PATH, SCORE_DATA_PATH)

							} else {
								fmt.Printf("删除失败!\n")

							}
						case 0:
							goto teacher_loop
						}
					}

				case 0:
					goto main_loop

				}

			}

		} else {
			fmt.Printf("%v\n", msg)
			goto main_loop

		}
	case "3": //管理员
		user := model.MANAGER_BUF[number]
		msg, ok := user.Login(number, password)
		if ok {
			for {
				func_choice := template.Manager_Main_Menu()
				switch func_choice {
				case 1:
					user.Show_Persional_info()
				case 2:
					user.Teacher_Info_List()
				case 3:
					user.Find_Teacher_Info()
				case 4:
					user.Add_Teacher_Info()
					model.Teacher{}.Write_To_File(TEACHER_DATA_PATH)
				case 5:
					user.Update_Teacher_Info()
					model.Teacher{}.Write_To_File(TEACHER_DATA_PATH)

				case 6:
					ok := user.Delete_Teacher_Info()
					if ok {
						model.Teacher{}.Write_To_File(TEACHER_DATA_PATH)

						fmt.Printf("删除成功!\n")
					} else {
						fmt.Printf("删除失败!\n")

					}
				case 0:
					goto main_loop
				}
			}
		} else {
			fmt.Printf("%v\n", msg)
			goto main_loop
		}
	default:
		fmt.Printf("账号不存在...")
		goto main_loop
	}

}
