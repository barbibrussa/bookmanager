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
            <v-icon
              title="Devolver"
              @click="selectBook(item.book_id); showReturnDialog();"
              small
              class="mr-2"
            >
              mdi-pencil
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
      <v-dialog
        v-model="returnDialog"
        persistent
        max-width="400"
      >
        <v-card>
          <v-card-title class="text-h5">
            ¿Desea devolver el libro?
          </v-card-title>
          <v-card-text>Confirme para la devolución o aprete cancelar</v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="green darken-1"
              text
              @click="returnDialog = false"
            >
              Cancelar
            </v-btn>
            <v-btn
              color="green darken-1"
              text
              @click="returnBook()"
            >
              Devolver
            </v-btn>
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
      returnDialog: false,
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
    selectBook(id) {
      this.selected = id;
    },
    showReturnDialog() {
      this.returnDialog = true;
    },
    returnBook() {
      client.put(`/books/${this.selected}/return`)
        .then((res) => console.log(res.data))
        .catch((err) => console.error(err));

      const i = this.list.findIndex((item) => item.ID === this.selected);
      this.list.splice(i, 1);

      this.returnDialog = false;
    },
  },
};
</script>

<style scoped>

</style>
