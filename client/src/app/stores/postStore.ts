import {makeAutoObservable, runInAction} from "mobx";
import {Post} from "../models/post";
import agent from "../api/agent";
import {store} from "./store";


export default class PostStore {
    posts: Map<number, Array<Post>> = new Map<number, Array<Post>>()

    get postList (){
        if (store.userStore.friends.length > 0 && this.posts.size > 0){
            const friendsNames = new Map<number, string>()
            for(const friend of store.userStore.friends){
                friendsNames.set(friend.id, `${friend.firstName} ${friend.lastName}`)
            }

            const friendsPosts = []
            for(const post of this.posts){
                const friendsId = post[0]
                const friendFullName = friendsNames.get(post[0])
                const friendPosts = post[1]
                friendsPosts.push([friendsId, friendFullName, friendPosts])
            }

            return friendsPosts
        } else {
            return []
        }
    }

    constructor() {
        makeAutoObservable(this)
    }

    getPosts = async (friendIDs: number[]) => {
        try {
            const posts = await agent.Posts.feed(friendIDs)
            runInAction(()=>this.posts = posts)
        }catch (e) {
            throw e
        }
    }



}