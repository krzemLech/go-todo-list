import { FC } from "react";
import { CheckIcon } from "@heroicons/react/24/solid";
import { XCircleIcon } from "@heroicons/react/24/outline";
import { useToggleTodo, useDeleteTodo } from "../api";

type Todo = {
  _id: string;
  title: string;
  completed: boolean;
};

type Props = {
  todos: Todo[];
};

export const TodoList: FC<Props> = ({ todos }) => {
  const { toggle: onToggle } = useToggleTodo();
  const { delete: onDelete } = useDeleteTodo();

  return (
    <ul
      role="list"
      className="divide-y divide-gray-700 bg-slate-500/10 px-4 py-0 rounded-md"
    >
      {todos.map((todo, idx) => (
        <li key={todo._id} className="flex justify-between gap-x-6 py-5">
          <div className="flex min-w-0 gap-x-4 items-center">
            <div
              className={`h-12 w-12 flex-none rounded-full px-5 py-3 ${
                todo.completed ? "bg-green-500/20" : "bg-slate-500/20"
              }`}
            >
              {idx + 1}
            </div>

            <p className="text-sm font-light text-white">{todo.title}</p>
          </div>
          <div className="hidden shrink-0 sm:flex sm:items-center gap-10">
            <button
              onClick={() =>
                onToggle({ id: todo._id, completed: !todo.completed })
              }
            >
              <CheckIcon
                className={`size-6 ${
                  todo.completed ? "text-green-500" : "text-green-500/30"
                }`}
              />
            </button>
            <button onClick={() => onDelete(todo._id)}>
              <XCircleIcon className="size-7 text-red-500/30 hover:text-red-500" />
            </button>
          </div>
        </li>
      ))}
    </ul>
  );
};
