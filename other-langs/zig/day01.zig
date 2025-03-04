const std = @import("std");
const utils = @import("utils.zig");

const isPartTwo = true;

pub fn main() !void {
    // init
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    var lefts = std.ArrayList(u32).init(arena.allocator());
    var rights = std.ArrayList(u32).init(arena.allocator());
    var rightNumOccurrences = std.AutoHashMap(u32, u32).init(arena.allocator());

    // get iterator over each line
    const args = try std.process.argsAlloc(arena.allocator());
    defer std.process.argsFree(arena.allocator(), args);
    if (args.len < 2) {
        std.debug.print("\x1b[31mERROR: must provide filename\x1b[0m\n", .{});
        return;
    }
    var lines_it = try utils.getLinesIteratorForFile(arena.allocator(), args[1]);

    // iterate over each line
    while (lines_it.peek() != null) {
        const line = lines_it.next().?;
        var space_it = std.mem.tokenizeAny(u8, line, " ");
        if (space_it.peek() == null) continue;
        const left_str = space_it.next().?;
        const right_str = space_it.next().?;
        const left = try std.fmt.parseInt(u32, left_str, 10);
        const right = try std.fmt.parseInt(u32, right_str, 10);
        try lefts.append(left);
        try rights.append(right);
        if (rightNumOccurrences.get(right) == null) {
            try rightNumOccurrences.put(right, 0);
        }
        try rightNumOccurrences.put(right, rightNumOccurrences.get(right).? + 1);
    }

    if (!isPartTwo) {
        // sort lists
        std.sort.heap(u32, lefts.items, {}, std.sort.asc(u32));
        std.sort.heap(u32, rights.items, {}, std.sort.asc(u32));

        // add up differences
        var sum: u64 = 0;
        for (lefts.items, rights.items) |l, r| {
            sum += if (l > r) (l - r) else (r - l);
        }
        std.debug.print("sum of diffs: {d}\n", .{sum});
    } else {
        var sum: u64 = 0;
        for (lefts.items) |val| {
            if (rightNumOccurrences.get(val)) |numOccurrences|
                sum += val * numOccurrences;
        }
        std.debug.print("sum of all pt. 2 vals: {d}\n", .{sum});
    }
}
