/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,ts}", "./node_modules/preline/preline.js"],
    theme: {
        screens: {
            sm: "640px",
            md: "768px",
            lg: "1024px",
            tbreak: "1350px",
            xl: "1280px",
        },
        extend: {},
    },
    plugins: [require("preline/plugin")],
};
