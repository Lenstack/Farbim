import {createGlobalStyle} from "styled-components";

export const GlobalStyle = createGlobalStyle`
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    font-size: ${props => props.theme.fonts.size};
    font-family: 'Rubik', sans-serif;
    color: ${props => props.theme.fonts.color};
  }  
`