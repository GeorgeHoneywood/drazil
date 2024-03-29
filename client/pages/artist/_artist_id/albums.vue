<template>
  <div v-if="!notFound">
    <v-breadcrumbs :items="breadcrumbs" />
    <v-data-table
      class="row-pointer"
      :headers="headers"
      :items="albums"
      :loading="loading"
      :mobile-breakpoint="0"
      :footer-props="{
        'items-per-page-options': [10, 25, 50]
      }"
      @click:row="clicked"
    >
      <template #[`item.albumArt`]="{ item }">
        <v-img :src="getAlbumArtLink($route.params.artist_id, item.id)" dark class="rounded" width="40px" />
      </template>
    </v-data-table>
  </div>
  <div v-else>
    <v-card>
      <v-card-title class="headline">
        Could not find this artist
      </v-card-title>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { SpecAlbum } from 'drazil-api'
import axios from 'axios'
import { api, getAlbumArtLink } from '~/util/api'

export default Vue.extend({
  async asyncData (context) {
    try {
      const res = await api.drazilListAlbums(context.route.params.artist_id)

      return {
        albums: res.data.albums!,
        artistName: res.data.artistName!,
        loading: false,
        breadcrumbs: [
          {
            text: 'Artists',
            disabled: false,
            to: {
              name: 'artists'
            },
            exact: true
          },
          {
            text: res.data.artistName!,
            disabled: false,
            to: {
              name: 'artist-artist_id-albums',
              params: {
                artist_id: context.route.params.artist_id
              }
            },
            exact: true
          }]
      }
    } catch (err) {
      if (axios.isAxiosError(err)) {
        if (err.response!.status === 404) {
          return { notFound: true }
        }
      }
    }
  },
  data () {
    return {
      headers: [
        {
          value: 'albumArt',
          sortable: false,
          width: '1%'
        },
        {
          text: 'Album name',
          value: 'name',
          sortable: true
        },
        {
          text: 'Year',
          value: 'year',
          sortable: true,
          width: '1%'
        },
        {
          text: 'Song count',
          value: 'songCount',
          sortable: true,
          width: '1%'
        }
      ],
      albums: [] as SpecAlbum[],
      loading: true,
      notFound: false,
      title: 'Albums',
      artistName: 'Loading...',
      breadcrumbs: []
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    clicked (row: any) {
      this.$router.push({
        name: 'artist-artist_id-album-album_id-songs',
        params: {
          artist_id: this.$route.params.artist_id,
          album_id: row.id
        }
      })
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    },
    getAlbumArtLink
  }
})
</script>

<style lang="css" scoped>
.row-pointer >>> tbody tr :hover {
  cursor: pointer;
}
</style>
