"use client";

import type { Monitor } from "~/types/status";
import { Badge } from "../ui/badge";
import { ArrowUpRightIcon } from "lucide-react";

export function StatusMonitor({ monitor }: { monitor: Monitor }) {
  return (
    <div className={"border-2 p-2 rounded-2xl ml-8"} key={monitor.id}>
      <span className={"inline-flex justify-center items-center"}>
        <div className={"text-sm flex justify-center items-center"}>
          <Badge color={monitor.isOnline ? "green" : "red"}>
            {monitor.isOnline ? (
              <img src={"/online.webp"} width={20} height={20} alt={"Online"} />
            ) : (
              <img
                src={"/offline.webp"}
                width={20}
                height={20}
                alt={"Offline"}
              />
            )}
          </Badge>
          <span>{monitor.name}</span>
          <span className={"inline-flex"}></span>
        </div>
        <span
          className={"p-2 flex justify-center items-center relative left-30"}
        >
          <Badge
            className={"flex justify-center items-center bg-white"}
            variant={"link"}
            asChild
          >
            <a
              className={"hover:text-blue-400"}
              href={monitor.external_url ?? "#"}
            >
              Open Link <ArrowUpRightIcon data-icon="inline-end" />
            </a>
          </Badge>
        </span>
      </span>
    </div>
  );
}
