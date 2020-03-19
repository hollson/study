// 默认对象{}也是键值对格式，但JS的对象的键必须是字符串。

// Map字典
(function () {
    // 方式1
    var m = new Map([['Michael', 95], ['Bob', 75]]);

    //方式2
    var m = new Map();
    m.set('Adam', 67);
    m.set('Bob', 59);
    m.has('Adam');
    m.get('Adam');
    m.delete('Bob');
    m.get('Adam'); // undefined

    console.log(m);
})();

// Set集合  
(function () {
    // 方式1
    var s = new Set(); // 空Set

    // 方式2
    var s = new Set([1, 2, 3, 3, '3']);  //过滤重复
    s.add(4);
    s.delete(2);

    console.log(s);
})();

