"use client";

import {type ReactNode, useEffect, useState} from "react";
import {motion} from "framer-motion";
import Markdown from "markdown-to-jsx";
import {CalendarDays, Signature} from "lucide-react";
import SyntaxHighlighter from "react-syntax-highlighter"

type C = { children?: ReactNode };

const mdOverrides = {
    h1: {
        component: ({children}: C) => (
            <h1 className="mt-12 mb-4 text-3xl font-semibold tracking-tight text-zinc-100">
                {children}
            </h1>
        ),
    },
    h2: {
        component: ({children}: C) => (
            <h2 className="mt-10 mb-3 text-2xl font-semibold tracking-tight text-zinc-100">
                {children}
            </h2>
        ),
    },
    h3: {
        component: ({children}: C) => (
            <h3 className="mt-8 mb-2 text-xl font-medium text-zinc-200">
                {children}
            </h3>
        ),
    },
    p: {
        component: ({children}: C) => (
            <p className="mb-5 leading-7 text-zinc-400">{children}</p>
        ),
    },
    ul: {
        component: ({children}: C) => (
            <ul className="mb-5 ml-5 list-disc space-y-1.5 text-zinc-400 marker:text-zinc-600">
                {children}
            </ul>
        ),
    },
    ol: {
        component: ({children}: C) => (
            <ol className="mb-5 ml-5 list-decimal space-y-1.5 text-zinc-400 marker:text-zinc-600">
                {children}
            </ol>
        ),
    },
    li: {
        component: ({children}: C) => <li className="pl-1">{children}</li>,
    },
    blockquote: {
        component: ({children}: C) => (
            <blockquote className="mb-5 border-l-2 border-zinc-700 pl-4 text-zinc-500">
                {children}
            </blockquote>
        ),
    },
    pre: {
        component: ({children}: C) => <>{children}</>,
    },
    code: {
        component: ({
                        className,
                        children,
                    }: {
            className?: string;
            children?: ReactNode;
        }) => {
            const isBlock = className?.startsWith("lang-");
            const language = className?.replace("lang-", "") ?? "";

            if (!isBlock) {
                return (
                    <code className="rounded bg-zinc-800 px-1.5 py-0.5 font-mono text-[0.85em] text-zinc-200">
                        {children}
                    </code>
                );
            }

            return (
                <div className="my-6 overflow-hidden rounded-lg border border-zinc-800 bg-zinc-900/50">
                    {language && (
                        <div className="border-b border-zinc-800 px-4 py-1.5 font-mono text-xs text-zinc-600">
                            {language}
                        </div>
                    )}
                    <SyntaxHighlighter
                        language={language}
                        customStyle={{
                            margin: 0,
                            padding: "1rem 1.25rem",
                            fontSize: "0.875rem",
                            background: "transparent",
                        }}
                    >
                        {String(children).replace(/\n$/, "")}
                    </SyntaxHighlighter>
                </div>
            );
        },
    },
    img: {
        component: ({src, alt}: { src?: string; alt?: string }) => (
            <img
                src={src}
                alt={alt}
                loading="lazy"
                className="mx-auto my-8 max-h-96 rounded-lg border border-zinc-800 object-contain"
            />
        ),
    },
    hr: {
        component: () => <hr className="my-10 border-0 border-t border-zinc-800"/>,
    },
    a: {
        component: ({href, children}: { href?: string; children?: ReactNode }) => (
            <a
                href={href}
                target="_blank"
                rel="noopener noreferrer"
                className="text-zinc-200 underline decoration-zinc-700 underline-offset-4 transition-colors hover:decoration-zinc-300"
            >
                {children}
            </a>
        ),
    },
    table: {
        component: ({children}: C) => (
            <div className="my-6 overflow-x-auto rounded-lg border border-zinc-800">
                <table className="w-full text-left text-sm">{children}</table>
            </div>
        ),
    },
    th: {
        component: ({children}: C) => (
            <th className="bg-zinc-900 px-4 py-2 font-medium text-zinc-200">
                {children}
            </th>
        ),
    },
    td: {
        component: ({children}: C) => (
            <td className="border-t border-zinc-800 px-4 py-2 text-zinc-400">
                {children}
            </td>
        ),
    },
};

function parseFrontmatter(raw: string) {
    const match = raw.match(/^---\r?\n([\s\S]*?)\r?\n---\r?\n?([\s\S]*)$/);
    if (!match) return {meta: {} as Record<string, string>, body: raw};

    const meta: Record<string, string> = {};
    match[1]
        .split(/\r?\n/)
        .filter(Boolean)
        .forEach((line) => {
            const i = line.indexOf(":");
            if (i === -1) return;
            meta[line.slice(0, i).trim()] = line
                .slice(i + 1)
                .trim()
                .replace(/^["']|["']$/g, "");
        });

    return {meta, body: match[2]};
}

export default function IncidencePage() {
    const [mdContent, setMdContent] = useState("");
    const [metadata, setMetadata] = useState<Record<string, string>>();
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const slug = window.location.pathname.split("/").filter(Boolean).pop();

        (async () => {
            try {
                const res = await fetch(`/api/status/incidences/${slug}/raw`);
                const {meta, body} = parseFrontmatter(await res.text());
                setMetadata(meta);
                setMdContent(body);
            } finally {
                setLoading(false);
            }
        })();
    }, []);

    return (
        <>
            <title>{metadata?.title ?? "Statis"}</title>
            <main className="min-h-screen px-4 py-16 sm:py-24">
                {!loading && (
                    <motion.article
                        initial={{opacity: 0, y: 12}}
                        animate={{opacity: 1, y: 0}}
                        transition={{duration: 0.35, ease: "easeOut"}}
                        className="mx-auto w-full max-w-3xl"
                    >
                        <header className="mb-10 border-b border-zinc-800 pb-8">
                            {metadata?.title && (
                                <h1 className="text-4xl font-semibold tracking-tight text-zinc-100 sm:text-5xl">
                                    {metadata.title}
                                </h1>
                            )}

                            <div className="mt-4 flex flex-wrap gap-x-5 gap-y-1 text-sm text-zinc-500">
                                {metadata?.date && (
                                    <span className="inline-flex items-center gap-1.5">
                    <CalendarDays className="size-4"/>
                                        {metadata.date}
                  </span>
                                )}
                                {metadata?.author && (
                                    <span className="inline-flex items-center gap-1.5">
                    <Signature className="size-4"/>
                                        {metadata.author}
                  </span>
                                )}
                            </div>
                        </header>

                        <Markdown options={{overrides: mdOverrides}}>{mdContent}</Markdown>
                    </motion.article>
                )}
            </main>
        </>
    );
}