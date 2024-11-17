'use client';

import { Container } from '@mantine/core';
import { ProfilePostsPage } from '@/app/[username]/_components/ProfilePosts/ProfilePostsPage/ProfilePostsPage';
import { useProfilePosts } from '@/app/[username]/_components/ProfilePosts/useProfilePosts';

const LIMIT = 12;

export const ProfilePosts = () => {
  const { user, page, setMorePages, loadMoreRef } = useProfilePosts();
  const postPages = [];
  if (user) {
    for (let i = 1; i <= page; i += 1) {
      postPages.push(
        <ProfilePostsPage
          key={i}
          userId={user!.id}
          index={i}
          limit={LIMIT}
          setMorePages={setMorePages}
        />
      );
    }
  }
  return (
    <Container my="xl">
      {postPages} <div ref={loadMoreRef} style={{ height: '1px', visibility: 'hidden' }} />
    </Container>
  );
};
