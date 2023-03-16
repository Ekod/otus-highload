import {makeAutoObservable, reaction, runInAction} from "mobx";
import {User, UserFormLogin, UserFormRegister} from "../models/user";
import agent from "../api/agent";
import {router} from "../router/Routes";
import {store} from "./store";

export default class UserStore {
    users: User[] = []
    friends: User[] = []
    user: User | null = null
    token: string | null = localStorage.getItem("jwt")

    constructor() {
        makeAutoObservable(this)

        reaction(
            () => this.token,
            token => {
                if (token) {
                    localStorage.setItem("jwt", token)
                } else {
                    localStorage.removeItem("jwt")
                }
            }
        )
    }

    get displayName() {
        if (this.user) {
            return `${this.user.firstName} ${this.user.lastName}`
        }
    }

    get isLoggedIn() {
        return !!this.user
    }

    loadUsers = async () => {
        try {
            const users = await agent.Users.list()
            runInAction(() => this.users = users)
        } catch (e) {
            throw e
        }
    }

    login = async (creds: UserFormLogin) => {
        try {
            const user = await agent.Account.login(creds)
            runInAction(() => {
                this.user = user
                this.token = user.token
            })
            router.navigate("/users")
            store.modalStore.closeModal()
        } catch (e) {
            throw e
        }
    }

    logout = () => {
        this.token = null
        this.user = null
        router.navigate("/")
    }

    register = async (creds: UserFormRegister) => {
        try {
            const user = await agent.Account.register(creds);
            runInAction(() => {
                this.user = user
                this.token = user.token
            });
            router.navigate('/users');
            store.modalStore.closeModal();
        } catch (error) {
            throw error;
        }
    }

    getUser = async () => {
        try {
            const user = await agent.Account.current()
            runInAction(() => this.user = user);
        } catch (e) {
            throw e
        }
    }

    searchUsers = async (firstName: string, lastName: string) => {
        try {
            const users = await agent.Users.search(firstName, lastName)
            runInAction(() => this.users = users)
        } catch (e) {
            throw e
        }
    }

    getFriends = async () => {
        try {
            const friends = await agent.Users.getFriends()
            runInAction(() => this.friends = friends)
        } catch (e) {
            throw e
        }
    }

    isFriend = (id: number) => {
        return this.friends.find((friend) => friend.id === id) !== undefined

    };

    becomeFriends = async (userID: number) => {
        try {
            await agent.Users.makeFriends(userID);
            await this.getFriends();
        } catch (e) {
            throw e
        }
    };

    removeFriend = async (userId: number) => {
        try {
            await agent.Users.removeFriend(userId);
            await this.getFriends();
        } catch (e) {
            throw e
        }
    };

    nullifyUsers = () => {
        this.users = []
    }
}



