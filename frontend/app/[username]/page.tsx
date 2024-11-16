import { ProfileCard } from '@/app/[username]/ProfileCard/ProfileCard';
import Loading from '@/components/loading';

export default async function Page({ params }: { params: Promise<{ username: string }> }) {
  const { username } = await params;
  if (!username) {
    return <Loading />;
  }
  return <ProfileCard username={username} />;
}
