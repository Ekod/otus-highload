import React from "react";
import {Container, Menu} from "semantic-ui-react";
import {NavLink} from "react-router-dom";

export default function NavBar() {
    return (
        <Menu inverted fixed="top">
            <Container>
                <Menu.Item as={NavLink} to="/" header>
                    <img src="/assets/logo.png" alt="logo" style={{marginRight: 10}}/>
                    Social
                </Menu.Item>
                <Menu.Item as={NavLink} to="/users" name="Users"/>
            </Container>
        </Menu>
    )
}