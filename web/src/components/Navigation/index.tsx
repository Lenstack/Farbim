import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"
import {IStyledProps} from "@/interfaces";

export const Navigation = styled.nav<IStyledProps>`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
`

export const Item = styled(ReachRouterLink)<IStyledProps>`
  text-decoration: none;
  color: ${props => props.theme.fonts.color};
  padding: 0.5rem;
  margin-top: 2rem;
  text-align: center;
  font-size: 0.9rem;
`
export const NGroup = styled.section``