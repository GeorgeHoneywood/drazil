<template>
  <div v-if="!notFound">
    <!-- <v-card>
      <v-card-title class="headline">
        <NuxtLink :to="`/artist/${$route.params.artist_id}/albums`">
          {{ artistName }}
        </NuxtLink>  / {{ albumName }}
      </v-card-title>
    </v-card> -->
    <v-breadcrumbs :items="breadcrumbs" />
    <v-data-table
      hide-default-footer
      :headers="headers"
      :items="songs"
      :loading="loading"
      :items-per-page="-1"
      @click:row="clicked"
    />
    <audio ref="player" style="width: 75%; position: fixed; bottom: 0; right: 0;" controls>
      <source :src="currentSong">
    </audio>
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
        loading: false,
        breadcrumbs: [
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
          },
          {
            text: res.data.albumName!,
            disabled: false,
            to: {
              name: 'artist-artist_id-album-album_id-songs',
              params: {
                artist_id: context.route.params.artist_id,
                album_id: context.route.params.album_id
              }
            },
            exact: true
          }
        ]
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
          text: '#',
          value: 'number',
          align: 'left',
          sortable: false,
          width: '1%'
        },
        {
          text: 'Song name',
          value: 'name',
          sortable: false
        }
      ],
      songs: [] as SpecSong[],
      loading: true,
      notFound: false,
      // artistCount: 0
      title: 'Songs',
      artistName: 'Loading...',
      albumName: 'Loading...',
      currentSong: '',
      breadcrumbs: []
    }
  },
  watch: {
    currentSong () {
      // @ts-ignore
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
