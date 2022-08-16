import {createGlobalStyle, DefaultTheme} from "styled-components";

export const Light: DefaultTheme = {
    colors: {
        bg: "#FFFFFF"
    }
}

export const Dark: DefaultTheme = {
    colors: {
        bg: "#000000"
    }
}

export const Themes = {
    Light,
    Dark
}

export const GlobalStyle = createGlobalStyle`
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    background-color: ${props => props.theme.colors.bg};
    font-size: 1rem;
    font-family: 'Rubik', sans-serif;
  }
`