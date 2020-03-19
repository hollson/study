#include "../include/func.h"

// 从 fromNum 加到 endNum
long sum(int fromNum, int endNum)
{
    int i;
    long result = 0;

    // 参数不符合规则，返回 -1
    if (fromNum < 0 || endNum < 0 || endNum)
    {
        return -1;
    }

    for (i = fromNum; i <= endNum; i++)
    {
        result += i;
    }

    // 返回大于等于0的值
    return result;
}

// 从 fromNum 乘到 endNum
long mult(int fromNum, int endNum)
{
    int i;
    long result = 1;

    // 参数不符合规则，返回 -1
    if (fromNum < 0 || endNum < 0 || endNum)
    {
        return -1;
    }

    for (i = fromNum; i <= endNum; i++)
    {
        result *= i;
    }

    // 返回大于等于0的值
    return result;
}