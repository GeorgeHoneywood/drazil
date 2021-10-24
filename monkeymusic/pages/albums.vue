<template>
  <v-row dense>
    <v-col
      v-for="album in albums"
      :key="album.id"
    >
      <v-card
        width="200"
        height="200"
        tile
        @click="showAlbum(album.artistId, album.id)"
      >
        <v-img
          :src="getAlbumArtLink(album.artistId, album.id)"
          class="white--text align-end"
          gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
        >
          <v-card-title v-text="album.name" />
        </v-img>

        <!-- <v-card-actions>
          <v-spacer />

          <v-btn icon>
            <v-icon>mdi-heart</v-icon>
          </v-btn>

          <v-btn icon>
            <v-icon>mdi-bookmark</v-icon>
          </v-btn>

          <v-btn icon>
            <v-icon>mdi-share-variant</v-icon>
          </v-btn>
        </v-card-actions> -->
      </v-card>
    </v-col>
  </v-row>
  <!-- <p>{{ albums }}</p> -->
</template>

<script lang="ts">
import Vue from 'vue'
import { SpecAllAlbumsReply } from 'monkey-api'
import { AxiosError } from 'axios'
import { api, getAlbumArtLink } from '~/util/api'

export default Vue.extend({
  asyncData () {
    return api.monkeyListAllAlbums().then((res) => {
      return { albums: res.data.albums!, loading: false }
    }).catch((err: AxiosError) => {
      if (err.response?.status === 404) {
        return { notFound: true }
      }
    })
  },
  data () {
    return {
      albums: [] as SpecAllAlbumsReply[],
      loading: true,
      title: 'Albums',
      notFound: false
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    clicked (row: any) {
      this.$router.push({ name: 'artist-artist_id-albums', params: { artist_id: row.id } })
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    },
    showAlbum (artistId: string, id: string) {
      console.log({ artistId }, { id })

      this.$router.push({
        name: 'artist-artist_id-album-album_id-songs',
        params: {
          artist_id: artistId,
          album_id: id
        }
      })
    },
    getAlbumArtLink
  }
})
</script>
