import { PaginationMetadata } from '@/types/api';

export interface Post {
  id: string;
  caption: string;
  userId: string;
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
