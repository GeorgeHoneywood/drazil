<template>
  <!-- <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6"> -->
  <div>
    <v-card>
      <v-card-title class="headline">
        See artists here!
      </v-card-title>
    </v-card>
    <v-data-table
      :headers="headers"
      :items="artists"
      :loading="loading"
      :footer-props="{
        'items-per-page-options': [10, 25, 50]
      }"
      @click:row="clicked"
    />
  </div>
  <!-- :server-items-length="artistCount" -->
  <!-- </v-col>
  </v-row> -->
</template>

<script lang="ts">
import Vue from 'vue'
import { MonkeyApi, SpecArtist } from 'monkey-api'

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
      artists: [] as SpecArtist[],
      loading: true,
      // artistCount: 0
      title: 'Artists'
    }
  },
  mounted () {
    this.updateTitle()
    this.getArtists()
  },
  methods: {
    getArtists () {
      this.loading = true

      const api = new MonkeyApi(undefined, 'http://localhost:8081')

      api.monkeyListArtists().then((res) => {
        if (res.data.artists) {
          this.artists = res.data.artists
        }

        // this.artistCount = res.data.artistCount
      }).catch((err) => {
        console.error(err)
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
