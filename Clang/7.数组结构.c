#include <stdio.h>

// 创建数组
void Array()
{
    int a[3];                   //定义整型数组(赋默认值)
    float b[3];                 //定义浮点数组(赋默认值)
    char c[8] = {97, 'b', 'c'}; //定义字符数组，并初始化值

    for (int i = 0; i < 8; i++)
    {
        printf("%d \n", c[i]);
    }
}

// 数组长度
void ArrayLenght()
{
    int a[10] = {0};
    int cnt = sizeof(a) / sizeof(a[0]);
    printf("cnt = %d\n", cnt);
}

// 二维数组
void D2Array()
{
    int a[3][2] = {{1, 2}, {3, 4}, {5, 6}};
    for (size_t i = 0; i < 3; i++)
    {
        printf("{%d,%d}\n", a[i][0], a[i][1]);
    }

    int b[3][2] = {
        1, 6,
        2, 5,
        3, 8};
    printf("%d\n", b[1][1]);
}

// 字符数组
void CharArray()
{
    // 字符串数组末尾有\0结束符,长度+1
    char a[6] = {'c', 'h', 'i', 'n', 'a', '\0'};
    char b[] = {"china"};
    printf("%C,%C\n", a[0], b[5]);
}

int main(int argc, char *argv[])
{
    // Array();
    // ArrLenght();
    D2Array();
    // CharArray();
}
