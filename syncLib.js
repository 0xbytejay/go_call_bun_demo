import fs from "fs";
import path from "path";

const srcDir = "/root/dev/bun/build/release/output/libs/linux/x86_64";
const destDir = "libs/linux/x86_64";

fs.readdirSync(destDir).forEach(file => {
    if (file.endsWith(".a")) {
        fs.unlinkSync(path.join(destDir, file));
        console.log(`Deleted: ${file}`);
    }
});

fs.readdirSync(srcDir).forEach(file => {
    if (file.endsWith(".a")) {
        fs.copyFileSync(path.join(srcDir, file), path.join(destDir, file));
        console.log(`Copied: ${file}`);
    }
});
