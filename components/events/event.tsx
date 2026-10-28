"use client"

import {Button} from "@/components/ui/button";
import {Event, EventType} from "@/types/events";
import {Badge} from "@/components/ui/badge";
import {ExternalLinkIcon} from "lucide-react";
import {Accordion, AccordionContent, AccordionItem, AccordionTrigger} from "@/components/ui/accordion";
import {useState} from "react";

export function IncidenceEvent({
                                   event
                               }: {
    event: Event
}) {

    const [isOpen, setOpen] = useState(false)

    const typeColor = (type: EventType) => {
        switch (type) {
            case EventType.INVESTIGATING: {
                return "#FF0000"
            }
            case EventType.UPDATE: {
                return "#808080"
            }
            case EventType.IDENTIFIED: {
                return "#FFA500"
            }
            case EventType.RESOLVED: {
                return "#00FF00"
            }
        }
    }

    return (
        <Accordion type="single" collapsible className="w-screen flex justify-center rounded-lg border">
            <AccordionItem key={event.id} value="item-1" className="border-b px-3 last:border-b-0">
                <AccordionTrigger
                    className={`hover:no-underline cursor-pointer`}
                    onClick={() => setOpen(!isOpen)}
                >
                    <span>
                        {event.title}
                    </span>
                    {/* bg-[#00FF00}] bg-[#FF0000] bg-[#808080] bg-[#FFA500]  */}
                    <span className={"right-1/12 absolute"}>
                        <Badge className={`bg-[${typeColor(event.type)}]`}>
                        {event.type.charAt(0).toUpperCase() + event.type.slice(1, event.type.length)}
                    </Badge>
                   </span>
                </AccordionTrigger>
                <AccordionContent>
                   <span className={"mb-10 p-2"}>
                           {event.description}
                   </span>
                    <hr className={"p-1 mt-2 mx-2 h-3"}/>
                    <div className={"mt-2 mb-5 flex"}>
                        <Button variant="outline" size="sm"
                                className="w-full text-zinc-700 font-extrabold hover:font-medium hover:cursor-pointer"
                                onClick={() => window.open(`/status/incidence/${event.id}`)}
                        >
                            <span>
                                Open Incidence
                            </span>
                            <span>
                                    <ExternalLinkIcon></ExternalLinkIcon>
                                </span>
                        </Button>
                    </div>
                </AccordionContent>
            </AccordionItem>
        </Accordion>

    )
}