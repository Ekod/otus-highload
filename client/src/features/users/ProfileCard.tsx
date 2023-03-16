import React from "react";
import {observer} from "mobx-react-lite";
import {Card, Image} from "semantic-ui-react";
import {useStore} from "../../app/stores/store";

export default observer(function ProfileCard() {
    const {userStore: {user}} = useStore();
    return (
        <Card>
            <Image src="/assets/user.png" wrapped ui={false}/>
            <Card.Content>
                <Card.Header>
                    {user?.firstName} {user?.lastName}
                </Card.Header>
                <Card.Meta>
                    <div>{`Пол: ${user?.gender}`}</div>
                    <div>{`Город: ${user?.city}`}</div>
                    <div>{`Возраст: ${user?.age}`}</div>
                </Card.Meta>
                <Card.Description>{`Интересы: ${user?.interests}`}</Card.Description>
            </Card.Content>
        </Card>
    );
});
