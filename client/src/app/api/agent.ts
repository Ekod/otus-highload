import axios, {AxiosError, AxiosResponse} from "axios";
import {User, UserFormLogin, UserFormRegister} from "../models/user";
import {toast} from "react-toastify";
import {router} from "../router/Routes";
import {store} from "../stores/store";
import {Post} from "../models/post";

axios.defaults.baseURL = process.env.REACT_APP_BASE_API_URL

axios.interceptors.request.use(config => {
    const token = store.userStore.token;
    if (token && config.headers) config.headers.Authorization = `Bearer ${token}`;
    return config;
})

axios.interceptors.response.use(async response => {
    return response;
}, (error: AxiosError) => {
    const {data, status, config} = error.response as AxiosResponse;
    switch (status) {
        case 400:
            toast.error('bad request')
            break;
        case 401:
            toast.error('unauthorized')
            router.navigate('/');
            break;
        case 403:
            toast.error('forbidden')
            router.navigate('/');
            break;
        case 404:
            router.navigate('/');
            break;
        default:
            router.navigate('/');
            break;
    }

    return Promise.reject(error);
})

const responseBody = <T>(response: AxiosResponse<T>) => response.data

const requests = {
    get: <T>(url: string) => axios.get<T>(url).then(responseBody),
    post: <T>(url: string, body: {}) => axios.post<T>(url, body).then(responseBody),
    del: (url: string) => axios.delete(url).then(responseBody),
}

const Users = {
    list: () => requests.get<User[]>("/users"),
    search: (firstName: string, lastName: string) => requests.get<User[]>(`/search-users?firstName=${firstName}&lastName=${lastName}`),
    getFriends: () => requests.get<User[]>("/friend/list"),
    makeFriends: (friendID: number) => requests.post("/friend/add", {friendID}),
    removeFriend: (friendId: number) => requests.del(`/friend/remove/${friendId}`),
}

const Account = {
    login: (user: UserFormLogin) => requests.post<User>("/user/login", user),
    register: (user: UserFormRegister) => requests.post<User>("/user/register", user),
    current: () => requests.get<User>("/user/info")
}

const Posts = {
    feed: (friendIDs: number[]) => requests.post<Map<number, Array<Post>>>("/post/feed", {friendIDs}),
}

const agent = {
    Users, Account, Posts
}

export default agent