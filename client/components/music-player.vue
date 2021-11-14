<template>
  <div class="mt-2 white" style="position: sticky; bottom: 0px; padding-bottom: 10px;">
    <v-slider
      v-model="percentComplete"
      :hide-details="true"
      :disabled="currentSong.id === undefined"
      max="100"
      min="0"
      step="0.1"
      @end="scrobble"
      @start="scrobbleStart"
    />
    <div class="mt-2 d-flex">
      <v-btn
        :disabled="false"
        elevation="1"
        color="secondary"
        class="mr-1"
        @click="toggle"
      >
        <v-icon v-if="!player.paused">
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

      <v-chip
        label
        medium
        class="ml-auto"
      >
        {{ fmtMSS(currentTime) }} / {{ fmtMSS(duration) }}
      </v-chip>

      <v-chip
        v-if="!$vuetify.breakpoint.mobile"
        label
        class="ml-1"
      >
        <v-slider
          v-model="player.volume"
          style="width: 150px"
          dense
          max="1"
          min="0"
          step="0.01"
          class="my-auto"
          prepend-icon="mdi-volume-high"
        />
      </v-chip>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { SpecSong } from 'drazil-api'
import { getSongLink, getAlbumArtLink } from '~/util/api'

function fmtMSS (s: number): string {
  const date = new Date(0)
  date.setSeconds(s)
  return date.toISOString().substr(14, 5)
}

export default Vue.extend({
  data () {
    return {
      songs: [] as SpecSong[],
      currentSong: {} as SpecSong,
      playing: false,
      shuffling: false,
      currentTime: 0,
      duration: 0,
      percentComplete: 0,
      scrobbling: false,
      player: {} as HTMLAudioElement
    }
  },
  watch: {
    currentSong () {
      this.player!.src = getSongLink(this.$route.params.artist_id,
        this.$route.params.album_id, this.currentSong.id!)
      this.player!.load()
      this.player!.play()

      if ('mediaSession' in navigator) {
        navigator.mediaSession.metadata = new MediaMetadata({
          title: this.currentSong.name,
          //   artist: this.artistName,
          //   album: this.albumName,
          artwork: [
            { src: `${getAlbumArtLink(this.$route.params.artist_id, this.$route.params.album_id)}` }
          ]
        })

        // navigator.mediaSession.setActionHandler('previoustrack', function () { /* Code excerpted. */ })
        navigator.mediaSession.setActionHandler('nexttrack', this.next)
      }
    }
  },
  mounted () {
    this.createListeners()
    this.player = new Audio()

    this.player.addEventListener('ended', this.next)
    this.player.addEventListener('timeupdate', this.timeUpdate)
    this.player.addEventListener('durationchange', this.durationUpdate)
    this.player.addEventListener('loadedmetadata', this.durationUpdate)

    if (this.$vuetify.breakpoint.mobile) {
      this.player.volume = 1
    } else {
      this.player.volume = 0.15
    }
  },
  beforeDestroy () {
    this.destroyListeners()

    this.player.pause()
    this.player = new Audio()
  },
  methods: {
    createListeners () {
      this.$nuxt.$on('songClicked', this.songClicked)
      this.$nuxt.$on('queue', this.queue)
    },
    destroyListeners () {
      this.$nuxt.$off('songClicked')
      this.$nuxt.$off('queue')
    },
    songClicked (song: SpecSong) {
      this.currentSong = song
    },
    queue (songs: SpecSong[]) {
      this.songs = songs
      console.log(this.songs)
    },
    next () {
      if (this.shuffling) {
        this.currentSong = this.songs[Math.floor(Math.random() * this.songs.length)]
        return
      }

      const nextSong = this.songs[this.currentSong.number!]

      if (!nextSong) {
        this.playing = false
        this.currentSong = {}
        return
      }

      this.currentSong = nextSong
    },
    play () {
      this.playing = true
      this.shuffling = false
      this.currentSong = this.songs[0]
    },
    toggle () {
      if (this.player!.paused) {
        this.player!.play()
      } else {
        this.player!.pause()
      }
    },
    shuffle () {
      this.shuffling = true
      this.playing = false
      this.currentSong = this.songs[Math.floor(Math.random() * this.songs.length)] // random :p
    },
    itemClass (item: any) {
      return this.currentSong.id === item.id ? 'primary lighten-1 white--text rounded-pill' : ''
    },
    timeUpdate () {
      if (this.scrobbling) {
        return
      }

      this.currentTime = this.player!.currentTime
      this.percentComplete = (this.currentTime / this.duration) * 100
    },
    durationUpdate () {
      this.duration = this.player!.duration
    },
    scrobble () {
      this.player!.currentTime = (this.percentComplete / 100) * this.player!.duration
      this.scrobbling = false
    },
    scrobbleStart () {
      this.scrobbling = true
    },
    fmtMSS,
    getSongLink,
    getAlbumArtLink
  }
})
</script>
