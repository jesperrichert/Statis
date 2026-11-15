export type Monitor = {
    identifier: string;
    is_parent: boolean;
    group: string | null;
    name: string;
    description: string;
    type: MonitorType;
    schedule_time: number;
    type_data: any | null;
    state: MonitorState;
    events: string[];
    affected_by: MonitorAffect[]
    external_url: string | null;
};

export type MonitorAffect = {
    name: string
    link: string
}

export enum MonitorState {
    ONLINE = "ONLINE",
    OFFLINE = "OFFLINE",
    DEGRADED = "DEGRADED",
    AFFECTED = "AFFECTED",
}

export enum MonitorType {
    GROUP = "GROUP",
    TCP = "TCP",
    HTTP = "HTTP",
    PING = "PING",
}

export const MonitorStateStyles: Record<
    MonitorState,
    { color: string; name: string; text: string }
> = {
    [MonitorState.DEGRADED]: {
        color: "bg-orange-500",
        text: "bg-orange-500/10 text-orange-400",
        name: "Degraded",
    },
    [MonitorState.AFFECTED]: {
        color: "bg-yellow-500",
        text: "bg-yellow-500/10 text-yellow-400",
        name: "Affected",
    },
    [MonitorState.OFFLINE]: {
        color: "bg-red-400",
        text: "bg-red-500/10 text-red-400",
        name: "Offline",
    },
    [MonitorState.ONLINE]: {
        color: "bg-green-500",
        text: "bg-emerald-500/10 text-emerald-400",
        name: "Online",
    },
};
