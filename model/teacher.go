package model

import (
	"fmt"
	"math"
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
		content, err := os.ReadFile(filename)
		// print("读了文件")
		if err != nil {
			path_list := strings.Split(filename, "/")
			dirs := path_list[:len(path_list)-1]
			os.MkdirAll(strings.Join(dirs, "/"), os.ModePerm) //创建目录
			os.Create(filename)                               //创建文件
		} else {
			TEACHER_BUF = make(map[string]Teacher)
			data_list := strings.Split(string(content), "\n")
			for _, v := range data_list {
				v_list := strings.Split(v, ",")
				if len(v_list) == 8 {
					num := v_list[0]
					name := v_list[1]
					major := v_list[2]
					birthday := v_list[3]
					gender, _ := strconv.Atoi(v_list[4])
					user_type, _ := strconv.Atoi(v_list[5])
					class := v_list[6]
					password := v_list[7]
					teacher := Teacher{
						Num:       num,
						Name:      name,
						Major:     major,
						Birthday:  birthday,
						Gender:    gender,
						User_type: user_type,
						Class:     class,
						Password:  password,
					}

					TEACHER_BUF[num] = teacher
				}

			}
		}
		return nil
	}
	// print("没有读文件")
	return nil

}

//写入文件
func (teacher Teacher) Write_To_File(filename string) {
	teacher_data := ""
	for _, v := range TEACHER_BUF {
		num := v.Num
		name := v.Name
		major := v.Major
		birthday := v.Birthday
		gender := strconv.Itoa(v.Gender)
		user_type := strconv.Itoa(v.User_type)
		class := v.Class
		password := v.Password

		teacher := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n", num, name, major, birthday, gender, user_type, class, password)
		teacher_data += teacher
	}
	os.WriteFile(filename, []byte(teacher_data), os.ModePerm)

}

