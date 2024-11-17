import { Container } from '@mantine/core';
import { FeedPage } from '@/app/_components/Feed/FeedPage/FeedPage';
import { useFeed } from '@/app/_components/Feed/useFeed';

const LIMIT = 4;

export const Feed = () => {
  const { user, page, setMorePages, loadMoreRef } = useFeed();
  const feedPages = [];
  if (user) {
    for (let i = 1; i <= page; i += 1) {
      feedPages.push(
        <FeedPage key={i} userId={user!.id} index={i} limit={LIMIT} setMorePages={setMorePages} />
      );
    }
  }

  return (
    <Container py="lg">
      {feedPages} <div ref={loadMoreRef} style={{ height: '1px', visibility: 'hidden' }} />
    </Container>
  );
};
