import {Wrapper, Aside, Content} from "./style"
import DashboardNavigation from "@/fixtures/dashboard.navigation.json"
import {Outlet} from "react-router-dom";
import {Navigation} from "@/components";

export const Dashboard = () => {
    return (
        <Wrapper>
            <Aside>
                <Navigation links={DashboardNavigation}/>
            </Aside>
            <Content>
                <Outlet/>
            </Content>
        </Wrapper>
    )
}