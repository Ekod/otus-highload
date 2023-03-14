import React from "react";
import {User} from "../../../app/models/user";
import {Item, Segment} from "semantic-ui-react";

interface Props {
    users: User[]
}

export default function ActivityList({users}: Props) {
    return (
        <Segment>
            <Item.Group divided>
                {users.map(user => (
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
                        </Item.Content>
                    </Item>
                ))}
            </Item.Group>
        </Segment>
    )
}