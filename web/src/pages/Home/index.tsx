import {Navigation} from "@/components";
import {Wrapper, Header, Content, Footer, LinkVersion} from "./style";
import {Outlet} from "react-router-dom";
import HomeNavigation from "@/fixtures/home.navigation.json"
import {ROUTES_PUBLIC} from "@/constants";

export const Home = () => {
    return (
        <Wrapper>
            <Header>
                <Navigation links={HomeNavigation}/>
            </Header>
            <Content>
                <Outlet/>
            </Content>
            <Footer>
                <LinkVersion to={ROUTES_PUBLIC.GITHUB}>Farm Management v0.0.1</LinkVersion>
            </Footer>
        </Wrapper>
    )
}
