import {Nav, Item, Photo, Group} from "./style"

interface INavProps {
    links: { to: string, label?: string }[]
    children?: | JSX.Element | JSX.Element[] | string | string[];
}

interface IItemProps {
    to: string
    label: string
}

interface IGroupProps {
    children?: | JSX.Element | JSX.Element[] | string | string[];
}

interface ILogoProps {
    to: string
    pathPhoto: string
}

export const Navigation = ({links, children, ...restProps}: INavProps) => {
    return (
        <Nav {...restProps}>
            <section>
                {
                    links.map((item, key) => <Item to={item.to} key={key}>{item.label}</Item>)
                }
            </section>
            {children}
        </Nav>
    )
}

Navigation.Item = ({to, label, ...restProps}: IItemProps) => {
    return (<Item to={to} {...restProps}>{label}</Item>)
}

Navigation.Group = ({children, ...restProps}: IGroupProps) => {
    return (<Group {...restProps}>{children}</Group>)
}

Navigation.Logo = ({to, pathPhoto, ...restProps}: ILogoProps) => {
    return (<Item to={to} {...restProps}><Photo src={pathPhoto}/></Item>)
}