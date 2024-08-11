import { createContext, useState, useContext, type FC } from "react";
import { AlertType, Alert } from "../types";

type AlertsContextType = {
  alerts: Alert[];
  setAlert: (type: AlertType, msg: string) => void;
  removeAlert: (id: string) => void;
};

export const AlertsContext = createContext<AlertsContextType>({
  alerts: [],
  setAlert: () => {},
  removeAlert: () => {},
});

type ProviderProps = {
  children: React.ReactNode;
};

export const AlertsContextProvider: FC<ProviderProps> = ({ children }) => {
  const [alerts, setAlerts] = useState<Alert[]>([]);

  const setAlert = (type: AlertType, msg: string) => {
    const id = Math.floor(Math.random() * 1000000) + "";
    const newAlert = { msg, type, id };
    setAlerts((prev) => [...prev, newAlert]);
  };

  const removeAlert = (id: string) => {
    setAlerts((prev) => prev.filter((alert) => alert.id !== id));
  };

  return (
    <AlertsContext.Provider value={{ alerts, setAlert, removeAlert }}>
      {children}
    </AlertsContext.Provider>
  );
};

export const useAlerts = () => useContext(AlertsContext);
