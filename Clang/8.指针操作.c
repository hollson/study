#include <stdio.h>

// 指针操作
void Pointer()
{
    char a = 'F';
    int f = 123;

    // 定义指针变量
    char *pa = &a;
    int *pb = &f;

    // 取指针值
    // 星号作用：1. 定义指针；2.指针取值
    printf("a=%c\n", *pa);
    printf("f=%d\n", *pb);

    *pa = 'C';
    *pb += 1;
    printf("mod a=%c\n", *pa);
    printf("mod f=%d\n", *pb);

    //指针长度是8字节(跟操作系统相关)
    printf("len(pa)=%d\n", (unsigned int)sizeof(pa));
    printf("len(pf)=%d\n", (unsigned int)sizeof(pb));

    printf("&a=%p\n", pa);
    printf("&f=%p\n", pb);
}

// 指针初始化(野指针)
void PointerInit()
{
    int *a;

    // 未初始化，很危险
    *a = 123;
}

// 变量地址(普通变量与数组变量)
void StringPointer()
{
    int a;
    int *p = &a;

    //普通指针
    printf("请输入一个整数：");
    scanf("%d", p);
    printf("a=%d\n", a);

    //数组指针：是第一个元素的地址
    char str[128];
    printf("请输入一个字符串：");
    scanf("%s", str);
    printf(" str=%s\n", str);
    printf("&str=%p\n", str);
    printf("&str=%p\n", &str[0]);
}

// 数组指针内存
void StringPointerMem()
{
    char a[] = "hello";
    int b[5] = {1, 2, 3, 4, 5};
    float c[5] = {1.1, 2.2, 3.3, 4.4, 5.5};
    double d[5] = {1.1, 2.2, 3.3, 4.4, 5.5};
    printf("char:   %p,%p,%p\n", &a[0], &a[1], &a[2]);
    printf("int:    %p,%p,%p\n", &b[0], &b[1], &b[2]);
    printf("float:  %p,%p,%p\n", &c[0], &c[1], &c[2]);
    printf("double: %p,%p,%p\n", &d[0], &d[1], &d[2]);
}

// 指针变量是左值，字符串名是地址常量
void StringLenByPointer()
{
    char str[] = "hello world";
    char *target = str;
    int count = 0;
    while (*target++ != '\0')
    {
        count++;
    }
    printf("字符串长度为：%d\n", count);
}

// 指针数组:是一个数组，每个数组元素存放一个指针变量
// 格式：int *p[n]
void PointerArray()
{
    char *p[3] = {
        "clang",
        "java",
        "python"};

    for (int i = 0; i < 3; i++)
    {
        printf("%s\n", p[i]);
    }
}

// 数组指针：指向数组的指针
// 格式：int (*p)[n]
void ArrayPointer()
{
    int tmp[3] = {1, 2, 3};
    int(*p)[3] = &tmp;
    for (int i = 0; i < 3; i++)
    {
        printf("%d\n", *(*p + i));
    }
}

// 二维数组
void ArrayD2()
{
    int arr[4][5] = {0};
    int i, j, k = 0;
    for (i = 0; i < 4; i++)
    {
        for (j = 0; j < 5; j++)
        {
            arr[i][j] = k++;
        }
    }

    printf("%p\n", *(arr + 1));
    printf("%p\n", arr[1]);
    printf("%p\n", &arr[1][0]);
    printf("%d\n", **(arr + 1));

    //加引用，减引用
    printf("%d\n", **(arr + 1) + 3);
    printf("%d\n", arr[1][3]);

    //结论：
    // *(arr+i)==arr[i]
    // *(*(arr+i)+j)=arr[i][j]
    // *(*(*(arr+i)+j)+k)=arr[i][j][k]
}

//指针加引用，减引用
void ArrayPointerRef()
{
    int arr[][3] = {{1, 2, 3}, {3, 4, 5}};
    int(*p)[3] = arr;

    printf("%d\n", **(p + 1));
    printf("%d\n", **(arr + 1));
    printf("%d\n", arr[1][0]);
    printf("%d\n", *(*(p + 1) + 2));
    printf("%d\n", *(*(arr + 1) + 2));
    printf("%d\n", arr[1][2]);
}

