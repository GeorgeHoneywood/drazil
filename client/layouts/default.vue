<template>
  <v-app>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
      color="primary"
    >
      <v-app-bar-nav-icon style="color: white" @click.stop="drawer = !drawer" />
      <v-toolbar-title class="white--text" v-text="title" />
      <v-spacer />
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
        <MusicPlayer />
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data () {
    return {
      clipped: true,
      drawer: false,
      items: [
        {
          icon: 'mdi-account-multiple',
          title: 'Artists',
          to: '/artists'
        },
        {
          icon: 'mdi-disc',
          title: 'Albums',
          to: '/albums'
        }
      ],
      miniVariant: false,
      title: 'welcome to drazil!'
    }
  },
  created () {
    this.setListeners()
  },
  mounted () {
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
      this.setDarkTheme(e.matches)
    })

    this.setDarkTheme(window.matchMedia('(prefers-color-scheme: dark)').matches)
  },
  methods: {
    setListeners () {
      this.$nuxt.$on('updateTitle', this.setTitle)
    },
    setTitle (title: string) {
      this.title = title || 'Monkey'
    },
    setDarkTheme (theme : boolean) {
      this.$vuetify.theme.dark = theme
    }
  }

})
</script>
