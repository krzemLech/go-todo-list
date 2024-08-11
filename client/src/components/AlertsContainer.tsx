import { useAlerts } from "../context/AlertsContext";
import { Alert } from "./Alert";

export const AlertsContainer = () => {
  const { alerts, removeAlert } = useAlerts();

  return (
    <div className="alerts-container absolute top-10 right-10 flex flex-col gap-4">
      {alerts.map((alert) => (
        <Alert
          key={alert.id}
          {...alert}
          onClose={() => removeAlert(alert.id)}
        />
      ))}
    </div>
  );
};
