import {Item, Navigation, Wrapper} from "@/components";
import {PROTECTED_ROUTES} from "@/constants";

export const Dashboard = () => {
    return (
        <Wrapper>
            <Navigation>
                <Item to={PROTECTED_ROUTES.DASHBOARD}>Lorem ipsum dolor.</Item>
                <Item to={PROTECTED_ROUTES.DASHBOARD}>Lorem ipsum dolor.</Item>
                <Item to={PROTECTED_ROUTES.DASHBOARD}>Lorem ipsum dolor.</Item>
                <Item to={PROTECTED_ROUTES.DASHBOARD}>Lorem ipsum dolor.</Item>
                <Item to={PROTECTED_ROUTES.DASHBOARD}>Lorem ipsum dolor.</Item>
            </Navigation>
        </Wrapper>
    )
}