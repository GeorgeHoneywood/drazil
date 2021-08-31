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
      :items-per-page="-1"
      :footer-props="{
        'items-per-page-options': [10, 25, 50, -1]
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
import { AxiosError } from 'axios'

export default Vue.extend({
  asyncData () {
    const api = new MonkeyApi(undefined, 'http://localhost:8081')

    return api.monkeyListArtists().then((res) => {
      return { artists: res.data.artists!, loading: false }
    }).catch((err: AxiosError) => {
      if (err.response?.status === 404) {
        return { notFound: true }
      }
    })
  },
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
  },
  methods: {
    clicked (row: any) {
      this.$router.push({ path: `/artist/${row.id}/albums` })
    },
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    }
  }
})
</script>
