import { HTTPError } from "../errors";

const BASE_URL = import.meta.env.DEV
  ? "http://localhost:3001/api/v1"
  : "api/v1";

type Method = "GET" | "POST" | "PATCH" | "DELETE";

type Params = {
  query?: Record<string, string | number>;
  params?: Record<string, string | number>;
  body?: Record<string, string | boolean>;
};

type useQueryParams = {
  queryKey: (string | number)[];
};

export const apiClient = (method: Method, url: string, params: Params) => {
  const urlObj = new URL(`${BASE_URL}/${url}`);
  if (params.query) {
    Object.entries(params.query).forEach(([key, value]) => {
      urlObj.searchParams.append(key, value.toString());
    });
  }
  return fetch(urlObj.href, {
    method,
    headers: {
      "Content-Type": "application/json",
    },
    body: params.body ? JSON.stringify(params.body) : undefined,
  }).then((res) => {
    if (!res.ok) {
      throw new HTTPError(res.status, res.statusText);
    }
    return res;
  });
};

export const fetchTodos = async ({
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  queryKey: [_, page, perPage],
}: useQueryParams) => {
  return apiClient("GET", "todos", { query: { page, perPage } }).then((res) =>
    res.json()
  );
};

export const addTodo = async (title: string) => {
  return apiClient("POST", "todos", { body: { title } }).then((res) =>
    res.json()
  );
};

export const toggleTodo = async ({
  id,
  completed,
}: Record<string, string | boolean>) => {
  return apiClient("PATCH", `todos/${id}`, { body: { completed } }).then(
    (res) => res.json()
  );
};

export const deleteTodo = async (id: string) =>
  apiClient("DELETE", `todos/${id}`, {}).then((res) => res.json());
