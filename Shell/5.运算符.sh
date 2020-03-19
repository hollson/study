#!/bin/bash

:<<EOF
表达式和运算符之间要有空格
运算符包括：
    算数运算符
    关系运算符
    布尔运算符
    字符串运算符
    文件测试运算符
EOF

#=================数字运算符================
a=10 b=20

val=`expr $a + $b`
echo "a + b : $val"

val=`expr $a \* $b`
echo "a * b : $val"

if [ $a != $b ]
then
   echo "a 不等于 b"
fi

#=================关系运算符================
:<<EOF
关系运算符只支持数字，不支持字符串，除非字符串的值是数字：
-eq 	相等         [ $a -eq $b ] 返回 false。
-ne 	不相等 	    [ $a -ne $b ] 返回 true。
-gt 	大于右边的 	  [ $a -gt $b ] 返回 false。
-lt 	小于右边的 	  [ $a -lt $b ] 返回 true。
-ge 	大于等于右边的 [ $a -ge $b ] 返回 false。
-le 	小于等于右边的 [ $a -le $b ] 返回 true。
EOF

if [ $a -eq $b ]
then
   echo "$a -eq $b : a 等于 b"
else
   echo "$a -eq $b: a 不等于 b"
fi

#=================布尔运算符================
:<<EOF
! 	非运算	[ ! false ] 返回 true。
-o 	或运算	[ $a -lt 20 -o $b -gt 100 ] 返回 true。
-a 	与运算	[ $a -lt 20 -a $b -gt 100 ] 返回 false。
EOF

if [ $a -lt 100 -a $b -gt 15 ]
then
   echo "$a 小于 100 且 $b 大于 15 : 返回 true"
else
   echo "$a 小于 100 且 $b 大于 15 : 返回 false"
fi

#=================逻辑运算符================
:<<EOF
&& 	逻辑的AND 	[[ $a -lt 100 && $b -gt 100 ]] 返回 false
|| 	逻辑的OR 	[[ $a -lt 100 || $b -gt 100 ]] 返回 true
EOF

if [[ $a -lt 100 && $b -gt 100 ]]
then
   echo "返回 true"
else
   echo "返回 false"
fi

#=================字符串运算符================
:<<EOF
= 	检测两个字符串是否相等，相等返回 true。 	[ $a = $b ] 返回 false。
!= 	检测两个字符串是否相等，不相等返回 true。 	[ $a != $b ] 返回 true。
-z 	检测字符串长度是否为0，为0返回 true。 	[ -z $a ] 返回 false。
-n 	检测字符串长度是否为0，不为0返回 true。 	[ -n "$a" ] 返回 true。
$ 	检测字符串是否为空，不为空返回 true。 	[ $a ] 返回 true。
EOF

a="abc"
b="efg"

if [ $a = $b ]
then
   echo "$a = $b : a 等于 b"
else
   echo "$a = $b: a 不等于 b"
fi
if [ $a != $b ]
then
   echo "$a != $b : a 不等于 b"
else
   echo "$a != $b: a 等于 b"
fi
if [ -z $a ]
then
   echo "-z $a : 字符串长度为 0"
else
   echo "-z $a : 字符串长度不为 0"
fi
if [ -n "$a" ]
then
   echo "-n $a : 字符串长度不为 0"
else
   echo "-n $a : 字符串长度为 0"
fi
if [ $a ]
then
   echo "$a : 字符串不为空"
else
   echo "$a : 字符串为空"
fi

#=================文件测试运算符================
:<<EOF
-b file 	检测文件是否是块设备文件，如果是，则返回 true。 	[ -b $file ] 返回 false。
-c file 	检测文件是否是字符设备文件，如果是，则返回 true。 	[ -c $file ] 返回 false。
-d file 	检测文件是否是目录，如果是，则返回 true。 	[ -d $file ] 返回 false。
-f file 	检测文件是否是普通文件（既不是目录，也不是设备文件），如果是，则返回 true。 	[ -f $file ] 返回 true。
-g file 	检测文件是否设置了 SGID 位，如果是，则返回 true。 	[ -g $file ] 返回 false。
-k file 	检测文件是否设置了粘着位(Sticky Bit)，如果是，则返回 true。 	[ -k $file ] 返回 false。
-p file 	检测文件是否是有名管道，如果是，则返回 true。 	[ -p $file ] 返回 false。
-u file 	检测文件是否设置了 SUID 位，如果是，则返回 true。 	[ -u $file ] 返回 false。
-r file 	检测文件是否可读，如果是，则返回 true。 	[ -r $file ] 返回 true。
-w file 	检测文件是否可写，如果是，则返回 true。 	[ -w $file ] 返回 true。
-x file 	检测文件是否可执行，如果是，则返回 true。 	[ -x $file ] 返回 true。
-s file 	检测文件是否为空（文件大小是否大于0），不为空返回 true。 	[ -s $file ] 返回 true。
-e file 	检测文件（包括目录）是否存在，如果是，则返回 true。 	[ -e $file ] 返回 true。
EOF

file="/var/www/runoob/test.sh"
if [ -r $file ]
then
   echo "文件可读"
else
   echo "文件不可读"
fi
if [ -w $file ]
then
   echo "文件可写"
else
   echo "文件不可写"
fi
if [ -x $file ]
then
   echo "文件可执行"
else
   echo "文件不可执行"
fi
if [ -f $file ]
then
   echo "文件为普通文件"
else
   echo "文件为特殊文件"
fi
if [ -d $file ]
then
   echo "文件是个目录"
else
   echo "文件不是个目录"
fi
if [ -s $file ]
then
   echo "文件不为空"
else
   echo "文件为空"
fi
if [ -e $file ]
then
   echo "文件存在"
else
   echo "文件不存在"
fi