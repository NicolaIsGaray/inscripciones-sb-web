export type School = "secundaria" | "deportiva";

export type Applicant = {
  name: string;
  instance: string;
  title: string;
  dni: string;
};

export type IconName =
  | "home" | "users" | "calendar" | "school" | "search"
  | "check" | "arrow" | "trash" | "book" | "mail"
  | "clock" | "plus" | "close" | "shield"
  | "chart" | "edit" | "download" | "upload" | "alert" | "info";

export type Member = { name: string; status?: "pending" | "empty" };
export type Group = { leader: string; solo?: boolean; members: Member[] };