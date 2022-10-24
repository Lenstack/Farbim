import styled from "styled-components";

export const Main = styled.main`
  width: 100vw;
  height: 100vh;
  background-color: ${props => props.theme.colors.background.primary};
  display: flex;
  justify-content: center;
  align-items: center;  
`