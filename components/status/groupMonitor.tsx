"use client"

import {Monitor} from "@/types/status";
import {MonitorBuilder} from "@/lib/MonitorBuilder";
import {Accordion, AccordionContent, AccordionItem, AccordionTrigger} from "@/components/ui/accordion";
import Link from "next/link";
import {Badge} from "@/components/ui/badge";
import {ArrowUpRightIcon} from "lucide-react";

export function GroupMonitor(
    {
        monitor,
        monitors
    }:
    {
        monitor: Monitor,
        monitors: Monitor[]
    }
) {

    return (
        <Accordion key={monitor.id} type="single" collapsible defaultValue="item-1">
            <AccordionItem value="item-1">
                <AccordionTrigger>
                    <span>
                        {monitor.name}
                    </span>
                </AccordionTrigger>
                <AccordionContent>
                    <div className={"inline-flex"}>
                        <div
                            className="inline-block h-[100px] min-h-[1em] w-0.5 self-stretch bg-neutral-100 dark:bg-white/10">
                        </div>
                        {
                            monitors.length > 0 ? monitors.map((monitor) => {
                                    return (
                                        <div className={"w-96"}>
                                            {new MonitorBuilder(monitor, monitors).toMonitor()}
                                        </div>

                                    )
                                }
                            ) : <div className={"w-96 p-2"}>No Monitor found</div>
                        }
                    </div>
                </AccordionContent>
            </AccordionItem>
        </Accordion>
    )
}