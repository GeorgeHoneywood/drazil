<template>
  <div>
    <v-breadcrumbs :items="breadcrumbs" />
    <v-data-table
      class="row-pointer"
      :mobile-breakpoint="0"
      :headers="headers"
      :items="artists"
      :loading="loading"
      :items-per-page="-1"
      :show-expand="true"
      :single-expand="true"
      :footer-props="{
        'items-per-page-options': [10, 25, 50, -1]
      }"
      @click:row="clicked"
      @item-expanded="expanded"
    >
      <template #expanded-item="{ headers }">
        <td
          :colspan="headers.length"
          class="pa-3"
        >
          <v-data-table
            hide-default-footer
            :mobile-breakpoint="0"
            :items="artistAlbums"
            :loading="loadingArtistAlbums"
            :headers="[{
                         value: 'albumArt',
                         sortable: false,
                         width: '1%'
                       },
                       {
                         text: 'Album name',
                         value: 'name'
                       }]"
            @click:row="showAlbum"
          >
            <template #[`item.albumArt`]="{ item }" style="width: 1%">
              <v-img
                :src="getAlbumArtLink(expandedArtistID, item.id)"
                dark
                class="rounded"
                width="32px"
              />
            </template>
          </v-data-table>
        </td>
      </template>
    </v-data-table>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { SpecArtist, SpecAlbum } from 'monkey-api'
import { AxiosError } from 'axios'
import { api, getAlbumArtLink } from '~/util/api'

export default Vue.extend({
  asyncData () {
    return api.monkeyListArtists().then((res) => {
      return { artists: res.data.artists!, loading: false }
    }).catch((err: AxiosError) => {
      if (err.response?.status === 404) {
        return { notFound: true }
      }
    })
  },
  data () {
    return {
      headers: [
        {
          text: 'Artist name',
          value: 'name',
          sortable: true
        }
      ],
      artists: [] as SpecArtist[],
      artistAlbums: [] as SpecAlbum[],
      loadingArtistAlbums: true,
      expandedArtistID: 0,
      loading: true,
      title: 'Artists',
      breadcrumbs: [
        {
          text: 'Artists',
          disabled: false,
          to: {
            name: 'artists'
          },
          exact: true
        }]
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
    async expanded ({ item, value } : {item: any, value: boolean}) {
      if (!value) {
        return
      }
      this.artistAlbums = []
      this.expandedArtistID = item.id
      this.loadingArtistAlbums = true

      this.artistAlbums = (await api.monkeyListAlbums(item.id)).data.albums!
      this.loadingArtistAlbums = false
    },
    showAlbum (row: any) {
      this.$router.push({
        name: 'artist-artist_id-album-album_id-songs',
        params: {
          artist_id: this.expandedArtistID.toString(),
          album_id: row.id
        }
      })
    },
    getAlbumArtLink
  }
})
</script>

<style lang="css">
.v-data-table__expanded.v-data-table__expanded__content {
  box-shadow: none !important;
}
</style>

<style lang="css" scoped>
.row-pointer >>> tbody tr :hover {
  cursor: pointer;
}
</style>
