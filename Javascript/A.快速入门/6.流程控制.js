// If条件语句
(function () {
    'use strict';  //严格模式
    var score = 75;
    if (score >= 80) {
        console.log('优秀');
    } else if (score >= 60) {
        console.log('及格');
    } else {
        console.log('不及格');
    }
})();

// For循环
(function () {
    var sum = 0;
    for (var i = 1; i <= 100; i++) {
        sum = sum + i;
    }
    console.log(sum);
})();


