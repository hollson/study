# !/bin/bash
# 定义变量
myvar="hello world"
echo ${myvar}

# 只读变量
readonly myvar
myvar="hello shell"
echo $myvar

# 删除变量
unset myvar
echo $myvar
