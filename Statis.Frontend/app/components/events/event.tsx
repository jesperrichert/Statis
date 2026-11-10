import { ExternalLinkIcon } from "lucide-react";
import { EventType, EventTypeStyles, type Event } from "../../../types/events";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "../ui/accordion";
import { Badge } from "../ui/badge";
import { Button } from "../ui/button";

export function IncidenceEvent({ event }: { event: Event }) {
  const style = EventTypeStyles[event.type];
  const time = new Date(event.date).toLocaleTimeString("de-DE", {
    hour: "2-digit",
    minute: "2-digit",
  });

  return (
    <div className="w-full overflow-hidden rounded-2xl border border-white/10 bg-white/4 backdrop-blur">
      <Accordion type="single" collapsible>
        <AccordionItem value={`event-${event.identifier}`} className="border-none">
          <AccordionTrigger className="cursor-pointer px-5 py-4 hover:bg-white/3 hover:no-underline">
            <div className="flex flex-1 items-center gap-3 text-left">
              <span className={`h-2.5 w-2.5 shrink-0 rounded-full ${style.dot}`} />
              <div className="flex min-w-0 flex-col">
                <span className="truncate font-semibold text-white">
                  {event.title}
                </span>
                <span className="text-xs text-white/40">{time} Uhr</span>
              </div>
              <Badge
                variant="outline"
                className={`ml-auto mr-3 shrink-0 ${style.badge}`}
              >
                {style.label}
              </Badge>
            </div>
          </AccordionTrigger>

          <AccordionContent className="px-5 pb-5">
            <p className="whitespace-pre-line text-sm leading-relaxed text-white/70">
              {event.description}
            </p>

            <div className="mt-4 border-t border-white/10 pt-4">
              <Button
                asChild
                variant="outline"
                size="sm"
                className="w-full gap-2 border-white/10 bg-transparent text-white hover:bg-white/10 hover:text-white"
              >
                <a
                  href={`/incidences/${event.identifier}`}
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Open Incidence
                  <ExternalLinkIcon className="h-4 w-4" />
                </a>
              </Button>
            </div>
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  );
}