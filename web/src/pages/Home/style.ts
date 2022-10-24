import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Wrapper = styled.section`
  height: 100vh;
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  background-color: ${props => props.theme.colors.background.primary};
`

export const Header = styled.header`
  grid-column: 3/11;
  margin: 2rem;

  nav div a:last-child {
    background-color: ${props => props.theme.colors.background.secondary};
    padding: 0.7rem;
    border-radius: 0.3rem;
  }
`

export const Content = styled.main`
  grid-column: 3/11;
  margin: 2rem;
`

export const Footer = styled.footer`
  grid-column: 3/11;
  margin: 2rem;
  display: flex;
  justify-content: end;
  align-items: end;
`

export const LinkVersion = styled(ReachRouterLink)`
  text-decoration: none;
  font-weight: lighter;
  font-size: 12px;
  color: ${props => props.theme.fonts.color};
`