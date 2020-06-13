######################################################################
#
# FILE:   install.sh
# BY:     CDFMLR
# UPDATE: 2020.06.14
#
# An Install Script for stus:
#     - back-end:  https://github.com/cdfmlr/stus.git
#     - front-end: https://github.com/cdfmlr/stus-front.git
#
# Copyright 2020 CDFMLR
# 
# Licensed under the Apache License, Version 2.0 (the "License"); 
# you may not use this file except in compliance with the License. 
# You may obtain a copy of the License at
# 
#        http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, 
# software distributed under the License is distributed on an 
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, 
# either express or implied. See the License for the specific 
# language governing permissions and limitations under the License.
#
#####################################################################

# 依赖检查
echo "👉 stus-install> 1/3 > Checke env..."

echo "git check: "
if command -v git; then
    echo "    Version `git --version` installed."
else
    echo "    git not found, please install it and try again."
    exit 1
fi

echo ""
echo "npm check: "
if command -v npm; then
    echo "    Version `npm -v` installed."
else
    echo "    npm not found, please install it and try again."
    exit 1
fi

echo ""
echo "golang check: "
if command -v go; then
    echo "    Version `go version` installed."
else
    echo "    golang not found, please install it and try again."
    exit 1
fi

# 拉取 git 仓库
echo ""
echo "👉 stus-install> 2/3 > Clone src..."

echo ""
echo "clone stus & stus-front from GitHub..."

mkdir stus
cd stus

mkdir dist
mkdir src
cd src

git clone https://github.com/cdfmlr/stus.git
git clone https://github.com/cdfmlr/stus-front.git

cd ..

# 编译
echo ""
echo "stus-install> 3/3 > Build src..."


# 编译后端
echo ""
echo "Build back-end..."

cd src/stus/main
go build main.go
echo "Done."
mv main ../../../dist/stus
cd ../../..

# 编译前端
echo ""
echo "build front-end..."

cd src/stus-front
if command -v vue; then
    echo "vue cli installed. Skip."
else
    echo "vue cli missing. install it: npm install -g @vue/cli"
    npm install -g @vue/cli
fi
echo "npm install ant-design-vue"
npm install ant-design-vue
echo "npm run build"
npm run build
echo "build done."
mv dist ../../dist/static
cd ../../..

echo ""
echo "👉 stus-install>> Done."
echo "stus 安装在 ./stus/dist"
echo "开始使用: "
echo "    $ cd ./stus/dist  # 必须进到目录再运行"
echo "    $ ./stus"
echo "然后你可以在 http://localhost:9001 访问 stus 服务。"
echo "更多使用方法请访问主页：https://github.com/cdfmlr/stus"
echo "---------------------------"
echo "Created by CDFMLR. All rights reserved."
