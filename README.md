# 基于go语言的史上最流弊的学生成绩管理系统





## 简介



<img src="imgs/image-20221022232831028.png" alt="image-20221022232831028" style="zoom:200%;" />





一个基于go语言的学生成绩管理系统，没有花里胡哨的界面，但有你想象不到的功能；没有mysql，redis做支撑，但文件管理也不赖
；没有高大上的高级语法，但经验告诉我：大道至简！



刚刚学习完go语言基础,纯当拿来练手用的。喜欢的小伙伴可以拿去玩玩~🤪🤪🤪



## 需求分析



### think before you do

写代码之前要搞明白自己要做什么，要开发出一个啥玩意，而不是一边写一遍想，那样也许能完成任务，但是你有时候会发现，当你在完成一个功能的时候，也许这个功能和之前的某个功能重复，但你为了代码的简洁性将这个两个功能的公共部分进行封装，也许很快就会封装完成，但是当开发到第n个功能的时候再次遇到了这样的情况，你是不是得再次封装？再者，如果在开发到第n个功能的时候，发现这个功能好像和第m个功能有点类似，此时你觉得有必要删除第m个功能，但是当你着手去删除的时候，发现它后面牵扯到一堆其他的功能，这个时候你会为了让你的程序显得更加完美而去做这件复杂的“删除”？种种诸如此类的问题都是因为一开始没有搞明白自己要干什么，对于要开发的程序没有一个大体的框架和认知！所以说，think before you do---很重要！在开始写代码之前，先好好思考自己要写一个什么样的程序，然后设计出程序的整体框架，框架的每个部分要实现什么功能？这些功能又可以由哪些子功能组成？实现这些子功能使用何种技术栈更加合理？



### 确定需求

在进行框架的设计之前，首先搞懂需求，以需求为驱动！

我打算实现一个学生成绩管理系统，见名知意，重点是对**学生成绩**的**管理**，当然，除了学生以外，还有教师，他们是充当学生的管理者，与此同时，对于教师的管理，权限就交给管理员；每个身份的用户都有自己特定的功能，不能越权操作。



简化需求：

- ​	面向用户
  - 学生
  - 教师
  - 管理员

- 功能划分	

  - 学生能够查询，比较成绩等

  - 教师能够管理学生的信息和成绩等

  - 管理可以管理教师的信息等

    

### 详细功能划分



#### **学生**

- 查看个人信息
- 查询成绩

  - 按照学号查询
  - 按照姓名查询
- 成绩PK

  - 能够和同班级的同学进行各科成绩以及总分的比较
- 成绩分析

  - 分析各科的成绩走势，绘制柱状图，计算方差，均值等来得出成绩是否稳定的结论

    

#### **教师**

- 查看个人信息

- 查看所有学生成绩（单科排序，总分排序，自己所管理的班级）

  - 可以查看所有学生的成绩
  - 能够选择成绩的排序方式，单科或则总分排序后显示

- 查询学生成绩

  - 功能同学生的查询成绩

- 成绩分析

  - 功能同学生的成绩分析，但是作用对象是自己所管理的某个班级的整体成绩分析系

- 管理班级

  - 学生信息列表

  - 查询学生信息

  - 更新学生信息

  - 新增学生信息

  - 删除学生信息

    

#### **管理员**

- 查看个人信息
- 教师信息列表
- 查询教师信息
- 更新教师信息
- 新增教师信息
- 删除教师信息



#### 所有用户

- 登录
  - 根据用户身份，跳转到各自对应的功能页面

- 登出
  - 退出系统，返回到登录界面，不退出程序



#### 其他

- 学生成绩的录入
  - 从指定文件录入学生成绩到系统中

- 定时写入，将buffer中的数据定时写入到文件中，以达到持久化的目的



## 系统设计



### **采用MTV的设计模式**

- M层：主要负责数据层面的交互，读取文件到buffer，写入文件等
- T层：主要用于内容的显示，各级菜单等
- V层：视图层，主要负责业务逻辑等



### 数据结构



#### 学生

```go
// 学生
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
}
```



#### 教师

```go

// 老师
type Teacher struct {
	num        string //职工号
	name       string //姓名
	major      string //专业
	birthday   string //出生日期
	gender     uint64 //性别 0：男 1：女
	user_type  uint64 //用户类型 1：学生 2：教师 3：管理员
	password   string //密码
	class_list string //管理的班级
}
```



#### 管理员

```go
// 管理者
type Manager struct {
	num       string //职工号
	name      string //姓名
	birthday  string //出生日期
	gender    uint64 //性别 0：男 1：女
	user_type uint64 //用户类型 1：学生 2：教师 3：管理员
	password  string //密码
}
```



#### 成绩

```go
// 成绩
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
```



## 存储



使用文件进行数据的存储。

管理员：manager.txt

教师：teacher.txt

学生：student.txt

成绩：score.txt



























