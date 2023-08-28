const std = @import("std");
const stuff = @import("stuff/rucksack_reorganization.zig");

pub fn main() !void {
    try stuff.rucksack_reorganization();
}
