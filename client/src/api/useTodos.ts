import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { addTodo, fetchTodos, toggleTodo, deleteTodo } from "./client";
import { useAlerts } from "../context/AlertsContext";
import { useEffect } from "react";

export const useGetTodos = () => {
  const { setAlert } = useAlerts();
  const { data, refetch, isError } = useQuery({
    queryKey: ["todos"],
    queryFn: fetchTodos,
  });

  useEffect(() => {
    if (isError) {
      console.error("An error occurred while fetching todos");
      setAlert("error", "An error occurred. Please try again.");
    }
  }, [isError]);

  return { data: data?.data, refetch };
};

export const useAddTodo = () => {
  const { setAlert } = useAlerts();
  const queryClient = useQueryClient();
  const { mutateAsync, isPending } = useMutation({
    mutationFn: addTodo,
    onSuccess: () => {
      setAlert("success", "Todos added");
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error) => {
      setAlert("error", "An error occurred. Please try again.");
      console.error(error);
    },
  });
  return { add: mutateAsync, loading: isPending };
};

export const useToggleTodo = () => {
  const { setAlert } = useAlerts();
  const queryClient = useQueryClient();
  const { mutateAsync } = useMutation({
    mutationFn: toggleTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error) => {
      setAlert("error", "An error occurred. Please try again.");
      console.error(error);
    },
  });
  return { toggle: mutateAsync };
};

export const useDeleteTodo = () => {
  const { setAlert } = useAlerts();
  const queryClient = useQueryClient();
  const { mutateAsync } = useMutation({
    mutationFn: deleteTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error) => {
      setAlert("error", "An error occurred. Please try again.");
      console.error(error);
    },
  });
  return { delete: mutateAsync };
};
