export interface Post {
  id: string
  title: string
  content: string
  id_user: number
  created_at: Date
}

export interface PostDAO {
  title: string
  content: string
}