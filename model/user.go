package model

//user.txt用于存储所有用户（学生，教师和管理员）以及他们的类型标识

//定义一个用户的接口
type User interface {
	Read_to_buffer(string) error //将文件内容读到缓冲区
	Login(string, string) bool   //登录
	Logout() error               //登出
	Show_info()                  //打印信息
	Create() error               //新增
	Find()                       //查找
	Update()                     //更新
	Delete() error               //删除
}

//学生
var STUDENT_BUF map[string]Student
var SCORE_BUF []Score
var Semester_List = []string{"大一上", "大一下", "大二上", "大二下", "大三上", "大三下", "大四上", "大四下"}
var Semester_Mapping = make(map[uint64]string, len(Semester_List))
var Gender_List = []string{"男", "女"}
var Gender_Mapping = make(map[uint64]string, len(Gender_List))
var User_Type_List = []string{"学生", "教师", "管理员"}
var User_Type_Mapping = make(map[uint64]string, len(User_Type_List))
var Major_List = []string{"信息与计算科学", "应用数学"}
var Major_Mapping = make(map[uint64]string, len(Major_List))

// 初始化
func init() {
	//初始化学期映射
	for i := 0; i < len(Semester_List); i++ {
		Semester_Mapping[uint64(i+1)] = Semester_List[i]
	}
	//初始化性别映射
	for i := 0; i < len(Gender_List); i++ {
		Gender_Mapping[uint64(i)] = Gender_List[i]

	}
	//初始化用户类型映射
	for i := 0; i < len(User_Type_List); i++ {
		User_Type_Mapping[uint64(i+1)] = User_Type_List[i]
	}

	//初始化专业列表
	for i := 0; i < len(Major_List); i++ {
		Major_Mapping[uint64(i+1)] = Major_List[i]
	}

}

//教师
var TEACHER_BUF map[string]Teacher

//管理员

var MANAGER_BUF map[string]Manager
