# 数据库系统开发实践 — stus学生管理系统

## 1. 项目概述

本次作业设计、开发一个**学生管理**数据库系统，旨在练习数据库系统的开发及使用。

### 1.1 实现目标

在本次作业中，我完成了学生管理系统 *stus* ( Stus Teaching Utility System 的递归缩写)的设计、开发。该系统基于 MySQL 数据库，使用 Web 全栈技术实现。通过在一台上服务器部署服务端程序，即可在任意终端的浏览器中使用该系统。

该系统面向教务管理、教师和学生三类使用者，分别实现了如下主要功能：

- 教务对学生信息、课程信息、教师信息的集中管理；
- 教师对所授课程学生、成绩的统计、管理；
- 学生对所修课程的成绩查询、以及选课系统。

### 1.2 开发环境

本次开发在 macOS Mojave 10.14.6 下完成。涉及的主要软件依赖如下：

| 项目     | 依赖   | 版本                                     |
| -------- | ------ | ---------------------------------------- |
| 数据库   | MySQL  | 8.0.15 for osx10.13 on x86_64 (Homebrew) |
| Web 后端 | Go     | 1.12 darwin/amd64                        |
| Web 前端 | Vue.js | 2.6.11 (@vue/cli 4.3.1)                  |

开发完成的服务端程序可以在 macOS，Linux，UNIX 系统运行（理论上兼容 Microsoft Windows，但没有经过测试，不保证实际可行）。前端网站可以在任意浏览器运行（不兼容IE，推荐使用最新的Chrome/Foxfire/Safari）。

## 2. 需求分析

### 2.1 业务活动情况

学生管理系统需要分别服务于教务管理人员、教师和学生三类使用者。

- 教务管理人员建立并管理学生学籍信息、教师信息以及开设课程信息，为课程分配教师，汇总、管理学生选课信息；教务管理人员拥有对所有数据的最高管理权限；
- 学生选课时，可以从学校开设的课程中选择自己需要或希望上的课程。学生还需要可以查询自己选的各门课的成绩；
- 教师可以查看教务为自己分配的课程（一位教师可以对应多门课程）、查看选课的学生，在完成一门课的学习后，教师需要录入学生的成绩，并且得到平均分、不及格人数等汇总信息；

为在最大程度上方便使用者，运行学生管理系统的环境要求要尽可能低，所以考虑使用 Web 应用开发技术，用户只需使用任意设备上的网页浏览器即可运行系统。系统应该响应式兼容各种尺寸的屏幕并保持设计的简洁性、现代化。

### 2.2 系统初步需求

设计一套实用的先进孤立系统以简化教务管理人员和教师的工作、方便学生查询并安排自己的学习情况。

- 系统性质：MIS 软件
- 使用者：教务管理人员，教师，学生
- 系统运行环境：网络运行

### 2.3 系统主要功能

- 学生管理：登记学生的基本信息（姓名，性别，班级等），并提供查询功能；
- 课程管理：登记课程的基本信息（课程名称，课程类型，学分等），并提供查询功能；
- 成绩管理：登记学生各门课程的考试成绩，并提供查询、统计功能；
- 教师管理：登记教师的基本信息（姓名、性别、年龄、职称等），并提供查询功能；
- 选课管理：登记学生选修课程，提供查询功能；
- 授课管理：登记教师教授课程，提供查询功能。

## 3. 数据库设计

### 3.1 数据字典

#### 3.1.1 学生

数据项：

| 数据项 | 含义说明             |
| ------ | -------------------- |
| sid    | 学号，学生的唯一标识 |
| sname  | 学生姓名             |
| sdept  | 学生所属院系         |
| sage   | 学生年龄             |
| sgrade | 学生年级             |
| ssex   | 学生性别             |
| sclass | 学生所属班级         |
| smajor | 学生主修专业         |

数据结构：

| 数据结构 | 含义说明                                               | 组成                                                  |
| -------- | ------------------------------------------------------ | ----------------------------------------------------- |
| Student  | 定义一个学生的有关信息，是学籍管理子系统的主体数据结构 | sid，sname，sdept，sage，sgrade，ssex，sclass，smajor |



数据存储：

