import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"
import {IStyledProps} from "@/interfaces";

export const Navigation = styled.nav<IStyledProps>`
  display: flex;
  flex-direction: column;
  justify-content: center;
  background-color: ${props => props.theme.colors.background.secondary};
  width: 5rem;
`

export const Item = styled(ReachRouterLink)<IStyledProps>`
  text-decoration: none;
  color: ${props => props.theme.fonts.color};
  padding: 0.5rem;
  text-align: center;
  font-size: ${props => props.theme.fonts.size};
`