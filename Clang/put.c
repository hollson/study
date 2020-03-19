#include <stdio.h>

#define MUL1(x, y) x *y
#define MUL2(x, y) (x) * (y)
#define MUL3(x, y) (x * y)

int main()
{
    int m = 2, n = 3;
    printf("%d,%d,%d\n", MUL1(m, m + n) * 2, MUL2(m, m + n) * 2, MUL3(m, m + n) * 2);
    return 0;
}