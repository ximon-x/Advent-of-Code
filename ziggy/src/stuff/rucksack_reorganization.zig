const std = @import("std");

pub fn rucksack_reorganization() !void {
    // Captures standard output
    const stdout = std.io.getStdOut().writer();

    // Opens the necessary file
    var file = try std.fs
        .cwd()
        .openFile("/home/simon/Projects/Personal/Advent-of-Code/ziggy/src/assets/rucksack_reorganization_input.txt", .{});

    // Defers closing the file
    defer file.close();

    var buf: [1024]u8 = undefined;
    while (try file.reader().readUntilDelimiterOrEof(&buf, '\n')) |line| {
        try stdout.print("Line: {s}!\n", .{line});
    }
}
