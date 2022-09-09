import {DefaultTheme} from "styled-components";

const Light: DefaultTheme = {
    colors: {
        background: {
            primary: "#F3F3F3",
            secondary: "#FAFAFA",
            third: "#ffffff"
        },
        info: "#43c6da",
        success: "#4be18d",
        warning: "#e59347",
        danger: "#e65454",
        white: "#FFFFFF",
        dark: "#1e1e1e",
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
            primary: "#141529",
            secondary: "#20213A",
            third: "#1F2340"
        },
        info: "#8AD7E3",
        success: "#8AE3B1",
        warning: "#E3B58A",
        danger: "#E38A8A",
        white: "#FFFFFF",
        dark: "#1e1e1e",
        gradiant:"linear-gradient(62deg, #67a1c1 0%, #65c7a4 100%)"
    }, fonts: {
        size: "0.8rem",
        color: "#dadada"
    }
}

export const Themes = {
    Light,
    Dark
}