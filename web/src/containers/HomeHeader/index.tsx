import {Nav} from "@/components";
import {ROUTES_HOME, ROUTES_PUBLIC} from "@/constants";
import Logo from "@/assets/Logo.svg";
import HomeNavigation from "@/fixtures/home.navigation.json";

export const HomeHeaderContainer = () => {
    return (
        <>
            <Nav.Group>
                <Nav.Logo to={ROUTES_HOME.MAIN} path={Logo}/>
            </Nav.Group>
            <Nav links={HomeNavigation}>
                <Nav.Item to={ROUTES_PUBLIC.SIGN_IN}>Sign in</Nav.Item>
                <Nav.Item to={ROUTES_PUBLIC.SIGN_UP} styleClass={"active"}>Get in sign up</Nav.Item>
            </Nav>
        </>
    )
}