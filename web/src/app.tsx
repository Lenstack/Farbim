import {BrowserRouter, Routes, Route} from "react-router-dom";
import {Home, NotFound, SignIn, SignUp} from "@/pages";
import {ROUTES} from "@/constants";
import {Wrapper} from "@/components";

export const App = () => {
    return (
        <BrowserRouter>
            <Wrapper>
                <Routes>
                    <Route path={ROUTES.SIGN_IN} element={<SignIn/>}/>
                    <Route path={ROUTES.SIGN_UP} element={<SignUp/>}/>
                    <Route path={ROUTES.HOME} element={<Home/>}/>
                    <Route path={ROUTES.NOT_FOUND} element={<NotFound/>}/>
                </Routes>
            </Wrapper>
        </BrowserRouter>
    )
}