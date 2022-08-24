import {SubmitHandler, useForm} from "react-hook-form";
import {Form, Group, Input, Link, Button, Error, Wrapper, Container, GroupLink} from "@/components";
import {Main, Header, Title} from "./style";
import {PUBLIC_ROUTES} from "@/constants";

interface IFormSignIn {
    email: string,
    password: string
}

export const SignIn = () => {
    const {register, handleSubmit, formState: {errors}} = useForm<IFormSignIn>();

    const onSubmit: SubmitHandler<IFormSignIn> = ({email, password}) => {
        console.log({email, password})
    }

    return (
        <Wrapper>
            <Main>
                <Container>
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
                            <Link to={PUBLIC_ROUTES.RESET_PASSWORD}>¿Forgot your password?</Link>
                            <Link to={PUBLIC_ROUTES.SIGN_UP}>¿You not have an account?</Link>
                        </GroupLink>
                    </Form>
                </Container>
            </Main>
        </Wrapper>
    )
};
