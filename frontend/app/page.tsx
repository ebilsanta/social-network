'use client';

import { useSession } from 'next-auth/react';
import Loading from '@/components/loading';
import { Feed } from './_components/Feed/Feed';
import { Welcome } from './_components/Welcome/Welcome';

export default function Page() {
  const { data: session, status } = useSession();
  if (status === 'loading') {
    return <Loading />;
  }
  if (session) {
    return <Feed />;
  }
  return (
    <>
      <Welcome />
    </>
  );
}
