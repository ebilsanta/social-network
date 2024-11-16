import { ProfileHeader } from '@/app/[username]/_components/ProfileHeader/ProfileHeader';
import Loading from '@/components/loading';

export default async function Page({ params }: { params: Promise<{ username: string }> }) {
  const { username } = await params;
  if (!username) {
    return <Loading />;
  }
  return <ProfileHeader username={username} />;
}
