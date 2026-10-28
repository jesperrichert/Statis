export type Monitor = {
    id: string
    isParent: boolean
    group: string | null
    name: string
    description: string
    type: MonitorType
    external_url: string | null
    checkData: any | null
    isOnline: boolean
    isEffected: boolean
}

export enum MonitorType {
    GROUP = "group",
    TCP = "tcp",
    HTTP = "http",
    PING = "ping"
}