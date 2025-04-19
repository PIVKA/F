<template>
  <div class="contacts-page container">
    <ContactList :contacts="contacts"/>
  </div>
</template>

<script>
import ContactList from '../components/ContactList'
import contacts from '../data_contacts/contacts'

export default {
data(){
  return{
    contacts
  }
},
components:{
  ContactList
}
};
</script>

<style>
</style>