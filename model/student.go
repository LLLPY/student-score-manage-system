package model

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Score struct {
	Num       string //学号
	Chinese   uint64 //语文
	Math      uint64 //数学
	English   uint64 //英语
	Physical  uint64 //物理
	Chemistry uint64 //化学
	Biology   uint64 //生物
	Sports    uint64 //体育
	Semester  uint64 //学期 1,2,3,4,5,6,7,8 从大一到大四总共8个学期
}

//学生
type Student struct {
	Num        string  //学号 学号的组成：年份 2022 编号 00001（00001到99999） 用户类型：1（1，2，3分别代表学生，教师和管理员）--》2022000011
	Name       string  //姓名
	Major      string  //专业
	Class      string  //班级
	Birthday   string  //出生日期
	Gender     uint64  //性别 0：男 1：女
	Semester   uint64  //年级 1：大一 2：大二 3：大三 4：大四
	User_type  uint64  //用户类型 1：学生 2：教师 3：管理员
	Password   string  //密码
	Score_list []Score //成绩列表
}

//将学生的成绩读取到缓冲
func (score Score) Read_to_buffer(filename string) (err error) {
	b, err2 := os.ReadFile(filename)
	if err2 != nil {
		fmt.Printf("读取成绩失败: %v\n", err2)
	}
	data_list := strings.Split(string(b), "\n")
	for _, v := range data_list {
		v_list := strings.Split(v, ",")
		num := v_list[0]
		chinese, _ := strconv.ParseUint(v_list[1], 10, 64)
		math, _ := strconv.ParseUint(v_list[2], 10, 64)
		english, _ := strconv.ParseUint(v_list[3], 10, 64)
		physical, _ := strconv.ParseUint(v_list[4], 10, 64)
		chemistry, _ := strconv.ParseUint(v_list[5], 10, 64)
		biology, _ := strconv.ParseUint(v_list[6], 10, 64)
		sports, _ := strconv.ParseUint(v_list[7], 10, 64)
		semester, _ := strconv.ParseUint(v_list[8], 10, 64)
		tmp_score := Score{Num: num, Chinese: chinese, Math: math, English: english, Physical: physical, Chemistry: chemistry, Biology: biology, Sports: sports, Semester: semester}
		SCORE_BUF = append(SCORE_BUF, tmp_score)
		fmt.Printf("tmp_score: %v\n", tmp_score)
	}
	return nil
}

//将文件中的数据读到缓冲区，用的时候直接从缓冲区取，而不需要每次都从文件中读取
func (student Student) Read_to_buffer(filename string) (err error) {

	if len(STUDENT_BUF) == 0 {
		b, err := os.ReadFile(filename)
		// print("读了文件")
		if err != nil {
			fmt.Printf("学生信息读取失败: %v\n", err)
			return err
		} else {
			STUDENT_BUF = make(map[string]Student)
			data_list := strings.Split(string(b), "\n")
			for _, v := range data_list {
				v_list := strings.Split(v, ",")
				num := v_list[0]
				name := v_list[1]
				major := v_list[2]
				class := v_list[3]
				birthday := v_list[4]
				gender, _ := strconv.ParseUint(v_list[5], 10, 64)
				semester, _ := strconv.ParseUint(v_list[6], 10, 64)
				user_type, _ := strconv.ParseUint(v_list[7], 10, 64)
				password := v_list[8]

				STUDENT_BUF[num] = Student{Num: num, Name: name, Major: major, Class: class, Birthday: birthday, Gender: gender, Semester: semester, User_type: user_type, Password: password}
				fmt.Printf("STUDENT_BUF[num]: %v\n", STUDENT_BUF[num])
			}
		}

		return nil
	}
	// print("没有读文件")
	return nil

}

//登录
func (student Student) Login(num string, password string) (ok bool) {
	s, ok := STUDENT_BUF[num]
	return s.Password == password
}

//登出
func (student Student) Logout() (err error) {
	student = Student{}

	return nil
}

//显示个人信息
func (student Student) Show_info() {
	fmt.Printf("\n%-32v\n", "===================个人信息====================")
	fmt.Printf("#   姓名：%-10v   学号：%-15s#\n", student.Name, student.Num)
	fmt.Printf("#   专业：%-10v班级：%-15s#\n", student.Major, student.Class)
	fmt.Printf("#   年级：%-10v  用户类型：%-9s#\n", Semester_Mapping[student.Semester], User_Type_Mapping[student.User_type])
	fmt.Printf("#   性别：%-10v    出生日期：%-11s#\n", Gender_Mapping[student.Gender], student.Birthday)
	fmt.Printf("===============================================\n")
	var a string
	fmt.Println("按任意键继续...")
	fmt.Scanln(&a)
}

