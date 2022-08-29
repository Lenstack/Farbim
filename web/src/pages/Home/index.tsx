import {GroupItem, Item, Navigation, Profile, Wrapper} from "@/components";
import {Header, Content} from "./styles";
import {HOME_ROUTES, NAVIGATION_ITEMS_HOME, PUBLIC_ROUTES} from "@/constants";
import Logo from "@/assets/Logo.svg";
import {Outlet} from "react-router-dom";

export const Home = () => {
    return (
        <Wrapper>
            <Header>
                <Item to={HOME_ROUTES.HOME}>
                    <Profile src={Logo} alt={"Logo"} width={"40"} height={"40"}/>
                </Item>
                <Navigation direction={"row"}>
                    {
                        NAVIGATION_ITEMS_HOME.map((item, index) => (
                            <Item to={item.path} key={index}>{item.name}</Item>))
                    }
                </Navigation>
                <GroupItem>
                    <Item to={PUBLIC_ROUTES.SIGN_IN}>Sign In</Item>
                    <Item to={PUBLIC_ROUTES.SIGN_UP}>Sign Up</Item>
                </GroupItem>
            </Header>
            <Content>
                <Outlet/>
            </Content>
        </Wrapper>
    )
}
