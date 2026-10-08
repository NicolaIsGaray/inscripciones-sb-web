import { Student } from "./student";

export class Group {
    id!: string;
    leader!: Student;
    members?: Student[];
    pending!: Student[];
    full?: boolean;
}