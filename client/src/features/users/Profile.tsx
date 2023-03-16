import React from "react";
import {observer} from "mobx-react-lite";
import {Grid} from "semantic-ui-react";
import ProfileCard from "./ProfileCard";
import FriendsInfo from "./FriendsInfo";

export default observer(function UserProfile() {
    return (
        <Grid>
            <Grid.Column width="10">
                <ProfileCard/>
            </Grid.Column>
            <Grid.Column width="6">
                <FriendsInfo/>
            </Grid.Column>
        </Grid>
    );
});
