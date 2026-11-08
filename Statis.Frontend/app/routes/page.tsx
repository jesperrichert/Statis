"use client";

import { motion } from "framer-motion";
import { useEffect, useState } from "react";
import { ArrowRight } from "lucide-react";
import { MonitorBuilder } from "~/lib/MonitorBuilder";
import type { Monitor } from "../../types/status";

export default function Page() {
  const [monitors, setMonitors] = useState<Monitor[]>([]);

  useEffect(() => {
    async function fetchData() {
      await fetchMonitors();
    }

    fetchData();
  }, []);

  const fetchMonitors = async () => {
    const response = await fetch("/api/status/monitors");
    const data = await response.json();
    setMonitors(data as Monitor[]);
  };

  return (
    <main className={"w-screen"}>
      <div className="rounded-2xl p-5">
        <motion.h1
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.2, duration: 0.6 }}
          className="text-4xl font-extrabold mb-8 text-white"
        >
          Uptime Monitors
        </motion.h1>

        <div className={"flex flex-col gap-5"}>
          {monitors.map((monitor) => {
            if (!monitor.is_parent) return <span key={monitor.id}></span>;
            return (
              <div>
                {new MonitorBuilder(
                  monitor,
                  monitors.filter((m) => m.id != monitor.id),
                ).toMonitor()}
              </div>
            );
          })}
        </div>
      </div>
    </main>
  );
}
