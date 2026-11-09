export type Event = {
  id: string;
  title: string;
  description: string;
  type: EventType;
  date: Date;
};

export enum EventType {
  RESOLVED = "resolved",
  UPDATE = "update",
  INVESTIGATING = "investigating",
  IDENTIFIED = "identified",
}


export const EventTypeStyles: Record<EventType, { label: string; dot: string; badge: string }> = {
  [EventType.INVESTIGATING]: {
    label: "Investigating",
    dot: "bg-red-500",
    badge: "bg-red-500/10 text-red-400 border-red-500/30",
  },
  [EventType.UPDATE]: {
    label: "Update",
    dot: "bg-zinc-400",
    badge: "bg-zinc-500/10 text-zinc-300 border-zinc-500/30",
  },
  [EventType.IDENTIFIED]: {
    label: "Identified",
    dot: "bg-orange-500",
    badge: "bg-orange-500/10 text-orange-400 border-orange-500/30",
  },
  [EventType.RESOLVED]: {
    label: "Resolved",
    dot: "bg-emerald-500",
    badge: "bg-emerald-500/10 text-emerald-400 border-emerald-500/30",
  },
};