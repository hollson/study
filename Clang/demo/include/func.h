#ifndef _FUNC_H
#define _FUNC_H

// 用宏定义来代替全局变量
#define OS "Linux"
#define WEB_URL "http://www.xxx.com"
#define WEB_NAME "XXX"

//也可以省略 extern，不过为了程序可读性，建议都写上
long sum(int, int);
long mult(int, int);
char *getWebName();
char *getWebURL();
#endif