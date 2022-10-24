import {SubmitHandler, useForm} from "react-hook-form";
import {Form} from "@/components";
import {Main} from "./style";
import {ROUTES_PUBLIC, ROUTES_DASHBOARD, ROUTES_HOME} from "@/constants";
import {useNavigate} from "react-router-dom";
import {setLocalStorage} from "@/utils";
import {SignInService} from "@/services";

interface IFormSignIn {
    email: string,
    password: string
}

export const SignIn = () => {
    const {register, handleSubmit, formState: {errors}} = useForm<IFormSignIn>();
    const navigate = useNavigate()

    const onSubmit: SubmitHandler<IFormSignIn> = async ({email, password}) => {
        const responseSignInService = await SignInService({email, password})
        const {AccessToken} = responseSignInService
        if (responseSignInService.Errors) {
            return
        }
        setLocalStorage('Token', AccessToken)
        navigate(ROUTES_DASHBOARD.MAIN)
    }

    return (
        <Main>
            <Form>
                <Form.Header>
                    <Form.Title>Sign In</Form.Title>
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
                <Form.Submit type={"submit"}>Sign In.</Form.Submit>
                <Form.Group>
                    <Form.Link to={ROUTES_PUBLIC.RECOVERY_PASSWORD}>¿Forgot your password?</Form.Link>
                    <Form.Link to={ROUTES_PUBLIC.SIGN_UP}>¿You not have an account?</Form.Link>
                </Form.Group>
                <Form.Group>
                    <Form.Link to={ROUTES_HOME.MAIN}>¿You want return to Home?</Form.Link>
                </Form.Group>
            </Form>
        </Main>
    )
};
