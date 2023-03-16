import React from "react";
import {Container, Dropdown, Image, Menu} from "semantic-ui-react";
import {Link, NavLink} from "react-router-dom";
import {useStore} from "../stores/store";
import {observer} from "mobx-react-lite";

export default observer(function NavBar() {
    const {userStore: {user, displayName, logout}} = useStore()
    return (
        <Menu inverted fixed="top">
            <Container>
                <Menu.Item as={NavLink} to="/" header>
                    <img src="/assets/logo.png" alt="logo" style={{marginRight: 10}}/>
                    Social
                </Menu.Item>
                <Menu.Item as={NavLink} to="/users" name="Users"/>
                {
                    user && <Menu.Item position="right">
                        <Image src="/assets/user.png" avatar spaced="right"/>
                        <Dropdown pointing="top left" text={displayName}>
                            <Dropdown.Menu>
                                <Dropdown.Item
                                    as={Link}
                                    to={`profile`}
                                    text="My profile"
                                    icon="user"
                                />
                                <Dropdown.Item onClick={logout} text="Logout" icon="power"/>
                            </Dropdown.Menu>
                        </Dropdown>
                    </Menu.Item>
                }
            </Container>
        </Menu>
    )
})