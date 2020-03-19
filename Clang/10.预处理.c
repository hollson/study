// 井号表示预处理命令

// 包含头文件
#include "stdio.h"
#include <stdio.h>

// 宏定义
#define MAX(a, b) (a > b) ? a : b

//条件编译
#if _WIN32
#include <windows.h>
#elif __linux__
#include <unistd.h>
#endif

// 预处理1：宏定义
void Define()
{
    int x, y, max;
    printf("input two numbers: ");
    scanf("%d %d", &x, &y);
    max = MAX(x, y);
    printf("max=%d\n", max);
}

// 预处理2：包含文件
void Include()
{
    // 尖括号：在包含文件目录中去查找(由系统的环境变量进行决定)。
    // 双引号：先在当前源文件目录查找, 未找见再去包含目录中查找。
    printf("#include表示包含文件的预处理命令\n");
}

// 预处理3：条件编译
void CondCompile()
{
#if _WIN32
    printf("当前为windows平台\n");
    Sleep(5000);
#elif __linux__
    printf("当前为linux平台\n");
    sleep(5);
#elif __Darwin__
    printf("当前为macos平台\n");
    sleep(5);
#endif
    puts("http://c.biancheng.net/");

    // 1. #error
    // 2. #line
    // 3. #pragma
}

// 预处理命令不是C语言本身的组成部分,不能直接对它们进行编译,
// C语言本只能将预处理之后的源程序进行编译处理,得到目标代码。
int main(void)
{
    // Define();
    // Include();
    CondCompile();
}