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