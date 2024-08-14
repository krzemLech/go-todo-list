import { TodoForm } from "./TodoForm";
import { TodoList } from "./TodoList";
import { useGetTodos } from "../hooks";
import { Pagination } from "./Pagination";

export const Todos = () => {
  const { data: todos = [], page, perPage, setPage, total } = useGetTodos();
  return (
    <div className="md:w-[800px] flex flex-col gap-10 mt-10 mx-auto">
      <TodoForm />
      {!!todos.length && <TodoList todos={todos} />}
      {!!todos.length && (
        <Pagination
          page={page}
          setPage={setPage}
          perPage={perPage}
          total={total}
        />
      )}
    </div>
  );
};
