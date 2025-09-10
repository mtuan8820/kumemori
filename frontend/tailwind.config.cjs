/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        dark: "var(--bg-dark)",
        main: "var(--bg)",
        light: "var(--bg-light)",
        text: "var(--text)",
        textmuted: "var(--text-muted)",
      }
    },
  },
  plugins: [],
};
