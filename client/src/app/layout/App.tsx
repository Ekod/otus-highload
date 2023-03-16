import React, {useEffect} from 'react';
import {Container} from 'semantic-ui-react';
import NavBar from "./NavBar";
import {Outlet, useLocation} from "react-router-dom";
import {ToastContainer} from "react-toastify";
import HomePage from "../../features/home/HomePage";
import {observer} from "mobx-react-lite";
import {useStore} from "../stores/store";
import ModalContainer from "../common/modal/ModalContainer";

function App() {
    const {userStore} = useStore()
    const location = useLocation()

    useEffect(() => {
        if (userStore.token) {
            userStore.getUser()
        }
    }, [userStore])

    return (
        <>
            <ModalContainer/>
            <ToastContainer position="bottom-right" hideProgressBar theme="colored"/>
            {location.pathname === '/' ? <HomePage/> : (
                <>
                    <NavBar/>
                    <Container style={{marginTop: '7em'}}>
                        <Outlet/>
                    </Container>
                </>
            )}
        </>
    );
}

export default observer(App);
