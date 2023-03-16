import {observer} from "mobx-react-lite";
import React from "react";
import {Button} from "semantic-ui-react";
import {useStore} from "../../../../app/stores/store";
import {FeaturesEnum} from "../../../../app/stores/featureStore";

export default observer(function FeatureButtons () {
    const {featureStore} = useStore()

    return (
        <>
            <Button
                content="Search Users"
                color="teal"
                onClick={() => featureStore.setCurrentFeature(FeaturesEnum.searchUsers)}
            />
            <Button
                content="List Friends"
                color="blue"
                onClick={() => featureStore.setCurrentFeature(FeaturesEnum.listFriends)}
            />
            <Button
                content="List Friends"
                color="green"
                onClick={() => featureStore.setCurrentFeature(FeaturesEnum.listFriendPosts)}
            />
        </>
    )
})