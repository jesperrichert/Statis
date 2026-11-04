export type Monitor = {
    id: string
    is_parent: boolean
    group: string | null
    name: string
    description: string
    type: MonitorType
    external_url: string | null
    check_data: any | null
}

export enum MonitorType {
    GROUP = "group",
    TCP = "tcp",
    HTTP = "http",
    PING = "ping"
}