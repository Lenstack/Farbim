import styled from "styled-components";

export const Wrapper = styled.section`
  min-height: 100vh;
  min-width: 100%;
  display: grid;
  grid-template-columns: repeat(12, 1fr);
`

export const Header = styled.header`
  grid-column: 2/12;
  display: flex;
  flex-direction: column;

  nav {
    padding: 2rem;
    a:last-child {
      background-color: ${props => props.theme.colors.warning};
      border-radius: 0.3rem;
      width: 120px;
      text-align: center;
      color: ${props => props.theme.colors.white};
    }
  }
`

export const Content = styled.main`
  grid-column: 2/12;
  grid-row: 2/12;
  background-color: aqua;
`

export const Footer = styled.footer`
  grid-column: 2/12;
  background-color: blueviolet;
`