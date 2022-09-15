import {Nav, Item, Photo} from "./style"

interface INavProps {
    links: { to: string, label?: string, photo?: string }[]
}

export const Navigation = ({links, ...restProps}: INavProps) => {
    return (
        <Nav {...restProps}>
            {
                links.map((item, key) => item.photo ?
                    <Item to={item.to} key={key}><Photo src={item.photo}/></Item> :
                    <Item to={item.to} key={key}>{item.label}</Item>
                )
            }
        </Nav>
    )
}