| 数据存储 | 说明               | 流入数据源   | 流出数据源                                                   | 组成    |
| -------- | ------------------ | ------------ | ------------------------------------------------------------ | ------- |
| 学生表   | 记录学生的基本情况 | 教务录入信息 | 教务查寻学生情况，学生查看自己的信息，教师查寻所授课程的选课学生 | Student |

#### 3.1.2 课程

数据项：

| 数据项 | 含义说明               |
| ------ | ---------------------- |
| cid    | 课程号，课程的唯一标识 |
| cname  | 课程名称               |
| ctype  | 课程类型（必修，选修） |
| cpoint | 课程学分               |
| cweek  | 开课周次               |
| ctime  | 上课时间               |
| caddr  | 上课教室地点           |

数据结构：

| 数据结构 | 含义说明                                               | 组成                                           |
| -------- | ------------------------------------------------------ | ---------------------------------------------- |
| Course   | 定义一门课程的有关信息，是课程管理子系统的主体数据结构 | cid, cname, ctype, cpoint, cweek, ctime, caddr |

数据存储：

| 数据存储 | 说明               | 流入数据源   | 流出数据源                                               | 组成   |
| -------- | ------------------ | ------------ | -------------------------------------------------------- | ------ |
| 课程表   | 记录课程的有关信息 | 教务录入信息 | 教务查寻课程信息，学生查询开设课程，教师查询自己所授课程 | Course |

#### 3.1.3 教师

数据项：

| 数据项 | 含义说明               |
| ------ | ---------------------- |
| tid    | 教工号，教师的唯一标识 |
| tname  | 教师姓名               |
| tdept  | 教师所属院系           |
| tsex   | 教师性别               |
| tpro   | 教师职称               |

数据结构：

| 数据结构 | 含义说明                                               | 组成                          |
| -------- | ------------------------------------------------------ | ----------------------------- |
| Teacher  | 定义一位教师的有关信息，是教师管理子系统的主体数据结构 | tid, tname, tdept, tsex, tpro |

数据存储：

| 数据存储 | 说明               | 流入数据源   | 流出数据源                                 | 组成    |
| -------- | ------------------ | ------------ | ------------------------------------------ | ------- |
| 教师表   | 记录教师的有关信息 | 教务录入信息 | 教务查寻，学生查询，教师查询自己的个人信息 | Teacher |

#### 3.1.4 学生选课、教师授课

数据项：

| 数据项 | 含义说明         |
| ------ | ---------------- |
| result | 学生所修课程成绩 |

数据储存：

| 数据存储   | 说明             | 流入数据源             | 流出数据源                   | 组成             |
| ---------- | ---------------- | ---------------------- | ---------------------------- | ---------------- |
| 学生选课表 | 记录学生选课信息 | 学生选课，教师录入成绩 | 教务查寻，学生查询，教师查询 | sid, cid, result |
| 教师授课表 | 记录教师授课信息 | 教务安排               | 教务查寻，学生查询，教师查询 | cid, tid         |

#### 3.1.5 系统用户信息

| 数据项   | 含义说明                                                     |
| -------- | ------------------------------------------------------------ |
| utype    | 用户类型（学生/教师/教务）                                   |
| uid      | 用户id，登录时用的账号；<br />对应学生、教师、教务分别为学号、教工号、教务管理员账号 |
| password | 系统登录密码                                                 |

数据结构：

| 数据结构 | 含义说明               | 组成                 |
| -------- | ---------------------- | -------------------- |
| Passwd   | 定义一条用户的账户信息 | utype, uid, password |

数据存储：

| 数据存储 | 说明               | 流入数据源     | 流出数据源                             | 组成   |
| -------- | ------------------ | -------------- | -------------------------------------- | ------ |
| 账户表   | 记录账户的有关信息 | 系统管理员创建 | 教务管理人员、教师、学生登录时验证信息 | Passwd |


#### 3.2 实体关系模型

根据需求分析及数据字典，系统需要设计*学生、教师、课程、选课、授课、账户*六张数据库表，作出 E-R 图如下：

![stus数据库的E-R图](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfpl0w8j8bj31420plju6.jpg)

在其中：

1. 性别可以用1位（boolean）来表示以节省空间；
2. 年龄为整数类型，满足大于0；
3. 学分为浮点型，可以表示小数，满足大于0，小于20；
4. 学生选课表中 sid、cid 必须对应在students、courses表中存在;
5. 教师授课表中 tid、cid 必须对应在teachers、courses表中存在；
6. 其余数据都可用字符型表示。





