"use client";

import { useEffect, useState } from "react";
import { motion } from "framer-motion";
import type { PageCallback } from "~/types/pages";
import Markdown from "markdown-to-jsx";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";

export default function IncidencePage() {
  const [mdContent, setMdContent] = useState<string>("");
  const [metadata, setMetadata] = useState<Record<string, any>>({});
  const [loading, setLoading] = useState<boolean>(true);
  const pathname = window.location.href.split("/").pop();

  useEffect(() => {
    async function fetch() {
      await fetchIncidence();
    }

    fetch();
  }, []);

  const fetchIncidence = async () => {
    const pageData = await fetch(`/api/status/incidence/${pathname}`, {
      method: "GET",
    });

    const data = (await pageData.json()) as PageCallback;

    console.log(data);

    setMetadata(data.metaData);
    setMdContent(data.markdown);
    setLoading(false);
  };

  return (
    <>
      <title>{metadata?.title ?? "Statis"}</title>
      <main className="flex items-center justify-center">
        <motion.section
          initial={{ opacity: 0, y: 24 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, ease: "easeOut" }}
          className="w-full max-w-3xl bg-zinc/90 dark:bg-zinc-900/80 rounded-2xl backdrop-blur-lg"
        >
          {/* Headline */}
          {metadata?.title && (
            <h1 className="text-4xl font-extrabold mb-6 text-center bg-gradient-to-r from-zinc-400 to-zinc-500 bg-clip-text text-transparent drop-shadow">
              {metadata.title}
            </h1>
          )}

          {/* Meta Info */}
          <div className="flex justify-center gap-6 mb-8 text-sm text-zinc-500 dark:text-zinc-400">
            {metadata?.date && (
              <span>
                <span className="font-semibold">📅</span> {metadata.date}
              </span>
            )}
            {metadata?.author && (
              <span>
                <span className="font-semibold">✍️</span> {metadata.author}
              </span>
            )}
          </div>

          {/* Markdown Content */}
          <Markdown
            options={{
              overrides: {
                div: {
                  component: ({ children }: { children: any }) => {
                    return <div className="text-white">{children}</div>;
                  },
                },
                h1: {
                  component: ({ children }: { children: any }) => (
                    <h1 className="text-white text-3xl font-bold mt-10 mb-4 bg-gradient-to-r from-zinc-400 to-zinc-500 bg-clip-text">
                      {children}
                    </h1>
                  ),
                },
                h2: {
                  component: ({ children }: { children: any }) => (
                    <h2 className="text-2xl font-bold mt-8 mb-3 text-white dark:text-zinc-300">
                      {children}
                    </h2>
                  ),
                },
                h3: {
                  component: ({ children }: { children: any }) => (
                    <h3 className="text-white text-xl font-semibold mt-6 mb-2 dark:text-zinc-300">
                      {children}
                    </h3>
                  ),
                },
                p: {
                  component: ({ children }: { children: any }) => (
                    <p className="text-white leading-relaxed dark:text-zinc-200 mb-4">
                      {children}
                    </p>
                  ),
                },
                ul: {
                  component: ({ children }: { children: any }) => (
                    <ul className="list-disc ml-6 mb-4 text-white dark:text-zinc-200 space-y-2">
                      {children}
                    </ul>
                  ),
                },
                ol: {
                  component: ({ children }: { children: any }) => (
                    <ol className="list-decimal ml-6 mb-4 text-white dark:text-zinc-200 space-y-2">
                      {children}
                    </ol>
                  ),
                },
                li: {
                  component: ({ children }: { children: any }) => (
                    <li className="pl-1">{children}</li>
                  ),
                },
                blockquote: {
                  component: ({ children }: { children: any }) => (
                    <blockquote className="border-l-4 border-zinc-400 pl-4 italic text-white dark:text-zinc-300 bg-zinc-50 dark:bg-zinc-800/50 py-2 mb-4 rounded">
                      {children}
                    </blockquote>
                  ),
                },
                code: {
                  component: ({
                    className,
                    children,
                  }: {
                    className: any;
                    children: any;
                  }) => {
                    const isBlock = className && className.startsWith("lang-");
                    const language = className
                      ? className.replace("lang-", "")
                      : "";

                    return isBlock ? (
                      <SyntaxHighlighter
                        language={language}
                        customStyle={{
                          borderRadius: "0.75rem",
                          margin: "1.5rem 0",
                          fontSize: "1rem",
                          background: "none",
                        }}
                      >
                        {String(children).replace(/\n$/, "")}
                      </SyntaxHighlighter>
                    ) : (
                      <code className="bg-zinc-100 dark:bg-zinc-800 px-1.5 py-0.5 rounded text-sm font-mono text-pink-600 dark:text-pink-400">
                        {children}
                      </code>
                    );
                  },
                },
                img: {
                  component: ({ src, alt }: { src: any; alt: any }) => (
                    <img
                      src={src}
                      alt={alt}
                      className="rounded-lg border shadow-md mx-auto my-6 max-h-80 object-contain"
                    />
                  ),
                },
                hr: {
                  component: () => (
                    <hr className="my-8 border-t border-zinc-300 dark:border-zinc-700" />
                  ),
                },
                a: {
                  component: ({
                    href,
                    children,
                  }: {
                    href: any;
                    children: any;
                  }) => (
                    <a
                      href={href}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-white dark:text-zinc-400 underline hover:text-zinc-500 transition"
                    >
                      {children}
                    </a>
                  ),
                },
              },
            }}
          >
            {mdContent}
          </Markdown>
        </motion.section>
      </main>
    </>
  );
}
