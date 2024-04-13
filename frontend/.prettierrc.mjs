// .prettierrc.mjs
/** @type {import("prettier").Config} */
export default {
  plugins: ["prettier-plugin-astro"],
  overrides: [
    {
      files: ["src/**/*.astro"],
      options: {
        parser: "astro",
        arrowParens: "avoid",
        singleQuote: true,
        bracketSpacing: true,
        endOfLine: "lf",
        semi: false,
        tabWidth: 2,
        trailingComma: "none",
      },
    },
    {
      files: ["*.js", "*.ts", "*.tsx", "*.jsx", "src/**/*.tsx"],
      options: {
        arrowParens: "avoid",
        singleQuote: true,
        bracketSpacing: true,
        endOfLine: "lf",
        semi: false,
        tabWidth: 2,
        trailingComma: "none",
      },
    },
  ],
};
