import { MonitorRow } from "~/components/status/statusMonitor"
import { MonitorType, type Monitor } from "../../types/status"
import { MonitorGroup } from "~/components/status/groupMonitor"


export class MonitorBuilder {
    monitor: Monitor
    otherMonitors: Monitor[]

    constructor(monitor: Monitor, otherMonitors: Monitor[] | null) {
        this.monitor = monitor
        this.otherMonitors = otherMonitors ?? []
    }

    public toMonitor() {
        switch (this.monitor.type) {
            case MonitorType.GROUP: {
                return this.buildGroupMonitor(this.monitor)
            }
            case MonitorType.HTTP: {
                return this.buildMonitor(this.monitor)
            }
            case MonitorType.TCP: {
                return this.buildMonitor(this.monitor)
            }
            case MonitorType.PING: {
                return this.buildMonitor(this.monitor)
            }
        }
    }

    private buildMonitor(data: Monitor) {
        return MonitorRow({
            monitor: data
        })
    }

    private buildGroupMonitor(data: Monitor) {
        return MonitorGroup({
            parent: data,
            children: this.otherMonitors.filter((m) => m.group == data.identifier)
        })
    }
}
