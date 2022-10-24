import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Nav = styled.nav`
  display: flex;
  align-items: center;
  justify-content: space-evenly;
`

export const Item = styled(ReachRouterLink)`
  text-decoration: none;  
  padding-right: 1rem;
  color: ${props => props.theme.fonts.color};
`

export const Photo = styled.img`
  width: 35px;
  height: 35px;
  object-fit: cover;
`

export const Group = styled.div``