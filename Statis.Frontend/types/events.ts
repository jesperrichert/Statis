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