//根据账号和密码获取一个学生
func Get_Student_By_Num(num string) (student *Student) {
	*student = STUDENT_BUF[num]
	return
}

//新增
func (student Student) Create() (err error) {
	var num, name, major, class, birthday, password string
	var gender, semester, major_choice uint64
	//姓名
	fmt.Print("请输入学号：")
	fmt.Scanln(&num)
	_, ok := STUDENT_BUF[num]
	if ok {
		fmt.Printf("该学号已存在！")
		return
	}
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	for i := 0; i < len(Major_Mapping); i++ {
		print(strconv.Itoa(i+1) + "." + Major_Mapping[uint64(i+1)] + "  ")
	}
	fmt.Printf("\n请输入专业：")
	fmt.Scanln(&major_choice)
	major = Major_Mapping[major_choice]
	fmt.Print("请输入班级：")
	fmt.Scanln(&class)
	fmt.Print("请输入生日：")
	fmt.Scanln(&birthday)
	for i := 0; i < len(Gender_Mapping); i++ {
		print(strconv.Itoa(i+1) + "." + Gender_Mapping[uint64(i)] + "  ")

	}
	fmt.Print("\n请输入性别：")
	fmt.Scanln(&gender)
	gender--
	for i := 0; i < len(Semester_Mapping); i++ {
		print(strconv.Itoa(i+1) + "." + Semester_Mapping[uint64(i+1)] + "  ")

	}
	fmt.Print("\n请输入学期：")
	fmt.Scanln(&semester)
	fmt.Print("请输入密码：")
	fmt.Scanln(&password)
	s := Student{}
	s.Name = name
	s.Num = num
	s.Major = major
	s.Class = class
	s.Birthday = birthday
	s.Gender = gender
	s.Semester = semester
	s.Password = password
	STUDENT_BUF[num] = s
	return nil
}

//查找学生成绩
func (student Student) Find() {
	var num, name, find_choice string

	fmt.Printf("\n1.学号查找 other.姓名查找\n")
	fmt.Printf("请选择查找方式：")
	fmt.Scanln(&find_choice)
	if find_choice == "1" {
		fmt.Printf("请输入学号：")
		fmt.Scanln(&num)
	} else {
		fmt.Printf("请输入姓名：")
		fmt.Scanln(&name)
	}

	var score_list []Score
	if num != "" {
		score_list = Find_by_Num(num)
		fmt.Printf("\n学号：%v\n", num)

	} else {
		score_list = Find_by_Name(name)
		fmt.Printf("\n姓名：%v\n", name)

	}

	if len(score_list) < 1 {
		fmt.Printf("\n无相关内容...\n")
	} else {
		fmt.Printf("\n%-6v%-10v%-3v%-3v%-3v%-3v%-3v%-3v%-3v\n", "姓名", "学号", "语文", "数学", "英语", "物理", "生物", "化学", "体育")
		for _, v := range score_list {
			fmt.Printf("%-6v%-12v%-5v%-5v%-5v%-5v%-5v%-5v%-5v\n", STUDENT_BUF[v.Num].Name, v.Num, v.Chinese, v.Math, v.English, v.Physical, v.Biology, v.Chemistry, v.Sports)

		}
		fmt.Printf("\n")
	}
	var a string
	fmt.Println("按任意键继续...")
	fmt.Scanln(&a)

}

//根据学号查找
func Find_by_Num(num string) (score_list []Score) {

	for _, v := range SCORE_BUF {
		if v.Num == num {
			score_list = append(score_list, v)
		}
	}
	return
}

//根据姓名查找
func Find_by_Name(name string) (score_list []Score) {

	for _, v := range SCORE_BUF {
		if strings.Index(STUDENT_BUF[v.Num].Name, name) != -1 {
			score_list = append(score_list, v)
		}
	}
	return
}

