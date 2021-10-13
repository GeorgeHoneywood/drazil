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
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data () {
    return {
      clipped: true,
      drawer: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'Welcome',
          to: '/'
        },
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
      title: 'Welcome to Monkey!'
    }
  },
  created () {
    this.setTitleListener()
  },
  methods: {
    setTitleListener () {
      this.$nuxt.$on('updateTitle', this.setTitle)
    },
    setTitle (title) {
      this.title = title || 'Monkey'
    }
  }
}
</script>
