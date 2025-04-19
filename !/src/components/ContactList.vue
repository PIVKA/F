<template>
  <div class="contact-page container">
      <Contact
        v-for="contact in contacts"
        :key="contact.phone"
        :contact="contact"
      />
  </div>
</template>

<script>
import Contact from '../components/Contact'
export default {
    props:{
        contacts: {
            type: Array,
            default: () => []
        }
    },
    components: {
        Contact
    }
}
</script>

<style>

</style>