//成绩PK
func (student Student) Score_Pk() {

	n := 1
	n_num_mapping := make(map[int]string, len(STUDENT_BUF))
	for k, v := range STUDENT_BUF {
		fmt.Printf("(%v.%v)  ", n, v.Name)
		if n%8 == 0 {
			fmt.Printf("\n")
		}
		n_num_mapping[n] = k
		n += 1
	}
	fmt.Printf("\n请选择要进行PK的同学：")
	fmt.Scanln(&n)

	semester_choice := 1

	for i := 1; i <= 8; i++ {
		fmt.Printf("(%v.%v)  ", i, Semester_Mapping[uint64(i)])

	}
	fmt.Printf("\n请选择学期：")
	fmt.Scanln(&semester_choice)
	fmt.Printf("\n")
	//找出自己和要PK的同学对应学期的成绩
	var self_score, other_score Score
	var flag int = 0
	for _, v := range SCORE_BUF {
		if v.Num == student.Num && v.Semester == uint64(semester_choice) {
			self_score = v
			flag++
		}
		if v.Num == n_num_mapping[n] && v.Semester == uint64(semester_choice) {
			other_score = v
			flag++
		}
		if flag == 2 {
			break
		}
	}

	//PK比较
	Pk("语文", self_score.Chinese, other_score.Chinese, 100)
	Pk("数学", self_score.Math, other_score.Math, 100)
	Pk("英语", self_score.English, other_score.English, 100)
	Pk("物理", self_score.Physical, other_score.Physical, 100)
	Pk("化学", self_score.Chemistry, other_score.Chemistry, 100)
	Pk("生物", self_score.Biology, other_score.Biology, 100)
	Pk("体育", self_score.Sports, other_score.Sports, 100)
	sum1 := self_score.Chinese + self_score.Math + self_score.English + self_score.Physical + self_score.Chemistry + self_score.Biology + self_score.Sports
	sum2 := other_score.Chinese + other_score.Math + other_score.English + other_score.Physical + other_score.Chemistry + other_score.Biology + other_score.Sports
	Pk("总分", sum1, sum2, 700)

}

//成绩比较的功能封装
func Pk_Part(course string, score1 uint64, score2 uint64, full_score uint64, flag string) {

	per1 := float64(float64(score1) / float64(full_score))
	per2 := float64(float64(score2) / float64(full_score))
	fmt.Printf("%v:我", course)
	for i := 0; i < int(per1*100)/2; i++ {
		fmt.Printf("#")

	}
	for i := 0; i < (50 - int(per1*100)/2); i++ {
		fmt.Printf("-")
	}

	fmt.Printf("%v分  ta:", score1)
	for i := 0; i < int(per2*100)/2; i++ {
		fmt.Printf("#")

	}
	for i := 0; i < (50 - int(per2*100)/2); i++ {
		fmt.Printf("-")
	}

	var per float64

	if score1 < score2 {
		fmt.Printf("%v分  ta超越了%v分:", score2, score2-score1)
		per = float64((score2 - score1)) / float64(score1)
		if score1 == 0 {
			per = 0 //除数可能为0
		}

	} else {
		fmt.Printf("%v分  我超越了%v分:", score2, score1-score2)
		per = float64((score1 - score2)) / float64(score2)
		if score2 == 0 {
			per = 0
		}

	}
	var tmp_per int = int(per)
	if per > 1 {
		tmp_per = 1 //占比可能超过100%
	}
	//计算百分比
	for i := 0; i < int(tmp_per*100)/2; i++ {
		fmt.Printf("#")

	}
	for i := 0; i < (50 - int(tmp_per*100)/2); i++ {
		fmt.Printf("-")
	}
	fmt.Printf("%.2f%%  占比(%.2f%%),%v!\n\n", per*100, per*100, flag)
}

//成绩PK
func Pk(course string, score1 uint64, score2 uint64, full_score uint64) {
	if score1 > score2 {
		Pk_Part(course, score1, score2, full_score, "完胜")

	} else if score1 < score2 {
		Pk_Part(course, score1, score2, full_score, "完败")

	} else {
		Pk_Part(course, score1, score2, full_score, "可敬的的对手")

	}
}

//更新
func (student Student) Update() {}

//更新学生信息
func (student *Student) Update_info(name string, major string, class string, birthday string, gender uint64, grade uint64, password string, score_list []Score) {

}

//删除
func (student Student) Delete() (err error) {
	//删除STUDENT_BUF中的内容
	delete(STUDENT_BUF, student.Num)

	//删除SCORE_BUF中的内容
	for i := 0; i < len(SCORE_BUF); i++ {
		if SCORE_BUF[i].Num == student.Num {
			SCORE_BUF = append(SCORE_BUF[:i], SCORE_BUF[i+1:]...)
			i--
		}

	}
	return nil

}
