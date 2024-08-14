import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toggleTodo } from "../api/client";
import { useAlerts } from "../context/AlertsContext";
import { messages } from "../messages";

export const useToggleTodo = () => {
  const { setAlert } = useAlerts();
  const queryClient = useQueryClient();
  const { mutateAsync } = useMutation({
    mutationFn: toggleTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error) => {
      setAlert("error", messages.AGAIN);
      console.error(error);
    },
  });
  return { toggle: mutateAsync };
};
