<template>
  <div class="session-expired">
    <h1 class="session-expired__title">Session Expired</h1>
    <p class="session-expired__message">
      Your session has expired. Please go back to the
      <router-link class="session-expired__link" to="/">Product List</router-link>.
    </p>
  </div>
</template>

<script>
export default {
  name: 'SessionExpired'
};
</script>

<style lang="scss" scoped>
@use '../styles/main.scss' as *;

.session-expired {
  padding: 2rem;
  text-align: center;
  color: #ff4d4f;
}

.session-expired__title {
  margin-bottom: 1rem;
}

.session-expired__message {
  font-size: 1rem;
}

.session-expired__link {
  color: $primary-color;
  text-decoration: none;
}
</style>
