package model

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//管理者
type Manager struct {
	num       string //职工号
	name      string //姓名
	birthday  string //出生日期
	gender    uint64 //性别 0：男 1：女
	user_type uint64 //用户类型 1：学生 2：教师 3：管理员
	password  string //密码
}

//将文件中的数据读到缓冲区，用的时候直接从缓冲区取，而不需要每次都从文件中读取
func (manager Manager) Read_to_buffer(filename string) (err error) {

	if len(MANAGER_BUF) == 0 {
		b, err := os.ReadFile(filename)
		// print("读了文件")
		if err != nil {
			fmt.Printf("管理员信息读取失败: %v\n", err)
			return err
		} else {
			MANAGER_BUF = make(map[string]Manager)
			data_list := strings.Split(string(b), "\n")
			for _, v := range data_list {
				v_list := strings.Split(v, ",")
				num := v_list[0]
				name := v_list[1]
				birthday := v_list[2]
				gender, _ := strconv.ParseUint(v_list[3], 10, 64)
				user_type, _ := strconv.ParseUint(v_list[4], 10, 64)
				password := v_list[5]
				MANAGER_BUF[num] = Manager{num: num, name: name, birthday: birthday, gender: gender, user_type: user_type, password: password}
				fmt.Printf("MANAGER_BUF[num]: %v\n", MANAGER_BUF[num])
			}
		}

		return nil
	}
	// print("没有读文件")
	return nil

}
