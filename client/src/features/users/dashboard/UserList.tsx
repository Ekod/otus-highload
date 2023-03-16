import React from "react";
import {User} from "../../../app/models/user";
import {Button, Item, Segment} from "semantic-ui-react";
import {observer} from "mobx-react-lite";
import {useStore} from "../../../app/stores/store";

export default observer(function UsersList() {
    const {userStore} = useStore()
    return (
        <Segment>
            <Item.Group divided>
                {userStore.users.map(user => (
                    <Item key={user.id}>
                        <Item.Content>
                            <Item.Header as="a">{user.firstName} {user.lastName}</Item.Header>
                            <Item.Meta>Age: {user.age}</Item.Meta>
                            <Item.Meta>Gender: {user.gender}</Item.Meta>
                            <Item.Meta>Email: {user.email}</Item.Meta>
                            <Item.Description>
                                <div>City: {user.city}</div>
                                <div>Interests: {user.interests}</div>
                            </Item.Description>
                            <Item.Extra>
                                {userStore.isFriend(user.id) ? (
                                    <>
                                        <Button
                                            floated="right"
                                            content="Удалить из Друзей"
                                            color="red"
                                            onClick={() => userStore.removeFriend(user.id)}
                                        />
                                    </>
                                ) : (
                                    <Button
                                        floated="right"
                                        content="Подружиться"
                                        color="blue"
                                        onClick={() => userStore.becomeFriends(user.id)}
                                    />
                                )}
                            </Item.Extra>
                        </Item.Content>
                    </Item>
                ))}
            </Item.Group>
        </Segment>
    )
})