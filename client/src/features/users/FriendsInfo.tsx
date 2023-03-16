import React, {useEffect} from "react";
import {Button, Header, Item, Segment} from "semantic-ui-react";
import {observer} from "mobx-react-lite";
import {useStore} from "../../app/stores/store";

export default observer(function FriendsInfo() {
    const {userStore} = useStore();

    useEffect(() => {
        userStore.getFriends().then(()=> userStore.nullifyUsers());
    }, [userStore]);

    return (
        <>
            <Header content="Друзья" textAlign="center"/>
            {userStore.friends.length! > 0 ? (
                userStore.friends.map((friend) => {
                    return (
                        <Segment.Group key={friend.id}>
                            <Segment>
                                <Item.Group>
                                    <Item>
                                        <Item.Content>
                                            <Item.Header>{`${friend.firstName} ${friend.lastName}`}</Item.Header>
                                            <Item.Meta>{friend.city}</Item.Meta>
                                            <Item.Meta>{friend.gender}</Item.Meta>
                                            <Item.Meta>{friend.age}</Item.Meta>
                                            <Item.Description>
                                                <div>{friend.interests}</div>
                                            </Item.Description>
                                            <Item.Extra>
                                                <Button
                                                    floated="right"
                                                    content="Удалить из Друзей"
                                                    color="red"
                                                    onClick={() => userStore.removeFriend(friend.id)}
                                                />
                                            </Item.Extra>
                                        </Item.Content>
                                    </Item>
                                </Item.Group>
                            </Segment>
                        </Segment.Group>
                    );
                })
            ) : (
                <Segment placeholder>
                    <Header icon>{`У вас друзей нет :(`}</Header>
                </Segment>
            )}
        </>
    );
});
