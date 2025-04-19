<template>
<div>
  <ul
    class="list-group list-group_contact justify-content-between align-items-center list-group-horizontal"
  >
    <li class="list-group-item list-group-item_contact">
      <img :src="contact.img" alt="" /><span>{{ contact.name }}</span>
    </li>
        <li class="list-group-item list-group-item_contact">{{ contact.status }}</li>
    <li class="list-group-item list-group-item_contact">{{ contact.phone }}</li>
    <li class="list-group-item list-group-item_contact">
      <i class="fa fa-comment" style="color: #ffd460"></i>
      <i class="fa fa-phone" style="color: #ffd460"></i>
    </li>
  </ul>
</div>

</template>

<script>
export default {
  props: {
    contact: {
      type: Object,
    },
  },
};
</script>

<style>
</style>