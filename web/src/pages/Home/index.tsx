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
                <Navigation links={HomeNavigation}>
                    <Navigation.Logo to={ROUTES_HOME.MAIN} pathPhoto={Logo}/>
                    <Navigation.Group>
                        <Navigation.Item to={ROUTES_PUBLIC.SIGN_IN} label={"Sign In"}/>
                        <Navigation.Item to={ROUTES_PUBLIC.SIGN_UP} label={"Sign Up"}/>
                    </Navigation.Group>
                </Navigation>
            </Header>
            <Content>
                <Outlet/>
            </Content>
            <Footer>
                <LinkVersion href={ROUTES_PUBLIC.GITHUB} target={"_blank"}>Farm Management v0.0.1</LinkVersion>
            </Footer>
        </Wrapper>
    )
}
