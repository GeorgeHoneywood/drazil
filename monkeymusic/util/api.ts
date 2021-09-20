import { MonkeyApi, Configuration } from 'monkey-api'

const config = new Configuration({})

export const api = new MonkeyApi(config, 'http://localhost:4444')

export const getSongLink = function (artistID: string, albumID: string, songID: string) {
  return `http://localhost:4444/v1/artist/${artistID}/album/${albumID}/song/${songID}/media`
}
