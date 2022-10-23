package model

//user.txt用于存储所有用户（学生，教师和管理员）以及他们的类型标识

//定义一个用户的接口
type User interface {
	Read_to_buffer(string) error //将文件内容读到缓冲区
	Login(string, string) bool   //登录
	Logout() error               //登出
	Show_info()                  //打印信息
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
//学生
var STUDENT_BUF map[string]Student
var SCORE_BUF []Score
var Semester_List = []string{"大一上", "大一下", "大二上", "大二下", "大三上", "大三下", "大四上", "大四下"}
var Semester_Mapping = make(map[int]string, len(Semester_List))
var Gender_List = []string{"男", "女"}
var Gender_Mapping = make(map[int]string, len(Gender_List))
var User_Type_List = []string{"学生", "教师", "管理员"}
var User_Type_Mapping = make(map[int]string, len(User_Type_List))
var Major_List = []string{"信息与计算科学", "应用数学"}
var Major_Mapping = make(map[int]string, len(Major_List))

// 初始化
func init() {
	//初始化学期映射
	for i := 0; i < len(Semester_List); i++ {
		Semester_Mapping[i+1] = Semester_List[i]
	}
	//初始化性别映射
	for i := 0; i < len(Gender_List); i++ {
		Gender_Mapping[i] = Gender_List[i]

	}
	//初始化用户类型映射
	for i := 0; i < len(User_Type_List); i++ {
		User_Type_Mapping[i+1] = User_Type_List[i]
	}

	//初始化专业列表
	for i := 0; i < len(Major_List); i++ {
		Major_Mapping[i+1] = Major_List[i]
	}

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//教师
var TEACHER_BUF map[string]Teacher

//对成绩进行排序
type ByChinese []Score
type ByMath []Score
type ByEnglish []Score
type ByPhysical []Score
type ByChemistry []Score
type ByBiology []Score
type BySports []Score
type BySum []Score

//按照语文成绩排序
func (s ByChinese) Len() int           { return len(s) }
func (s ByChinese) Less(i, j int) bool { return s[i].Chinese > s[j].Chinese }
func (s ByChinese) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//按照数学成绩排序
func (s ByMath) Len() int           { return len(s) }
func (s ByMath) Less(i, j int) bool { return s[i].Math > s[j].Math }
func (s ByMath) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//按照英语成绩排序
func (s ByEnglish) Len() int           { return len(s) }
func (s ByEnglish) Less(i, j int) bool { return s[i].English > s[j].English }
func (s ByEnglish) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//按照物理成绩排序
func (s ByPhysical) Len() int           { return len(s) }
func (s ByPhysical) Less(i, j int) bool { return s[i].Physical > s[j].Physical }
func (s ByPhysical) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//按照化学成绩排序
func (s ByChemistry) Len() int           { return len(s) }
func (s ByChemistry) Less(i, j int) bool { return s[i].Chemistry > s[j].Chemistry }
func (s ByChemistry) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//按照生物成绩排序
func (s ByBiology) Len() int           { return len(s) }
func (s ByBiology) Less(i, j int) bool { return s[i].Biology > s[j].Biology }
func (s ByBiology) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//按照体育成绩排序
func (s BySports) Len() int           { return len(s) }
func (s BySports) Less(i, j int) bool { return s[i].Sports > s[j].Sports }
func (s BySports) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//按照总分成绩排序
func (s BySum) Len() int { return len(s) }
func (s BySum) Less(i, j int) bool {
	return s[i].Chinese+s[i].Math+s[i].English+s[i].Physical+s[i].Chemistry+s[i].Biology+s[i].Sports > s[j].Chinese+s[j].Math+s[j].English+s[j].Physical+s[j].Chemistry+s[j].Biology+s[j].Sports
}
func (s BySum) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//管理员

var MANAGER_BUF map[string]Manager
