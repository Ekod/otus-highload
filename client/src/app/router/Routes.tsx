import {createBrowserRouter, RouteObject} from "react-router-dom";
import App from "../layout/App";
import UserDashboard from "../../features/users/dashboard/UserDashboard";

export const routes: RouteObject[] = [
    {
        path: "/",
        element: <App/>,
        children: [
            {
                path: "users",
                element: <UserDashboard/>
            }
        ]
    }
]

export const router = createBrowserRouter(routes)