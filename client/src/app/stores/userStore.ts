import {makeAutoObservable} from "mobx";
import {User} from "../models/user";
import agent from "../api/agent";

export default class UserStore {
    users: User[] = []

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


    setUsers = (users: User[]) => {
        this.users = users
    }
}