import styled from "styled-components";

interface IProps {
    Column?: string
    Row?: string
    Border?: string
    Padding?: string
    Margin?: string
    Width?: string
    Height?: string
}

export const Wrapper = styled.div<IProps>`
  display: grid;
  grid-template-columns: repeat(11, 1fr);
  grid-template-rows: auto;
  min-height: 100vh;
  min-width: 100%;
  background-color: ${props => props.theme.colors.background.primary};
`

export const Container = styled.section<IProps>`
  background-color: ${props => props.theme.colors.background.secondary};
  border-radius: 0.375rem;
  padding: ${props => props.Padding};
  grid-column: ${props => props.Column};
  grid-row: ${props => props.Row};
  height: ${props => props.Height};
  width: ${props => props.Width};
`

export const Header = styled.header<IProps>`
  grid-column: ${props => props.Column};
  grid-row: ${props => props.Row};
`

export const Aside = styled.aside<IProps>`
  grid-column: ${props => props.Column};
  grid-row: ${props => props.Row};
`

export const Main = styled.main<IProps>`
  grid-column: ${props => props.Column};
  grid-row: ${props => props.Row};
`

export const Footer = styled.footer<IProps>`
  grid-column: ${props => props.Column};
  grid-row: ${props => props.Row};
`

export const Section = styled.section<IProps>`
  grid-column: ${props => props.Column};
  grid-row: ${props => props.Row};
`

export const Article = styled.article<IProps>`
  grid-column: ${props => props.Column};
  grid-row: ${props => props.Row};
`