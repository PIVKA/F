<template>
  <div>
    <li
      class="list-group-item p-3 d-flex justify-content-between align-items-start"
    >
      <div class="ms-2 me-auto">
        <img width="150px" class="avatar" :src="chat.img" alt="chat.title" />
        <div class="fw">
          <div class="fw-bold fs-1">
            {{ chat.title }}
          </div>
          <div class="fw fs-4">
          {{ chat.message }}
          </div>
        </div>
      </div>
      <span class="badge bg-primary rounded-pill">{{ chat.unchecked }}</span>
    </li>
  </div>
</template>

<script>
export default {
  props: {
    chat: {
      type: Object,
    },
  },
};
</script>

<style>
</style>