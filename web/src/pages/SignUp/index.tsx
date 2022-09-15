import {Button, Error, Form, Header, Title, Group, GroupLink, Input, Link} from "@/components";
import {ROUTES_PUBLIC, ROUTES_DASHBOARD} from "@/constants";
import {SubmitHandler, useForm} from "react-hook-form";
import {Main} from "./style";
import {useFetch} from "@/hooks";
import {useNavigate} from "react-router-dom";

interface IFormSignUp {
    name: string
    email: string
    password: string
    confirm_password: string
}

export const SignUp = () => {
    const {register, handleSubmit, formState: {errors}} = useForm<IFormSignUp>();
    const navigate = useNavigate()

    const onSubmit: SubmitHandler<IFormSignUp> = async ({email, password, confirm_password}) => {

        navigate(ROUTES_DASHBOARD.MAIN)
    }

    return (
        <>
            <Main>

                <Header>
                    <Title>Create an account</Title>
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
                    <Group>
                        <Input {...register('confirm_password', {required: 'This Field Is Required'})} type={'password'}
                               id={'confirm_password'} placeholder={"Confirm your password."}
                               autoComplete={"off"}/>
                        <Error>{errors.confirm_password?.message}</Error>
                    </Group>
                    <Button type={"submit"}>Sign Up.</Button>
                    <GroupLink>
                        <Link to={ROUTES_PUBLIC.RECOVERY_PASSWORD}>¿Forgot your password?</Link>
                        <Link to={ROUTES_PUBLIC.SIGN_IN}>¿Already have an account?</Link>
                    </GroupLink>
                </Form>

            </Main>
        </>
    )
}