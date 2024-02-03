/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./view/**/*.templ", "./public/preline/dist/*.js"],
    theme: {
        extend: {},
    },
    plugins: [require("@tailwindcss/forms"), require("./public/preline/plugin")],
};
