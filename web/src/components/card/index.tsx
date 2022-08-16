import {Container} from "./style";
import {ReactNode} from "react";

interface ICardProps {
    children?: ReactNode
}

export const Card = ({children}: ICardProps) => {
    return (
        <Container>{children}</Container>
    )
}