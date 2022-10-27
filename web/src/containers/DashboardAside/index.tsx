import {Nav} from "@/components";
import DashboardNavigationItems from "@/fixtures/dashboard.navigation.json"
import {ROUTES_HOME} from "@/constants";
import Logo from "@/assets/Logo.svg";

export const DashboardAsideContainer = () => {
    return (
        <>
            <Nav.Group>
                <Nav.Logo to={ROUTES_HOME.MAIN} path={Logo}/>
            </Nav.Group>
            <Nav links={DashboardNavigationItems}/>
        </>
    )
}