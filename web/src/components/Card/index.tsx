import {Container, Header, Body, Footer} from "./style"
import {ReactNode} from "react";

interface ICardProps {
    children?: ReactNode | ReactNode[]
}

export const Card = ({children, ...restProps}: ICardProps) => {
    return (
        <Container {...restProps}>
            {children}
        </Container>
    )
}

Card.Header = ({children, ...restProps}: ICardProps) => {
    return (<Header{...restProps}>{children}</Header>)
}

Card.Body = ({children, ...restProps}: ICardProps) => {
    return (<Body{...restProps}>{children}</Body>)
}

Card.Footer = ({children, ...restProps}: ICardProps) => {
    return (<Footer{...restProps}>{children}</Footer>)
}