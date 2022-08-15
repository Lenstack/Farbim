import {ThemeProvider} from "styled-components";
import {Theme} from "@/globalStyle";
import {SignIn} from "@/pages";

export const App = () => {
    return (
        <ThemeProvider theme={Theme}>
            <SignIn/>
        </ThemeProvider>
    )
}