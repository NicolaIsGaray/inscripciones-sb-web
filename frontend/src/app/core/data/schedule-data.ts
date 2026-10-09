import { ScheduleExtra, ScheduleStage } from '../models/schedule.model';

/**
 * Calendario de ingreso a 1.º año 2027 de la Escuela 4-084.
 * Todas las fechas son de 2026 (proceso de admisión, cursada 2027).
 * Fuente: MEMO-2026-275-DGEMZA.
 */
export const scheduleStages: ScheduleStage[] = [
  {
    id: 'primera',
    order: 1,
    title: 'Primera instancia',
    icon: 'users',
    tone: 'blue',
    categories: [
      {
        id: 'discapacidad',
        category: 'Discapacidad',
        priority: 'P1',
        steps: {
          preinscripcion: '9 OCT',
          publicacion: '13 OCT',
          inscripcion: '15 OCT',
        },
      },
      {
        id: 'hermanos',
        category: 'Hermanos/as',
        priority: 'P2',
        steps: {
          preinscripcion: '13 Y 14 OCT',
          publicacion: '14 OCT',
          inscripcion: '15 Y 16 OCT',
        },
      },
      {
        id: 'hijos-personal',
        category: 'Hijos/as del personal',
        priority: 'P1',
        steps: {
          preinscripcion: '19 OCT',
          publicacion: '20 OCT',
          inscripcion: '23 OCT',
        },
      },
      {
        id: 'origen-300m',
        category: 'Origen a 300 m',
        priority: 'P3',
        note: 'Escuela de origen a 300 m, si corresponde',
        steps: {
          preinscripcion: '19 Y 20 OCT',
          publicacion: '21 OCT',
          inscripcion: '23 OCT',
        },
      },
      {
        id: 'federados',
        category: 'Federados (Ed. Física)',
        priority: 'P4',
        note: 'La publicación solo aclara Origen y cercanía',
        steps: {
          preinscripcion: '19 Y 20 OCT',
          publicacion: '21 OCT',
          inscripcion: '23 OCT',
        },
      },
      {
        id: 'cercania',
        category: 'Cercanía del domicilio',
        priority: 'P5',
        steps: {
          preinscripcion: '19 Y 20 OCT',
          publicacion: '21 OCT',
          inscripcion: '23 OCT',
        },
      },
      {
        id: 'abanderados',
        category: 'Abanderados y escoltas titulares',
        priority: 'P1',
        note: 'La publicación es el mismo día, a las 21:00 hs',
        steps: {
          preinscripcion: '26 Y 27 OCT',
          publicacion: '27 OCT',
          inscripcion: '28 Y 29 OCT',
        },
      },
    ],
    banner: {
      date: '2 NOV',
      label: 'Publicación en la escuela del listado final de ingresantes por primera instancia.',
    },
  },
  {
    id: 'segunda',
    order: 2,
    title: 'Segunda instancia',
    icon: 'school',
    tone: 'gold',
    events: [
      { date: '4 AL 8 NOV', label: 'Elección de escuela' },
      { date: '18 NOV', label: 'Consulta de resultados en GEI' },
      { date: '19 AL 20 NOV', label: 'Renuncia al banco' },
      { date: '21 NOV', label: 'Publicación en la web institucional' },
      { date: '22 NOV', label: 'Inscripción en la escuela' },
    ],
  },
  {
    id: 'tercera',
    order: 3,
    title: 'Tercera instancia',
    subtitle: 'Supervisión',
    icon: 'shield',
    tone: 'pink',
    events: [
      {
        date: '24 AL 27 NOV',
        label: 'Confirmación e inscripción en la escuela secundaria.',
      },
      { date: '27 NOV', label: 'Publicación en la web institucional' },
    ],
  },
  {
    id: 'cuarta',
    order: 4,
    title: 'Cuarta instancia',
    subtitle: 'Dirección de línea',
    icon: 'book',
    tone: 'cream',
    events: [
      { date: '30 NOV AL 2 DIC', label: 'Correo y formulario de Dirección de línea.' },
      { date: '2 DIC · 18:00', label: 'Cierre.' },
      {
        date: '4 AL 10 DIC',
        label: 'Asignación e inscripción en la escuela secundaria.',
      },
    ],
  },
];

export const scheduleExtras: ScheduleExtra[] = [
  {
    id: 'datos',
    icon: 'clock',
    title: 'Datos de la escuela',
    items: [
      { label: 'Horarios de atención', value: '8:30 a 12:00 / 14:30 a 17:30' },
      { label: 'Dirección', value: 'French 870 · San Martín · Mendoza' },
      { label: 'Periodo', value: 'Fechas de 2026' },
      { label: 'Normativa', value: 'MEMO-2026-275-DGEMZA' },
    ],
  },
  {
    id: 'documentacion',
    icon: 'book',
    title: 'Documentación',
    items: [
      { label: 'DNI y partida', value: 'Original y copia' },
      { label: 'Prioridad', value: 'Sumá la documentación que acredite tu prioridad' },
      { label: 'Publicaciones', value: 'Se realizan únicamente en la web de la institución' },
    ],
  },
  {
    id: 'cooperadora',
    icon: 'users',
    title: 'Cooperadora',
    items: [
      { label: 'Un estudiante', value: '$35.000' },
      { label: 'Dos hermanos/as', value: '$30.000 por estudiante' },
      { label: 'Tres hermanos/as', value: '$28.000 por estudiante' },
    ],
  },
];