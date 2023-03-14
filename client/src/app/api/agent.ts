import axios, {AxiosError, AxiosResponse} from "axios";
import {User, UserForm} from "../models/user";
import {toast} from "react-toastify";
import {router} from "../router/Routes";

axios.defaults.baseURL = process.env.REACT_APP_BASE_API_URL

// axios.interceptors.request.use(config => {
//     const token = store.commonStore.token;
//     if (token && config.headers) config.headers.Authorization = `Bearer ${token}`;
//     return config;
// })

axios.interceptors.response.use(async response => {
    return response;
}, (error: AxiosError) => {
    const { data, status, config } = error.response as AxiosResponse;
    switch (status) {
        case 400:
            toast.error('bad request')
            break;
        case 403:
            toast.error('forbidden')
            router.navigate('/login');
            break;
        case 404:
            router.navigate('/not-found');
            break;
    }
    return Promise.reject(error);
})

const responseBody = <T>(response: AxiosResponse<T>) => response.data

const requests = {
    get: <T>(url: string) => axios.get<T>(url).then(responseBody),
    post: <T>(url: string, body: {}) => axios.post<T>(url, body).then(responseBody),
}

const Users = {
    list: () => requests.get<User[]>("/users")
}

const Account = {
    login: (user: UserForm) => requests.post<User>("/login", user),
    register: (user: UserForm) => requests.post<User>("/register", user),
}

const agent = {
    Users, Account
}

export default agent