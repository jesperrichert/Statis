export type Event = {
    id: string
    title: string
    description: string
    type: EventType
}

export enum EventType {
    RESOLVED = "resolved",
    UPDATE = "update",
    INVESTIGATING = "investigating",
    IDENTIFIED = "identified"
}