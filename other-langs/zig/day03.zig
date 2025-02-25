const std = @import("std");
const utils = @import("utils.zig");

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    // ensure user provided filename
    const args = try std.process.argsAlloc(arena.allocator());
    if (args.len < 2) {
        utils.printError("must provide filename");
        return;
    }

    // get contents of file
    const filename = args[1];
    const file = std.fs.cwd().openFile(filename, .{}) catch {
        utils.printError("could not open file");
        std.process.abort();
    };
    const file_stat = try file.stat();
    const buf = try file.readToEndAlloc(arena.allocator(), file_stat.size);
}
