import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"
import {IStyledProps} from "@/interfaces";

export const Navigation = styled.nav<IStyledProps>`
  display: flex;
  flex-direction: ${props => props.direction};
`

export const Item = styled(ReachRouterLink)<IStyledProps>`
  text-decoration: none;
  color: ${props => props.theme.fonts.color};
  padding: 1rem;
  margin-right: 0.5rem;
  font-size: ${props => props.theme.fonts.size};
`
export const GroupItem = styled.div``