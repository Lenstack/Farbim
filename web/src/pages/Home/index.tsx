import {Wrapper} from "@/components";
import {Link} from "react-router-dom";
import {PUBLIC_ROUTES} from "@/constants";

export const Home = () => {
    return (
        <Wrapper>
            Home
            <Link to={PUBLIC_ROUTES.SIGN_IN}/>
        </Wrapper>
    )
}
