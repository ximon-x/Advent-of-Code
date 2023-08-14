use std::fs;

pub fn calorie_counting() {
    let file_path = "Sensitive Personal Data";
    let mut best_elf: [u32; 2] = [0, 0];
    let mut curr_elf: [u32; 2] = [0, 0];

    let contents =
        fs::read_to_string(file_path).expect("could not read from file");

    for line in contents.lines() {
        match line {
            "" => {
                if curr_elf[1] > best_elf[1] {
                    best_elf[1] = curr_elf[1];
                    best_elf[0] = curr_elf[0];
                }

                curr_elf[0] += 1;
                curr_elf[1] = 0;
            }
            _ => {
                let calories: u32 =
                    line.trim().parse().expect("could not convert to u32");

                curr_elf[1] += calories;
            }
        }
    }

    println!(
        "The best elf is {} with {} points",
        best_elf[0], best_elf[1]
    );
}
