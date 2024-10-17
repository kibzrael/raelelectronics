/** @type {import('tailwindcss').Config} */

const plugin = require("tailwindcss/plugin");

module.exports = {
  content: ["./templates/**/*.html"],
  theme: {
    extend: {
      colors: {
        secondary: "#F45E0C",
        error: "#C91433",
        "error-light": "#FAE7EB",
        success: "#146C43",
        "success-light": "#D1F7E5",

        primary: {
          DEFAULT: "#0C68F4",
          25: "#E4EEFE",
          50: "#AECDFB",
          75: "#78ABF9",
          100: "#428AF6",
          200: "#2779F5",
          300: "#0C68F4",
          400: "#0951BE",
          500: "#063A88",
          600: "#052E6D",
          700: "#042352",
          900: "#021736",
        },
      },
      boxShadow: {
        shadow1: "-2px 2px 15px -1px rgba(113, 113, 113, 0.12)",
        shadow2: "-2px 2px 20px -1px rgba(113, 113, 113, 0.20)",
        shadow3: "-2px 2px 10px -1px rgba(113, 113, 113, 0.15)",
      },
    },
  },
  plugins: [
    plugin(({ addUtilities }) => {
      addUtilities({
        ".font-display-1": { "@apply text-4xl md:text-6xl font-semibold": {} },
        ".font-display-2": { "@apply text-3xl sm:text-5xl font-semibold": {} },
        ".font-h1": { "@apply font-medium text-3xl md:text-4xl": {} },
        ".font-h2": { "@apply font-medium text-2xl md:text-3xl": {} },
        ".font-h3": { "@apply font-medium text-xl md:text-3xl": {} },
        ".font-h4": { "@apply font-medium text-lg md:text-2xl": {} },
        ".font-h5": { "@apply font-medium text-base md:text-xl": {} },
        ".font-h6": { "@apply font-medium text-sm md:text-lg": {} },
        ".font-body-xl": { "@apply font-light text-lg md:text-xl": {} },
        ".font-body-lg": { "@apply font-light text-base md:text-lg": {} },
        ".font-body-md": { "@apply font-light text-sm md:text-base": {} },
        ".font-body-sm": { "@apply font-light text-xs md:text-sm": {} },
        ".font-body-xs": { "@apply font-light text-xs md:text-xs": {} },
        ".font-caption-lg": { "@apply font-medium text-base md:text-lg": {} },
        ".font-caption-md": { "@apply font-medium text-sm md:text-base": {} },
        ".font-caption-sm": { "@apply font-medium text-xs md:text-sm": {} },
        ".font-button-lg": { "@apply font-semibold text-base md:text-lg": {} },
        ".font-button-sm": { "@apply font-medium text-sm md:text-base": {} },
        ".font-overline-lg": { "@apply text-base md:text-base font-light": {} },
        ".font-overline-sm": { "@apply text-xs md:text-xs font-light": {} },
        ".font-underline-primary": {
          "@apply text-base md:text-base underline": {},
        },
      });
    }),
  ],
};
