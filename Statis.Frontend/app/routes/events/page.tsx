"use client";

import { useEffect, useMemo, useState } from "react";
import { motion } from "framer-motion";
import { IncidenceEvent } from "~/components/events/event";
import type { Event } from "../../../types/events";

type DayGroup = { key: string; date: Date; events: Event[] };

export default function StatusEventPage() {
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("/api/status/events")
      .then((res) => res.json() as Promise<Event[]>)
      .then(setEvents)
      .catch(console.error)
      .finally(() => setLoading(false));
  }, []);

  const groups = useMemo<DayGroup[]>(() => {
    const map = new Map<string, DayGroup>();

    for (const e of events) {
      const date = new Date(e.date);
      const key = `${date.getFullYear()}-${date.getMonth()}-${date.getDate()}`;
      const group = map.get(key) ?? { key, date, events: [] };
      group.events.push(e);
      map.set(key, group);
    }

    return [...map.values()]
      .sort((a, b) => b.date.getTime() - a.date.getTime())
      .map((g) => ({
        ...g,
        events: g.events.sort(
          (a, b) => new Date(b.date).getTime() - new Date(a.date).getTime(),
        ),
      }));
  }, [events]);

  return (
    <main className="px-4 py-10 sm:px-8">
      <div className="mx-auto max-w-2xl">
        <motion.header
          initial={{ opacity: 0, y: -16 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="mb-10 text-center"
        >
          <h1 className="text-4xl font-extrabold tracking-tight text-white">
            Incident Events
          </h1>
          <p className="mt-1 text-sm text-white/50">
           History of All Outages and Updates
          </p>
        </motion.header>

        {loading && (
          <p className="text-center text-white/40">Loading Events…</p>
        )}

        {!loading && groups.length === 0 && (
          <p className="rounded-2xl border border-dashed border-white/10 p-8 text-center text-white/40">
            No events found
          </p>
        )}

        <div className="flex flex-col gap-10">
          {groups.map((group, i) => (
            <motion.section
              key={group.key}
              initial={{ opacity: 0, y: 12 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 0.1 + i * 0.07, duration: 0.4 }}
            >
              <h2 className="mb-3 flex items-center gap-3 text-sm font-semibold uppercase tracking-wider text-white/50">
                {group.date.toLocaleDateString("en-GB", {
                  weekday: "long",
                  day: "2-digit",
                  month: "long",
                  year: "numeric",
                })}
                <span className="h-px flex-1 bg-white/10" />
              </h2>

              <div className="flex flex-col gap-3">
                {group.events.map((event) => (
                  <IncidenceEvent key={event.id} event={event} />
                ))}
              </div>
            </motion.section>
          ))}
        </div>
      </div>
    </main>
  );
}