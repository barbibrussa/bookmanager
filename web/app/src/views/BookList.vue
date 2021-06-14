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
          <th class="text-left">
            Actions
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
          <td>
            <v-icon
              small
              class="mr-2"
              @click="showInfo(item.ID)"
            >
              mdi-alert-circle-outline
            </v-icon>
            <v-icon
              small
              class="mr-2"
              @click="borrowBook(item.ID)"
            >
              mdi-pencil
            </v-icon>
            <v-icon
              small
              @click="deleteBook(item.ID)"
            >
              mdi-delete
            </v-icon>
          </td>
        </tr>
        </tbody>
      </template>
    </v-simple-table>
  </div>
</template>

<script>
import client from '@/api/client';

export default {
  name: 'BookList',
  beforeCreate() {
    client.get('/books').then(
      async (response) => {
        this.list = response.data;
      },
    );
  },
  data() {
    return {
      list: [],
      book: {},
    };
  },
  methods: {
    showInfo(id) {
      client.get(`/books/${id}`).then((res) => console.log(res.data));
    },
    deleteBook(id) {
      client.delete(`/books/${id}`)
        .then((res) => console.log(res.data))
        .catch((err) => console.error(err));

      const i = this.list.findIndex((item) => item.id === id);
      this.list.splice(i, 1);
    },
    borrowBook(id) {
      client.post(`/books/${id}/borrow`, {
        first_name: prompt('Ingrese el nombre'),
        last_name: prompt('Ingrese el apellido'),
        dni: prompt('Ingrese el DNI'),
        phone_number: prompt('Ingrese el numero de telefono'),
      }).then((res) => console.log(res));
    },
  },
};
</script>

<style scoped>

</style>
