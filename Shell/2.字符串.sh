#!/bin/bash

# 字符串
str1=hello
str2='hello'
str3="hello" #推荐
echo $str1 $str2 $str3

# 字符串拼接
greet="hello"
echo "$greet world"
echo "${greet} world"

# 字符串长度
echo ${#greet}

# 子字符串
hw="hello world"
echo ${hw:0:5}





