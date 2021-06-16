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
            Nombre
          </th>
          <th class="text-left">
            Apellido
          </th>
          <th class="text-left">
            DNI
          </th>
          <th class="text-left">
            Número de telefono
          </th>
          <th class="text-left">
            Préstamo
          </th>
          <th class="text-left">
            Fecha límite de devolución
          </th>
          <th class="text-left">
            Acciones
          </th>
        </tr>
        </thead>
        <tbody>
        <tr
          v-for="item in list"
          :key="item.ID"
        >
          <td>{{ item.ID }}</td>
          <td>{{ item.first_name }}</td>
          <td>{{ item.last_name }}</td>
          <td>{{ item.dni }}</td>
          <td>{{ item.phone_number }}</td>
          <td>{{ item.borrowed_at }}</td>
          <td>{{ item.deadline }}</td>
          <td>
            <v-icon title="Informacion"
                    small
                    class="mr-2"
                    @click="showInfo(item.book_id)"
            >
              mdi-alert-circle-outline
            </v-icon>
          </td>
        </tr>
        </tbody>
      </template>
    </v-simple-table>
    <div>
      <v-dialog
        v-model="infoDialog"
        transition="dialog-top-transition"
        max-width="600"
      >
        <v-card>
          <v-toolbar
            color="primary"
            dark
          >Información</v-toolbar>
          <v-card-title>{{ book.title }}</v-card-title>
          <v-card-text>
            <div class="my-1 text-subtitle-">Autor: {{ book.author }}</div>
            <div>ID: {{ book.ID }}</div>
          </v-card-text>
          <v-card-actions class="justify-end">
            <v-btn
              text
              @click="infoDialog = false"
            >Close</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
import client from '@/api/client';

export default {
  name: 'Checkouts',
  beforeCreate() {
    client.get('/checkouts').then(
      async (response) => {
        this.list = response.data;
      },
    );
  },
  data() {
    return {
      list: [],
      book: {},
      infoDialog: false,
    };
  },
  methods: {
    showInfo(id) {
      client.get(`/books/${id}`)
        .then((res) => {
          this.book = res.data;
          this.infoDialog = true;
        });
    },
  },
};
</script>

<style scoped>

</style>
