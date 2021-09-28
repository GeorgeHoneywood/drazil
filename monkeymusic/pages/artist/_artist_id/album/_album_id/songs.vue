<template>
  <div v-if="!notFound">
    <v-breadcrumbs :items="breadcrumbs" />
    <div class="d-flex mb-2">
      <v-btn
        elevation="1"
        color="primary"
        class="mr-2 ml-4"
        @click="play"
      >
        <v-icon>
          mdi-playlist-plus
        </v-icon>
      </v-btn>
      <v-btn
        elevation="1"
        color="secondary"
        class="mr-auto"
        @click="shuffle"
      >
        <v-icon>
          mdi-shuffle
        </v-icon>
      </v-btn>

      <v-chip
        v-if="currentSong.name"
        label
        medium
      >
        <v-icon left>
          mdi-volume-high
        </v-icon>
        {{ currentSong.name }}
      </v-chip>
      <v-img
        class="ml-2 rounded"
        :src="`http://localhost:4444/v1/artist/${$route.params.artist_id}/album/${$route.params.album_id}/art`"
        max-width="40px"
        max-height="40px"
      />
    </div>
    <v-data-table
      class="row-pointer"
      hide-default-footer
      :headers="headers"
      :items="songs"
      :loading="loading"
      :items-per-page="-1"
      :item-class="itemClass"
      @click:row="clicked"
    >
      <template #[`item.playing`]="{ item }">
        <v-icon
          v-if="currentSong.id == item.id"
          style="color: white"
        >
          mdi-play
        </v-icon>
      </template>
    </v-data-table>
    <audio
      v-show="currentSong.id"
      ref="player"
      style="position: fixed; bottom: 0; right: 0; display: block"

      preload="auto"
      @ended="next"
      @timeupdate="timeUpdate"
      @durationchange="durationUpdate"
    >
      <source v-if="currentSong.id" :src="`http://localhost:4444/v1/artist/${$route.params.artist_id}/album/${$route.params.album_id}/song/${currentSong.id}/media`">
    </audio>
    <div class="d-flex mt-2" style="position: sticky; bottom: 10px">
      <v-btn
        v-if="$refs.player"
        elevation="1"
        color="secondary"
        class="mr-1"
        @click="toggle"
      >
        <v-icon v-if="playState">
          mdi-pause
        </v-icon>
        <v-icon v-else>
          mdi-play
        </v-icon>
      </v-btn>
      <v-btn
        elevation="1"
        color="secondary"
        class=""
        @click="next"
      >
        <v-icon>
          mdi-skip-next
        </v-icon>
      </v-btn>

      <v-slider
        v-model="percentComplete"
        :hide-details="true"
        :disabled="currentSong.id === undefined"
        max="100"
        min="0"
        step="0.1"
        class="mx-2 px-2 white"
        @end="scrobble"
      />
    </div>
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
          text: '',
          value: 'playing',
          sortable: false,
          width: '1%'
        },
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
      currentSong: {} as SpecSong,
      breadcrumbs: [],
      playing: false,
      shuffling: false,
      playState: true,
      currentTime: 0,
      duration: 0,
      percentComplete: 0
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
      this.playState = true
    },
    next () {
      if (this.shuffling) {
        this.currentSong = this.songs[Math.floor(Math.random() * this.songs.length)]
        return
      }

      this.currentSong = this.songs[this.currentSong.number!]
    },
    play () {
      this.playState = true
      this.playing = true
      this.shuffling = false
      this.currentSong = this.songs[0]
    },
    toggle () {
      // @ts-ignore
      if (this.$refs.player!.paused) {
        // @ts-ignore
         this.$refs.player!.play()
         this.playState = true
      } else {
        // @ts-ignore
        this.$refs.player!.pause()
        this.playState = false
      }
    },
    shuffle () {
      this.playState = true
      this.shuffling = true
      this.playing = false
      this.currentSong = this.songs[Math.floor(Math.random() * this.songs.length)] // random :p
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    },
    itemClass (item: any) {
      return this.currentSong.id === item.id ? 'primary lighten-1 white--text rounded-pill' : ''
    },
    timeUpdate () {
      // @ts-ignore
      this.currentTime = this.$refs.player.currentTime
      this.percentComplete = (this.currentTime / this.duration) * 100
    },
    durationUpdate () {
      // @ts-ignore
      this.duration = this.$refs.player.duration
    },
    scrobble () {
      // @ts-ignore
      this.$refs.player.currentTime = (this.percentComplete / 100) * this.$refs.player.duration
    }

  }
})
</script>

<style lang="css" scoped>
.row-pointer >>> tbody tr :hover {
  cursor: pointer;
}
</style>
