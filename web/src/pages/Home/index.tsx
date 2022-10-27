import {Wrapper, Header, Content} from "./style";
import {Outlet} from "react-router-dom";
import {HomeHeaderContainer} from "@/containers";

export const Home = () => {
    return (
        <Wrapper>
            <Header>
                <HomeHeaderContainer/>
            </Header>
            <Content>
                <Outlet/>
            </Content>
        </Wrapper>
    )
}
