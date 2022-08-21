import {Main, Container, Header, Form, Label, Group, Input, Link, Button, Error} from "@/components";
import {ROUTES} from "@/constants";
import {SubmitHandler, useForm} from "react-hook-form";

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
        <Main Column={"4/9"} Row={"3/5"}>
            <Container Padding={"3rem"} Height={"20rem"}>
                <Header>
                    <h3>Sign In.</h3>
                </Header>
                <Form onSubmit={handleSubmit(onSubmit)} method={'POST'}>
                    <Group>
                        <Label htmlFor={'email'}>Email</Label>
                        <Input {...register('email', {required: 'This Field Is Required'})} type={'email'}
                               id={'email'}/>
                        <Error>{errors.email?.message}</Error>
                    </Group>
                    <Group>
                        <Label htmlFor={'password'}>Password</Label>
                        <Input {...register('password', {required: 'This Field Is Required'})} type={'password'}
                               id={'password'}/>
                        <Error>{errors.password?.message}</Error>
                        <Group>
                            <Link to={ROUTES.SIGN_UP}>Not have account?</Link>
                            <Link to={ROUTES.RESET_PASSWORD}>Forgotten Password?.</Link>
                        </Group>
                    </Group>
                    <Button type={"submit"}>Sign In.</Button>
                </Form>
            </Container>
        </Main>
    )
};
