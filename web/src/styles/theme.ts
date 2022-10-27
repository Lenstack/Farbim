import {DefaultTheme} from "styled-components";

const Light: DefaultTheme = {
    colors: {
        background: {
            primary: "#e4f0f4",
            secondary: "#464444",
            third: "#fafafa"
        },
        info: "#43c6da",
        success: "#4be18d",
        warning: "#e59347",
        danger: "#e65454",
        white: "#FFFFFF",
        dark: "#1B1B1B",
    },
    fonts: {
        size: "0.8rem",
        color: "#2c2c2c"
    },
    logo: {
        color: "invert(25%) sepia(11%) saturate(120%) brightness(88%) contrast(83%)"
    }
}

const Dark: DefaultTheme = {
    colors: {
        background: {
            primary: "#1f2323",
            secondary: "#558D65",
            third: "#2c3333"
        },
        info: "#8AD7E3",
        success: "#8AE3B1",
        warning: "#E3B58A",
        danger: "#E38A8A",
        white: "#FFFFFF",
        dark: "#1B1B1B",

    }, fonts: {
        size: "0.8rem",
        color: "#ffffff"
    },
    logo: {
        color: "invert(100%) sepia(1%) saturate(44%) hue-rotate(111deg) brightness(140%) contrast(66%)"
    }
}

export const Themes = {
    Light,
    Dark
}