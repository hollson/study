#!/bin/bash

# 整体定义
array=(aaa bbb ccc )

# 分量定义
array[0]=aaaa
array[1]="bbbb"
array[2]=cccc

# 读取数组
echo ${array[1]} #某个
echo ${array[@]} #所有

# 数组长度
len=${#array[@]}
len=${#array[*]}
echo $len

# 元素长度
len=${#array[2]}
echo $len