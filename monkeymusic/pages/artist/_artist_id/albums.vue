<template>
  <!-- <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6"> -->
  <div v-if="!notFound">
    <v-card>
      <v-card-title class="headline">
        {{ artistName }}
      </v-card-title>
    </v-card>
    <v-data-table
      :headers="headers"
      :items="albums"
      :loading="loading"
      :footer-props="{
        'items-per-page-options': [10, 25, 50]
      }"
      @click:row="clicked"
    >
      <template #[`item.albumArt`]="{ item }">
        <v-img :src="item.albumArt" dark class="rounded" width="40px" />
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
  <!-- :server-items-length="artistCount" -->
  <!-- </v-col>
  </v-row> -->
</template>

<script lang="ts">
import Vue from 'vue'
import { MonkeyApi, SpecAlbum } from 'monkey-api'
import { AxiosError } from 'axios'

export default Vue.extend({
  asyncData (context) {
    const api = new MonkeyApi(undefined, 'http://localhost:8081')

    return api.monkeyListAlbums(context.route.params.artist_id).then((res) => {
      return { albums: res.data.albums!, artistName: res.data.artistName!, loading: false }
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
          // text: 'Art',
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
      albums: [] as SpecAlbum[],
      loading: true,
      notFound: false,
      // artistCount: 0
      title: 'Albums',
      artistName: 'Loading...'
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    clicked (row: any) {
      this.$router.push({ path: `/artist/${this.$route.params.artist_id}/album/${row.id}/songs` })
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    }
  }
})
</script>
