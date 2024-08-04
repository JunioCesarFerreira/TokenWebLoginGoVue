<template>
  <div>
    <h2>Protected Page</h2>
    <div v-if="isAuthenticated">
      <p>This page is protected and requires authentication.</p>
      <p>Welcome to the protected area!</p>
      <button @click="fetchMessage">Fetch Message</button>
      <p v-if="apiMessage">{{ apiMessage }}</p>
    </div>
    <div v-else>
      <p>You are not authenticated. Please log in to access this page.</p>
    </div>
  </div>
</template>

<script>

const API_URL = 'http://localhost:8082';
const PROTECTED_ENDPOINT = '/protected';

export default {
  name: 'ProtectedPage',
  data() {
    return {
      isAuthenticated: false,
      apiMessage: ''
    };
  },
  mounted() {
    this.checkAuthentication();
  },
  methods: {
    checkAuthentication() {
      const token = sessionStorage.getItem('authToken');
      console.log("token local storage", token)
      this.isAuthenticated = (token !== "null");
      console.log("is auth = ", this.isAuthenticated)
    },
    async fetchMessage() {
      try {
        const token = sessionStorage.getItem('authToken');
        const response = await fetch(`${API_URL}${PROTECTED_ENDPOINT}`, {
          method: 'GET',
          headers: {
            'Authorization': `${token}`
          }
        });

        if (response.ok) {
          const data = await response.json();
          this.apiMessage = data.message;
        } else {
          this.apiMessage = 'Error fetching message';
        }
      } catch (error) {
        console.error('Error:', error);
        this.apiMessage = 'An error occurred';
      }
    }
  }
};
</script>
