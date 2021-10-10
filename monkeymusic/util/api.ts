import { MonkeyApi, Configuration } from 'monkey-api'

const config = new Configuration({})

export const apiBase = 'http://192.168.1.169:4444'

export const api = new MonkeyApi(config, apiBase)

export const getSongLink = function (artistID: string, albumID: string, songID: string) {
  return `${apiBase}/v1/artist/${artistID}/album/${albumID}/song/${songID}/media`
}

export const getAlbumArtLink = function (artistID: string, albumID: string) {
  return `${apiBase}/v1/artist/${artistID}/album/${albumID}/art`
}
