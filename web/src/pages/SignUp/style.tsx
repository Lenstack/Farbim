import styled from "styled-components";
import {IStyledProps} from "@/interfaces";
import {Breakpoints} from "@/styles";

export const Main = styled.main<IStyledProps>`
  display: flex;
  justify-content: center;
  align-items: center;
  grid-column: 3/11;

  @media (${Breakpoints.md}) {
    grid-column: 4/10;
  }
`

export const Header = styled.header<IStyledProps>`
  display: flex;
  justify-content: center;
  margin-bottom: 1.5rem;
`

export const Title = styled.h3<IStyledProps>`
  color: ${props => props.theme.fonts.color};
  font-size: 1.2rem;
  font-weight: bold;
`