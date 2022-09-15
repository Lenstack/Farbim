import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Wrapper = styled.section`
  min-height: 100vh;
  min-width: 100%;
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  background-color: ${props => props.theme.colors.background.primary};
`

export const Header = styled.header`
  grid-column: 2/12;
  display: flex;
  flex-direction: column;
  padding: 0 2rem 0 2rem;

  nav {
    padding-top: 2rem;

    a:last-child {
      background-color: ${props => props.theme.colors.dark};
      border-radius: 0.3rem;
      width: 100px;
      text-align: center;
      color: ${props => props.theme.colors.white};
    }
  }
`

export const Content = styled.main`
  grid-column: 2/12;
  grid-row: 2/12;
  padding: 0 2rem 0 2rem;
`

export const Footer = styled.footer`
  grid-column: 2/12;
  display: flex;
  justify-content: end;
  align-items: center;
  padding: 0 2rem 0 2rem;
`

export const LinkVersion = styled(ReachRouterLink)`
  text-decoration: none;
  font-weight: lighter;
  font-size: 12px;
  color: ${props => props.theme.fonts.color};
`