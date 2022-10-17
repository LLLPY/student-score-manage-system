# 基于go语言的学生成绩管理系统



## 简介



一个基于go语言的学生成绩管理系统。
刚刚学习完go语言基础,纯当拿来练手用的。





## 需求分析



面向的用户：教师，学生，管理员

管理范围：某一个班级的所有学生



教师拥有查询修改删除增加学生信息的权限

学生拥有成绩查询的权限

管理者拥有管理管理教师和学生的权限









### 1.登录

- 所用只有登录后才可使用

  

### 2.登出

- 退出系统



### 3.增

- 能够新增学生的信息和成绩

  

### 4.删

- 能够删除学生的信息和成绩



### 5.改

- 能够修改学生的信息和成绩



### 6.查

- 能够根据学号或者姓名进行查找



### 7.统计

- 对学生成绩进行总分排序，单科排序
- 计算学生的平均分，方差，中位数，极差等等



### 8.其他

- 成绩pk
- 历史成绩





## 系统设计



采用MTV的设计模式

- M层：主要负责数据层面的内容，读取写入文件等
- T层：主要用于内容的显示，各级菜单等
- V层：视图层，主要负责业务逻辑



对象：管理员，教师，学生

每一类对象定义一个结构体















存储，暂时使用文件存储。

管理员：manager.txt

教师：teacher.txt

学生：student.txt































