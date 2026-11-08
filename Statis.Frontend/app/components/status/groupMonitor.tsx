import type { Monitor } from "../../../types/status";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion";
import { MonitorRow, StatusDot } from "./statusMonitor";


export function MonitorGroup({
  parent,
  children,
}: {
  parent: Monitor;
  children: Monitor[];
}) {
  const onlineCount = children.filter((c) => true).length;
  const allOnline = onlineCount === children.length;

  return (
    <div className="overflow-hidden rounded-2xl border border-white/10 bg-white/4 backdrop-blur">
      <Accordion type="single" collapsible defaultValue={`group-${parent.id}`}>
        <AccordionItem value={`group-${parent.id}`} className="border-none">
          <AccordionTrigger className="px-5 py-4 hover:bg-white/3 hover:no-underline">
            <div className="flex flex-1 items-center gap-3 text-left">
              <StatusDot online={allOnline} />
              <span className="text-xl font-semibold text-white">
                {parent.name}
              </span>
              <span className="ml-auto mr-3 text-sm text-white/50">
                {onlineCount}/{children.length} online
              </span>
            </div>
          </AccordionTrigger>
          <AccordionContent className="px-5 pb-5">
            {children.length > 0 ? (
              <div className="grid gap-2 sm:grid-cols-2 xl:grid-cols-3">
                {children.map((child) => (
                  <MonitorRow key={child.id} monitor={child} />
                ))}
              </div>
            ) : (
              <p className="rounded-xl border border-dashed border-white/10 p-4 text-center text-sm text-white/40">
                Keine Monitore gefunden
              </p>
            )}
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  );
}