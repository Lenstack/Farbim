import styled from "styled-components";

export const Wrapper = styled.div`
  width: 100vw;
  height: 100vh;
  background-color: ${props => props.theme.colors.background.primary};
  display: grid;
  grid-template-columns: repeat(12, 1fr);
`

export const Main = styled.main`
  grid-column: 2/12;
  display: flex;
  justify-content: center;
  align-items: center;
`