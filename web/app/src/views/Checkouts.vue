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
            ID Libro
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
          <td>{{ item.book_id }}</td>
        </tr>
        </tbody>
      </template>
    </v-simple-table>
  </div>
</template>

<script>
export default {
  name: 'Checkouts',
  beforeCreate() {
    fetch('http://localhost:3030/checkouts').then(
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
