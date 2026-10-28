"use client";

import { Outlet } from "react-router";
import { House, Calendars } from "lucide-react";

export default function RootLayout() {
  return (
    <>
      <header className="border-b border-zinc-700">
        <div className="max-w-6xl mx-auto px-6 py-4 flex items-center justify-between">
          <nav className="flex gap-8 text-sm items-center justify-center font-medium text-zinc-300">
            <div>
              <img src="/statis.svg" alt="Statis" width={50} height={50} />
            </div>
            <a
              href="/"
              className="hover:text-zinc-400 transition-colors inline-flex gap-3"
              aria-label="homepage"
            >
              <House size={35} />
              <span className={"text-2xl font-extrabold"}>Home</span>
            </a>
            <a
              href="/status/events"
              className="hover:text-zinc-400 transition-colors inline-flex gap-3"
              aria-label="contact"
            >
              <Calendars size={35} />
              <span className={"text-2xl font-extrabold"}>Events</span>
            </a>
          </nav>
        </div>
      </header>

      <main className="grow max-w-6xl mx-auto px-6 py-8">
        <Outlet />
      </main>
    </>
  );
}
