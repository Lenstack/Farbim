import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Container = styled.nav`
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
  width: 110px;
  height: auto;
  filter: ${props => props.theme.logo.color};
`

export const Group = styled.section``