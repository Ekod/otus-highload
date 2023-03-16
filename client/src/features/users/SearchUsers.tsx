import {observer} from "mobx-react-lite";
import {useStore} from "../../app/stores/store";
import {Form, Formik} from "formik";
import {Button, Header} from "semantic-ui-react";
import MyTextInput from "../../app/common/form/MyTextInput";

export default observer(function SearchUsers() {
    const {userStore} = useStore()
    return (
        <Formik
            initialValues={{firstName: "", lastName: ""}}
            onSubmit={
                ({firstName, lastName}) => userStore.searchUsers(firstName, lastName)
            }>
            {({handleSubmit, isSubmitting}) => (
                <Form
                    className="ui form"
                    onSubmit={handleSubmit}
                    autoComplete="off"
                >
                    <Header as='h2' content='Search Users' color="teal" textAlign="center"/>
                    <MyTextInput placeholder="First Name" name="firstName"/>
                    <MyTextInput placeholder="Last Name" name="lastName"/>
                    <Button loading={isSubmitting} positive content="Search" type="submit" fluid/>
                </Form>
            )}
        </Formik>
    )
})