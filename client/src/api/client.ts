const BASE_URL = import.meta.env.DEV
  ? "http://localhost:3001/api/v1"
  : "api/v1";

type Method = "GET" | "POST" | "PATCH" | "DELETE";

export const apiClient = (
  method: Method,
  url: string,
  body?: Record<string, string | boolean>
) => {
  return fetch(`${BASE_URL}/${url}`, {
    method,
    headers: {
      "Content-Type": "application/json",
    },
    body: body ? JSON.stringify(body) : undefined,
  }).then((res) => {
    if (!res.ok) throw new Error(res.statusText);
    return res;
  });
};

export const fetchTodos = async () =>
  apiClient("GET", "todos").then((res) => res.json());

export const addTodo = async (title: string) => {
  return apiClient("POST", "todos", { title }).then((res) => res.json());
};

export const toggleTodo = async ({
  id,
  completed,
}: Record<string, string | boolean>) => {
  return apiClient("PATCH", `todos/${id}`, { completed }).then((res) =>
    res.json()
  );
};

export const deleteTodo = async (id: string) =>
  apiClient("DELETE", `todos/${id}`).then((res) => res.json());
