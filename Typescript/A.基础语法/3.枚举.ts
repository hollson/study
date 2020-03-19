// 枚举类型 - 常量值
enum Direction {
    Up = 1,  //默认从0开始
    Down,
    Left,
    Right
}

console.log(Direction.Left);           //读枚举值
console.log(Direction[Direction.Left]);//读枚举名称


// 枚举类型 - 计算值
enum FileAccess {
    None,
    Read = 1 << 1,
    Write = 1 << 2,
    ReadWrite = Read | Write,
    G = "123".length
}

// 常量枚举
const enum Enum {
    A = 1,
    B = A * 2
}
console.log(Enum.B)

// 外部枚举
declare enum DeclareEnum {
    A = 1,
    B,
    C = 2
}
