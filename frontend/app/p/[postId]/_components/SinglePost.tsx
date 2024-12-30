'use client';

import useSWR from 'swr';
import { AspectRatio, Avatar, Card, Center, Flex, Image, Text } from '@mantine/core';
import Loading from '@/components/Loading';
import NotFound from '@/components/NotFound';
import Unauthorized from '@/components/Unauthorized';
import { UnauthorizedError } from '@/lib/errors';
import { PostAPI } from '@/lib/post-api';

export const SinglePost = ({ postId }: { postId: string }) => {
  const fetchPost = async () => {
    const response = await PostAPI.getPostById(postId);
    return response.data;
  };

  const { data: post, error, isLoading } = useSWR(`/api/posts/${postId}`, fetchPost);
  if (isLoading) {
    return <Loading />;
  }
  if (error) {
    if (error instanceof UnauthorizedError) {
      return <Unauthorized />;
    }
    return <NotFound />;
  }
  if (!post) {
    return <NotFound />;
  }

  return (
    <Center mt="xl">
      <Card p="md" radius="md" component="a">
        <Flex gap={10} mb={10} align="center">
          <Avatar src={post?.user.image} radius="xl" />
          <Text fw={600}>{post?.user.username}</Text>
        </Flex>
        <AspectRatio ratio={1920 / 1080}>
          <Image src={post?.image} />
        </AspectRatio>
        <Text c="dimmed" size="xs" tt="uppercase" fw={700} mt="md">
          {post?.createdAt && new Date(post.createdAt.seconds * 1000).toLocaleDateString()}
        </Text>
        <Text mt={5}>{post?.caption}</Text>
      </Card>
    </Center>
  );
};
