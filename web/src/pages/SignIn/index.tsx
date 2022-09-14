import {SubmitHandler, useForm} from "react-hook-form";
import {Form, Group, Header, Title, Input, Link, Button, Error, GroupLink} from "@/components";
import {Main} from "./style";
import {ROUTES_PUBLIC, ROUTES_DASHBOARD} from "@/constants";
import {FetchingApi} from "@/services/base";
import {useNavigate} from "react-router-dom";

interface IFormSignIn {
    email: string,
    password: string
}

export const SignIn = () => {
    const {register, handleSubmit, formState: {errors}} = useForm<IFormSignIn>();
    const navigate = useNavigate()

    const signInApi = async (email: string, password: string) => {
        const optionsSignInApi = {
            method: "POST",
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({email: email, password: password})
        }
        return await FetchingApi("authentication/sign_in", optionsSignInApi)
    }

    const refreshTokenApi = async (token: string) => {
        const optionsRefreshTokenApi = {
            method: "POST",
            headers: {
                'Authorization': 'Bearer ' + token,
                'Content-Type': 'application/json'
            }
        }
        return await FetchingApi("authorization/refresh_token", optionsRefreshTokenApi)
    }

    const onSubmit: SubmitHandler<IFormSignIn> = async ({email, password}) => {
        const token = await signInApi(email, password).then(response => {
            return response.TokenAccess
        }).catch(err => console.log(err))

        if (!token) return

        const refreshToken = await refreshTokenApi(token).then(response => {
            return response.TokenRefresh
        }).catch(err => console.log(err))

        window.localStorage.setItem("RefreshToken", refreshToken)
        navigate(ROUTES_DASHBOARD.MAIN)
    }

    return (
        <>
            <Main>
                <Header>
                    <Title>Welcome to management</Title>
                </Header>
                <Form onSubmit={handleSubmit(onSubmit)} method={'POST'}>
                    <Group>
                        <Input {...register('email', {required: 'This Field Is Required'})} type={'email'}
                               id={'email'} placeholder={"Your email."}
                               autoComplete={"on"}/>
                        <Error>{errors.email?.message}</Error>
                    </Group>
                    <Group>
                        <Input {...register('password', {required: 'This Field Is Required'})} type={'password'}
                               id={'password'} placeholder={"Your password."}
                               autoComplete={"off"}/>
                        <Error>{errors.password?.message}</Error>
                    </Group>
                    <Button type={"submit"}>Sign In.</Button>
                    <GroupLink>
                        <Link to={ROUTES_PUBLIC.RECOVERY_PASSWORD}>¿Forgot your password?</Link>
                        <Link to={ROUTES_PUBLIC.SIGN_UP}>¿You not have an account?</Link>
                    </GroupLink>
                </Form>
            </Main>
        </>
    )
};