// Void指针：可以指向任何地址(谨慎转换，谨慎使用)
void VoidPointer()
{
    int num = 1024;
    int *pi = &num;
    char *ps = "hello";
    void *pv;

    pv = pi;
    printf("pi=%p, pv=%p\n", pi, pv);
    printf("pv=%p\n", (void *)pv);

    pv = ps;
    printf("ps=%p, pv=%p\n", ps, pv);
    printf("pv=%s\n", (char *)pv);
}

// NULL指针
// 原型：#define NULL ((void *)0)
void NullPointer()
{
    int *p1;        //危险
    int *p2 = NULL; //推荐
    printf("*p1 =%d\n", *p1);
    printf("NULL=%d\n", *p2);
}

void PointerPointer()
{
    char *langs[] =
        {
            "c++",
            "java",
            "ruby",
            "pthon",
            "golang"};
    char **by;
    char **tar[4];

    by = &langs[5];
    tar[0] = &langs[0];
    tar[1] = &langs[1];
    tar[2] = &langs[2];
    tar[3] = &langs[3];

    printf("my tar is %s\n", **tar);
    printf("my langs:\n");
    for (int i = 0; i < 4; i++)
    {
        printf("%s\n", *tar[i]);
    }
}

void PointerPointerDemo()
{
    int num = 99;
    int *p = &num;
    int **pp = &p;

    printf("num:%d\n", num);
    printf("*p:%d\n", *p);
    printf("**p:%d\n", **pp);
    printf("&p:%p,pp:%p\n", &p, pp);
    printf("&num:%p,p:%p,*pp:%p\n", &num, p, *pp);
}

void ArrayD2Pointer()
{
    int arr[3][4] = {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12}};
    int i, j;
    for (i = 0; i < 3; i++)
    {
        for (j = 0; j < 4; j++)
        {
            printf("%4d", *(*(arr + i) + j));
            /* code */
        }
        printf("\n");
    }
}

void ArrayD2PointerV3()
{
    int arr[3][4] = {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12}};
    int(*p)[3][4] = &arr;
    int i, j;

    for (i = 0; i < 3; i++)
    {
        for (j = 0; j < 4; j++)
        {
            printf("%4d", *(*(*p + i) + j));
        }
        printf("\n");
    }
}

// 常量与指针
// 指针可以修改为指向不同的常量/变量
// 可以通过解引用来读取指针指向的数据
void ConstWithPointer()
{
    int num = 1;
    const int cnum = 2;
    const int *pc = &cnum;

    printf("cnum:%d,&cnum:%p\n", cnum, &cnum);
    printf("*pc:%d,pc:%p\n", *pc, pc);

    pc = &num;
    printf("num:%d,&num:%p\n", num, &num);
    printf("*pc:%d,pc:%p\n", *pc, pc);

    num = 3;
    printf("*pc:%d,pc:%p\n", *pc, pc);
}

// 常量指针
void ConstPointer()
{
    int num = 1;
    const int cnum = 2;
    int *const p = &num;

    *p = 1024;
    printf("*p:%d\n", *p);

    // p = &cnum;  //不可修改
    // printf("p:%d\n", *p);
}

void Bar()
{
    struct INFO
    {
        int a;
        char b;
        double c;
    };

    struct INFO *p1;
    void *p2;

    printf("struct point size is :%ld\n", sizeof(p1));
    printf("void point size is :%ld\n", sizeof(p2));
}

//
int PointArray()
{
    int i;
    int a[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 0};
    int *p = a;
    for (i = 0; i < 10; i++)
    {
        printf("P Value:%d   a Value :%d\n", *(p++), *(a + i));
    }
    printf("\n");
    return 0;
}

void PointString()
{
    char *str = "www.dotcpp.com";
    char string[] = "Welcome to dotcpp.com";
    str[0] = 'C'; //试图修改str指向的常量区的字符串内容
    printf("%s", str);
    // return 0;
}

int main(int argc, char *argv[])
{
    // Foo();
    // Bar();
    // PointArray();
    // PointString();

    // Pointer();
    // PointerString();
    // ArrayPointerMem();
    // StringLenByPointer();
    // PointerArray();
    // ArrayPointer();
    // ArrayD2();
    // ArrayPointerRef();
    // VoidPointer();
    // NullPointer();
    // PointerPointer();

    // PointerPointerDemo();
    // ArrayD2PointerV3();
    ConstPointer();
}
