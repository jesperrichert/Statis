"use client";

import {motion} from "framer-motion";
import {useEffect, useState} from "react";
import {MonitorBuilder} from "~/lib/MonitorBuilder";
import type {Monitor} from "../../types/status";

export default function Page() {
    const [monitors, setMonitors] = useState<Monitor[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch("/api/status/monitors")
            .then((res) => res.json() as Promise<Monitor[]>)
            .then(setMonitors)
            .catch(console.error)
            .finally(() => setLoading(false));
    }, []);

    return (
        <main className={"w-screen"}>
            <div className="rounded-2xl p-5">
                <motion.h1
                    initial={{opacity: 0, y: -20}}
                    animate={{opacity: 1, y: 0}}
                    transition={{delay: 0.2, duration: 0.6}}
                    className="text-4xl font-extrabold mb-8 text-white"
                >
                    Uptime Monitors
                </motion.h1>

                <div className={"flex flex-col gap-5"}>
                    {loading && (
                        <p className="text-center text-white/40">Loading Monitors...</p>
                    )}

                    {!loading && (monitors && monitors.length == 0) && (
                        <p className="rounded-2xl border border-dashed border-white/10 p-8 text-center text-white/40">
                            No monitors found
                        </p>
                    )}

                    {monitors && monitors.length > 0 && monitors.map((monitor) => {
                        if (!monitor.is_parent) return <span key={monitor.identifier}></span>;
                        return (
                            <div>
                                {new MonitorBuilder(
                                    monitor,
                                    monitors.filter((m) => m.identifier != monitor.identifier),
                                ).toMonitor()}
                            </div>
                        );
                    })}
                </div>
            </div>
        </main>
    );
}