据此就可以做成更详细的设计图：

![stus系统的数据库EER图，由 MySQL Workbench 逆向工程生成](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfpbn9qeaaj310y0u07bs.jpg)

## 4. 软件设计

以上述设计的数据库为基础，现考虑整体软件系统设计，如下图所示：

![stus系统整体设计示意图](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfprw76k3kj31eb0owgor.jpg)



### 4.1 前端设计

从前端的角度看，系统提供网页服务，用户可以在浏览器中运行 stus 系统的 Web App （网页版），用户打开网页后，首先登录系统，然后系统会按照用户的不同角色（教务管理、学生、教师）显示不同的操作页面，提供该角色需要的不同功能：

| 角色         | 提供功能                                                     |
| ------------ | ------------------------------------------------------------ |
| 教务管理人员 | 学生管理，教师管理，课程管理，教师任课管理，学生选课管理     |
| 教师         | 查看自己的个人信息、任课信息，课程管理（所任课程学生查询，成绩录入、统计查询） |
| 学生         | 查看自己的个人信息，成绩查询，学生选课                       |

该 Web App 中所有对数据的增删改查操作都通过调用系统后端提供的 Restful Web API 实现。同时为提高系统的可用性、可拓展性，这一套 API 也可暴露给专业用户，支持用户自己编程使用该系统。

### 4.2 后端设计

从后端的角度看，管理员通过部署 stus 服务端程序来开启系统服务。服务端程序通过一套 ORM (对象关系映射)连接、读写数据库，并分别将学生、教师、教务管理所需的操作封装成 Web API，由一个“Server”抽象结构统一路由调用。

在处理前端发送的请求时，后端程序首先要通过 Access Token 系统认证用户（Access Token 认证系统和用户登录 API 配合使用）。Access Token 认证系统会对未登录或拒不提供 Access Token 信息的用户拒绝访问，对经过登录的合法用户按照不同角色（学生、教师、教务管理）提供的不同操作权限处理用户的具体请求内容。

### 4.3 API 设计

#### 4.3.1 Login API

>  Login API 提供用户的登录操作。

任何人都可以使用该接口。

```markdown
请求：
	POST /api/login
说明：
	用户登录
参数：
	- utype: 用户类型，"student" 或 "teacher" 或 "admin"
	- uid: 用户ID，对应于 student|teacher|admin 分别是 sid|tid|aid
	- password: 用户的登录密码
响应：
	成功时，返回一个 assess token：
		{"token": access_token}
	失败时，返回错误信息：
		{"error": error_description}
```

⚠️【注意】：登录后会返回由 Access Token 系统为用户分配的身份识别码（Access Token），在调用其他所有 API 时都必须在 Request Header 中包含一条 `{"token": access_token}`。



#### 4.3.2 Core API

> Core API 提供对数据的基础增删改查操作。

教务管理人员使用 Core API 完成所有管理操作，学生、教师只能有限制地使用部分 Core API。

```markdown
请求：
	GET|POST|PUT|DELETE /api/core/student
说明：
	获取、添加、修改、删除指定学生信息

请求：
	GET|POST|PUT|DELETE /api/core/course
说明：
	获取、添加、修改、删除指定课程信息

请求：
	GET|POST|PUT|DELETE /api/core/teacher
说明：
	获取、添加、修改、删除指定教师信息

请求：
	GET|POST|PUT|DELETE /api/core/ct
说明：
	获取、添加、修改、删除指定教师授课信息

请求：
	GET|POST|PUT|DELETE /api/core/sc
说明：
	获取、添加、修改、删除指定学生选课信息

/* 以上请求的参数、响应形式均相同如下：*/

参数(GET, DELETE)：
	query: 指定查询的 SQL WHERE 条件，不可为空
参数(POST, PUT)：
	record: 指定要添加或修改的记录（包含全部字段）的 JSON 表示


响应(POST, PUT, DELETE)：
	成功时，返回一个success：
		{"success": "success"}
	失败时，返回错误信息：
		{"error": error_description}
响应(GET)：
	失败时，返回错误信息：
		{"error": error_description}
	成功时，返回查询到的所有记录信息(同数据库设计中的字段)：
		[{...}, ...]
		例如，以某种条件查询学生成功时返回：
            [
                {
                    "sid": "201810000897",
                    "sname": "露卡",
                    "sdept": "数理",
                    "smajor": "信息与计算科学",
                    "sage": 13,
                    "ssex": false,
                    "sgrade": "2018",
                    "sclass": "1808"
                },
                ...
            ]
```

