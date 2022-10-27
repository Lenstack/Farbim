import {Container, Item, Logo, Group} from "./style"

interface INavProps {
    links: { id: string, to: string, label?: string, logo?: string }[]
    children?: | JSX.Element | JSX.Element[] | string | string[];
}

interface IGroupProps {
    children?: | JSX.Element | JSX.Element[] | string | string[];
}

interface IItemProps {
    to: string
    styleClass?: string
    children?: | JSX.Element | JSX.Element[] | string | string[];
}

interface ILogoProps {
    to: string
    path: string
    alt?: string
}

export const Nav = ({links, children, ...restProps}: INavProps) => {
    return (
        <Container>
            <Group>
                {
                    links.map((item) =>
                        <Item to={item.to} key={item.id} {...restProps} >
                            {item.label}
                        </Item>
                    )
                }
            </Group>
            {
                children ? <Group>{children}</Group> : null
            }
        </Container>
    )
}

Nav.Item = ({to, styleClass, children, ...restProps}: IItemProps) => {
    return (<Item to={to} className={styleClass} {...restProps}>{children}</Item>)
}

Nav.Logo = ({to, path, alt, ...restProps}: ILogoProps) => {
    return (<Item to={to} {...restProps} ><Logo src={path} alt={alt}/></Item>)
}

Nav.Group = ({children, ...restProps}: IGroupProps) => {
    return (<Group {...restProps}>{children}</Group>)
}