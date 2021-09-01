<template>
  <!-- <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6"> -->
  <div v-if="!notFound">
    <v-card>
      <v-card-title class="headline">
        {{ artistName }} / {{ albumName }}
      </v-card-title>
    </v-card>
    <v-data-table
      hide-default-footer
      :headers="headers"
      :items="songs"
      :loading="loading"
      :items-per-page="-1"
      @click:row="clicked"
    />
    <audio ref="player" style="width: 100%" controls>
      <source :src="currentSong">
      Your browser does not support the audio tag.
    </audio>
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
import axios from 'axios'

export default Vue.extend({
  async asyncData (context) {
    const api = new MonkeyApi(undefined, 'http://localhost:8081')

    try {
      const res = await api.monkeyListSongs(context.route.params.artist_id, context.route.params.album_id)

      return {
        songs: res.data.songs!,
        artistName: res.data.artistName!,
        albumName: res.data.albumName!,
        loading: false
      }
    } catch (err) {
      if (axios.isAxiosError(err)) {
        if (err.response!.status === 404) {
          return { notFound: true }
        }
      }
    }

    // songs: res.data.songs!, loading: false

    // ).catch((err: AxiosError) => {
    //   if (err.response?.status === 404) {
    //     return { notFound: true }
    //   }
    // })
  },
  data () {
    return {
      headers: [
        {
          text: '#',
          value: 'number',
          align: 'left',
          sortable: true,
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
      artistName: 'Loading...',
      albumName: 'Loading...',
      currentSong: ''
    }
  },
  watch: {
    currentSong () {
      this.$refs.player!.load()
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    clicked (row: any) {
      this.currentSong = row.path
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    }
  }
})
</script>
