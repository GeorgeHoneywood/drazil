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
        :loading="loading"
        tile
        @click="showAlbum(album.artistId, album.id)"
      >
        <v-img
          :src="getAlbumArtLink(album.artistId, album.id)"
          class="align-end"
          gradient="to bottom, rgba(0,0,0,0), rgba(0,0,0,.5)"
          height="100%"
          aspect-ratio="1"
        >
          <v-card-title
            class="white--text text-subtitle-1"
            style="word-break: normal"
            v-text="album.name"
          />
        </v-img>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from 'vue'
import { SpecAllAlbum } from 'drazil-api'
import axios from 'axios'
import { api, getAlbumArtLink } from '~/util/api'

export default Vue.extend({
  data () {
    return {
      albums: Array.from({ length: 25 }, () => { return {} }) as SpecAllAlbum[],
      loading: true,
      title: 'Albums',
      notFound: false
    }
  },
  mounted () {
    this.loadAlbums()
    this.updateTitle()
  },

  methods: {
    async loadAlbums () {
      try {
        const res = await api.drazilListAllAlbums()
        this.albums = res.data.albums!
        this.loading = false
      } catch (err) {
        if (axios.isAxiosError(err)) {
          if (err.response?.status === 404) {
            this.notFound = true
          }
        }
      }
    },
    clicked (row: any) {
      this.$router.push({ name: 'artist-artist_id-albums', params: { artist_id: row.id } })
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    },
    showAlbum (artistId: string, id: string) {
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
