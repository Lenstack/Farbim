import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"

export const Form = styled.form``

export const Group = styled.section``

export const Label = styled.label``

export const Input = styled.input``

export const Button = styled.button``

export const TextArea = styled.textarea``

export const Error = styled.span`
  color: ${props => props.theme.colors.danger};
`

export const Link = styled(ReachRouterLink)``