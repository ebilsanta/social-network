import { useEffect } from 'react';
import useSWR from 'swr';
import { SimpleGrid } from '@mantine/core';
import { Post } from '@/app/_components/Feed/Post/Post';
import { FeedAPI } from '@/lib/feed-api';

interface FeedPageProps {
  userId: string;
  index: number;
  limit: number;
  setMorePages: (value: boolean) => void;
}

export const FeedPage = ({ userId, index, limit, setMorePages }: FeedPageProps) => {
  const fetchFeedPage = async () => {
    const response = await FeedAPI.getFeed(userId, index, limit);
    return response;
  };
  const { data, isLoading } = useSWR(
    `/api/feeds/${userId}?page=${index}&limit=${limit}`,
    fetchFeedPage
  );
  useEffect(() => {
    if (data && !data.pagination.nextPage) {
      setMorePages(false);
    }
  }, [data]);

  if (!data || isLoading) {
    return null;
  }

  return (
    <SimpleGrid cols={{ base: 1 }} mt="lg">
      {data.data.map((post) => (
        <Post key={post.id} post={post} />
      ))}
    </SimpleGrid>
  );
};
