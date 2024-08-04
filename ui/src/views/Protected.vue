<template>
  <div>
    <h2>Protected Page</h2>
    <div v-if="isAuthenticated">
      <p>This page is protected and requires authentication.</p>
      <p>Welcome to the protected area!</p>
      <button @click="fetchMessage">Fetch Message</button>
      <p v-if="apiMessage">{{ apiMessage }}</p>
      <p>Time until logout: {{ inactivityTimeLeft }} seconds</p>
    </div>
    <div v-else>
      <p>You are not authenticated. Please log in to access this page.</p>
      <button @click="goToLogin">Go to Login</button>
    </div>
  </div>
</template>

<script>
import { 
  API_URL, 
  PROTECTED_ENDPOINT, 
  INACTIVITY_LIMIT, 
  RENEW_ENDPOINT, 
  RENEW_LIMIT 
} from '@/config.js';

export default {
  name: 'ProtectedPage',
  data() {
    return {
      isAuthenticated: false,
      apiMessage: '',
      lastActivity: Date.now(),
      inactivityTimeLeft: INACTIVITY_LIMIT / 1000
    };
  },

  mounted() {
    this.checkAuthentication(); // Verifica se está autenticado
    this.startInactivityTimer(); // Inicia temporizador de inatividade
    this.startRenewalTimer(); // Inicia temporizador de renovação de token
  },
  
  unmounted() {
    clearInterval(this.inactivityTimer);
    clearInterval(this.renewalTimer);
    document.removeEventListener('mousemove', this.resetActivity);
    document.removeEventListener('keydown', this.resetActivity);
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
          this.isAuthenticated = false;
        }
      } catch (error) {
        console.error('Error:', error);
        this.apiMessage = 'An error occurred';
      }
    },

    async renewToken() {
      try {
        const token = sessionStorage.getItem('authToken');
        const response = await fetch(`${API_URL}${RENEW_ENDPOINT}`, {
          method: 'GET',
          headers: {
            'Authorization': `${token}`
          }
        });

        if (response.ok) {
          const data = await response.json();
          sessionStorage.setItem('authToken', data.token);
        }
      } catch (error) {
        console.error('Error renewing token:', error);
        this.logout();
      }
    },

    goToLogin() {
      this.$router.push('/login');
    },

    startInactivityTimer() {
      document.addEventListener('mousemove', this.resetActivity);
      document.addEventListener('keydown', this.resetActivity);

      this.inactivityTimer = setInterval(() => {
        const timeSinceLastActivity = Date.now() - this.lastActivity;
        this.inactivityTimeLeft = Math.max(0, Math.floor((INACTIVITY_LIMIT - timeSinceLastActivity) / 1000));

        if (timeSinceLastActivity > INACTIVITY_LIMIT) {
          this.logout();
        }
      }, 1000);
    },

    startRenewalTimer() {
      this.renewalTimer = setInterval(() => {
        this.renewToken();
      }, RENEW_LIMIT);
    },

    resetActivity() {
      this.lastActivity = Date.now();
    },

    logout() {
      sessionStorage.setItem('authToken', null);
      this.$router.push('/login');
    }
  }
};
</script>
