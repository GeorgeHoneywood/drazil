import { MonkeyApi, Configuration } from 'monkey-api'

const config = new Configuration({})

export const apiBase = process.env.baseUrl || window.location.host
console.log('baseUrl', process.env.baseUrl)

export const api = new MonkeyApi(config, apiBase)

export const getSongLink = function (artistID: string, albumID: string, songID: string) {
  return `${apiBase}/v1/artist/${artistID}/album/${albumID}/song/${songID}/media`
}

export const getAlbumArtLink = function (artistID: string, albumID: string) {
  return `${apiBase}/v1/artist/${artistID}/album/${albumID}/art`
}
