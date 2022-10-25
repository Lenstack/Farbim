import styled from "styled-components";

export const Wrapper = styled.section`
  min-width: 100%;
  min-height: 100vh;
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  background-color: ${props => props.theme.colors.background.primary};
`

export const Aside = styled.aside`
  grid-area: 1 / 1 / 6 / 2;
  background-color: ${props => props.theme.colors.background.third};


  display: flex;

`

export const Content = styled.section`
  grid-area: 1 / 2 / 6 / 13;
`