export const API_URL = 'http://localhost:8082';
export const LOGIN_ENDPOINT = '/login';
export const PROTECTED_ENDPOINT = '/protected';
export const RENEW_ENDPOINT = '/renew';

export const SUCCESS_MESSAGE = 'Login successful';
export const FAILURE_MESSAGE = 'Login failed';
export const ERROR_MESSAGE = 'An error occurred';

export const INACTIVITY_LIMIT = 2 * 60 * 1000; // 2 minutos sem atividade
export const RENEW_LIMIT = 50 * 1000; // Renova token a cada 50 segundos, pois os tokens expiram em 1 minuto