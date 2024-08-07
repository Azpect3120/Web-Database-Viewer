/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./web/templates/*.html",
    "./web/templates/**/*.html"
  ],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        primary: {
          100: "#EF8275",
          200: "#98023E",
        },
        secondary: {
          100: "#E6E6E6",
          200: "#808D8E",
          300: "#1E2019",
        },
      },
    },
  },
  plugins: [],
}