由于 Core API 权限极高，使用时务必慎重。







#### 4.3.3 Student API

> Student API 提供学生所特需的操作。

该接口只应该被学生调用。

成绩查询：

```markdown
请求：
	/api/student/exam_result
说明：
	学生查询自己的各科成绩
参数：
	- sid: 学生学号
响应：
	失败时，返回错误信息：
		{"error": error_description}
	成功时，返回查询到的课程、成绩信息：
        [
            {
                "cid": "w22223",
                "cname": "奋斗学",
                "result": 97
            },
            ...
        ]
```

开设课程查询：

```markdown
请求：
	/api/student/courses
说明：
	学生查询所有开设的课程，在选课时使用
参数：
	无
响应：
	失败时，返回错误信息：
		{"error": error_description}
	成功时，返回查询到的课程、授课老师信息：
        [
            {
                "cid": "w22223",
                "cname": "奋斗学",
                "ctype": "公选",
                "cpoint": 1,
                "cweek": "5",
                "ctime": "205",
                "caddr": "教十八楼213",
                "tid": "t001",
                "tname": "肖灯",
                "tdept": "数理",
                "tsex": true,
                "tpro": "讲师"
            },
            ...
        ]
```

已选课程查询：

```markdown
请求：
	/api/student/courses
说明：
	学生查询自己已选的课程，在选课时使用
参数：
	- sid: 学生学号
响应：
	失败时，返回错误信息：
		{"error": error_description}
	成功时，返回查询到的课程、授课老师信息，与 /api/student/courses 类似：
		[{...}, ...]
```

#### 4.3.4 Teacher API

> Teacher API 提供教师所特需的操作。

该接口只应该被教师调用。

任课查询：

```markdown
请求：
	/api/teacher/courses
说明：
	教师查询自己担当的课程
参数：
	- tid: 教师工号
响应：
	失败时，返回错误信息：
		{"error": error_description}
	成功时，返回查询到的课程即选课人数信息：
        [
            {
                "cid": "w22223",
                "cname": "奋斗学",
                "ctype": "公选",
                "cpoint": 1,
                "cweek": "5",
                "ctime": "205",
                "caddr": "教十八楼213",
                "student_num": 1
            }
        ]
```

课程详情查询：

```markdown
请求：
	/api/teacher/data_of_course
说明：
	教师查询自己担当的某一课程的具体学生信息
参数：
	- cid: 课程号
响应：
	失败时，返回错误信息：
		{"error": error_description}
	成功时，返回查询到的课程选课学生基本信息、成绩、已经课程的统计信息：
        {
            "students": [
                {
                    "sid": "201810000999",
                    "sname": "张三",
                    "sdept": "数理",
                    "smajor": "信息与计算科学",
                    "sage": 20,
                    "ssex": true,
                    "sgrade": "2018",
                    "sclass": "1809",
                    "result": 12	// 学生成绩
                }
            ],
            "count": 1,				// 学生人数
            "average": 12,			// 平均分数
            "best": 12,				// 最高分数
            "worst": 12,			// 最低分数
            "not_pass_count": 1		// 挂科人数
        }
```







## 5. 系统开发

### 5.1 数据库实现

本系统使用 MySQL 作为数据库。

在开始后端开发前，可以按照设计图，手动完成数据库的建立。但我采取的方法是在后端开发中，使用 ORM 完成数据库的建立。

手动定义数据表，可以参考 [附录A]()。

### 5.2 后端实现

本项目后端使用 Go 语言编写。程序实现的结构图（逆向工程得到，有部分手动修改）：

![stus后端UML](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfqin6pdszj31zj0r2tfo.jpg)


