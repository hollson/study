#include <stdio.h>
#include <string.h>


//typedef

struct INFO
{
    int num;
    char str[256];
};

int main(void)
{
    struct INFO A;
    A.num = 2014;
    strcpy(A.str, "Welcome to dotcpp.com");
    printf("This year is %d %s\n", A.num, A.str);
    return 0;
}