export type AlertType = "success" | "info" | "error";

export type Alert = {
  msg: string;
  type: AlertType;
  id: string;
};
