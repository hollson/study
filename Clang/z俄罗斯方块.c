/********************************************************************/
// Name:俄-罗-斯-方-块 V1.0
// Author:Giant
// Time: 2015/5/17
/********************************************************************/
#include "graphics.h"
#include "time.h"
#include "math.h"
#include "dos.h"
#include "conio.h"
#define NULL 0
#define False 0
#define True 1
#define REVOLVE 5
#define DOWN 2
#define LEFT 3
#define RIGHT 4
#define TIMER 0x1c
#define ESC 27
#define DOWN_MAX 420
void interrupt (*oldtimer)(void);
static unsigned grade = 0;
char msg[10] = "Grade : ";
char msg1[10] = "\0";
typedef struct boxes
{
    unsigned int box;
    int color;
    int next;
} DESIGN;
DESIGN a[19] = {35008, 2, 1, /*  The 19 shapes of boxes ,includes size,color and the next number*/
                3712, 2, 2,
                50240, 2, 3,
                736, 2, 0,
                1100, YELLOW, 5,
                2272, YELLOW, 6,
                51328, YELLOW, 7,
                3616, YELLOW, 4,
                35904, 4, 9,
                1728, 4, 8,
                19584, 12, 11,
                3168, 12, 10,
                1248, 5, 13,
                35968, 5, 14,
                3648, 5, 15,
                19520, 5, 12,
                34952, 6, 17,
                3840, 6, 16,
                3264, 10, 18};
