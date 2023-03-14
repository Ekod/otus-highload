import axios, {AxiosError, AxiosResponse} from "axios";
import {User} from "../models/user";
import {toast} from "react-toastify";
import {router} from "../router/Routes";

axios.defaults.baseURL = "http://localhost:5000/api"

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
        case 403:
            toast.error('forbidden')
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

const agent = {
    Users
}

export default agent