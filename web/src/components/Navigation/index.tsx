import {Nav, Item, Logo, Group} from "./style"

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

export const Navigation = ({links, children, ...restProps}: INavProps) => {
    return (
        <Nav>
            <Group>
                {
                    links.map((item) =>
                        <Item to={item.to} key={item.id} {...restProps} >
                            {item.label}
                        </Item>
                    )
                }
            </Group>
            {children}
        </Nav>
    )
}

Navigation.Item = ({to, styleClass, children, ...restProps}: IItemProps) => {
    return (<Item to={to} className={styleClass} {...restProps}>{children}</Item>)
}

Navigation.Logo = ({to, path, alt, ...restProps}: ILogoProps) => {
    return (<Item to={to} {...restProps} ><Logo src={path} alt={alt}/></Item>)
}

Navigation.Group = ({children, ...restProps}: IGroupProps) => {
    return (<Group {...restProps}>{children}</Group>)
}