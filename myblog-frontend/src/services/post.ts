import { Post, PostDAO } from './../models/Post';
import axios from './axios'

const apiUrl = '/posts'

const responseToPost = (p: Post) => ({
  ...p,
  created_at: new Date(p.created_at)
})
export class PostService{
  static async create(data: PostDAO): Promise<Post>{
    try {
      const post = (await axios.post<Post>(apiUrl, data)).data
      return responseToPost(post)
    } catch (error: any) {
      throw new Error(error.response.data.error)
    }
  }
  
  static async getPosts(): Promise<Post[]>{
    return (await axios.get<Post[]>(apiUrl)).data
      .map(responseToPost)
  }
  
  static async getById(id: number): Promise<Post>{
    try {
      const post = (await axios.get<Post>(`${apiUrl}/${id}`)).data
      return responseToPost(post)
    } catch (error: any) {
      throw new Error(error.response.data.error)
    }
  }
  static async delete(id: number) {
    try {
      await axios.delete(`${apiUrl}/${id}`)
    } catch (error: any) {
      throw new Error(error.response.data.error)
    }
  }
}
