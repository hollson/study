#!/bin/bash

echo "第1个参: $1"
echo "第2个参: ${2}"
echo "所有参数: $@"
echo "所有参数: $*"
echo "参数总数: $#"
echo "脚本进程ID: $$"
echo "后台进程ID: $!"
echo "当前选项: $-"
echo "最后状态: $?"

# 执行 ./xxx.sh a b c d e

:<<EOF
###$*与$@区别###
*：理解为整体所有(传递了一个参数）
@：理解为元素所有
EOF

for i in "$*"; do
    echo $i
done

for i in "$@"; do
    echo $i
done

