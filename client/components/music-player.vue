<template>
  <v-card class="mt-2 px-2" style="position: sticky; bottom: 8px; padding-bottom: 10px;">
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
        class="mr-1"
        :disabled="playHead - 1 < 0"
        @click="prev"
      >
        <v-icon>
          mdi-skip-previous
        </v-icon>
      </v-btn>
      <v-btn
        elevation="1"
        color="secondary"
        :disabled="playHead + 1 >= queue.length"
        @click="next"
      >
        <v-icon>
          mdi-skip-next
        </v-icon>
      </v-btn>

      <v-chip
        v-if="!$vuetify.breakpoint.mobile"
        label
        class="ml-auto my-auto"
        @wheel.stop.prevent="volumeScroll"
        v-text="currentSong.name"
      />

      <v-chip
        label
        medium
        class="ml-1 my-auto"
      >
        {{ fmtMSS(currentTime) }} / {{ fmtMSS(duration) }}
      </v-chip>

      <v-chip
        v-if="!$vuetify.breakpoint.mobile"
        label
        class="ml-1 my-auto"
        @wheel.stop.prevent="volumeScroll"
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
    <v-snackbar
      v-model="snackbar"
      :top="true"
      :right="true"
      :timeout="2000"
    >
      {{ snackbarText }}

      <template #action="{ attrs }">
        <v-btn
          color="blue"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-card>
</template>

<script lang="ts">
import Vue from 'vue'
import { getSongLink, getAlbumArtLink, FlatSong } from '~/util/api'

function fmtMSS (s: number): string {
  const date = new Date(0)
  date.setSeconds(s)
  return date.toISOString().substr(14, 5)
}

export default Vue.extend({
  data () {
    return {
      queue: [] as FlatSong[],
      currentSong: {} as FlatSong,
      playing: false,
      shuffling: false,
      currentTime: 0,
      duration: 0,
      percentComplete: 0,
      scrobbling: false,
      player: new Audio(),
      playHead: 0,
      snackbar: false,
      snackbarText: ''
    }
  },
  watch: {
    currentSong () {
      this.player.src = getSongLink(
        this.currentSong.artistId,
        this.currentSong.albumId,
        this.currentSong.id!)
      this.player.load()
      this.player.play()

      if ('mediaSession' in navigator) {
        navigator.mediaSession.metadata = new MediaMetadata({
          title: this.currentSong.name,
          artist: this.currentSong.albumName,
          album: this.currentSong.artistName,
          artwork: [
            { src: `${getAlbumArtLink(this.$route.params.artist_id, this.$route.params.album_id)}` }
          ]
        })

        navigator.mediaSession.setActionHandler('previoustrack', this.prev)
        navigator.mediaSession.setActionHandler('nexttrack', this.next)
      }
    }
  },
  mounted () {
    this.createListeners()

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
      this.$nuxt.$on('enqueue', this.enqueue)
    },
    destroyListeners () {
      this.$nuxt.$off('enqueue')
    },
    enqueue (songs: FlatSong[]) {
      if (this.queue.length === 0) {
        this.currentSong = songs[0]
        this.playHead = 0
      }
      this.queue = [...this.queue, ...songs]

      this.snackbarText = `${songs.length} added to queue`
      this.snackbar = true
    },
    next () {
      if (this.playHead + 1 >= this.queue.length) { return }

      this.playHead++
      this.currentSong = this.queue[this.playHead]
    },
    prev () {
      if (this.playHead - 1 < 0) { return }

      this.playHead--
      this.currentSong = this.queue[this.playHead]
    },
    toggle () {
      if (this.player!.paused) {
        this.player!.play()
      } else {
        this.player!.pause()
      }
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
    volumeScroll (e: WheelEvent) {
      let current = this.player.volume

      if (e.deltaY < 0) {
        current += 0.075
      } else {
        current -= 0.075
      }

      if (current > 1 || current < 0) {
        return
      }

      this.player.volume = current
    },
    fmtMSS,
    getSongLink,
    getAlbumArtLink
  }
})
</script>
