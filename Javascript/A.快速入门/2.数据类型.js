
// 数据类型
(function DataType() {
    var a = typeof (8);                  //"number"
    var b = typeof (1.2);                //"number"
    var c = typeof ("11");               //"string"
    var d = typeof ('a');                //"string"
    var e = typeof (true);               //"boolean"
    var f = typeof (function () { });    //"function"
    var g = typeof ([1, 2, 3]);          //"object"
    var h = typeof (x);                  //"undefined"
    var i = 0 / 0;                       // NaN
    var j = Infinity;                    //无限大 

    console.log(a, b, c, d, e, f, g, h, i, j);
})();

// 变量值
(function VarValue() {
    var a = false == 0;             // true，会自动转换数据类型，比较诡异
    var b = false === 0;            // false，不自动转换数据类型
    var c = NaN === NaN;            // false
    var d = isNaN(NaN);             // true
    var e = 1 / 3 === (1 - 2 / 3);  //false， 浮点数的相等比较
    var f = Math.abs(1 / 3 - (1 - 2 / 3)) < 0.0000001; // true

    console.log(a, b, c, d, e, f);
})();



