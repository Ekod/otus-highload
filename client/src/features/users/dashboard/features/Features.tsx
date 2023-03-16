import {observer} from "mobx-react-lite";
import React from "react";
import {useStore} from "../../../../app/stores/store";
import SearchUsers from "../../SearchUsers";
import {FeaturesEnum} from "../../../../app/stores/featureStore";
import FeatureButtons from "./FeatureButtons";
import FriendsInfo from "../../FriendsInfo";
import FriendsPosts from "./FriendsPosts";

export default observer(function Features () {
    const {featureStore} = useStore()

    const renderFeature = () => {
        switch (featureStore.currentFeature){
            case FeaturesEnum.searchUsers:
                return <SearchUsers />
            case FeaturesEnum.listFriends:
                return <FriendsInfo />
            case FeaturesEnum.listFriendPosts:
                return <FriendsPosts />
            default:
                return <FeatureButtons />
        }
    }
    return (
        <>
            {renderFeature()}
        </>
    )
})