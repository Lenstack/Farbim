import {Navigation} from "@/components";
import {Wrapper, Header, Content, Footer} from "./style";
import HomeNavigation from "@/fixtures/home.navigation.json"
import {Outlet} from "react-router-dom";

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
               Lorem ipsum dolor sit amet, consectetur adipisicing elit. Libero, sit.
            </Footer>
        </Wrapper>
    )
}
