<template>
  <!-- <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6"> -->
  <div>
    <v-card>
      <v-card-title class="headline">
        Collection
      </v-card-title>
    </v-card>
    <v-data-table
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
        >
          <v-data-table
            hide-default-footer
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
              <v-img :src="item.albumArt" dark class="rounded" width="32px" />
            </template>
            <!-- <thead style="display: none">
                <tr>
                  <th style="width: 1%" />
                  <th />
                </tr>
              </thead> -->
            <!-- <tbody>
              <tr v-for="album in artistAlbums" :key="album.name" @click="showAlbum">
                <td style="width: 1%">

                </td>
                <td>{{ album.name }}</td>
              </tr>
            </tbody> -->
          </v-data-table>
        </td>
      </template>
    </v-data-table>
  </div>
</template>

  <!-- :server-items-length="artistCount" -->
  <!-- </v-col>
  </v-row> -->

<script lang="ts">
import Vue from 'vue'
import { MonkeyApi, SpecArtist } from 'monkey-api'
import { AxiosError } from 'axios'
import { SpecAlbum } from '~/api-client'

export default Vue.extend({
  asyncData () {
    const api = new MonkeyApi(undefined, 'http://localhost:8081')

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
      // artistCount: 0
      title: 'Artists'
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    clicked (row: any) {
      this.$router.push({ path: `/artist/${row.id}/albums` })
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

      const api = new MonkeyApi(undefined, 'http://localhost:8081')

      this.artistAlbums = (await api.monkeyListAlbums(item.id)).data.albums!
      this.loadingArtistAlbums = false
    },
    showAlbum (row: any) {
      this.$router.push({ path: `/artist/${this.expandedArtistID}/album/${row.id}/songs` })
    }
  }
})
</script>
