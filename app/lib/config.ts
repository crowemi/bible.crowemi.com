export const CONFIG = GetConfig();

export type Config = {
    uri: string;
    databaseID: string;
    projectID: string;
}

export function GetConfig(): Config {
    const config = process.env.CONFIG;
    if (config) {
        const ret = JSON.parse(Buffer.from(config, 'base64').toString('utf-8')) as Config;
        return ret;
    }
    throw new Error("Config not found");
}