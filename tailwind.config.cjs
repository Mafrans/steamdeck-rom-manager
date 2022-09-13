/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{ts,tsx}", "./src-uploader/**/*.{ts,tsx}"],
  theme: {
    extend: {
      screens: {
        xs: "320px",
      },
    },
  },
  plugins: [],
};
