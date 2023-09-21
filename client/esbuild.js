require("esbuild").buildSync({
    entryPoints: [
        "./src/online-stats-app.ts"
    ],
    bundle: true,
    minify: process.env.MINIFY === "true",
    outdir: "../public/assets/js",
});
