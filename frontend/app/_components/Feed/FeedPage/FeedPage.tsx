import useSWR from 'swr';
import { SimpleGrid } from '@mantine/core';
import { Post } from '@/app/_components/Feed/Post/Post';
import { FeedAPI } from '@/lib/feed-api';

interface FeedPageProps {
  userId: string;
  index: number;
  limit: number;
}

export const FeedPage = ({ userId, index, limit }: FeedPageProps) => {
  const fetchFeedPage = async () => {
    const response = await FeedAPI.getFeed(userId, index, limit);
    return response.data;
  };
  const { data } = useSWR(`/api/feeds/${userId}?page=${index}&limit=${limit}`, fetchFeedPage);
  return (
    <SimpleGrid cols={{ base: 1 }}>
      {data && data.map((post) => <Post key={post.id} post={post} />)}
    </SimpleGrid>
  );
};
