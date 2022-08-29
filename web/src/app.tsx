import {BrowserRouter, Routes, Route} from "react-router-dom";
import {About, Dashboard, Home, NotFound, Prices, RecoveryPassword, SignIn, SignUp} from "@/pages";
import {PUBLIC_ROUTES, PROTECTED_ROUTES, HOME_ROUTES} from "@/constants";

export const App = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path={PUBLIC_ROUTES.SIGN_IN} element={<SignIn/>}/>
                <Route path={PUBLIC_ROUTES.SIGN_UP} element={<SignUp/>}/>
                <Route path={PUBLIC_ROUTES.RECOVERY_PASSWORD} element={<RecoveryPassword/>}/>
                <Route path={HOME_ROUTES.HOME} element={<Home/>}>
                    <Route path={HOME_ROUTES.ABOUT} element={<About/>}/>
                    <Route path={HOME_ROUTES.PRICES} element={<Prices/>}/>
                </Route>
                <Route path={PROTECTED_ROUTES.DASHBOARD} element={<Dashboard/>}>

                </Route>
                <Route path={PUBLIC_ROUTES.NOT_FOUND} element={<NotFound/>}/>
            </Routes>
        </BrowserRouter>
    )
}