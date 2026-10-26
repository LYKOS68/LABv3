// Farben freigegeben am 2026-10-26, gemessen aus hero1 (Leder), hero2 (Kassette)
// und dem Gold-Schriftzug. Quelle: .data/Brand.md §4.
module.exports = {
  content: ['./views/**/*.templ', './*.go'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Karla', '"Helvetica Neue"', 'sans-serif'],
        head: ['Montserrat', '"Helvetica Neue"', 'sans-serif'],
      },
      colors: {
        ink: '#0b0c0e',
        leather: {
          dark: '#2a1f15',
          DEFAULT: '#4d4135',
        },
        paper: '#efe6cf',
        gold: {
          shine: '#f5e6b2',
          light: '#e1cf9a',
          DEFAULT: '#cab682',
          brass: '#9a8553',
        },
        spray: '#6a8652',
        cassette: '#213a24',
        neon: {
          DEFAULT: '#52e158',
          light: '#8dff6e',
        },
      },
    },
  },
  plugins: [],
}
