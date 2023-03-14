import React, {useEffect} from "react";
import {Grid} from "semantic-ui-react";
import ActivityList from "./UserList";
import {useStore} from "../../../app/stores/store";
import {observer} from "mobx-react-lite";

function UserDashboard() {
    const {userStore} = useStore()
    useEffect(() => {
        userStore.loadUsers()
    }, [userStore])

    return (
        <Grid>
            <Grid.Column width="10">
                <ActivityList users={userStore.users}/>
            </Grid.Column>
        </Grid>
    )
}

export default observer(UserDashboard)