 当代码块拥有资源时，它被称为所有权。 代码块创建一个包含资源的对象。 当控件到达块的末尾时，对象将被销毁，资源将被释放。
 fn main()
 {
     let x=10;
     let y=x;  //x的所有权被转移到变量y，并且变量x被销毁
     println!("value of x :{}",x);
 }


// 所有权和函数
// #以下是一些复制类型：
//  所有整数类型，如u32。
//  布尔类型-bool，其值为true或false。
//  所有浮动类型，如f64。
//  字符类型，如char。
fn main() {
    // 所有权转移
    let s = String::from("Yiibai");
    take_ownership(s);

    // 复制特性
    let ch = 'a';
    moves_copy(ch);
    println!("{}", ch);
}

fn take_ownership(s: String) {
    println!("{}", s);
}

fn moves_copy(c: char) {
    println!("{}", c);
}


fn main()  
{  
  let x= gives_ownership();  
  println!("value of x is {}",x);  
 }  
 
 //从函数返回值也会转移所有权
fn gives_ownership()->u32  
{  
     let y=100;  
     y  
}

