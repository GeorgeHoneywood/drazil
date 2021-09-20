<template>
  <div v-if="!notFound">
    <v-breadcrumbs :items="breadcrumbs" />
    <div class="mb-2">
      <v-btn
        elevation="1"
        color="primary"
        class="mr-1"
        @click="play"
      >
        <v-icon left>
          mdi-play
        </v-icon>
        Play
      </v-btn>
      <v-btn
        elevation="1"
        color="secondary"
        class="mr-1"
        @click="shuffle"
      >
        <v-icon left>
          mdi-shuffle
        </v-icon>
        Shuffle
      </v-btn>
      <v-img
        class="float-right rounded"
        :src="albumArt"
        width="40px"
      />
      <v-chip
        class="float-right"
        label
        medium
      >
        <v-icon left>
          mdi-volume-high
        </v-icon>
        {{ currentSong.name }}
      </v-chip>
    </div>
    <v-data-table
      class="row-pointer"
      hide-default-footer
      :headers="headers"
      :items="songs"
      :loading="loading"
      :items-per-page="-1"
      @click:row="clicked"
    />
    <audio
      v-show="currentSong.path"
      ref="player"
      style="position: fixed; bottom: 0; right: 0;"
      controls
      preload="auto"
      @ended="next"
    >
      <source :src="`http://localhost:4444/v1/artist/${$route.params.artist_id}/album/${$route.params.album_id}/song/${currentSong.id}/media`">
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
import { SpecSong } from 'monkey-api'
import axios from 'axios'
import { api } from '~/util/api'

export default Vue.extend({
  async asyncData (context) {
    try {
      const res = await api.monkeyListSongs(context.route.params.artist_id, context.route.params.album_id)

      return {
        songs: res.data.songs!,
        artistName: res.data.artistName!,
        albumName: res.data.albumName!,
        albumArt: res.data.albumArt!,
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
      title: 'Songs',
      artistName: 'Loading...',
      albumName: 'Loading...',
      albumArt: '',
      currentSong: {} as SpecSong,
      breadcrumbs: [],
      playing: false,
      shuffling: false
    }
  },
  watch: {
    currentSong () {
      // @ts-ignore
      this.$refs.player!.load()
      // @ts-ignore
      this.$refs.player!.play()
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    clicked (row: SpecSong) {
      this.currentSong = row
    },
    next () {
      if (this.playing) {
        // somewhat confusingly this just plays the next song
        this.currentSong = this.songs[this.currentSong.number!]
      }

      if (this.shuffling) {
        this.currentSong = this.songs[Math.floor(Math.random() * this.songs.length)]
      }
    },
    play () {
      this.playing = true
      this.shuffling = false
      this.currentSong = this.songs[0]
    },
    shuffle () {
      this.shuffling = true
      this.playing = false
      this.currentSong = this.songs[Math.floor(Math.random() * this.songs.length)] // random :p
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    }
  }
})
</script>

<style lang="css" scoped>
.row-pointer >>> tbody tr :hover {
  cursor: pointer;
}
</style>
