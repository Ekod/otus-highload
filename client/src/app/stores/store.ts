import UserStore from "./userStore";
import {createContext, useContext} from "react";
import ModalStore from "./modalStore";
import FeatureStore from "./featureStore";
import PostStore from "./postStore";

interface Store {
    userStore: UserStore
    modalStore: ModalStore
    featureStore: FeatureStore
    postStore: PostStore
}

export const store: Store = {
    userStore: new UserStore(),
    modalStore: new ModalStore(),
    featureStore: new FeatureStore(),
    postStore: new PostStore()
}

export const StoreContext = createContext(store)

export function useStore() {
    return useContext(StoreContext)
}