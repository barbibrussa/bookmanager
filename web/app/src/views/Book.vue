<template>
  <v-container>
    <v-card>
      <v-card-title>{{ book.title }} - {{ id }}</v-card-title>
      <v-card-subtitle>{{ book.author }}</v-card-subtitle>
    </v-card>
    <v-card :key="review.ID" v-for="review of reviews">
      <v-card-text> {{ review.review }}</v-card-text>
      <v-rating
        :value = "review.score / 2"
        color="amber"
        dense
        half-increments
        readonly
        size="14"
      ></v-rating>
    </v-card>
  </v-container>
</template>

<script>

import client from '@/api/client';

export default {
  name: 'Book',
  beforeMount() {
    this.id = this.$route.params.id;

    client.get(`/books/${this.id}`).then(
      async (response) => {
        this.book = response.data;
      },
    );

    client.get(`/books/${this.id}/reviews`).then(
      async (response) => {
        this.reviews = response.data;
      },
    );
  },
  data() {
    return {
      id: 0,
      book: {},
      reviews: [],
    };
  },
};
</script>

<style scoped>

</style>
