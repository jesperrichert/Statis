import { ArrowUpRightIcon } from "lucide-react";
import {
  MonitorState,
  MonitorStateStyles,
  type Monitor,
} from "../../../types/status";
import { cn } from "~/lib/utils";

export function StatusDot({ color }: { color: string }) {
  return (
    <span className="relative flex h-3 w-3 shrink-0">
      <span
        className={cn(
          "absolute inline-flex h-full w-full animate-ping rounded-full  opacity-60",
          color,
        )}
      />

      <span
        className={cn("relative inline-flex h-3 w-3 rounded-full", color)}
      />
    </span>
  );
}

export function MonitorRow({ monitor }: { monitor: Monitor }) {
  const state = MonitorStateStyles[monitor.state];

  return (
    <div className="border rounded-xl border-white/10 bg-white/3 hover:bg-white/[0.07] flex flex-col">
      <div
        className={"group flex items-center gap-3 px-4 py-3 transition-colors"}
      >
        <StatusDot color={state.color} />
        <span className="flex-1 truncate text-sm font-medium text-white">
          {monitor.name}
        </span>
        <span
          className={cn(
            "rounded-full px-2 py-0.5 text-xs font-medium",
            state.text,
          )}
        >
          {state.name}
        </span>
        {monitor.external_url && (
          <a
            href={monitor.external_url}
            target="_blank"
            className="rounded-lg p-1.5 text-white/40 transition-colors hover:bg-white/10 hover:text-white"
          >
            <ArrowUpRightIcon className="h-4 w-4" />
          </a>
        )}
      </div>

      {monitor.effected_by.length > 0 ? (
        <span className="text-xs text-zinc-600 ml-5 mb-2">
          Effected by{" "}
          <span className="gap-2 inline-flex">
            {" "}
            {monitor.effected_by.map((e, i) => (
              <>
                <i>
                  {" "}
                  <a href={e.link}>{e.name}</a>{" "}
                  {monitor.effected_by.length - 1 == i ? "" : ", "}
                </i>
              </>
            ))}
          </span>
        </span>
      ) : (
        <></>
      )}
    </div>
  );
}
