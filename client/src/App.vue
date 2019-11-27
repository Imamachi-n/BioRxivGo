<template>
  <v-app>
    <v-app-bar app color="teal darken-4" dark>
      <v-btn icon to="/">
        <v-icon>mdi-home</v-icon>
      </v-btn>
      <v-toolbar-title>
        <v-icon class="mb-1 mr-1">fas fa-dna</v-icon>
        <b>BioRxivGo:</b>
        The Curated Preprint Server for Genomics & BioInformatics
      </v-toolbar-title>
    </v-app-bar>

    <v-content>
      <ArticleList />
    </v-content>
  </v-app>
</template>

<script>
import ArticleList from "./components/ArticleList";

export default {
  name: "App",

  components: {
    ArticleList,
  },

  data: () => ({
    //
  }),
};
</script>
