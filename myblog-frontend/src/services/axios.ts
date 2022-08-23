import { TokenService } from './token';
import axios from 'axios';

const instance = axios.create({
  baseURL: "http://localhost:5000/api",
  headers: {
    'Authorization': `Bearer ${TokenService.getToken()}`
  }
})

export default instance