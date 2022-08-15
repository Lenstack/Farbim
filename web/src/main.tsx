import React from 'react'
import {createRoot} from 'react-dom/client'
import {App} from "./app";
import {GlobalStyle, Theme} from "@/globalStyle";

createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <GlobalStyle/>
        <App/>
    </React.StrictMode>
)