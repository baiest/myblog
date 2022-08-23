import { Post, PostDAO } from './../models/Post';
import axios from './axios'

const apiUrl = '/posts'

export class PostService{
  static async create(data: PostDAO): Promise<Post>{
      try {
        return (await axios.post<Post>(apiUrl, data)).data
      } catch (error: any) {
        throw new Error(error.response.data.error)
      }
  }
  
  static async getPosts(): Promise<Post[]>{
    return (await axios.get<Post[]>(apiUrl)).data
      .map((p: Post) => ({
        ...p,
        created_at: new Date(p.created_at)
      }))
  }
}
