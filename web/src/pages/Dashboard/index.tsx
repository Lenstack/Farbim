import {Wrapper, Aside, Content} from "./style"
import {Outlet} from "react-router-dom";
import {DashboardAsideContainer} from "@/containers";

export const Dashboard = () => {
    return (
        <Wrapper>
            <Aside>
                <DashboardAsideContainer/>
            </Aside>
            <Content>
                <Outlet/>
            </Content>
        </Wrapper>
    )
}