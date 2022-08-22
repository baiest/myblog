export interface Request<T> {
  loading: boolean
  data: T
  error: string | null
} 