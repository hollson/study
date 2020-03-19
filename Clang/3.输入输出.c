#include <stdio.h>

int main(void)
{
}

//字符输出函数putchar
void PutChar()
{
    char x = 'X';
    putchar('A');  /*输出大写字母A */
    putchar(x);    /*输出字符变量x的值*/
    putchar('\n'); /*换行*/
}

//字符输入函数getchar
void GetChar()
{
    char c;        /*定义字符变量c*/
    c = getchar(); /*将读取的字符赋值给字符变量c*/
}

//格式化输出函数printf
void PrintF()
{
    int a = 12;
    float b = 3.1415;
    char c = 'A';
    printf("%d\n", a);
    printf("o%o\n", a);
    printf("0x%x\n", a);
    printf("%3.2f\n", b);
    printf("%c\n", c);
    getchar();

    /*使用可变宽度输出字段*/
    unsigned width, precision;
    int number = 256;
    double weight = 25.5;
    printf("Please input number's width:\n");
    scanf("%d", &width);
    printf("The number is: %*d\n", width, number);
    printf("Then please input width and precision:\n");
    scanf("%d %d", &width, &precision);
    printf("Weight = %*.*f\n", width, precision, weight);
}

// 格式化输入函数scanf
void ScanF()
{
    float a;
    scanf("%10f", &a); //正确
    // scanf("%10.2f", &a); //错误

    /*用*跳过scanf接收的数字*/
    int num;
    printf("Please enter three number:\n");
    scanf("%*d %*d %d", &num);
    printf("The last number is %d\n", num);
}