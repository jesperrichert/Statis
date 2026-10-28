"use server";

import {NextRequest, NextResponse} from "next/server";
import {Event} from "@/types/events";

export async function GET(req: NextRequest) {
    const page = req.url.split("/").pop();
    
    const response = await fetch(
        process.env.INCIDENCE_DIRECTORY?.replaceAll("{page}", `${page}`) as string
    );

    const data = await response.text();

    const metadataRaw = data.split("---")[1];
    const markdown = data.split("---")[2] || "";

    const meta: Record<string, any> = {};
    if (metadataRaw) {
        metadataRaw
            .trim()
            .split("\n")
            .forEach((line) => {
                const [key, value] = line.split(":").map((s) => s.trim());
                if (key && value) meta[key] = value.replace(/"/g, "");
            });
    }

    return NextResponse.json({
        markdown: markdown,
        metaData: meta,
    });
}