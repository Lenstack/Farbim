import {BrowserRouter, Routes, Route} from "react-router-dom";
import {ThemeProvider} from "styled-components";
import {Theme} from "@/globalStyle";
import {NotFound, SignIn} from "@/pages";
import {ROUTES} from "@/constants";

export const App = () => {
    return (
        <ThemeProvider theme={Theme}>
            <BrowserRouter>
                <Routes>
                    <Route path={ROUTES.SIGN_IN} element={<SignIn/>}/>
                    <Route path={ROUTES.NOT_FOUND} element={<NotFound/>}/>
                </Routes>
            </BrowserRouter>
        </ThemeProvider>
    )
}