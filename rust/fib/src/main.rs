// Generate the nth nonacci number
use std::io;

fn main() {
    println!("Enter a number");

    let mut temp = String::new();
    io::stdin()
        .read_line(&mut temp)
        .expect("failed to read line");

    let n: u32 = temp.trim().parse().expect("a number");

    let mut a = 0;
    let mut b = 1;

    if n == 1 {
        println!("{}", a);
    } else if n == 2 {
        println!("{}", b);
    } else {
        for _ in 2..n {
            let c = a + b;
            a = b;
            b = c;
        }
        println!("{}", b);
    }
}
