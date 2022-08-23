import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Form = styled.form``

export const Group = styled.section`
  display: block;
  justify-content: center;
  align-items: center;
  margin-top: 0.5rem;
`

export const Label = styled.label`
  display: block;
`

export const Input = styled.input`
  margin: 0.5rem 0 0.5rem 0;
  padding: 0.8rem;
  width: 100%;
  height: 2.2rem;
  border: none;
  border-radius: 0.5rem;
`

export const Button = styled.button``

export const TextArea = styled.textarea``

export const Error = styled.span`
  color: ${props => props.theme.colors.danger};
`

export const Link = styled(ReachRouterLink)``