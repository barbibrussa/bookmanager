<template>
  <div>
    <v-simple-table>
      <template v-slot:default>
        <thead>
        <tr>
          <th class="text-left">
            ID
          </th>
          <th class="text-left">
            Title
          </th>
          <th class="text-left">
            Author
          </th>
        </tr>
        </thead>
        <tbody>
        <tr
          v-for="item in list"
          :key="item.ID"
        >
          <td>{{ item.ID }}</td>
          <td>{{ item.title }}</td>
          <td>{{ item.author }}</td>
        </tr>
        </tbody>
      </template>
    </v-simple-table>
  </div>
</template>

<script>
export default {
  name: 'BookList',
  beforeCreate() {
    fetch('http://localhost:3030/books').then(
      async (response) => {
        this.list = await response.json();
      },
    );
  },
  data() {
    return {
      list: [],
    };
  },
};
</script>

<style scoped>

</style>
