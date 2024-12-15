import { PaginationMetadata } from '@/types/api';
import { User } from '@/types/user';

export interface Post {
  id: string;
  caption: string;
  user: User;
  image: string;
  createdAt: {
    seconds: number;
    nanos: number;
  };
}

export interface GetPostByIdResponse {
  data: Post;
}

export interface GetPostsByUserIdResponse {
  data: Post[];
  pagination: PaginationMetadata;
}
