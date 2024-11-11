import { PaginationMetadata } from '@/types/api';

export interface CreateUserRequest {
  email: string;
  image: string;
  username: string;
  name: string;
  id: string;
}

export interface User {
  id: string;
  email: string;
  name: string;
  username: string;
  postCount: number;
  followerCount: number;
  followingCount: number;
  image: string;
  createdAt: {
    seconds: number;
    nanos: number;
  };
}

export interface GetUsersResponse {
  data: User[];
  pagination: PaginationMetadata;
}
