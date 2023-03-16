import {ErrorMessage, Form, Formik} from "formik";
import {observer} from "mobx-react-lite";
import {Button, Header} from "semantic-ui-react";
import MyTextInput from "../../app/common/form/MyTextInput";
import {useStore} from "../../app/stores/store";
import * as Yup from 'yup';
import ValidationError from "../errors/ValidationError";
import {Gender} from "../../app/models/user";
import React from "react";

export default observer(function RegisterForm() {
    const {userStore} = useStore();
    return (
        <Formik
            initialValues={{
                firstName: "",
                lastName: "",
                age: 0,
                gender: Gender.male,
                interests: "",
                city: "",
                email: "",
                password: "",
                error: null,
            }}
            onSubmit={(values, {setErrors}) =>
                userStore.register(values).catch(error => setErrors({error: error}))}
            validationSchema={Yup.object({
                firstName: Yup.string().required(),
                lastName: Yup.string().required(),
                email: Yup.string().required(),
                password: Yup.string().required(),
            })}
        >
            {({handleSubmit, isSubmitting, errors, isValid, dirty}) => (
                <Form className='ui form error' onSubmit={handleSubmit} autoComplete='off'>
                    <Header as='h2' content='Sign up to Reactivities' color="teal" textAlign="center"/>
                    <MyTextInput placeholder="First Name" name='firstName'/>
                    <MyTextInput placeholder="Last Name" name='lastName'/>
                    <MyTextInput placeholder="Email" name='email'/>
                    <MyTextInput placeholder="Password" name='password' type='password'/>
                    <MyTextInput placeholder="Age" name='age' type="number"/>
                    <MyTextInput placeholder="Gender" name='gender'/>
                    <MyTextInput placeholder="City" name='city'/>
                    <MyTextInput placeholder="Interests" name='interests'/>
                    <ErrorMessage name='error' render={() =>
                        <ValidationError errors={errors.error}/>}/>
                    <Button
                        disabled={!isValid || !dirty || isSubmitting}
                        loading={isSubmitting}
                        positive content='Register'
                        type="submit" fluid
                    />
                </Form>
            )}

        </Formik>
    )
})