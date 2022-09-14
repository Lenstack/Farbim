import {Form, Item, Navigation, Profile, Wrapper} from "@/components";
import {Aside, Content} from "./style"
import {PROTECTED_ROUTES, NAVIGATION_ITEMS_DASHBOARD, PUBLIC_ROUTES} from "@/constants";
import Logo from "@/assets/Logo.svg"
import {Outlet} from "react-router-dom";

export const Dashboard = () => {
    return (
        <Wrapper>
            <Aside>
                <Item to={PROTECTED_ROUTES.DASHBOARD}>
                    <Profile src={Logo} alt={"Logo"} width={"40"} height={"40"}/>
                </Item>
                <Navigation direction={"column"}>
                    {
                        NAVIGATION_ITEMS_DASHBOARD.map((item, index) => (
                            <Item to={PROTECTED_ROUTES.DASHBOARD + item.path} key={index}>{item.name}</Item>))
                    }
                </Navigation>
                <Item to={PROTECTED_ROUTES.DASHBOARD + PROTECTED_ROUTES.PROFILE}>
                    <Profile src={Logo} alt={"Profile"}/>
                </Item>
                <Form>
                    <Item to={PUBLIC_ROUTES.SIGN_IN}>Logout</Item>
                </Form>
            </Aside>
            <Content>
                <Outlet/>
            </Content>
        </Wrapper>
    )
}