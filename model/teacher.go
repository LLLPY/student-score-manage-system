package model

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"student-score-manage-system/utils"
)

// 老师
type Teacher struct {
	Num       string //职工号
	Name      string //姓名
	Major     string //专业
	Birthday  string //出生日期
	Gender    int    //性别 0：男 1：女
	User_type int    //用户类型 1：学生 2：教师 3：管理员
	Class     string //管理的班级
	Password  string //密码

}

// 将文件中的数据读到缓冲区，用的时候直接从缓冲区取，而不需要每次都从文件中读取
func (teacher Teacher) Read_to_buffer(filename string) (err error) {

	if len(TEACHER_BUF) == 0 {
		b, err := os.ReadFile(filename)
		// print("读了文件")
		if err != nil {
			fmt.Printf("教师信息读取失败: %v\n", err)
			return err
		} else {
			TEACHER_BUF = make(map[string]Teacher)
			data_list := strings.Split(string(b), "\n")
			for _, v := range data_list {
				v_list := strings.Split(v, ",")
				num := v_list[0]
				name := v_list[1]
				major := v_list[2]
				birthday := v_list[3]
				gender, _ := strconv.Atoi(v_list[4])
				user_type, _ := strconv.Atoi(v_list[5])
				class := v_list[6]
				password := v_list[7]
				TEACHER_BUF[num] = Teacher{Num: num, Name: name, Major: major, Birthday: birthday, Gender: gender, User_type: user_type, Class: class, Password: password}
			}
		}
		return nil
	}
	// print("没有读文件")
	return nil

}

//登录
func (teacher Teacher) Login(num string, password string) (ok bool) {
	s, ok := TEACHER_BUF[num]
	return s.Password == password
}

// 显示个人信息
func (teacher Teacher) Show_info() {
	name := teacher.Name
	num := teacher.Num
	major := teacher.Major
	class := teacher.Class
	gender := Gender_Mapping[teacher.Gender]
	user_type := User_Type_Mapping[teacher.User_type]
	birthday := teacher.Birthday
	fmt.Printf("teacher: %v\n", teacher)
	fmt.Printf("\n===================个人信息===================\n")
	fmt.Printf("#   %s%s%-20s\n", "姓名："+name, strings.Repeat("　", (18-len([]byte(name)))/3), "职工号："+num)
	fmt.Printf("#   %s%s%-20s\n", "专业："+major, strings.Repeat("　", (18-len([]byte(major)))/3), "班级："+class)
	fmt.Printf("#   %s%s%-20s\n", "性别："+gender, strings.Repeat("　", (18-len([]byte(gender)))/3), "用户类型："+user_type)
	fmt.Printf("#   %s\n", "出生日期："+birthday)
	fmt.Printf("==============================================\n")
	utils.Legal_input_string("按任意键继续...", map[string]string{})
}

//学生成绩列表
func (teacher Teacher) Student_Score_List() {

	var score_list []Score
	for _, v := range SCORE_BUF {
		student := STUDENT_BUF[v.Num]
		if student.Major == teacher.Major && student.Class == teacher.Class {
			score_list = append(score_list, v)
		}

	}
	sort_set := map[int]string{
		1: "默认不排序",
		2: "按照语文成绩",
		3: "按照数学成绩",
		4: "按照英语成绩",
		5: "按照物理成绩",
		6: "按照化学成绩",
		7: "按照生物成绩",
		8: "按照体育成绩",
		9: "按照总分成绩",
	}
	for i := 1; i <= len(sort_set); i++ {
		fmt.Printf("%v.%s ", i, sort_set[i])

	}
	sort_choice := utils.Legal_input_int("\n请选择排序的方式:", sort_set)

	switch sort_choice {
	case 1:
	case 2:
		sort.Sort(ByChinese(score_list))
	case 3:
		sort.Sort(ByMath(score_list))
	case 4:
		sort.Sort(ByEnglish(score_list))
	case 5:
		sort.Sort(ByPhysical(score_list))
	case 6:
		sort.Sort(ByChemistry(score_list))
	case 7:
		sort.Sort(ByBiology(score_list))
	case 8:
		sort.Sort(BySports(score_list))
	case 9:
		sort.Sort(BySum(score_list))
	}
	Show_score_list(score_list)

}

