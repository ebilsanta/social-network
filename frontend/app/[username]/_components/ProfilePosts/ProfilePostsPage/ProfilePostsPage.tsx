import { useEffect } from 'react';
import useSWR from 'swr';
import { AspectRatio, Grid, Image } from '@mantine/core';
import { PostAPI } from '@/lib/post-api';

interface ProfilePostsPageProps {
  userId: string;
  index: number;
  limit: number;
  setMorePages: (value: boolean) => void;
}

export const ProfilePostsPage = ({ userId, index, limit, setMorePages }: ProfilePostsPageProps) => {
  const fetchProfilePosts = async () => {
    const response = await PostAPI.getPostsByUserId(userId, index, limit);
    return response;
  };
  const { data, isLoading } = useSWR(
    `/api/posts/user/${userId}?page=${index}&limit=${limit}`,
    fetchProfilePosts
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
    <Grid gutter="md">
      {data.data.map((post) => (
        <Grid.Col key={post.id} span={{ base: 12, md: 4 }}>
          <AspectRatio ratio={1}>
            <Image src={post.image} />
          </AspectRatio>
        </Grid.Col>
      ))}
    </Grid>
  );
};
