import {Wrapper, Aside, Content} from "./style"
import {Outlet} from "react-router-dom";

export const Dashboard = () => {
    return (
        <Wrapper>
            <Aside>
            </Aside>
            <Content>
                <Outlet/>
            </Content>
        </Wrapper>
    )
}