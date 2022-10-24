import {Container, Header, Group, Title, SubTitle, Label, Input, Link, Submit, Error} from "./style"
import {ReactNode} from "react"

interface IFormProps {
    children?: ReactNode | ReactNode[]
}

export const Form = ({children, ...restProps}: IFormProps) => {
    return (<Container{...restProps}>{children}</Container>)
}

Form.Header = ({children, ...restProps}: IFormProps) => {
    return (<Header{...restProps}>{children}</Header>)
}

Form.Group = ({children, ...restProps}: IFormProps) => {
    return (<Group{...restProps}>{children}</Group>)
}

Form.Title = ({children, ...restProps}: IFormProps) => {
    return (<Title{...restProps}>{children}</Title>)
}

Form.SubTitle = ({children, ...restProps}: IFormProps) => {
    return (<SubTitle{...restProps}>{children}</SubTitle>)
}

Form.Label = ({children, ...restProps}: IFormProps) => {
    return (<Label{...restProps}>{children}</Label>)
}

Form.Input = ({children, ...restProps}: IFormProps) => {
    return (<Input{...restProps}>{children}</Input>)
}

Form.Link = ({children, ...restProps}: any) => {
    return (<Link{...restProps}>{children}</Link>)
}

Form.Submit = ({children, ...restProps}: any) => {
    return (<Submit{...restProps}>{children}</Submit>)
}

Form.Error = ({children, ...restProps}: any) => {
    return (<Error{...restProps}>{children}</Error>)
}
