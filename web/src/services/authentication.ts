import {BaseUrl} from "@/services";
import axios from "axios";
import {User} from "@/models";

const signInUrl = BaseUrl + "authentication/sign_in"
const signUpUrl = BaseUrl + "authentication/sign_up"
const refreshTokenUrl = BaseUrl + "authorization/refresh_token"

export const SignInService = ({email, password}: User) => {
    return axios.post(signInUrl, {"Email": email,"Password": password})
        .then(response => response.data)
        .catch(error => error.response.data)
}

export const SignUpService = ({email, password}: User) => {
    return axios.post(signUpUrl, {"Email": email,"Password": password})
        .then(response => response.data)
        .catch(error => error.response.data)
}

export const RefreshTokenService = (refreshToken: string) => {
    return axios.post(refreshTokenUrl, {}, {headers: {'Authorization': `Bearer ${refreshToken}`}})
        .then(response => response.data)
        .catch(error => error.response.data)
}