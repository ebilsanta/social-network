'use client';

import { ReactNode } from 'react';
import { Navbar } from '@/app/_components/Navbar/Navbar';
import { useUser } from '@/providers/user-provider';
import classes from './Layout.module.css';

export const Layout = ({ children }: { children: ReactNode }) => {
  const { user } = useUser();

  return (
    <div className={`${user ? classes.containerNavbar : classes.container}`}>
      {user && <Navbar user={user} />}
      <div>{children}</div>
    </div>
  );
};
