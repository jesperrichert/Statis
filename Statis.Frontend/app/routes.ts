import { type RouteConfig, index, layout, route } from "@react-router/dev/routes";

export default [
  layout("routes/layout.tsx", [
    index("routes/page.tsx"),
    route("/events", "routes/events/page.tsx"),
    route("/incidences/:id", "routes/incidence/page.tsx"),
  ]),
] satisfies RouteConfig;
