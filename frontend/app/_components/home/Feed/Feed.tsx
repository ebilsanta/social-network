import { Container } from '@mantine/core';
import { FeedPage } from '@/app/_components/Home/Feed/FeedPage/FeedPage';
import { useFeed } from '@/app/_components/Home/Feed/useFeed';

const LIMIT = 4;

export const Feed = () => {
  const { user, page, loadMoreRef } = useFeed();
  const feedPages = [];
  if (user) {
    for (let i = 1; i <= page; i += 1) {
      feedPages.push(<FeedPage key={i} userId={user!.id} index={i} limit={LIMIT} />);
    }
  }

  return (
    <Container py="xl">
      {feedPages} <div ref={loadMoreRef} style={{ height: '1px', visibility: 'hidden' }}></div>
    </Container>
  );
};
