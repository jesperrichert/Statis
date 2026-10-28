"use server";

import {NextResponse} from "next/server";
import {Monitor, MonitorType} from "@/types/status";
import ping from "ping";
import {check as checkTCP} from "tcp-port-used"

export async function GET() {
    const response = await fetch(process.env.STATUS_DATA as string, {
        method: "GET"
    })
    const data = await response.json() as Monitor[]

    const monitors = await Promise.all(
        data.map(async (m) => {

            let isOnline: boolean = false
            switch (m.type) {
                case MonitorType.GROUP: {
                    isOnline = false
                }
                    break;
                case MonitorType.TCP: {
                    try {
                        isOnline = await checkTCP(m.checkData.port ? m.checkData.port : m.checkData.split(":")[1] ?? 80, m.checkData.host ? m.checkData.host : m.checkData.split(":")[0])
                    } catch (e) {
                        isOnline = false
                    }
                }
                    break;
                case MonitorType.PING: {
                    let res = await ping.promise.probe(m.checkData.host ? m.checkData.host : m.checkData, {
                        timeout: m.checkData.timeout ?? 30000,
                        min_reply: m.checkData.min_reply ?? 10
                    });
                    isOnline = res.alive
                }
                    break;
                case MonitorType.HTTP: {
                    const statusResponse = await fetch(m.checkData as string)
                    isOnline = statusResponse.ok
                }
                    break;
            }

            return {
                isOnline: isOnline,
                name: m.name,
                checkData: m.checkData,
                id: m.id,
                external_url: m.external_url,
                isParent: m.isParent,
                group: m.group,
                type: m.type
            } satisfies Monitor;
        }) as unknown as Monitor[]
    )


    return NextResponse.json({
        monitors: monitors
    })
}
