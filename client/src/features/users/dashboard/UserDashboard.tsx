import React, {useEffect} from "react";
import {Container} from "semantic-ui-react";
import {useStore} from "../../../app/stores/store";
import {observer} from "mobx-react-lite";
import UsersList from "./UserList";
import Features from "./features/Features";

function UserDashboard() {
    const {userStore} = useStore()
    useEffect(() => {
        userStore.loadUsers()
        userStore.getFriends();
    }, [userStore])

    return (
        <>
            <Container text>
                <Features />
                {userStore.users.length > 0 ? <UsersList /> : null}
            </Container>

        </>

    )
}

export default observer(UserDashboard)