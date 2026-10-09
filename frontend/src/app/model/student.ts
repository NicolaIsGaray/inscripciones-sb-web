import { Invitations } from "./invitations";

export class Student {
    id!: string;
    dni!: string;
    name!: string;
    email!: string;
    instance!: string;
    title!: string;
    school!: 'secundaria' | 'deportiva';
    hasGroup!: boolean;
    alone!: boolean;
    confirmed!: boolean;
    invitations?: Invitations[];
}