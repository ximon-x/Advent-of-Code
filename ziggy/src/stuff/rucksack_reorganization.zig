const std = @import("std");
const testing = std.testing;

fn generate_alphabets() [52]usize {
    var chars: [52]usize = undefined;

    for (0..26) |i| {
        chars[i] = i + 'a';
        chars[i + 26] = i + 'A';
    }

    return chars;
}

fn find_unique(line: []u8) usize {
    const n = line.len;
    var unique: usize = 0;

    for (0..n / 2) |i| {
        for (n / 2..n) |j| {
            if (line[i] == line[j]) {
                unique = line[j];
            }
        }

        if (unique != 0) {
            return unique;
        }
    }

    return unique;
}

pub fn rucksack_reorganization() !void {
    const alphabets = generate_alphabets();
    var score: usize = 0;

    // Opens the necessary file
    var file = try std.fs
        .cwd()
        .openFile("/home/simon/Projects/Personal/Advent-of-Code/ziggy/src/assets/rucksack_reorganization_input.txt", .{});

    // Defers closing the file
    defer file.close();

    var buf: [1024]u8 = undefined;
    while (try file.reader().readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const unique = find_unique(line);

        for (0..alphabets.len) |i| {
            if (unique == alphabets[i]) {
                score += i + 1;
            }
        }
    }
    std.debug.print("Final Score: {d}\n", .{score});
}
