import {Button, Container, Error, Form, Header, Title, Group, GroupLink, Input, Link, Wrapper} from "@/components";
import {PUBLIC_ROUTES} from "@/constants";
import {SubmitHandler, useForm} from "react-hook-form";
import {Main} from "./style";

interface IFormSignUp {
    name: string
    email: string
    password: string
    confirm_password: string
}

export const SignUp = () => {

    const {register, handleSubmit, formState: {errors}} = useForm<IFormSignUp>();

    const onSubmit: SubmitHandler<IFormSignUp> = ({name, email, password, confirm_password}) => {
        console.log({name, email, password, confirm_password})
    }

    return (<Wrapper>
        <Main>
            <Container>
                <Header>
                    <Title>Create an account</Title>
                </Header>
                <Form onSubmit={handleSubmit(onSubmit)} method={'POST'}>
                    <Group>
                        <Input {...register('name', {required: 'This Field Is Required'})} type={'text'}
                               id={'name'} placeholder={"Your full name."}
                               autoComplete={"on"}/>
                        <Error>{errors.name?.message}</Error>
                    </Group>
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
                        <Link to={PUBLIC_ROUTES.RECOVERY_PASSWORD}>¿Forgot your password?</Link>
                        <Link to={PUBLIC_ROUTES.SIGN_IN}>¿Already have an account?</Link>
                    </GroupLink>
                </Form>
            </Container>
        </Main>
    </Wrapper>)
}