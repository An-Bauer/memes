const plugin = require('tailwindcss/plugin')

module.exports = {
  content: ["./web/**/*.{html,js}"],
  purge: [],
  darkMode: false,
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [
    plugin(function ({ addVariant }) {
      addVariant('bla', '&.bla');
    }),
  ],
}