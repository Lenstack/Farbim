import styled from "styled-components";
import {IStyledProps} from "@/interfaces";

export const Wrapper = styled.section`
  min-height: 100vh;
  width: 100%;
  background-color: ${props => props.theme.colors.background.primary};
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  grid-template-rows: auto;
  grid-gap: 0.5rem;
`

export const Grid = styled.section<IStyledProps>`
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  grid-template-rows: auto;
  grid-gap: 1rem;
`

export const Column = styled.section<IStyledProps>``

export const Row = styled.section<IStyledProps>``

export const Container = styled.section<IStyledProps>`
  background-color: ${props => props.theme.colors.background.secondary};
  padding: ${props => props.padding + 'rem'};
  border-radius: 0.5rem;  
`