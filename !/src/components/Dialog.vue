<template>
  <div class="d-flex justify-content-between">
    <div class="dialog-window">
      <div class="dialog-message"></div>
      <div class="form-floating d-flex justify-content-between">
        <textarea
          v-model="message"
          class="form-control"
          placeholder="Введите сообщение..."
          id="floatingTextarea2"
          style="height: 100px; width: 850px"
        ></textarea>
        <div class="btn-group">
          <button @click="sendMessage" class="btn btn-outline-light">Send</button>
        </div>
      </div>
    </div>
    <div class="dialog-board">
      <li @click="addEmoji" class="emoji" v-for="smile in emoji" :key="smile.slug">
        {{ smile.character }}
      </li>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: "",
      emoji: [],
    };
  },
  // https://emoji-api.com/emojis?access_key=31abf45e925d0c113e328c0037644b43b2d0f071
  mounted() {
    const axios = require("axios");

    axios
      .get(
        "https://emoji-api.com/emojis?access_key=31abf45e925d0c113e328c0037644b43b2d0f071"
      )
      .then((res) => {
        return res;
      })
      .then((res) => {
        console.log(res.data);
        this.emoji = res.data;
      })
      .catch((err) => {
        console.log(err);
      });
  },
  methods: {
    sendMessage() {
      const dialogMessage = document.querySelector(".dialog-message");

      dialogMessage.innerHTML += `<div class="message">${this.message}</div>`;
    },
    addEmoji(){
      this.message += `${event.target.innerHTML}`
    }
  },
};
</script>

<style>
</style>