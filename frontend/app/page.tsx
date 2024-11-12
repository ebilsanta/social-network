'use client';

import { useSession } from 'next-auth/react';
import { Loader } from '@mantine/core';
import { Feed } from './_components/Feed/Feed';
import { Welcome } from './_components/Welcome/Welcome';

export default function Page() {
  const { data: session, status } = useSession();
  if (status === 'loading') {
    return <Loader />;
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
