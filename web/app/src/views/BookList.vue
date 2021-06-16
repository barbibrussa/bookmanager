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
            Título
          </th>
          <th class="text-left">
            Autor
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
          <td>{{ item.title }}</td>
          <td>{{ item.author }}</td>
          <td>
            <v-icon title="Informacion"
              small
              class="mr-2"
              @click="showInfo(item.ID)"
            >
              mdi-alert-circle-outline
            </v-icon>
            <v-dialog
              v-model="dialog"
              persistent
              max-width="600px"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-icon title="Pedir prestado" @click="selectBook(item.ID)"
                  small
                  class="mr-2"
                  v-bind="attrs"
                  v-on="on"
                >
                  mdi-pencil
                </v-icon>
              </template>
              <v-card>
                <v-card-title>
                  <span class="text-h5">Prestar libro</span>
                </v-card-title>
                <v-card-text>
                  <v-container>
                    <v-row>
                      <v-col
                        cols="12"
                        sm="6"
                        md="4"
                      >
                        <v-text-field v-model="firstName"
                          label="Nombre*"
                          required
                        ></v-text-field>
                      </v-col>
                      <v-col
                        cols="12"
                        sm="6"
                        md="4"
                      >
                        <v-text-field v-model="lastName"
                          label="Apellido*"
                          persistent-hint
                          required
                        ></v-text-field>
                      </v-col>
                      <v-col
                        cols="12"
                        sm="6"
                        md="4"
                      >
                        <v-text-field v-model="dni"
                          label="DNI*"
                          persistent-hint
                          required
                        ></v-text-field>
                      </v-col>
                      <v-col
                        cols="12"
                        sm="6"
                        md="4"
                      >
                        <v-text-field v-model="phoneNumber"
                          label="Número de teléfono++*"
                          persistent-hint
                          required
                        ></v-text-field>
                      </v-col>
                    </v-row>
                  </v-container>
                  <small>*campos obligatorios</small>
                </v-card-text>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn
                    color="blue darken-1"
                    text
                    @click="dialog = false"
                  >
                    Cerrar
                  </v-btn>
                  <v-btn
                    color="blue darken-1"
                    text
                    @click="borrowBook()"
                  >
                    Confirmar
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
            <v-dialog
              v-model="dialog"
              persistent
              max-width="290"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-icon title="Borrar libro" @click="selectBook(item.ID)"
                  small v-bind="attrs"
                        v-on="on"
                >
                mdi-delete
                </v-icon>
              </template>
              <v-card>
                <v-card-title class="text-h5">
                  ¿Desea borrar el libro?
                </v-card-title>
                <v-card-text>Confirme para eliminarlo o aprete cancelar</v-card-text>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn
                    color="green darken-1"
                    text
                    @click="dialog = false"
                  >
                    Cancelar
                  </v-btn>
                  <v-btn
                    color="green darken-1"
                    text
                    @click="deleteBook()"
                  >
                    Eliminar
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
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
      dialog: false,
      firstName: '',
      lastName: '',
      dni: '',
      phoneNumber: '',
    };
  },
  methods: {
    showInfo(id) {
      client.get(`/books/${id}`).then((res) => console.log(res.data));
    },
    deleteBook() {
      client.delete(`/books/${this.selected}`)
        .then((res) => console.log(res.data))
        .catch((err) => console.error(err));

      const i = this.list.findIndex((item) => item.id === this.selected);
      this.list.splice(i, 1);
    },
    selectBook(id) {
      this.selected = id;
    },
    borrowBook() {
      client.post(`/books/${this.selected}/borrow`, {
        first_name: this.firstName,
        last_name: this.lastName,
        dni: this.dni,
        phone_number: this.phoneNumber,
      }).then((res) => console.log(res));
      this.dialog = false;
    },
  },
};
</script>

<style scoped>

</style>
