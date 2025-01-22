import { Container } from '@mantine/core';
import { FeedPage } from '@/app/_components/Feed/FeedPage/FeedPage';
import { useFeed } from '@/app/_components/Feed/useFeed';
import PostModal from '@/components/PostModal/PostModal';

const LIMIT = 3;

export const Feed = () => {
  const { user, page, setMorePages, loadMoreRef, postId } = useFeed();
  const feedPages = [];
  if (user) {
    for (let i = 1; i <= page; i += 1) {
      feedPages.push(
        <FeedPage key={i} userId={user!.id} index={i} limit={LIMIT} setMorePages={setMorePages} />
      );
    }
  }

  return (
    <Container pb="lg" size="xs">
      {postId && <PostModal postId={postId} />}
      {feedPages} <div ref={loadMoreRef} style={{ height: '1px', visibility: 'hidden' }} />
    </Container>
  );
};
