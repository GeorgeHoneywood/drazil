<template>
  <div v-if="!notFound">
    <v-breadcrumbs compact :items="breadcrumbs" />
    <div class="d-flex mb-2">
      <v-btn
        elevation="1"
        color="primary"
        class="mr-2"
        icon
      >
        <!--  @click="play" -->
        <v-icon>
          mdi-playlist-plus
        </v-icon>
      </v-btn>
      <v-btn
        elevation="1"
        color="secondary"
        class="mr-auto"
        icon
      >
        <!--  @click="shuffle" -->
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
          <v-icon style="color: white">
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
              <v-card-title class="text-h6 primary white--text">
                <span
                  class="text-truncate d-inline-block"
                >{{ item.name }}</span>
                <v-spacer />
                <v-btn

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

export default Vue.extend({
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
      breadcrumbs: [] as any[]

    }
  },
  watch: {
    title: {
      handler () {
        this.$nuxt.$emit('updateTitle', this.title)
      },
      immediate: true
    }
  },
  beforeMount () {
    this.loadSongs()
  },
  methods: {
    async loadSongs () {
      try {
        const res = await api.drazilListSongs(this.$route.params.artist_id, this.$route.params.album_id)

        this.songs = res.data.songs!
        this.artistName = res.data.artistName!
        this.albumName = res.data.albumName!
        this.loading = false
        this.title = `${res.data.artistName} - ${res.data.albumName}`
        this.breadcrumbs = [
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
                artist_id: this.$route.params.artist_id
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
                artist_id: this.$route.params.artist_id,
                album_id: this.$route.params.album_id
              }
            },
            exact: true
          }
        ]

        this.$nuxt.$emit('queue', this.songs)
      } catch (err) {
        if (axios.isAxiosError(err)) {
          if (err.response!.status === 404) {
            return { notFound: true }
          }
        }
      }
    },
    clicked (clicked : SpecSong) {
      this.$nuxt.$emit('songClicked', clicked)
      this.currentSong = clicked
    },
    itemClass (item: any) {
      return this.currentSong.id === item.id ? 'primary lighten-1 white--text rounded-pill' : ''
    },
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
