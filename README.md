# stus

Stus Teaching Utility System  |  Stus 学生管理系统

![stus logo](https://tva1.sinaimg.cn/large/007S8ZIlgy1gfrlz7cnnlj30dw0dwq34.jpg)

> This stuff is just a term project for my database course. DO NOT USE IT in any realistic scenario.

*Stus* (recursive acronym for Stus Teaching Utility System) is a student management system. This system is based on MySQL database and implemented via Web full-stack tech. By deploying a server-side application on your server, this system can be accessed in the browsers of any terminal.

The system is designed for three types of users, namely, academic administration, teachers and students, and it implements the following main functions:

- Centralized management of student information, course information, and faculty information by Academic Affairs.
- Statistics and management of students and grades in the courses taught by teachers.
- Student access to grades for courses taken, and the course selection system.

 *Stus* 是一个基于 MySQL 数据库，使用 Web 全栈技术实现的**学生管理系统**。通过在一台上服务器部署服务端程序，即可在任意终端的浏览器中使用该系统。

该系统面向教务管理、教师和学生三类使用者，分别实现了如下主要功能：

- 教务对学生信息、课程信息、教师信息的集中管理；
- 教师对所授课程学生、成绩的统计、管理；
- 学生对所修课程的成绩查询、以及选课系统。

The front-end of stus is at https://github.com/cdfmlr/stus-front

## Getting Started

- For **Linux / MacOS**:

```sh
wget -N --no-check-certificate https://raw.githubusercontent.com/cdfmlr/stus/master/install.sh
bash install.sh
```

- For **Windows**:

WSL (recommended) or building from source manually.

- From **Source**:

Require: `git`, `npm`, `go>=12`

```sh
git clone https://github.com/cdfmlr/stus.git	# clone back end
git clone https://github.com/cdfmlr/stus-front.git	# clone front end

# build back end
cd stus/main
go build main.go
# You will get the bin file: `main`
main --help		# see usages
cd ../..

# build front end
cd stus-front
npm install -g @vue/cli		# install vue cli if not exist
npm install ant-design-vue
npm run build	# web ui (static ) will be built to ./dist
cd ..

# Run stus
cd stus/main
./main -static ../../stus-front/dist
```

## Documentation

Full documentation can be found at `docs/stus.md` (zh-cn).

## License

Copyright 2020 CDFMLR                                                   

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

