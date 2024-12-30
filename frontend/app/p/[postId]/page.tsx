import { SinglePost } from '@/app/p/[postId]/_components/SinglePost';

export default async function Page({ params }: { params: Promise<{ postId: string }> }) {
  const { postId } = await params;
  return <SinglePost postId={postId} />;
}
