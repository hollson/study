function identity<T>(arg: T): T {
    return arg;
}

let output1 = identity<string>("myString"); //泛型参数
let output2 = identity("myString");         //类型推论 
console.log(output1, output2);

// 泛型变量
