import {Article, Aside, Footer, Header, Main, Wrapper} from "./style";
import {ReactNode} from "react";

interface IProps {
    children?: ReactNode,
}

export const Layout = ({children}: IProps) => {
    return (
        <Wrapper>{children}</Wrapper>
    )
}

Layout.Aside = ({children}: IProps) => {
    return (
        <Aside>{children}</Aside>
    )
}

Layout.Header = ({children}: IProps) => {
    return (
        <Header>{children}</Header>
    )
}

Layout.Main = ({children}: IProps) => {
    return (
        <Main>{children}</Main>
    )
}

Layout.Article = ({children}: IProps) => {
    return (
        <Article>{children}</Article>
    )
}

Layout.Footer = ({children}: IProps) => {
    return (
        <Footer>{children}</Footer>
    )
}