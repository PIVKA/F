<template>
      <div class="chat-page container">
       <ChatList :chatUser="chatUser" class="list-group list-group-numbered"/>
      </div>
</template>

<script>
import ChatList from '../components/ChatList.vue'
import chatUser from '../data/chatUser';
export default {
    data(){
        return{
            chatUser
        }
    },
   components:{
    ChatList,
   }
}
</script>

<style>

</style>