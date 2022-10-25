import {Navigation} from "@/components";
import {Wrapper, Header, Content, Footer, LinkVersion} from "./style";
import {Outlet} from "react-router-dom";
import HomeNavigation from "@/fixtures/home.navigation.json"
import {ROUTES_HOME, ROUTES_PUBLIC} from "@/constants";
import Logo from "@/assets/Logo.svg"

export const Home = () => {
    return (
        <Wrapper>
            <Header>
                <Navigation.Group>
                    <Navigation.Logo to={ROUTES_HOME.MAIN} path={Logo} alt={"logo"}/>
                </Navigation.Group>
                <Navigation links={HomeNavigation}>
                    <Navigation.Group>
                        <Navigation.Item to={ROUTES_PUBLIC.SIGN_IN}>Sign In</Navigation.Item>
                        <Navigation.Item to={ROUTES_PUBLIC.SIGN_UP} styleClass={"active"}>Sign Up</Navigation.Item>
                    </Navigation.Group>
                </Navigation>
            </Header>
            <Content>
                <Outlet/>
            </Content>
            <Footer>
                <LinkVersion to={ROUTES_HOME.MAIN} target={"_blank"}>Management v0.0.1</LinkVersion>
            </Footer>
        </Wrapper>
    )
}
