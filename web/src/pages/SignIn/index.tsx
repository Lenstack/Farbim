import {SubmitHandler, useForm} from "react-hook-form";
import {Form, Group, Header, Title, Input, Link, Button, Error, GroupLink} from "@/components";
import {Main} from "./style";
import {ROUTES_PUBLIC, ROUTES_DASHBOARD} from "@/constants";
import {useNavigate} from "react-router-dom";
import {setLocalStorage} from "@/utils";
import {RefreshTokenService, SignInService} from "@/services";

interface IFormSignIn {
    email: string,
    password: string
}

export const SignIn = () => {
    const {register, handleSubmit, formState: {errors}} = useForm<IFormSignIn>();
    const navigate = useNavigate()

    const onSubmit: SubmitHandler<IFormSignIn> = async ({email, password}) => {
        const responseSignInService = await SignInService({email, password})
        const {TokenAccess} = responseSignInService
        if (responseSignInService.Errors) {
            return
        }

        const responseRefreshToken = await RefreshTokenService(TokenAccess)
        const {TokenRefresh} = responseRefreshToken
        if (responseRefreshToken.Errors) {
            return
        }
        setLocalStorage('token', TokenRefresh)
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
