const std = @import("std");

pub fn getLinesIteratorForFile(allocator: std.mem.Allocator, filename: []const u8) !std.mem.SplitIterator(u8, .scalar) {
    const file = try std.fs.cwd().openFile(filename, .{});
    defer file.close();
    const buf = try file.readToEndAlloc(allocator, std.math.maxInt(u64));
    const it = std.mem.splitScalar(u8, buf, '\n');
    return it;
}

pub fn intAbs(val: i32) i32 {
    return if (val >= 0) val else (val ^ -1) + 1;
}

pub fn countTokensScalar(comptime t: type, buf: []const u8, delimiter: u8) u32 {
    var it = std.mem.tokenizeScalar(t, buf, delimiter);
    var num_tokens: u32 = 0;
    while (it.next() != null) num_tokens += 1;
    return num_tokens;
}
