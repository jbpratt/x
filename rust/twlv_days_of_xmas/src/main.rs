// Prints the song "The Twelve Days of Christmas"
fn main() {
    let days = [
        "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth",
        "tenth", "eleventh", "twelveth",
    ];
    let pres = [
        "A partridge in a pear tree",
        "Two turtle doves",
        "Three French hens",
        "Four calling birds",
        "Five gold rings",
        "Six geese a laying",
        "Seven swans a swimming",
        "Eight maids a milking",
        "Nine ladies dancing",
        "Ten lords a leaping",
        "Eleven pipers piping",
        "12 drummers drumming",
    ];
    for i in 0..12 {
        println!(
            "On the {} day of Christmas, my true love gave to me...",
            days[i]
        );
        for x in (0..i + 1).rev() {
            println!("{}", pres[x]);
        }
    }
}
