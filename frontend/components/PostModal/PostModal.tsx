import { useRouter } from 'next/navigation';
import useSWR from 'swr';
import { AspectRatio, Card, Image, Modal, Notification, Text } from '@mantine/core';
import { PostAPI } from '@/lib/post-api';

const PostModal = ({ postId }: { postId: string }) => {
  const fetchPost = async () => {
    const response = await PostAPI.getPostById(postId);
    return response.data;
  };
  const { data: post, error } = useSWR(`/api/posts/${postId}`, fetchPost);
  const router = useRouter();

  const handleClose = () => {
    router.push('/');
  };

  if (error) {
    return (
      <Notification color="red" title="Failed to load post">
        Please try again. If the error persists, please contact thaddeusleezx@gmail.com
      </Notification>
    );
  }

  return (
    <Modal opened centered onClose={handleClose} size="1000px">
      <Card p="md" radius="md" onClick={() => {}}>
        <AspectRatio ratio={1920 / 1080}>
          <Image src={post?.image} />
        </AspectRatio>
        <Text c="dimmed" size="xs" tt="uppercase" fw={700} mt="md">
          {post?.createdAt && new Date(post.createdAt.seconds * 1000).toLocaleDateString()}
        </Text>
        <Text mt={5}>{post?.caption}</Text>
      </Card>
    </Modal>
  );
};

export default PostModal;
