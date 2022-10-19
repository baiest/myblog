import { User } from './User';
export interface Post {
  id: number
  title: string
  content: string
  user: User
  created_at: Date
}

export interface PostDAO {
  title: string
  content: string
}