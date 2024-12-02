const std = @import("std");
const utils = @import("utils.zig");

const IS_PART_TWO = true;

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const args = try std.process.argsAlloc(arena.allocator());
    if (args.len < 2) {
        std.debug.print("\x1b[31mERROR: must provide filename\x1b[0m\n", .{});
        return;
    }
    var lines_it = try utils.getLinesIteratorForFile(arena.allocator(), args[1]);

    var num_good_lines: u32 = 0;
    while (lines_it.next()) |line| {
        if (try isLineSafe(line)) num_good_lines += 1 else if (IS_PART_TWO) {
            const num_numbers = utils.countTokensScalar(u8, line, ' ');
            for (0..num_numbers) |i_to_omit| {
                var num_it = std.mem.tokenizeScalar(u8, line, ' ');
                var sliced_str = std.ArrayList(u8).init(arena.allocator());
                var i: usize = 0;
                while (num_it.next()) |num_str| : (i += 1) {
                    const num_plus_space = try std.fmt.allocPrint(arena.allocator(), "{s} ", .{num_str});
                    if (i != i_to_omit) try sliced_str.appendSlice(num_plus_space);
                }

                if (try isLineSafe(sliced_str.items)) {
                    num_good_lines += 1;
                    break;
                }
            }
        }
    }

    std.debug.print("num good lines: {d}\n", .{num_good_lines});
}

const MAX_DIFF = 3;
fn isLineSafe(line: []const u8) !bool {
    var is_increasing: ?bool = null;
    var last_num: i32 = -1;

    var nums_it = std.mem.tokenizeScalar(u8, line, ' ');
    while (nums_it.next()) |num_str| {
        const num = try std.fmt.parseInt(i32, num_str, 10);
        if (last_num == -1) {
            last_num = num;
            continue;
        } else {
            // difference must be between 1 and 3
            const abs_diff = utils.intAbs(num - last_num);
            if (abs_diff > MAX_DIFF or abs_diff == 0) return false;

            // must always either increase or decrease
            if (is_increasing == null) {
                is_increasing = num > last_num;
            } else {
                if (is_increasing.? and num < last_num) return false;
                if (!is_increasing.? and num > last_num) return false;
            }

            last_num = num;
        }
    }
    return true;
}
