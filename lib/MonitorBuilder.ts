import {Monitor, MonitorType} from "@/types/status";
import {StatusMonitor} from "@/components/status/statusMonitor";
import {GroupMonitor} from "@/components/status/groupMonitor";
import {bool} from "sharp";

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
                return this.buildHttpMonitor(this.monitor)
            }
            case MonitorType.TCP: {
                return this.buildTcpMonitor(this.monitor)
            }
            case MonitorType.PING: {
                return this.buildPingMonitor(this.monitor)
            }
        }
    }

    private buildTcpMonitor(data: Monitor) {
        return StatusMonitor({
            monitor: data
        })
    }

    private buildHttpMonitor(data: Monitor) {
        return StatusMonitor({
            monitor: data
        })
    }

    private buildGroupMonitor(data: Monitor) {
        return GroupMonitor({
            monitor: data,
            monitors: this.otherMonitors.filter((m) => m.group == data.id)
        })
    }

    private buildPingMonitor(data: Monitor) {
        return StatusMonitor({
            monitor: data
        })
    }

}
