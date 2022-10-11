#!/bin/sh


###
 # @Author: Mamba24 akateason@qq.com
 # @Date: 2022-10-12 01:25:06
 # @LastEditors: Mamba24 akateason@qq.com
 # @LastEditTime: 2022-10-12 01:52:27
 # @FilePath: /go/installAll.sh
 # @Description: 
 # 
 # Copyright (c) 2022 by Mamba24 akateason@qq.com, All Rights Reserved. 
### 

echo 'pull ...'
git pull

cp -r sender/. /usr/local/bin/
echo '\n\n\n安装成功.🚀\n\n\n'
