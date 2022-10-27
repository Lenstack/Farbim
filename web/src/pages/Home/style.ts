import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Wrapper = styled.section`
  height: 100vh;
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  grid-template-rows: repeat(12, 1fr);
  grid-gap: 1rem;
  background-color: ${props => props.theme.colors.background.primary};
`

export const Header = styled.header`
  grid-column: 4/10;
  grid-row: 1/3;
  display: grid;
  align-items: center;
  
  nav {
    grid-column: 2/4;
  }
  
  nav section .active:last-child {
    background-color: ${props => props.theme.colors.background.secondary};
    color: ${props => props.theme.colors.white};
    border-radius: 0.3rem;
    padding: 0.7rem;
  }
`

export const Content = styled.main`
  grid-column: 3/11;
  grid-row: 3/12;
  margin: 2rem;
`

export const Footer = styled.footer`
  grid-column: 3/11;
  grid-row: 6/6;
  margin-bottom: 3rem;
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