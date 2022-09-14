import {BrowserRouter, Routes, Route} from "react-router-dom";
import {Dashboard, Home, NotFound, RecoveryPassword, SignIn, SignUp} from "@/pages";
import {ROUTES_PUBLIC, ROUTES_DASHBOARD} from "@/constants";

export const App = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path={ROUTES_PUBLIC.SIGN_IN} element={<SignIn/>}/>
                <Route path={ROUTES_PUBLIC.SIGN_UP} element={<SignUp/>}/>
                <Route path={ROUTES_PUBLIC.RECOVERY_PASSWORD} element={<RecoveryPassword/>}/>
                <Route path={ROUTES_PUBLIC.MAIN} element={<Home/>}>
                </Route>
                <Route path={ROUTES_DASHBOARD.MAIN} element={<Dashboard/>}>
                </Route>
                <Route path={ROUTES_PUBLIC.NOT_FOUND} element={<NotFound/>}/>
            </Routes>
        </BrowserRouter>
    )
}