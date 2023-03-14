import {createBrowserRouter, RouteObject} from "react-router-dom";
import App from "../layout/App";
import UserDashboard from "../../features/users/dashboard/UserDashboard";
import LoginForm from "../../features/users/LoginForm";

export const routes: RouteObject[] = [
    {
        path: "/",
        element: <App/>,
        children: [
            {
                path: "users",
                element: <UserDashboard/>
            },
            {
                path: "login",
                element: <LoginForm/>
            }
        ]
    }
]

export const router = createBrowserRouter(routes)