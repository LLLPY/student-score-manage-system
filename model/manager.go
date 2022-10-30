package model

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"student-score-manage-system/utils"
)

// 管理者
type Manager struct {
	Num       string //职工号
	Name      string //姓名
	Birthday  string //出生日期
	Gender    int    //性别 0：男 1：女
	User_type int    //用户类型 1：学生 2：教师 3：管理员
	Password  string //密码
}

// 将文件中的数据读到缓冲区，用的时候直接从缓冲区取，而不需要每次都从文件中读取
func (manager Manager) Read_to_buffer(filename string) (err error) {

	if len(MANAGER_BUF) == 0 {
		content, err := os.ReadFile(filename)
		if err != nil {
			path_list := strings.Split(filename, "/")
			dirs := path_list[:len(path_list)-1]
			os.MkdirAll(strings.Join(dirs, "/"), os.ModePerm) //创建目录
			os.Create(filename)                               //创建文件
		} else {
			MANAGER_BUF = make(map[string]Manager)
			data_list := strings.Split(string(content), "\n")
			for _, v := range data_list {
				v_list := strings.Split(v, ",")
				if len(v_list) == 6 {
					num := v_list[0]
					name := v_list[1]
					birthday := v_list[2]
					gender, _ := strconv.Atoi(v_list[3])
					user_type, _ := strconv.Atoi(v_list[4])
					password := v_list[5]
					manager := Manager{Num: num, Name: name, Birthday: birthday, Gender: gender, User_type: user_type, Password: password}
					MANAGER_BUF[num] = manager

				}
			}
		}
	}
	//初始化一个管理员
	MANAGER_BUF["2022000013"] = Manager{
		Name:      "root",
		Num:       "2022000013",
		Gender:    0,
		User_type: 3,
		Birthday:  "2022-10-12",
		Password:  "1234",
	}
	return nil

}

// 登录
func (manager Manager) Login(num string, password string) (msg string, ok bool) {
	fmt.Printf("MANAGER_BUF: %v\n", MANAGER_BUF)
	ok = true
	//账号校验
	if len(num) != 10 {
		ok = false
		msg = "账号长度不合法..."
	}
	year, _ := strconv.Atoi(num[:4])
	if year < 2022 {
		ok = false
		msg = "账号不合法..."
	}

	user_type := string(num[len(num)-1])
	if user_type != "1" && user_type != "2" && user_type != "3" {
		ok = false
		msg = "用户类型不合法..."
	}

	s, exit := MANAGER_BUF[num]
	if !exit {
		ok = false
		msg = "账号不存在..."
	} else {
		ok = s.Password == password
		if ok == false {
			msg = "密码不正确..."
		}
	}
	return msg, ok

}

// 显示个人信息
func (manager Manager) Show_Persional_info() {
	name := manager.Name
	num := manager.Num
	gender := Gender_Mapping[manager.Gender]
	user_type := User_Type_Mapping[manager.User_type]
	birthday := manager.Birthday
	fmt.Printf("\n===================个人信息===================\n")
	fmt.Printf("#   %s%s%-20s\n", "姓名："+name, strings.Repeat("　", (18-len([]byte(name)))/3), "职工号："+num)
	fmt.Printf("#   %s%s%-20s\n", "性别："+gender, strings.Repeat("　", (18-len([]byte(gender)))/3), "用户类型："+user_type)
	fmt.Printf("#   %s\n", "出生日期："+birthday)
	fmt.Printf("==============================================\n")
	utils.Legal_input_string("按任意键继续...", map[string]string{})
}

//打印教师信息列表
func Show_Teacher_Info(teacher_list []Teacher) {
	fmt.Printf("\n%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s\n",
		"姓名", strings.Repeat("　", (20-len([]byte("姓名")))/3),
		"学号", strings.Repeat("　", (20-len([]byte("学号")))/3),
		"专业", strings.Repeat("　", (25-len([]byte("专业")))/3),
		"班级", strings.Repeat("　", (15-len([]byte("班级")))/3),
		"性别", strings.Repeat("　", (12-len([]byte("性别")))/3),
		"出生日期", strings.Repeat("　", (20-len([]byte("出生日期")))/3),
	)
	for _, v := range teacher_list {
		fmt.Printf("%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s\n",
			v.Name, strings.Repeat("　", (20-len([]byte(v.Name)))/3),
			v.Num, strings.Repeat("　", (15-len([]byte(v.Num)))/3),
			v.Major, strings.Repeat("　", (25-len([]byte(v.Major)))/3),
			v.Class, strings.Repeat("　", (15-len([]byte(v.Class)))/3),
			Gender_Mapping[v.Gender], strings.Repeat("　", (10-len([]byte(Gender_Mapping[v.Gender])))/3),
			v.Birthday, strings.Repeat("　", (20-len([]byte(v.Birthday)))/3),
		)
	}
	fmt.Printf("\n")
}

