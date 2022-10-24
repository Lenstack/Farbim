import {DefaultTheme} from "styled-components";

const Light: DefaultTheme = {
    colors: {
        background: {
            primary: "#e8fae5",
            secondary: "#3e4a3c",
            third: "#fafafa"
        },
        info: "#43c6da",
        success: "#4be18d",
        warning: "#e59347",
        danger: "#e65454",
        white: "#FFFFFF",
        dark: "#1B1B1B",
        gradiant:"linear-gradient(62deg, #FBAB7E 0%, #F7CE68 100%)"
    },
    fonts: {
        size: "0.8rem",
        color: "#2c2c2c"
    }
}

const Dark: DefaultTheme = {
    colors: {
        background: {
            primary: "#202525",
            secondary: "#589e79",
            third: "#2c3333"
        },
        info: "#8AD7E3",
        success: "#8AE3B1",
        warning: "#E3B58A",
        danger: "#E38A8A",
        white: "#FFFFFF",
        dark: "#1B1B1B",
        gradiant:"linear-gradient(62deg, #67a1c1 0%, #65c7a4 100%)"
    }, fonts: {
        size: "0.8rem",
        color: "#ffffff"
    }
}

export const Themes = {
    Light,
    Dark
}