import { Applicant, School } from '../models/school.model';

export const applicantsData: Record<School, Applicant[]> = {
  secundaria: [
    { name: "Agostina Lucero", instance: "Primera instancia", title: "CUD", dni: "48123890" },
    { name: "Benjamín Sosa", instance: "Primera instancia", title: "Hermanos", dni: "47905612" },
    { name: "Catalina Ríos", instance: "Primera instancia", title: "Hijo de Personal", dni: "48221467" },
    { name: "Dante Quiroga", instance: "Primera instancia", title: "Escuela cercana a 300m", dni: "47887654" },
    { name: "Emma Funes", instance: "Primera instancia", title: "Cercanía a Domicilio", dni: "48310987" },
    { name: "Felipe Correa", instance: "Primera instancia", title: "Abanderado/Escolta", dni: "48001234" },
    { name: "Guadalupe Molina", instance: "Segunda instancia", title: "Foro Virtual GEI", dni: "47994421" },
    { name: "Hilario Vega", instance: "Tercera instancia", title: "Supervisión", dni: "48123456" },
  ],
  deportiva: [
    { name: "Ignacio Pereyra", instance: "Primera instancia", title: "Hermanos", dni: "48098765" },
    { name: "Josefina Castro", instance: "Primera instancia", title: "Hijos de Personal", dni: "48222678" },
    { name: "Lautaro Méndez", instance: "Primera instancia", title: "Escuela a 300m", dni: "48111222" },
    { name: "Malena Bustos", instance: "Primera instancia", title: "Cercanía a Domicilio", dni: "47987654" },
  ],
};