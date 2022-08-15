import React from 'react'
import {createRoot} from 'react-dom/client'
import {App} from "./app";
import {GlobalStyle, Theme} from "@/globalStyle";
import {ThemeProvider} from "styled-components";

createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <ThemeProvider theme={Theme}>
            <GlobalStyle/>
            <App/>
        </ThemeProvider>
    </React.StrictMode>
)