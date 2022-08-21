import 'styled-components';

declare module 'styled-components' {
    export interface DefaultTheme {
        colors: {
            background: {
                primary: string,
                secondary: string
            },
            info: string,
            success: string,
            warning: string,
            danger: string,
            white: string
        },
        fonts: {
            size: string
            color: string
        }
    }
}