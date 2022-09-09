import styled from "styled-components";

export const Aside = styled.aside`
  background-color: ${props => props.theme.colors.background.secondary};
  grid-column: 1/1;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center; 
`

export const Content = styled.section`
  grid-column: 2/13;
  padding: 1rem;
`