<template>
  <v-row dense>
    <v-col
      v-for="album in albums"
      :key="album.id"
      cols="6"
      sm="3"
      md="3"
      lg="2"
    >
      <v-card
        tile
        @click="showAlbum(album.artistId, album.id)"
      >
        <v-responsive :aspect-ratio="1/1">
          <v-img
            :src="getAlbumArtLink(album.artistId, album.id)"
            class="align-end"
            gradient="to bottom, rgba(0,0,0,0), rgba(0,0,0,.5)"
            height="100%"
          >
            <v-card-title
              class="white--text text-subtitle-1"
              style="word-break: normal"
              v-text="album.name"
            />
          </v-img>
        </v-responsive>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from 'vue'
import { SpecAllAlbumsReply } from 'drazil-api'
import { AxiosError } from 'axios'
import { api, getAlbumArtLink } from '~/util/api'

export default Vue.extend({
  asyncData () {
    return api.drazilListAllAlbums().then((res) => {
      return { albums: res.data.albums!, loading: false }
    }).catch((err: AxiosError) => {
      if (err.response?.status === 404) {
        return { notFound: true }
      }
    })
  },
  data () {
    return {
      albums: [] as SpecAllAlbumsReply[],
      loading: true,
      title: 'Albums',
      notFound: false
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    clicked (row: any) {
      this.$router.push({ name: 'artist-artist_id-albums', params: { artist_id: row.id } })
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    },
    showAlbum (artistId: string, id: string) {
      console.log({ artistId }, { id })

      this.$router.push({
        name: 'artist-artist_id-album-album_id-songs',
        params: {
          artist_id: artistId,
          album_id: id
        }
      })
    },
    getAlbumArtLink
  }
})
</script>
