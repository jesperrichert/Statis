"use client"

import {useEffect, useState} from "react";
import {motion} from "framer-motion";
import { IncidenceEvent } from "~/components/events/event";

export default function StatusEventPage() {

    const [events, setEvents] = useState<Event[]>([])

    useEffect(() => {
        async function fetchData() {
            await fetchEvents()
        }

        fetchData()
    }, []);

    const fetchEvents = async () => {
        const response = await fetch("/api/status/events")
        const data = await response.json()
        setEvents(data.events as Event[])
    }

    return (
        <div key={"status"}>

            <motion.h1
                initial={{opacity: 0, y: -20}}
                animate={{opacity: 1, y: 0}}
                transition={{delay: 0.2, duration: 0.6}}
                className="text-4xl font-extrabold text-center mb-8 bg-linear-to-r from-zinc-300 to-zinc-500 bg-clip-text text-transparent"
            >
                Incident Events
            </motion.h1>

            <div className={"flex justify-center items-center"}>
                {
                    events.length > 0 ? (events.map((e) => {
                        return (
                            <div className={"flex justify-center items-center w-96"}>
                                <IncidenceEvent event={e} />
                            </div>
                        )
                    })) : (<span className={"flex justify-center items-center"}>No Events found</span>)
                }
            </div>
        </div>
    )
}