在对数据库的连接、读写中使用了开源第三方 ORM 库 —— [GORM](http://gorm.io)，其余部分代码均在 Go 语言标准库下完成。在 GORM 的支持下，除目前系统使用的 MySQL 数据库，还可以选用 PostgreSQL、SQLite3 和 SQL Server，具体的使用方式详见 [GORM 文档](https://gorm.io/docs/connecting_to_the_database.html)。 

后端的源代码可以从 [https://github.com/cdfmlr/stus](https://github.com/cdfmlr/stus) 获取（在 Apache 2.0 协议下开放源代码）。

### 5.3 前端实现

本项目前端以 Vue.js 为基础，用 [Ant Design Vue](https://antdv.com/docs/vue/introduce-cn/) UI 框架实现。程序实现的结构图（逆向工程得到）：

![stus 前端结构图](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfqqlpixe9j324m0u0h1v.jpg)

前端的源代码可以从[https://github.com/cdfmlr/stus-front](https://github.com/cdfmlr/stus-front) 获取（在 Apache 2.0 协议下开放源代码）。

## 6. 使用测试

部署后端程序后，在浏览器打开，即可看到登录界面：

![stus 运行截图，登录页面](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfqqtvhltmj31f40u07wi.jpg)

选择不同身份登录即可看到主要的操作页面，例如，学生登录后的个人信息主页：

![stus 运行截图，学生信息界面](../Library/Application Support/typora-user-images/image-20200613163127326.png)

教师登录后对所授的某一课程的管理页面：

![stus 运行截图，教师课程管理页面](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfqqzx11i9j312k0u0h3y.jpg)

教务管理员登录后的学生管理页面：

![image-20200613163834254](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfqr2ffxy2j31dt0u0wzl.jpg)

系统服务端程序的核心组件均通过了单元测试，以及前端系统涉及的各种增删改查操作也都通过了实际使用测试，可以正常运行。

## 7. 自我评价与总结

本次开发实现的 stus 系统，基本做到了*简化教务管理人员和教师的工作、方便学生查询并安排自己的学习情况*的设计目标，达到了预期。



系统实现的 Web UI 界面简洁、美观，几乎没有操作系统的限制，可以在任何浏览器免安装使用，比较符合当今移动为先、云为先的产品设计理念。响应式设计让用户在手机等非常规桌面型屏幕上也能获得良好的使用体验。



在开发过程中，练习了关系型数据库的设计和开发使用。除了实际应用的需要，我在设计 API 的时候还考虑了对 SQL 有针对性的练习：

- 在 Core API 中，练习了 SQL 的基本增删改查；
- 在 Student API 中，练习了跨表的关联查询（学生选课中显示课程和授课老师关联查询得到的综合信息）；
- 在 Teacher API 中，练习了聚合信息的查询（教师的课程成绩信息统计）；



但这次开发中也有诸多不足：

1、 系统功能不够完善，比如欠缺“班主任查询本班学生成绩”等实际中常用的操作；

2、 没有充分利用、练习 SQL 数据库的部分重要功能，比如视图、权限管理和触发器；


3、 系统安全性低，无法用于实际生产环境。只靠 Access Token 系统保证基础的安全远远不够，系统存在有大量安全隐患，比如缺乏对 SQL 注入的检测、规避机制，也没有利用 SQL 提供的视图、角色权限、触发器等途径从数据库层面保障数据安全；


4、 未经过高压、高并发测试和针对性调试，不能很好的保障系统在压力下工作的正确性、安全性。

5、 直接使用了开源的第三方 ORM 库，大大降低了对嵌入使用 SQL 的练习[^注1]。原本计划在项目中自己实现一个简易的 ORM，以加深对实际生产使用中比较常用和重要的连接池、事务机制、自动关联的对象关系映射、钩子操作、自动迁移等等方面的理解。但由于短时间内实在难以完成一个功能较为全面、满足该项目需求的 ORM 系统，我最终还是选择了直接调用自己日常开发中常用的 Gorm。

6、 在设计数据库时，考虑的范式要求比较低，数据库虽然易于构建，但存在一些数据冗余及其他问题，也对后期拓展、维护不太友好。（问题最突出的是 students 表，属性太多，而且对院系、班级等字段直接使用了字符型储存，这个设计相当失败；应该采取建立院系表、班级表，然后在 students 中存系号、班级编号等的策略）







> [^注1] 在 Go 语言中，标准库里提供了一套 SQL 接口(database/sql)，这套接口只是抽象接口，需要配合第三方实现的 Driver 使用（比如有 MySQL、Sqlite、ODBC 等等的 Driver），这套接口使用原生的嵌入式 SQL，我个人认为，这种实现相比于使用 ORM 库，没有太大的优势，尤其是写出来的代码可读性、可维护性都比较差，所以在以前玩过几次之后就没有再使用了，所以这次开发也没有考虑这一套。

# 附录 A. 数据库定义

取自 MySQL 请求得到的原始DDL（有部分手动修改）：

```mysql
CREATE DATABASE `stus` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */

CREATE TABLE `courses` (
  `cid` varchar(255) NOT NULL,
  `cname` varchar(255) DEFAULT NULL,
  `ctype` varchar(255) DEFAULT NULL,
  `cpoint` double DEFAULT NULL,
  `cweek` varchar(255) DEFAULT NULL,
  `ctime` varchar(255) DEFAULT NULL,
  `caddr` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

CREATE TABLE `students` (
  `sid` varchar(255) NOT NULL,
  `sname` varchar(255) DEFAULT NULL,
  `sdept` varchar(255) DEFAULT NULL,
  `sage` int(11) DEFAULT NULL,
  `ssex` tinyint(1) DEFAULT NULL,
  `sgrade` varchar(255) DEFAULT NULL,
  `sclass` varchar(255) DEFAULT NULL,
  `smajor` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

CREATE TABLE `teachers` (
  `tid` varchar(255) NOT NULL,
  `tname` varchar(255) DEFAULT NULL,
  `tdept` varchar(255) DEFAULT NULL,
  `tsex` tinyint(1) DEFAULT NULL,
  `tpro` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci


CREATE TABLE `student_course_relations` (
  `sid` varchar(255) NOT NULL,
  `cid` varchar(255) NOT NULL,
  `result` double DEFAULT NULL,
  PRIMARY KEY (`sid`,`cid`),
  KEY `student_course_relations_cid_courses_cid_foreign` (`cid`),
  CONSTRAINT `student_course_relations_cid_courses_cid_foreign` 
    FOREIGN KEY (`cid`) REFERENCES `courses` (`cid`) 
    ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `student_course_relations_sid_students_sid_foreign` 
    FOREIGN KEY (`sid`) REFERENCES `students` (`sid`) 
    ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci


CREATE TABLE `course_teacher_relations` (
  `cid` varchar(255) NOT NULL,
  `tid` varchar(255) NOT NULL,
  PRIMARY KEY (`cid`,`tid`),
  KEY `course_teacher_relations_tid_teachers_tid_foreign` (`tid`),
  CONSTRAINT `course_teacher_relations_cid_courses_cid_foreign` 
    FOREIGN KEY (`cid`) REFERENCES `courses` (`cid`) 
    ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `course_teacher_relations_tid_teachers_tid_foreign` 
    FOREIGN KEY (`tid`) REFERENCES `teachers` (`tid`) 
    ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

CREATE TABLE `passwds` (
  `utype` varchar(255) NOT NULL,
  `uid` varchar(255) NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`utype`,`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
```

# 附录 B. 源代码及部署方法

项目涉及的前后端均托管于 GitHub，在 Apache 2.0 协议下开放源代码：

- 后端：[https://github.com/cdfmlr/stus](https://github.com/cdfmlr/stus)
- 前端：[https://github.com/cdfmlr/stus-front](https://github.com/cdfmlr/stus-front)

要部署 stus 服务，macOS / Linux / UNIX 用户可以使用后端仓库中提供的 `install.sh` 脚本，对于 Windows 用户，需要自行参考 `install.sh`，手动完成编译、安装等操作，推荐使用 WSL 运行 `install.sh` 完成操作。



# 参考

[1] CDFMLR. SQL 基本使用[EB/OL].https://blog.csdn.net/u012419550/article/details/104439413 ,2020-02-22.

[2] Jinzhu, GORM Contributors. GORM[EB/OL].https://gorm.io/ ,2020.

[3] CDFMLR. Golang 实战——微信公众号课程提醒系统[EB/OL].https://blog.csdn.net/u012419550/article/details/104781073 ,2020-03-10.

[4] CDFMLR. CiFa-front[EB/OL].https://github.com/cdfmlr/CiFa ,2020-05.

[5] Ant Design Team. Ant Design of Vue[EB/OL].https://antdv.com/docs/vue/introduce-cn/ ,2020.