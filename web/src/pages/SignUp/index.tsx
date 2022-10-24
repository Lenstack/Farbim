import {Form} from "@/components";
import {ROUTES_PUBLIC, ROUTES_DASHBOARD, ROUTES_HOME} from "@/constants";
import {SubmitHandler, useForm} from "react-hook-form";
import {Main} from "./style";
import {useNavigate} from "react-router-dom";
import {SignUpService} from "@/services";

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
        if (password != confirm_password) return

        const responseSignInService = await SignUpService({email, password})
        const {Code, Message} = responseSignInService

        if (responseSignInService.Errors) return
        navigate(ROUTES_DASHBOARD.MAIN)
    }

    return (
        <Main>
            <Form>
                <Form.Header>
                    <Form.Title>Create an account</Form.Title>
                </Form.Header>
                <Form.Group>
                    <Form.Label>Email</Form.Label>
                    <Form.Input/>
                    <Form.Error>{errors.email?.message}</Form.Error>
                </Form.Group>
                <Form.Group>
                    <Form.Label>Password</Form.Label>
                    <Form.Input/>
                    <Form.Error>{errors.password?.message}</Form.Error>
                </Form.Group>
                <Form.Group>
                    <Form.Label>Confirm Password</Form.Label>
                    <Form.Input/>
                    <Form.Error>{errors.confirm_password?.message}</Form.Error>
                </Form.Group>
                <Form.Submit type={"submit"}>Sign Up.</Form.Submit>
                <Form.Group>
                    <Form.Link to={ROUTES_PUBLIC.RECOVERY_PASSWORD}>¿Forgot your password?</Form.Link>
                    <Form.Link to={ROUTES_PUBLIC.SIGN_IN}>¿Already have an account?</Form.Link>
                </Form.Group>
                <Form.Group>
                    <Form.Link to={ROUTES_HOME.MAIN}>¿You want return to Home?</Form.Link>
                </Form.Group>
            </Form>
        </Main>
    )
}