#include <stdio.h>

// 1. If语句
void IfCase(int flag)
{
    if (flag >= 90)
    {
        printf("Good\n");
    }
    else if (flag >= 60)
    {
        printf("passed\n");
    }
    else
    {
        printf("failed\n");
    }
}

// 2. Switch语句
void SwitchCase(int value)
{
    switch (value)
    {
    case 1:
        printf("one\n");
        break;
    case 2:
        printf("two\n");
        break;
    case 3:
        printf("three\n");
        break;
    default:
        printf("other\n");
        break;
    }
}

// 3. For循环
void ForLoop()
{
    for (int i = 0; i < 5; i++)
    {
        printf("%d\n", i);
    }
}

// 4. While循环
void WhileLoop()
{
    int i = 0;
    while (i++ < 5)
    {
        printf("%d\n", i);
    }
}

// 5. DoWhile循环
void DoWhileLoop()
{
    int i = 0;
    do
    {
        printf("%d\n", i);
        i++;
    } while (i < 5);
}

int main()
{
    // IfCase(10);
    // SwitchCase(2);
    // ForLoop();
    WhileLoop();
    DoWhileLoop();
}