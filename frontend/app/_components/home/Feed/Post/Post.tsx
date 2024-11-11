import { AspectRatio, Card, Image, Text } from '@mantine/core';
import { Post as APIPost } from '@/types/post';
import classes from './Post.module.css';

export const Post = ({ post }: { post: APIPost }) => {
  const { createdAt, caption, image } = post;
  return (
    <Card p="md" radius="md" component="a" href="#" className={classes.card}>
      <AspectRatio ratio={1920 / 1080}>
        <Image src={image} />
      </AspectRatio>
      <Text c="dimmed" size="xs" tt="uppercase" fw={700} mt="md">
        {createdAt && createdAt.seconds}
      </Text>
      <Text className={classes.title} mt={5}>
        {caption}
      </Text>
    </Card>
  );
};
