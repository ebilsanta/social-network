'use client';

import { useSession } from 'next-auth/react';
import { Loader } from '@mantine/core';
import { Home } from './_components/home/home';
import { Welcome } from './_components/welcome/welcome';

export default function HomePage() {
  const { data: session, status } = useSession();
  if (status === 'loading') {
    return <Loader />;
  }
  if (session) {
    return <Home />;
  }
  return (
    <>
      <Welcome />
    </>
  );
}
