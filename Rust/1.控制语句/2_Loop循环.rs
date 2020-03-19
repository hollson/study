fn main() {
    with_loop();
    with_for_scope();
    with_for_iter();
    with_while();
}


// For范围循环
fn with_for_scope() {
    let mut result;
    for i in 1..11 {
        result = 2 * i;
        println!("2*{}={} ", i, result);
    }
}

// For迭代循环
fn with_for_iter() {
    let fruits = ["mango", "apple", "banana", "litchi", "watermelon"];
    for a in fruits.iter() {
        println!("{} ", a);
    }

}

// While循环
fn with_while() {
    let mut i = 1;
    while i <= 10 {
        println!("{} ", i);
        i = i + 1;
    }
    println!()
}