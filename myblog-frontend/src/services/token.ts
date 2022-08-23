export class TokenService {
  static  saveToken(token: string){
    localStorage.setItem('token', token)
  }
  
  static getToken() {
    return localStorage.getItem('token')
  }
  
  static removeToken(){
    localStorage.removeItem('token')
  }
}