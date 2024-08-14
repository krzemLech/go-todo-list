import { useMutation, useQueryClient } from "@tanstack/react-query";
import { deleteTodo } from "../api/client";
import { useAlerts } from "../context/AlertsContext";
import { messages } from "../messages";

export const useDeleteTodo = () => {
  const { setAlert } = useAlerts();
  const queryClient = useQueryClient();
  const { mutateAsync } = useMutation({
    mutationFn: deleteTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error) => {
      setAlert("error", messages.AGAIN);
      console.error(error);
    },
  });
  return { delete: mutateAsync };
};
