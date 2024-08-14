import { FC } from "react";

type PaginationProps = {
  page: number;
  perPage: number;
  setPage: (page: number) => void;
  total: number;
};

export const Pagination: FC<PaginationProps> = ({
  page,
  setPage,
  perPage,
  total,
}) => {
  const buttonClass =
    "relative inline-flex ml-3 items-center rounded-md bg-transparent px-3 py-2 text-sm font-thin text-slate-400 ring-1 ring-inset ring-slate-400 hover:bg-slate-700 focus-visible:outline-offset-0 disabled:ring-slate-700 disabled:text-slate-700 disabled:cursor-not-allowed disabled:hover:bg-transparent";
  return (
    <nav
      aria-label="Pagination"
      className="flex items-center justify-between rounded-md bg-slate-500/10 px-4 py-3 sm:px-6"
    >
      <div className="hidden sm:block">
        <p className="text-sm text-slate-500">
          Showing{" "}
          <span className="font-medium">{(page - 1) * perPage + 1}</span> to{" "}
          <span className="font-medium">
            {page * perPage > total ? total : page * perPage}
          </span>{" "}
          of <span className="font-medium">{total}</span> results
        </p>
      </div>
      <div className="flex flex-1 justify-between sm:justify-end">
        <button
          disabled={page === 1}
          onClick={() => setPage(page - 1)}
          className={buttonClass}
        >
          Previous
        </button>
        <button
          disabled={page * perPage >= total}
          onClick={() => setPage(page + 1)}
          className={buttonClass}
        >
          Next
        </button>
      </div>
    </nav>
  );
};
