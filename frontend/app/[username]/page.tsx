import { Container, Divider } from '@mantine/core';
import { ProfileHeader } from '@/app/[username]/_components/ProfileHeader/ProfileHeader';
import { ProfilePosts } from '@/app/[username]/_components/ProfilePosts/ProfilePosts';
import Loading from '@/components/Loading';

export default async function Page({ params }: { params: Promise<{ username: string }> }) {
  const { username } = await params;
  if (!username) {
    return <Loading />;
  }
  return (
    <Container px="md">
      <ProfileHeader username={username} />
      <Divider />
      <ProfilePosts />
    </Container>
  );
}
