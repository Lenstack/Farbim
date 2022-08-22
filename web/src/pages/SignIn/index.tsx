import {SubmitHandler, useForm} from "react-hook-form";
import {Form, Label, Group, Input, Link, Button, Error, Heading, Wrapper, Container} from "@/components";
import {Main} from "./style";
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
            <Main column={"5/9"}>
                <Container padding={"3"}>
                    <Heading>
                        <h3>Sign In.</h3>
                    </Heading>
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
                                <Link to={PUBLIC_ROUTES.SIGN_UP}>Not have account?</Link>
                                <Link to={PUBLIC_ROUTES.RESET_PASSWORD}>Forgotten Password?.</Link>
                            </Group>
                        </Group>
                        <Button type={"submit"}>Sign In.</Button>
                    </Form>
                </Container>
            </Main>
        </Wrapper>
    )
};
