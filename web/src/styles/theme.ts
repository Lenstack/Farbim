import {DefaultTheme} from "styled-components";

const Light: DefaultTheme = {
    colors: {
        background: {
            primary: "#E8E8E1",
            secondary: "#F6F6F6",
            third: "#fbfbfb"
        },
        info: "#8AD7E3",
        success: "#8AE3B1",
        warning: "#E3B58A",
        danger: "#E38A8A",
        white: "#FFFFFF"
    },
    fonts: {
        size: "1rem",
        color: "#222823"
    }
}

const Dark: DefaultTheme = {
    colors: {
        background: {
            primary: "#191C19",
            secondary: "#222823",
            third: "#2c332e"
        },
        info: "#8AD7E3",
        success: "#8AE3B1",
        warning: "#E3B58A",
        danger: "#E38A8A",
        white: "#FFFFFF"
    }, fonts: {
        size: "1rem",
        color: "#FFFFFF"
    }
}

export const Themes = {
    Light,
    Dark
}