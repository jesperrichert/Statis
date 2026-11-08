import { ArrowUpRightIcon } from "lucide-react";
import type { Monitor } from "../../../types/status";
import { cn } from "~/lib/utils";

export function StatusDot({ online }: { online: boolean }) {
  return (
    <span className="relative flex h-3 w-3 shrink-0">
      {online && (
        <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-60" />
      )}
      <span
        className={cn(
          "relative inline-flex h-3 w-3 rounded-full",
          online ? "bg-emerald-500" : "bg-red-500",
        )}
      />
    </span>
  );
}

export function MonitorRow({ monitor }: { monitor: Monitor }) {
  const online = true;

  return (
    <div
      className={cn(
        "group flex items-center gap-3 rounded-xl border px-4 py-3 transition-colors",
        "border-white/10 bg-white/3 hover:bg-white/[0.07]",
        !online && "border-red-500/30 bg-red-500/5",
      )}
    >
      <StatusDot online={online} />
      <span className="flex-1 truncate text-sm font-medium text-white">
        {monitor.name}
      </span>
      <span
        className={cn(
          "rounded-full px-2 py-0.5 text-xs font-medium",
          online
            ? "bg-emerald-500/10 text-emerald-400"
            : "bg-red-500/10 text-red-400",
        )}
      >
        {online ? "Online" : "Offline"}
      </span>
      {monitor.external_url && (
        <a
          href={monitor.external_url}
          target="_blank"
          rel="noopener noreferrer"
          aria-label={`${monitor.name} öffnen`}
          className="rounded-lg p-1.5 text-white/40 transition-colors hover:bg-white/10 hover:text-white"
        >
          <ArrowUpRightIcon className="h-4 w-4" />
        </a>
      )}
    </div>
  );
}