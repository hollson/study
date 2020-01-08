// 1. 布尔
var isDone = false;
// 2. 数字
var decLiteral = 6; //十进制
var hexLiteral = 0xf00d; //十六进制
var binaryLiteral = 10; //二进制
var octalLiteral = 484; //八进制
// 3. 字符串 (单引号、双引号、模版)
var nick = "bob";
var nick_name = "Gene";
var age = 37;
var sentence = "Hello, my name is " + nick_name + ".\nI'll be " + (age + 1) + " years old next month.";
// 4. 数组（一般数组 和 泛型数组）
var list1 = [1, 2, 3];
var list2 = [1, 2, 3];
// 5. 元组
var x;
x = ['hello', 10]; //赋值
x[0] += ' world';
console.log(x); //读取 
// 6. 枚举 
var Color;
(function (Color) {
    Color[Color["Red"] = 1] = "Red";
    Color[Color["Green"] = 2] = "Green";
    Color[Color["Blue"] = 3] = "Blue";
})(Color || (Color = {})); //(可指定起始索引值)
var c1 = Color.Green; //读枚举下标
var c2 = Color[2]; //读枚举字符
// 7. Any
var notSure = 4;
notSure = "maybe a string instead";
notSure = false;
// 8. Void
var unusable = undefined; //只能为它赋undefined和null：
function warnUser() {
    console.log("This is my warning message");
}
// 9. Null和Undefined
var u = undefined;
var n = null;
// 10. Never (永不存在的值的类型)
function error(msg) {
    throw new Error(msg);
}
function fail() {
    return error("failed");
}
function infiniteLoop() {
    while (true) { }
}
create({ prop: 0 }); // OK
create(null); // OK
// 12. 断言 (泛型推断 和 as推断)
var some = "this is a string";
var len1 = some.length;
var len2 = some.length;
