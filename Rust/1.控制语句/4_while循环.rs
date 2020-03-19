fn main() {
    foo();
    bar();
}

fn foo() {
    let mut i = 1;
    while i <= 5 {
        print!("{}", i);
        print!(" ");
        i = i + 1;
    }
    println!()
}

fn bar() {
    let array = [10, 20, 30, 40, 50, 60];
    let mut i = 0;
    while i < 6 {
        print!("{}", array[i]);
        print!(" ");
        i = i + 1;
    }
    println!()
}
