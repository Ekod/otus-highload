import UserStore from "./userStore";
import {createContext, useContext} from "react";

interface Store {
    userStore: UserStore
}

export const store: Store = {
    userStore: new UserStore()
}

export const StoreContext = createContext(store)

export function useStore() {
    return useContext(StoreContext)
}