import {Form} from "@/components";
import {ROUTES_PUBLIC} from "@/constants";
import {SubmitHandler, useForm} from "react-hook-form";
import {Main, Wrapper} from "./style";

interface IFormSendEmail {
    email: string,
}

export const RecoveryPassword = () => {
    const {register, handleSubmit, formState: {errors}} = useForm<IFormSendEmail>();

    const onSubmit: SubmitHandler<IFormSendEmail> = ({email}) => {
        console.log({email})
    }

    return (
        <Wrapper>
            <Main>
                <Form>
                    <Form.Header>
                        <Form.Title>Recovery forgotten password</Form.Title>
                        <Form.SubTitle>Enter the email associated with your account and we’ll send you a recovery
                            link</Form.SubTitle>
                    </Form.Header>
                    <Form.Group>
                        <Form.Input/>
                        <Form.Error>{errors.email?.message}</Form.Error>
                    </Form.Group>
                    <Form.Submit type={"submit"}>Send recovery link.</Form.Submit>
                    <Form.Group>
                        <Form.Link to={ROUTES_PUBLIC.SIGN_IN}>Return to sign in</Form.Link>
                    </Form.Group>
                </Form>
            </Main>
        </Wrapper>
    )
}