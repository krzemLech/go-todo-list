import { useQuery } from "@tanstack/react-query";
import { fetchTodos } from "../api/client";
import { useAlerts } from "../context/AlertsContext";
import { useEffect, useState } from "react";
import { messages } from "../messages";

export const useGetTodos = () => {
  const { setAlert } = useAlerts();
  const [page, setPage] = useState(1);
  const [perPage, setPerPage] = useState(5);
  const { data, refetch, isError } = useQuery({
    queryKey: ["todos", page, perPage],
    queryFn: fetchTodos,
  });

  useEffect(() => {
    if (isError) {
      console.error("An error occurred while fetching todos");
      setAlert("error", messages.SERVER);
    }
  }, [isError]);

  return {
    data: data?.data || [],
    refetch,
    page,
    setPage,
    perPage,
    setPerPage,
    total: data?.total,
  };
};