//登录
func (teacher Teacher) Login(num string, password string) (msg string, ok bool) {
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

	s, exit := TEACHER_BUF[num]
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
func (teacher Teacher) Show_Persional_info() {
	name := teacher.Name
	num := teacher.Num
	major := teacher.Major
	class := teacher.Class
	gender := Gender_Mapping[teacher.Gender]
	user_type := User_Type_Mapping[teacher.User_type]
	birthday := teacher.Birthday
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

//计算平均分和方差
func cal_avage_square(score_map map[int]int) {

}

//成绩分析
func (teacher Teacher) Analyse_Class_Score() {

	//计算自己班的各科成绩的平均分

	chinese_sum_mapping := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	math_sum_mapping := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	english_sum_mapping := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	physical_sum_mapping := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	chemistry_sum_mapping := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	biology_sum_mapping := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	sports_sum_mapping := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	all_sum_mapping := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	n := 0

	for _, s := range STUDENT_BUF {
		if s.Major == teacher.Major && s.Class == teacher.Class {
			n++ //统计班级人数
		}
	}
	for _, v := range SCORE_BUF {
		student := STUDENT_BUF[v.Num]
		if student.Major == teacher.Major && student.Class == teacher.Class {
			chinese_sum_mapping[v.Semester] += v.Chinese
			math_sum_mapping[v.Semester] += v.Math
			english_sum_mapping[v.Semester] += v.English
			physical_sum_mapping[v.Semester] += v.Physical
			chemistry_sum_mapping[v.Semester] += v.Chemistry
			biology_sum_mapping[v.Semester] += v.Biology
			sports_sum_mapping[v.Semester] += v.Sports
		}

	}

	for i := 1; i <= len(Semester_List); i++ {
		all_sum_mapping[i] = chinese_sum_mapping[i] + math_sum_mapping[i] + english_sum_mapping[i] + physical_sum_mapping[i] + chemistry_sum_mapping[i] + biology_sum_mapping[i] + sports_sum_mapping[i]

	}

	var s1, s2, s3, s4, s5, s6, s7, s8, s9 []string
	var score_sum, score_avage, score_square, score_n int

	subject_choice := utils.Legal_input_int("1.语文 2.数学 3.英语 4.物理 5.化学 6.生物 7.体育 8.总分\n请选择你要分析的内容:", map[int]string{1: "语文", 2: "数学", 3: "英语", 4: "物理", 5: "化学", 6: "生物", 7: "体育", 8: "总分"})
	switch subject_choice {
	case 1:
		s1 = get_pillar(Semester_List[0], chinese_sum_mapping[1]/n, "语文")
		s2 = get_pillar(Semester_List[1], chinese_sum_mapping[2]/n, "语文")
		s3 = get_pillar(Semester_List[2], chinese_sum_mapping[3]/n, "语文")
		s4 = get_pillar(Semester_List[3], chinese_sum_mapping[4]/n, "语文")
		s5 = get_pillar(Semester_List[4], chinese_sum_mapping[5]/n, "语文")
		s6 = get_pillar(Semester_List[5], chinese_sum_mapping[6]/n, "语文")
		s7 = get_pillar(Semester_List[6], chinese_sum_mapping[7]/n, "语文")
		s8 = get_pillar(Semester_List[7], chinese_sum_mapping[8]/n, "语文")
		s9 = get_pillar("语文平均分", 0, "语文")

		for k, v := range chinese_sum_mapping {
			if chinese_sum_mapping[k] != 0 {
				score_sum += v / n
				score_n++
			}
		}

		score_avage = score_sum / score_n
		for k, v := range chinese_sum_mapping {
			if chinese_sum_mapping[k] != 0 {
				score_square += int(math.Pow(float64(v/n-score_avage), 2))
			}
		}

	case 2:
		s1 = get_pillar(Semester_List[0], math_sum_mapping[1]/n, "数学")
		s2 = get_pillar(Semester_List[1], math_sum_mapping[2]/n, "数学")
		s3 = get_pillar(Semester_List[2], math_sum_mapping[3]/n, "数学")
		s4 = get_pillar(Semester_List[3], math_sum_mapping[4]/n, "数学")
		s5 = get_pillar(Semester_List[4], math_sum_mapping[5]/n, "数学")
		s6 = get_pillar(Semester_List[5], math_sum_mapping[6]/n, "数学")
		s7 = get_pillar(Semester_List[6], math_sum_mapping[7]/n, "数学")
		s8 = get_pillar(Semester_List[7], math_sum_mapping[8]/n, "数学")
		s9 = get_pillar("数学平均分", 0, "数学")
		for k, v := range math_sum_mapping {
			if math_sum_mapping[k] != 0 {
				score_sum += v / n
				score_n++
			}
		}

		score_avage = score_sum / score_n
		for k, v := range math_sum_mapping {
			if math_sum_mapping[k] != 0 {
				score_square += int(math.Pow(float64(v/n-score_avage), 2))
			}
		}
	case 3:
		s1 = get_pillar(Semester_List[0], english_sum_mapping[1]/n, "英语")
		s2 = get_pillar(Semester_List[1], english_sum_mapping[2]/n, "英语")
		s3 = get_pillar(Semester_List[2], english_sum_mapping[3]/n, "英语")
		s4 = get_pillar(Semester_List[3], english_sum_mapping[4]/n, "英语")
		s5 = get_pillar(Semester_List[4], english_sum_mapping[5]/n, "英语")
		s6 = get_pillar(Semester_List[5], english_sum_mapping[6]/n, "英语")
		s7 = get_pillar(Semester_List[6], english_sum_mapping[7]/n, "英语")
		s8 = get_pillar(Semester_List[7], english_sum_mapping[8]/n, "英语")
		s9 = get_pillar("英语平均分", 0, "英语")
		for k, v := range english_sum_mapping {
			if english_sum_mapping[k] != 0 {
				score_sum += v / n
				score_n++
			}
		}

		score_avage = score_sum / score_n
		for k, v := range english_sum_mapping {
			if english_sum_mapping[k] != 0 {
				score_square += int(math.Pow(float64(v/n-score_avage), 2))
			}
		}
	case 4:
		s1 = get_pillar(Semester_List[0], physical_sum_mapping[1]/n, "物理")
		s2 = get_pillar(Semester_List[1], physical_sum_mapping[2]/n, "物理")
		s3 = get_pillar(Semester_List[2], physical_sum_mapping[3]/n, "物理")
		s4 = get_pillar(Semester_List[3], physical_sum_mapping[4]/n, "物理")
		s5 = get_pillar(Semester_List[4], physical_sum_mapping[5]/n, "物理")
		s6 = get_pillar(Semester_List[5], physical_sum_mapping[6]/n, "物理")
		s7 = get_pillar(Semester_List[6], physical_sum_mapping[7]/n, "物理")
		s8 = get_pillar(Semester_List[7], physical_sum_mapping[8]/n, "物理")
		s9 = get_pillar("物理平均分", 0, "物理")
		for k, v := range physical_sum_mapping {
			if physical_sum_mapping[k] != 0 {
				score_sum += v / n
				score_n++
			}
		}

		score_avage = score_sum / score_n
		for k, v := range physical_sum_mapping {
			if physical_sum_mapping[k] != 0 {
				score_square += int(math.Pow(float64(v/n-score_avage), 2))
			}
		}
	case 5:
		s1 = get_pillar(Semester_List[0], chemistry_sum_mapping[1]/n, "化学")
		s2 = get_pillar(Semester_List[1], chemistry_sum_mapping[2]/n, "化学")
		s3 = get_pillar(Semester_List[2], chemistry_sum_mapping[3]/n, "化学")
		s4 = get_pillar(Semester_List[3], chemistry_sum_mapping[4]/n, "化学")
		s5 = get_pillar(Semester_List[4], chemistry_sum_mapping[5]/n, "化学")
		s6 = get_pillar(Semester_List[5], chemistry_sum_mapping[6]/n, "化学")
		s7 = get_pillar(Semester_List[6], chemistry_sum_mapping[7]/n, "化学")
		s8 = get_pillar(Semester_List[7], chemistry_sum_mapping[8]/n, "化学")
		s9 = get_pillar("化学平均分", 0, "化学")
		for k, v := range chemistry_sum_mapping {
			if chemistry_sum_mapping[k] != 0 {
				score_sum += v / n
				score_n++
			}
		}

		score_avage = score_sum / score_n
		for k, v := range chemistry_sum_mapping {
			if chemistry_sum_mapping[k] != 0 {
				score_square += int(math.Pow(float64(v/n-score_avage), 2))
			}
		}
	case 6:
		s1 = get_pillar(Semester_List[0], biology_sum_mapping[1]/n, "生物")
		s2 = get_pillar(Semester_List[1], biology_sum_mapping[2]/n, "生物")
		s3 = get_pillar(Semester_List[2], biology_sum_mapping[3]/n, "生物")
		s4 = get_pillar(Semester_List[3], biology_sum_mapping[4]/n, "生物")
		s5 = get_pillar(Semester_List[4], biology_sum_mapping[5]/n, "生物")
		s6 = get_pillar(Semester_List[5], biology_sum_mapping[6]/n, "生物")
		s7 = get_pillar(Semester_List[6], biology_sum_mapping[7]/n, "生物")
		s8 = get_pillar(Semester_List[7], biology_sum_mapping[8]/n, "生物")
		s9 = get_pillar("生物平均分", 0, "生物")
		for k, v := range biology_sum_mapping {
			if biology_sum_mapping[k] != 0 {
				score_sum += v / n
				score_n++
			}
		}

		score_avage = score_sum / score_n
		for k, v := range biology_sum_mapping {
			if biology_sum_mapping[k] != 0 {
				score_square += int(math.Pow(float64(v/n-score_avage), 2))
			}
		}
	case 7:
		s1 = get_pillar(Semester_List[0], sports_sum_mapping[1]/n, "体育")
		s2 = get_pillar(Semester_List[1], sports_sum_mapping[2]/n, "体育")
		s3 = get_pillar(Semester_List[2], sports_sum_mapping[3]/n, "体育")
		s4 = get_pillar(Semester_List[3], sports_sum_mapping[4]/n, "体育")
		s5 = get_pillar(Semester_List[4], sports_sum_mapping[5]/n, "体育")
		s6 = get_pillar(Semester_List[5], sports_sum_mapping[6]/n, "体育")
		s7 = get_pillar(Semester_List[6], sports_sum_mapping[7]/n, "体育")
		s8 = get_pillar(Semester_List[7], sports_sum_mapping[8]/n, "体育")
		s9 = get_pillar("体育平均分", 0, "体育")
		for k, v := range sports_sum_mapping {
			if sports_sum_mapping[k] != 0 {
				score_sum += v / n
				score_n++
			}
		}

		score_avage = score_sum / score_n
		for k, v := range sports_sum_mapping {
			if sports_sum_mapping[k] != 0 {
				score_square += int(math.Pow(float64(v/n-score_avage), 2))
			}
		}
	case 8:
		s1 = get_pillar(Semester_List[0], all_sum_mapping[1]/n, "总分")
		s2 = get_pillar(Semester_List[1], all_sum_mapping[2]/n, "总分")
		s3 = get_pillar(Semester_List[2], all_sum_mapping[3]/n, "总分")
		s4 = get_pillar(Semester_List[3], all_sum_mapping[4]/n, "总分")
		s5 = get_pillar(Semester_List[4], all_sum_mapping[5]/n, "总分")
		s6 = get_pillar(Semester_List[5], all_sum_mapping[6]/n, "总分")
		s7 = get_pillar(Semester_List[6], all_sum_mapping[7]/n, "总分")
		s8 = get_pillar(Semester_List[7], all_sum_mapping[8]/n, "总分")
		s9 = get_pillar("总分平均分", 0, "总分")
		for k, v := range all_sum_mapping {
			if all_sum_mapping[k] != 0 {
				score_sum += v / n
				score_n++
			}
		}

		score_avage = score_sum / score_n
		for k, v := range all_sum_mapping {
			if all_sum_mapping[k] != 0 {
				score_square += int(math.Pow(float64(v/n-score_avage), 2))
			}
		}
	}
	s9[2] = "        "
	fmt.Printf("\n平均分: %v\n", score_avage)
	fmt.Printf("方  差: %v\n", score_square)
	//打印柱状图
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Printf("%v%v%v%v%v%v%v%v%v\n",
			s1[i], s2[i], s3[i], s4[i], s5[i], s6[i], s7[i], s8[i], s9[i],
		)
	}
	fmt.Printf("\n")

}

/////////////////////////管理班级////////////////////////////////////////////

//打印学生信息列表
func Show_Student_Info(student_list []Student) {
	fmt.Printf("\n%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s\n",
		"姓名", strings.Repeat("　", (20-len([]byte("姓名")))/3),
		"学号", strings.Repeat("　", (20-len([]byte("学号")))/3),
		"专业", strings.Repeat("　", (25-len([]byte("专业")))/3),
		"班级", strings.Repeat("　", (15-len([]byte("班级")))/3),
		"年级", strings.Repeat("　", (12-len([]byte("年级")))/3),
		"性别", strings.Repeat("　", (12-len([]byte("性别")))/3),
		"出生日期", strings.Repeat("　", (20-len([]byte("出生日期")))/3),
	)
	for _, v := range student_list {
		fmt.Printf("%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s%-s\n",
			v.Name, strings.Repeat("　", (20-len([]byte(v.Name)))/3),
			v.Num, strings.Repeat("　", (15-len([]byte(v.Num)))/3),
			v.Major, strings.Repeat("　", (25-len([]byte(v.Major)))/3),
			v.Class, strings.Repeat("　", (13-len([]byte(v.Class)))/3),
			Semester_Mapping[v.Semester], strings.Repeat("　", (15-len([]byte(Semester_Mapping[v.Semester])))/3),
			Gender_Mapping[v.Gender], strings.Repeat("　", (10-len([]byte(Gender_Mapping[v.Gender])))/3),
			v.Birthday, strings.Repeat("　", (20-len([]byte(v.Birthday)))/3),
		)
	}
	fmt.Printf("\n")
}

//学生信息列表
func (teacher Teacher) Student_Info_List() {
	var student_list []Student
	for _, v := range STUDENT_BUF {
		if v.Class == teacher.Class {
			student_list = append(student_list, v)
		}
	}
	Show_Student_Info(student_list)

}

//查询学生信息
func (teacher Teacher) Find_Student_Info() {
	var num, name string
	find_choice := utils.Legal_input_int("\n1.学号查找 2.姓名查找\n请选择查询方式:", map[int]string{1: "学号", 2: "姓名"})
	if find_choice == 1 {
		fmt.Printf("请输入学号：")
		fmt.Scanln(&num)
	} else {
		fmt.Printf("请输入姓名：")
		fmt.Scanln(&name)
	}

	var student_list []Student
	if num != "" {
		fmt.Printf("\n学号: %v\n", num)
		s, ok := STUDENT_BUF[num]
		if ok {
			student_list = append(student_list, s)
		}
	} else {
		fmt.Printf("\n姓名: %v\n", name)
		for _, v := range STUDENT_BUF {
			if v.Class == teacher.Class && strings.Index(v.Name, name) != -1 {
				student_list = append(student_list, v)
			}
		}
	}
	if len(student_list) == 0 {
		fmt.Printf("无相关信息...\n")
	} else {
		Show_Student_Info(student_list)
	}
}

// 新增学生信息
func (teacher Teacher) Add_Student_Info() bool {
	var num, name, major, class, birthday, password string
	var gender, semester int
	major = teacher.Major
	class = teacher.Class
	//学号
	num = utils.Legal_input_string("请输入学号：", map[string]string{})
	_, ok := STUDENT_BUF[num]
	if ok {
		fmt.Printf("该学号已存在！")
		return false
	}
	name = utils.Legal_input_string("请输入姓名：", map[string]string{})
	birthday = utils.Legal_input_string("请输入生日：", map[string]string{})
	for i := 0; i < len(Gender_Mapping); i++ {
		print(strconv.Itoa(i+1) + "." + Gender_Mapping[i] + "  ")
	}
	gender = utils.Legal_input_int("请输入性别：", map[int]string{1: "男", 2: "女"})
	gender--
	for i := 0; i < len(Semester_Mapping); i++ {
		print(strconv.Itoa(i+1) + "." + Semester_Mapping[i+1] + "  ")

	}
	semester = utils.Legal_input_int("\n请输入学期：", Semester_Mapping)
	password = utils.Legal_input_string("请输入密码：", map[string]string{})
	s := Student{
		Name:     name,
		Num:      num,
		Major:    major,
		Class:    class,
		Birthday: birthday,
		Gender:   gender,
		Semester: semester,
		Password: password,
	}

	STUDENT_BUF[num] = s
	return true
}

//更新学生信息
func (teacher Teacher) Update_Student_Info() {
	num := utils.Legal_input_string("请输入要修改的学生的学号：", map[string]string{})
	_, ok := STUDENT_BUF[num]
	if ok == false {
		fmt.Printf("无此学号的相关信息...\n")
	} else {
		student := STUDENT_BUF[num]
		name := utils.Legal_input_string("请输入新的姓名(旧："+student.Name+",不修改请直接回车):", map[string]string{})
		if name != "" {
			student.Name = name
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
		major := utils.Legal_input_int("请输入新的专业(旧："+student.Major+"):", Major_Mapping)
		student.Major = Major_Mapping[major]
		class := utils.Legal_input_string("请输入新的班级(旧："+student.Class+",不修改请直接回车):", map[string]string{})
		if class != "" {
			student.Class = class
		}
		for i := 1; ; i++ {
			s, ok := Semester_Mapping[i]
			if ok {
				fmt.Printf("%v.%s ", i, s)
				if i%8 == 0 {
					fmt.Printf("\n")
				}
			} else {
				break
			}

		}
		semester := utils.Legal_input_int("请输入新的年级(旧："+Semester_Mapping[student.Semester]+"):", Semester_Mapping)
		student.Semester = semester

		gender := utils.Legal_input_int("[1.男 2.女] 请输入新的性别(旧："+Gender_Mapping[student.Gender]+"):", map[int]string{1: "男", 2: "女"})
		gender--
		student.Gender = gender
		birthday := utils.Legal_input_string("请输入新的生日(旧："+student.Birthday+",不修改请直接回车):", map[string]string{})
		if birthday != "" {
			student.Birthday = birthday
		}
		STUDENT_BUF[num] = student

	}

}

// 删除学生信息
func (teacher Teacher) Delete_Student_Info() bool {
	num := utils.Legal_input_string("请输入要删除的学生的学号：", map[string]string{})
	_, ok := STUDENT_BUF[num]
	if ok == false {
		fmt.Printf("无此学号的相关信息...\n")
	} else {
		Show_Student_Info([]Student{STUDENT_BUF[num]})
		confirm := utils.Legal_input_int("是否确认删除([1.是 2:否]):", map[int]string{1: "是", 2: "否"})
		if confirm == 1 {
			//删除STUDENT_BUF中的内容
			delete(STUDENT_BUF, num)

			//删除SCORE_BUF中的内容
			for i := 0; i < len(SCORE_BUF); i++ {
				if SCORE_BUF[i].Num == num {
					SCORE_BUF = append(SCORE_BUF[:i], SCORE_BUF[i+1:]...)
					i--
				}
			}
			return true

		} else {
			return false
		}
	}
	return false

}
