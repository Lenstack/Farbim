import styled from "styled-components";

export const Wrapper = styled.section`
  min-width: 100%;
  min-height: 100vh;
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  background-color: ${props => props.theme.colors.background.primary};
`

export const Aside = styled.aside`
  grid-column: 1/2;
  background-color: ${props => props.theme.colors.background.secondary};

  nav {
    flex-direction: column;
  }
`

export const Content = styled.section`
  grid-column: 2/13;
`