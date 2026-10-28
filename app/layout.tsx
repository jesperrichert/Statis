"use client";

import Link from "next/link";
import Image from "next/image";
import "./globals.css";
import {House, Calendars} from "lucide-react"

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode;
}) {
    return (
        <html lang="de">
        <head>
            <title>Statis</title>
            <meta name="description" content="Jespersen.space"/>
            <meta name="viewport" content="width=device-width, initial-scale=1"/>
            <link rel="icon" href="/statis.svg"/>
        </head>
        <body className="min-h-screen flex flex-col bg-zinc-900 text-zinc-100 transition-colors duration-300">
        <header className="border-b border-zinc-700">
            <div className="max-w-6xl mx-auto px-6 py-4 flex items-center justify-between">
                <nav className="flex gap-8 text-sm items-center justify-center font-medium text-zinc-300">
                    <div>
                        <Image
                            src="/statis.svg"
                            alt="Statis"
                            width={50}
                            height={50}
                        />
                    </div>
                    <Link
                        href="/"
                        className="hover:text-zinc-400 transition-colors inline-flex gap-3"
                        aria-label="homepage"
                    >
                        <House size={35}/>
                        <span className={"text-2xl font-extrabold"}>Home</span>
                    </Link>
                    <Link
                        href="/status/events"
                        className="hover:text-zinc-400 transition-colors inline-flex gap-3"
                        aria-label="contact"
                    >
                        <Calendars size={35} />
                        <span className={"text-2xl font-extrabold"}>Events</span>
                    </Link>
                </nav>
            </div>
        </header>

        <main className="grow max-w-6xl mx-auto px-6 py-8">
            {children}
        </main>
        </body>
        </html>
    );
}
