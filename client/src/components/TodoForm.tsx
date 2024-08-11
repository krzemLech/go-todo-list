import { useState, FC } from "react";
import { PlusIcon, ArrowPathIcon } from "@heroicons/react/24/solid";
import { useAddTodo } from "../api";

export const TodoForm: FC = () => {
  const [text, setText] = useState("");

  const { add, loading } = useAddTodo();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    add(text)
      .then(() => setText(""))
      .catch(console.error);
  };

  return (
    <div className="form-wrapper w-full">
      <form className="form flex gap-10 font-light" onSubmit={handleSubmit}>
        <input
          className="block w-full rounded-md border-0 bg-white/5 py-1.5 text-white shadow-sm ring-1 ring-inset ring-white/10 focus:ring-2 focus:ring-inset focus:ring-indigo-500 sm:text-sm sm:leading-6 p-4 outline-none placeholder:font-extralight"
          type="text"
          placeholder="Type your todo here..."
          value={text}
          onChange={(e) => setText(e.target.value)}
          ref={(input) => input?.focus()}
        />
        <button
          type="submit"
          className="rounded-full px-3 py-3 text-sm text-white shadow-sm hover:bg-indigo-600 font-extralight bg-indigo-500 disabled:bg-indigo-800 disabled:text-indigo-700"
          disabled={loading || !text}
        >
          {loading ? (
            <ArrowPathIcon className="size-6 text-slate-300 animate-spin" />
          ) : (
            <PlusIcon className="size-6 text-slate-200" />
          )}
        </button>
      </form>
    </div>
  );
};
