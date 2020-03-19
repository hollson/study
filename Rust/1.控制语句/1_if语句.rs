fn main() {
    foo();
    bar();
}

// If普通语句
fn normal() {
    let a = 5;
    let b = 6;
    if a != b {
        if a > b {
            println!("a is greater than b");
        } else {
            println!("a is less than b");
        }
    } else {
        println!("a is equal to b");
    }
}

// If三目运算
fn bar() {
    let a = if true { 1 } else { 2 };
    println!("value of a is: {}", a);
}
