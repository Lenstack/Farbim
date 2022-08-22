import styled from "styled-components";
import {IStyledProps} from "@/interfaces";

export const Wrapper = styled.section`
  min-height: 100vh;
  min-width: 100%;
  background-color: ${props => props.theme.colors.background.primary};
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  grid-template-rows: auto;
  grid-gap: 1rem;
`

export const Main = styled.main<IStyledProps>`
  grid-row: ${props => props.row};
  grid-column: ${props => props.column};
  display: flex;
  justify-content: center;
  align-items: center;
`