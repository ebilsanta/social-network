'use client';

import { AspectRatio, Avatar, Card, Flex, Image, Text } from '@mantine/core';
import { Post as APIPost } from '@/types/post';
import classes from './Post.module.css';

export const Post = ({ post }: { post: APIPost }) => {
  const { createdAt, caption, image, user } = post;
  const handlePostClick = (id: string) => {
    window.history.pushState(null, '', `/p/${id}`);
  };
  return (
    <Card
      p="md"
      radius="md"
      component="a"
      className={classes.card}
      onClick={() => handlePostClick(post?.id)}
    >
      <Flex gap={10} mb={10} align="center">
        <Avatar src={user.image} radius="xl" />
        <Text fw={600}>{user.username}</Text>
      </Flex>
      <AspectRatio ratio={1920 / 1080}>
        <Image src={image} />
      </AspectRatio>
      <Text c="dimmed" size="xs" tt="uppercase" fw={700} mt="md">
        {createdAt && new Date(createdAt.seconds * 1000).toLocaleDateString()}
      </Text>
      <Text className={classes.title} mt={5}>
        {caption}
      </Text>
    </Card>
  );
};
