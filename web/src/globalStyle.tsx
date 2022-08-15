import {createGlobalStyle} from "styled-components";

export const Theme = {
    isDark: {
        bgColor: "#0A0911",
        primary: "#101012"
    },
    isLight: {
        bgColor: "#FFFCF6",
        primary: "#FFFFFF"
    }
}

export const GlobalStyle = createGlobalStyle`
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    background-color: #FFFCF6;
    font-size: 0.875rem;
    font-family: 'Rubik', sans-serif;
  }
`