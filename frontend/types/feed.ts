import { PaginationMetadata } from '@/types/api';
import { Post } from '@/types/post';

export interface GetFeedResponse {
  data: Post[];
  pagination: PaginationMetadata;
}
