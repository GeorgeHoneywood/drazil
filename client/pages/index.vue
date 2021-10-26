<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card>
        <v-card-text>
          <p>
            drazil is a web-based music player.
          </p>

          <p>
            think plex, but just for music.
          </p>
        </v-card-text>
        <v-card-actions>
          <v-btn
            text
            to="artists"
          >
            Artists
          </v-btn>
          <v-btn
            text
            to="albums"
          >
            Albums
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  data () {
    return {
      title: 'drazil!'
    }
  },
  mounted () {
    this.updateTitle()
  },
  methods: {
    updateTitle () {
      this.$nuxt.$emit('updateTitle', this.title)
    }
  }
})
</script>
