import React from 'react'
import {createRoot} from 'react-dom/client'
import {App} from "./app";
import {GlobalStyle, Themes} from "@/styles";
import {ThemeProvider} from "styled-components";

createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <ThemeProvider theme={Themes['Dark']}>
            <GlobalStyle/>
            <App/>
        </ThemeProvider>
    </React.StrictMode>
)