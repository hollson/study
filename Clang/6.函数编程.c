#include <stdio.h>

// 声明函数(类似接口)
int Sum(int a, int b);
// 函数实现
int Sum(int a, int b)
{
    return a + b;
}

// 1. 自动变量
void AutoVar(int a)
{
    auto int b, c = 3; /*定义 b,c 为自动变量*/
    printf("%d,%d,%d\n", a, b, c);
}

// 2. 外部变量
void OutVar()
{
    ;
}

// 3. 静态变量
void StaticVar()
{
    ;
}

// 4. 寄存器变量
void RegisterVar()
{
    ;
    // register
}
// 保存变量当前值的存储单元有两类, 一类是内存, 另一类是 CPU 的寄存 器。
// 变量的存储类型关系到变量的存储位置, C 语言中定义了 4 种存储属性, 即自动变量、外部变量、静 态变量和寄存器变量, 它关系到变量在内存中的存放位置, 由此决定了变量的保留时间和变量的作用范围。
int main(int argc, char *argv[])
{
    // AutoVar(2);
    // OutVar();
    // StaticVar();
    // RegisterVar();
    printf("%d\n", Sum(2, 3));
}
