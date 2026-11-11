export type Monitor = {
    identifier: string;
    is_parent: boolean;
    group: string | null;
    name: string;
    description: string;
    type: MonitorType;
    schedule_time: number;
    check_data: any | null;
    state: MonitorState;
    effected_by: MonitorEffect[];
    external_url: string | null;
};

export type MonitorEffect = {
    name: string
    link: string
}

export enum MonitorState {
    ONLINE = "online",
    OFFLINE = "offline",
    DEGRADED = "degraded",
    EFFECTED = "effected",
}

export enum MonitorType {
    GROUP = "group",
    TCP = "tcp",
    HTTP = "http",
    PING = "ping",
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
    [MonitorState.EFFECTED]: {
        color: "bg-yellow-500",
        text: "bg-yellow-500/10 text-yellow-400",
        name: "Effected",
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
