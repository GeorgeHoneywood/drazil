import { DrazilApi, Configuration } from 'drazil-api'
import { SpecSong } from '~/api-client'

export type FlatSong = SpecSong & {
  albumName: string,
  artistName: string,
  artistId: string,
  albumId:string
  }

const config = new Configuration({})

export const apiBase = process.env.baseUrl || `${window.location.protocol}//${window.location.host}`

export const api = new DrazilApi(config, apiBase)

export const getSongLink = function (artistID: string, albumID: string, songID: string) {
  return `${apiBase}/v1/artist/${artistID}/album/${albumID}/song/${songID}/media`
}

export const getAlbumArtLink = function (artistID: string, albumID: string) {
  return `${apiBase}/v1/artist/${artistID}/album/${albumID}/art`
}
