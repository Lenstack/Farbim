import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Container = styled.form`
  display: flex;
  flex-direction: column;
`

export const Header = styled.header`
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
`

export const Group = styled.section`
  padding: 1rem 0 1rem 0;
`

export const Title = styled.h1`
  font-size: 1.2rem;
`

export const SubTitle = styled.h2`
  font-size: 0.8rem;
  font-weight: lighter;
`

export const Label = styled.label``

export const Input = styled.input`
  background-color: ${props => props.theme.colors.background.third};
  color: ${props => props.theme.fonts.color};
  border: none;
  border-radius: 0.3rem;
  width: 100%;
  height: 2.2rem;
  padding: 0.5rem;
  margin-top: 0.5rem;
`

export const Link = styled(ReachRouterLink)`
  text-decoration: none;
  color: ${props => props.theme.fonts.color};
`

export const Submit = styled.button`
  border: none;
  background-color: ${props => props.theme.colors.background.secondary};
  color: ${props => props.theme.colors.white};  
  border-radius: 0.3rem;
  width: 100%;
  height: 2.5rem;
`

export const Error = styled.span`
  color: ${props => props.theme.colors.danger};
`