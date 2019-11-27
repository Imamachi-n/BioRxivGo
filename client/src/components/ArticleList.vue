<template>
  <v-container>
    <div v-for="article in $store.getters['getArticles']" :key="article.Doi">
      <ArticleItem :article="article" />
    </div>
  </v-container>
</template>

<script>
import ArticleItem from "./ArticleItem";
export default {
  name: "ArticleList",
  components: {
    ArticleItem,
  },
  async created() {
    this.$store.dispatch("getArticlesAllAction");
  },
  data: () => ({
    //
  }),
};
</script>
