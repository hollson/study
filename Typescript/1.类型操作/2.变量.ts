// 1. javascript里面的var
// var声明可以在包含它的函数，模块，命名空间或全局作用域内部任何位置被访问
// 也就是var会覆盖它作用域之外的值
console.log("===================var==================")
function foo(shouldInitialize: boolean) {
    if (shouldInitialize) {
        var x = 10;
    }
    return x;
}

console.log(foo(true));  // returns '10'x
console.log(foo(false)); // returns 'undefined'

// 2. javascript-var引发的问题
console.log("===================isseue==================")
for (var i = 0; i < 10; i++) {
    setTimeout(function () {
        console.log(i);
    }, 100 * i);
}
// 分析：setTimeout在若干毫秒后执行一个函数，并且是在for循环结束后。 
// for循环结束后，i的值为10。 所以当函数被调用的时候，它会打印出 10！

// 3. javascript-var 补救措施
console.log("===================solve==================")
for (var i = 0; i < 10; i++) {
    (function (i) {
        setTimeout(function () { console.log(i); }, 100 * i);
    })(i);
}

// 4. let 完美上场
console.log("===================let==================")
for (let i = 0; i < 10; i++) {
    setTimeout(function () {
        console.log(i);
    }, 100 * i);
}


// 5. let作用域
console.log("===================domain==================")
function f(condition, x) {
    if (condition) {
        let x = 100;
        return x;
    }
    return x;
}

f(false, 0); // returns 0
f(true, 0);  // returns 100


// 6. cost常量
console.log("===================const==================")
const numLivesForCat = 9;
const kitty = {
    name: "Aurora",
    numLives: numLivesForCat,
}

// // Error
// kitty = {
//     name: "Danielle",
//     numLives: numLivesForCat
// };

// all "okay"
kitty.name = "Rory";
kitty.name = "Kitty";
kitty.numLives--;

console.log(kitty)
