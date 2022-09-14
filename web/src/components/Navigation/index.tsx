import {Nav, Item} from "./style"

interface INavProps {
    links: { to: string, label?: string, photo?: string }[]
}

export const Navigation = ({links, ...restProps}: INavProps) => {
    return (
        <Nav {...restProps}>
            {links.map((item, key) => (<Item to={item.to} key={key}>{item.label}</Item>))}
        </Nav>
    )
}