import { FC } from "react";
import { AlertType } from "../types";
import { XMarkIcon } from "@heroicons/react/20/solid";
import {
  ExclamationCircleIcon,
  InformationCircleIcon,
  CheckBadgeIcon,
} from "@heroicons/react/24/outline";

type AlertProps = {
  type: AlertType;
  msg: string;
  onClose: () => void;
};

const alertColors = {
  success: "bg-green-500/20",
  info: "bg-blue-500/20",
  error: "bg-red-500/30",
};

const btnColors = {
  success: "hover:text-green-500",
  info: "hover:text-blue-500",
  error: "hover:text-red-500",
};

const textColors = {
  success: "text-green-400",
  info: "text-blue-400",
  error: "text-red-400",
};

export const Alert: FC<AlertProps> = ({ type, msg, onClose }) => {
  return (
    <div className={`p-4 flex ${alertColors[type]}`}>
      {type === "error" && (
        <ExclamationCircleIcon className="size-6 text-red-500 mr-4" />
      )}
      {type === "info" && (
        <InformationCircleIcon className="size-6 text-blue-500 mr-4" />
      )}
      {type === "success" && (
        <CheckBadgeIcon className="size-6 text-green-500 mr-4" />
      )}
      <div className={textColors[type]}>{msg}</div>
      <div className="ml-auto pl-4">
        <button>
          <XMarkIcon
            className={"size-4 text-slate-900 mt-1" + btnColors[type]}
            onClick={onClose}
          />
        </button>
      </div>
    </div>
  );
};
