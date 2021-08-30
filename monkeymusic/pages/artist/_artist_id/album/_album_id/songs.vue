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
      :items="songs"
      :loading="loading"
      :footer-props="{
        'items-per-page-options': [10, 25, 50]
      }"
      @click:row="clicked"
    />
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
import { MonkeyApi, SpecSong } from 'monkey-api'
import { AxiosError } from 'axios'

export default Vue.extend({
  asyncData (context) {
    const api = new MonkeyApi(undefined, 'http://localhost:8081')

    return api.monkeyListSongs(context.route.params.artist_id, context.route.params.album_id).then((res) => {
      return { songs: res.data.songs!, loading: false }
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
          text: 'Track number',
          value: 'number',
          align: 'left',
          sortable: false,
          width: '1%'
        },
        {
          text: 'Song name',
          value: 'name',
          sortable: true
        }
      ],
      songs: [] as SpecSong[],
      loading: true,
      notFound: false,
      // artistCount: 0
      title: 'Songs',
      artistName: 'Loading...'
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    clicked (row: any) {
      console.log(row)
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    }
  }
})
</script>
