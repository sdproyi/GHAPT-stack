/** @type {import('tailwindcss').Config} */
export const content = [
  './templates/**/*.templ',
  './views/**/*.templ',
];
export const theme = {
  extend: {
    colors:{
      background: "#F1F1F1"
    },
    fontFamily: {
      poppins: ['poppins'],
    },
  },
}