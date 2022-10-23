package utils

import (
	"fmt"
	"strconv"
)

//合法的输入一个字符
func Legal_input_string(tips string, legal_set map[string]string) (s string) {
	fmt.Printf(tips) //输出提示
	legal_list := ""

	for v, k := range legal_set {
		legal_list += v + k + ", "
	}
	if len(legal_list) > 0 {
		legal_list = legal_list[:len(legal_list)-2] //去掉最后一个逗号

	}
	for {
		fmt.Scanln(&s)
		_, ok := legal_set[s]
		if ok || len(legal_set) == 0 { //如果长度为0，说明可以输入任意字符
			break
		} else {
			fmt.Printf("输入有误，请重新输入(合法输入列表：[%v])：", legal_list)

		}
	}

	return

}

//合法的输入一个数字
func Legal_input_int(tips string, legal_set map[int]string) (s int) {
	fmt.Printf(tips) //输出提示

	//找出索引范围
	max_index := 1
	for k, _ := range legal_set {
		if k > max_index {
			max_index = k
		}
	}

	legal_list := ""
	for i := 1; i <= max_index; i++ {
		legal_list += strconv.Itoa(i) + "." + legal_set[i] + ", "
		if i%8 == 0 {
			legal_list += "\n"
		}

	}
	if len(legal_list) > 0 {
		legal_list = legal_list[:len(legal_list)-2] //去掉最后一个逗号

	}
	for {
		fmt.Scanln(&s)
		_, ok := legal_set[s]
		if ok || len(legal_set) == 0 { //如果长度为0，说明可以输入任意字符
			break
		} else {
			fmt.Printf("输入有误，请重新输入(合法输入列表：[%v])：", legal_list)

		}
	}

	return

}
