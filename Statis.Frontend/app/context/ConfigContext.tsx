import {createContext, useEffect, useState} from "react";
import type {Config} from "../../types/config";

export const ConfigContext = createContext<Config>({} as Config);

export function ConfigContextProvider() {
    const [config, setConfig] = useState<Config | null>();

    useEffect(() => {
        if (!config)
            fetch("/config.json")
                .then((res) => res.json())
                .then((value: Config) => setConfig(value))
                .catch(_ => {
                });
    })

    return (<ConfigContext value={config ?? {} as Config}/>);
}
