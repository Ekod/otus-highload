import {observer} from "mobx-react-lite";
import {useStore} from "../../../../app/stores/store";
import {Button, Item, Segment} from "semantic-ui-react";
import React from "react";
import {Post} from "../../../../app/models/post";

export default observer(function FriendsPosts(){
    const {postStore} = useStore()
    return (
        <>
            <Segment>
                <Item.Group divided>
                    {postStore.postList.map(post => {
                        const userID = post[0] as number
                        const userName = post[1] as string
                        const userPosts = post[2] as Post[]
                        return(
                            <Item key={userID}>
                                <Item.Content>
                                    <Item.Header as="a">{userName}</Item.Header>
                                    <Item.Description>
                                        {userPosts.map(post=> <div>{post}</div>)}
                                    </Item.Description>
                                    <Item.Extra>

                                    </Item.Extra>
                                </Item.Content>
                            </Item>
                            )})}
                </Item.Group>
            </Segment>
        </>
    )
})