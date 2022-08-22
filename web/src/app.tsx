import {BrowserRouter, Routes, Route} from "react-router-dom";
import {Home, NotFound, SignIn, SignUp} from "@/pages";
import {PUBLIC_ROUTES} from "@/constants";

export const App = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path={PUBLIC_ROUTES.SIGN_IN} element={<SignIn/>}/>
                <Route path={PUBLIC_ROUTES.SIGN_UP} element={<SignUp/>}/>
                <Route path={PUBLIC_ROUTES.HOME} element={<Home/>}/>
                <Route path={PUBLIC_ROUTES.NOT_FOUND} element={<NotFound/>}/>
            </Routes>
        </BrowserRouter>
    )
}