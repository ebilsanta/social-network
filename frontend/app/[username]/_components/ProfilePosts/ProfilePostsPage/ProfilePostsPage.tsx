import { useEffect } from 'react';
import useSWR from 'swr';
import { AspectRatio, Card, Image, SimpleGrid } from '@mantine/core';
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
    <SimpleGrid cols={{ base: 1, xs: 2, md: 3 }}>
      {data.data.map((post) => (
        <Card key={post.id} radius="sm" component="a" href="#" m={0} p={0}>
          <AspectRatio ratio={1}>
            <Image src={post.image} />
          </AspectRatio>
        </Card>
      ))}
    </SimpleGrid>
  );
};
