// 1. 布尔
let isDone: boolean = false;


// 2. 数字
let decLiteral: number = 6;          //十进制
let hexLiteral: number = 0xf00d;     //十六进制
let binaryLiteral: number = 0b1010;  //二进制
let octalLiteral: number = 0o744;    //八进制


// 3. 字符串 (单引号、双引号、模版)
let nick: string = "bob";
let nick_name: string = `Gene`;

let age: number = 37;
let sentence: string = `Hello, my name is ${nick_name}.
I'll be ${age + 1} years old next month.`;


// 4. 数组（一般数组 和 泛型数组）
let list1: number[] = [1, 2, 3];
let list2: Array<number> = [1, 2, 3];


// 5. 元组
let x: [string, number];
x = ['hello', 10];      //赋值
x[0] += ' world';
console.log(x);         //读取 


// 6. 枚举 
enum Color { Red = 1, Green, Blue } //(可指定起始索引值)
let c1: Color = Color.Green;        //读枚举下标
let c2: string = Color[2];          //读枚举字符


// 7. Any
let notSure: any = 4;
notSure = "maybe a string instead";
notSure = false;


// 8. Void
let unusable: void = undefined;  //只能为它赋undefined和null：

function warnUser(): void {
    console.log("This is my warning message");
}


// 9. Null和Undefined
let u: undefined = undefined;
let n: null = null;


// 10. Never (永不存在的值的类型)
function error(msg: string): never {  // 返回无法达到的终点
    throw new Error(msg);
}

function fail() {                     // 推断的返回值类型为never
    return error("failed");
}


function infiniteLoop(): never {      // 返回无法达到的终点
    while (true) {
    }
}


// 11. Object (除number，string，boolean，symbol，null或undefined之外的类型)
declare function create(o: object | null): void;

create({prop: 0}); // OK
create(null); // OK


// 12. 断言 (泛型推断 和 as推断)
let some: any = "this is a string";
let len1: number = (<string>some).length;
let len2: number = (some as string).length;






