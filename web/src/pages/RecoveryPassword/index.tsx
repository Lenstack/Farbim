import {
    Button,
    Error,
    Form,
    Group,
    GroupLink,
    Header,
    Title,
    SubTitle,
    Input,
    Link,

} from "@/components";
import {ROUTES_PUBLIC} from "@/constants";
import {SubmitHandler, useForm} from "react-hook-form";
import {Main} from "./style";

interface IFormSendEmail {
    email: string,
}

export const RecoveryPassword = () => {
    const {register, handleSubmit, formState: {errors}} = useForm<IFormSendEmail>();

    const onSubmit: SubmitHandler<IFormSendEmail> = ({email}) => {
        console.log({email})
    }

    return (
        <>
            <Main>
                <Form onSubmit={handleSubmit(onSubmit)} method={'POST'}>
                    <Header>
                        <Title>Recovery forgotten password</Title>
                        <SubTitle>Enter the email associated with your account and we’ll send you a recovery
                            link</SubTitle>
                    </Header>
                    <Group>
                        <Input {...register('email', {required: 'This Field Is Required'})} type={'email'}
                               id={'email'} placeholder={"Your email."}
                               autoComplete={"on"}/>
                        <Error>{errors.email?.message}</Error>
                    </Group>
                    <Button type={"submit"}>Send recovery link.</Button>
                    <GroupLink>
                        <Link to={ROUTES_PUBLIC.SIGN_IN}>Return to sign in</Link>
                    </GroupLink>
                </Form>
            </Main>
        </>
    )
}