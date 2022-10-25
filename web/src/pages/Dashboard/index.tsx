import {Wrapper, Aside, Content} from "./style"
import {Outlet} from "react-router-dom";
import {Navigation} from "@/components";
import DashboardNavigationItems from "@/fixtures/dashboard.navigation.json"

export const Dashboard = () => {
    return (
        <Wrapper>
            <Aside>
                <Navigation links={DashboardNavigationItems}/>
            </Aside>
            <Content>
                <Outlet/>
            </Content>
        </Wrapper>
    )
}