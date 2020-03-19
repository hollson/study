//1. 类定义
class Greeter {
    greeting: string;
    constructor(message: string) {
        this.greeting = message;
    }
    greet() {
        return "Hello, " + this.greeting;
    }
}

let greeter = new Greeter("world");
console.log(greeter.greet());

// ==========================================
//2. 继承
class Animal {
    name: string;
    constructor(theName: string) { this.name = theName; }
    move(distanceInMeters: number = 0) {
        console.log(`${this.name} moved ${distanceInMeters}m.`);
    }
}
//super()会执行基类的构造方法。
class Snake extends Animal {
    constructor(name: string) { super(name); }
    move(distanceInMeters = 5) {
        console.log("Slithering...");
        super.move(distanceInMeters);
    }
}

class Horse extends Animal {
    constructor(name: string) { super(name); }
    move(distanceInMeters = 45) {
        console.log("Galloping...");
        super.move(distanceInMeters);
    }
}

let sam = new Snake("Sammy the Python");
let tom: Animal = new Horse("Tommy the Palomino");

sam.move();
tom.move(34);

// ==========================================

//3. 访问保护
// 默认为public
class Animal {
    public name: string;
    public constructor(theName: string) { this.name = theName; }
    public move(distanceInMeters: number) {
        console.log(`${this.name} moved ${distanceInMeters}m.`);
    }
}
//当成员被标记成private时，它就不能在声明它的类的外部访问。比如：
class Animal {
    private name: string;
    constructor(theName: string) { this.name = theName; }
}

new Animal("Cat").name; // Error: 'name' is private;

//protected：修饰符与private修饰符的行为很相似，但有一点不同，protected成员在派生类中仍然可以访问。例如：

//readonly关键字将属性设置为只读的。 只读属性必须在声明时或构造函数里被初始化。



//4. 存取器
//TypeScript支持通过getters / setters来截取对对象成员的访问。 它能帮助你有效的控制对对象成员的访问。
//下面来看如何把一个简单的类改写成使用get和set。 首先，我们从一个没有使用存取器的例子开始。

//5. 静态属性


//6. 抽象类

//7. 构造函数


