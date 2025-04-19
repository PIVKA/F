<template>
 <div class="chat-page container">
    <Chat
      v-for="chat in chatUser"
      :key="chat.id" 
      :chat="chat"
   />
 </div>
</template>

<script>
import Chat from '../components/Chat';

export default {
    props:{
        chatUser:{
            type: Array,
            default:() => [],
        }
    },
    components:{
        Chat
    }
}
</script>

<style>

</style>