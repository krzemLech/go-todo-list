import { TodoForm } from "./TodoForm";
import { TodoList } from "./TodoList";
import { useGetTodos } from "../api";

export const Todos = () => {
  const { data: todos = [] } = useGetTodos();
  return (
    <div className="w-[800px] flex flex-col gap-10 mt-10 mx-auto">
      <TodoForm />
      {!!todos.length && <TodoList todos={todos} />}
    </div>
  );
};
