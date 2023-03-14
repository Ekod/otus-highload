export interface User {
    id: number
    firstName: string
    lastName: string
    age: number
    gender: string
    interests: string
    city: string
    email: string
    token: string
}

export interface UserForm {
    email: string
    password: string
}