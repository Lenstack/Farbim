import {BrowserRouter, Routes, Route} from "react-router-dom";
import {Dashboard, Home, NotFound, RecoveryPassword, SignIn, SignUp} from "@/pages";
import {ROUTES_PUBLIC, ROUTES_HOME, ROUTES_DASHBOARD} from "@/constants";
import {HomeFaqContainer, HomeHowWorkContainer, HomePricingContainer} from "@/containers";

export const App = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path={ROUTES_PUBLIC.SIGN_IN} element={<SignIn/>}/>
                <Route path={ROUTES_PUBLIC.SIGN_UP} element={<SignUp/>}/>
                <Route path={ROUTES_PUBLIC.RECOVERY_PASSWORD} element={<RecoveryPassword/>}/>
                <Route path={ROUTES_HOME.MAIN} element={<Home/>}>
                    <Route index element={<HomeHowWorkContainer/>}/>
                    <Route path={ROUTES_HOME.HOW_WORKS} element={<HomeHowWorkContainer/>}/>
                    <Route path={ROUTES_HOME.PRICING} element={<HomePricingContainer/>}/>
                    <Route path={ROUTES_HOME.FAQ} element={<HomeFaqContainer/>}/>
                </Route>
                <Route path={ROUTES_DASHBOARD.MAIN} element={<Dashboard/>}>
                </Route>
                <Route path={ROUTES_PUBLIC.NOT_FOUND} element={<NotFound/>}/>
            </Routes>
        </BrowserRouter>
    )
}