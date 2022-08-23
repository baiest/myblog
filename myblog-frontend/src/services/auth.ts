import { TokenService } from './token';
import { Auth, AuthResponse } from './../models/Auth';
import axios from "./axios"

const apiUrl = '/auth'

export class AuthService {
  
  static async login(data: Auth): Promise<Error | null>{
    try {
      const response = await axios.post<AuthResponse>(`http://localhost:5000/api${apiUrl}/login`, data)
      TokenService.saveToken(response.data.token)
      return null
    } catch (error: any) {
      throw error
    }
  }
}