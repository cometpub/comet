import Handlebars from "handlebars";
import fs from "fs";
import path from "path";
import esbuild from "esbuild";

const OUTDIR = path.resolve("../static");
const OUTDIR_CSS = path.resolve("../static/css");

function buildCSS() {
  esbuild.buildSync({
    bundle: true,
    minify: false,
    entryPoints: ["src/css/*.css"],
    loader: {
      ".eot": "file",
      ".svg": "file",
      ".ttf": "file",
      ".woff": "file",
    },
    outdir: OUTDIR_CSS,
    external: ["/static/*"],
  });
}

function registerPartials() {
  const partialsDir = path.resolve("./src/xslt/partials");

  let files = fs.readdirSync(partialsDir);

  for (const file of files) {
    const filepath = path.join(partialsDir, file);
    const name = path.parse(filepath).name;
    const template = fs.readFileSync(filepath).toString();

    Handlebars.registerPartial(name, template);
  }

  // register bundled CSS as partials
  files = fs.readdirSync(OUTDIR_CSS);

  for (const file of files) {
    const filepath = path.join(OUTDIR_CSS, file);
    const name = `css-${path.parse(filepath).name}`;
    const template = fs.readFileSync(filepath).toString();

    Handlebars.registerPartial(name, template);
  }
}

function renderXslt() {
  const xsltDir = path.resolve("./src/xslt");

  if (!fs.existsSync(OUTDIR)) {
    fs.mkdirSync(OUTDIR);
  }

  const files = fs.readdirSync(xsltDir);

  for (const file of files) {
    const filepath = path.join(xsltDir, file);
    const name = path.parse(filepath).name;
    const stats = fs.statSync(filepath);

    if (stats.isFile()) {
      const source = fs.readFileSync(filepath).toString();
      const template = Handlebars.compile(source);
      const result = template();

      fs.writeFileSync(path.join(OUTDIR, `${name}.xsl`), result);
    }
  }
}

function main() {
  buildCSS();
  registerPartials();
  renderXslt();
}

main();
