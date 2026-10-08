import { Invitations } from "./invitations";

export class Student {
    id!: string;
    dni!: string;
    name!: string;
    email!: string;
    instance!: string;
    title!: string;
    hasGroup!: boolean;
    alone!: boolean;
    confirmed!: boolean;
    invitations?: Invitations[];
}