"use server";

import {NextResponse} from "next/server";
import {Event} from "@/types/events";

export async function GET() {
    const response = await fetch(process.env.EVENTS_DATA as string, {
        method: "GET"
    })
    const data = await response.json() as Event[]
    return NextResponse.json({
        events: data
    })
}
