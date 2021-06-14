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
            First name
          </th>
          <th class="text-left">
            Last name
          </th>
          <th class="text-left">
            DNI
          </th>
          <th class="text-left">
            Phone number
          </th>
          <th class="text-left">
            Borrowed at
          </th>
          <th class="text-left">
            Deadline
          </th>
          <th class="text-left">
            Book ID
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
