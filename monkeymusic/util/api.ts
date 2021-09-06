import { MonkeyApi, Configuration } from 'monkey-api'

const config = new Configuration({})
export const api = new MonkeyApi(config, 'http://localhost:8081')
