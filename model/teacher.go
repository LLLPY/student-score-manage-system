package model

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 老师
type Teacher struct {
	num       string //职工号
	name      string //姓名
	major     string //专业
	birthday  string //出生日期
	gender    uint64 //性别 0：男 1：女
	user_type uint64 //用户类型 1：学生 2：教师 3：管理员
	password  string //密码

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
				gender, _ := strconv.ParseUint(v_list[4], 10, 64)
				user_type, _ := strconv.ParseUint(v_list[5], 10, 64)
				password := v_list[6]
				TEACHER_BUF[num] = Teacher{num: num, name: name, major: major, birthday: birthday, gender: gender, user_type: user_type, password: password}
			}
		}
		return nil
	}
	// print("没有读文件")
	return nil

}
