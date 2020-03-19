#include <stdio.h>
#include "include/func.h"

// gcc liby异常：https://www.cnblogs.com/humz/p/10592602.html
int main()
{
    int n1 = 1, n2 = 10;
    printf("从%d加到%d的和为%ld\n", n1, n2, sum(n1, n2));
    printf("从%d乘到%d的积为%ld\n", n1, n2, mult(n1, n2));

    // printf("OS：%s\n", OS);
    printf("Power By %s（%s）", getWebName(), getWebURL());
    return 0;
}