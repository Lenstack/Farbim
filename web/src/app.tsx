import {BrowserRouter, Routes, Route} from "react-router-dom";
import {NotFound, SignIn} from "@/pages";
import {ROUTES} from "@/constants";

export const App = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path={ROUTES.SIGN_IN} element={<SignIn/>}/>
                <Route path={ROUTES.NOT_FOUND} element={<NotFound/>}/>
            </Routes>
        </BrowserRouter>
    )
}