const ffi = require("ffi");

try {
  const lib = ffi.Library("libprint", { print: ["void", ["string"]] });
  lib.print("Hello world");
} catch {
  console.log("Missing dynamic library");
}
