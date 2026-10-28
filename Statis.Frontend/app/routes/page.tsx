"use client";

import {motion} from "framer-motion";
import {useEffect, useState} from "react";
import {ArrowRight} from "lucide-react";
import type { Monitor } from "~/types/status";
import { MonitorBuilder } from "~/lib/MonitorBuilder";

export default function Page() {
  const [monitors, setMonitors] = useState<Monitor[]>([])

  useEffect(() => {
    async function fetchData() {
      await fetchMonitors()
    }

    fetchData()
  }, []);

  const fetchMonitors = async () => {
    const response = await fetch("/api/status/monitors")
    const data = await response.json()
    setMonitors(data.monitors as Monitor[])
  }

  return (
      <main className={"w-screen"}>
        <div
            className="dark:bg-zinc-900/80 rounded-2xl"
        >
          <motion.h1
              initial={{opacity: 0, y: -20}}
              animate={{opacity: 1, y: 0}}
              transition={{delay: 0.2, duration: 0.6}}
              className="text-4xl font-extrabold mb-8 bg-linear-to-r from-zinc-300 to-zinc-500 bg-clip-text text-transparent"
          >
            Uptime Monitors
          </motion.h1>

          <div className={"flex flex-col w-100"}>
            {
              monitors.map((monitor) => {
                if (!monitor.isParent) return <span key={monitor.id}></span>
                return (
                    <div className={""}>
                      {new MonitorBuilder(monitor, monitors.filter((m) => m.id != monitor.id)).toMonitor()}
                    </div>
                )
              })
            }
          </div>
          <div
              onClick={() => window.open("/status/events")}
              className={"border-2 p-2 rounded-2xl mt-5 inline-flex w-full cursor-pointer hover:text-zinc-400"}
          >
                    <span>
                        Browse Events
                    </span>
            <span className={"left-62 relative"}>
                        <ArrowRight/>
                    </span>
          </div>
        </div>
      </main>
  );
}
