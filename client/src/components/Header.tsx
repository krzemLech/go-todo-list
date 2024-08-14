import { GoIcon } from "./GoIcon";
import { MongoIcon } from "./MongoIcon";
import { ReactIcon } from "./ReactIcon";

export const Header = () => {
  return (
    <header className="bg-slate-700/50 rounded-md p-4 w-full justify-between items-center flex px-10 mb-10">
      <div className="flex gap-3 items-center max-sm:mx-auto max-sm:translate-x-4">
        <GoIcon />
        <span className="text-2xl text-slate-500 ml-2">+</span>
        <ReactIcon />
        <span className="text-2xl text-slate-500 ml-2">+</span>
        <MongoIcon />
      </div>
      <h1 className="sm:first-line:text-2xl text-slate-500 max-sm:absolute top-1 right-1 text-xl">
        Todos with GO
      </h1>
    </header>
  );
};
