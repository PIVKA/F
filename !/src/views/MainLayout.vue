<template>
    <div class="container">
        <main-header/>
        <router-view/>
    </div>
  </template>
  
  <script>
  import MainHeader from '../components/MainHeader'
  export default {
      components: {
          MainHeader,
      }
  }
  </script>
  
  <style>
  
  </style>