int screeninarry[19][21] = {0};
int TimerCounter = 0;
void interrupt newtimer(void)
{
    (*oldtimer)();
    TimerCounter++;
}
void SetTimer(void interrupt (*IntProc)(void))
{
    oldtimer = getvect(TIMER);
    disable();
    setvect(TIMER, IntProc);
    enable();
}
void KillTimer()
{
    disable();
    setvect(TIMER, oldtimer);
    enable();
}
void background()
{
    int x1 = getmaxx() / 2 - 200, y1 = 60, x2 = getmaxx() / 2 + 200, y2 = 420; /* x,y is    maingameinterface  of coordinate*/
    int i;
    setlinestyle(0, 1, 3);
    /* set bkground */
    setfillstyle(1, BLUE);
    bar(0, 0, getmaxx(), getmaxy());
    /*   maingameinterface  */
    setcolor(WHITE);
    setlinestyle(0, 1, 3);
    rectangle(x1, y1, x2, y2); /*   18* 20  & 360 * 400 */
    setfillstyle(1, LIGHTBLUE);
    bar(x1, y1, x2, y2);
    /* set outbox left and right  */
    setcolor(WHITE);
    setlinestyle(0, 1, 3);
    rectangle(x2 + 20, 80, x2 + 100, 160);
    setfillstyle(1, LIGHTBLUE);      /*   setfill maingameinterface     */
    bar(x2 + 20, 80, x2 + 100, 160); /*	  setfill occupation box	*/
    i = 20;
    setlinestyle(0, 1, 3);
    while (i < 80)
    {
        line(x2 + 20, 80 + i, x2 + 100, 80 + i);
        i = i + 20;
    }
    i = 20;
    while (i < 80)
    {
        line(x2 + 20 + i, 80, x2 + 20 + i, 160);
        i = i + 20;
    }
    /* this is about of the game ,include opreta and author and so on*/
    setlinestyle(0, 1, 3);
    rectangle(x2 + 10, 240, x2 + 110, DOWN_MAX);
    bar(x2 + 10, 240, x2 + 110, DOWN_MAX);
    moveto(x2 + 10, 250);
    outtext("Welcomeplay!");
    moveto(x2 + 10, 270);
    outtext("'P' is pause");
    moveto(x2 + 10, 290);
    outtext("'R' is renew");
    moveto(x2 + 10, 310);
    outtext("'Esc' is exit");
    moveto(x2 + 10, 340);
    outtext(" ----by Giant");
}
/*                               */
void box(unsigned int box, int color) /* box difine shape,color */
{
    int x = getmaxx() / 2 + 200 + 20, y = 80; /* set dutum of x,y  */
    int i = getmaxx() / 2 + 200 + 20, n;
    unsigned int max = box;
    unsigned int mask = 32768;
    /* init boxes */
    int k = 20;
    setfillstyle(1, LIGHTBLUE);
    bar(x, 80, x + 80, 160);
    k = 20;
    setcolor(WHITE);
    setlinestyle(0, 1, 3);
    while (k < 80)
    {
        line(x, 80 + k, x + 80, 80 + k);
        k = k + 20;
    }
    k = 20;
    while (k < 80)
    {
        line(x + k, 80, x + k, 160);
        k = k + 20;
    }
    /* boxes it is ok */
    setfillstyle(1, color);
    setlinestyle(0, 1, 3);
    for (n = 0; n < 16; n++)
    {
        if (mask & max)
        {
            rectangle(x, y, x + 20, y + 20);
            bar(x, y, x + 20, y + 20);
        }
        x = x + 20;
        if (x - i == 80)
        {
            x = (x - i) % 80 + i;
            y = y + 20;
        }
        mask = mask >> 1;
    }
}
double killfullline(int lasty)
{
    int vga = getmaxx() / 2 - 200;
    int x = vga, y = lasty, yy = lasty;
    int n = 0;
    int count = 0, i = 0;
    int color;
    int score = 0;
    for (y = lasty; y >= lasty - 60; y = y - 20)
    {
        count = 0;
        x = vga;
        for (i = 0; i < 20; i++, x = x + 20)
        {
            if (screeninarry[(y - 60) / 20][(x - vga) / 20] == 2)
            {
                count++;
            }
        }
        if (20 == count)
        {
            n = 0;
            score = 0;
            x = vga;
            n++;
            score = score + 10 * pow(2, n);
            setfillstyle(1, LIGHTBLUE);
            setcolor(LIGHTBLUE);
            for (x = vga, i = 0; i < 20; x = x + 20, i++) /* clean fullline */
            {
                rectangle(x, y, x + 20, y + 20);
                bar(x, y, x + 20, y + 20);
                screeninarry[(y - 60) / 20][(x - vga) / 20] = 0;
            }
            for (x = vga, i = 0, yy = y; yy >= 60; x = x + 20, i++)
            {
                if (i == 20)
                {
                    i = 0;
                    x = vga;
                    /* killfullline(yy); */
                    yy = yy - 20;
                }
                if (screeninarry[(yy - 20 - 60) / 20][(x - vga) / 20] == 2)
                {
                    color = getpixel(x + 10, yy - 10); /* get  color of this box*/
                    screeninarry[(yy - 20 - 60) / 20][(x - vga) / 20] = 0;
                    setfillstyle(1, LIGHTBLUE);
                    bar(x, yy - 20, x + 20, yy);
                    setcolor(LIGHTBLUE);
                    rectangle(x, yy - 20, x + 20, yy);
                    setfillstyle(1, color);
                    setlinestyle(0, 1, 3);
                    setcolor(WHITE);
                    rectangle(x, yy, x + 20, yy + 20);
                    bar(x, yy, x + 20, yy + 20);
                    screeninarry[(yy - 60) / 20][(x - vga) / 20] = 2;
                }
            }
            y = y + 20;
        } /*if (20==count) */
    }     /*	for	*/
    return score;
} /*	killfullline	*/
int drawshape(int startx, int starty, int n, int flag) /*  set start,coordinate of x,y */
{
    int x, y;
    int vga = getmaxx() / 2 - 200;
    int i = 0;
    int var;
    int x1 = getmaxx() / 2 - 200, y1 = 60, x2 = getmaxx() / 2 + 200, y2 = 420;
    int bottom;
    unsigned int mask = 32768;
    unsigned int max = a[n].box;
    setfillstyle(1, LIGHTBLUE);
    if (flag == 0)
        var = 1;
    else
        var = 2;
    setfillstyle(1, BLUE);
    bar(getmaxx() / 2 - 120, 20, getmaxx() / 2 + 120, 50);
    settextstyle(1, 0, 3);
    outtextxy(getmaxx() / 2 - 100, 20, msg);
    /*setfillstyle(1,BLUE);
    	bar(x1,0,x2,y1);*/
    setlinestyle(0, 1, 3);
    setcolor(WHITE);
    rectangle(x1, y1, x2, y2);
    x = startx, y = starty;
    for (i = 0; i < 16; i++) /* draw picture of shape */
    {
        if (x - startx == 80)
        {
            x = startx;
            y = y + 20;
        }
        if (mask & max)
        {
            setfillstyle(1, a[n].color);
            setlinestyle(0, 1, 3);
            rectangle(x, y, x + 20, y + 20);
            bar(x, y, x + 20, y + 20);
            screeninarry[(y - 60) / 20][(x - vga) / 20] = var;
            bottom = y;
        }
        x = x + 20;
        mask = mask >> 1;
    }
    /*setfillstyle(1,BLUE);
    	bar(x1,0,x2,y1);*/
    setlinestyle(0, 1, 3);
    setcolor(WHITE);
    rectangle(x1, y1, x2, y2);
    /*	delay	*/
    return bottom;
    /* in */
}
void cleanshape(int startx, int starty, int n)
{
    int x = startx, y = starty;
    int vga = getmaxx() / 2 - 200;
    int i = 0;
    unsigned int max;
    unsigned int mask = 32768;
    max = a[n].box;
    for (i = 0; i < 16; i++) /* clean shape to bk */
    {
        if (x - startx == 80)
        {
            x = startx;
            y = y + 20;
        }
        if (max & mask)
        {
            setfillstyle(1, LIGHTBLUE);
            bar(x, y, x + 20, y + 20);
            setcolor(LIGHTBLUE);
            rectangle(x, y, x + 20, y + 20);
            screeninarry[(y - 60) / 20][(x - vga) / 20] = 0;
        }
        x = x + 20;
        mask = mask >> 1;
    }
}
int checkshape(int startx, int starty, unsigned int shapebox, int direction)
{
    int i = 0, x = startx, y = starty, tempx = startx, tempy = starty;
    int vga = getmaxx() / 2 - 200;
    unsigned int mask = 32768;
    unsigned int max = a[shapebox].box;
    /**/
    if (direction == LEFT)
    {
        for (x = startx, y = starty, i = 0; i < 16; x = x + 20, i++)
        {
            if (x - startx == 80)
            {
                x = startx;
                y = y + 20;
            }
            if (x < vga || screeninarry[(y - 60) / 20][(x - vga) / 20] == 2)
                return False;

            mask = mask >> 1;
        }
        return True;
    }
    else if (direction == RIGHT)
    {
        mask = 32768;
        for (x = startx, y = starty, i = 0; i < 16; x = x + 20, i++)
        {
            if (x - startx == 80)
            {
                x = startx;
                y = y + 20;
            }
            if (max & mask)
                if (x + 20 > getmaxx() - vga || screeninarry[(y - 60) / 20][(x - vga) / 20] == 2)
                    return False;
            mask = mask >> 1;
        }
        return True;
    }
    else if (direction == DOWN)
    {
        mask = 32768;
        for (x = startx, y = starty, i = 0; i < 16; x = x + 20, i++)
        {
            if (x - startx == 80)
            {
                x = startx;
                y = y + 20;
            }
            if (max & mask)
                if (y >= DOWN_MAX || screeninarry[(y - 60) / 20][(x - vga) / 20] == 2)
                    return False;
            mask = mask >> 1;
        }
        return True;
    }
    else if (direction == REVOLVE)
    {
        mask = 32768;
        for (x = startx, y = starty, i = 0; i < 16; x = x + 20, i++)
        {
            if (x - startx == 80)
            {
                x = startx;
                y = y + 20;
            }
            if (max & mask)
                if (screeninarry[(y - 60) / 20][(x - vga) / 20] == 2 || y >= DOWN_MAX || x + 20 > getmaxx() - vga || x < vga)
                    return False;
            mask = mask >> 1;
        }
        return True;
    }
    else
        printf("direction is error!\n ");
}
int main()
{
    int drive, mode;
    int i = 0;
    int bottom = 0;
    int xofshape, xofbox;
    int flag;
    int startx = 150, starty = 0;
    char ch, ch1;
    char direction;
    srand((unsigned)time(0));
    drive = DETECT;
    initgraph(&drive, &mode, "C:\\TC20\\BGI");
    /* srand((unsigned)time(0));  */
    background();
    memset(screeninarry, 0, 19 * 21 * sizeof(int));
    SetTimer(newtimer);
    while (1)
    {
        setfillstyle(1, BLUE);
        bar(getmaxx() / 2 - 120, 20, getmaxx() / 2 + 120, 50);
        settextstyle(1, 0, 3);
        outtextxy(getmaxx() / 2 - 100, 20, msg);
        if (i == 0)
        {
            xofshape = rand() % 19;
            xofbox = rand() % 19;
        }
        else
        {
            xofshape = xofbox;
            xofbox = rand() % 19;
        }
        i++;
        box(a[xofbox].box, a[xofbox].color);
        startx = getmaxx() / 2 - 200 + 120;
        starty = 60;
        while (1)
        {
            if (bioskey(1))
            {
                direction = bioskey(0);
                if (direction == ESC)
                {
                    setcolor(WHITE);
                    settextstyle(1, 0, 3);
                    moveto(getmaxx() / 2 - 150, 440);
                    outtext("Are you sure to exit ?(y/n)");
                    ch = getch();
                    if (ch == 'y' || ch == 'Y')
                        exit(0);
                    else
                    {
                        setfillstyle(1, BLUE);
                        bar(getmaxx() / 2 - 150, 430, getmaxx() / 2 + 180, 500);
                    }
                }
                if (direction == 'r' || direction == 'R')
                {
                    setcolor(WHITE);
                    settextstyle(1, 0, 3);
                    moveto(getmaxx() / 2 - 100, 440);
                    outtext("pausing......");
                    ch = getch();
                    setfillstyle(1, BLUE);
                    bar(getmaxx() / 2 - 120, 430, getmaxx() / 2 + 120, 490);
                }
                if (direction == 'a' || direction == 'A')
                {
                    if (checkshape(startx - 20, starty, xofshape, LEFT))
                    {
                        cleanshape(startx, starty, xofshape);
                        startx = startx - 20;
                        drawshape(startx, starty, xofshape, 0);
                    }
                }
                if (direction == 'd' || direction == 'D')
                {
                    if (checkshape(startx + 20, starty, xofshape, RIGHT))
                    {
                        cleanshape(startx, starty, xofshape);
                        startx = startx + 20;
                        drawshape(startx, starty, xofshape, 0);
                    }
                }
                if (direction == 's' || direction == 'S')
                {
                    if (checkshape(startx, starty + 20, xofshape, DOWN))
                    {
                        TimerCounter = TimerCounter + 18;
                    }
                }
                if (direction == 'w' || direction == 'W')
                {
                    if (checkshape(startx, starty, a[xofshape].next, REVOLVE))
                    {
                        cleanshape(startx, starty, xofshape);
                        xofshape = a[xofshape].next;
                        drawshape(startx, starty, xofshape, 0);
                    }
                }
            }
            if (TimerCounter >= 18)
            {
                if (checkshape(startx, starty + 20, xofshape, DOWN))
                {
                    drawshape(startx, starty, xofshape, 0);
                    cleanshape(startx, starty, xofshape);
                    starty = starty + 20;
                    drawshape(startx, starty, xofshape, 0);
                    TimerCounter = 0;
                }
                else /* can not down */
                {
                    if (starty <= 60)
                    {
                        setcolor(WHITE);
                        while (1)
                        {
                            if (bioskey(1))
                                exit(1);
                            moveto(getmaxx() / 2 - 100, 450);
                            bar(getmaxx() / 2 - 120, 20, getmaxx() / 2 + 120, 50);
                            settextstyle(1, 0, 3);
                            outtextxy(getmaxx() / 2 - 100, 20, msg);
                            outtext("Game  over ! ! !");
                            delay(2000);
                            setfillstyle(1, BLUE);
                            bar(getmaxx() / 2 - 120, 440, getmaxx() / 2 + 120, 460);
                        }
                    }
                    /*  down is ok var=2 */
                    bottom = drawshape(startx, starty, xofshape, 1);
                    grade = grade + killfullline(bottom);
                    itoa(grade, msg1, 10);
                    strcpy(msg + 8, msg1);
                    break;
                }
            }
        }
    }
    closegraph();
    KillTimer();
}