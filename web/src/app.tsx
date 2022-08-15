import {BrowserRouter, Routes, Route} from "react-router-dom";
import {ThemeProvider} from "styled-components";
import {Theme} from "@/globalStyle";
import {SignIn} from "@/pages";

export const App = () => {
    return (
        <ThemeProvider theme={Theme}>
            <BrowserRouter>
                <Routes>
                    <Route path={"/"} element={<SignIn/>}/>
                </Routes>
            </BrowserRouter>
        </ThemeProvider>
    )
}