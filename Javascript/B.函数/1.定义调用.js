// 由于JavaScript的函数也是一个对象

// ================================================
// 函数定义方式1
function abs(x) {
    if (x >= 0) {
        return x;
    } else {
        return -x;
    }
}

// 函数定义方式2
var abs = function (x) {
    if (x >= 0) {
        return x;
    } else {
        return -x;
    }
};

var a = abs(10, 'blablabla'); // 10
var b = abs(-9, 'haha', 'hehe', null); // 9
console.log(a, b)

// ================================================

// 关键字arguments
//它只在函数内部起作用，并且永远指向当前函数的调用者传入的所有参数。arguments类似Array但它不是一个Array
function foo(x) {
    console.log('x = ' + x); // 10
    for (var i = 0; i < arguments.length; i++) {
        console.log('arg ' + i + ' = ' + arguments[i]); // 10, 20, 30
    }
}
foo(10, 20, 30);


// ================================================
// rest参数,固定参数之外的剩余参数
function foo(a, b, ...rest) {
    console.log('a = ' + a);
    console.log('b = ' + b);
    console.log(rest);
}

foo(1, 2, 3, 4, 5);
// 结果:
// a = 1
// b = 2
// Array [ 3, 4, 5 ]

foo(1);
// 结果:
// a = 1
// b = undefined
// Array []
