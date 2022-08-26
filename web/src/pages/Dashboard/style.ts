import styled from "styled-components";

export const Aside = styled.aside`
  background-color: ${props => props.theme.colors.background.secondary};
  grid-column: 1/1;
  display: flex;
  align-items: center;
  justify-content: center;
`

export const Content = styled.section`
  grid-column: 3/12;
`