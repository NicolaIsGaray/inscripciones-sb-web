import { IconName } from './school.model';

export type StageTone = "blue" | "gold" | "pink" | "cream";

/** Las tres columnas de la matriz de la primera instancia */
export type ScheduleColumns = {
  preinscripcion: string;
  publicacion: string;
  inscripcion: string;
};

/** Una categoría de la primera instancia, con sus tres fechas */
export type ScheduleCategory = {
  id: string;
  category: string;
  priority?: string;
  note?: string;
  steps: ScheduleColumns;
};

/** Una fecha suelta, usada en las instancias 2, 3 y 4 */
export type ScheduleEvent = {
  date: string;
  label: string;
  note?: string;
};

export type ScheduleStage = {
  id: string;
  order: number;
  title: string;
  subtitle?: string;
  icon: IconName;
  tone: StageTone;
  /** Si viene `categories`, la etapa se dibuja como matriz de 3 columnas */
  categories?: ScheduleCategory[];
  /** Si viene `events`, la etapa se dibuja como lista de fechas */
  events?: ScheduleEvent[];
  /** Fila destacada al pie de una etapa */
  banner?: ScheduleEvent;
};

/** Dato suelto de la esquina: horarios, documentación, cooperadora */
export type ScheduleExtra = {
  id: string;
  icon: IconName;
  title: string;
  items: { label: string; value: string }[];
};