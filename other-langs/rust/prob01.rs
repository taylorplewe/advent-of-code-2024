use std::fs;

const IS_PART_TWO: bool = true;

fn main() {
    let contents = fs::read_to_string("assets/prob.txt")
        .expect("Something went wrong reading the file");

    let mut lefts: Vec<i32> = Vec::new();
    let mut rights: Vec<i32> = Vec::new();

    for line in contents.lines() {
        let mut sides = line.split_whitespace();
        let left = sides.next().unwrap().parse::<i32>().unwrap();
        let right = sides.next().unwrap().parse::<i32>().unwrap();

        lefts.push(left);
        rights.push(right);
    }

    if IS_PART_TWO {
        let mut total_dist = 0;
        for key in lefts {
            total_dist += rights.iter().filter(|num| **num == key).count() as i32 * key;
        }
        println!("total dist: {}", total_dist);
    } else {
        lefts.sort();
        rights.sort();
    
        let mut total_dist = 0;
        for i in 0..lefts.len() {
            total_dist += (lefts[i] - rights[i]).abs();
        }
    
        println!("total dist: {}", total_dist);
    }
}
