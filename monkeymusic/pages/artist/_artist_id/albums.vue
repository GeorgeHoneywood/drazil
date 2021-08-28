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
      :items="albums"
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
import { MonkeyApi, SpecAlbum } from 'monkey-api'
import { AxiosError } from 'axios'

export default Vue.extend({
  data () {
    return {
      headers: [
        {
          text: 'ID',
          value: 'id',
          align: 'left',
          sortable: false,
          width: '1%'
        },
        {
          text: 'Artist name',
          value: 'name',
          sortable: true
        }
      ],
      albums: [] as SpecAlbum[],
      loading: true,
      notFound: false,
      // artistCount: 0
      title: 'Albums',
      artistName: 'Loading...'
    }
  },
  mounted () {
    this.updateTitle()
    this.getAlbum()
  },
  methods: {
    getAlbum () {
      this.loading = true

      const api = new MonkeyApi(undefined, 'http://localhost:8081')

      api.monkeyListAlbums(this.$route.params.artist_id).then((res) => {
        this.albums = res.data.albums!
        this.artistName = res.data.artistName!

        // this.artistCount = res.data.artistCount
      }).catch((err: AxiosError) => {
        if (err.response?.status === 404) {
          this.notFound = true
        }
      }).finally(() => { this.loading = false })
    },
    clicked (row: any) {
      console.log('hello', row.id)
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    }
  }
})
</script>