//教师信息列表
func (manager Manager) Teacher_Info_List() {
	var teacher_list []Teacher
	for _, v := range TEACHER_BUF {

		teacher_list = append(teacher_list, v)
	}
	Show_Teacher_Info(teacher_list)

}

//查询教师信息
func (manager Manager) Find_Teacher_Info() {
	var num, name string
	find_choice := utils.Legal_input_int("\n1.学号查找 2.姓名查找\n请选择查询方式:", map[int]string{1: "学号", 2: "姓名"})
	if find_choice == 1 {
		fmt.Printf("请输入学号：")
		fmt.Scanln(&num)
	} else {
		fmt.Printf("请输入姓名：")
		fmt.Scanln(&name)
	}

	var teacher_list []Teacher
	if num != "" {
		fmt.Printf("\n学号: %v\n", num)
		s, ok := TEACHER_BUF[num]
		if ok {
			teacher_list = append(teacher_list, s)
		}
	} else {
		fmt.Printf("\n姓名: %v\n", name)
		for _, v := range TEACHER_BUF {
			if strings.Index(v.Name, name) != -1 {
				teacher_list = append(teacher_list, v)
			}
		}
	}
	if len(teacher_list) == 0 {
		fmt.Printf("无相关信息...\n")
	} else {
		Show_Teacher_Info(teacher_list)
	}
}

// 新增教师信息
func (manager Manager) Add_Teacher_Info() bool {
	var num, name, major, class, birthday, password string
	var gender int
	//学号
	num = utils.Legal_input_string("请输入工号：", map[string]string{})
	_, ok := TEACHER_BUF[num]
	if ok {
		fmt.Printf("该工号已存在！")
		return false
	}
	name = utils.Legal_input_string("请输入姓名：", map[string]string{})
	class = utils.Legal_input_string("请输入新的班级:", map[string]string{})
	major = Major_Mapping[utils.Legal_input_int("请输入新的专业:", Major_Mapping)]
	birthday = utils.Legal_input_string("请输入生日：", map[string]string{})
	for i := 0; i < len(Gender_Mapping); i++ {
		print(strconv.Itoa(i+1) + "." + Gender_Mapping[i] + "  ")
	}
	gender = utils.Legal_input_int("请输入性别：", map[int]string{1: "男", 2: "女"})
	gender--

	password = utils.Legal_input_string("请输入密码：", map[string]string{})
	s := Teacher{
		Name:     name,
		Num:      num,
		Major:    major,
		Class:    class,
		Birthday: birthday,
		Gender:   gender,
		Password: password,
	}

	TEACHER_BUF[num] = s
	return true
}

//更新教师信息
func (manager Manager) Update_Teacher_Info() {
	num := utils.Legal_input_string("请输入要修改的教师的工号：", map[string]string{})
	_, ok := TEACHER_BUF[num]
	if ok == false {
		fmt.Printf("无此工号的相关信息...\n")
	} else {
		teacher := TEACHER_BUF[num]
		name := utils.Legal_input_string("请输入新的姓名(旧："+teacher.Name+",不修改请直接回车):", map[string]string{})
		if name != "" {
			teacher.Name = name
		}
		for i := 1; ; i++ {
			s, ok := Major_Mapping[i]
			if ok {
				fmt.Printf("%v.%s ", i, s)
				if i%8 == 0 {
					fmt.Printf("\n")
				}
			} else {
				break
			}

		}
		major := utils.Legal_input_int("请输入新的专业(旧："+teacher.Major+"):", Major_Mapping)
		teacher.Major = Major_Mapping[major]
		class := utils.Legal_input_string("请输入新的班级(旧："+teacher.Class+",不修改请直接回车):", map[string]string{})
		if class != "" {
			teacher.Class = class
		}

		gender := utils.Legal_input_int("[1.男 2.女] 请输入新的性别(旧："+Gender_Mapping[teacher.Gender]+"):", map[int]string{1: "男", 2: "女"})
		gender--
		teacher.Gender = gender
		birthday := utils.Legal_input_string("请输入新的生日(旧："+teacher.Birthday+",不修改请直接回车):", map[string]string{})
		if birthday != "" {
			teacher.Birthday = birthday
		}
		TEACHER_BUF[num] = teacher

	}

}

// 删除教师信息
func (manager Manager) Delete_Teacher_Info() bool {
	num := utils.Legal_input_string("请输入要删除的教师的工号：", map[string]string{})
	_, ok := TEACHER_BUF[num]
	if ok == false {
		fmt.Printf("无此工号的相关信息...\n")
	} else {
		Show_Teacher_Info([]Teacher{TEACHER_BUF[num]})
		confirm := utils.Legal_input_int("是否确认删除([1.是 2:否]):", map[int]string{1: "是", 2: "否"})
		if confirm == 1 {
			//删除TEACHER_BUF中的内容
			delete(TEACHER_BUF, num)
			return true

		} else {
			return false
		}
	}
	return false

}
