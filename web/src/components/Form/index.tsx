import styled from "styled-components";
import {Link as ReachRouterLink} from "react-router-dom"
import {IStyledProps} from "@/interfaces";

export const Form = styled.form``

export const Group = styled.section`
  margin-top: 0.8rem;
`

export const GroupLink = styled.section`
  display: flex;
  margin-top: 0.8rem;
  justify-content: space-between;
`

export const Label = styled.label`
  display: block;
  margin: 0.5rem 0 0.5rem 0;
  font-size: 0.8rem;
`

export const Input = styled.input`
  margin: 0.5rem 0 0.5rem 0;
  padding: 0.9rem;
  width: 100%;
  height: 2.5rem;
  border: none;
  border-radius: 0.5rem;
  background-color: ${props => props.theme.colors.background.secondary};
  color: ${props => props.theme.fonts.color};

  ::placeholder {
    color: ${props => props.theme.fonts.color};
  }
`

export const Button = styled.button`
  margin-top: 1.5rem;
  width: 100%;
  height: 3rem;
  border: none;
  border-radius: 0.5rem;
  padding: 0.6rem;
  background-image: ${props => props.theme.colors.gradiant};
  text-align: center;
  text-transform: uppercase;
  color: ${props => props.theme.fonts.color};
  display: block;
  cursor: pointer;
  outline: none;
`

export const TextArea = styled.textarea``


export const Error = styled.span`
  color: ${props => props.theme.colors.danger};
  font-size: 0.8rem;
  display: flex;
  justify-content: flex-end;
  margin-top: 0.5rem;
`

export const Link = styled(ReachRouterLink)<IStyledProps>`
  margin: 0.5rem;
  text-decoration: none;
  color: ${props => props.theme.fonts.color};
  font-size: 0.8rem;

  :last-child {
    color: ${props => props.theme.colors.info};
  }
`