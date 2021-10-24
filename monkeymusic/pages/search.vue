<template>
  <div>
    <v-text-field
      v-model="token"
      placeholder="Search"
      solo
      @keyup="search"
    />

    <div v-if="!notFound">
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
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'
import { api, getAlbumArtLink } from '~/util/api'
import { SpecAllAlbum, SpecSearchReply } from '~/api-client'

export default Vue.extend({
//   async asyncData () {
//     try {
//       const res = await api.monkeySearch('cheese')

  //       return {
  //         albums: res.data.albums!,
  //         artists: res.data.artists!,
  //         songs: res.data.songs!,
  //         loading: false
  //       }
  //     } catch (err) {
  //       if (axios.isAxiosError(err)) {
  //         if (err.response!.status === 404) {
  //           return { notFound: true }
  //         }
  //       }
  //     }
  //   },
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
        }

      ],
      loading: true,
      notFound: false,
      title: 'Search',
      token: '',
      result: {} as SpecSearchReply
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    async search () {
      try {
        const res = await api.monkeySearch(this.token)

        this.result = res.data
      } catch (err) {
        if (axios.isAxiosError(err)) {
          if (err.response!.status === 404) {
            console.log('error')
          }
        }
      }
    },
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
