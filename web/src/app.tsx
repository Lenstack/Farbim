import {BrowserRouter, Routes, Route} from "react-router-dom";
import {Dashboard, Faq, Home, HowWorks, NotFound, Pricing, RecoveryPassword, SignIn, SignUp} from "@/pages";
import {ROUTES_PUBLIC, ROUTES_HOME, ROUTES_DASHBOARD} from "@/constants";

export const App = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path={ROUTES_PUBLIC.SIGN_IN} element={<SignIn/>}/>
                <Route path={ROUTES_PUBLIC.SIGN_UP} element={<SignUp/>}/>
                <Route path={ROUTES_PUBLIC.RECOVERY_PASSWORD} element={<RecoveryPassword/>}/>
                <Route path={ROUTES_HOME.MAIN} element={<Home/>}>
                    <Route index element={<HowWorks/>}/>
                    <Route path={ROUTES_HOME.HOW_WORKS} element={<HowWorks/>}/>
                    <Route path={ROUTES_HOME.PRICING} element={<Pricing/>}/>
                    <Route path={ROUTES_HOME.FAQ} element={<Faq/>}/>
                </Route>
                <Route path={ROUTES_DASHBOARD.MAIN} element={<Dashboard/>}>
                </Route>
                <Route path={ROUTES_PUBLIC.NOT_FOUND} element={<NotFound/>}/>
            </Routes>
        </BrowserRouter>
    )
}