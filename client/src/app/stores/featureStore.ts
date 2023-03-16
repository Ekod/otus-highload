import {makeAutoObservable, reaction} from "mobx";

export enum FeaturesEnum {
    buttons = "buttons",
    searchUsers = "searchUsers",
    listFriends = "listFriends",
    listFriendPosts = "listFriendPosts",
}

export default class FeatureStore {
    currentFeature: FeaturesEnum = FeaturesEnum.buttons

    constructor() {
        makeAutoObservable(this)
    }

    setCurrentFeature = (feature: FeaturesEnum) => {
        this.currentFeature = feature
    }


}