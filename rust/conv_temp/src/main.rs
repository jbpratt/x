// Convert temperatures from Fahrenheit to Celsius
use std::io;

fn main() {
    println!("Enter a temp in Fahrenheit");

    let mut temp = String::new();
    io::stdin()
        .read_line(&mut temp)
        .expect("Failed to read line");

    let temp: f32 = temp.trim().parse().expect("a float 32");

    println!("{}", ((temp - 32.0) * (5.0 / 9.0)))
}
