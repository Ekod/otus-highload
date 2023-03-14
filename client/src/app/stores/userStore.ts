import {makeAutoObservable} from "mobx";
import {User, UserForm} from "../models/user";
import agent from "../api/agent";

export default class UserStore {
    users: User[] = []
    user: User | null = null

    constructor() {
        makeAutoObservable(this)
    }

    loadUsers = async () => {
        try {
            const users = await agent.Users.list()
            this.setUsers(users)
        } catch (e) {
            console.log(e)
        }
    }

    get isLoggedIn() {
        return !!this.user
    }

    setUsers = (users: User[]) => {
        this.users = users
    }

    login = async (creds: UserForm) => {
        try {
            const user = await agent.Account.login(creds)
        } catch (e) {
            throw e
        }
    }
}



