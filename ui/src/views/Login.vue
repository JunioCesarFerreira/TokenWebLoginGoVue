<template>
  <div class="login-container">
    <h2 class="login-title">Login</h2>
    <form @submit.prevent="login" class="form-container">
      <div class="input-group">
        <label for="userId">ID: </label>
        <input type="text" id="userId" v-model="userid" class="input-field" @input="validateInput" />
      </div>
      <div class="input-group">
        <label for="password">Pwd:</label>
        <input type="password" id="password" v-model="hash" class="input-field" />
      </div>
      <button type="submit" class="login-button">Login</button>
    </form>
    <p v-if="message" class="error-message">{{ message }}</p>
  </div>
</template>

<script>
import CryptoJS from 'crypto-js';
import { 
  API_URL, 
  LOGIN_ENDPOINT, 
  SUCCESS_MESSAGE, 
  FAILURE_MESSAGE, 
  ERROR_MESSAGE 
} from '@/config.js';

export default {
  name: 'LoginForm',
  data() {
    return {
      userid: '',
      hash: '',
      message: ''
    };
  },
  methods: {
    validateInput(event) {
      const value = event.target.value;
      if (value.includes('<') || value.includes('>')) {
        event.target.value = value.replace(/[<>]/g, '');
        this.message = 'Invalid characters removed';
      }
    },
    async login() {
      try {
        const hashedPwd = CryptoJS.SHA512(this.hash).toString(CryptoJS.enc.Hex);
        const response = await this.makeLoginRequest(this.userid, hashedPwd);

        if (response.ok) {
          const data = await response.json();
          console.log(SUCCESS_MESSAGE);
          this.handleSuccess(data.token);
        } else {
          console.log(FAILURE_MESSAGE);
          this.handleFailure();
        }
      } catch (error) {
        console.error('Error:', error);
        this.handleError();
      }
    },
    async makeLoginRequest(userId, hashedPwd) {
      return await fetch(`${API_URL}${LOGIN_ENDPOINT}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          userId,
          pass: hashedPwd
        })
      });
    },
    handleSuccess(token) {
      this.message = SUCCESS_MESSAGE;
      sessionStorage.setItem('authToken', token);
      this.$router.push('/protected');
    },
    handleFailure() {
      sessionStorage.setItem('authToken', null);
      this.message = FAILURE_MESSAGE;
    },
    handleError() {
      this.message = ERROR_MESSAGE;
    }
  }
};
</script>

<style>
.login-container {
  text-align: center;
  padding: 20px;
}
.login-title {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}
.form-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #f8f8f8;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
}
.input-group {
  margin: 10px 0;
  display: flex;
  align-items: center;
}
label {
  display: inline-block;
  width: 80px;
  font-weight: bold;
}
.input-field {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  width: 100%;
}
.login-button {
  padding: 10px 20px;
  font-size: 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}
.login-button:hover {
  background-color: #0056b3;
}
.error-message {
  color: red;
  margin-top: 10px;
  font-size: 14px;
}
</style>

