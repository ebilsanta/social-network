'use client';

import { useSession, signIn, signOut } from 'next-auth/react';

export default function Component() {
  const { data: session } = useSession();
  if (session) {
    console.log(session.user);
  }
  if (session) {
    return (
      <>
        Signed in as {session.user.email} <br />
        <button type="submit" onClick={() => signOut()}>Sign out</button>
      </>
    );
  }
  return (
    <>
      Not signed in <br />
      <button type="submit" onClick={() => signIn()}>Sign in</button>
    </>
  );
}
