#include <stdio.h>

void hello()
{
    printf("hello world\n");
}
// 被调函数必须位于主调函数之前
int main(int argc, char *argv[])
{
    hello();
    return 0;
}
