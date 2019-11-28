<template>
  <v-container>
    <v-card class="mx-auto" outlined>
      <v-container>
        <!-- Title -->
        <div class="headline mb-1">{{ article.Title }}</div>

        <!-- Author -->
        <v-list-item-subtitle>{{ article.Author }}</v-list-item-subtitle>

        <!-- Published -->
        <v-list-item-subtitle>{{ article.Published }}</v-list-item-subtitle>

        <!-- DOI -->
        <v-list-item-subtitle>
          <a :href="`https://doi.org/${article.Doi}`" target="_blank">doi: {{ article.Doi }}</a>
        </v-list-item-subtitle>

        <!-- Description -->
        <v-card-text>{{ article.Description }}</v-card-text>

        <v-card-actions>
          <v-btn
            color="green darken-4"
            class="pt-0 pb-0"
            :href="article.Link"
            text
            target="_blank"
          >Link</v-btn>
        </v-card-actions>
      </v-container>
    </v-card>
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
