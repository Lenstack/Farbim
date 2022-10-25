import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Nav = styled.nav`
  display: flex;  
  justify-content: space-between;
`

export const Item = styled(ReachRouterLink)`
  text-decoration: none;  
  padding-right: 1rem;
  padding-left: 1rem;
  color: ${props => props.theme.fonts.color};
`

export const Logo = styled.img`
  width: 50px;
  height: 50px;
`

export const Group = styled.div``