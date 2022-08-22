import { Post } from './../models/Post';
import axios from 'axios'

const apiUrl = '/users'

export class PostService{
  static async getPosts(): Promise<Post[]>{
    return (await axios.get(`http://localhost:5000/api${apiUrl}`)).data
  }
}
