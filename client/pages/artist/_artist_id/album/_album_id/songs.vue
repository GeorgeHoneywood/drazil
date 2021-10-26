<template>
  <div v-if="!notFound">
    <v-breadcrumbs compact :items="breadcrumbs" />
    <div class="d-flex mb-2">
      <v-btn
        elevation="1"
        color="primary"
        class="mr-2"
        icon
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
        icon
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
        class="ml-2"
      >
        <v-icon left>
          mdi-volume-high
        </v-icon>
        {{ currentSong.name }}
      </v-chip>
      <v-img
        class="ml-2 rounded"
        :src="getAlbumArtLink($route.params.artist_id, $route.params.album_id)"
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
      mobile-breakpoint="0"
      @click:row="clicked"
    >
      <template #[`item.playing`]="{ item }">
        <template
          v-if="currentSong.id == item.id"
        >
          <v-icon v-if="player.paused" style="color: white">
            mdi-pause
          </v-icon>
          <v-icon v-else style="color: white">
            mdi-play
          </v-icon>
        </template>
      </template>
      <template #[`item.lyrics`]="{ item }">
        <v-dialog
          v-if="item.lyrics"
          max-width="500"
        >
          <template #activator="{ on, attrs }">
            <v-btn
              v-bind="attrs"
              icon
              v-on="on"
            >
              <v-icon>
                mdi-view-headline
              </v-icon>
            </v-btn>
          </template>
          <template #default="dialog">
            <v-card>
              <v-card-title class="text-h6 primary lighten-2 white--text">
                <span>{{ albumName }} - {{ item.name }}</span>
                <v-spacer />
                <v-btn
                  class="ml-auto"
                  icon
                  @click="dialog.value = false"
                >
                  <v-icon
                    style="color: white"
                  >
                    mdi-close
                  </v-icon>
                </v-btn>
              </v-card-title>

              <v-card-text
                class="mt-4"
                style="white-space: pre-line"
                v-text="item.lyrics"
              />
            </v-card>
          </template>
        </v-dialog>
      </template>
    </v-data-table>
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
          class="ml-1"
          @click="toggleVolumeSlider"
        >
          <v-slider
            v-if="volumeSlider"
            dense
            max="1"
            min="0"
            step="0.01"
            vertical
            style="height: 50px; z-index: 999; width: 14px; position: fixed"
          />
          <v-icon v-else>
            mdi-volume-high
          </v-icon>
        </v-chip>
      </div>
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
import { SpecSong } from 'drazil-api'
import axios from 'axios'
import { api, getSongLink, getAlbumArtLink } from '~/util/api'

function fmtMSS (s: number): string {
  return ((s - (s %= 60)) / 60 + (s > 9 ? ':' : ':0') + s).split('.')[0]
}

export default Vue.extend({
  async asyncData (context) {
    try {
      const res = await api.drazilListSongs(context.route.params.artist_id, context.route.params.album_id)

      return {
        songs: res.data.songs!,
        artistName: res.data.artistName!,
        albumName: res.data.albumName!,
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
        },
        {
          text: 'Lyrics',
          value: 'lyrics',
          sortable: false,
          align: 'right'
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
      currentTime: 0,
      duration: 0,
      percentComplete: 0,
      scrobbling: false,
      volumeSlider: false,
      player: {} as HTMLAudioElement
    }
  },
  watch: {
    currentSong () {
      this.player!.src = getSongLink(this.$route.params.artist_id,
        this.$route.params.album_id, this.currentSong.id!)
      this.player!.volume = 0.15
      this.player!.load()
      this.player!.play()

      if ('mediaSession' in navigator) {
        navigator.mediaSession.metadata = new MediaMetadata({
          title: this.currentSong.name,
          artist: this.artistName,
          album: this.albumName,
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
    this.player = new Audio()

    this.player?.addEventListener('ended', this.next)
    this.player?.addEventListener('timeupdate', this.timeUpdate)
    this.player?.addEventListener('durationchange', this.durationUpdate)
    this.player?.addEventListener('loadedmetadata', this.durationUpdate)

    this.updateTitle()
  },
  beforeDestroy () {
    // console.log('hello')
    this.player?.pause()
    this.player = new Audio()
  },
  methods: {
    clicked (row: SpecSong) {
      this.currentSong = row
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
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
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
    toggleVolumeSlider () {
      this.volumeSlider = !this.volumeSlider
    },
    fmtMSS,
    getSongLink,
    getAlbumArtLink
  }
})
</script>

<style lang="css" scoped>
.row-pointer >>> tbody tr :hover {
  cursor: pointer;
}
</style>
