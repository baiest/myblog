import { Post } from './../models/Post';
import axios from './axios'

const apiUrl = '/posts'

export class PostService{
  static async getPosts(): Promise<Post[]>{
    return (await axios.get<Post[]>(apiUrl)).data
      .map((p: Post) => ({
        ...p,
        created_at: new Date(p.created_at)
      }))
  }
}
