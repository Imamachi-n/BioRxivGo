<template>
  <v-container>
    <p>{{ article.Title }}</p>
  </v-container>
</template>

<script>
export default {
  name: "ArticleItem",
  props: ["article"],
  data: () => ({
    //
  }),
};
</script>
