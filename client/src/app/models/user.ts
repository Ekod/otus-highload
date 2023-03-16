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

export interface UserFormLogin {
    email: string
    password: string
}

export enum Gender {
    male = "male",
    female = "female",
}

export interface UserFormRegister {
    firstName: string
    lastName: string
    age: number
    gender: Gender
    interests: string
    city: string
    email: string
    password: string
}