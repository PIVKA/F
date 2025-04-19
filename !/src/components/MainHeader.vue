<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary p-3">
    <div class="container-fluid">
      <router-link class="navbar-brand" to="/" href="#">SMTH</router-link>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarTogglerDemo02"
        aria-controls="navbarTogglerDemo02"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarTogglerDemo02">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link to="/" class="nav-link">Chat</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/contacts" class="nav-link">Contacts</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/dialog" >Dialog</router-link>
          </li>
        </ul>
        <form class="d-flex" role="search">
          <input
            class="form-control me-2"
            type="search"
            placeholder="Search"
            aria-label="Search"
          />
          <button
            class="btn btn-warning"
            style="color: black !important"
            type="submit"
          >
            Search
          </button>
        </form>
      </div>
    </div>
  </nav>
</template>

<script>
export default {};
</script>

<style lang="scss">
</style>