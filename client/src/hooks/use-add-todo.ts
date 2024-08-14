import { useMutation, useQueryClient } from "@tanstack/react-query";
import { addTodo } from "../api/client";
import { useAlerts } from "../context/AlertsContext";
import { HTTPError } from "../errors";
import { messages } from "../messages";

export const useAddTodo = () => {
  const { setAlert } = useAlerts();
  const queryClient = useQueryClient();
  const { mutateAsync, isPending } = useMutation({
    mutationFn: addTodo,
    onSuccess: () => {
      setAlert("success", messages.SUCCESS_ADDED);
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error: HTTPError) => {
      if (error.status === 503) {
        setAlert("error", messages.TOO_MANY_TODOS);
      } else if (error.status === 400) {
        setAlert("error", messages.PROFANE);
      } else {
        setAlert("error", messages.AGAIN);
      }
    },
  });
  return { add: mutateAsync, loading: isPending };
};
