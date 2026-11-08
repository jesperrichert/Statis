"use client";

import { useEffect, useState } from "react";
import { motion } from "framer-motion";
import { IncidenceEvent } from "~/components/events/event";
import type { Event } from "../../../types/events";

export default function StatusEventPage() {
  const [events, setEvents] = useState<Record<string, Event[]>>({});

  useEffect(() => {
    async function fetchData() {
      await fetchEvents();
    }

    fetchData();
  }, []);

  const fetchEvents = async () => {
    const response = await fetch("/api/status/events");
    const data = (await response.json()) as Event[];
    const timedEvents: Record<string, Event[]> = {};
    data.forEach((e) => {
      const date = new Date(e.date);
      const key = `${date.getFullYear()}-${date.getUTCMonth()}-${date.getUTCDay()}`;

      console.log(key);

      const old = timedEvents[key]?.length >= 1 ? timedEvents[key] : [];
      timedEvents[key] = [...old, e];
    });

    setEvents(timedEvents);
  };

  return (
    <div
      key={"status"}
      className="w-screen flex flex-col justify-center items-center"
    >
      <motion.h1
        initial={{ opacity: 0, y: -20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.2, duration: 0.6 }}
        className="text-4xl font-extrabold text-center text-white p-3"
      >
        Incident Events
      </motion.h1>

      <div className={"flex justify-center items-center"}>
        {Object.keys(events).length > 0 ? (
          Object.values(events).map((e) => {
            const date = new Date(e[0].date);
            return (
              <>
                <div className="flex flex-col justify-center items-center">
                  <span className="flex font-bold p-2">
                    {date.getDay()}.{date.getMonth()}.{date.getFullYear()}
                  </span>
                  <div>
                    {e.map((event) => {
                      return (
                        <div className="p-2">
                          <IncidenceEvent event={event} />
                        </div>
                      );
                    })}
                  </div>
                  <span className="p-4 mx-4">
                    <hr className="w-60 border-t-4" />
                  </span>
                </div>
              </>
            );
          })
        ) : (
          <span className={"flex justify-center items-center"}>
            No Events found
          </span>
        )}
      </div>
    </div>
  );
}