//查询学生成绩
func (teacher Teacher) Search_Student_Score() {
	Student{}.Find()
}

//成绩分析
func (teacher Teacher) Analyse_Class_Score() {

	//计算自己班的各科成绩的平均分

	var  chinese_sum_mapping, math_sum_mapping, english_sum_mapping, physical_sum_mapping, chemistry_sum_mapping, biology_sum_mapping, sports_sum_mapping,all_sum_mapping map[int]int{1:0,2:0,3:0,4:0,5:0,6:0,7:0,8:0}
	n := 0


	for _, v := range SCORE_BUF {
		student := STUDENT_BUF[v.Num]
		if student.Major == teacher.Major && student.Class == teacher.Class {
			switch student.Semester {
			case 1:
				chinese_sum_mapping[student.Semester] += v.Chinese
				n += 1 //班级人数加一

			case 2:
				math_sum_mapping[student.Semester] += v.Math
			case 3:
				english_sum_mapping[student.Semester] += v.English
			case 4:
				physical_sum_mapping[student.Semester] += v.Physical
			case 5:
				chemistry_sum_mapping[student.Semester] += v.Chemistry
			case 6:
				biology_sum_mapping[student.Semester] += v.Biology
			case 7:
				sports_sum_mapping[student.Semester] += v.Sports
			}

		}

	}
	for i := 1; i <= len(Semester_List); i++ {
		all_sum_mapping[i] = chinese_sum_mapping[i] + math_sum_mapping[i] + english_sum_mapping[i] + physical_sum_mapping[i] + chemistry_sum_mapping[i] + biology_sum_mapping[i] + sports_sum_mapping[i]

	}

	var s1, s2, s3, s4, s5, s6, s7, s8, s9 []string

	subject_choice := utils.Legal_input_int("1.语文 2.数学 3.英语 4.物理 5.化学 6.生物 7.体育 8.总分\n请选择你要分析的内容:", map[int]string{1: "语文", 2: "数学", 3: "英语", 4: "物理", 5: "化学", 6: "生物", 7: "体育", 8: "总分"})
	switch subject_choice {
	case 1:
		s1 = get_pillar(Semester_List[0], chinese_sum_mapping[1]/n, "语文")
		s2 = get_pillar(Semester_List[1], math_sum_mapping[2]/n, "语文")
		s3 = get_pillar(Semester_List[2], english_sum_mapping[3]/n, "语文")
		s4 = get_pillar(Semester_List[3], physical_sum_mapping[4]/n, "语文")
		s5 = get_pillar(Semester_List[4], chemistry_sum_mapping[5]/n, "语文")
		s6 = get_pillar(Semester_List[5], biology_sum_mapping[6]/n, "语文")
		s7 = get_pillar(Semester_List[6], sports_sum_mapping[7]/n, "语文")
		s8 = get_pillar(Semester_List[7], all_sum_mapping[8]/n, "语文")
		s9 = get_pillar("   语文     ", 0, "语文")

	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:

	}

	//打印柱状图
	fmt.Printf("\n")
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Printf("%v%v%v%v%v%v%v%v%v\n",
			s1[i], s2[i], s3[i], s4[i], s5[i], s6[i], s7[i], s8[i], s9[i],
		)
	}
	fmt.Printf("\n")

}

// 新增
func (teacher Teacher) Create() (err error) {
	var num, name, major, class, birthday, password string
	var gender, semester int
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
	major = teacher.Major
	class = teacher.Class
	fmt.Print("请输入生日：")
	fmt.Scanln(&birthday)
	for i := 0; i < len(Gender_Mapping); i++ {
		print(strconv.Itoa(i+1) + "." + Gender_Mapping[i] + "  ")

	}
	fmt.Print("\n请输入性别：")
	fmt.Scanln(&gender)
	gender--
	for i := 0; i < len(Semester_Mapping); i++ {
		print(strconv.Itoa(i+1) + "." + Semester_Mapping[i+1] + "  ")

	}
	semester = utils.Legal_input_int("\n请输入学期：", Semester_Mapping)
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

// 更新
func (student Student) Update() {}

// 更新学生信息
func (student *Student) Update_info(name string, major string, class string, birthday string, gender int, grade int, password string, score_list []Score) {

}

// 删除
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
