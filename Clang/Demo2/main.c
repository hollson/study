
#include <stdio.h>
#include "cal.h"

int main()
{
    int n = 10;
    int m = 5;
    int result;
    result = add(n, m);
    printf("n+m=%i\n", result);
    result = sub(n, m);
    printf("n-m=%i\n", result);
    return 0;
}