"use client";

import {Outlet} from "react-router";
import {Calendars, House} from "lucide-react";

export default function RootLayout() {
    return (
        <div>
            <header className="border-b border-zinc-700">
                <div className="flex ml-2 items-center justify-between">
                    <nav
                        className="flex flex-row my-5 gap-8 text-sm items-center justify-center font-medium text-zinc-300">
                        <div>
                            <img src="/statis.svg" alt="Statis" width={50} height={50}/>
                        </div>
                        <a
                            href="/"
                            className="hover:text-zinc-400 transition-colors inline-flex gap-2"
                            aria-label="homepage"
                        >
                            <House/>
                            <span className={"text-xl font-extrabold"}>Home</span>
                        </a>
                        <a
                            href="/events"
                            className="hover:text-zinc-400 transition-colors inline-flex gap-2"
                            aria-label="contact"
                        >
                            <Calendars/>
                            <span className={"text-xl font-extrabold"}>Events</span>
                        </a>
                    </nav>
                </div>
            </header>

            <Outlet/>
        </div>
    );
}
