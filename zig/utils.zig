const std = @import("std");

pub fn getLinesIteratorForFile(allocator: std.mem.Allocator, filename: []const u8) !std.mem.SplitIterator(u8, .scalar) {
    const file = try std.fs.cwd().openFile(filename, .{});
    defer file.close();
    const buf = try file.readToEndAlloc(allocator, std.math.maxInt(u64));
    const it = std.mem.splitScalar(u8, buf, '\n');
    return it;
}
