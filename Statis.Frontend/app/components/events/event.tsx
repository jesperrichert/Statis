import { ExternalLinkIcon } from "lucide-react";
import { useState } from "react";
import { EventType, type Event } from "../../../types/events";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "../ui/accordion";
import { Badge } from "../ui/badge";
import { Button } from "../ui/button";

export function IncidenceEvent({ event }: { event: Event }) {
  const [isOpen, setOpen] = useState(false);

  const typeColor = (type: EventType) => {
    switch (type) {
      case EventType.INVESTIGATING: {
        return "#FF0000";
      }
      case EventType.UPDATE: {
        return "#808080";
      }
      case EventType.IDENTIFIED: {
        return "#FFA500";
      }
      case EventType.RESOLVED: {
        return "#00FF00";
      }
    }
  };

  return (
    <Accordion
      type="single"
      collapsible
      className="w-150 flex justify-center rounded-lg border"
    >
      <AccordionItem
        key={event.id}
        value="item-1"
        className="border-b px-3 last:border-b-0"
      >
        <AccordionTrigger
          className={`hover:no-underline cursor-pointer`}
          onClick={() => setOpen(!isOpen)}
        >
          <span>{event.title}</span>
          {/* bg-[#00FF00}] bg-[#FF0000] bg-[#808080] bg-[#FFA500]  */}
          <span className={"right-1/12 absolute"}>
            <Badge className={`bg-[${typeColor(event.type)}]`}>
              {event.type.charAt(0).toUpperCase() +
                event.type.slice(1, event.type.length)}
            </Badge>
          </span>
        </AccordionTrigger>
        <AccordionContent>
          <span className={"mb-10 p-2"}>{event.description}</span>
          <hr className={"p-1 mt-2 mx-2 h-3"} />
          <div className={"mt-2 mb-5 flex"}>
            <Button
              variant="outline"
              size="sm"
              className="w-full text-zinc-700 font-extrabold hover:font-medium hover:cursor-pointer"
              onClick={() => window.open(`/incidences/${event.id}`)}
            >
              <span>Open Incidence</span>
              <span>
                <ExternalLinkIcon></ExternalLinkIcon>
              </span>
            </Button>
          </div>
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  );
}
